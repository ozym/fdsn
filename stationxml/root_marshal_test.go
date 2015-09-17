package stationxml

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMarshalling_Root(t *testing.T) {

	var tests = []struct {
		x string
		r FDSNStationXML
	}{
		{"testdata/fdsn.xml", FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "SeisComP3",
			Sender:        "WEL(GNS_Test)",
			Created:       MustParseDateTime("2015-08-28T11:10:40"),
		},
		},
	}

	for _, test := range tests {
		if err := Validate(test.r); err != nil {
			t.Errorf("FDSNStationXML struct is not valid: %s", err)
		}
	}

	for _, test := range tests {
		x, err := ioutil.ReadFile(test.x)
		if err != nil {
			t.Error(err)
		}

		y, err := test.r.Marshal()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(string(x), string(y)) {
			t.Errorf("FDSNStationXML Marshal:\n\t%q\n\t%q\n", string(x), string(y))
		}
	}

	for _, test := range tests {
		x, err := ioutil.ReadFile(test.x)
		if err != nil {
			t.Error(err)
		}

		var s FDSNStationXML
		err = xml.Unmarshal(x, &s)
		if err != nil {
			t.Error(err)
		}

		y, err := s.Marshal()
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(string(x), string(y)) {
			t.Errorf("Unmarshal/Marshal: \n\t%q\n\t%q\n", string(x), string(y))
		}
	}
}
