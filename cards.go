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

// Card status

const (
	CardStatusNone = iota
	CardStatusExhausted
	CardStatusDrying
)

// Card pile type

const (
	CardPile = iota
	CardPileHand
	CardPileTable
	CardPileDiscard
	CardPileInkwell
)

type PlayingCard struct {
	game    *Game
	owner   *Player
	UUID    uuid.UUID    `json:"uuid"`
	Details *CardDetails `json:"details"`
	Damage  int          `json:"damage"`
	Status  int          `json:"status"`
}

type PlayingCardPile struct {
	game     *Game
	owner    *Player
	content  []*PlayingCard
	pileType int
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

func (card *PlayingCard) HasNoStatus() bool {
	return card.Status == CardStatusNone
}

func (card *PlayingCard) IsExhausted() bool {
	return card.Status == CardStatusExhausted
}

func (card *PlayingCard) SetStatus(status int) {
	card.Status = status
	card.DispatchState()
}

func (card *PlayingCard) DispatchState() {
	card.game.DispatchEventToEveryone(NewCardStateUpdateEvent(card))
}

func (pile *PlayingCardPile) IsPile() bool {
	return pile.pileType == CardPile
}

func (pile *PlayingCardPile) IsHand() bool {
	return pile.pileType == CardPileHand
}

func (pile *PlayingCardPile) IsTable() bool {
	return pile.pileType == CardPileTable
}

func (pile *PlayingCardPile) IsDiscard() bool {
	return pile.pileType == CardPileDiscard
}

func (pile *PlayingCardPile) IsInkwell() bool {
	return pile.pileType == CardPileInkwell
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

func (pile *PlayingCardPile) FindFirst(filter func(card *PlayingCard) bool) (int, *PlayingCard) {
	for index, card := range pile.content {
		if filter(card) {
			return index, card
		}
	}
	return 0, nil
}

func (pile *PlayingCardPile) FindByUUID(uuid uuid.UUID) (int, *PlayingCard) {
	return pile.FindFirst(func(card *PlayingCard) bool {
		return card.UUID == uuid
	})
}

func (pile *PlayingCardPile) Add(cards []*PlayingCard) {
	for _, card := range cards {
		pile.content = append(pile.content, card)
	}
	pile.DispatchState()
}

func (pile *PlayingCardPile) PickCard(index int) {
	pile.content = append(pile.content[0:index], pile.content[index+1:len(pile.content)]...)
	pile.DispatchState()
}

func (pile *PlayingCardPile) DispatchState() {
	// Control who sees cards
	if pile.IsHand() {
		pile.game.DispatchEventToOthers(pile.owner, NewCardPileHiddenUpdateEvent(pile))
		pile.game.DispatchEvent(pile.owner, NewCardPileUpdateEvent(pile))
	} else if pile.IsPile() || pile.IsDiscard() || pile.IsInkwell() {
		pile.game.DispatchEventToEveryone(NewCardPileHiddenUpdateEvent(pile))
	} else if pile.IsTable() {
		pile.game.DispatchEventToEveryone(NewCardPileUpdateEvent(pile))
	}
}

func (pile *PlayingCardPile) Length() int {
	return len(pile.content)
}

func (pile *PlayingCardPile) NoStatusCount() int {
	count := 0
	for _, card := range pile.content {
		if card.HasNoStatus() {
			count++
		}
	}
	return count
}

func (pile *PlayingCardPile) ResetStatus() {
	for _, card := range pile.content {
		if !card.HasNoStatus() {
			card.SetStatus(CardStatusNone)
		}
	}
}
