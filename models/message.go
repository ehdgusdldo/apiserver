package models

// Message message struct
type Message struct {
	Message string `example:"success"`
}

// FailedMessage failed message sturct
type FailedMessage struct {
	Message string `example:"failed"`
	err     interface{}
}
