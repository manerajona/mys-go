package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type randomResult struct {
	Number int
	Parity string
}

type ViewData struct {
	Items []randomResult
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	n, _ := strconv.Atoi(req.FormValue("num"))
	// NOTE: Ignoring potential errors in convertion
	sequence := calculate(n)

	err := tpl.ExecuteTemplate(w, "index.gohtml", ViewData{sequence})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func calculate(count int) []randomResult {
	const b = 36
	sequence := []randomResult{}
	for ; count > 0; count-- {
		x := rand.Intn(b)
		par := parityOf(x)
		sequence = append(sequence, randomResult{x, par})
	}
	return sequence
}

func parityOf(v int) string {
	switch {
	case v == 0:
		return "Cero"
	case v%2 == 0:
		return "Par"
	default:
		return "Impar"
	}
}
