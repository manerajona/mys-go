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
  lambda:=15.0
  for count > 0 {
    x := -1.0 * lambda * math.Log(1.0 - rand.Float64())
    sequence = append(sequence, math.Round(x*1000)/1000)
    count--
  }
  return sequence
}
