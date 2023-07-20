package main

import (
	"CarSim/car"
	"CarSim/parts"
	"fmt"
)

func main() {
	car := &car.Car{}
	wheels := car.WheelManipulator()
	door := car.DoorManipulator()
	motor := car.MotorManipulator()
	seatBelt := car.SeatbeltManipulator()

	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)
	seatBelt.SetState(parts.SeatbeltStateEngauged)
	motor.SetRPM(100, seatBelt)
	motor.SetWheelsRPM(wheels)

	car.Print()

}
