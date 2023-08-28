package slicex

func RemoveIndex[T any](key int, slice []T) []T {
	return append(slice[:key], slice[key+1:]...)
}
