package calc

import (
	"testing"
)

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

func TestCalculateInt(t *testing.T) {
	type args struct {
		input     string
		operation string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantFloat float32
	}{
		{
			name: "Test Add",
			args: args{"1 + 49", "add"},
			want: 50,
		},
		{
			name: "Test Add",
			args: args{"294724 + 324284", "add"},
			want: 619008,
		},
		{
			name: "Test Add",
			args: args{"1 + 1", "add"},
			want: 2,
		},
		{
			name: "Test Subtract",
			args: args{"3 - 1", "subtract"},
			want: 2,
		},
		{
			name: "Test Subtract",
			args: args{"379879 - 198", "subtract"},
			want: 379681,
		},
		{
			name: "Test Multiply",
			args: args{"37 * 98", "multiply"},
			want: 3626,
		},
		{
			name: "Test Multiply",
			args: args{"1 * 1", "multiply"},
			want: 1,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateInt(tt.args.input, tt.args.operation); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetermineNumTypes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test int",
			args: args{"1 * 1"},
			want: "int",
		},
		{
			name: "Test float",
			args: args{"1.0 * 1.0"},
			want: "float",
		},
		{
			name: "Test float 2",
			args: args{"3090.1 / 1"},
			want: "float",
		},
		{
			name: "Test float 3",
			args: args{"1 - 1.0989"},
			want: "float",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineNumTypes(tt.args.input); got != tt.want {
				t.Errorf("DetermineIntOrFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateFloat(t *testing.T) {
	type args struct {
		input     string
		operation string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Add",
			args: args{"0.5 + 0.5", "add"},
			want: 1.0,
		},
		{
			name: "Test Add",
			args: args{"12.913828 + 0.7871", "add"},
			want: 13.700928,
		},
		{
			name: "Test Add",
			args: args{"1.283203 + 0.1212", "add"},
			want: 1.404403,
		},
		{
			name: "Test Add",
			args: args{"0.404403 + 0.002", "add"},
			want: 0.406403,
		},
		{
			name: "Test Subtract",
			args: args{"0.5 - 0.1", "subtract"},
			want: 0.4,
		},
		{
			name: "Test Subtract",
			args: args{"1.232321 - 0.01", "subtract"},
			want: 1.222321,
		},
		{
			name: "Test Subtract",
			args: args{"1.222321 - 3.101", "subtract"},
			want: -1.878679,
		},
		{
			name: "Test Divide",
			args: args{"1.222321 / 3.101", "divide"},
			want: 0.3941699451789745,
		},
		{
			name: "Test Divide",
			args: args{"4000 / 0.333", "divide"},
			want: 12012.012012012012,
		},
		{
			name: "Test Divide",
			args: args{"5 / 4", "divide"},
			want: 1.25,
		},
		{
			name: "Test Divide",
			args: args{"5 / -4", "divide"},
			want: -1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFloat(tt.args.input, tt.args.operation); got != tt.want {
				t.Errorf("CalculateFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
