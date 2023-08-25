package main

func main() {
	initGameData()

	game := NewGame("test1", "test2")

	game.Start()

	// Test
	_, card := game.CurrentTurn().Hand.FindFirst(func(card *PlayingCard) bool {
		return card.IsTypeGlimmer()
	})
	game.CurrentTurn().PlayCharacter(card.UUID)
}
