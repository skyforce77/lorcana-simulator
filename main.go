package main

func main() {
	initGameData()

	game := Game{
		players: []*Player{
			newPlayer("test1", gameData.Decks[0]),
			newPlayer("test2", gameData.Decks[1]),
		},
	}

	game.Start()
}
