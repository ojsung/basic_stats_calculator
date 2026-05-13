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

var formTmpl = template.Must(template.ParseFS(templateFS, "templates/base.html", "templates/form.html"))

var cdfTmpl = template.Must(template.ParseFS(templateFS, "templates/base.html", "templates/cdf.html"))

var pvalueTmpl = template.Must(template.ParseFS(templateFS, "templates/base.html", "templates/pvalue.html"))

type formData struct {
	P, N, K   string
	Error     string
	Result    string
	Pct       string
	ActiveTab string
}

type termRow struct {
	K           string
	Probability string
	Cumulative  string
}

type cdfData struct {
	P, N, K       string
	Error         string
	Cumulative    string
	CumulativePct string
	Terms         []termRow
	ActiveTab     string
}

type pvalueData struct {
	P, N, K   string
	Tail      string
	Error     string
	PValue    string
	PValuePct string
	ActiveTab string
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	formTmpl.Execute(w, formData{ActiveTab: "binomial"}) //nolint:errcheck
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		formTmpl.Execute(w, formData{Error: "Could not parse form.", ActiveTab: "binomial"}) //nolint:errcheck
		return
	}
	d := formData{P: r.FormValue("p"), N: r.FormValue("n"), K: r.FormValue("k"), ActiveTab: "binomial"}
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
	pctFloat := bu.PrecFloat().Mul(&result, big.NewFloat(100))
	d.Result = bu.ToStr(&result, 10)
	d.Pct = bu.ToStr(pctFloat, 4)
	formTmpl.Execute(w, d) //nolint:errcheck
}

func cdfFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	cdfTmpl.Execute(w, cdfData{ActiveTab: "cdf"}) //nolint:errcheck
}

func cdfCalculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		cdfTmpl.Execute(w, cdfData{Error: "Could not parse form.", ActiveTab: "cdf"}) //nolint:errcheck
		return
	}
	d := cdfData{P: r.FormValue("p"), N: r.FormValue("n"), K: r.FormValue("k"), ActiveTab: "cdf"}
	p, ok := new(big.Float).SetString(d.P)
	if !ok {
		d.Error = "Invalid value for p — must be a decimal number between 0 and 1."
		cdfTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	n, err := strconv.ParseInt(d.N, 10, 64)
	if err != nil {
		d.Error = "Invalid value for n — must be a whole number."
		cdfTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	k, err := strconv.ParseInt(d.K, 10, 64)
	if err != nil {
		d.Error = "Invalid value for k — must be a whole number."
		cdfTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	cumulative, terms, calcErr := calculator.CumulativeBinomialProbability(p, n, k)
	if calcErr != nil {
		d.Error = calcErr.Error()
		cdfTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	pctFloat := new(big.Float).Mul(&cumulative, big.NewFloat(100))
	d.Cumulative = bu.ToStr(&cumulative, 10)
	d.CumulativePct = bu.ToStr(pctFloat, 4)
	runningSum := bu.PrecFloat().SetInt64(0)
	d.Terms = make([]termRow, len(terms))
	for i, term := range terms {
		prob := new(big.Float).Copy(&term)
		runningSum.Add(runningSum, prob)
		d.Terms[i] = termRow{
			K:           strconv.FormatInt(int64(i), 10),
			Probability: bu.ToStr(prob, 6),
			Cumulative:  bu.ToStr(runningSum, 6),
		}
	}
	cdfTmpl.Execute(w, d) //nolint:errcheck
}

func pvalueFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pvalueTmpl.Execute(w, pvalueData{ActiveTab: "pvalue"}) //nolint:errcheck
}

func pvalueCalculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		pvalueTmpl.Execute(w, pvalueData{Error: "Could not parse form.", ActiveTab: "pvalue"}) //nolint:errcheck
		return
	}
	d := pvalueData{
		P: r.FormValue("p"), N: r.FormValue("n"), K: r.FormValue("k"),
		Tail: r.FormValue("tail"), ActiveTab: "pvalue",
	}
	if d.Tail == "" {
		d.Tail = "two"
	}
	p, ok := new(big.Float).SetString(d.P)
	if !ok {
		d.Error = "Invalid value for p — must be a decimal number between 0 and 1."
		pvalueTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	n, err := strconv.ParseInt(d.N, 10, 64)
	if err != nil {
		d.Error = "Invalid value for n — must be a whole number."
		pvalueTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	k, err := strconv.ParseInt(d.K, 10, 64)
	if err != nil {
		d.Error = "Invalid value for k — must be a whole number."
		pvalueTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	pval, calcErr := calculator.BinomialPValue(p, n, k, d.Tail)
	if calcErr != nil {
		d.Error = calcErr.Error()
		pvalueTmpl.Execute(w, d) //nolint:errcheck
		return
	}
	pctFloat := bu.PrecFloat().Mul(&pval, big.NewFloat(100))
	d.PValue = bu.ToStr(&pval, 10)
	d.PValuePct = bu.ToStr(pctFloat, 4)
	pvalueTmpl.Execute(w, d) //nolint:errcheck
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
	http.HandleFunc("/cdf", cdfFormHandler)
	http.HandleFunc("/cdf/calculate", cdfCalculateHandler)
	http.HandleFunc("/pvalue", pvalueFormHandler)
	http.HandleFunc("/pvalue/calculate", pvalueCalculateHandler)
	fmt.Printf("Listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
