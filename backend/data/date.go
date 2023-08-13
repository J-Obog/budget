package data

import (
	"fmt"
	"strconv"
	"time"
)

type Date struct {
	Month int
	Day   int
	Year  int
}

func NewDate(month int, day int, year int) Date {
	return Date{
		Month: month,
		Day:   day,
		Year:  year,
	}
}

func (d *Date) IsValid() bool {
	monthString := strconv.Itoa(d.Month)
	if d.Month < 10 {
		monthString = "0" + monthString
	}

	dayString := strconv.Itoa(d.Day)
	if d.Day < 10 {
		dayString = "0" + dayString
	}

	_, err := time.Parse(time.DateOnly, fmt.Sprintf("%d-%s-%s", d.Year, monthString, dayString))

	return err == nil
}
