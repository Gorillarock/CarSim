// Package gearbox is where all the different types of gearbox's structures will be written.
package gearbox

// GearboxManipulator represents gearbox's objects main methods.
type GearboxManipulator interface {
	SetGear(int)
	GetGear() int
	GetGearRatio() float64
	Print()
}
