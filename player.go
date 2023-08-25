package main

import (
	"github.com/google/uuid"
	"log"
)

type Player struct {
	Game *Game
	UUID uuid.UUID
	Name string
	Pile *PlayingCardPile
	Hand *PlayingCardPile
	Ink  int
}

func newPlayer(game *Game, name string) *Player {
	player := &Player{
		game,
		uuid.New(),
		name,
		nil,
		nil,
		0,
	}

	player.Pile = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		false,
		true,
	}

	player.Hand = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		true,
		false,
	}

	return player
}

func (player *Player) InitDeck(deck *Deck) {
	playingCards := make([]*PlayingCard, deck.CardsAmount)
	counter := 0

	for typ, count := range deck.DeckDefinition {
		for i := 0; i < count; i++ {
			playingCards[counter] = &PlayingCard{
				uuid.New(),
				typ,
				0,
			}
			counter++
		}
	}

	player.Pile.content = playingCards
	player.Pile.DispatchState()
}

func (player *Player) HandleEvent(event Event) {
	//TODO
	log.Printf("EVENT[%s] %s", player.Name, EventAsJson(event))
}

func (player *Player) PileToHand(count int) {
	picked := player.Pile.Pick(count)
	player.Hand.Add(picked)
}
