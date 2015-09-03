package fdsn

import (
	"fmt"
)

// Representation of floating-point numbers used as measurements. min: 0, max: 360
type Azimuth struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  *float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError *float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (a Azimuth) IsValid() error {
	if a.Unit != "" && a.Unit != "DEGREES" {
		return fmt.Errorf("azimuth invalid unit: %s", a.Unit)
	}
	if a.Value < 0 || a.Value > 360 {
		return fmt.Errorf("azimuth outside range: %g", a.Value)
	}
	return nil
}
