package data

type Date struct {
	Month int `json:"month"`
	Day   int `json:"day"`
	Year  int `json:"year"`
}

func NewDate(month int, day int, year int) Date {
	return Date{
		Month: month,
		Day:   day,
		Year:  year,
	}
}
