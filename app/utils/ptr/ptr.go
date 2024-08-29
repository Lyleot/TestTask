package ptr

func AsRef[T any](a T) *T {
	return &a
}

func DeRef[T any](a *T) T {
	var zero T
	if a == nil {
		return zero
	}
	return *a
}
