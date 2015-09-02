package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestLatitude_Marshal(t *testing.T) {

	testLatitude := `<Latitude datum="WGS84" unit="DEGREES">45</Latitude>`

	l := &Latitude{
		Datum: "WGS84",
		Unit:  "DEGREES",
		Value: 45.0,
	}

	x, err := xml.Marshal(&l)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testLatitude {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testLatitude, ""}, "\n=========\n"))
	}
}

func TestLatitude_String(t *testing.T) {

	var tests = []struct {
		s string
		x Latitude
	}{
		{
			`{"Datum":"WGS84","Unit":"DEGREES","Value":45}`,
			Latitude{
				Datum: "WGS84",
				Unit:  "DEGREES",
				Value: 45.0,
			},
		}, {
			`{"Value":45}`,
			Latitude{
				Value: 45.0,
			},
		}, {
			`{"Value":0}`,
			Latitude{},
		}}

	for _, test := range tests {
		if test.x.String() != test.s {
			t.Error(strings.Join([]string{"string mismatch:", test.x.String(), test.s, ""}, "\n****\n"))
		}
	}
}

func TestLatitude_Valid(t *testing.T) {

	var tests = []Latitude{
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

func TestLatitude_InValid(t *testing.T) {
	var tests = []Latitude{
		Latitude{
			Unit: "UNIT",
		},
		Latitude{
			Value: 100,
		},
		Latitude{
			Value: -100,
		},
	}

	for _, n := range tests {
		if err := Validate(n); err == nil {
			t.Errorf("latitude struct should be invalid: %s", n)
		}
	}
}
