package structs

type Hero struct {
	Id               string
	Name             string
	Faction          string
	Primary          string
	Strength         []float64
	Agility          []float64
	Intelligence     []float64
	HP               []float64
	Mana             []float64
	Damage           []string
	Armor            []float64
	AttacksPerSecond []float64
	MovementSpeed    float64
	TurnRate         float64
	SightRange       []float64
	AttackRange      string
	MissileSpeed     string
	AttackDuration   []float64
	CastDuration     []float64
	BaseAttackTime   float64
	ProfilePicture   string
}

func (h *Hero) MergeWith(merge Hero) {
	h.Name = merge.Name
	h.Faction = merge.Faction
	h.Primary = merge.Primary
	h.Strength = merge.Strength
	h.Agility = merge.Agility
	h.Intelligence = merge.Intelligence
	h.HP = merge.HP
	h.Mana = merge.Mana
	h.Damage = merge.Damage
	h.Armor = merge.Armor
	h.AttacksPerSecond = merge.AttacksPerSecond
	h.MovementSpeed = merge.MovementSpeed
	h.TurnRate = merge.TurnRate
	h.SightRange = merge.SightRange
	h.AttackRange = merge.AttackRange
	h.MissileSpeed = merge.MissileSpeed
	h.AttackDuration = merge.AttackDuration
	h.CastDuration = merge.CastDuration
	h.BaseAttackTime = merge.BaseAttackTime
	h.ProfilePicture = merge.ProfilePicture
}
