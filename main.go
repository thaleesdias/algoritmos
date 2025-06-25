package main

import (
	"algoritmos/leetcode"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go \"frase para inverter\"")
		return
	}

	input := os.Args[1]
	result := leetcode.ReverseEachWord(input)
	fmt.Println("Resultado:", result)
}
