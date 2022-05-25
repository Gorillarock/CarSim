# CarSim
Interfaces and Singletons

## Details for Assignment

Create a new branch for the following features.  Then make a PR into master branch when complete.

1) Organize the code into multiple files.
2) Create single gear function which sets wheels rpm to half of motor rpm ( 2(motor):1(wheels) )  (Assume wheels are not effected by inertia)

3) Add seatbelt functionality to car.
	Requirements for Seatbelt:
		a) Only needs 1 seatbelt for driver.
		b) seatbelt can be engauged or disengauged.
		c) seatbelt defaults to disengauged
		d) seabelt must be worn before motor rpm can be increased beyond 0.

4) Create Gearbox which allows for at least 2 different gears. (gear is the ratio between motor rpm and wheels rpm | example: 2:1 or 4:1)
	Requirements for Gearbox:
		a) Gearbox must have getter and setter.
		b) Gearbox must have at least 2 gears ( 2:1 and 4:1 ).
		c) Wheels rpm must be set based on current gear in relation to motor rpm.

5) Create gas which increases motor rpm by 1 every hundredth of a second and maxes the motor rpm out at 7000 rpm.
