package dota2Wiki

import (
	"dota/conn/scraperconn"
	"dota/scrapers/dota2Wiki/parsers"
	"dota/scrapers/dota2Wiki/parsers/structs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetHeroes(heroes []string) {
	for _, h := range heroes {
		sHero := getKeyFromDota2Wiki(h)
		name, _ := parsers.ParseHeroName(sHero)
		faction, _ := parsers.ParseHeroFaction(sHero)
		strength, _ := parsers.ParseHeroStrength(sHero)
		agility, _ := parsers.ParseHeroAgility(sHero)
		intelligence, _ := parsers.ParseHeroIntelligence(sHero)
		primary, _ := parsers.ParseHeroPrimary(sHero)
		hero := structs.Hero{
			name,
			faction,
			strength,
			agility,
			intelligence,
			primary,
		}
		log.Printf("\n%+v", hero)

		bHero, _ := json.Marshal(hero)
		scraperconn.PostHero(bHero)
	}
}

func GetItems(items []string) {
	for _, i := range items {
		sItem := getKeyFromDota2Wiki(i)
		name, _ := parsers.ParseItemName(sItem)
		cost, _ := parsers.ParseItemCost(sItem)
		disassemble, _ := parsers.ParseItemDisassemble(sItem)
		item := structs.Item{
			name,
			cost,
			disassemble,
		}
		log.Printf("\n%+v", item)

		bItem, _ := json.Marshal(item)
		scraperconn.PostItem(bItem)
	}
}

func getKeyFromDota2Wiki(key string) string {
	res, err := http.Get("http://dota2.gamepedia.com/" + key)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	return string(result)
}
