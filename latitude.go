package fdsn

// Type for latitude coordinate.
type Latitude struct {
	LatitudeBase
	Datum string `xml:"datum,attr,omitempty" json:",omitempty"` // WGS84
}

func (l Latitude) IsValid() error {

	if err := l.LatitudeBase.IsValid(); err != nil {
		return err
	}

	return nil
}
