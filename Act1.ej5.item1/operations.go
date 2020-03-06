package main

import (
	"math"
	"math/rand"
)

const (
	e                 float64 = 2.71828
	DECIMAL_PRECISION         = 1000
)

func uniformeContinua(a int, b int, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()
		x := rnd*float64(b-a) + float64(a)
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func uniformeDiscreta(a int, b int, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()
		x := rnd*float64(b-a) + float64(a)
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func exponencial(lambda int, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()
		x := -1.0 * float64(lambda) * math.Log(1.0-rnd)
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func UContUDiscExp(a1 int, b1 int, a2 int, b2 int, lambda int, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()

		x1 := math.Round(rnd)*float64(b1-a1) + float64(a1)
		x2 := rnd*float64(b2-a2) + float64(a2)
		x3 := -1.0 * float64(lambda) * math.Log(1.0-rnd)

		xn := []float64{round(x1), round(x2), round(x3)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func empirica(n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()
		x := rnd * 3
		if x > 1 {
			if x <= 2 {
				x = 2 - x
			} else {
				x = 0
			}
		}
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func poisson(lambda float64, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		nfact := factorial(i) //n!
		x := (math.Pow(lambda, float64(i)) * math.Pow(e, -lambda)) / float64(nfact)
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, 0)
	}
	return sequence
}

func normal(a float64, b float64, n int) []aleatoryNumber {
	sequence := []aleatoryNumber{}
	for i := 0; i <= n; i++ {
		rnd := rand.Float64()
		x := rnd*b + a
		xn := []float64{round(x)}
		addAleatorio(&sequence, i, xn, rnd)
	}
	return sequence
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return (n * factorial(n-1))
}

func round(x float64) float64 {
	return math.Round(x*DECIMAL_PRECISION) / DECIMAL_PRECISION
}

func addAleatorio(arr *[]aleatoryNumber, id int, x []float64, rnd float64) {
	*arr = append(*arr, aleatoryNumber{
		Id:        id,
		Generated: x,
		Rnd:       round(rnd),
	})
}
