// Package wheels is where all the different types of wheel's structures will be written.
package wheels

// ReferenceMotor represents the motor object reference needed by this package.
type ReferenceMotor interface {
	GetRPM() int
}