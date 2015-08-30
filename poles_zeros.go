package fdsn

type PolesZeros struct {
	// Same meaning as Equipment:resourceId.</xs:documentation>
	ResourceId string `xml:"resourceId,attr"`

	// A name given to this filter.
	Name string `xml:"name,attr"`

	Description *string `xml:",omitempty"`

	// The units of the data as input from the perspective of data acquisition.
	// After correcting data for this response, these would be the resulting units.
	InputUnits Units

	// The units of the data as output from the perspective of data acquisition.
	// These would be the units of the data prior to correcting for this response.
	OutputUnits Units

	PzTransferFunctionType string
	NormalizationFactor    FloatValue
	NormalizationFrequency Frequency
	Zeros                  []PoleZero `xml:"Zero,omitempty"`
	Poles                  []PoleZero `xml:"Pole,omitempty"`
}
