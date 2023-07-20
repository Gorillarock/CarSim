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
	SetRPM(int) error
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
