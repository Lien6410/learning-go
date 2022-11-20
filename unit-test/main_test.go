package main

import "testing"

// file name has to be xxx_test.go

// function name has to be TestXxxxx

// cli: go test
// cli: go test -v
// cli: go test -cover
// cli: go test -coverprofile=coverage.out && go tool cover -html=coverage.out

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 2.0, 50.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Expect an erro but did not get one.")
			}
		} else {
			if err != nil {
				t.Error("Expect no erro but got one.", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expect %f but got %f", tt.expected, got)
		}

	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Errorf("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Errorf("Did not get an error when we should have")
// 	}
// }
