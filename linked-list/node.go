package main

type Node struct {
	data int
	next *Node
}

func NewNode(data int) Node {
	return Node{data: data}
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}
