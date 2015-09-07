package fdsn

import (
	"encoding/json"
	"fmt"
)

// Representation of floating-point numbers used as measurements. min: 0, max: 360
type Azimuth struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (a *Azimuth) String() string {
	j, err := json.Marshal(&a)
	if err == nil {
		return string(j)
	}
	return ""
}

func (a *Azimuth) IsValid() error {
	if a == nil {
		return nil
	}

	if a.Unit != "" && a.Unit != "DEGREES" {
		return fmt.Errorf("azimuth invalid unit: %s", a.Unit)
	}
	if a.Value < 0 || a.Value > 360 {
		return fmt.Errorf("azimuth outside range: %g", a.Value)
	}
	if a.PlusError < 0.0 {
		return fmt.Errorf("azimuth plus error shouldn't be negative: %g", a.PlusError)
	}
	if a.MinusError < 0.0 {
		return fmt.Errorf("azimuth minus error shouldn't be negative: %g", a.MinusError)
	}

	return nil
}
