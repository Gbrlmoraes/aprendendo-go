package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// A função swap está retornando dois valores
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
