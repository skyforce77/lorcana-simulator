package main

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// Card types

const CardTypeGlimmer = "glimmer"
const CardTypeItem = "item"
const CardTypeAction = "action"
const CardTypeSong = "song"

type PlayingCard struct {
	UUID    uuid.UUID    `json:"uuid"`
	Details *CardDetails `json:"details"`
	Damage  int          `json:"damage"`
}

type PlayingCardPile struct {
	game    *Game
	owner   *Player
	content []*PlayingCard
	isHand  bool
	isPile  bool
}

func (card *PlayingCard) IsTypeGlimmer() bool {
	return card.Details.Type == CardTypeGlimmer
}

func (card *PlayingCard) IsTypeItem() bool {
	return card.Details.Type == CardTypeItem
}

func (card *PlayingCard) IsTypeAction() bool {
	return card.Details.Type == CardTypeAction
}

func (card *PlayingCard) IsTypeSong() bool {
	return card.Details.Type == CardTypeSong
}

func (card *PlayingCard) IsDead() bool {
	return card.IsTypeGlimmer() &&
		card.Damage >= card.Details.Willpower
}

func (pile *PlayingCardPile) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pile.content), func(i, j int) {
		pile.content[i], pile.content[j] = pile.content[j], pile.content[i]
	})
	pile.DispatchState()
}

func (pile *PlayingCardPile) Pick(count int) []*PlayingCard {
	toPick := count
	if len(pile.content) < count {
		toPick = len(pile.content)
	}

	cards := pile.content[len(pile.content)-toPick : len(pile.content)]
	pile.content = pile.content[0 : len(pile.content)-toPick]

	pile.DispatchState()
	return cards
}

func (pile *PlayingCardPile) Add(cards []*PlayingCard) {
	for _, card := range cards {
		pile.content = append(pile.content, card)
	}
	pile.DispatchState()
}

func (pile *PlayingCardPile) DispatchState() {
	// Control who sees cards
	if pile.isHand {
		pile.game.DispatchEventToOthers(pile.owner, NewCardCountUpdateEvent(pile))
		pile.game.DispatchEvent(pile.owner, NewCardUpdateEvent(pile))
	} else {
		pile.game.DispatchEventToEveryone(NewCardCountUpdateEvent(pile))
	}
}
