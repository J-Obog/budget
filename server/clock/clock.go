package clock

import "github.com/J-Obog/paidoff/data"

type Clock interface {
	Now() int64
	IsDateValid(date data.Date) bool
	FromDate(date data.Date) int64
	CurrentMonth() int
	CurrentYear() int
	DateFromStamp(timestamp int64) data.Date
}
