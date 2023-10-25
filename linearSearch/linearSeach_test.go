package linear

import (
	"testing"
)

var arr []int8 = []int8{1, 2, 3, 5, 8, 9, 44, 77, 108}

func TestAnyValue(t *testing.T) {
	found := LinearSearch(arr, 44)
	if !found {
		t.Fatal("Value 44 (any) not found")
	}
}

func TestFirstValue(t *testing.T) {
	found := LinearSearch(arr, 1)
	if !found {
		t.Fatal("Value 1 (first) not found")
	}
}

func TestLastValue(t *testing.T) {
	found := LinearSearch(arr, 77)
	if !found {
		t.Fatal("Value 77 (last) not found")
	}
}

func TestMiddleValue(t *testing.T) {
	found := LinearSearch(arr, 8)
	if !found {
		t.Fatal("Value 8 (middle) not found")
	}
}
