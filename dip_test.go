package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestDip_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x Dip
	}{
		{
			`<Dip unit="DEGREES" plusError="2" minusError="2">1</Dip>`,
			Dip{Value: 1, Unit: "DEGREES", PlusError: 2, MinusError: 2},
		}, {
			`<Dip>1</Dip>`,
			Dip{Value: 1.0},
		}, {
			`<Dip>0</Dip>`,
			Dip{},
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

func TestDip_Valid(t *testing.T) {

	var tests = []Dip{
		Dip{Value: 1, Unit: "DEGREES"},
		Dip{Value: 1.0},
		Dip{},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("azimuth is invalid: %v (%s)", a, err)
		}
	}
}

func TestDip_InValid(t *testing.T) {

	var tests = []Dip{
		Dip{Unit: "degrees"},
		Dip{Value: -91.0},
		Dip{Value: 91.0},
		Dip{PlusError: -1.0},
		Dip{MinusError: -1.0},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("azimuth should be invalid: %v", a)
		}
	}
}
