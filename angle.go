package fdsn

// min: -360 max: 360
type Angle struct {
	// DEGREES
	Unit *string `xml:"unit,attr,omitempty"`

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  *float64 `xml:"plusError,attr,omitempty"`
	MinusError *float64 `xml:"minusError,attr,omitempty"`

	Value float64 `xml:",chardata"`
}
