// Package seatbelt is where all the different types of seatbelt's structures will be written.
package seatbelt

import (
	"fmt"
)

// Seatbelt represents a seatbelt type.
type Seatbelt struct {
	engaged bool
}

// NewSeatbelt function is used to create new instances of the Seatbelt structure
func NewSeatbelt(specs map[string]bool) *Seatbelt {
	engaged, ok := specs["engaged"]
	if !ok{ engaged=false; }
    return &Seatbelt{
        engaged: engaged,
    }
}

// SetEngaged function is the setter method for the property engaged of the structure
func (m *Seatbelt) SetEngaged(engaged bool) {
	m.engaged = engaged
}

// IsEngaged function is the getter method for the property engaged of the structure
func (m *Seatbelt) IsEngaged() bool {
	return m.engaged
}

// Print function its used print the structure info.
func (m *Seatbelt) Print() {
	fmt.Printf("Seatbelt: \n\tEngaged: %t\n", m.engaged)
}