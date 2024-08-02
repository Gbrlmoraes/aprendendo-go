package main

import "fmt"

// Ao declarar os argumentos da função, como x e y são do mesmo tipo,
// podemos omitir o tipo dos dois, menos do último
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
