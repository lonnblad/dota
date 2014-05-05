package structs

type Hero struct {
	Id           string
	Name         string
	Faction      string
	Primary      string
	Strength     []float64
	Agility      []float64
	Intelligence []float64
}

func (h *Hero) MergeWith(merge Hero) {
	h.Name = merge.Name
	h.Faction = merge.Faction
	h.Primary = merge.Primary
	h.Strength = merge.Strength
	h.Agility = merge.Agility
	h.Intelligence = merge.Intelligence
}
