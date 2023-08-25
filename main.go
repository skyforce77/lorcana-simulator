package main

import (
	"log"
	"time"
)

func main() {
	initGameData()

	game := NewGame("test1", "test2")

	game.Start()

	// Test
	game.CurrentTurn().DrawCards(20)

	log.Println("FIRST INK")
	time.Sleep(time.Second)
	_, card := game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.Details.IsInkwell()
	})
	game.CurrentTurn().ToInk(card.UUID)
	log.Println("SECOND INK")
	_, card = game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.Details.IsInkwell()
	})
	game.CurrentTurn().ToInk(card.UUID)

	// Next turn to be able to use ink
	log.Println("TURNS")
	time.Sleep(time.Second)
	game.Next()
	game.Next()

	log.Println("PLAY CHARACTER")
	time.Sleep(time.Second)
	_, card = game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.IsTypeGlimmer() && card.Details.Cost <= 2 // We have too inks available
	})
	game.CurrentTurn().PlayCharacter(card.UUID)
}
