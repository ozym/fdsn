package fdsn

type FrequencyRangeGroup struct {
	FrequencyStart float64
	FrequencyEnd   float64
	// Variation in decibels within the specified range.
	FrequencyDBVariation float64
}

func (f FrequencyRangeGroup) IsValid() error {
	return nil
}
