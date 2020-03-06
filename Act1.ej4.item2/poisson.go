package main

import (
	"math"
)

func calculateByPoisson(count int) map[int]float64 {
	const e, lambda = 2.71828, 2.0
	sequence := make(map[int]float64)
	for index := 1; count >= index; index++ {

		n := index - 1
		nFact := factorial(n) //n!

		x := (math.Pow(lambda, float64(n)) * math.Pow(e, -lambda)) / float64(nFact)
		sequence[index] = math.Round(x*1000) / 1000
	}
	return sequence
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return (n * factorial(n-1))
}
