package main

import (
  "fmt"
  "math/rand"
  "math"
)

func main() {
  fmt.Println(calculate(99))
}

func calculate(count int) []float64 {

  sequence := []float64{}
  lambda:=15.0
  for count > 0 {
    x := -1.0 * lambda * math.Log(1.0 - rand.Float64())
    sequence = append(sequence, math.Round(x*100)/100)
    count--
  }
  return sequence
}
