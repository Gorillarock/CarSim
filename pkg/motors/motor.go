// Package motors is where all the different types of motor's structures will be written.
package motors

import (
	"fmt"
)

// Motor represents a motor type.
type Motor struct {
	on  bool
	rpm int
}

// NewMotor function is used to create new instances of the Motor structure
func NewMotor(specs map[string]interface{}) *Motor {
	on:=false
	_, ok := specs["on"]
	if ok{ 
		on, ok = specs["on"].(bool)
	}
	if !ok{ on=false; }
	rpm:=0
	_, ok = specs["rpm"]
	if ok{ 
		rpm, ok = specs["rpm"].(int)
	}
	if !ok{ rpm=0; }
    return &Motor{
        on: on,
		rpm: rpm,
    }
}

// SetOn function is the setter method for the property on of the structure
func (m *Motor) SetOn(on bool) {
	m.on = on
}

// IsOn function is the getter method for the property on of the structure
func (m *Motor) IsOn() bool {
	return m.on
}

// SetRPM function is the setter method for the property rpm of the structure
func (m *Motor) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	m.rpm = rpm
	return nil
}

// GetRPM function is the getter method for the property rpm of the structure
func (m *Motor) GetRPM() int {
	return m.rpm
}

// Print function its used print the structure info.
func (m *Motor) Print() {
	fmt.Printf("Motor: \n\tOn: %t, \n\tRPM: %d\n", m.on, m.rpm)
}