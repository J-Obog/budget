package clock

import "github.com/J-Obog/paidoff/data"

type Clock interface {
	Now() int64
	MonthEnd(timestamp int64) int64
	MonthStart(timestamp int64) int64
	IsDateValid(date data.Date) bool
	FromDate(date data.Date) int64
	CurrentMonth() int
	CurrentYear() int
	DateFromStamp(timestamp int64) data.Date
}
