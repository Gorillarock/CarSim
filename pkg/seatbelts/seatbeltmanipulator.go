// Package seatbelt is where all the different types of seatbelt's structures will be written.
package seatbelt

// SeatbeltManipulator represents seatbelt's objects main methods.
type SeatbeltManipulator interface {
	SetEngaged(bool)
	IsEngaged() bool
	Print()
}