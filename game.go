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
		newPlayer(game, player1),
		newPlayer(game, player2),
	}

	for _, player := range game.players {
		player.InitDeck(gameData.Decks[0])
	}

	return game
}

func (game *Game) DispatchEvent(player *Player, event Event) {
	player.HandleEvent(event)
}

func (game *Game) DispatchEventToOthers(player *Player, event Event) {
	for _, other := range game.players {
		if other != player {
			other.HandleEvent(event)
		}
	}
}

func (game *Game) DispatchEventToEveryone(event Event) {
	for _, player := range game.players {
		player.HandleEvent(event)
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
	game.DispatchEventToEveryone(NewGameUpdateEvent(game))
}
