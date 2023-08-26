package main

import (
	"errors"
	"github.com/google/uuid"
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

func NewPlayer(game *Game, name string) *Player {
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

	player.Pile = NewPlayingCardPile(player, CardPile)
	player.Hand = NewPlayingCardPile(player, CardPileHand)
	player.Table = NewPlayingCardPile(player, CardPileTable)
	player.Discard = NewPlayingCardPile(player, CardPileDiscard)
	player.Inkwell = NewPlayingCardPile(player, CardPileInkwell)

	game.DispatchPacket(player, NewPlayerUUIDAssignedPacket(player))

	return player
}

func (player *Player) InitDeck(deck *Deck) {
	playingCards := make([]*PlayingCard, deck.CardsAmount)
	counter := 0

	for typ, count := range deck.DeckDefinition {
		for i := 0; i < count; i++ {
			playingCards[counter] = NewPlayingCard(typ, player)
			counter++
		}
	}

	player.Pile.content = playingCards
	player.Pile.DispatchState()
}

func (player *Player) HandlePacket(packet Packet) {
	//TODO
	//log.Printf("EVENT[%s] %s", player.Name, PacketAsJson(packet))
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

	event := NewLuaPlacedEvent()
	err := card.TriggerLuaEvent(event)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(event)

	if !event.IsCancelled() {
		player.Hand.PickCard(index)
		player.Table.Add([]*PlayingCard{card})
		card.SetStatus(CardStatusDrying)
	}

	return nil
}
