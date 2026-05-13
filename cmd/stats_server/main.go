package main

import (
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
	"github.com/ojsung/basic_stats_calculator/pkg/calculator"
)

const formHTML = `<!DOCTYPE html>
<html>
<head><title>Binomial Probability Calculator</title></head>
<body>
<h1>Binomial Probability Calculator</h1>
%s
<form method="POST" action="/calculate">
  <label>p (chance of success, 0-1):<br><input type="text" name="p"></label><br><br>
  <label>n (number of trials):<br><input type="text" name="n"></label><br><br>
  <label>k (number of successes):<br><input type="text" name="k"></label><br><br>
  <input type="submit" value="Calculate">
</form>
</body>
</html>`

const resultHTML = `<!DOCTYPE html>
<html>
<head><title>Result</title></head>
<body>
<h1>Result</h1>
<p>P(X = %s) with p=%s, n=%s: <strong>%s</strong> (%s%%)</p>
<a href="/">Calculate again</a>
</body>
</html>`

const calcErrorHTML = `<!DOCTYPE html>
<html>
<head><title>Error</title></head>
<body>
<h1>Error</h1>
<p>%s</p>
<a href="/">Try again</a>
</body>
</html>`

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, formHTML, "")
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, formHTML, "<p>Could not parse form.</p>")
		return
	}
	p, ok := new(big.Float).SetString(r.FormValue("p"))
	if !ok {
		fmt.Fprintf(w, formHTML, "<p>Invalid value for p — must be a decimal number between 0 and 1.</p>")
		return
	}
	n, err := strconv.ParseInt(r.FormValue("n"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, formHTML, "<p>Invalid value for n — must be a whole number.</p>")
		return
	}
	k, err := strconv.ParseInt(r.FormValue("k"), 10, 64)
	if err != nil {
		fmt.Fprintf(w, formHTML, "<p>Invalid value for k — must be a whole number.</p>")
		return
	}
	result, calcErr := calculator.CalculateBinomialProbability(p, n, k)
	if calcErr != nil {
		fmt.Fprintf(w, calcErrorHTML, calcErr.Error())
		return
	}
	resultStr := bu.ToStr(&result, 10)
	pctFloat := new(big.Float).Mul(&result, big.NewFloat(100))
	pctStr := bu.ToStr(pctFloat, 4)
	fmt.Fprintf(w, resultHTML, r.FormValue("k"), r.FormValue("p"), r.FormValue("n"), resultStr, pctStr)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Printf("Listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
