package fdsn

// Sample rate in samples per second.
type SampleRate struct {
	Unit string `xml:"unit,attr,omitempty"` // SAMPLES/S

	Value float64 `xml:",chardata"`
}

func (s SampleRate) IsValid() error {
	return nil
}
