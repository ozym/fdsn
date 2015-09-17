package stationxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFDSNStationXML_Marshal(t *testing.T) {

	var tests = []struct {
		s string
		x FDSNStationXML
	}{
		{
			strings.Join([]string{
				"<FDSNStationXML xmlns=\"http://www.fdsn.org/xml/station/1\" schemaVersion=\"1.0\">",
				"<Source>Test Source</Source>",
				"<Sender>Test Sender</Sender>",
				"<Created>2015-08-28T06:09:37</Created>",
				"</FDSNStationXML>",
			}, ""),
			FDSNStationXML{
				NameSpace:     "http://www.fdsn.org/xml/station/1",
				SchemaVersion: "1.0",
				Source:        "Test Source",
				Sender:        "Test Sender",
				Created:       MustParseDateTime("2015-08-28T06:09:37"),
			},
		}, {
			strings.Join([]string{
				"<FDSNStationXML xmlns=\"http://www.fdsn.org/xml/station/1\" schemaVersion=\"1.0\">",
				"<Source>Test Source</Source>",
				"<ModuleURI>Module URI</ModuleURI>",
				"<Created>2015-08-28T06:09:37</Created>",
				"</FDSNStationXML>",
			}, ""),
			FDSNStationXML{
				NameSpace:     "http://www.fdsn.org/xml/station/1",
				SchemaVersion: "1.0",
				Source:        "Test Source",
				ModuleURI:     "Module URI",
				Created:       MustParseDateTime("2015-08-28T06:09:37"),
			},
		}}

	for _, test := range tests {

		s, err := xml.Marshal(test.x)
		if err != nil {
			t.Error(err)
		}

		if (string)(s) != test.s {
			t.Error(strings.Join([]string{"marshalling mismatch:", (string)(s), test.s, ""}, "\n=========\n"))
		}
	}
}

func TestFDSNStationXML_Valid(t *testing.T) {

	x := FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "Test Source",
		Sender:        "",
		Created:       MustParseDateTime("2015-08-28T06:09:37"),
	}
	if err := Validate(x); err != nil {
		t.Errorf("FDSNStationXML struct is not valid: %s", err)
	}
}

func TestFDSNStationXML_InValid(t *testing.T) {
	var tests = []FDSNStationXML{
		FDSNStationXML{
			NameSpace:     "bad http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "Test Source",
			Created:       MustParseDateTime("2015-08-28T06:09:37"),
		},
		FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "bad 1.0",
			Source:        "Test Source",
			Created:       MustParseDateTime("2015-08-28T06:09:37"),
		},
		FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "",
			Created:       MustParseDateTime("2015-08-28T06:09:37"),
		},
		FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "Test Source",
		},
	}

	for _, x := range tests {
		if err := Validate(x); err == nil {
			t.Errorf("FDSNStationXML struct should be invalid: %s", x)
		}
	}
}
