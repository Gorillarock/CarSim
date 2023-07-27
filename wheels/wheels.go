package wheels

import (
	"fmt"
	"test_code/motor"
)

type Wheels struct {
	Rpm   int
	Size  int
	Count int
}

type WheelsManipulator interface {
	SetRPM(int) error
	SetSize(int) error
	GetRPM() int
	GetSize() int
	GetCount() int
	SetHalfMotorRPM(motor.MotorManipulator) error
	Print()
}

func (w *Wheels) SetRPM(rpm int) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	w.Rpm = rpm
	return nil
}

func (w *Wheels) SetHalfMotorRPM(m motor.MotorManipulator) error {
	motorRpm := m.GetRPM()
	rpm := int(motorRpm / 2)
	err := w.SetRPM(rpm)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wheels) GetRPM() int {
	return w.Rpm
}

func (w *Wheels) SetSize(size int) error {
	if size < 0 {
		return fmt.Errorf("Negative size")
	}
	w.Size = size
	return nil
}

func (w *Wheels) GetSize() int {
	return w.Size
}

func (w *Wheels) GetCount() int {
	return w.Count
}

func (w *Wheels) Print() {
	fmt.Printf("Wheels: \n\tRPM: %d, \n\tSize: %d, \n\tCount:%d\n", w.Rpm, w.Size, w.Count)
}
