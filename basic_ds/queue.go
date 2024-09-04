package main

import "fmt"

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(item int) {
	q.elements = append(q.elements, item)
}

func (q *Queue) Dequeue() int {
	length := len(q.elements)

	if length == 0 {
		fmt.Println("queue is empty")
		return -1
	}

	// first item to be fetched
	item := q.elements[0]

	// move the index to next element
	q.elements = q.elements[1:]
	return item

}

func main() {
	queue := &Queue{}
	queue.Enqueue(3)
	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(1)
	queue.Enqueue(4)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.elements)

}
