package main

import (
  "fmt"
  "math/rand"
  "math"
  "strconv"
)

/***************************************************
-------------------USER INTERFACE-------------------
****************************************************/

func main() {
  for {
    fmt.Println(">>>>>>>>>>>>>>>>>>>>MENU<<<<<<<<<<<<<<<<<<<<")
    fmt.Println("1. dl champagne.")
    fmt.Println("2. copas champagne.")
    fmt.Println("3. segundos entre pedido de copas.")
    fmt.Println("4. opcion 1.2.3 juntas.")
    fmt.Println("5. Segundos en extraer copa.")
    fmt.Println("6. personas por auto.")
    fmt.Println("7. minutos entre autos")
    fmt.Println("8. help")
    fmt.Println("0. salir")
    // input
    var n int
    _, err := fmt.Scanf("%d", &n)
    if err != nil {
      helpMenu()
    }
    // Exit or calculate
    if(n==0) {
      break
    } else {
      calculate(n)
    }
    fmt.Println("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-")
  }
}

func calculate(n int) {
  switch n {
    case 1:
        templateSimple(uniformeContinua(13, 17, 10), "dl champagne", "U[13;17]", "13 + (17-3) * rnd()")
    case 2:
        templateSimple(uniformeDiscreta(20, 25, 10), "copas champagne", "U[20;25]", "20 + (25-20) * rnd()")
    case 3:
        templateSimple(exponencial(10, 10), "segundos entre pedido de copas", "E[10]", "-1 * 10 * Ln[1.0 - rnd()]")
    case 4:
        templateXTres(UcUdExp(13, 17, 20, 25, 10, 10), 10)
    case 5:
        templateSimple(empirica(10), "Segundos en extraer copa", "Empirica", "f(x)= x si 0<=x<=1; f(x)= 2-x si 1<=x<=2; f(x)= 0 para cualquier otro caso;")
    case 6:
        templateSimple(poisson(1.9, 10), "personas por auto", "P[1.9]", "POW(1.9, n) * POW(e, -1.9) / n!")
    case 7:
        templateHora(normal(3, 0.5, 10), 21, "minutos entre autos", "N[3,0.5]", "rnd() * 0.5 + 3")
    default:
        helpMenu()
  }
}

func templateSimple(s []float64, titulo string, distribucion string, formula string) {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>"+titulo+"<<<<<<<<<<<<<<<<<<<<")
  fmt.Println("Distribución: "+distribucion)
  fmt.Println("Formula: "+formula)
  for k, v := range s {
    fmt.Println("id:", k)
    fmt.Println("NA:", v)
  }
}

func templateXTres(s []float64, n int) {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>dl champagne<<<<<<<<<<<<<<<<<<<<")
  for k, v := range s {
        if(k==n) {
          fmt.Println(">>>>>>>>>>>>>>>>>>>>copas champagne<<<<<<<<<<<<<<<<<<<<")
        } else if(k==n*2) {
            fmt.Println(">>>>>>>>>>>>>>>>>>>>csegundos entre pedido de copas<<<<<<<<<<<<<<<<<<<<")
        }
        fmt.Println("id:", k)
        fmt.Println("NA:", v)
    }
}

func templateHora(s []float64, hora int, titulo string, distribucion string, formula string) {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>"+titulo+"<<<<<<<<<<<<<<<<<<<<")
  fmt.Println("Distribución: "+distribucion)
  fmt.Println("Formula: "+formula)
  fmt.Println(s)
  var min int
  for k, v := range s {
    fmt.Println("id:", k)
    min += int(v)
    fmt.Println("NA:", strconv.Itoa(hora) + ":" + strconv.Itoa(min) + "hs")
  }
}

func helpMenu() {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>HELP<<<<<<<<<<<<<<<<<<<<")
  fmt.Println("Op 1) genera N.A. con dist. continua entre 13 y 17.")
  fmt.Println("Op 2) genera N.A. con dist. uniforme entre 20 y 25.")
  fmt.Println("Op 3) genera N.A. con dist. exponencial con 1/lambda=10.")
  fmt.Println("Op 4) genera N.A. con dist en 1, 2 y 3.")
  fmt.Println("Op 5) genera N.A. con dist. empirica.")
  fmt.Println("Op 6) genera N.A. con dist. poisson con media=1,8.")
  fmt.Println("Op 7) genera N.A. con dist. normal con media=3 min, desvio=0.5 min a partir de las 21hs")
}

/***************************************************
-------------------PROGRAM LOGIC--------------------
****************************************************/
const e float64 = 2.71828

func uniformeContinua(a int, b int, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    rnd := rand.Intn(b - a)
    x := rnd + a
    sequence = append(sequence, float64(x))
    n--
  }
  return sequence
}

func uniformeDiscreta(a int, b int, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    rnd := rand.Float64();
    x := rnd * float64(b - a) + float64(a)
    sequence = append(sequence, math.Round(x*100)/100)
    n--
  }
  return sequence
}

func exponencial(lambda int, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    rnd := rand.Float64();
    x := -1.0 * float64(lambda) * math.Log(1.0 - rnd)
    sequence = append(sequence, math.Round(x*100)/100)
    n--
  }
  return sequence
}

func UcUdExp(a1 int, b1 int, a2 int, b2 int, lambda int, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    rnd := rand.Float64();
    x1 := math.Round(rnd) * float64(b1 - a1) + float64(a1)
    sequence = append(sequence, x1)
    x2 := rnd * float64(b2 - a2) + float64(a2)
    sequence = append(sequence, math.Round(x2*100)/100)
    x3 := -1.0 * float64(lambda) * math.Log(1.0 - rnd)
    sequence = append(sequence, math.Round(x3*100)/100)
    n--
  }
  return sequence
}

func empirica(n int) []float64 {

  sequence := []float64{}

  for n > 0 {
    x := rand.Float64() * 3

    if (x > 1) {
      if x <= 2 {
        x = 2 - x
      } else {
        x = 0
      }
    }
    sequence = append(sequence, math.Round(x*100)/100)
    n--
  }
  return sequence
}

func poisson(lambda float64, n int) []float64 {

  sequence := []float64{}

  for n > 0 {
    nfact := factorial(n) //n!
    x := (math.Pow(lambda, float64(n)) * math.Pow(e, -lambda)) / float64(nfact)
    sequence = append(sequence, math.Round(x*100)/100)
    n--
  }
  return sequence
}

func normal(a float64, b float64, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    x := rand.Float64() * b + a
    sequence = append(sequence, math.Round(x*100)/100)
    n--
  }
  return sequence
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return (n * factorial(n-1))
}
