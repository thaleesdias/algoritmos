package main

import (
	"algoritmos/busca"
	"fmt"
)

func main() {

	valores := []int{1, 3, 5, 6, 8, 13, 14, 15, 16, 22, 25}

	pos := busca.Binary_Search(valores, 25)

	if pos != -1 {
		fmt.Printf("Valor encontrado na posição %d\n", pos)
	} else {
		fmt.Println("Valor não encontrado")
	}
}
