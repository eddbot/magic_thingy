package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func getRawResults() (RawResults, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.mtgo.com/en/mtgo/decklist/modern-league-2023-06-16", nil)

	if err != nil {
		return RawResults{}, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return RawResults{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return RawResults{}, err
	}

	rgx := regexp.MustCompile(`window\.MTGO\.decklists\.data = (.+);`)

	bd := rgx.FindString(string(body))

	replacements := map[string]string{
		";":                             "",
		"window.MTGO.decklists.data = ": "",
	}

	for oldWord, newWord := range replacements {
		bd = strings.Replace(bd, oldWord, newWord, -1)
	}
	var rawResults RawResults

	if err := json.Unmarshal([]byte(bd), &rawResults); err != nil {
		log.Fatal(err)
	}
	return rawResults, nil
}

type RawResults struct {
	ID    string `json:"_id"`
	Decks []struct {
		Player string `json:"player"`
		Deck   []struct {
			Sb        bool `json:"SB"`
			DeckCards []struct {
				CardAttributes struct {
					Color    string `json:"COLOR"`
					Cost     int    `json:"COST"`
					CardCode int    `json:"Card_Code"`
					Name     string `json:"NAME"`
					Rarity   string `json:"RARITY"`
					Set      string `json:"Set"`
					Type     string `json:"Type"`
				} `json:"CARD_ATTRIBUTES,omitempty"`
				Quantity     int `json:"Quantity"`
				RightLOGINID int `json:"Right_LOGINID"`
			} `json:"DECK_CARDS"`
		} `json:"deck"`
		Loginid int `json:"LOGINID"`
	} `json:"decks"`
	EventName string    `json:"event_name"`
	Date      time.Time `json:"date"`
	EventType string    `json:"event_type"`
}
