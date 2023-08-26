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

	log.Println("FIRST INK")
	_, card := game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.Details.IsInkwell()
	})
	game.CurrentTurn().ToInk(card.UUID)
	log.Println("SECOND INK")
	_, card = game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.Details.IsInkwell()
	})
	game.CurrentTurn().ToInk(card.UUID)

	log.Println("PLAY CHARACTER")
	_, card = game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.IsTypeGlimmer() && card.Details.Cost <= 2 // We have too inks available
	})
	game.CurrentTurn().PlayCharacter(card.UUID)

	gameData.Moves[0].Execute(card)
}
