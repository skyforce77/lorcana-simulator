package main

type Player struct {
	Name string
	Pile *PlayingCardPile
	Hand *PlayingCardPile
	Ink  int
}

func newPlayer(name string, deck *Deck) *Player {
	player := &Player{
		name,
		pileFromDeck(deck),
		&PlayingCardPile{
			make([]*PlayingCard, 0),
		},
		0,
	}
	return player
}

func (player *Player) PileToHand(count int) {
	picked := player.Pile.Pick(count)
	player.Hand.Add(picked)
}
