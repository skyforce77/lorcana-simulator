package main

type Game struct {
	players []*Player
}

func NewGame() *Game {
	game := &Game{
		players: []*Player{
			nil,
			nil,
		},
	}

	game.players = []*Player{
		newPlayer(game, "test1"),
		newPlayer(game, "test2"),
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
		player.PileToHand(7)
	}
}
