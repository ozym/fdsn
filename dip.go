package fdsn

import (
	"fmt"
)

// Instrument dip in degrees down from horizontal.
// Together azimuth and dip describe the direction of the sensitive axis of the instrument. min: -90, max: 90
type Dip struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (d Dip) IsValid() error {
	if d.Unit != "" && d.Unit != "DEGREES" {
		return fmt.Errorf("dip invalid unit: %s", d.Unit)
	}

	if d.Value < -90 || d.Value > 90 {
		return fmt.Errorf("dip outside range: %g", d.Value)
	}
	return nil
}
