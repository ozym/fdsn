package fdsn

import (
// "encoding/xml"
)

// Complex type for sensitivity and frequency ranges.
// This complex type can be used to represent both overall sensitivities and individual stage gains. The
// FrequencyRangeGroup is an optional construct that defines a pass band in Hertz ( FrequencyStart and
// FrequencyEnd) in which the SensitivityValue is valid within the number of decibels specified in FrequencyDBVariation.
type Gain struct {
	// A scalar that, when applied to the data values, converts the data to different units (e.g. Earth units)
	Value float64
	// The frequency (in Hertz) at which the Value is valid.
	Frequency float64
}

func (g Gain) IsValid() error {
	return nil
}

/*
func (g Gain) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// if it's a smallish integer then use that...
	switch {
	case g.Value > 1.0e+10:
		return e.EncodeElement(g, start)
	case (float64)(int64(g.Value)) == (float64)(g.Value):
		return e.EncodeElement(struct {
			Value     int64
			Frequency float64
		}{int64(g.Value), g.Frequency}, start)
	default:
		return e.EncodeElement(g, start)
	}
}
*/
