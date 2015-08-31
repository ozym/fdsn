package fdsn

import (
	"encoding/xml"
	"fmt"
)

const (
	STATUS_UNKNOWN uint = iota
	STATUS_OPEN
	STATUS_CLOSED
	STATUS_PARTIAL
)

var restrictedStatusLookup = []string{
	STATUS_UNKNOWN: "unknown",
	STATUS_OPEN:    "open",
	STATUS_CLOSED:  "closed",
	STATUS_PARTIAL: "partial",
}

var restrictedStatusMap = map[string]uint{
	"unknown": STATUS_UNKNOWN,
	"open":    STATUS_OPEN,
	"closed":  STATUS_CLOSED,
	"partial": STATUS_PARTIAL,
}

type RestrictedStatus struct {
	Status uint
}

func (r *RestrictedStatus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(r.Status) < len(restrictedStatusLookup)) {
		return fmt.Errorf("invalid nominal entry: %d", r.Status)
	}
	return e.EncodeElement(restrictedStatusLookup[r.Status], start)
}

func (r *RestrictedStatus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := restrictedStatusMap[s]; !ok {
		return fmt.Errorf("invalid nominal value: %s", s)
	}

	r.Status = restrictedStatusMap[s]

	return nil
}

func (r *RestrictedStatus) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if !(int(r.Status) < len(restrictedStatusLookup)) {
		return xml.Attr{}, fmt.Errorf("invalid nominal entry: %d", r.Status)
	}

	return xml.Attr{Name: name, Value: restrictedStatusLookup[r.Status]}, nil
}

func (r *RestrictedStatus) UnmarshalXMLAttr(attr xml.Attr) error {

	if _, ok := restrictedStatusMap[attr.Value]; !ok {
		return fmt.Errorf("invalid nominal value: %s", attr.Value)
	}

	*r = RestrictedStatus{Status: restrictedStatusMap[attr.Value]}

	return nil
}
