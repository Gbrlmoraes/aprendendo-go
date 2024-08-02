package main

import (
	"fmt"
	"math"
)

// A expressão genérica T(v) converte o valor v para o tipo T
func main() {
	var x, y int = 3, 4

	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Println(x, y, z)
}

// Outros exemplos
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)

// Ou, dentro de funções, de forma mais simples
// i := 42
// f := float64(i)
// u := uint(f)
