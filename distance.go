package fdsn

// Extension of Float for distances, elevations, and depths.
type Distance struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // METERS

	//Expressing uncertainties or errors with a positive and a negative component.
	// Both values should be given as positive integers, but minus_error is understood to actually be negative.
	PlusError  float64 `xml:"plusError,attr,omitempty" json:",omitempty"`
	MinusError float64 `xml:"minusError,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (d Distance) IsValid() error {
	return nil
}
