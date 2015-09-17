package stationxml

import (
	"testing"
)

func TestCfTransferFunctionType_String(t *testing.T) {

	var tests = []struct {
		r CfTransferFunctionType
		s string
	}{
		{CfFunctionUnknown, "UNKNOWN"},
		{CfFunctionAnalogRadiansPerSecond, "ANALOG (RADIANS/SECOND)"},
		{CfFunctionAnalogHertz, "ANALOG (HERTZ)"},
		{CfFunctionDigital, "DIGITAL"},
	}

	for _, x := range tests {
		if x.r.String() != x.s {
			t.Errorf("invalid string: %s != %s", x.r.String(), x.s)
		}

	}
}

func TestCfTransferFunctionType_Valid(t *testing.T) {

	var tests = []CfTransferFunctionType{
		CfFunctionUnknown,
		CfFunctionAnalogRadiansPerSecond,
		CfFunctionAnalogHertz,
		CfFunctionDigital,
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("restricted status is invalid: %v (%s)", a, err)
		}
	}
}

func TestCfTransferFunctionType_InValid(t *testing.T) {

	var tests = []CfTransferFunctionType{
		CfFunctionDigital + 1,
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("restricted status should be invalid: %v", a)
		}
	}
}
