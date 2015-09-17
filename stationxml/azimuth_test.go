package stationxml

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
			`<Azimuth plusError="2" minusError="2" unit="DEGREES">1</Azimuth>`,
			Azimuth{Float: Float{Value: 1, Unit: "DEGREES", UncertaintyDouble: UncertaintyDouble{PlusError: 2, MinusError: 2}}},
		}, {
			`<Azimuth>1</Azimuth>`,
			Azimuth{Float: Float{Value: 1.0}},
		}, {
			`<Azimuth>0</Azimuth>`,
			Azimuth{},
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

func TestAzimuth_Valid(t *testing.T) {

	var tests = []Azimuth{
		Azimuth{Float: Float{Value: 1.0, Unit: "DEGREES"}},
		Azimuth{Float: Float{Value: 1.0}},
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
		Azimuth{Float: Float{Value: 1.0, Unit: "degrees"}},
		Azimuth{Float: Float{Value: -1.0}},
		Azimuth{Float: Float{Value: 361.0}},
		Azimuth{Float: Float{UncertaintyDouble: UncertaintyDouble{PlusError: -1.0}}},
		Azimuth{Float: Float{UncertaintyDouble: UncertaintyDouble{MinusError: -1.0}}},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("azimuth should be invalid: %v", a)
		}
	}
}
