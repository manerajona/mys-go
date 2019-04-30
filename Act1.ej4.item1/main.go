package main

import (
  "fmt"
  "math/rand"
)

func main() {
  var n int
  fmt.Println("Ingrese la cantidad de rondas : ")
  _, err := fmt.Scanf("%d", &n)
  if err != nil {
    fmt.Println(err)
  }
  printResult(calculate(n))
}

func printResult(res []int) {
  fmt.Println("Resultado: ")
  for _,v := range res {
    parity := "cero"
    if v!=0 {
      if v%2==0 {
        parity = "par"
      } else {
        parity = "impar"
      }
    }
    fmt.Printf("%d (%v) \n", v, parity)
  }
}

const a=0
const b=36
func calculate(count int) []int {
  sequence := []int{}
  for count > 0 {
    x := rand.Intn(b - a) + a
    sequence = append(sequence, x)
    count--
  }
  return sequence
}
