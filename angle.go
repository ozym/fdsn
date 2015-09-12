package fdsn

import (
	"fmt"
)

type Angle struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"`

	// Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (a Angle) IsValid() error {

	switch {
	case a.Unit != "" && a.Unit != "DEGREES":
		return fmt.Errorf("invalid unit: %s", a.Unit)
	case a.Value < -360 || a.Value > 360:
		return fmt.Errorf("angle outside range: %g", a.Value)
	case a.PlusError < 0.0:
		return fmt.Errorf("angle plus error shouldn't be negative: %g", a.PlusError)
	case a.MinusError < 0.0:
		return fmt.Errorf("angle minus error shouldn't be negative: %g", a.MinusError)
	}

	return nil
}
