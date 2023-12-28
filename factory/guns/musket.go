package guns

type Musket struct {
	Name  string
	Power int
}

func (g *Musket) SetName(name string) {
	g.Name = name
}

func (g *Musket) SetPower(power int) {
	g.Power = power
}

func (g *Musket) GetName() string {
	return g.Name
}

func (g *Musket) GetPower() int {
	return g.Power
}
