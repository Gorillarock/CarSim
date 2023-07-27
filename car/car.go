package car

import (
	"fmt"
	"test_code/doors"
	"test_code/motor"
	"test_code/wheels"
)

const (
	wheelCount       = 4
	defaultWheelSize = 20
)

type Car struct {
	wheels *wheels.Wheels
	door   *doors.Door
	motor  *motor.Motor
}

type CarManipulator interface {
	WheelManiputlator() wheels.WheelsManipulator
	DoorManipulator() doors.DoorManipulator
	MotorManipulator() motor.MotorManipulator
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

func (c *Car) MotorManipulator() motor.MotorManipulator {
	if c.motor == nil {
		c.motor = &motor.Motor{}
	}
	return c.motor
}

func (c *Car) SetMotorAndWheels(motorRpm int) error {
	err := c.MotorManipulator().SetRPM(motorRpm)
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
}
