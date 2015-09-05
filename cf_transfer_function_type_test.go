package fdsn

import (
	"testing"
)

func TestCfTransferFunctionType_String(t *testing.T) {

	var tests = []struct {
		r CfTransferFunctionType
		s string
	}{
		{CfTransferFunctionType{CF_FUNCTION_UNKNOWN}, `"UNKNOWN"`},
		{CfTransferFunctionType{CF_FUNCTION_ANALOG_RADIANS_PER_SECOND}, `"ANALOG (RADIANS/SECOND)"`},
		{CfTransferFunctionType{CF_FUNCTION_ANALOG_HERTZ}, `"ANALOG (HERTZ)"`},
		{CfTransferFunctionType{CF_FUNCTION_DIGITAL}, `"DIGITAL"`},
	}

	for _, x := range tests {
		if x.r.String() != x.s {
			t.Errorf("invalid string: %s != %s", x.r.String(), x.s)
		}

	}
}

func TestCfTransferFunctionType_Valid(t *testing.T) {

	var tests = []CfTransferFunctionType{
		CfTransferFunctionType{CF_FUNCTION_UNKNOWN},
		CfTransferFunctionType{CF_FUNCTION_ANALOG_RADIANS_PER_SECOND},
		CfTransferFunctionType{CF_FUNCTION_ANALOG_HERTZ},
		CfTransferFunctionType{CF_FUNCTION_DIGITAL},
	}

	for _, a := range tests {

		if err := a.IsValid(); err != nil {
			t.Errorf("restricted status is invalid: %v (%s)", a, err)
		}
	}
}

func TestCfTransferFunctionType_InValid(t *testing.T) {

	var tests = []CfTransferFunctionType{
		CfTransferFunctionType{CF_FUNCTION_DIGITAL + 1},
	}

	for _, a := range tests {
		if err := a.IsValid(); err == nil {
			t.Errorf("restricted status should be invalid: %v", a)
		}
	}
}
