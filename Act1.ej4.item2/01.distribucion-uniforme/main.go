package main

import (
  "fmt"
  "math/rand"
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

func printResult(res []int) {
  fmt.Println("Resultado: ")
  for k,v := range res {
    fmt.Printf("%v) %v \n", k, v)
  }
}

const a=20
const b=50
func calculate(count int) []int {
  sequence := []int{}
  for count > 0 {
    x := rand.Intn(b - a) + a
    sequence = append(sequence, x)
    count--
  }
  return sequence
}
