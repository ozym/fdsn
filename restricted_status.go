package fdsn

import (
	"encoding/json"
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

func (r *RestrictedStatus) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if !(int(r.Status) < len(restrictedStatusLookup)) {
		return xml.Attr{}, fmt.Errorf("invalid restricted value: %d", r.Status)
	}

	return xml.Attr{Name: name, Value: restrictedStatusLookup[r.Status]}, nil
}

func (r *RestrictedStatus) UnmarshalXMLAttr(attr xml.Attr) error {

	if _, ok := restrictedStatusMap[attr.Value]; !ok {
		return fmt.Errorf("invalid restricted value: %s", attr.Value)
	}

	*r = RestrictedStatus{Status: restrictedStatusMap[attr.Value]}

	return nil
}

func (r *RestrictedStatus) MarshalJSON() ([]byte, error) {
	if !(int(r.Status) < len(restrictedStatusLookup)) {
		return nil, fmt.Errorf("invalid restricted value: %d", r.Status)
	}
	return json.Marshal(restrictedStatusLookup[r.Status])
}

func (r *RestrictedStatus) UnmarshalJSON(data []byte) error {
	var b []byte
	err := json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	s := string(b)

	if _, ok := restrictedStatusMap[s]; !ok {
		return fmt.Errorf("invalid restricted value: %s", s)
	}

	*r = RestrictedStatus{Status: restrictedStatusMap[s]}

	return nil
}

func (r *RestrictedStatus) String() string {

	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r *RestrictedStatus) IsValid() error {
	if r == nil {
		return nil
	}

	if !(int(r.Status) < len(restrictedStatusLookup)) {
		return fmt.Errorf("invalid restricted value: %d", r.Status)
	}

	return nil
}
