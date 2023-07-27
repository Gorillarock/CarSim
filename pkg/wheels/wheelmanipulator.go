// Package wheels is where all the different types of wheel's structures will be written.
package wheels

// WheelsManipulator represents wheel's objects main methods.
type WheelsManipulator interface {
	SetRPM(int) error
	SetSize(int) error
	GetRPM() int
	GetSize() int
	GetCount() int
	SetGear_2_1(ReferenceMotor) error
	Print()
}