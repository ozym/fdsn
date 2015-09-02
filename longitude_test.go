package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestLongitude_Marshal(t *testing.T) {

	testLongitude := `<Longitude datum="WGS84" unit="DEGREES">45</Longitude>`

	l := &Longitude{
		Datum: "WGS84",
		Unit:  "DEGREES",
		Value: 45.0,
	}

	x, err := xml.Marshal(&l)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testLongitude {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testLongitude, ""}, "\n=========\n"))
	}
}

func TestLongitude_String(t *testing.T) {

	var tests = []struct {
		s string
		x Longitude
	}{
		{
			`{"Datum":"WGS84","Unit":"DEGREES","Value":45}`,
			Longitude{
				Datum: "WGS84",
				Unit:  "DEGREES",
				Value: 45.0,
			},
		}, {
			`{"Value":45}`,
			Longitude{
				Value: 45.0,
			},
		}, {
			`{"Value":0}`,
			Longitude{},
		}}

	for _, test := range tests {
		if test.x.String() != test.s {
			t.Error(strings.Join([]string{"string mismatch:", test.x.String(), test.s, ""}, "\n****\n"))
		}
	}
}

func TestLongitude_Valid(t *testing.T) {

	var tests = []Longitude{
		{
			Datum: "WGS84",
			Unit:  "DEGREES",
			Value: 45.0,
		}, {
			Value: 45.0,
		}}

	for _, l := range tests {
		if err := Validate(l); err != nil {
			t.Errorf("latitude struct is not valid: %s", err)
		}
	}
}

func TestLongitude_InValid(t *testing.T) {
	var tests = []Longitude{
		Longitude{
			Unit: "UNIT",
		},
		Longitude{
			Value: 200,
		},
		Longitude{
			Value: -200,
		},
	}

	for _, n := range tests {
		if err := Validate(n); err == nil {
			t.Errorf("latitude struct should be invalid: %s", n)
		}
	}
}
