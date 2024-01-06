package calculator_test

import (
	"01testing/calculator"
	"testing"
)

func TestDivide_WhenNumsAre10And5(t *testing.T) {
	got := calculator.Divide(10, 5)

	if got != 2 {
		t.Errorf("Divide() = %v, want %v", got, 2)
	}
}

func TestDivide_WhenNumsAre5And3(t *testing.T) {

	got := calculator.Divide(5, 3)

	if got != 1 {
		t.Errorf("Divide() = %v, want %v", got, 1)
	}
}

/* Following is called as parameterized testing. All the above testcases which just differs in parameters(basically, repetative testcases) are put into a single test. */
func TestDivide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "check when numerator 10 and denomenator is 5", args: args{a: 10, b: 5}, want: 2},
		{"check when numerator 5 and denomenator is 3", args{a: 5, b: 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := calculator.Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
