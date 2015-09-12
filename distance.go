package fdsn

// Extension of Float for distances, elevations, and depths.
type Distance struct {
	Float

	Unit string `xml:"unit,attr,omitempty" json:",omitempty"` // METERS
	UncertaintyDouble
}

func (d Distance) IsValid() error {

	if err := Validate(d.Float); err != nil {
		return err
	}
	if err := Validate(d.UncertaintyDouble); err != nil {
		return err
	}

	return nil
}
