package array

import (
	"reflect"
	"testing"
)

func TestForEach(t *testing.T) {
	// Test case 1: Empty array
	a := Array[int]{}
	a.ForEach(func(value int, index int) {
		t.Errorf("Unexpected call to callback function with value %v and index %v", value, index)
	})

	// Test case 2: Array with one element
	a = Array[int]{1}
	var result int
	a.ForEach(func(value int, index int) {
		result = value
	})
	if result != 1 {
		t.Errorf("Expected callback function to be called with value 1, but got %v", result)
	}

	// Test case 3: Array with multiple elements
	a = Array[int]{1, 2, 3, 4, 5}
	var sum int
	a.ForEach(func(value int, index int) {
		sum += value
	})
	if sum != 15 {
		t.Errorf("Expected sum of array elements to be 15, but got %v", sum)
	}
}

// write test case for rest of the functions in iterators.go

func TestFilter(t *testing.T) {
	a := Array[int]{5, 6, 7, 8, 9}
	b := a.Filter(func(value int, index int) bool {
		return value >= 8
	})

	if len(b) != 2 {
		t.Errorf("Expected length of array to be 2, but got %v", len(b))
	}

}

func TestMap(t *testing.T) {
	a := Array[int]{5, 6, 7, 8, 9}
	b := Map[int, int](a, func(value int, index int) int {
		return value * 2
	})

	if len(b) != 5 {
		t.Errorf("Expected length of array to be 5, but got %v", len(b))
	}

	expected := Array[int]{10, 12, 14, 16, 18}
	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Expected %v, but got %v", expected, b)
	}
}

func TestReduce(t *testing.T) {
	a := Array[int]{1, 2, 3, 4, 5}
	b := Reduce[int, map[string]int](a, func(pV map[string]int, cV int, index int) map[string]int {
		pV["sum"] += cV
		pV["count"]++
		return pV
	}, map[string]int{"sum": 0, "count": 0})

	if b["sum"] != 15 {
		t.Errorf("Expected sum of array elements to be 15, but got %v", b["sum"])
	}
	if b["count"] != 5 {
		t.Errorf("Expected count of array elements to be 5, but got %v", b["count"])
	}
}
