package stationxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestVoltage_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x Voltage
	}{
		{
			`<Voltage plusError="2" minusError="2" unit="VOLTS">1</Voltage>`,
			Voltage{Float: Float{Value: 1, Unit: "VOLTS", UncertaintyDouble: UncertaintyDouble{PlusError: 2, MinusError: 2}}},
		}, {
			`<Voltage>1</Voltage>`,
			Voltage{Float: Float{Value: 1.0}},
		}, {
			`<Voltage>0</Voltage>`,
			Voltage{},
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

func TestVoltage_Valid(t *testing.T) {

	var tests = []Voltage{
		Voltage{Float: Float{Value: 1, Unit: "VOLTS"}},
		Voltage{Float: Float{Value: 1.0}},
		Voltage{},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("voltage is invalid: %v (%s)", a, err)
		}
	}
}

func TestVoltage_InValid(t *testing.T) {

	var tests = []Voltage{
		Voltage{Float: Float{Unit: "volts"}},
		Voltage{Float: Float{UncertaintyDouble: UncertaintyDouble{PlusError: -1.0}}},
		Voltage{Float: Float{UncertaintyDouble: UncertaintyDouble{MinusError: -1.0}}},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("voltage should be invalid: %v", a)
		}
	}
}
