package fdsn

import (
	"fmt"
)

// Base latitude type. Because of the limitations of schema, defining this type and then extending
// it to create the real latitude type is the only way to restrict values while adding datum as an attribute.
type LatitudeBase struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (l LatitudeBase) IsValid() error {

	if l.Unit != "" && l.Unit != "DEGREES" {
		return fmt.Errorf("invalid latitude unit: %s", l.Unit)
	}
	if l.Value < -90 || l.Value > 90 {
		return fmt.Errorf("longitude outside range: %g", l.Value)
	}

	return nil
}
