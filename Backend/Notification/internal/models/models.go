package models

type Error struct {
	Error string `json:"error"`
}

type KafkaMessage struct {
	MessageType string `json:"messageType"`
	Message     string `json:"message"`
}

type MailMessage struct {
	Mail     string `json:"mail"`
	CheckInt int    `json:"checkInt"`
}
