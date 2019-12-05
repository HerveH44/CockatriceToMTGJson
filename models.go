package main

import "encoding/xml"

type CockatriceSets struct {
	XMLName xml.Name `xml:"cockatrice_carddatabase"`
	Version string   `xml:"version,attr"`
	Sets    []struct {
		Name        string `xml:"name"`
		Longname    string `xml:"longname"`
		Settype     string `xml:"settype"`
		Releasedate string `xml:"releasedate"`
	} `xml:"sets>set"`
	Cards []struct {
		Name       string `xml:"name"`
		Text       string `xml:"text"`
		Token      string `xml:"token"`
		Cipt       string `xml:"cipt"`
		Upsidedown string `xml:"upsidedown"`
		Tablerow   string `xml:"tablerow"`
		Set        struct {
			Text   string `xml:",chardata"`
			Rarity string `xml:"rarity,attr"`
			Uuid   string `xml:"uuid,attr"`
			Num    string `xml:"num,attr"`
			Muid   string `xml:"muid,attr"`
			Picurl string `xml:"picurl,attr"`
		} `xml:"set"`
		Related        string `xml:"related"`
		ReverseRelated string `xml:"reverse-related"`
		Prop           struct {
			Text            string `xml:",chardata"`
			Layout          string `xml:"layout"`
			Side            string `xml:"side"`
			Type            string `xml:"type"`
			Maintype        string `xml:"maintype"`
			Manacost        string `xml:"manacost"`
			Cmc             string `xml:"cmc"`
			Colors          string `xml:"colors"`
			Coloridentity   string `xml:"coloridentity"`
			Pt              string `xml:"pt"`
			Loyalty         string `xml:"loyalty"`
			FormatStandard  string `xml:"format-standard"`
			FormatCommander string `xml:"format-commander"`
			FormatModern    string `xml:"format-modern"`
			FormatPauper    string `xml:"format-pauper"`
		} `xml:"prop"`
	} `xml:"cards>card"`
}

type MTGJsonSet struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	ReleaseDate string   `json:"releaseDate"`
	BaseSetSize int      `json:"baseSetSize"`
	Cards       []Card   `json:"cards"`
	Booster     []string `json:"booster"`
}

func (set *MTGJsonSet) addCard(card Card) []Card {
	set.Cards = append(set.Cards, card)
	return set.Cards
}

type Card struct {
	Name              string   `json:"name"`
	Number            int      `json:"number"`
	Layout            string   `json:"layout"`
	Names             []string `json:"names"`
	Loyalty           string   `json:"loyalty"`
	Power             string   `json:"power"`
	Toughness         string   `json:"toughness"`
	ConvertedManaCost int      `json:"convertedManaCost"`
	Colors            []string `json:"colors"`
	Types             []string `json:"types"`
	Supertypes        []string `json:"supertypes"`
	ManaCost          string   `json:"manaCost"`
	URL               string   `json:"url"`
	Rarity            string   `json:"rarity"`
	ScryfallID        string   `json:"scryfallId"`
	Side              string   `json:"side"`
	IsAlternative     bool     `json:"isAlternative"`
}
