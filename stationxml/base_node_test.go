package stationxml

import (
	"testing"
)

func TestBaseNode_Valid(t *testing.T) {

	var tests = []BaseNode{
		BaseNode{Code: "CODE"},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("base node is invalid: %v (%s)", a, err)
		}
	}
}

func TestBaseNode_InValid(t *testing.T) {

	var tests = []BaseNode{
		BaseNode{},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("base node should be invalid: %v", a)
		}
	}
}
