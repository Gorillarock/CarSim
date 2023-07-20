package car

import "CarSim/parts"

type CarManipulator interface {
	WheelManiputlator() parts.WheelsManipulator
	DoorManipulator() parts.DoorManipulator
	MotorManipulator() parts.MotorManipulator
	SeatbeltManipulator() parts.SeatbeltManipulator
	GearboxManipulator() parts.GearboxManipulator
	Print()
}
