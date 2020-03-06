package main

import (
	"math/rand"
)

func calculateByUniform(count int) map[int]float64 {
	const a, b = 20, 50
	sequence := make(map[int]float64)
	for index := 1; count >= index; index++ {
		x := rand.Intn(b-a) + a
		sequence[index] = float64(x)
	}
	return sequence
}
