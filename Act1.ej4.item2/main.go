package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type ViewData struct {
	Items map[int]float64
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	dist := req.FormValue("dist")
	n, _ := strconv.Atoi(req.FormValue("num"))

	var sequence map[int]float64
	switch dist {
	case "uniform":
		sequence = calculateByUniform(n)
	case "exp":
		sequence = calculateByExp(n)
	case "empiric":
		sequence = calculateByEmpiric(n)
	case "poisson":
		sequence = calculateByPoisson(n)
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", ViewData{sequence})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
