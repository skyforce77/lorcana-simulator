package main

import (
	"log"
)

func main() {
	initGameData()

	game := NewGame("test1", "test2")

	game.Start()

	// Test
	game.CurrentTurn().DrawCards(20)

	for i := 0; i < 4; i++ {
		_, card := game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
			return card.Details.IsInkwell()
		})
		err := game.CurrentTurn().ToInk(card.UUID)
		if err != nil {
			panic(err)
		}
	}

	log.Println("PLAY CHARACTER")
	_, card := game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.IsTypeGlimmer() && card.Details.Cost <= 2 // We have too inks available
	})
	card.Details = cards["1:18"]
	card.InitMoves()

	err := card.game.CurrentTurn().PlayCharacter(card.UUID)
	if err != nil {
		panic(err)
	}
}
