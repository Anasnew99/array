package array

type Array[T comparable] []T

func New[T comparable](arr ...T) Array[T] {
	a := Array[T](arr)
	return a
}

func (a *Array[T]) Push(arr ...T) int {
	*a = append([]T(*a), arr...)

	return len(*a)
}

func (a *Array[T]) Pop() T {
	if len(*a) == 0 {
		panic("Array is empty")
	}
	item := []T(*a)[len(*a)-1]
	*a = []T(*a)[:len(*a)-1]
	return item
}
