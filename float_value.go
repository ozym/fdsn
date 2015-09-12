package fdsn

import (
	"encoding/xml"
)

// A workaround to match XML encoding of large numbers.
type FloatValue float64

func (f *FloatValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// if it's an integer then use that...
	switch {
	case f == nil:
		return nil
	case *f > 1.0e+10:
		return e.EncodeElement(*f, start)
	case (float64)(int64(*f)) == (float64)(*f):
		return e.EncodeElement((int64)(*f), start)
	default:
		return e.EncodeElement(*f, start)
	}
}

func (f FloatValue) IsValid() error {
	return nil
}
