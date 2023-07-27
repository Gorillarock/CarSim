// Package doors is where all the different types of doors's structures will be written.
package doors

// DoorManipulator represents door's objects main methods.
type DoorManipulator interface {
	SetOpen(bool) error
	SetLocked(bool)
	IsOpen() bool
	IsLocked() bool
	Print()
}