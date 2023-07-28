package car

import (
	"test_code/doors"
	"test_code/motor"
	"test_code/seatbelt"
	"test_code/wheels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCar_SetMotorAndWheels(t *testing.T) {
	type fields struct {
		wheels   *wheels.Wheels
		door     *doors.Door
		motor    *motor.Motor
		seatbelt *seatbelt.Seatbelt
	}
	type args struct {
		motorRpm int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				wheels: &wheels.Wheels{},
				door:   &doors.Door{},
				motor:  &motor.Motor{},
				seatbelt: &seatbelt.Seatbelt{
					Engaged: true,
				},
			},
			args: args{
				motorRpm: 4,
			},
			wantErr: false,
		},

		{
			name: "Negative - seatbelt not engaged",
			fields: fields{
				wheels:   &wheels.Wheels{},
				door:     &doors.Door{},
				motor:    &motor.Motor{},
				seatbelt: &seatbelt.Seatbelt{},
			},
			args: args{
				motorRpm: 4,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				wheels:   tt.fields.wheels,
				door:     tt.fields.door,
				motor:    tt.fields.motor,
				seatbelt: tt.fields.seatbelt,
			}
			if err := c.SetMotorAndWheels(tt.args.motorRpm); (err != nil) != tt.wantErr {
				t.Errorf("Car.SetMotorAndWheels() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		assert := assert.New(t)
		assert.Equal(tt.fields.wheels.GetRPM(), tt.fields.motor.GetRPM()/2)
	}
}
