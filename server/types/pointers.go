package types

func Ptr[T any](val T) *T {
	ptr := new(T)
	*ptr = val
	return ptr
}

func StringPtr(v string) *string {
	return Ptr[string](v)
}
