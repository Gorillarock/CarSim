// Package main is where the main logic of the program is written.
package main

// Run with: go run main.go

import (
	"fmt"
	"time"
	cr "test_code/pkg/cars"
)

// main function is where the main logic of the program runs.
func main() {
	// Get car and internal parts objects
	car := &cr.Car{}
	wheels := car.WheelManipulator()
	gearbox := car.GearboxManipulator()
	seatbelt := car.SeatbeltManipulator()
	door := car.DoorManipulator()
	motor := car.MotorManipulator()
	// Run sequence
	car.Print()

	fmt.Println("Turning On Motor")

	door.SetLocked(false)
	door.SetOpen(true)
	fmt.Println("** Getting Into Car. **")
	door.SetOpen(false)
	door.SetLocked(true)
	motor.SetOn(true)

	fmt.Println("** Car Started. **")
	car.Print()

	fmt.Println("** Accelerating Car... **")
	err:= car.SetMotorRpm(100)
	if err!=nil{ fmt.Println("Ups: "+err.Error()) }
	fmt.Println("** Putting On Seatbelt **")
	seatbelt.SetEngaged(true)

	for i := 0; i < 3; i++ {
		fmt.Println("** Setting Up GearBox And Starting to Accelerate... **")
		gearbox.SetGear(i)
		gearbox.Print()
		ticker := time.NewTicker(1 * time.Millisecond)
		errCh := make(chan error)
		go func() {
			start := time.Now()
			for range ticker.C {
				if err := car.IncreaseSpeed(1,7000); err != nil {
					errCh <- err
					elapsed := time.Since(start)
					fmt.Printf("Took %s seconds\n", elapsed)
					return
				}
			}
		}()
		err = <-errCh
		wheels.Print()
		if i<2{ fmt.Println("** Stopping Car And Starting Over With A Different Gearbox. **") }
		car.SetMotorRpm(0)
	}

	fmt.Println("** Stopping Car. **")
	motor.SetOn(false)
	car.Print()

}
