package main

import (
	"fmt"
	"test_code/car"
)

func main() {
	car := &car.Car{}

	door := car.DoorManipulator()
	motor := car.MotorManipulator()
	seatbelt := car.SeatbeltManipulator()

	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)
	seatbelt.SetEngaged(true)

	car.SetMotorAndWheels(4)

	car.Print()

}
