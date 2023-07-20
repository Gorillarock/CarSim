package car

import (
	"CarSim/constant"
	"CarSim/parts"
	"fmt"
)

type Car struct {
	wheels *parts.Wheels
	door   *parts.Door
	motor  *parts.Motor
}

func (c *Car) WheelManipulator() parts.WheelsManipulator {
	if c.wheels == nil {
		c.wheels = &parts.Wheels{Count: constant.WheelCount, Size: constant.DefaultWheelSize}
	}
	return c.wheels
}

func (c *Car) DoorManipulator() parts.DoorManipulator {
	if c.door == nil {
		c.door = &parts.Door{Locked: true}
	}
	return c.door
}

func (c *Car) MotorManipulator() parts.MotorManipulator {
	if c.motor == nil {
		c.motor = &parts.Motor{}
	}
	return c.motor
}

func (c *Car) Print() {
	fmt.Printf("Car: \n")
	w := c.WheelManipulator()
	w.Print()
	d := c.DoorManipulator()
	d.Print()
	m := c.MotorManipulator()
	m.Print()
}
