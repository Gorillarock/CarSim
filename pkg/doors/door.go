// Package doors is where all the different types of doors's structures will be written.
package doors

import (
	"fmt"
)

// Door represents a door type.
type Door struct {
	open   bool
	locked bool
}

// NewDoor function is used to create new instances of the Door structure
func NewDoor(specs map[string]bool) *Door {
	open, ok := specs["open"]
	if !ok{ open=false; }
	locked, ok := specs["locked"]
	if !ok{ locked=false; }
    return &Door{
        open: open,
		locked: locked,
    }
}

// SetOpen function is the setter method for the property open of the structure
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

// IsOpen function is the getter method for the property open of the structure
func (d *Door) IsOpen() bool {
	return d.open
}

// SetLocked function is the setter method for the property locked of the structure
func (d *Door) SetLocked(locked bool) {
	d.locked = locked
}

// IsLocked function is the getter method for the property locked of the structure
func (d *Door) IsLocked() bool {
	return d.locked
}

// Print function its used print the structure info.
func (d *Door) Print() {
	fmt.Printf("Door: \n\tOpen: %t, \n\tLocked: %t\n", d.open, d.locked)
}
