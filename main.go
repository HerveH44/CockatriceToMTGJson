package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		xmlFile       = flag.String("xml", "", "the cockatrice xml v4 file")
		jsonDirectory = flag.String("dir", "", "the output directory for json files")
	)
	flag.Parse()

	if *xmlFile == "" {
		log.Fatal("xml file required")
	}

	if *jsonDirectory == "" {
		log.Fatal("jsonDirectory required")
	}

	file, err := ioutil.ReadFile(*xmlFile)
	if err != nil {
		log.Fatal("could not open the file xml.xml")
	}

	err = os.MkdirAll(*jsonDirectory, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	var d CockatriceSets
	if err = xml.Unmarshal(file, &d); err != nil {
		log.Println(err)
		log.Fatal("could not parse file")
	}
	m := make(map[string]*MTGJsonSet)

	for _, set := range d.Sets {
		m[set.Name] = &MTGJsonSet{
			Code:        set.Name,
			Name:        set.Longname,
			Type:        set.Settype,
			ReleaseDate: set.Releasedate,
		}
	}

	for _, c := range d.Cards {
		if set, ok := m[c.Set.Text]; ok {
			convertedManaCost, _ := strconv.ParseInt(c.Prop.Cmc, 10, 64)
			number, _ := strconv.ParseInt(c.Set.Num, 10, 64)

			powerToughness := strings.Split(c.Prop.Pt, "/")
			power := ""
			toughness := ""
			if len(powerToughness) == 2 {
				power = powerToughness[0]
				toughness = powerToughness[1]
			}
			card := Card{
				Name:              c.Name,
				Number:            int(number),
				Layout:            c.Prop.Layout,
				Names:             []string{c.Name},
				Loyalty:           c.Prop.Loyalty,
				Power:             power,
				Toughness:         toughness,
				ConvertedManaCost: int(convertedManaCost),
				Colors:            strings.Split(c.Prop.Colors, ""),
				Types:             makeTypes(c.Prop.Type),
				Supertypes:        []string{},
				ManaCost:          c.Prop.Manacost,
				URL:               c.Set.Picurl,
				Rarity:            c.Set.Rarity,
				ScryfallID:        "",
				Side:              "a",
				IsAlternative:     false,
			}
			set.addCard(card)
		}
	}

	for k := range m {
		jsonSet := m[k]
		jsonSet.BaseSetSize = len(jsonSet.Cards)
		set, err := json.MarshalIndent(jsonSet, "", "    ")
		if err != nil {
			log.Println(err)
		}

		err = ioutil.WriteFile(*jsonDirectory+"/"+jsonSet.Code+".json", set, 0644)
		if err != nil {
			log.Println(err)
		}
	}
}

func makeTypes(cardTypes string) (ret []string) {
	for _, cardType := range strings.Split(cardTypes, "—") {
		if trimmedType := strings.TrimSpace(cardType); trimmedType != "" {
			ret = append(ret, trimmedType)
		}
	}
	return
}
