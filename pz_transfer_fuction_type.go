package fdsn

import (
	"encoding/xml"
	"fmt"
)

const (
	PZ_FUNCTION_UNKNOWN uint = iota
	PZ_FUNCTION_LAPLACE_RADIANS_PER_SECOND
	PZ_FUNCTION_LAPLACE_HERTZ
	PZ_FUNCTION_LAPLACE_Z_TRANSFORM
)

var pzFunctionLookup = []string{
	PZ_FUNCTION_UNKNOWN:                    "UNKNOWN",
	PZ_FUNCTION_LAPLACE_RADIANS_PER_SECOND: "LAPLACE (RADIANS/SECOND)",
	PZ_FUNCTION_LAPLACE_HERTZ:              "LAPLACE (HERTZ)",
	PZ_FUNCTION_LAPLACE_Z_TRANSFORM:        "DIGITAL (Z-TRANSFORM)",
}

var pzFunctionMap = map[string]uint{
	"UNKNOWN":                  PZ_FUNCTION_UNKNOWN,
	"LAPLACE (RADIANS/SECOND)": PZ_FUNCTION_LAPLACE_RADIANS_PER_SECOND,
	"LAPLACE (HERTZ)":          PZ_FUNCTION_LAPLACE_HERTZ,
	"DIGITAL (Z-TRANSFORM)":    PZ_FUNCTION_LAPLACE_Z_TRANSFORM,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type PzTransferFunctionType struct {
	Type uint
}

func (f PzTransferFunctionType) IsValid() error {

	if !(int(f.Type) < len(pzFunctionLookup)) {
		return fmt.Errorf("invalid transfer function type: %d", f.Type)
	}

	return nil
}

func (f *PzTransferFunctionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(f.Type) < len(pzFunctionLookup)) {
		return fmt.Errorf("invalid function entry: %d", f.Type)
	}
	return e.EncodeElement(pzFunctionLookup[f.Type], start)
}

func (f *PzTransferFunctionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := pzFunctionMap[s]; !ok {
		return fmt.Errorf("invalid function: %s", s)
	}

	f.Type = pzFunctionMap[s]

	return nil
}
