package fdsn

import (
	"encoding/xml"
	"time"
)

type DateTime struct {
	time.Time
}

const DateTimeFormat = "2006-01-02T15:04:05"

func Now() DateTime {
	return DateTime{time.Now()}
}

func Parse(s string) (DateTime, error) {
	x, err := time.Parse(DateTimeFormat, s)
	return DateTime{x}, err
}

func MustParse(s string) DateTime {
	x, err := time.Parse(DateTimeFormat, s)
	if err != nil {
		panic(err)
	}
	return DateTime{x}
}

func (t *DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.Format(DateTimeFormat), start)
}

func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	x, err := time.Parse(DateTimeFormat, s)
	if err != nil {
		return nil
	}
	*t = DateTime{x}

	return nil
}

func (t *DateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: t.Format(DateTimeFormat)}, nil
}

func (t *DateTime) UnmarshalXMLAttr(attr xml.Attr) error {

	x, err := time.Parse(DateTimeFormat, attr.Value)
	if err != nil {
		return nil
	}
	*t = DateTime{x}

	return nil
}
