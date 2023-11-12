package array

import (
	"reflect"
)

type sortParams[T interface{}] func(a T, b T) bool

func merge[T any](a, b Array[T], compareFn sortParams[T]) []T {
	size, i, j := len(a)+len(b), 0, 0
	result := make(Array[T], size, size)

	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			// all the elements of a have been sorted
			result[k] = b[j]
			j++
		case j == len(b):
			// all the elements of b have been sorted
			result[k] = a[i]
			i++
		case compareFn(a[i], b[j]):
			// element of a is less than element of b
			result[k] = a[i]
			i++
		default:
			// element of b is less than element of a
			result[k] = b[j]
			j++
		}
	}
	return result
}

func mergeSort[T interface{}](a Array[T], compareFn sortParams[T]) Array[T] {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	return merge(mergeSort(a[:mid], compareFn), mergeSort(a[mid:], compareFn), compareFn)
}

// compareFn is an optional parameter. only one argument is needed
func (a *Array[T]) Sort(compareFn ...sortParams[T]) {
	var comparator sortParams[T]

	if len(compareFn) > 0 {
		comparator = compareFn[0]
	} else {
		// If no comparator is provided, use the default for the type
		var zero T
		switch reflect.TypeOf(zero).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			comparator = sortParams[T](reflect.ValueOf(func(a T, b T) bool {
				return reflect.ValueOf(a).Int() < reflect.ValueOf(b).Int()
			}).Interface().(func(T, T) bool))
		case reflect.Float32, reflect.Float64:
			comparator = sortParams[T](reflect.ValueOf(func(a T, b T) bool {
				return reflect.ValueOf(a).Float() < reflect.ValueOf(b).Float()
			}).Interface().(func(T, T) bool))
		case reflect.String:
			comparator = sortParams[T](reflect.ValueOf(func(a T, b T) bool {
				return reflect.ValueOf(a).String() < reflect.ValueOf(b).String()
			}).Interface().(func(T, T) bool))
		default:
			panic("Unsupported type or provide one comparator function")
		}
	}

	*a = mergeSort(*a, comparator)
}
