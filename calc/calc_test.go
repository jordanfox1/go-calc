package calc

import "testing"

func TestDetermineOperation(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Addition",
			args: args{input: "+"},
			want: "add",
		},
		{
			name: "Test Addition",
			args: args{"5 + 300"},
			want: "add",
		},
		{
			name: "Test Subtraction",
			args: args{"5 - 300"},
			want: "subtract",
		},
		{
			name: "Test Multiplication",
			args: args{"5 * 20"},
			want: "multiply",
		},
		{
			name: "Test Multiplication",
			args: args{"*"},
			want: "multiply",
		},
		{
			name: "Test Division",
			args: args{"5 / 20"},
			want: "divide",
		},
		{
			name: "Test Invalid",
			args: args{""},
			want: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineOperation(tt.args.input); got != tt.want {
				t.Errorf("DetermineOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
