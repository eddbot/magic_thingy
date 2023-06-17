package main

import "log"

func decker() {
	// we need to parse this into a noice format
	rawResults, err := getRawResults()

	if err != nil {
		log.Fatal(err)
	}
	results := CreateResults(rawResults)

	filtered := results.findDecks("amulet")

	_ = filtered
}
