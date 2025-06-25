package leetcode

import (
	"strings"
)

func reverseWord(word string) string {
	r := []rune(word) // transforma a string em slice de runas (pra suportar acentos etc)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // inverte
	}
	return string(r)
}

func ReverseEachWord(sentence string) string {
	words := strings.Fields(sentence) // divide por espaço
	for i, word := range words {
		words[i] = reverseWord(word) // inverte cada palavra
	}
	return strings.Join(words, " ") // junta de novo com espaço
}
