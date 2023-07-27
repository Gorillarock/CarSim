// Package cars is where all the different types of car's structures will be written.
package cars

import (
	wh "test_code/pkg/wheels"
	mt "test_code/pkg/motors"
	do "test_code/pkg/doors"
	stb "test_code/pkg/seatbelts"
	gb "test_code/pkg/gearboxes"
)

// CarManipulator represents car's objects main methods.
type CarManipulator interface {
	WheelManiputlator() wh.WheelsManipulator
	DoorManipulator() do.DoorManipulator
	MotorManipulator() mt.MotorManipulator
	SeatbeltManipulator() stb.SeatbeltManipulator
	GearboxManipulator() gb.GearboxManipulator
	SetMotorRpm(rpm int) error
	IncreaseSpeed(rpm,maxrpm int) error
	Print()
}