package data

type Message struct {
	Id        string      `json:"id"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type EntityDeletionMessage struct {
	EntityId string `json:"entityId"`
}

type TransactionsModifiedMessage struct {
	AccountId string `json:"accountId"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
}
