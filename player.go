package main

import (
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

func (player *Player) PlayCharacter(uuid uuid.UUID) bool {
	index, card := player.Hand.FindByUUID(uuid)
	log.Println(index, card, card.IsTypeGlimmer())

	if card == nil {
		return false
	}

	if !card.IsTypeGlimmer() {
		return false
	}

	player.Hand.PickCard(index)
	player.Table.Add([]*PlayingCard{card})
	card.SetStatus(CardStatusExhausted)

	return true
}
