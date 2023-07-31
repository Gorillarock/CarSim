package gearbox

import "fmt"

type Gearbox struct {
	Gear int
}

type GearboxManipulator interface {
	SetGear(int)
	GetGear() int
	Print()
}

func (g *Gearbox) SetGear(gear int) {
	g.Gear = gear
}

func (g *Gearbox) GetGear() int {
	return g.Gear
}

func (g *Gearbox) Print() {
	fmt.Printf("Gearbox: \n\tGear: %d\n", g.Gear)
}
