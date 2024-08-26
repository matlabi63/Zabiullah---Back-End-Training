package main

import (
	"fmt"
)

// Stack structure holds unique items using a map for uniqueness and a slice for ordering
type Stack struct {
	items map[string]bool
	order []string
}

// NewStack initializes and returns a new Stack instance
func NewStack() *Stack {
	return &Stack{
		items: make(map[string]bool),
		order: make([]string, 0),
	}
}

// push adds an item to the stack if it's not already present
func (s *Stack) push(item string) {
	if !s.items[item] {
		s.items[item] = true
		s.order = append(s.order, item)
	}
}

// pop removes and returns the top item from the stack
func (s *Stack) pop() string {
	if len(s.order) == 0 {
		return ""
	}

	lastIndex := len(s.order) - 1
	item := s.order[lastIndex]
	s.order = s.order[:lastIndex]
	delete(s.items, item)
	return item
}

// isEmpty checks if the stack is empty
func (s *Stack) isEmpty() bool {
	return len(s.order) == 0
}

// values returns all items in the stack in the order they were pushed
func (s *Stack) values() []string {
	return s.order
}

func main() {
	stack := NewStack()

	stack.push("item1")
	stack.push("item2")
	stack.push("item3")
	stack.push("item1") 

	fmt.Println("Stack values:", stack.values())

	fmt.Println("Pop item:", stack.pop())
	fmt.Println("Pop item:", stack.pop())

	fmt.Println("Is stack empty?", stack.isEmpty())

	fmt.Println("Stack values after pops:", stack.values())
}
