package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestLatitude_Marshal(t *testing.T) {

	testLatitude := `<Latitude unit="DEGREES" datum="WGS84">45</Latitude>`

	l := &Latitude{
		Datum: "WGS84",
		LatitudeBase: LatitudeBase{
			Unit:  "DEGREES",
			Value: 45.0,
		},
	}

	x, err := xml.Marshal(&l)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testLatitude {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testLatitude, ""}, "\n=========\n"))
	}
}

func TestLatitude_Valid(t *testing.T) {

	var tests = []Latitude{
		{
			Datum: "WGS84",
			LatitudeBase: LatitudeBase{
				Unit:  "DEGREES",
				Value: 45.0,
			},
		}, {
			LatitudeBase: LatitudeBase{
				Value: 45.0,
			},
		}}

	for _, l := range tests {
		if err := Validate(&l); err != nil {
			t.Errorf("latitude struct is not valid: %s", err)
		}
	}
}

func TestLatitude_InValid(t *testing.T) {
	var tests = []Latitude{
		Latitude{
			LatitudeBase: LatitudeBase{
				Unit: "UNIT",
			},
		},
		Latitude{
			LatitudeBase: LatitudeBase{
				Value: 100,
			},
		},
		Latitude{
			LatitudeBase: LatitudeBase{
				Value: -100,
			},
		},
	}

	for _, n := range tests {
		if err := Validate(&n); err == nil {
			t.Errorf("latitude struct should be invalid: %s", n)
		}
	}
}
