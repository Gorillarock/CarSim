// Package wheels is where all the different types of wheel's structures will be written.
package wheels

// Run with: go test -v -run=TestSetGear_2_1 

import(  
	"math/rand"
	"testing"
	"time"
)

// mockMotor represents the structure of a motor
type mockMotor struct {
	Rpm int
}

// GetRPM function get the simulated rpm's of a motor
func (m *mockMotor) GetRPM() int {
	m.Rpm = rand.Intn(1000)
	return m.Rpm
}

// TestSetGear_2_1 function run the test for the SetGear_2_1 method in wheels
func TestSetGear_2_1(t *testing.T) {
	wheel := NewWheels(nil)
	motor := &mockMotor{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wheel.SetGear_2_1(motor)
		wrpm := wheel.GetRPM()
		t.Logf("Test: motor RPM: %d, got: %d",motor.Rpm,wrpm)
		if wrpm != (motor.Rpm/2) {
			t.Errorf("Incorrect wheel RPM, got: %d, want: %d.",wrpm,(motor.Rpm/2))
		}
	}
}