package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestFDSNStationXML_Marshal(t *testing.T) {

	testFDSNStationXML := strings.Join([]string{
		"<FDSNStationXML xmlns=\"http://www.fdsn.org/xml/station/1\" schemaVersion=\"1.0\">",
		"<Source>Test Source</Source>",
		"<Sender>Test Sender</Sender>",
		"<Created>2015-08-28T06:09:37</Created>",
		"</FDSNStationXML>",
	}, "")

	x := FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "Test Source",
		Sender:        MapString("Test Sender"),
		Created:       MustParse("2015-08-28T06:09:37"),
	}

	s, err := xml.Marshal(&x)
	if err != nil {
		t.Error(err)
	}

	if (string)(s) != testFDSNStationXML {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(s), testFDSNStationXML, ""}, "\n=========\n"))
	}
}

func TestFDSNStationXML_Valid(t *testing.T) {

	x := FDSNStationXML{
		NameSpace:     "http://www.fdsn.org/xml/station/1",
		SchemaVersion: "1.0",
		Source:        "Test Source",
		Sender:        MapString("Test Sender"),
		Created:       MustParse("2015-08-28T06:09:37"),
	}
	if err := Validate(x); err != nil {
		t.Errorf("FDSNStationXML struct is not valid: %s", err)
	}
}

func TestFDSNStationXML_NotValid(t *testing.T) {
	var tests = []FDSNStationXML{
		FDSNStationXML{
			NameSpace:     "bad http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "Test Source",
			Created:       MustParse("2015-08-28T06:09:37"),
		},
		FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "bad 1.0",
			Source:        "Test Source",
			Created:       MustParse("2015-08-28T06:09:37"),
		},
		FDSNStationXML{
			NameSpace:     "http://www.fdsn.org/xml/station/1",
			SchemaVersion: "1.0",
			Source:        "",
			Created:       MustParse("2015-08-28T06:09:37"),
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
