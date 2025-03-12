package main

import "fmt"

func reverseLinkedList(head *Node) Node {
	var previous *Node = nil
	var current *Node = head
	var next *Node = head.next

	for {
		current.next = previous
		previous = current
		current = next

		if current == nil {
			break
		}

		next = next.next
	}

	return *previous
}

func runReverseLinkedList() {
	root := createLinkedList(10)

	reversed := reverseLinkedList(&root)

	fmt.Println("Linked List's Lenght: ", getLinkedListLength(&reversed))
	fmt.Print("Reversed Linked List Items: ")

	printLinkedList(&reversed)
}
