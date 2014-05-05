package structs

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
