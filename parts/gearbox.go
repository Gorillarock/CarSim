package parts

import "fmt"

type Gearbox struct {
	Motor       MotorManipulator
	Wheels      WheelsManipulator
	CurrentGear int
}

func (g *Gearbox) NewGearbox(motor MotorManipulator, wheels WheelsManipulator) *Gearbox {
	return &Gearbox{
		Motor:       motor,
		Wheels:      wheels,
		CurrentGear: 2,
	}
}

func (g *Gearbox) UpdateWheelsRPM() {
	if g.Motor == nil || g.Wheels == nil {
		return
	}
	motorRPM := g.Motor.GetRPM()
	wheelsRPM := motorRPM / g.CurrentGear
	g.Wheels.SetRPM(wheelsRPM)
}

func (g *Gearbox) SetGear(gear int) error {
	if gear > 0 {
		g.CurrentGear = gear
	} 
	return fmt.Errorf("invalid gear, must be greater than zero")
}
