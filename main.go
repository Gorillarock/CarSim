package main

import (
	"CarSim/car"
	"CarSim/parts"
	"fmt"
	"time"
)

func main() {
	car := &car.Car{}
	wheels := car.WheelManipulator()
	door := car.DoorManipulator()
	motor := car.MotorManipulator()
	seatBelt := car.SeatbeltManipulator()

	gearbox := car.GearboxManipulator()

	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)
	seatBelt.SetState(parts.SeatbeltStateEngauged)
	fmt.Println("** Increasing RPM directly to motor**")
	motor.SetRPM(100, seatBelt)
	motor.SetWheelsRPM(wheels)
	car.Print()
	fmt.Println("** Change Gear**")
	gearbox.SetGear(4)
	gearbox.UpdateWheelsRPM()
	car.Print()
	fmt.Println("** Change Gear**")
	motor.SetRPM(900, seatBelt)
	gearbox.SetGear(2)
	gearbox.UpdateWheelsRPM()
	car.Print()

	maxReachedChan := make(chan bool)

	go increaseSpeed(motor,seatBelt, 1, 7000, maxReachedChan)

	<-maxReachedChan

	fmt.Println("** Increased RPM until reaches 7000**")
	car.Print()
}

func increaseSpeed(motor parts.MotorManipulator,
	seatBelt parts.SeatbeltManipulator,
	increment, max int,
	maxReached chan bool) {
	for rpm := motor.GetRPM(); rpm <= max; rpm += increment {
		motor.SetRPM(rpm, seatBelt)
		if rpm >= max {
			maxReached <- true
		}
		time.Sleep(time.Millisecond)
	}
}
