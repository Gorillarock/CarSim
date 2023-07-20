package parts

import "fmt"

type Gearbox struct {
	Motor       MotorManipulator
	Wheels      WheelsManipulator
	CurrentGear string
}

func (g *Gearbox) SetMotor(motor MotorManipulator) {
	g.Motor = motor
}

func (g *Gearbox) SetWheels(wheels WheelsManipulator) {
	g.Wheels = wheels
}

func (g *Gearbox) UpdateWheelsRPM() {
	if g.Motor == nil || g.Wheels == nil {
		return
	}

	motorRPM := g.Motor.GetRPM()
	gearRatio := 0
	if g.CurrentGear == "2:1" {
		gearRatio = 2
	} else if g.CurrentGear == "4:1" {
		gearRatio = 4
	}

	wheelsRPM := motorRPM / gearRatio
	g.Wheels.SetRPM(wheelsRPM)
}

func (g *Gearbox) SetGear(gear string) error {
	switch gear {
	case "2:1", "4:1":
		g.CurrentGear = gear
	default:
		return fmt.Errorf("invalid gear, only 2:1 and 4:1")
	}

	return nil
}
