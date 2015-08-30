package fdsn

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
	FrequencyRangeGroups []FrequencyRangeGroup `xml:"FrequencyRangeGroup,omitempty"`
}
