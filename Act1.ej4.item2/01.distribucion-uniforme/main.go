package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println(calculate(99))
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
