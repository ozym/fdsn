package fdsn

import (
	"fmt"
)

// Type for latitude coordinate. min: -180 max: 180
type Longitude struct {
	Datum string `xml:"datum,attr,omitempty"` // WGS84
	Unit  string `xml:"unit,attr,omitempty"`  // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty"`

	Value float64 `xml:",chardata"`
}

func (l Longitude) IsValid() error {
	if l.Value < -180 || l.Value > 180 {
		return fmt.Errorf("longitude outside range: %g", l.Value)
	}
	return nil
}
