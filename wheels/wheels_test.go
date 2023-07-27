package wheels_test

import (
	"test_code/car"
	"test_code/motor"
	"test_code/wheels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWheels_SetHalfMotorRPM(t *testing.T) {
	type fields struct {
		Rpm   int
		Size  int
		Count int
	}
	type args struct {
		m motor.MotorManipulator
	}
	car := &car.Car{}
	car.MotorManipulator().SetRPM(2)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			fields: fields{
				Rpm:   1,
				Size:  20,
				Count: 4,
			},
			args: args{
				m: car.MotorManipulator(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &wheels.Wheels{
				Rpm:   tt.fields.Rpm,
				Size:  tt.fields.Size,
				Count: tt.fields.Count,
			}
			if err := w.SetHalfMotorRPM(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Wheels.SetHalfMotorRPM() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert := assert.New(t)
			assert.Equal(tt.fields.Rpm, tt.args.m.GetRPM()/2)
		})
	}
}
