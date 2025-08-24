package strain

// Implement the "Keep" and "Discard" function in this file.
// func Keep given a collection and a predicate on the collection's elements, and returns a new collection containing those elements where the predicate is true
func Keep[T any](in []T, predicate func(v T) bool) (result []T) {
	result = make([]T, 0, len(in))
	for _, v := range in {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return
}

// func Discard given a collection and a predicate on the collection's elements, and returns a new collection containing those elements where the predicate is false
func Discard[T any](in []T, predicate func(v T) bool) (result []T) {
	result = make([]T, 0, len(in))
	for _, v := range in {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return
}
