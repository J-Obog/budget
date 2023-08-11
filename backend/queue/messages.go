package queue

type Message struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`
}

type CategoryDeletedMessage struct {
	AccountId  string `json:"accountId"`
	CategoryId string `json:"categoryId"`
}
