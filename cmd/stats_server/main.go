package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"math/big"
	"net/http"
	"os"
	"strconv"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
	"github.com/ojsung/basic_stats_calculator/pkg/calculator"
)

//go:embed templates/*.html
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

var formTmpl = template.Must(template.ParseFS(templateFS, "templates/form.html"))

type formData struct {
	P, N, K string
	Error   string
	Result  string
	Pct     string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	formTmpl.Execute(w, formData{}) //nolint:errcheck
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		formTmpl.Execute(w, formData{Error: "Could not parse form."}) //nolint:errcheck
		return
	}
	d := formData{P: r.FormValue("p"), N: r.FormValue("n"), K: r.FormValue("k")}
	p, ok := new(big.Float).SetString(d.P)
	if !ok {
		d.Error = "Invalid value for p — must be a decimal number between 0 and 1."
		formTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	n, err := strconv.ParseInt(d.N, 10, 64)
	if err != nil {
		d.Error = "Invalid value for n — must be a whole number."
		formTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	k, err := strconv.ParseInt(d.K, 10, 64)
	if err != nil {
		d.Error = "Invalid value for k — must be a whole number."
		formTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	result, calcErr := calculator.CalculateBinomialProbability(p, n, k)
	if calcErr != nil {
		d.Error = calcErr.Error()
		formTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	pctFloat := new(big.Float).Mul(&result, big.NewFloat(100))
	d.Result = bu.ToStr(&result, 10)
	d.Pct = bu.ToStr(pctFloat, 4)
	formTmpl.Execute(w, d) //nolint:errcheck
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	staticSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSub))))
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Printf("Listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
