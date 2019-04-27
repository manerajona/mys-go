package main

import (
  "fmt"
  "math/rand"
  "math"
)
This can be used to generate random floats in
    // other ranges, for example `5.0 <= f' < 1
func main() {
  fmt.Println(calculate(99))
}

func calculate(count int) []float64 {

  sequence := []float64{}

  for count > 0 {
    x := rand.Float64() * 3

    if (x > 1) {
      if x <= 2 {
        x = 2 - x
      } else {
        x = 0
      }
    }
    sequence = append(sequence, math.Round(x*100)/100)
    count--
  }
  return sequence
}
