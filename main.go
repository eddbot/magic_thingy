package main

import (
	"fmt"
	"log"
)

func main() {

	// we need to parse this into a noice format
	rawResults, err := getRawResults()

	if err != nil {
		log.Fatal(err)
	}
	results := CreateResults(rawResults)

	filtered := results.findDecks("amulet")

	_ = filtered

	trie := CreateNameTrie()

	x := trie.getFuzzy("the")

	fmt.Println(x)

	for {
		fmt.Println("enter a card name")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(trie.getFuzzy(input))
	}

}
