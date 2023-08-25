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

type PlayerUUIDAssignedEvent struct {
	ID     string    `json:"id"`
	Player uuid.UUID `json:"player"`
}

func NewPlayerUUIDAssignedEvent(player *Player) *PlayerUUIDAssignedEvent {
	return &PlayerUUIDAssignedEvent{
		"PlayerUUIDAssignedEvent",
		player.UUID,
	}
}

type CardPileCountUpdateEvent struct {
	ID     string    `json:"id"`
	Player uuid.UUID `json:"player"`
	Cards  int       `json:"cards"`
	Type   int       `json:"type"`
}

func NewCardPileCountUpdateEvent(pile *PlayingCardPile) *CardPileCountUpdateEvent {
	return &CardPileCountUpdateEvent{
		"CardPileCountUpdateEvent",
		pile.owner.UUID,
		len(pile.content),
		pile.pileType,
	}
}

type CardPileUpdateEvent struct {
	ID     string         `json:"id"`
	Player uuid.UUID      `json:"player"`
	Cards  []*PlayingCard `json:"cards"`
	Type   int            `json:"type"`
}

func NewCardPileUpdateEvent(pile *PlayingCardPile) *CardPileUpdateEvent {
	return &CardPileUpdateEvent{
		"CardPileUpdateEvent",
		pile.owner.UUID,
		pile.content,
		pile.pileType,
	}
}

type GameUpdateEvent struct {
	ID         string    `json:"id"`
	PlayerTurn uuid.UUID `json:"playerTurn"`
}

func NewGameUpdateEvent(game *Game) *GameUpdateEvent {
	return &GameUpdateEvent{
		"GameUpdateEvent",
		game.players[game.turn].UUID,
	}
}

type CardStateUpdateEvent struct {
	ID       string    `json:"id"`
	Status   int       `json:"status"`
	CardUUID uuid.UUID `json:"cardUUID"`
}

func NewCardStateUpdateEvent(card *PlayingCard) *CardStateUpdateEvent {
	return &CardStateUpdateEvent{
		"CardStateUpdateEvent",
		card.Status,
		card.UUID,
	}
}
