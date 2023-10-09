package cli

import "testing"

func Test_validateInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test valid",
			args: args{"1 - 1.0989"},
			want: "1 - 1.0989",
		},
		{
			name: "Test valid",
			args: args{"1 + 1"},
			want: "1 + 1",
		},
		{
			name: "Test valid",
			args: args{"19 / 1.1"},
			want: "19 / 1.1",
		},
		{
			name:    "Test invalid",
			args:    args{"1 + 1 + 2"},
			want:    "invalid",
			wantErr: true,
		},
		{
			name:    "Test invalid",
			args:    args{"1 + 1 / 2"},
			want:    "invalid",
			wantErr: true,
		},
		{
			name:    "Test invalid",
			args:    args{"1 + 1 / 2 * 8 0000"},
			want:    "invalid",
			wantErr: true,
		},
		{
			name:    "Test invalid",
			args:    args{"test "},
			want:    "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
