package fdsn

// A tolerance value, measured in seconds per sample, used as a
// threshold for time error detection in data from the channel.
type ClockDrift struct {
	Unit string `xml:"unit,attr,omitempty" json:",omitempty"`

	Value float64 `xml:",chardata"`
}

func (c ClockDrift) IsValid() error {
	return nil
}
