package guns

type AK47 struct {
	Name  string
	Power int
}

func (g *AK47) SetName(name string) {
	g.Name = name
}

func (g *AK47) SetPower(power int) {
	g.Power = power
}

func (g *AK47) GetName() string {
	return g.Name
}

func (g *AK47) GetPower() int {
	return g.Power
}
