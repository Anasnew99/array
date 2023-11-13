package array

// ForEach iterates over the array and calls the callback function for each element.
// The callback function is passed the value and index of the element.
func (a Array[T]) ForEach(cb func(value T, index int)) {
	for i, v := range a {
		cb(v, i)
	}
}

// Filter iterates over the array and calls the callback function for each element.
// The callback function is passed the value and index of the element.
// It returns a new Array containing the elements for which the callback function returned true.
func (a Array[T]) Filter(cb func(value T, index int) bool) Array[T] {
	var result = New[T]()
	for i, v := range a {
		if cb(v, i) {
			result = append(result, v)
		}
	}
	return result
}

// Map iterates over the array and calls the callback function for each element.
// The callback function is passed the value and index of the element.
// It returns a new Array containing the elements returned by the callback function.
func Map[T interface{}, K any](a Array[T], cb func(value T, index int) K) Array[K] {
	var result = New[K]()
	for i, v := range a {
		result = append(result, cb(v, i))
	}
	return result
}

// Reduce iterates over the array and calls the callback function for each element.
// The callback function is passed the previous value, current value and index of the element.
// It returns the final value.
func Reduce[T interface{}, K any](a Array[T], cb func(pV K, cV T, index int) K, initial K) K {
	var result = initial
	for i, v := range a {
		result = cb(result, v, i)
	}
	return result
}
