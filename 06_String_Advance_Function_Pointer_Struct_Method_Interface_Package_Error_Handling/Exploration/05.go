package main

import (
	"fmt"
)

// UniqueStack is a stack that maintains unique items
type UniqueStack struct {
	stack []interface{}            // slice to store stack elements
	set   map[interface{}]struct{} // map to ensure uniqueness
}

// NewUniqueStack creates a new instance of UniqueStack
func NewUniqueStack() *UniqueStack {
	return &UniqueStack{
		stack: []interface{}{},
		set:   make(map[interface{}]struct{}),
	}
}

// Push adds an item to the stack if it's not already present
func (us *UniqueStack) Push(item interface{}) {
	if _, exists := us.set[item]; !exists {
		us.stack = append(us.stack, item)
		us.set[item] = struct{}{}
	}
}

// Pop removes and returns the item from the top of the stack
func (us *UniqueStack) Pop() (interface{}, bool) {
	if us.IsEmpty() {
		return nil, false
	}
	// Get the last item
	item := us.stack[len(us.stack)-1]
	// Remove the item from the stack and the set
	us.stack = us.stack[:len(us.stack)-1]
	delete(us.set, item)
	return item, true
}

// IsEmpty checks if the stack is empty
func (us *UniqueStack) IsEmpty() bool {
	return len(us.stack) == 0
}

// Values returns all items in the stack
func (us *UniqueStack) Values() []interface{} {
	copyStack := make([]interface{}, len(us.stack))
	copy(copyStack, us.stack)
	return copyStack
}

func main() {
	stack := NewUniqueStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(2) // This won't be added again

	fmt.Println("Stack values:", stack.Values()) // Output: Stack values: [1 2 3]

	if item, ok := stack.Pop(); ok {
		fmt.Println("Popped item:", item) // Output: Popped item: 3
	}

	fmt.Println("Stack values after pop:", stack.Values()) // Output: Stack values after pop: [1 2]
	fmt.Println("Is stack empty?", stack.IsEmpty())        // Output: Is stack empty? false
}
