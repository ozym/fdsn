package stationxml

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
			`<Dip plusError="2" minusError="2" unit="DEGREES">1</Dip>`,
			Dip{Float: Float{Value: 1, Unit: "DEGREES", UncertaintyDouble: UncertaintyDouble{PlusError: 2, MinusError: 2}}},
		}, {
			`<Dip>1</Dip>`,
			Dip{Float: Float{Value: 1.0}},
		}, {
			`<Dip>0</Dip>`,
			Dip{},
		}}

	for _, test := range tests {

		s, err := xml.Marshal(test.x)
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
		Dip{Float: Float{Value: 1, Unit: "DEGREES"}},
		Dip{Float: Float{Value: 1.0}},
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
		Dip{Float: Float{Unit: "degrees"}},
		Dip{Float: Float{Value: -91.0}},
		Dip{Float: Float{Value: 91.0}},
		Dip{Float: Float{UncertaintyDouble: UncertaintyDouble{PlusError: -1.0}}},
		Dip{Float: Float{UncertaintyDouble: UncertaintyDouble{MinusError: -1.0}}},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("azimuth should be invalid: %v", a)
		}
	}
}
