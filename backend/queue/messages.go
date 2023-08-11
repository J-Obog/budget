package queue

type Message struct {
	Id   string `json:"id"`
	Body []byte `json:"body"`
}

type CategoryDeletedMessage struct {
	AccountId  string `json:"accountId"`
	CategoryId string `json:"categoryId"`
}
