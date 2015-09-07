package fdsn

import (
	"encoding/json"
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

func (d *Dip) String() string {
	j, err := json.Marshal(d)
	if err == nil {
		return string(j)
	}
	return ""
}

func (d *Dip) IsValid() error {

	if d == nil {
		return nil
	}

	if d.Unit != "" && d.Unit != "DEGREES" {
		return fmt.Errorf("dip invalid unit: %s", d.Unit)
	}

	if d.Value < -90 || d.Value > 90 {
		return fmt.Errorf("dip outside range: %g", d.Value)
	}
	if d.PlusError < 0.0 {
		return fmt.Errorf("dip plus error shouldn't be negative: %g", d.PlusError)
	}
	if d.MinusError < 0.0 {
		return fmt.Errorf("dip minus error shouldn't be negative: %g", d.MinusError)
	}

	return nil
}
