// Package motors is where all the different types of motor's structures will be written.
package motors

// MotorManipulator represents motor's objects main methods.
type MotorManipulator interface {
	SetOn(bool)
	SetRPM(int) error
	IsOn() bool
	GetRPM() int
	Print()
}

