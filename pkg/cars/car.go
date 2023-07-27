// Package cars is where all the different types of car's structures will be written.
package cars

import (
	"fmt"
	wh "test_code/pkg/wheels"
	mt "test_code/pkg/motors"
	do "test_code/pkg/doors"
	stb "test_code/pkg/seatbelts"
	gb "test_code/pkg/gearboxes"
)

const (
	wheelCount       = 4
	defaultWheelSize = 20
)

// Car represents a car type.
type Car struct {
	wheels *wh.Wheels
	door   *do.Door
	motor  *mt.Motor
	driverSeatbelt *stb.Seatbelt
	gearbox *gb.GearBox
}

// WheelManipulator function gets the reference to the WheelsManipulator in the car.
func (c *Car) WheelManipulator() wh.WheelsManipulator {
	if c.wheels == nil {
		c.wheels = wh.NewWheels(map[string]int{"count": wheelCount, "size": defaultWheelSize})
	}
	return c.wheels
}

// DoorManipulator function gets the reference to the DoorManipulator in the car.
func (c *Car) DoorManipulator() do.DoorManipulator {
	if c.door == nil {
		c.door = do.NewDoor(map[string]bool{"locked": true})
	}
	return c.door
}

// MotorManipulator function gets the reference to the MotorManipulator in the car.
func (c *Car) MotorManipulator() mt.MotorManipulator {
	if c.motor == nil {
		c.motor = mt.NewMotor(map[string]interface{}{"seatbelt": c.driverSeatbelt})
	}
	return c.motor
}

// SeatbeltManipulator function gets the reference to the SeatbeltManipulator in the car.
func (c *Car) SeatbeltManipulator() stb.SeatbeltManipulator {
	if c.driverSeatbelt == nil {
		c.driverSeatbelt = stb.NewSeatbelt(nil)
	}
	return c.driverSeatbelt
}

// GearboxManipulator function gets the reference to the GearboxManipulator in the car.
func (c *Car) GearboxManipulator() gb.GearboxManipulator {
	if c.gearbox == nil {
		c.gearbox = gb.NewGearBox(nil)
	}
	return c.gearbox
}

// SetMotorRpm function sets the car rpm's considerating the interactions between car parts.
func (c *Car) SetMotorRpm(rpm int) error {
	seatbelt := c.SeatbeltManipulator()
	if seatbelt!=nil && !seatbelt.IsEngaged() && rpm>0{  //Seatbelt security
		return fmt.Errorf("Not Seatbelt Engaged")
	}
	c.MotorManipulator().SetRPM(rpm)
	c.WheelManipulator().SetRPM(int(float64(rpm)*c.GearboxManipulator().GetGearRatio()))
	return nil;
}

// IncreaseSpeed function its used to increase the car rpm's considering a max limit.
func (c *Car) IncreaseSpeed(rpm,maxrpm int) error {
	if c.MotorManipulator().GetRPM()>=maxrpm{
		return fmt.Errorf("Max speed reached")
	}
	if (rpm+c.MotorManipulator().GetRPM())>maxrpm{ rpm=rpm+c.MotorManipulator().GetRPM()-maxrpm; }
	return c.SetMotorRpm(rpm+c.MotorManipulator().GetRPM())
}

// Print function its used print the structure info.
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