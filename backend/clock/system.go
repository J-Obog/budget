package clock

import "time"

type SystemClock struct{}

func NewSystemClock() *SystemClock {
	return &SystemClock{}
}

func (c *SystemClock) Now() int64 {
	return time.Now().UnixNano()
}

func (c *SystemClock) CurrentMonth() int {
	return int(time.Now().Month())
}

func (c *SystemClock) CurrentYear() int {
	return time.Now().Year()
}
