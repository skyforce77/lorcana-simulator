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

type CardPileHiddenUpdateEvent struct {
	ID     string      `json:"id"`
	Player uuid.UUID   `json:"player"`
	Cards  []uuid.UUID `json:"cards"`
	Type   int         `json:"type"`
}

func NewCardPileHiddenUpdateEvent(pile *PlayingCardPile) *CardPileHiddenUpdateEvent {
	cards := make([]uuid.UUID, len(pile.content))

	for index, card := range pile.content {
		cards[index] = card.UUID
	}

	return &CardPileHiddenUpdateEvent{
		"CardPileHiddenUpdateEvent",
		pile.owner.UUID,
		cards,
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
