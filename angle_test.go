package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestAngle_Marshal(t *testing.T) {

	var p float64 = 2
	var m float64 = -2

	var tests = []struct {
		s string
		x Angle
	}{
		{
			`<Angle unit="DEGREES" plusError="2" minusError="-2">1</Angle>`,
			Angle{
				Value:      1,
				Unit:       "DEGREES",
				PlusError:  &p,
				MinusError: &m,
			},
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

func TestAngle_String(t *testing.T) {

	var p float64 = 2
	var m float64 = -2

	var tests = []struct {
		s string
		x Angle
	}{
		{
			`{"Unit":"DEGREES","PlusError":2,"MinusError":-2,"Value":1}`,
			Angle{
				Value:      1,
				Unit:       "DEGREES",
				PlusError:  &p,
				MinusError: &m,
			},
		}, {
			`{"Value":1}`,
			Angle{Value: 1.0},
		}, {
			`{"Value":0}`,
			Angle{},
		}}

	for _, test := range tests {

		if test.x.String() != test.s {
			t.Error(strings.Join([]string{"string mismatch:", test.x.String(), test.s, ""}, "\n****\n"))
		}
	}
}

func TestAngle_Valid(t *testing.T) {

	var tests = []Angle{
		Angle{
			Value: 1,
			Unit:  "DEGREES",
		},
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
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("angle should be invalid: %v", a)
		}
	}
}
