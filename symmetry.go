package fdsn

import (
	"encoding/xml"
	"fmt"
)

const (
	SYMMETRY_UNKNOWN uint = iota
	SYMMETRY_NONE
	SYMMETRY_EVEN
	SYMMETRY_ODD
)

var symmetryLookup = []string{
	SYMMETRY_UNKNOWN: "UNKNOWN",
	SYMMETRY_NONE:    "NONE",
	SYMMETRY_EVEN:    "EVEN",
	SYMMETRY_ODD:     "ODD",
}

var symmetryMap = map[string]uint{
	"UNKNOWN": SYMMETRY_UNKNOWN,
	"NONE":    SYMMETRY_NONE,
	"EVEN":    SYMMETRY_EVEN,
	"ODD":     SYMMETRY_ODD,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type Symmetry struct {
	Type uint
}

func (t Symmetry) IsValid() error {

	if !(int(t.Type) < len(symmetryLookup)) {
		return fmt.Errorf("invalid symmetry entry: %d", t.Type)
	}

	return nil
}

func (t *Symmetry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(t.Type) < len(symmetryLookup)) {
		return fmt.Errorf("invalid symmetry entry: %d", t.Type)
	}
	return e.EncodeElement(symmetryLookup[t.Type], start)
}

func (t *Symmetry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := symmetryMap[s]; !ok {
		return fmt.Errorf("invalid function: %s", s)
	}

	t.Type = symmetryMap[s]

	return nil
}
