package stationxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSecond_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x Second
	}{
		{
			`<Second plusError="2" minusError="2" unit="SECONDS">1</Second>`,
			Second{Float: Float{Value: 1, Unit: "SECONDS", UncertaintyDouble: UncertaintyDouble{PlusError: 2, MinusError: 2}}},
		}, {
			`<Second>1</Second>`,
			Second{Float: Float{Value: 1.0}},
		}, {
			`<Second>0</Second>`,
			Second{},
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

func TestSecond_Valid(t *testing.T) {

	var tests = []Second{
		Second{Float: Float{Value: 1, Unit: "SECONDS"}},
		Second{Float: Float{Value: 1.0}},
		Second{},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("second is invalid: %v (%s)", a, err)
		}
	}
}

func TestSecond_InValid(t *testing.T) {

	var tests = []Second{
		Second{Float: Float{Unit: "seconds"}},
		Second{Float: Float{UncertaintyDouble: UncertaintyDouble{PlusError: -1.0}}},
		Second{Float: Float{UncertaintyDouble: UncertaintyDouble{MinusError: -1.0}}},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("second should be invalid: %v", a)
		}
	}
}
