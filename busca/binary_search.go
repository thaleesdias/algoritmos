package busca

func Binary_Search(arr [11]int, alvo int) int {

	inicio := 0
	fim := len(arr) - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2

		if arr[meio] == alvo {
			return meio // achou!
		} else if arr[meio] < alvo {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}

	return -1 // não achou
}
