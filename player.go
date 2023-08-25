package main

type Player struct {
	Name string
	Pile *PlayingCardPile
	Hand *PlayingCardPile
}

func newPlayer(name string, deck *Deck) *Player {
	player := &Player{
		name,
		pileFromDeck(deck),
		&PlayingCardPile{
			make([]*PlayingCard, 0),
		},
	}
	return player
}

func (player *Player) pileToHand(count int) {
	picked := player.Pile.Pick(count)
	player.Hand.Add(picked)
}
