package main

import "fmt"

// Podemos declarar uma variável da seguinte forma
var c, java bool

// No caso de uma atribuição, o tipo pode ser omitido
var python, golang = true, "yes!"

func main() {

	// Também é possível declarar variáveis dentro de funções
	var i int

	// Dentro de uma função, podemos usar a declaração curta
	k := 3

	fmt.Println(i, c, python, java, golang, k)
}

// Variáveis declaradas sem um valor padrão possuem um "valor zero"
// 0 para tipos numéricos
// false para tipos boleanos
// "" para tipos strings
