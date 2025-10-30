package model

type EventPayload struct {
	Pattern string     `json:"pattern"`
	Data    any `json:"data"`
}

type Event interface {
	GetId() string
}
