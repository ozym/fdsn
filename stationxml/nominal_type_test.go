package stationxml

import (
	"testing"
)

func TestNominalType_String(t *testing.T) {

	var tests = []struct {
		n NominalType
		s string
	}{
		{Nominal, "NOMINAL"},
		{Calculated, "CALCULATED"},
	}

	for _, x := range tests {
		if x.n.String() != x.s {
			t.Errorf("invalid string: %s != %s", x.n.String(), x.s)
		}

	}
}
