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
		Code:             "NZ",
		Description:      "New Zealand National Seismograph Network",
		RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
		StartDate:        MustParsePtr("1980-01-01T00:00:00"),
	}

	x, err := xml.Marshal(&n)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testNetwork {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testNetwork, ""}, "\n=========\n"))
	}
}

func TestNetwork_String(t *testing.T) {

	var tests = []struct {
		s string
		x Network
	}{
		{
			`{"Code":"XX","StartDate":"1980-01-01T00:00:00","EndDate":"1980-01-01T00:00:00","RestrictedStatus":"open","AlternateCode":"YY","HistoricalCode":"ZZ","Description":"D"}`,
			Network{
				Code:             "XX",
				StartDate:        MustParsePtr("1980-01-01T00:00:00"),
				EndDate:          MustParsePtr("1980-01-01T00:00:00"),
				RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
				AlternateCode:    "YY",
				HistoricalCode:   "ZZ",
				Description:      "D",
			},
		}, {
			`{"Code":"XX"}`,
			Network{
				Code: "XX",
			},
		}, {
			`{"Code":""}`,
			Network{},
		}}

	for _, test := range tests {
		if test.x.String() != test.s {
			t.Error(strings.Join([]string{"string mismatch:", test.x.String(), test.s, ""}, "\n****\n"))
		}
	}
}

func TestNetwork_Valid(t *testing.T) {

	n := Network{
		Code:             "NZ",
		Description:      "New Zealand National Seismograph Network",
		RestrictedStatus: &RestrictedStatus{STATUS_OPEN},
		StartDate:        MustParsePtr("1980-01-01T00:00:00"),
	}

	if err := Validate(&n); err != nil {
		t.Errorf("network struct is not valid: %s", err)
	}
}

func TestNetwork_InValid(t *testing.T) {
	var tests = []Network{
		Network{
			Code: "",
		},
	}

	for _, n := range tests {
		if err := Validate(&n); err == nil {
			t.Errorf("network struct should be invalid: %s", n)
		}
	}
}
