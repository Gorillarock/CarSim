package car

import (
	"fmt"
	"test_code/doors"
	"test_code/gearbox"
	"test_code/motor"
	"test_code/seatbelt"
	"test_code/wheels"
)

const (
	wheelCount       = 4
	defaultWheelSize = 20
	defaultGear      = 2
)

type Car struct {
	wheels   *wheels.Wheels
	door     *doors.Door
	motor    *motor.Motor
	seatbelt *seatbelt.Seatbelt
	gearbox  *gearbox.Gearbox
}

type CarManipulator interface {
	WheelManiputlator() wheels.WheelsManipulator
	DoorManipulator() doors.DoorManipulator
	MotorManipulator() motor.MotorManipulator
	SeatbeltManipulator() seatbelt.SeatbeltManipulator
	GearboxManipulator() gearbox.Gearbox
	Print()
}

func (c *Car) WheelManipulator() wheels.WheelsManipulator {
	if c.wheels == nil {
		c.wheels = &wheels.Wheels{Count: wheelCount, Size: defaultWheelSize}
	}
	return c.wheels
}

func (c *Car) DoorManipulator() doors.DoorManipulator {
	if c.door == nil {
		c.door = &doors.Door{Locked: true}
	}
	return c.door
}

func (c *Car) SeatbeltManipulator() seatbelt.SeatbeltManipulator {
	if c.seatbelt == nil {
		c.seatbelt = &seatbelt.Seatbelt{
			Engaged: false,
		}
	}
	return c.seatbelt
}

func (c *Car) MotorManipulator() motor.MotorManipulator {
	if c.motor == nil {
		c.motor = &motor.Motor{}
	}
	return c.motor
}

func (c *Car) GearboxManipulator() gearbox.GearboxManipulator {
	if c.gearbox == nil {
		c.gearbox = &gearbox.Gearbox{
			Gear: defaultGear,
		}
	}
	return c.gearbox
}

func (c *Car) SetMotorRPM(motorRpm int) error {
	if !c.seatbelt.IsEngaged() {
		return fmt.Errorf("seabelt must be engaged before motor rpm can be increased")
	}
	err := c.MotorManipulator().SetRPM(motorRpm)
	if err != nil {
		return err
	}
	err = c.UpdateWheelsSpeed()
	if err != nil {
		return err
	}
	return nil
}

func (c *Car) ChangeGear(gear int) {
	c.GearboxManipulator().SetGear(gear)
	c.UpdateWheelsSpeed()
}

func (c *Car) UpdateWheelsSpeed() error {
	motorRpm := c.MotorManipulator().GetRPM()
	// check the gearbox relation
	gear := c.GearboxManipulator().GetGear()

	err := c.WheelManipulator().SetRPM(motorRpm / gear)
	if err != nil {
		return err
	}
	return nil
}

func (c *Car) Print() {
	fmt.Printf("Car: \n")
	w := c.WheelManipulator()
	w.Print()
	d := c.DoorManipulator()
	d.Print()
	m := c.MotorManipulator()
	m.Print()
	s := c.SeatbeltManipulator()
	s.Print()
	g := c.GearboxManipulator()
	g.Print()
}
