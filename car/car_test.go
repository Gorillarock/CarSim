package car

import (
	"test_code/doors"
	"test_code/gearbox"
	"test_code/motor"
	"test_code/seatbelt"
	"test_code/wheels"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCar_SetMotorRPM(t *testing.T) {
	type fields struct {
		wheels   *wheels.Wheels
		door     *doors.Door
		motor    *motor.Motor
		seatbelt *seatbelt.Seatbelt
		gearbox  *gearbox.Gearbox
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
				gearbox: &gearbox.Gearbox{
					Gear: 2,
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
				gearbox: &gearbox.Gearbox{
					Gear: 2,
				},
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
			if err := c.SetMotorRPM(tt.args.motorRpm); (err != nil) != tt.wantErr {
				t.Errorf("Car.SetMotorRPM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		assert := assert.New(t)
		assert.Equal(tt.fields.wheels.GetRPM(), tt.fields.motor.GetRPM()/tt.fields.gearbox.Gear)
	}
}

func TestCar_ChangeGear(t *testing.T) {
	type fields struct {
		wheels   *wheels.Wheels
		door     *doors.Door
		motor    *motor.Motor
		seatbelt *seatbelt.Seatbelt
		gearbox  *gearbox.Gearbox
	}
	type args struct {
		gear int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				gearbox: &gearbox.Gearbox{
					Gear: 2,
				},
			},
			args: args{
				gear: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				wheels:   tt.fields.wheels,
				door:     tt.fields.door,
				motor:    tt.fields.motor,
				seatbelt: tt.fields.seatbelt,
				gearbox:  tt.fields.gearbox,
			}
			c.ChangeGear(tt.args.gear)
		})
	}
}

func TestCar_SetWheels(t *testing.T) {
	type fields struct {
		wheels   *wheels.Wheels
		door     *doors.Door
		motor    *motor.Motor
		seatbelt *seatbelt.Seatbelt
		gearbox  *gearbox.Gearbox
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Car{
				wheels:   tt.fields.wheels,
				door:     tt.fields.door,
				motor:    tt.fields.motor,
				seatbelt: tt.fields.seatbelt,
				gearbox:  tt.fields.gearbox,
			}
			if err := c.SetWheels(); (err != nil) != tt.wantErr {
				t.Errorf("Car.SetWheels() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
