package motor

import (
	"fmt"
)

type Motor struct {
	on  bool
	rpm int
}

type MotorManipulator interface {
	SetOn(bool)
	SetRPM(int) error
	IsOn() bool
	GetRPM() int
	Print()
}

func (m *Motor) SetOn(on bool) {
	m.on = on
}

func (m *Motor) IsOn() bool {
	return m.on
}

func (m *Motor) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	m.rpm = rpm
	return nil
}

func (m *Motor) GetRPM() int {
	return m.rpm
}

func (m *Motor) Print() {
	fmt.Printf("Motor: \n\tOn: %t, \n\tRPM: %d\n", m.on, m.rpm)
}
