// Package array provides a generic array implementation in Go.
package array

// Array is a generic data structure that can hold elements of any type.
type sortable interface {
	int8 | int | int16 | int32 | int64 | float64 | float32 | string
}
type Array[T interface{}] []T
type SortedArray[T sortable] Array[T]

// New creates a new Array[T] from a variadic number of arguments.
// It accepts zero or more arguments of any type and returns an Array containing these elements.
func New[T interface{}](arr ...T) Array[T] {
	a := Array[T](arr)
	return a
}
