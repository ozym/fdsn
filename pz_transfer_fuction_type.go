package fdsn

import (
	"encoding/xml"
	"fmt"
)

const (
	FUNCTION_UNKNOWN uint = iota
	FUNCTION_LAPLACE_RADIANS_PER_SECOND
	FUNCTION_LAPLACE_HERTZ
	FUNCTION_LAPLACE_Z_TRANSFORM
)

var functionLookup = []string{
	FUNCTION_UNKNOWN:                    "UNKNOWN",
	FUNCTION_LAPLACE_RADIANS_PER_SECOND: "LAPLACE (RADIANS/SECOND)",
	FUNCTION_LAPLACE_HERTZ:              "LAPLACE (HERTZ)",
	FUNCTION_LAPLACE_Z_TRANSFORM:        "DIGITAL (Z-TRANSFORM)",
}

var functionMap = map[string]uint{
	"UNKNOWN":                  FUNCTION_UNKNOWN,
	"LAPLACE (RADIANS/SECOND)": FUNCTION_LAPLACE_RADIANS_PER_SECOND,
	"LAPLACE (HERTZ)":          FUNCTION_LAPLACE_HERTZ,
	"DIGITAL (Z-TRANSFORM)":    FUNCTION_LAPLACE_Z_TRANSFORM,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type PzTransferFunctionType struct {
	Type uint
}

func (f *PzTransferFunctionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(f.Type) < len(functionLookup)) {
		return fmt.Errorf("invalid function entry: %d", f.Type)
	}
	return e.EncodeElement(functionLookup[f.Type], start)
}

func (f *PzTransferFunctionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := functionMap[s]; !ok {
		return fmt.Errorf("invalid function: %s", s)
	}

	f.Type = functionMap[s]

	return nil
}
