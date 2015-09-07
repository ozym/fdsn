package fdsn

import (
	"encoding/json"
	"fmt"
)

type Frequency struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // HERTZ

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (f Frequency) String() string {
	j, err := json.Marshal(&f)
	if err == nil {
		return string(j)
	}
	return ""
}

func (f *Frequency) IsValid() error {
	if f == nil {
		return nil
	}

	if f.Unit != "" && f.Unit != "HERTZ" {
		return fmt.Errorf("invalid unit: %s", f.Unit)
	}
	if f.PlusError < 0.0 {
		return fmt.Errorf("frequency plus error shouldn't be negative: %g", f.PlusError)
	}
	if f.MinusError < 0.0 {
		return fmt.Errorf("frequency minus error shouldn't be negative: %g", f.MinusError)
	}

	return nil
}
