package data

type Category struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	Color     uint   `json:"color"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

type CategoryUpdate struct {
	Id        string
	AccountId string
	Name      string
	Color     uint
}
