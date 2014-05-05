package businesslogic

import (
	"dota/structs"
)

type Logic struct {
	data DataLayer
}

func (l *Logic) SetDataMap(data DataLayer) {
	l.data = data
}

func (l *Logic) GetHeroes() ([]structs.Hero, error) {
	return l.data.GetHeroes()
}

func (l *Logic) GetHeroById(id string) (structs.Hero, error) {
	return l.data.GetHeroById(id)
}

func (l *Logic) GetItems() ([]structs.Item, error) {
	return l.data.GetItems()
}

func (l *Logic) GetItemById(id string) (structs.Item, error) {
	return l.data.GetItemById(id)
}

func (l *Logic) AddHero(h structs.Hero) (structs.Hero, error) {
	hero_new := l.data.AddHero()
	hero_new.MergeWith(h)
	l.data.PatchHero(hero_new)
	return hero_new, nil
}

func (l *Logic) AddItem(i structs.Item) (structs.Item, error) {
	item_new := l.data.AddItem()
	item_new.MergeWith(i)
	l.data.PatchItem(item_new)
	return item_new, nil
}

type DataLayer interface {
	AddHero() structs.Hero
	PatchHero(structs.Hero) structs.Hero
	GetHeroes() ([]structs.Hero, error)
	GetHeroById(string) (structs.Hero, error)
	AddItem() structs.Item
	PatchItem(structs.Item) structs.Item
	GetItems() ([]structs.Item, error)
	GetItemById(string) (structs.Item, error)
}

type OpenAPILogic interface {
	GetHeroes() ([]structs.Hero, error)
	GetHeroById(id string) (structs.Hero, error)
	GetItems() ([]structs.Item, error)
	GetItemById(id string) (structs.Item, error)
}

type ScraperAPILogic interface {
	AddHero(structs.Hero) (structs.Hero, error)
	AddItem(structs.Item) (structs.Item, error)
}
