package db

type Account struct {
	Id           string
	Email        string
	Password     string
	EmailEnabled bool
	Tombstone    *int64
	CreatedAt    int64
	UpdatedAt    int64
}

type Category struct {
	Id        string
	Name      string
	ImageUrl  string
	CreatedAt int64
	UpdatedAt int64
}

type GoalType uint

const (
	GoalType_Expense GoalType = 0
	GoalType_Savings GoalType = 1
)

type Goal struct {
	Id            string
	AccountId     string
	CategoryId    string
	Month         int64
	Year          int64
	Name          string
	CurrentAmount float64
	TargetAmount  float64
	GoalType      GoalType
	CreatedAt     int64
	UpdatedAt     int64
}
