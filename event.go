package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

type EventHandler interface {
	HandleEvent(Event)
}

type Event interface{}

func EventAsJson(event Event) string {
	str, _ := json.Marshal(event)
	return string(str)
}

type CardCountUpdateEvent struct {
	ID     string    `json:"id"`
	Player uuid.UUID `json:"player"`
	Cards  int       `json:"cards"`
	IsHand bool      `json:"isHand"`
	IsPile bool      `json:"isPile"`
}

func NewCardCountUpdateEvent(pile *PlayingCardPile) *CardCountUpdateEvent {
	return &CardCountUpdateEvent{
		"CardCountUpdateEvent",
		pile.owner.UUID,
		len(pile.content),
		pile.isHand,
		pile.isPile,
	}
}

type CardUpdateEvent struct {
	ID     string         `json:"id"`
	Player uuid.UUID      `json:"player"`
	Cards  []*PlayingCard `json:"cards"`
	IsHand bool           `json:"isHand"`
	IsPile bool           `json:"isPile"`
}

func NewCardUpdateEvent(pile *PlayingCardPile) *CardUpdateEvent {
	return &CardUpdateEvent{
		"CardUpdateEvent",
		pile.owner.UUID,
		pile.content,
		pile.isHand,
		pile.isPile,
	}
}
