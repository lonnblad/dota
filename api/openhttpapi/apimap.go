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
		Id:               h.Id,
		Name:             h.Name,
		Faction:          h.Faction,
		Strength:         h.Strength,
		Agility:          h.Agility,
		Intelligence:     h.Intelligence,
		Primary:          h.Primary,
		HP:               h.HP,
		Mana:             h.Mana,
		Damage:           h.Damage,
		Armor:            h.Armor,
		AttacksPerSecond: h.AttacksPerSecond,
		MovementSpeed:    h.MovementSpeed,
		TurnRate:         h.TurnRate,
		SightRange:       h.SightRange,
		AttackRange:      h.AttackRange,
		MissileSpeed:     h.MissileSpeed,
		AttackDuration:   h.AttackDuration,
		CastDuration:     h.CastDuration,
		BaseAttackTime:   h.BaseAttackTime,
		ProfilePicture:   h.ProfilePicture,
		Type:             "hero",
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
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Faction          string    `json:"faction"`
	Strength         []float64 `json:"strength"`
	Agility          []float64 `json:"agility"`
	Intelligence     []float64 `json:"intelligence"`
	Primary          string    `json:"primary"`
	HP               []float64 `json:"hp"`
	Mana             []float64 `json:"mana"`
	Damage           []string  `json:"damage"`
	Armor            []float64 `json:"armor"`
	AttacksPerSecond []float64 `json:"attackspersecond"`
	MovementSpeed    float64   `json:"movementspeed"`
	TurnRate         float64   `json:"turnrate"`
	SightRange       []float64 `json:"sightrange"`
	AttackRange      string    `json:"attackrange"`
	MissileSpeed     string    `json:"missilespeed"`
	AttackDuration   []float64 `json:"attackduration"`
	CastDuration     []float64 `json:"castduration"`
	BaseAttackTime   float64   `json:"baseattacktime"`
	ProfilePicture   string    `json:"profilepicture"`
	Type             string    `json:"type"`
}

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
	Type string `json:"type"`
}
