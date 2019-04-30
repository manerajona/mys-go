package main

import (
  "fmt"
  "math/rand"
  "math"
  "os"
  "text/tabwriter"
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
    menu()
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
        // Dist. uniforme continua entre 13 y 17
        sequence := uniformeContinua(13, 17, n)
        printSimpleTable(sequence, op1, "U[13;17]", "13 + (17-3) * rnd")
    case 2:
        // Dist. uniforme discreta entre 20 y 25
        sequence := uniformeDiscreta(20, 25, n)
        printSimpleTable(sequence, op2, "U[20;25]", "20 + (25-20) * rnd")
    case 3:
        // Dist. exponencial con lambda**-1 = 10
        sequence := exponencial(10, n)
        printSimpleTable(sequence, op3, "E[10]", "-1 * 10 * Ln[ 1 - rnd ]")
    case 4:
        // Distribuciones op 1, 2 y 3
        sequence := UContUDiscExp(13, 17, 20, 25, 10, n)
        tableX3(sequence)
    case 5:
        // Distribución empirica
        sequence := empirica(n)
        printSimpleTable(sequence, op5, "Empirica", "si 0<=x<=1 : fx = x, si 1<=x<=2 : fx = 2-x, sobrante : fx = 0")
    case 6:
        // Dist. Poisson con media = 1.8
        sequence := poisson(1.9, n)
        printSimpleTable(sequence, op6, "P[1.9]", "POW(1.9, n) * POW(e, -1.9) / n!")
    case 7:
        HORA_INICIO := 21
        // Dist. normal con media = 3 min y desv. = 0.5 min (a partir de las 21hs)
        sequence := normal(3, 0.5, n)
        printTimeTable(sequence, op7, "N[3,0.5]", "rnd() * 0.5 + 3", HORA_INICIO)
    default:
        helpMenu()
  }
}

func printSimpleTable(fields []NumeroAleatorio, titulo string, distribucion string, formula string) {
  w := new(tabwriter.Writer)
  printHeader(w, titulo, distribucion, formula)
  for _, field := range fields {
    id:=strconv.Itoa(field.Id)
    na:=strconv.FormatFloat(field.NGenerado[0], 'f', -1, 64)
    rnd:=strconv.FormatFloat(field.Rnd, 'f', -1, 64)
    fmt.Fprintln(w, id+"\t"+na+"\t"+rnd+"\t")
    fmt.Fprintln(w)
  }
  printLine()
}

func tableX3(fields []NumeroAleatorio) {
  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	w.Flush()
  printLine()
  fmt.Fprintln(w, "Simulado\t Generado\t Random\t")
  for _, field := range fields {
    id:=strconv.Itoa(field.Id)
    na:=arrF64toString(field.NGenerado)
    rnd:=strconv.FormatFloat(field.Rnd, 'f', -1, 64)
    fmt.Fprintln(w, id+"\t"+na+"\t"+rnd+"\t")
    fmt.Fprintln(w)
  }
  printLine()
}

func arrF64toString(arr []float64) string {
  var str string
  i:=0
  for _, x := range arr {
    i++
    str+=" op"+strconv.Itoa(i)+") "+strconv.FormatFloat(x, 'f', -1, 64)
  }
  return str
}

func printTimeTable(fields []NumeroAleatorio, titulo string, distribucion string, formula string, horaInicio int) {
  w := new(tabwriter.Writer)
  printHeader(w, titulo, distribucion, formula)
  var min float64
  for _, field := range fields {
    id:=strconv.Itoa(field.Id)
    min += field.NGenerado[0]
    na:= strconv.Itoa(horaInicio) + ":" + strconv.FormatFloat(round(min), 'f', -1, 64)+"hs"
    rnd:=strconv.FormatFloat(field.Rnd, 'f', -1, 64)
    fmt.Fprintln(w, id+" \t"+na+" \t"+rnd)
    fmt.Fprintln(w)
  }
  printLine()
}

func printHeader(w *tabwriter.Writer, titulo string, distribucion string, formula string) {
  w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	w.Flush()
  printLine()
  fmt.Println("- Titulo: "+titulo)
  fmt.Println("- Distribución: "+distribucion)
  fmt.Println("- Formula: "+formula)
  printLine()
  fmt.Fprintln(w, "Simulado \tGenerado \tRandom\t")
}

func printLine() {
  fmt.Println("-----------------------------------------------")
}

func menu() {
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

func uniformeContinua(a int, b int, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64()
    x := rnd * float64(b - a) + float64(a)
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func uniformeDiscreta(a int, b int, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64();
    x := rnd * float64(b - a) + float64(a)
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func exponencial(lambda int, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64();
    x := -1.0 * float64(lambda) * math.Log(1.0 - rnd)
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func UContUDiscExp(a1 int, b1 int, a2 int, b2 int, lambda int, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64();

    x1 := math.Round(rnd) * float64(b1 - a1) + float64(a1)
    x2 := rnd * float64(b2 - a2) + float64(a2)
    x3 := -1.0 * float64(lambda) * math.Log(1.0 - rnd)

    xn := []float64{round(x1), round(x2), round(x3)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func empirica(n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64()
    x := rnd * 3
    if x > 1 {
      if x <= 2 {
        x = 2 - x
      } else {
        x = 0
      }
    }
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func poisson(lambda float64, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    nfact := factorial(i) //n!
    x := (math.Pow(lambda, float64(i)) * math.Pow(e, -lambda)) / float64(nfact)
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, 0)
  }
  return sequence
}

func normal(a float64, b float64, n int) []NumeroAleatorio {
  sequence := []NumeroAleatorio{}
  for i:=0; i<=n; i++ {
    rnd := rand.Float64()
    x := rnd * b + a
    xn := []float64{round(x)}
    sequence = addAleatorio(sequence, i, xn, rnd)
  }
  return sequence
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return (n * factorial(n-1))
}

func round(x float64) float64{
  return math.Round(x*DECIMAL_PRECISION)/DECIMAL_PRECISION
}

func addAleatorio(arr []NumeroAleatorio, id int, x []float64, rnd float64) []NumeroAleatorio {
  na := NumeroAleatorio {
    Id: id,
    NGenerado: x,
    Rnd: round(rnd),
  }
  arr = append(arr, na)
  return arr
}

type NumeroAleatorio struct {
  Id int
  NGenerado []float64
  Rnd float64
}
