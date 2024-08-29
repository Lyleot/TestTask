package ptr

// AsRef создает указатель на значение и возвращает его.
// Полезно для создания указателя из значения.
func AsRef[T any](a T) *T {
	return &a
}

// DeRef возвращает значение, на которое указывает указатель.
// Если указатель равен nil, возвращает значение по умолчанию для типа T.
func DeRef[T any](a *T) T {
	var zero T
	if a == nil {
		return zero
	}
	return *a
}
