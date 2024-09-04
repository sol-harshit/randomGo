package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) Push(item int) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() int {
	// checking the height of the stack
	top := len(s.elements)
	if top == 0 {
		fmt.Println("stack is empty")
		return -1
	}

	elem := s.elements[top-1]
	s.elements = s.elements[:top-1]

	return elem
}

func main() {
	s := &Stack{}
	s.Push(2)
	s.Push(4)
	s.Push(8)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
