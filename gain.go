package fdsn

import (
	"encoding/json"
)

// Complex type for sensitivity and frequency ranges.
// This complex type can be used to represent both overall sensitivities and individual stage gains. The
// FrequencyRangeGroup is an optional construct that defines a pass band in Hertz ( FrequencyStart and
// FrequencyEnd) in which the SensitivityValue is valid within the number of decibels specified in FrequencyDBVariation.
type Gain struct {
	// A scalar that, when applied to the data values, converts the data to different units (e.g. Earth units)
	Value FloatValue
	// The frequency (in Hertz) at which the Value is valid.
	Frequency float64
}

func (g *Gain) String() string {

	j, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(j)
}

func (g *Gain) IsValid() error {
	if err := Validate(&g.Value); err != nil {
		return err
	}
	return nil
}
