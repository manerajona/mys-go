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
  fmt.Println(calculate(n))
}

func calculate(count int) []int {
  sequence := []int{}
  a:=20
  b:=50
  for count > 0 {
    x := rand.Intn(b - a) + a
    sequence = append(sequence, x)
    count--
  }
  return sequence
}
