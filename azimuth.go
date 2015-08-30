package fdsn

// Representation of floating-point numbers used as measurements. min: 0, max: 360
type Azimuth struct {
	Unit *string `xml:"unit,attr,omitempty"` // DEGREES

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  *float64 `xml:"plusError,attr,omitempty"`
	MinusError *float64 `xml:"minusError,attr,omitempty"`

	Value float64 `xml:",chardata"`
}
