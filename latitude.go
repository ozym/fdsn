package fdsn

import (
	"fmt"
)

// Type for latitude coordinate. min: -90 max: 90
type Latitude struct {
	Datum string `xml:"datum,attr,omitempty" json:",omitempty"` // WGS84
	Unit  string `xml:"unit,attr,omitempty" json:",omitempty"`  // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (l Latitude) IsValid() error {
	if l.Value < -90 || l.Value > 90 {
		return fmt.Errorf("latitude outside range: %g", l.Value)
	}
	return nil
}
