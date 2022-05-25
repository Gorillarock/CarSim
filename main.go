package main

import (
	"fmt"
)

const (
	wheelCount       = 4
	defaultWheelSize = 20
)

type Wheels struct {
	rpm   int
	size  int
	count int
}

type WheelsManipulator interface {
	SetRPM(int) error
	SetSize(int) error
	GetRPM() int
	GetSize() int
	GetCount() int
	Print()
}

func (w *Wheels) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	w.rpm = rpm
	return nil
}

func (w *Wheels) GetRPM() int {
	return w.rpm
}

func (w *Wheels) SetSize(size int) error {
	if size < 0 {
		return fmt.Errorf("Negative size")
	}
	w.size = size
	return nil
}

func (w *Wheels) GetSize() int {
	return w.size
}

func (w *Wheels) GetCount() int {
	return w.count
}

func (w *Wheels) Print() {
	fmt.Printf("Wheels: \n\tRPM: %d, \n\tSize: %d, \n\tCount:%d\n", w.rpm, w.size, w.count)
}

type Door struct {
	open   bool
	locked bool
}

type DoorManipulator interface {
	SetOpen(bool) error
	SetLocked(bool)
	IsOpen() bool
	IsLocked() bool
	Print()
}

func (d *Door) SetOpen(open bool) error {
	if !d.locked {
		d.open = open
		return nil
	}
	if d.open {
		d.open = open
		return nil
	}
	return fmt.Errorf("Door is locked, cannot open.")
}

func (d *Door) IsOpen() bool {
	return d.open
}

func (d *Door) SetLocked(locked bool) {
	d.locked = locked
}

func (d *Door) IsLocked() bool {
	return d.locked
}

func (d *Door) Print() {
	fmt.Printf("Door: \n\tOpen: %t, \n\tLocked: %t\n", d.open, d.locked)
}

type Motor struct {
	on  bool
	rpm int
}

type MotorManipulator interface {
	SetOn(bool)
	SetRPM(int) error
	IsOn() bool
	GetRPM() int
	Print()
}

func (m *Motor) SetOn(on bool) {
	m.on = on
}

func (m *Motor) IsOn() bool {
	return m.on
}

func (m *Motor) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	m.rpm = rpm
	return nil
}

func (m *Motor) GetRPM() int {
	return m.rpm
}

func (m *Motor) Print() {
	fmt.Printf("Motor: \n\tOn: %t, \n\tRPM: %d\n", m.on, m.rpm)
}

type Car struct {
	wheels *Wheels
	door   *Door
	motor  *Motor
}

type CarManipulator interface {
	WheelManiputlator() WheelsManipulator
	DoorManipulator() DoorManipulator
	MotorManipulator() MotorManipulator
	Print()
}

func (c *Car) WheelManipulator() WheelsManipulator {
	if c.wheels == nil {
		c.wheels = &Wheels{count: wheelCount, size: defaultWheelSize}
	}
	return c.wheels
}

func (c *Car) DoorManipulator() DoorManipulator {
	if c.door == nil {
		c.door = &Door{locked: true}
	}
	return c.door
}

func (c *Car) MotorManipulator() MotorManipulator {
	if c.motor == nil {
		c.motor = &Motor{}
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

func main() {
	car := &Car{}
	//wheels := car.WheelManipulator()
	door := car.DoorManipulator()
	motor := car.MotorManipulator()

	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)

}
