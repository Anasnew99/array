package array

// Push adds a variadic number of arguments to the end of the array.
// It accepts zero or more arguments of any type and appends them to the end of the array.
// It returns the new length of the array.
func (a *Array[T]) Push(arr ...T) int {
	*a = append([]T(*a), arr...)
	return len(*a)
}

// Pop removes the last element from the array and returns it.
// If the array is empty, it panics with the message "Array is empty".
// It returns the removed element.
func (a *Array[T]) Pop() T {
	if len(*a) == 0 {
		panic("Array is empty")
	}
	item := []T(*a)[len(*a)-1]
	*a = []T(*a)[:len(*a)-1]
	return item
}

// Splice removes a number of elements from the array starting at a given index.
// It accepts an index and a count as arguments.
// It removes 'count' elements from the array starting at 'index'.
// It returns a new Array containing the removed elements.
func (a *Array[T]) Splice(index int, count int) Array[T] {
	newArray := make([]T, count)
	copy(newArray, (*a)[index:index+count])
	*a = append([]T(*a)[:index], []T(*a)[index+count:]...)
	return newArray
}
