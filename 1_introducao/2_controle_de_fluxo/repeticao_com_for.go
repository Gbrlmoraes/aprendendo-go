package main

import "fmt"

// Laço de repetição for básico possui apenas três componentes separados por ponto e vírgula:
// - A instrução inicial: executada antes da primeira iteração
// - A expressão de condição: avaliada antes de cada iteração
// - A pós-instrução: executado no final de cada iteração
func for_basico() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Somatório for básico:", sum)
}

// A primeira e a terceira instrução podem ser omitidas, Ex:
func for_basico_2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Somatório for básico 2:", sum)
}

// Obs: Omitir a condição em um laço torna ele infinito
func main() {
	for_basico()
	for_basico_2()
}
