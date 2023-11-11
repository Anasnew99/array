package array

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	arr := New[int](1, 2, 3)
	if len(arr) != 3 {
		t.Errorf("New[int](1, 2, 3) = %d; want 3", len(arr))
	}
}

func TestPush(t *testing.T) {
	arr := New[int](1, 2, 3)
	arr.Push(4)
	if len(arr) != 4 {
		t.Errorf("Push(4) = %d; want 4", len(arr))
	}
}

func TestPop(t *testing.T) {
	arr := New[int](1, 2, 3)
	arr.Push(4)
	last := arr.Pop()
	if last != 4 || len(arr) != 3 {
		t.Errorf("Pop() = %d; want 4", last)
		t.Errorf("After Pop, len(arr) = %d; want 3", len(arr))
	}
}

func TestSplice(t *testing.T) {
	arr := New[int](1, 2, 3)
	arr.Push(4)
	fmt.Println(arr)
	spliced := arr.Splice(1, 2)
	fmt.Println(spliced)
	if len(arr) != 2 || len(spliced) != 2 || spliced[0] != 2 || spliced[1] != 3 {
		t.Errorf("Splice(1, 2) = %d; want 2", len(spliced))
		t.Errorf("After Splice, len(arr) = %d; want 2", len(arr))
	}

}
