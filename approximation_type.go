package fdsn

import (
	"encoding/xml"
	"fmt"
)

const (
	APPROXIMATION_UNKNOWN uint = iota
	APPROXIMATION_MACLAURIN
)

var approximationLookup = []string{
	APPROXIMATION_UNKNOWN:   "UNKNOWN",
	APPROXIMATION_MACLAURIN: "MACLAURIN",
}

var approximationMap = map[string]uint{
	"UNKNOWN":   APPROXIMATION_UNKNOWN,
	"MACLAURIN": APPROXIMATION_MACLAURIN,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type ApproximationType struct {
	Type uint
}

func (f *ApproximationType) Valid() bool {

	if int(f.Type) < len(approximationLookup) {
		return true
	}

	return false
}

func (f *ApproximationType) String() string {

	if f.Valid() {
		return approximationLookup[f.Type]
	}

	return ""
}

func (f *ApproximationType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if f.Valid() {
		return e.EncodeElement(approximationLookup[f.Type], start)
	}

	return fmt.Errorf("invalid function entry: %d", f.Type)
}

func (f *ApproximationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := approximationMap[s]; !ok {
		return fmt.Errorf("invalid function: %s", s)
	}

	f.Type = approximationMap[s]

	return nil
}
