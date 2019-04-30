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
  printResult(calculate(n))
}

func printResult(res []float64) {
  fmt.Println("Resultado: ")
  for k,v := range res {
    fmt.Printf("%v) %v \n", k, v)
  }
}

const e float64 = 2.71828
const lambda=2.0
func calculate(count int) []float64 {
  sequence := []float64{}
  for n:=0; n <=count; n++ {
    nfact := factorial(n) //n!
    x := (math.Pow(lambda, float64(n)) * math.Pow(e, -lambda)) / float64(nfact)
    sequence = append(sequence, math.Round(x*1000)/1000)
  }
  return sequence
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return (n * factorial(n-1))
}
