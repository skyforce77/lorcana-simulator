package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/goccy/go-yaml"
)

//go:embed resources/gameData.yaml
var rawGameData []byte

type CardAction struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type CardDetails struct {
	Name          string        `json:"name"`
	Title         string        `json:"title"`
	Cost          int           `json:"cost"`
	Inkwell       int           `json:"inkwell"`
	Strength      int           `json:"strength"`
	Willpower     int           `json:"willpower"`
	Color         int           `json:"color"`
	Type          string        `json:"type"`
	Action        string        `json:"action"`
	Actions       []string      `json:"actions"`
	Flavour       string        `json:"flavour"`
	Lore          int           `json:"lore"`
	Illustrator   string        `json:"illustrator"`
	ID            int           `json:"id"`
	Rarity        string        `json:"rarity"`
	Image         string        `json:"image"`
	FranchiseID   int           `json:"franchise_id"`
	Traits        []string      `json:"traits"`
	FrontImage    string        `json:"FrontImage"`
	FrontImageAlt string        `json:"FrontImageAlt"`
	BackImage     string        `json:"BackImage"`
	Amount        int           `json:"amount"` // For deck use only
	UniqueID      string        // Calculated
	Moves         []*CardAction // Calculated
}

type CardSet struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Promo  bool           `json:"promo"`
	Amount int            `json:"amount"`
	Cards  []*CardDetails `json:"cards"`
}

type Deck struct {
	ID             int
	Name           string         `json:"name"`
	Details        []*CardDetails `json:"cards"`
	DeckDefinition map[*CardDetails]int
	CardsAmount    int
}

type GameData struct {
	Sets  []*CardSet    `json:"sets"`
	Decks []*Deck       `json:"decks"`
	Moves []*CardAction `json:"moves"`
}

var gameData GameData
var cards = make(map[string]*CardDetails)

func initGameData() {
	err := yaml.Unmarshal(rawGameData, &gameData)
	if err != nil {
		panic(err)
	}

	processCards()
	processDecks()
}

func processCards() {
	// Collect moves
	moves := make(map[string]*CardAction)
	for _, move := range gameData.Moves {
		moves[move.ID] = move
	}

	// Compute cards unique ids and card map
	for _, set := range gameData.Sets {
		for _, card := range set.Cards {
			card.UniqueID = fmt.Sprintf("%d:%d", set.ID, card.ID)

			if val, ok := cards[card.UniqueID]; ok {
				panic(fmt.Sprintf("Card ID %s is already taken by %s - %s (meant to be replaced by %s - %s)",
					card.UniqueID, val.Name, val.Title, card.Name, card.Title))
			}

			card.Moves = make([]*CardAction, len(card.Actions))
			for index, action := range card.Actions {
				if move, ok := moves[action]; ok {
					card.Moves[index] = move
				} else {
					panic(fmt.Sprintf("Card %s - %s tries to use an undefined action %s",
						card.Name, card.Title, action))
				}
			}

			cards[card.UniqueID] = card
		}
	}
}

func processDecks() {
	// Attach cards definitions
	for _, deck := range gameData.Decks {
		err := deck.Import()
		if err != nil {
			panic(err)
		}
	}
}

func (card *CardDetails) IsInkwell() bool {
	return card.Inkwell != 0
}

func (deck *Deck) Import() error {
	cardAmount := 0
	deck.DeckDefinition = make(map[*CardDetails]int)

	for cardId, card := range deck.Details {
		cardAmount += card.Amount
		var found = false

		for _, cardType := range cards {
			if cardType.Name == card.Name && cardType.Title == card.Title {
				found = true
				deck.Details[cardId] = cardType
				deck.DeckDefinition[cardType] = card.Amount
				break
			}
		}

		if !found {
			return errors.New(fmt.Sprintf("Card not found for deck %s (meant to be %s - %s)",
				deck.Name, card.Name, card.Title))
		}
	}

	if cardAmount != 60 {
		return errors.New(fmt.Sprintf("Deck %s should have 60 cards (currently %d)",
			deck.Name, cardAmount))
	}

	deck.CardsAmount = cardAmount

	return nil
}
