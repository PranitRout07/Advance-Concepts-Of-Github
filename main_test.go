package main

import (
	"log"
	"testing"
)

var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"valid-data", 90.0, 2.0, 45.0, false},
	{"valid-data",1000.0,20.0,50.0,false},
	{"invalid-data", 89.0, 0.0, 0.0, true},
	{"valid-data", 500.0, 2.0, 250.0, false},
	{"valid-data", 200.0, 4.0, 50.0, false},
	{"valid-data", 600.0, 2.0, 300.0, false},
	{"valid-data", 550.0, 11.0, 50.0, false},
	{"valid-data", 250.0, 2.0, 125.0, false},
	{"valid-data", 800.0, 2.0, 400.0, false},
	{"valid-data", 12.0, 2.0, 6.0, false},

}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Not Expected an error but did get one")
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
		log.Println(tt.name)

	}
}