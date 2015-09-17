package stationxml

import (
	"testing"
)

func TestApproximationType_String(t *testing.T) {

	var tests = []struct {
		a ApproximationType
		s string
	}{
		{ApproximationTypeUnknown, "UNKNOWN"},
		{ApproximationTypeMaclaurin, "MACLAURIN"},
	}

	for _, x := range tests {
		if x.a.String() != x.s {
			t.Errorf("invalid string: %v (%s)", x.a, x.s)
		}

	}
}

func TestApproximationType_Valid(t *testing.T) {

	var tests = []ApproximationType{
		ApproximationTypeUnknown,
		ApproximationTypeMaclaurin,
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("restricted status is invalid: %s (%s)", a, err)
		}
	}
}

func TestApproximationType_InValid(t *testing.T) {

	var tests = []ApproximationType{
		ApproximationTypeMaclaurin + 1,
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("restricted status should be invalid: %v", a)
		}
	}
}
