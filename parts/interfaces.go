package parts

type DoorManipulator interface {
	SetOpen(bool) error
	SetLocked(bool)
	IsOpen() bool
	IsLocked() bool
	Print()
}

type MotorManipulator interface {
	SetOn(bool)
	SetRPM(rpm int, sb SeatbeltManipulator) error
	IsOn() bool
	GetRPM() int
	SetWheelsRPM(wm WheelsManipulator) error
	Print()
}

type WheelsManipulator interface {
	SetRPM(int) error
	SetSize(int) error
	GetRPM() int
	GetSize() int
	GetCount() int
	Print()
}

type SeatbeltManipulator interface {
	SetState(state SeatbeltState) error
	GetState() SeatbeltState
	Print()
}
type GearboxManipulator interface {
	NewGearbox(motor MotorManipulator, wheels WheelsManipulator) *Gearbox
	SetGear(gear string) error
	UpdateWheelsRPM()
}
