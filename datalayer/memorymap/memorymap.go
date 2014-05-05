package memorymap

import (
	"dota/structs"
	"errors"
	"strconv"
)

type MemoryMap struct {
	heroids map[string]int
	heroes  []structs.Hero
	itemids map[string]int
	items   []structs.Item
}

func NewMemory() MemoryMap {
	m := MemoryMap{}
	m.heroids = make(map[string]int)
	m.heroes = make([]structs.Hero, 0)
	m.itemids = make(map[string]int)
	m.items = make([]structs.Item, 0)
	return m
}

func (m *MemoryMap) LoadData() {
	m.AddHero()
	m.AddHero()
	m.AddItem()
}

func (m *MemoryMap) AddHero() structs.Hero {
	x := len(m.heroes)
	id := strconv.Itoa(x)
	hero := structs.Hero{Id: id}
	m.heroids[id] = x
	m.heroes = append(m.heroes, hero)
	return hero
}

func (m *MemoryMap) PatchHero(h structs.Hero) structs.Hero {
	m.heroes[m.heroids[h.Id]] = h
	return m.heroes[m.heroids[h.Id]]
}

func (m *MemoryMap) GetHeroes() ([]structs.Hero, error) {
	return m.heroes, nil
}

func (m *MemoryMap) GetHeroById(id string) (structs.Hero, error) {
	if i, present := m.heroids[id]; present {
		return m.heroes[i], nil
	}
	return structs.Hero{}, errors.New("404")
}

func (m *MemoryMap) AddItem() structs.Item {
	x := len(m.items)
	id := strconv.Itoa(x)
	item := structs.Item{Id: id}
	m.itemids[id] = x
	m.items = append(m.items, item)
	return item
}

func (m *MemoryMap) PatchItem(i structs.Item) structs.Item {
	m.items[m.itemids[i.Id]] = i
	return m.items[m.itemids[i.Id]]
}

func (m *MemoryMap) GetItems() ([]structs.Item, error) {
	return m.items, nil
}

func (m *MemoryMap) GetItemById(id string) (structs.Item, error) {
	if i, present := m.itemids[id]; present {
		return m.items[i], nil
	}
	return structs.Item{}, errors.New("404")
}
