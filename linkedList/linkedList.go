package linkedlist

import "fmt"

// Nó da lista duplamente ligada
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// Lista duplamente ligada
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Adiciona no final
func (l *DoublyLinkedList) AddLast(value int) {
	newNode := &Node{value: value}

	if l.tail == nil {

		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

// Adiciona no início
func (l *DoublyLinkedList) AddFirst(value int) {
	newNode := &Node{value: value}

	if l.head == nil {

		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
}

// Remove do início
func (l *DoublyLinkedList) RemoveFirst() {
	if l.head == nil {
		fmt.Println("Lista vazia")
		return
	}

	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		// Lista ficou vazia
		l.tail = nil
	}
}

// Remove do final
func (l *DoublyLinkedList) RemoveLast() {
	if l.tail == nil {
		fmt.Println("Lista vazia")
		return
	}

	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		// Lista ficou vazia
		l.head = nil
	}
}

// Imprime da esquerda pra direita
func (l *DoublyLinkedList) PrintForward() {
	for current := l.head; current != nil; current = current.next {
		fmt.Printf("%d <-> ", current.value)
	}
	fmt.Println("nil")
}

// Imprime da direita pra esquerda
func (l *DoublyLinkedList) PrintBackward() {
	for current := l.tail; current != nil; current = current.prev {
		fmt.Printf("%d <-> ", current.value)
	}
	fmt.Println("nil")
}
