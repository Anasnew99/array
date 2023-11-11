package array

import (
	"reflect"
)

type sortParams[T interface{}] func(a T, b T) bool

func merge[T interface{}](a Array[T], mid int, compareFn sortParams[T]) {
	// merge two sorted arrays
	i := 0
	j := mid
	for i < j && j < len(a) {
		if compareFn((a)[i], (a)[j]) {
			// move i to the right
			i++
		} else {
			temp := (a)[j]
			a[j] = (a)[i]
			(a)[i] = temp
			i++
			j++
		}
	}

}

func mergeSort[T interface{}](a Array[T], compareFn sortParams[T]) {
	if len(a) <= 1 {
		return
	}
	mid := len(a) / 2
	mergeSort(a[:mid], compareFn)
	mergeSort(a[mid:], compareFn)
	merge(a, mid, compareFn)

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

	mergeSort(*a, comparator)
}
