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
		BaseNode: BaseNode{Code: "NZ",
			Description:      "New Zealand National Seismograph Network",
			RestrictedStatus: StatusOpen,
			StartDate:        MustParse("1980-01-01T00:00:00"),
		},
	}

	x, err := xml.Marshal(&n)
	if err != nil {
		t.Error(err)
	}

	if (string)(x) != testNetwork {
		t.Error(strings.Join([]string{"marshalling mismatch:", (string)(x), testNetwork, ""}, "\n=========\n"))
	}
}

func TestNetwork_Valid(t *testing.T) {

	n := Network{
		BaseNode: BaseNode{
			Code:             "NZ",
			Description:      "New Zealand National Seismograph Network",
			RestrictedStatus: StatusOpen,
			StartDate:        MustParse("1980-01-01T00:00:00"),
		},
	}

	if err := Validate(&n); err != nil {
		t.Errorf("network struct is not valid: %s", err)
	}
}

func TestNetwork_InValid(t *testing.T) {
	var tests = []Network{
		Network{
			BaseNode: BaseNode{
				Code: "",
			},
		},
	}

	for _, n := range tests {
		if err := Validate(&n); err == nil {
			t.Errorf("network struct should be invalid: %s", n)
		}
	}
}
