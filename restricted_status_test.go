package fdsn

import (
	"testing"
)

func TestRestrictedStatus_String(t *testing.T) {

	var tests = []struct {
		r RestrictedStatus
		s string
	}{
		{RestrictedStatus{STATUS_OPEN}, `"open"`},
		{RestrictedStatus{STATUS_CLOSED}, `"closed"`},
		{RestrictedStatus{STATUS_PARTIAL}, `"partial"`},
	}

	for _, x := range tests {
		if x.r.String() != x.s {
			t.Errorf("invalid string: %s != %s", x.r.String(), x.s)
		}

	}
}

func TestRestrictedStatus_Valid(t *testing.T) {

	var tests = []RestrictedStatus{
		RestrictedStatus{STATUS_UNKNOWN},
		RestrictedStatus{STATUS_OPEN},
		RestrictedStatus{STATUS_CLOSED},
		RestrictedStatus{STATUS_PARTIAL},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("restricted status is invalid: %v (%s)", a, err)
		}
	}
}

func TestRestrictedStatus_InValid(t *testing.T) {

	var tests = []RestrictedStatus{
		RestrictedStatus{STATUS_PARTIAL + 1},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("restricted status should be invalid: %v", a)
		}
	}
}
