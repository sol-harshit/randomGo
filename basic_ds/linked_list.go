package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	// create a new node
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// traversing the newNode
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}
	list.Insert(4)
	list.Insert(3)
	list.Insert(55)
	list.Insert(11)
	list.Insert(59)

	list.Print()
}
