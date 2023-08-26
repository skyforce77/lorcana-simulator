package main

import (
	"errors"
	"github.com/google/uuid"
	lua "github.com/yuin/gopher-lua"
	"log"
)

type Player struct {
	Game    *Game
	UUID    uuid.UUID
	Name    string
	Pile    *PlayingCardPile
	Hand    *PlayingCardPile
	Table   *PlayingCardPile
	Discard *PlayingCardPile
	Inkwell *PlayingCardPile
}

func newPlayer(game *Game, name string) *Player {
	player := &Player{
		game,
		uuid.New(),
		name,
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	player.Pile = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		CardPile,
	}

	player.Hand = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		CardPileHand,
	}

	player.Table = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		CardPileTable,
	}

	player.Discard = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		CardPileDiscard,
	}

	player.Inkwell = &PlayingCardPile{
		game,
		player,
		make([]*PlayingCard, 0),
		CardPileInkwell,
	}

	game.DispatchEvent(player, NewPlayerUUIDAssignedEvent(player))

	return player
}

func (player *Player) InitDeck(deck *Deck) {
	playingCards := make([]*PlayingCard, deck.CardsAmount)
	counter := 0

	for typ, count := range deck.DeckDefinition {
		for i := 0; i < count; i++ {
			playingCards[counter] = &PlayingCard{
				player.Game,
				player,
				uuid.New(),
				typ,
				0,
				CardStatusNone,
				make(map[string]*lua.LFunction),
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

func (player *Player) DrawCards(count int) {
	picked := player.Pile.Pick(count)
	player.Hand.Add(picked)
}

func (player *Player) ToInk(uuid uuid.UUID) error {
	index, card := player.Hand.FindByUUID(uuid)

	if card == nil {
		return errors.New("can't find card")
	}

	if !card.Details.IsInkwell() {
		return errors.New("should be inkwell")
	}

	if !card.HasNoStatus() {
		return errors.New("should be ready")
	}

	player.Hand.PickCard(index)
	player.Inkwell.Add([]*PlayingCard{card})
	return nil
}

func (player *Player) PlayCharacter(uuid uuid.UUID) error {
	index, card := player.Hand.FindByUUID(uuid)

	if card == nil {
		return errors.New("can't find card")
	}

	if !card.IsTypeGlimmer() {
		return errors.New("should be glimmer")
	}

	if player.Inkwell.NoStatusCount() < card.Details.Cost {
		return errors.New("not enough ink")
	}

	player.Hand.PickCard(index)
	player.Table.Add([]*PlayingCard{card})
	card.SetStatus(CardStatusDrying)

	return nil
}
