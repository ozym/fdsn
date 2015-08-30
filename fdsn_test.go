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

	x := &FDSNStationXML{
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