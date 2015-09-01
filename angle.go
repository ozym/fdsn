package fdsn

import (
	"encoding/json"
	"fmt"
)

type Angle struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"`

	// Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  *float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError *float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (a Angle) String() string {
	j, err := json.Marshal(&a)
	if err == nil {
		return string(j)
	}
	return ""
}

func (a Angle) IsValid() error {
	if a.Unit != "" && a.Unit != "DEGREES" {
		return fmt.Errorf("invalid unit: %s", a.Unit)
	}
	if a.Value < -360 || a.Value > 360 {
		return fmt.Errorf("angle outside range: %g", a.Value)
	}
	return nil
}
