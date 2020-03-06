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

type aleatoryNumber struct {
	Id        int
	Generated []float64
	Rnd       float64
}

type ViewData struct {
	Distribution, Method string
	Items                []aleatoryNumber
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	op := req.FormValue("op")
	n, _ := strconv.Atoi(req.FormValue("num"))
	// NOTE: Ignoring potential errors in convertion

	var (
		distribution, method string
		sequence             []aleatoryNumber
	)
	switch op {
	case "op1":
		// Dist. uniforme continua entre 13 y 17
		sequence = uniformeContinua(13, 17, n)
		distribution = "U[13;17]"
		method = "13 + (17-3) * rnd"
	case "op2":
		// Dist. uniforme discreta entre 20 y 25
		sequence = uniformeDiscreta(20, 25, n)
		distribution = "U[20;25]"
		method = "20 + (25-20) * rnd"
	case "op3":
		// Dist. exponencial con lambda**-1 = 10
		sequence = exponencial(10, n)
		distribution = "E[10]"
		method = "-1 * 10 * Ln[ 1 - rnd ]"
	case "op4":
		// Distribuciones op 1, 2 y 3
		sequence = UContUDiscExp(13, 17, 20, 25, 10, n)
		distribution = "U[13;17], U[20;25] y E[10]"
		method = ""
	case "op5":
		// Distribución empirica
		sequence = empirica(n)
		distribution = "Empírica"
		method = "si 0<=x<=1 : fx = x, si 1<=x<=2 : fx = 2-x, sobrante : fx = 0"
	case "op6":
		// Dist. Poisson con media = 1.8
		sequence = poisson(1.9, n)
		distribution = "P[1.9]"
		method = "POW(1.9, n) * POW(e, -1.9) / n!"
	case "op7":
		// Dist. normal con media = 3 min y desv. = 0.5 min (a partir de las 21hs)
		sequence = normal(3, 0.5, n)
		distribution = "N[3,0.5]"
		method = "rnd() * 0.5 + 3"
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", ViewData{distribution, method, sequence})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
