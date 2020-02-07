package models

type Message struct {
	Message string `example:"success"`
}

type FailedMessage struct {
	Message string `example:"failed"`
	err     interface{}
}
