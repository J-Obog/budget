package models

type Category struct {
	Id        string `json:"id"`
	AccountId string `json:"accountId"`
	Name      string `json:"name"`
	Color     int    `json:"color"`
	UpdatedAt int    `json:"updatedAt"`
	CreatedAt int    `json:"createdAt"`
}
