package array

import (
	"testing"
)

func TestFind(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	findFunc := func(item int, index int) bool {
		return item == 3
	}
	expectedValue := 3
	expectedIndex := 2

	actualValue, actualIndex := arr.Find(findFunc)

	if actualValue != expectedValue {
		t.Errorf("Find() returned %v, expected %v", actualValue, expectedValue)
	}

	if actualIndex != expectedIndex {
		t.Errorf("Find() returned index %v, expected %v", actualIndex, expectedIndex)
	}
}

func TestFindIndex(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	findFunc := func(item int, index int) bool {
		return item == 3
	}
	expectedIndex := 2

	actualIndex := arr.FindIndex(findFunc)

	if actualIndex != expectedIndex {
		t.Errorf("FindIndex() returned index %v, expected %v", actualIndex, expectedIndex)
	}
}

func TestFindLast(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	findFunc := func(item int, index int) bool {
		return item == 3
	}
	expectedValue := 3
	expectedIndex := 2

	actualValue, actualIndex := arr.FindLast(findFunc)

	if actualValue != expectedValue {
		t.Errorf("FindLast() returned %v, expected %v", actualValue, expectedValue)
	}

	if actualIndex != expectedIndex {
		t.Errorf("FindLast() returned index %v, expected %v", actualIndex, expectedIndex)
	}
}
