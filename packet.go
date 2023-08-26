package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

type PacketHandler interface {
	HandlePacket(Packet)
}

type Packet interface{}

func PacketAsJson(packet Packet) string {
	str, _ := json.Marshal(packet)
	return string(str)
}

type PlayerUUIDAssignedPacket struct {
	ID     string    `json:"id"`
	Player uuid.UUID `json:"player"`
}

func NewPlayerUUIDAssignedPacket(player *Player) *PlayerUUIDAssignedPacket {
	return &PlayerUUIDAssignedPacket{
		"PlayerUUIDAssignedPacket",
		player.UUID,
	}
}

type CardPileHiddenUpdatePacket struct {
	ID     string      `json:"id"`
	Player uuid.UUID   `json:"player"`
	Cards  []uuid.UUID `json:"cards"`
	Type   int         `json:"type"`
}

func NewCardPileHiddenUpdatePacket(pile *PlayingCardPile) *CardPileHiddenUpdatePacket {
	cards := make([]uuid.UUID, len(pile.content))

	for index, card := range pile.content {
		cards[index] = card.UUID
	}

	return &CardPileHiddenUpdatePacket{
		"CardPileHiddenUpdatePacket",
		pile.owner.UUID,
		cards,
		pile.pileType,
	}
}

type CardPileUpdatePacket struct {
	ID     string         `json:"id"`
	Player uuid.UUID      `json:"player"`
	Cards  []*PlayingCard `json:"cards"`
	Type   int            `json:"type"`
}

func NewCardPileUpdatePacket(pile *PlayingCardPile) *CardPileUpdatePacket {
	return &CardPileUpdatePacket{
		"CardPileUpdatePacket",
		pile.owner.UUID,
		pile.content,
		pile.pileType,
	}
}

type GameUpdatePacket struct {
	ID         string    `json:"id"`
	PlayerTurn uuid.UUID `json:"playerTurn"`
}

func NewGameUpdatePacket(game *Game) *GameUpdatePacket {
	return &GameUpdatePacket{
		"GameUpdatePacket",
		game.players[game.turn].UUID,
	}
}

type CardStateUpdatePacket struct {
	ID       string    `json:"id"`
	Status   int       `json:"status"`
	CardUUID uuid.UUID `json:"cardUUID"`
}

func NewCardStateUpdatePacket(card *PlayingCard) *CardStateUpdatePacket {
	return &CardStateUpdatePacket{
		"CardStateUpdatePacket",
		card.Status,
		card.UUID,
	}
}
