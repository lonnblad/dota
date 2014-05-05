package structs

type Hero struct {
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
}

type Item struct {
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Disassemble bool   `json:"disassemble"`
}
