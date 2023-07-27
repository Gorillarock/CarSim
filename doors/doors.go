package doors

import (
	"fmt"
)

type Door struct {
	open   bool
	Locked bool
}

type DoorManipulator interface {
	SetOpen(bool) error
	SetLocked(bool)
	IsOpen() bool
	IsLocked() bool
	Print()
}

func (d *Door) SetOpen(open bool) error {
	if !d.Locked {
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
	d.Locked = locked
}

func (d *Door) IsLocked() bool {
	return d.Locked
}

func (d *Door) Print() {
	fmt.Printf("Door: \n\tOpen: %t, \n\tLocked: %t\n", d.open, d.Locked)
}
