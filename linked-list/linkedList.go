package main

import "fmt"

func createLinkedList(size int) Node {
	root := NewNode(0)

	current := &root

	for i := 1; i <= size; i++ {
		newNode := NewNode(i)
		current.SetNext(&newNode)
		current = current.next
	}

	return root
}

func getLinkedListLength(head *Node) int {
	count := 0

	current := head

	for {
		count++
		if current.next == nil {
			break
		}
		current = current.next
	}

	return count
}

func printLinkedList(head *Node) {
	current := head

	for {
		fmt.Printf("%d", current.data)
		if current.next == nil {
			break
		}
		fmt.Print(" -> ")
		current = current.next
	}
}

func runLinkedList() {
	root := createLinkedList(10)

	fmt.Println("Linked List's Lenght: ", getLinkedListLength(&root))
	fmt.Print("Linked List Items: ")

	printLinkedList(&root)
}
