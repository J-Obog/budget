package data

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
