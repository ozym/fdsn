package fdsn

import (
	"encoding/json"
	"fmt"
)

// Type for latitude coordinate. min: -180 max: 180
type Longitude struct {
	Datum string `xml:"datum,attr,omitempty" json:",omitempty"` // WGS84
	Unit  string `xml:"unit,attr,omitempty" json:",omitempty"`  // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (l *Longitude) String() string {

	j, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(j)
}

func (l *Longitude) IsValid() error {
	if l == nil {
		return nil
	}

	if l.Unit != "" && l.Unit != "DEGREES" {
		return fmt.Errorf("invalid latitude unit: %s", l.Unit)
	}
	if l.Value < -180 || l.Value > 180 {
		return fmt.Errorf("longitude outside range: %g", l.Value)
	}

	return nil
}
