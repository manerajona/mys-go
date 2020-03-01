package main

import (
	"math"
	"math/rand"
)

func calculateByEmpiric(count int) map[int]float64 {
	sequence := make(map[int]float64)
	for index := 1; count >= index; index++ {
		x := rand.Float64() * 3
		if x > 1 {
			if x <= 2 {
				x = 2 - x
			} else {
				x = 0
			}
		}
		sequence[index] = math.Round(x*1000) / 1000
	}
	return sequence
}
