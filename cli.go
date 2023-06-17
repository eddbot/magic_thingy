package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cli() {
	trie := CreateNameTrie()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("enter a card name")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		fmt.Println(trie.getFuzzy(input))
	}

}
