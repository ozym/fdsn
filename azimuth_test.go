package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestAzimuth_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x Azimuth
	}{
		{
			`<Azimuth unit="DEGREES" plusError="2" minusError="2">1</Azimuth>`,
			Azimuth{Value: 1, Unit: "DEGREES", PlusError: 2, MinusError: 2},
		}, {
			`<Azimuth>1</Azimuth>`,
			Azimuth{Value: 1.0},
		}, {
			`<Azimuth>0</Azimuth>`,
			Azimuth{},
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

func TestAzimuth_Valid(t *testing.T) {

	var tests = []Azimuth{
		Azimuth{Value: 1, Unit: "DEGREES"},
		Azimuth{Value: 1.0},
		Azimuth{},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("azimuth is invalid: %v (%s)", a, err)
		}
	}
}

func TestAzimuth_InValid(t *testing.T) {

	var tests = []Azimuth{
		Azimuth{Unit: "degrees"},
		Azimuth{Value: -1.0},
		Azimuth{Value: 361.0},
		Azimuth{PlusError: -1.0},
		Azimuth{MinusError: -1.0},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("azimuth should be invalid: %v", a)
		}
	}
}
