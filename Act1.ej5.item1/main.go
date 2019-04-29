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

const op1 = "dl champagne"
const op2 = "copas champagne"
const op3 = "segundos entre pedido de copas"
const op4 = "opcion 1.2.3 juntas"
const op5 = "Segundos en extraer copa"
const op6 = "personas por auto"
const op7 = "minutos entre autos"
const op8 = "help"
const op0 = "salir"

func main() {
  for {
    fmt.Println(">>>>>>>>>>>>>>>>>>>>MENU<<<<<<<<<<<<<<<<<<<<")
    fmt.Println("1."+op1)
    fmt.Println("2."+op2)
    fmt.Println("3."+op3)
    fmt.Println("4."+op4)
    fmt.Println("5."+op5)
    fmt.Println("6."+op6)
    fmt.Println("7."+op7)
    fmt.Println("8."+op8)
    fmt.Println("0."+op0)
    // Leer input
    var op int
    _, err := fmt.Scanf("%d", &op)
    if err != nil {
      helpMenu()
      fmt.Println("status: error")
    }
    // Salir o calcular
    if(op==0) {
      break
    } else {
      calculate(op)
    }
    fmt.Println("status: ok")
  }
}

func calculate(op int) {
  const n = 10
  switch op {
    case 1:
        sequence1 := uniformeContinua(13, 17, n)
        templateSimple(sequence1, op1, "U[13;17]", "13 + (17-3) * rnd()")
    case 2:
        sequence2 := uniformeDiscreta(20, 25, n)
        templateSimple(sequence2, op2, "U[20;25]", "20 + (25-20) * rnd()")
    case 3:
        sequence3 := exponencial(10, n)
        templateSimple(sequence3, op3, "E[10]", "-1 * 10 * Ln[1.0 - rnd()]")
    case 4:
        sequence4 := UcUdExp(13, 17, 20, 25, 10, n)
        templateXTres(sequence4, n)
    case 5:
        sequence5 := empirica(n)
        templateSimple(sequence5, op5, "Empirica", "f(x)= x si 0<=x<=1; f(x)= 2-x si 1<=x<=2; f(x)= 0 para cualquier otro caso;")
    case 6:
        sequence6 := poisson(1.9, n)
        templateSimple(sequence6, op6, "P[1.9]", "POW(1.9, n) * POW(e, -1.9) / n!")
    case 7:
        HORA_INICIO := 21
        sequence7 := normal(3, 0.5, n)
        templateHora(sequence7, op7, "N[3,0.5]", "rnd() * 0.5 + 3", HORA_INICIO)
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
  fmt.Println(">>>>>>>>>>>>>>>>>>>>"+op1+"<<<<<<<<<<<<<<<<<<<<")
  for k, v := range s {
        if(k==n) {
          fmt.Println(">>>>>>>>>>>>>>>>>>>>"+op2+"<<<<<<<<<<<<<<<<<<<<")
        } else if(k==n*2) {
            fmt.Println(">>>>>>>>>>>>>>>>>>>>"+op3+"<<<<<<<<<<<<<<<<<<<<")
        }
        fmt.Println("id:", k)
        fmt.Println("NA:", v)
    }
}

func templateHora(s []float64, titulo string, distribucion string, formula string, horaInicio int,) {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>"+titulo+"<<<<<<<<<<<<<<<<<<<<")
  fmt.Println("Distribución: "+distribucion)
  fmt.Println("Formula: "+formula)
  fmt.Println(s)
  var min int
  for k, v := range s {
    fmt.Println("id:", k)
    min += int(v)
    fmt.Println("NA:", strconv.Itoa(horaInicio) + ":" + strconv.Itoa(min) + "hs")
  }
}

func helpMenu() {
  fmt.Println(">>>>>>>>>>>>>>>>>>>>HELP<<<<<<<<<<<<<<<<<<<<")
  fmt.Println("Op 1) genera N.A. con dist. continua entre 13 y 17.")
  fmt.Println("Op 2) genera N.A. con dist. uniforme entre 20 y 25.")
  fmt.Println("Op 3) genera N.A. con dist. exponencial con 1/lambda=10.")
  fmt.Println("Op 4) genera N.A. con dist. en 1, 2 y 3.")
  fmt.Println("Op 5) genera N.A. con dist. empirica.")
  fmt.Println("Op 6) genera N.A. con dist. poisson con media=1,8.")
  fmt.Println("Op 7) genera N.A. con dist. normal con media=3 min, desvio=0.5 min a partir de las 21hs")
}

/***************************************************
-------------------PROGRAM LOGIC--------------------
****************************************************/
const e float64 = 2.71828
const DECIMAL_PRECISION = 1000

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
    sequence = append(sequence, math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION)
    n--
  }
  return sequence
}

func exponencial(lambda int, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    rnd := rand.Float64();
    x := -1.0 * float64(lambda) * math.Log(1.0 - rnd)
    sequence = append(sequence, math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION)
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
    sequence = append(sequence, math.Round(x2*DECIMAL_PRECISION)/DECIMAL_PRECISION)
    x3 := -1.0 * float64(lambda) * math.Log(1.0 - rnd)
    sequence = append(sequence, math.Round(x3*DECIMAL_PRECISION)/DECIMAL_PRECISION)
    n--
  }
  return sequence
}

func empirica(n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    x := rand.Float64() * 3
    if x > 1 {
      if x <= 2 {
        x = 2 - x
      } else {
        x = 0
      }
    }
    sequence = append(sequence, math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION)
    n--
  }
  return sequence
}

func poisson(lambda float64, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    nfact := factorial(n) //n!
    x := (math.Pow(lambda, float64(n)) * math.Pow(e, -lambda)) / float64(nfact)
    sequence = append(sequence, math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION)
    n--
  }
  return sequence
}

func normal(a float64, b float64, n int) []float64 {
  sequence := []float64{}
  for n > 0 {
    x := rand.Float64() * b + a
    sequence = append(sequence, math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION)
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
