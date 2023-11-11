package array

import (
	"testing"
)

func TestSort(t *testing.T) {
	arr := New[int](4, 1, 5, 2, 8)

	arr.Sort(func(a int, b int) bool {
		return a < b
	})

	arr2 := New[int](4, 1, 5, 2, 8)
	arr2.Sort()

	if arr[0] != 1 || arr[1] != 2 || arr[2] != 4 || arr[3] != 5 || arr[4] != 8 {
		t.Errorf("Not sorted correctly")
	}

	if arr2[0] != 1 || arr2[1] != 2 || arr2[2] != 4 || arr2[3] != 5 || arr2[4] != 8 {
		t.Errorf("Not sorted correctly")
	}
}
