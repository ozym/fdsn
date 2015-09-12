package fdsn

import (
	"testing"
)

func TestRestrictedStatus_String(t *testing.T) {

	var tests = []struct {
		r RestrictedStatus
		s string
	}{
		{StatusOpen, "open"},
		{StatusClosed, "closed"},
		{StatusPartial, "partial"},
	}

	for _, x := range tests {
		if x.r.String() != x.s {
			t.Errorf("invalid string: %s != %s", x.r.String(), x.s)
		}

	}
}

func TestRestrictedStatus_Valid(t *testing.T) {

	var tests = []RestrictedStatus{
		StatusUnknown,
		StatusOpen,
		StatusClosed,
		StatusPartial,
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("restricted status is invalid: %v (%s)", a, err)
		}
	}
}

func TestRestrictedStatus_InValid(t *testing.T) {

	var tests = []RestrictedStatus{
		statusEnd,
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("restricted status should be invalid: %v", a)
		}
	}
}
