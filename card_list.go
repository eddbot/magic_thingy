package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type CardList map[string][]string

func (c CardList) getName(name string) []string {
	return c[name]
}

func (c CardList) getFuzzy(seq string) []string {

	var results []string
	names := map[string]bool{}

	for k, v := range c {
		if strings.Contains(k, seq) {
			for _, word := range v {
				names[word] = true
			}
		}

		if len(names) >= 10 {
			break
		}
	}
	for k := range names {
		results = append(results, k)
	}
	return results
}

func CreateNameTrie() CardList {

	names := CardList{}

	f, err := os.Open("cards.json")

	if err != nil {
		log.Fatal(err)
	}

	cards, err := io.ReadAll(f)

	if err != nil {
		log.Fatal(err)
	}

	var cc SryfallCards

	err = json.Unmarshal(cards, &cc)

	if err != nil {
		log.Fatal(err)
	}

	for _, x := range cc {
		name := strings.Split(strings.ToLower(x.Name), "")
		for i := 0; i < len(name); i++ {
			bb := name[0 : i+1]
			chars := strings.Join(bb, "")
			if _, ok := names[chars]; ok {
				names[chars] = append(names[chars], x.Name)
			} else {
				names[chars] = []string{x.Name}
			}
		}
	}
	return names
}

type SryfallCards []struct {
	Object        string `json:"object"`
	ID            string `json:"id"`
	OracleID      string `json:"oracle_id"`
	MultiverseIds []int  `json:"multiverse_ids"`
	MtgoID        int    `json:"mtgo_id"`
	MtgoFoilID    int    `json:"mtgo_foil_id,omitempty"`
	TcgplayerID   int    `json:"tcgplayer_id"`
	CardmarketID  int    `json:"cardmarket_id"`
	Name          string `json:"name"`
	Lang          string `json:"lang"`
	ReleasedAt    string `json:"released_at"`
	URI           string `json:"uri"`
	ScryfallURI   string `json:"scryfall_uri"`
	Layout        string `json:"layout"`
	HighresImage  bool   `json:"highres_image"`
	ImageStatus   string `json:"image_status"`
	ImageUris     struct {
		Small      string `json:"small"`
		Normal     string `json:"normal"`
		Large      string `json:"large"`
		Png        string `json:"png"`
		ArtCrop    string `json:"art_crop"`
		BorderCrop string `json:"border_crop"`
	} `json:"image_uris"`
	ManaCost      string  `json:"mana_cost"`
	Cmc           float64 `json:"cmc"`
	TypeLine      string  `json:"type_line"`
	OracleText    string  `json:"oracle_text"`
	Colors        []any   `json:"colors"`
	ColorIdentity []any   `json:"color_identity"`
	Keywords      []any   `json:"keywords"`
	Legalities    struct {
		Standard        string `json:"standard"`
		Future          string `json:"future"`
		Historic        string `json:"historic"`
		Gladiator       string `json:"gladiator"`
		Pioneer         string `json:"pioneer"`
		Explorer        string `json:"explorer"`
		Modern          string `json:"modern"`
		Legacy          string `json:"legacy"`
		Pauper          string `json:"pauper"`
		Vintage         string `json:"vintage"`
		Penny           string `json:"penny"`
		Commander       string `json:"commander"`
		Oathbreaker     string `json:"oathbreaker"`
		Brawl           string `json:"brawl"`
		Historicbrawl   string `json:"historicbrawl"`
		Alchemy         string `json:"alchemy"`
		Paupercommander string `json:"paupercommander"`
		Duel            string `json:"duel"`
		Oldschool       string `json:"oldschool"`
		Premodern       string `json:"premodern"`
		Predh           string `json:"predh"`
	} `json:"legalities"`
	Games           []string `json:"games"`
	Reserved        bool     `json:"reserved"`
	Foil            bool     `json:"foil"`
	Nonfoil         bool     `json:"nonfoil"`
	Finishes        []string `json:"finishes"`
	Oversized       bool     `json:"oversized"`
	Promo           bool     `json:"promo"`
	Reprint         bool     `json:"reprint"`
	Variation       bool     `json:"variation"`
	SetID           string   `json:"set_id"`
	Set             string   `json:"set"`
	SetName         string   `json:"set_name"`
	SetType         string   `json:"set_type"`
	SetURI          string   `json:"set_uri"`
	SetSearchURI    string   `json:"set_search_uri"`
	ScryfallSetURI  string   `json:"scryfall_set_uri"`
	RulingsURI      string   `json:"rulings_uri"`
	PrintsSearchURI string   `json:"prints_search_uri"`
	CollectorNumber string   `json:"collector_number"`
	Digital         bool     `json:"digital"`
	Rarity          string   `json:"rarity"`
	FlavorText      string   `json:"flavor_text,omitempty"`
	CardBackID      string   `json:"card_back_id"`
	Artist          string   `json:"artist"`
	ArtistIds       []string `json:"artist_ids"`
	IllustrationID  string   `json:"illustration_id"`
	BorderColor     string   `json:"border_color"`
	Frame           string   `json:"frame"`
	FullArt         bool     `json:"full_art"`
	Textless        bool     `json:"textless"`
	Booster         bool     `json:"booster"`
	StorySpotlight  bool     `json:"story_spotlight"`
	EdhrecRank      int      `json:"edhrec_rank"`
	Prices          struct {
		Usd       string `json:"usd"`
		UsdFoil   any    `json:"usd_foil"`
		UsdEtched any    `json:"usd_etched"`
		Eur       string `json:"eur"`
		EurFoil   any    `json:"eur_foil"`
		Tix       string `json:"tix"`
	} `json:"prices"`
	RelatedUris struct {
		Gatherer                  string `json:"gatherer"`
		TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
		TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
		Edhrec                    string `json:"edhrec"`
	} `json:"related_uris"`
	SecurityStamp string `json:"security_stamp,omitempty"`
	Preview       struct {
		Source      string `json:"source"`
		SourceURI   string `json:"source_uri"`
		PreviewedAt string `json:"previewed_at"`
	} `json:"preview,omitempty"`
	Power     string `json:"power,omitempty"`
	Toughness string `json:"toughness,omitempty"`
	PennyRank int    `json:"penny_rank,omitempty"`
}
