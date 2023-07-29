package types

func Ptr[T any](val T) *T {
	ptr := new(T)
	*ptr = val
	return ptr
}

func StringPtr(v string) *string {
	return Ptr[string](v)
}

func IntPtr(v int) *int {
	return Ptr[int](v)
}

func Int64Ptr(v int64) *int64 {
	return Ptr[int64](v)
}

func Float64Ptr(v float64) *float64 {
	return Ptr[float64](v)
}
