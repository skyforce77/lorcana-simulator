package main

import "fmt"

type Game struct {
	players []*Player
}

func (game *Game) Start() {
	player := game.players[0]
	player.pileToHand(7)

	fmt.Printf("pile has %d cards\n", len(player.Pile.content))
	fmt.Printf("hand has %d cards\n", len(player.Hand.content))
}
