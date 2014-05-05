package structs

type Item struct {
	Id          string
	Name        string
	Disassemble bool
	Cost        int
}

func (i *Item) MergeWith(merge Item) {
	i.Name = merge.Name
	i.Disassemble = merge.Disassemble
	i.Cost = merge.Cost
}
