package data

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string
	IsActivated bool  `json:"isActivated"`
	IsDeleted   bool  `json:"isDeleted"`
	CreatedAt   int64 `json:"createdAt"`
	UpdatedAt   int64 `json:"updatedAt"`
}

type AccountUpdate struct {
	Id        string
	Name      string
	Timestamp int64
}
