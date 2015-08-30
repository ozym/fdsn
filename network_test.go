package fdsn

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNetwork_Marshal(t *testing.T) {

	testNetwork := strings.Join([]string{
		"<Network code=\"NZ\" startDate=\"1980-01-01T00:00:00\" restrictedStatus=\"open\">",
		"<Description>New Zealand National Seismograph Network</Description>",
		"</Network>",
	}, "")

	n := &Network{
		Code:             MapString("NZ"),
		Description:      MapString("New Zealand National Seismograph Network"),
		RestrictedStatus: MapString("open"),
		StartDate:        MapDateTime("1980-01-01T00:00:00"),
	}

	x, err := xml.Marshal(&n)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testNetwork {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testNetwork, ""}, "\n=========\n"))
	}
}
