package models

type Account struct {
	Id          string
	Name        string
	Email       string
	Password    string
	IsActivated bool
	CreatedAt   int64
	UpdatedAt   int64
}
