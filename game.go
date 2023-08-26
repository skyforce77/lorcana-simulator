package main

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	uuid        uuid.UUID
	players     []*Player
	turn        int
	lastUsed    time.Time
	turnTimeout *time.Time
}

func NewGame(player1 string, player2 string) *Game {
	game := &Game{
		uuid: uuid.New(),
		players: []*Player{
			nil,
			nil,
		},
		turn:        0,
		lastUsed:    time.Now(),
		turnTimeout: nil,
	}

	game.players = []*Player{
		NewPlayer(game, player1),
		NewPlayer(game, player2),
	}

	for _, player := range game.players {
		player.InitDeck(gameData.Decks[0])
	}

	return game
}

func (game *Game) DispatchPacket(player *Player, event Packet) {
	player.HandlePacket(event)
}

func (game *Game) DispatchPacketToOthers(player *Player, event Packet) {
	for _, other := range game.players {
		if other != player {
			other.HandlePacket(event)
		}
	}
}

func (game *Game) DispatchPacketToEveryone(event Packet) {
	for _, player := range game.players {
		player.HandlePacket(event)
	}
}

func (game *Game) Start() {
	for _, player := range game.players {
		player.DrawCards(7)
	}

	game.DispatchState()
}

func (game *Game) Next() {
	if game.turn >= len(game.players)-1 {
		game.turn = 0
	} else {
		game.turn++
	}

	game.CurrentTurn().Inkwell.ResetStatus()
	game.CurrentTurn().Table.ResetStatus()
	game.DispatchState()
}

func (game *Game) CurrentTurn() *Player {
	return game.players[game.turn]
}

func (game *Game) DispatchState() {
	game.DispatchPacketToEveryone(NewGameUpdatePacket(game))
}
