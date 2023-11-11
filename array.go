package array

type Array[T interface{}] []T

func New[T interface{}](arr ...T) Array[T] {
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

func (a *Array[T]) Splice(index int, count int) Array[T] {

	newArray := make([]T, count)
	copy(newArray, (*a)[index:index+count])
	*a = append([]T(*a)[:index], []T(*a)[index+count:]...)

	return newArray

}
