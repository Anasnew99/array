// Package array provides a generic array implementation in Go.
package array

// Array is a generic data structure that can hold elements of any type.
type Array[T interface{}] []T

// New creates a new Array[T] from a variadic number of arguments.
// It accepts zero or more arguments of any type and returns an Array containing these elements.
func New[T interface{}](arr ...T) Array[T] {
	a := Array[T](arr)
	return a
}
