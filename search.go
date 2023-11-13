package array

// Find returns the first element along with index in the array that satisfies the provided testing function.
// If no elements in the array satisfy the testing function, it returns index as -1.
// It accepts a testing function as an argument.
// The testing function accepts an element of the array as an argument and returns a boolean.

func (a Array[T]) Find(findFunc func(item T, index int) bool) (T, int) {
	for i, v := range a {
		if findFunc(v, i) {
			return v, i
		}
	}
	var zero T
	return zero, -1
}

// FindIndex returns the index of the first element in the array that satisfies the provided testing function.
func (a Array[T]) FindIndex(findFunc func(item T, index int) bool) int {
	_, i := a.Find(findFunc)
	return i
}

// FindLast returns the last element along with index in the array that satisfies the provided testing function.
func (a Array[T]) FindLast(findFunc func(item T, index int) bool) (T, int) {
	for i := len(a) - 1; i >= 0; i-- {
		if findFunc(a[i], i) {
			return a[i], i
		}
	}
	var zero T
	return zero, -1
}
