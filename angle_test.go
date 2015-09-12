package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestAngle_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x Angle
	}{
		{
			`<Angle unit="DEGREES" plusError="2" minusError="2">1</Angle>`,
			Angle{Value: 1, Unit: "DEGREES", PlusError: 2, MinusError: 2},
		}, {
			`<Angle>1</Angle>`,
			Angle{Value: 1.0},
		}, {
			`<Angle>0</Angle>`,
			Angle{},
		}}

	for _, test := range tests {

		s, err := xml.Marshal(&test.x)
		if err != nil {
			t.Error(err)
		}

		if (string)(s) != test.s {
			t.Error(strings.Join([]string{"mismatch:", (string)(s), test.s, ""}, "\n****\n"))
		}
	}
}

func TestAngle_Valid(t *testing.T) {

	var tests = []Angle{
		Angle{Value: 1, Unit: "DEGREES"},
		Angle{Value: 1.0},
		Angle{},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("angle is invalid: %v (%s)", a, err)
		}
	}
}

func TestAngle_InValid(t *testing.T) {

	var tests = []Angle{
		Angle{Unit: "degrees"},
		Angle{Value: -361.0},
		Angle{Value: 361.0},
		Angle{PlusError: -1.0},
		Angle{MinusError: -1.0},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("angle should be invalid: %v", a)
		}
	}
}
