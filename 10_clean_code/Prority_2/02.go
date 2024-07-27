package main

import "fmt"

// Set represents a collection of unique integers
type Set struct {
	elements map[int]struct{}
}

// NewSet creates and returns a new Set
func NewSet() *Set {
	return &Set{
		elements: make(map[int]struct{}),
	}
}

// Add inserts an element into the set
func (s *Set) Add(element int) {
	s.elements[element] = struct{}{}
}

// GetElements returns a slice of all elements in the set
func (s *Set) GetElements() []int {
	result := make([]int, 0, len(s.elements))
	for element := range s.elements {
		result = append(result, element)
	}
	return result
}

// Remove deletes an element from the set
func (s *Set) Remove(element int) {
	delete(s.elements, element)
}

func main() {
	mySet := NewSet()

	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(1)
	mySet.Add(3)
	mySet.Add(4)
	mySet.Add(5)

	mySet.Remove(100) // Element 100 is not in the set, so no change

	fmt.Println(mySet.GetElements())
}
