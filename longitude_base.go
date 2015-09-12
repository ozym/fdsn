package fdsn

import (
	"fmt"
)

type LongitudeBase struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (l LongitudeBase) IsValid() error {

	if l.Unit != "" && l.Unit != "DEGREES" {
		return fmt.Errorf("invalid longitude unit: %s", l.Unit)
	}
	if l.Value < -180 || l.Value > 180 {
		return fmt.Errorf("longitude outside range: %g", l.Value)
	}

	return nil
}
