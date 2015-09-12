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

func (f ApproximationType) String() string {

	if int(f.Type) < len(approximationLookup) {
		return approximationLookup[f.Type]
	}

	return approximationLookup[0]
}

func (f ApproximationType) IsValid() error {

	if !(int(f.Type) < len(approximationLookup)) {
		return fmt.Errorf("invalid approximation type: %d", f.Type)
	}

	return nil
}

func (f *ApproximationType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if int(f.Type) < len(approximationLookup) {
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
