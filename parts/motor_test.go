package parts

import "testing"

func TestMotor_SetWheelsRPM(t *testing.T) {
	type fields struct {
		On  bool
		Rpm int
	}
	type args struct {
		wm WheelsManipulator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "motor on - 2 rpm",
			fields: fields{
				On: true,
				Rpm: 2,
			},
			args: args{
				wm: mockWheelManipulator(Wheels{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Motor{
				On:  tt.fields.On,
				Rpm: tt.fields.Rpm,
			}
			if err := m.SetWheelsRPM(tt.args.wm); (err != nil) != tt.wantErr {
				t.Errorf("Motor.SetWheelsRPM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func mockWheelManipulator(wheels Wheels) WheelsManipulator {
	return &wheels
}
