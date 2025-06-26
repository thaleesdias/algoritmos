package busca

import "fmt"

func Binary_Search(arr []int, alvo int) int {

	inicio := 0
	fim := len(arr) - 1
	cont := 0

	for inicio <= fim {
		cont++
		meio := (inicio + fim) / 2

		if arr[meio] == alvo {
			fmt.Println("Número de iterações:", cont)
			return meio // achou!
		} else if arr[meio] < alvo {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	fmt.Println("Número de iterações:", cont)

	return -1 // não achou
}
