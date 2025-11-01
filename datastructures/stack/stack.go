package stack

import "fmt"

type Stack struct {
	items []int
}

// NewStack creates and returns a new empty stack
func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = s.items[:0]
}

func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Print("Stack (top to bottom): ")
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Printf("%d ", s.items[i])
	}
	fmt.Println()
}

func (s *Stack) ToSlice() []int {
	result := make([]int, len(s.items))
	for i := 0; i < len(s.items); i++ {
		result[i] = s.items[len(s.items)-1-i]
	}
	return result
}
