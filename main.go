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

	gearbox :=car.GearboxManipulator()
	gearbox.SetMotor(motor)
	gearbox.SetWheels(wheels)

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
	gearbox.SetGear("4:1")
	gearbox.UpdateWheelsRPM()
	car.Print()
	fmt.Println("** Change Gear**")
	motor.SetRPM(900, seatBelt)
	gearbox.SetGear("2:1")
	gearbox.UpdateWheelsRPM()
	
	car.Print()
	

}
