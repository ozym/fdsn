package fdsn

import (
	"encoding/json"
)

type FrequencyRangeGroup struct {
	FrequencyStart float64
	FrequencyEnd   float64
	// Variation in decibels within the specified range.
	FrequencyDBVariation float64
}

func (f *FrequencyRangeGroup) String() string {

	j, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(j)
}

func (f *FrequencyRangeGroup) IsValid() error {
	return nil
}
