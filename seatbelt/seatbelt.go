package seatbelt

import "fmt"

type Seatbelt struct {
	Engaged bool
}

type SeatbeltManipulator interface {
	SetEngaged(bool)
	IsEngaged() bool
	Print()
}

func (s *Seatbelt) SetEngaged(engaged bool) {
	s.Engaged = engaged
}

func (s *Seatbelt) IsEngaged() bool {
	return s.Engaged
}

func (s *Seatbelt) Print() {
	fmt.Printf("Seatbelt: \n\tEngaged: %t\n", s.Engaged)
}
