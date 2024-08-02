package main

// Importando os pacotes fmt e math/rand
// Também é possível usar várias funções de import separadas, mesmo não sendo uma boa prática.
import (
	"fmt"
	"math/rand"
)

// Por convenção o nome do pacote que será usado é
// o último elemento do caminho de importação
func main() {
	fmt.Println("Meu número favorito é", rand.Intn(10))
}

// Obs:
// O nome das funções importadas sempre começa com letra maiúscula.
