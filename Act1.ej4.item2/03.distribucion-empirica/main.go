package main

import (
  "fmt"
  "math/rand"
  "math"
)

func main() {
  var n int
  fmt.Println("Ingrese la cantidad de num. aleatorios a generar : ")
  _, err := fmt.Scanf("%d", &n)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(calculate(n))
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
    sequence = append(sequence, math.Round(x*1000)/1000)
    count--
  }
  return sequence
}
