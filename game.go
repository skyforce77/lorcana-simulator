package main

type Game struct {
	players []*Player
}

func (game *Game) Start() {
	for _, player := range game.players {
		player.PileToHand(7)
	}
}
