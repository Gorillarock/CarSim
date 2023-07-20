package parts

import "fmt"

type Door struct {
	Open   bool
	Locked bool
}

func (d *Door) SetOpen(Open bool) error {
	if !d.Locked {
		d.Open = Open
		return nil
	}
	if d.Open {
		d.Open = Open
		return nil
	}
	return fmt.Errorf("Door is Locked, cannot Open.")
}

func (d *Door) IsOpen() bool {
	return d.Open
}

func (d *Door) SetLocked(Locked bool) {
	d.Locked = Locked
}

func (d *Door) IsLocked() bool {
	return d.Locked
}

func (d *Door) Print() {
	fmt.Printf("Door: \n\tOpen: %t, \n\tLocked: %t\n", d.Open, d.Locked)
}
