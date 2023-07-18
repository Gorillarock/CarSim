# CarSim
Interfaces and Singletons

## Details for Assignment

Create a new branch for the following features.  Then make a PR into master branch when complete.

### Go

1) Organize the code into multiple files.
2) Create single gear function which sets wheels rpm to half of motor rpm ( 2(motor):1(wheels) )  (Assume wheels are not effected by inertia)
3) Create a test for the function created in step 2.

4) Add seatbelt functionality to car.
	Requirements for Seatbelt:
		a) Only needs 1 seatbelt for the vehicle (assume it is for the driver).
		b) seatbelt can be engauged or disengauged (2 states only).
		c) seatbelt defaults to disengauged
		d) seabelt must be engauged before motor rpm can be increased beyond 0.

5) Create Gearbox which allows for at least 2 different gears. (gear is the ratio between motor rpm and wheels rpm | example: one gear may be "2:1" and another may be "4:1")
	Requirements for Gearbox:
		a) Gearbox must have getter and setter functions.
		b) Gearbox must have at least 2 gears ( 2:1 and 4:1 ).
		c) Wheels rpm must be set based on current gear in relation to motor rpm.
		d) NOTE: A "Car" will have a gearbox as a component.  All other components should not interact without the car itself.
			IE - Gearbox does not directly manipulate wheels.  Motor does not directly manipulate gearbox.
   			All linkage between components happens at the "Car" level.

6) Create a function called "increaseSpeed" which increases motor rpm by some amount (param 1), until some maximum amount (param 2) is acheived.  Provide the ability for the function to notify the caller in some way that it has hit the maximum.  Call the function so that the rpm of the engine is increased by 1 every millisecond and maxes out when the motor rpm reaches 7000.  Then print the rpm of the wheels.


### Docker

1) Create a Dockerfile for the program.
   	Requirements for Dockerfile:
   		a) Should use golang:alpine as builder.
   		b) Should be able to run the docker image as a container using "docker run -it --rm <name_of_image>" and have it run the main function for the program.


### Makefile (Some Go)

1) Create a basic makefile which will run all of the repo's tests.
   	Requirements for Makefile:
   		a) Command for test execution should be "make test"
   		b) Command should first handle clearing any elements in the testcache, before running the actual tests.
