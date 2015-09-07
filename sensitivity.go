package fdsn

import (
	"encoding/json"
)

// Sensitivity and frequency ranges.
// The FrequencyRangeGroup is an optional construct that defines a pass band in Hertz
// (FrequencyStart and FrequencyEnd) in which the SensitivityValue is valid within the
// number of decibels specified in FrequencyDBVariation.
type Sensitivity struct {
	// A scalar that, when applied to the data values, converts the data to different units (e.g. Earth units)
	Value FloatValue
	// The frequency (in Hertz) at which the Value is valid.
	Frequency float64
	// The units of the data as input from the perspective of data acquisition.
	// After correcting data for this response, these would be the resulting units.
	InputUnits Units
	// The units of the data as output from the perspective of data acquisition.
	// These would be the units of the data prior to correcting for this response.
	OutputUnits Units
	// The frequency range for which the SensitivityValue is valid within the dB variation specified.
	FrequencyRangeGroups []FrequencyRangeGroup `xml:"FrequencyRangeGroup,omitempty" json:",omitempty"`
}

func (s *Sensitivity) String() string {

	j, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(j)
}

func (s *Sensitivity) IsValid() error {
	if s == nil {
		return nil
	}

	if err := Validate(&s.Value); err != nil {
		return err
	}
	if err := Validate(&s.InputUnits); err != nil {
		return err
	}
	if err := Validate(&s.OutputUnits); err != nil {
		return err
	}
	for _, f := range s.FrequencyRangeGroups {
		if err := Validate(&f); err != nil {
			return err
		}
	}

	return nil
}

func (s *Sensitivity) Copy(level Level) *Sensitivity {

	if s == nil {
		return nil
	}

	switch {
	case level < CHANNEL_LEVEL:
		return nil
	case level > CHANNEL_LEVEL:
		return s
	}

	return &Sensitivity{
		Value:      s.Value,
		Frequency:  s.Frequency,
		InputUnits: s.InputUnits,
	}
}
