package car

import (
	"fmt"
	"test_code/doors"
	"test_code/motor"
	"test_code/seatbelt"
	"test_code/wheels"
)

const (
	wheelCount       = 4
	defaultWheelSize = 20
)

type Car struct {
	wheels   *wheels.Wheels
	door     *doors.Door
	motor    *motor.Motor
	seatbelt *seatbelt.Seatbelt
}

type CarManipulator interface {
	WheelManiputlator() wheels.WheelsManipulator
	DoorManipulator() doors.DoorManipulator
	MotorManipulator() motor.MotorManipulator
	SeatbeltManipulator() seatbelt.SeatbeltManipulator
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

func (c *Car) setMotorRPM(motorRpm int) error {
	if !c.seatbelt.IsEngaged() {
		return fmt.Errorf("seabelt must be engaged before motor rpm can be increased")
	}
	err := c.MotorManipulator().SetRPM(motorRpm)
	if err != nil {
		return err
	}
	return nil
}
func (c *Car) SetMotorAndWheels(motorRpm int) error {
	err := c.setMotorRPM(motorRpm)
	if err != nil {
		return err
	}
	err = c.WheelManipulator().SetRPM(motorRpm / 2)
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
}
