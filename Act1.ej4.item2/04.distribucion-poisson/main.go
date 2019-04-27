package main

import (
  "fmt"
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

  const e float64 = 2.71828
  sequence := []float64{}

  lambda:=2.0
  for count > 0 {
    n := count
    nfact := factorial(n) //n!
    x := (math.Pow(lambda, float64(n)) * math.Pow(e, -lambda)) / float64(nfact)
    sequence = append(sequence, math.Round(x*1000)/1000)
    count--
  }
  return sequence
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return (n * factorial(n-1))
}
