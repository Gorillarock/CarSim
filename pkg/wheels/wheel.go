// Package wheels is where all the different types of wheel's structures will be written.
package wheels

import (
	"fmt"
)

// Wheels represents a wheel type.
type Wheels struct {
	rpm   int
	size  int
	count int
}

// NewWheels function is used to create new instances of the Wheels structure
func NewWheels(specs map[string]int) *Wheels {
	rpm, ok := specs["rpm"]
	if !ok{ rpm=0; }
	size, ok := specs["size"]
	if !ok{ size=0; }
	count, ok := specs["count"]
	if !ok{ count=0; }
    return &Wheels{
        rpm: rpm,
		size: size,
		count: count,
    }
}

// SetRPM function is the setter method for the property rpm of the structure
func (w *Wheels) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	w.rpm = rpm
	return nil
}

// GetRPM function is the getter method for the property rpm of the structure
func (w *Wheels) GetRPM() int {
	return w.rpm
}

// SetSize function is the setter method for the property size of the structure
func (w *Wheels) SetSize(size int) error {
	if size < 0 {
		return fmt.Errorf("Negative size")
	}
	w.size = size
	return nil
}

// GetSize function is the getter method for the property size of the structure
func (w *Wheels) GetSize() int {
	return w.size
}

// GetCount function is the getter method for the property count of the structure
func (w *Wheels) GetCount() int {
	return w.count
}

// SetGear_2_1 is the function that is used to set the rpm's of the wheel to the half of the motor rpm's
func (w *Wheels) SetGear_2_1(motor ReferenceMotor) error {
	rpm:= motor.GetRPM()
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	w.rpm = rpm/2
	return nil
}

// Print function its used print the structure info.
func (w *Wheels) Print() {
	fmt.Printf("Wheels: \n\tRPM: %d, \n\tSize: %d, \n\tCount:%d\n", w.rpm, w.size, w.count)
}