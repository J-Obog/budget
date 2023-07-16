package queue

type Message struct {
	Id   string      `json:"id"`
	Data interface{} `json:"data"`
}

type CategoryDeletedMessage struct {
	CategoryId string `json:"categoryId"`
}
