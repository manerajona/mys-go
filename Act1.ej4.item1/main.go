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
	sequence := calculate(n)

	err := tpl.ExecuteTemplate(w, "index.gohtml", ViewData{sequence})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

const a = 0
const b = 36

func calculate(count int) []randomResult {
	sequence := []randomResult{}
	for count > 0 {
		x := rand.Intn(b-a) + a
		par := getParity(x)
		rr := randomResult{x, par}
		sequence = append(sequence, rr)
		count--
	}
	return sequence
}

func getParity(v int) string {
	parity := "cero"
	if v != 0 {
		if v%2 == 0 {
			parity = "par"
		} else {
			parity = "impar"
		}
	}
	return parity
}
