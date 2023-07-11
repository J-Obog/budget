package api

const (
	MaxAccountNameChars       int = 150
	MinAccountNameChars       int = 1
	MaxBudgetDescriptionChars int = 200
)

func checkBudgetDesciption(description string) error {
	return nil
}

func checkAccountName(name string) error {
	return nil
}

func checkDate(month int, day int, year int) error {
	return nil
}
