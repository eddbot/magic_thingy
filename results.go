package main

import (
	"strings"
)

type Result struct {
	Player    string
	Main      map[string]int
	Sideboard map[string]int
	Archetype string
}

type Results []Result

func CreateResults(rawResults RawResults) Results {
	results := Results{}
	for _, r := range rawResults.Decks {
		result := Result{Player: r.Player, Main: map[string]int{}, Sideboard: map[string]int{}}
		for _, d := range r.Deck {
			for _, c := range d.DeckCards {
				if !d.Sb {
					result.Main[c.CardAttributes.Name] += c.Quantity
				} else {
					result.Sideboard[c.CardAttributes.Name] += c.Quantity
				}
			}
		}
		results.Add(result)
	}
	return results
}

func (r *Results) Add(res Result) {
	*r = append(*r, res)
}

func (r *Results) findDecks(cardName string) []Result {

	decks := Results{}

	for _, res := range *r {
		for k := range res.Main {
			if strings.Contains(strings.ToLower(k), strings.ToLower(cardName)) {
				decks = append(decks, res)
			}
		}
	}
	return decks
}
