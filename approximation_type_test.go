package fdsn

import (
	"testing"
)

func TestApproximationType(t *testing.T) {

	var tests = []struct {
		a ApproximationType
		s string
	}{
		{ApproximationType{APPROXIMATION_UNKNOWN}, "UNKNOWN"},
		{ApproximationType{APPROXIMATION_MACLAURIN}, "MACLAURIN"},
	}

	for _, x := range tests {
		if x.a.String() != x.s {
			t.Errorf("invalid string: %v", &x)
		}

	}
}
