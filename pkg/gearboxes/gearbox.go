// Package gearbox is where all the different types of gearbox's structures will be written.
package gearbox

import (
	"fmt"
)

const (
	Gear1 int = iota
	Gear1_2
	Gear1_4
)

// GearBox represents a gearbox type.
type GearBox struct {
	gear   int
}

// NewGearBox function is used to create new instances of the GearBox structure
func NewGearBox(specs map[string]interface{}) *GearBox {
	gear:=Gear1
	_, ok := specs["gear"]
	if ok{ 
		gear, ok = specs["gear"].(int)
	}
	if !ok{ gear=Gear1; }
    return &GearBox{
        gear: gear,
    }
}

// SetGear function is the setter method for the property gear of the structure
func (d *GearBox) SetGear(gear int) {
	d.gear = gear
}

// GetGear function is the getter method for the property gear of the structure
func (d *GearBox) GetGear() int {
	return d.gear
}

// GetGearRatio function is used to return the gear ratio of the gearbox.
func (d *GearBox) GetGearRatio() float64 {
	switch d.gear {
	case Gear1_2:
		return 1.0/2
	case Gear1_4:
		return 1.0/4
	default:
		return 1
	}
}

// getGearRatioStr function is used to return the gear ratio of the gearbox as string.
func (d *GearBox) getGearRatioStr() string {
	switch d.gear {
	case Gear1_2:
		return "1/2"
	case Gear1_4:
		return "1/4"
	default:
		return "1"
	}
}

// Print function its used print the structure info.
func (d *GearBox) Print() {
	fmt.Printf("GearBox: \n\tGear: %s\n",d.getGearRatioStr())
}
