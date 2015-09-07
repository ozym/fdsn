package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestStation_Marshal(t *testing.T) {

	testStation := strings.Join([]string{
		"<Station code=\"ABAZ\" startDate=\"2008-10-13T00:00:00\" restrictedStatus=\"open\">",
		"<Latitude>-36.600224</Latitude>",
		"<Longitude>174.832333</Longitude>",
		"<Elevation>74</Elevation>",
		"<Site>",
		"<Name>Army Bay</Name>",
		"</Site>",
		"<CreationDate>2008-10-13T00:00:00</CreationDate>",
		"</Station>",
	}, "")

	n := &Station{
		Code:             "ABAZ",
		Site:             Site{Name: "Army Bay"},
		StartDate:        MustParsePtr("2008-10-13T00:00:00"),
		RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
		Latitude:         Latitude{Value: -36.600224},
		Longitude:        Longitude{Value: 174.832333},
		Elevation:        Distance{Value: 74},
		CreationDate:     MustParse("2008-10-13T00:00:00"),
	}

	x, err := xml.Marshal(&n)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testStation {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testStation, ""}, "\n=========\n"))
	}
}

func TestStation_String(t *testing.T) {

	var tests = []struct {
		s string
		x Station
	}{
		{
			`{"Code":"ABAZ","StartDate":"2008-10-13T00:00:00","RestrictedStatus":"open","Latitude":{"Value":-36.600224},"Longitude":{"Value":174.832333},"Elevation":{"Value":74},"Site":{"Name":"Army Bay"},"CreationDate":"2008-10-13T00:00:00"}`,
			Station{
				Code:             "ABAZ",
				Site:             Site{Name: "Army Bay"},
				StartDate:        MustParsePtr("2008-10-13T00:00:00"),
				RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
				Latitude:         Latitude{Value: -36.600224},
				Longitude:        Longitude{Value: 174.832333},
				Elevation:        Distance{Value: 74},
				CreationDate:     MustParse("2008-10-13T00:00:00"),
			},
		}, {
			`{"Code":"XX","Latitude":{"Value":0},"Longitude":{"Value":0},"Elevation":{"Value":0},"Site":{"Name":""},"CreationDate":"0001-01-01T00:00:00"}`,
			Station{
				Code:      "XX",
				Latitude:  Latitude{Value: 0.0},
				Longitude: Longitude{Value: 0.0},
			},
		}, {
			`{"Code":"","Latitude":{"Value":0},"Longitude":{"Value":0},"Elevation":{"Value":0},"Site":{"Name":""},"CreationDate":"0001-01-01T00:00:00"}`,
			Station{},
		}}

	for _, test := range tests {
		if test.x.String() != test.s {
			t.Error(strings.Join([]string{"string mismatch:", test.x.String(), test.s, ""}, "\n****\n"))
		}
	}
}

func TestStation_Valid(t *testing.T) {

	s := Station{
		Code:             "ABAZ",
		Site:             Site{Name: "Army Bay"},
		StartDate:        MustParsePtr("2008-10-13T00:00:00"),
		RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
		Latitude:         Latitude{Value: -36.600224},
		Longitude:        Longitude{Value: 174.832333},
		Elevation:        Distance{Value: 74},
		CreationDate:     MustParse("2008-10-13T00:00:00"),
	}

	if err := Validate(&s); err != nil {
		t.Errorf("station struct is not valid: %s", err)
	}
}

func TestStation_InValid(t *testing.T) {
	var tests = []Station{
		Station{
			Code: "",
		},
		Station{
			Site: Site{},
		},
		Station{
			Code: "C",
			Site: Site{},
		},
		Station{
			Code: "",
			Site: Site{Name: "N"},
		},
	}

	for _, s := range tests {
		if err := Validate(&s); err == nil {
			t.Errorf("station struct should be invalid: %s", s)
		}
	}
}
