package queue

type AccountDeletionMessage struct {
	AccountId string `json:"accountId"`
}

type NotificationMessage struct {
	AccountEmail string `json:"acountEmail"`
}
