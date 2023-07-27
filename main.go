package main

import (
	"fmt"
	"test_code/car"
)

func main() {
	car := &car.Car{}
	//wheels := car.WheelManipulator()
	door := car.DoorManipulator()
	motor := car.MotorManipulator()

	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)

}
