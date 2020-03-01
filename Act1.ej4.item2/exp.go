package main

import (
	"math"
	"math/rand"
)

func calculateByExp(count int) map[int]float64 {
	const lambda = 15.0
	sequence := make(map[int]float64)

	for index := 1; count >= index; index++ {
		x := -1.0 * lambda * math.Log(1.0-rand.Float64())
		sequence[index] = math.Round(x*1000) / 1000
	}
	return sequence
}
