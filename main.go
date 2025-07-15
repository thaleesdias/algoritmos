package main

import (
	linkedlist "algoritmos/linkedList"
	"fmt"
)

func main() {

	list := linkedlist.DoublyLinkedList{}

	list.AddLast(10)
	list.AddLast(20)
	list.AddFirst(5)

	fmt.Println("Lista normal:")
	list.PrintForward() // 5 <-> 10 <-> 20 <-> nil

	fmt.Println("Lista reversa:")
	list.PrintBackward() // 20 <-> 10 <-> 5 <-> nil

	list.RemoveFirst()
	list.RemoveLast()

	fmt.Println("Depois de remover:")
	list.PrintForward() // 10 <-> nil
}
