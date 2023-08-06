package clock

type Clock interface {
	Now() int64
	CurrentMonth() int
	CurrentYear() int
}
