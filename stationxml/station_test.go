package stationxml

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

	n := Station{
		BaseNode: BaseNode{
			Code:             "ABAZ",
			StartDate:        MustParseDateTimePtr("2008-10-13T00:00:00"),
			RestrictedStatus: StatusOpen,
		},
		Site: Site{Name: "Army Bay"},
		Latitude: Latitude{
			LatitudeBase: LatitudeBase{
				Float: Float{
					Value: -36.600224,
				},
			},
		},
		Longitude: Longitude{
			LongitudeBase: LongitudeBase{
				Float: Float{
					Value: 174.832333,
				},
			},
		},
		Elevation:    Distance{Float: Float{Value: 74}},
		CreationDate: MustParseDateTime("2008-10-13T00:00:00"),
	}

	x, err := xml.Marshal(n)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testStation {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testStation, ""}, "\n=========\n"))
	}
}

func TestStation_Valid(t *testing.T) {

	s := Station{
		BaseNode: BaseNode{
			Code:             "ABAZ",
			StartDate:        MustParseDateTimePtr("2008-10-13T00:00:00"),
			RestrictedStatus: StatusOpen,
		},
		Site: Site{Name: "Army Bay"},
		Latitude: Latitude{
			LatitudeBase: LatitudeBase{
				Float: Float{
					Value: -36.600224,
				},
			},
		},
		Longitude: Longitude{
			LongitudeBase: LongitudeBase{
				Float: Float{
					Value: 174.832333,
				},
			},
		},
		Elevation:    Distance{Float: Float{Value: 74}},
		CreationDate: MustParseDateTime("2008-10-13T00:00:00"),
	}

	if err := Validate(s); err != nil {
		t.Errorf("station struct is not valid: %s", err)
	}
}

func TestStation_InValid(t *testing.T) {
	var tests = []Station{
		Station{
			BaseNode: BaseNode{
				Code: "",
			},
		},
		Station{
			Site: Site{},
		},
		Station{
			BaseNode: BaseNode{
				Code: "C",
			},
			Site: Site{},
		},
		Station{
			BaseNode: BaseNode{
				Code: "",
			},
			Site: Site{Name: "N"},
		},
	}

	for _, s := range tests {
		if err := Validate(s); err == nil {
			t.Errorf("station struct should be invalid: %s", s)
		}
	}
}
