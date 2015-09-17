package stationxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestLongitude_Marshal(t *testing.T) {

	testLongitude := `<Longitude unit="DEGREES" datum="WGS84">45</Longitude>`

	l := Longitude{
		Datum: "WGS84",
		LongitudeBase: LongitudeBase{
			Float: Float{
				Unit:  "DEGREES",
				Value: 45.0,
			},
		},
	}

	x, err := xml.Marshal(l)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testLongitude {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testLongitude, ""}, "\n=========\n"))
	}
}

func TestLongitude_Valid(t *testing.T) {

	var tests = []Longitude{
		{
			Datum: "WGS84",
			LongitudeBase: LongitudeBase{
				Float: Float{
					Unit:  "DEGREES",
					Value: 45.0,
				},
			},
		}, {
			LongitudeBase: LongitudeBase{
				Float: Float{
					Value: 45.0,
				},
			},
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
			LongitudeBase: LongitudeBase{
				Float: Float{
					Unit: "UNIT",
				},
			},
		},
		Longitude{
			LongitudeBase: LongitudeBase{
				Float: Float{
					Value: 200,
				},
			},
		},
		Longitude{
			LongitudeBase: LongitudeBase{
				Float: Float{
					Value: -200,
				},
			},
		},
	}

	for _, n := range tests {
		if err := Validate(n); err == nil {
			t.Errorf("latitude struct should be invalid: %s", n)
		}
	}
}
