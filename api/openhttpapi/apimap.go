package openhttpapi

import (
	"dota/structs"
)

func mapToAPIHeroes(hs []structs.Hero) (heroes []Hero) {
	for _, h := range hs {
		heroes = append(heroes, mapToAPIHero(h))
	}
	return
}

func mapToAPIHero(h structs.Hero) (hero Hero) {
	hero = Hero{
		Id:      h.Id,
		Name:    h.Name,
		Faction: h.Faction,
		Type:    "hero",
	}
	return
}

func mapToAPIItems(is []structs.Item) (items []Item) {
	for _, i := range is {
		items = append(items, mapToAPIItem(i))
	}
	return
}

func mapToAPIItem(i structs.Item) (item Item) {
	item = Item{
		Id:   i.Id,
		Name: i.Name,
		Cost: i.Cost,
		Type: "item",
	}
	return
}

type Hero struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Faction string `json:"faction"`
	Type    string `json:"type"`
}

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
	Type string `json:"type"`
}
