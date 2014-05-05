package scraperapi

import (
	"dota/structs"
)

func mapFromAPIHeroToHero(h Hero) (hero structs.Hero) {
	hero = structs.Hero{
		Id:           "",
		Name:         h.Name,
		Faction:      h.Faction,
		Strength:     h.Strength,
		Agility:      h.Agility,
		Intelligence: h.Intelligence,
		Primary:      h.Primary,
	}
	return
}

func mapFromAPIItemToItem(i Item) (item structs.Item) {
	item = structs.Item{
		Id:          "",
		Name:        i.Name,
		Cost:        i.Cost,
		Disassemble: i.Disassemble,
	}
	return
}

type Hero struct {
	Name         string    `json:"name"`
	Faction      string    `json:"faction"`
	Strength     []float64 `json:"strength"`
	Agility      []float64 `json:"agility"`
	Intelligence []float64 `json:"intelligence"`
	Primary      string    `json:"primary"`
}

type Item struct {
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Disassemble bool   `json:"disassemble"`
}
