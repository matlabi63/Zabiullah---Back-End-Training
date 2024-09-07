package main

import (
	"testing"
)
func TestMain(t *testing.T) {
	main()
}

// TestRectangleArea tests the RectangleArea function.
func TestRectangleArea(t *testing.T) {
    tests := []struct {
        length, width int
        expected      string
    }{
        {4, 5, "even rectangle"},
        {3, 3, "odd rectangle"},
        {6, 7, "even rectangle"},
        {8, 4, "even rectangle"},
    }

    for _, test := range tests {
        result := RectangleArea(test.length, test.width)
        if result != test.expected {
            t.Errorf("RectangleArea(%d, %d) = %s; expected %s", test.length, test.width, result, test.expected)
        }
    }
}

// TestRectanglePerimeter tests the RectanglePerimeter function.
func TestRectanglePerimeter(t *testing.T) {
    tests := []struct {
        length, width int
        expected      bool
    }{
        {2, 3, true},
        {5, 5, true},
        {10, 10, true},
        {4, 7, false},
    }

    for _, test := range tests {
        result := RectanglePerimeter(test.length, test.width)
        if result != test.expected {
            t.Errorf("RectanglePerimeter(%d, %d) = %v; expected %v", test.length, test.width, result, test.expected)
        }
    }
}

// TestSquareArea tests the SquareArea function.
func TestSquareArea(t *testing.T) {
    tests := []struct {
        side     int
        expected string
    }{
        {2, "even square"},
        {3, "odd square"},
        {4, "even square"},
        {5, "odd square"},
    }

    for _, test := range tests {
        result := SquareArea(test.side)
        if result != test.expected {
            t.Errorf("SquareArea(%d) = %s; expected %s", test.side, result, test.expected)
        }
    }
}

// TestSquarePerimeter tests the SquarePerimeter function.
func TestSquarePerimeter(t *testing.T) {
    tests := []struct {
        side     int
        expected bool
    }{
        {5, false},
        {10, true},
        {8, false},
        {9, false},
    }

    for _, test := range tests {
        result := SquarePerimeter(test.side)
        if result != test.expected {
            t.Errorf("SquarePerimeter(%d) = %v; expected %v", test.side, result, test.expected)
        }
    }
}
