package main

import (
	"math/rand"
	"time"
)

type PlayingCard struct {
	Type   *CardType
	Damage int
}

type PlayingCardPile struct {
	content []*PlayingCard
}

func (pile *PlayingCardPile) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pile.content), func(i, j int) {
		pile.content[i], pile.content[j] = pile.content[j], pile.content[i]
	})
}

func (pile *PlayingCardPile) Pick(count int) []*PlayingCard {
	toPick := count
	if len(pile.content) < count {
		toPick = len(pile.content)
	}

	cards := pile.content[len(pile.content)-toPick : len(pile.content)]
	pile.content = pile.content[0 : len(pile.content)-toPick]

	return cards
}

func (pile *PlayingCardPile) Add(cards []*PlayingCard) {
	for _, card := range cards {
		pile.content = append(pile.content, card)
	}
}

func pileFromDeck(deck *Deck) *PlayingCardPile {
	playingCards := make([]*PlayingCard, deck.CardsAmount)
	counter := 0

	for typ, count := range deck.DeckDefinition {
		for i := 0; i < count; i++ {
			playingCards[counter] = &PlayingCard{
				typ,
				0,
			}
			counter++
		}
	}

	return &PlayingCardPile{
		playingCards,
	}
}
