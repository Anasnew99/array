package array

import (
	"testing"
)

func TestNew(t *testing.T) {
	arr := New[int](1, 2, 3)
	if len(arr) != 3 {
		t.Errorf("New[int](1, 2, 3) = %d; want 3", len(arr))
	}
}
