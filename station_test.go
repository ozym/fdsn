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
		StartDate:        MapDateTime("2008-10-13T00:00:00"),
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
