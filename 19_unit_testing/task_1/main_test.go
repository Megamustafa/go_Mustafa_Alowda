package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length, width int
		expected      string
	}{
		{4, 5, "even rectangle"},
		{3, 5, "odd rectangle"},
		{6, 2, "even rectangle"},
		{7, 3, "odd rectangle"},
	}

	for _, tt := range tests {
		result := RectangleArea(tt.length, tt.width)
		if result != tt.expected {
			t.Errorf("RectangleArea(%d, %d) = %v; want %v", tt.length, tt.width, result, tt.expected)
		}
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length, width int
		expected      bool
	}{
		{4, 6, true},
		{5, 5, true},
		{7, 3, true},
		{8, 2, true},
		{0, 1, false},
		{10, 0, true},
		{20, 20, true},
	}

	for _, tt := range tests {
		result := RectanglePerimeter(tt.length, tt.width)
		if result != tt.expected {
			t.Errorf("RectanglePerimeter(%d, %d) = %v; want %v", tt.length, tt.width, result, tt.expected)
		}
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		s        int
		expected string
	}{
		{4, "even square"},
		{5, "odd square"},
		{6, "even square"},
		{7, "odd square"},
	}

	for _, tt := range tests {
		result := SquareArea(tt.s)
		if result != tt.expected {
			t.Errorf("SquareArea(%d) = %v; want %v", tt.s, result, tt.expected)
		}
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		s        int
		expected bool
	}{
		{4, false},
		{10, true},
		{9, false},
		{6, false},
		{0, false},
		{10, true},
	}

	for _, tt := range tests {
		result := SquarePerimeter(tt.s)
		if result != tt.expected {
			t.Errorf("SquarePerimeter(%d) = %v; want %v", tt.s, result, tt.expected)
		}
	}
}
