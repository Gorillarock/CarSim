package parts

import "fmt"

type SeatbeltState string

var (
	SeatbeltStateEngauged    SeatbeltState = "engauged"
	SeatbeltStateDisengauged SeatbeltState = "disengauged"
)

type Seatbelt struct {
	State *SeatbeltState
}

func (s *Seatbelt) GetState() SeatbeltState {
	if s.State == nil {
		return SeatbeltStateDisengauged
	}
	return *s.State
}

func (s *Seatbelt) SetState(state SeatbeltState) error {
	s.State = &state
	return nil
}

func (s *Seatbelt) Print() {
	fmt.Printf("Seatbelt: \n\tStatus: %s\n", s.GetState())
}
