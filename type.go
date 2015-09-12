package fdsn

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

const (
	TYPE_UNKNOWN uint = iota
	TYPE_TRIGGERED
	TYPE_CONTINUOUS
	TYPE_HEALTH
	TYPE_GEOPHYSICAL
	TYPE_WEATHER
	TYPE_FLAG
	TYPE_SYNTHESIZED
	TYPE_INPUT
	TYPE_EXPERIMENTAL
	TYPE_MAINTENANCE
	TYPE_BEAM
)

var typeLookup = []string{
	TYPE_UNKNOWN:      "UNKNOWN",
	TYPE_TRIGGERED:    "TRIGGERED",
	TYPE_CONTINUOUS:   "CONTINUOUS",
	TYPE_HEALTH:       "HEALTH",
	TYPE_GEOPHYSICAL:  "GEOPHYSICAL",
	TYPE_WEATHER:      "WEATHER",
	TYPE_FLAG:         "FLAG",
	TYPE_SYNTHESIZED:  "SYNTHESIZED",
	TYPE_INPUT:        "INPUT",
	TYPE_EXPERIMENTAL: "EXPERIMENTAL",
	TYPE_MAINTENANCE:  "MAINTENANCE",
	TYPE_BEAM:         "BEAM",
}

var typeMap = map[string]uint{
	"UNKNOWN":      TYPE_UNKNOWN,
	"TRIGGERED":    TYPE_TRIGGERED,
	"CONTINUOUS":   TYPE_CONTINUOUS,
	"HEALTH":       TYPE_HEALTH,
	"GEOPHYSICAL":  TYPE_GEOPHYSICAL,
	"WEATHER":      TYPE_WEATHER,
	"FLAG":         TYPE_FLAG,
	"SYNTHESIZED":  TYPE_SYNTHESIZED,
	"INPUT":        TYPE_INPUT,
	"EXPERIMENTAL": TYPE_EXPERIMENTAL,
	"MAINTENANCE":  TYPE_MAINTENANCE,
	"BEAM":         TYPE_BEAM,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type Type struct {
	Type uint
}

func (t Type) IsValid() error {

	if !(int(t.Type) < len(typeLookup)) {
		return fmt.Errorf("invalid type: %d", t.Type)
	}

	return nil
}

func (t Type) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(t.Type) < len(typeLookup)) {
		return fmt.Errorf("invalid type: %d", t.Type)
	}
	return e.EncodeElement(typeLookup[t.Type], start)
}

func (t *Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := typeMap[s]; !ok {
		return fmt.Errorf("invalid type: %s", s)
	}

	t.Type = typeMap[s]

	return nil
}

func (t Type) MarshalJSON() ([]byte, error) {
	if !(int(t.Type) < len(typeLookup)) {
		return nil, fmt.Errorf("invalid type: %d", t.Type)
	}
	return json.Marshal(typeLookup[t.Type])
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var b []byte
	err := json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	s := string(b)

	if _, ok := typeMap[s]; !ok {
		return fmt.Errorf("invalid type: %s", s)
	}

	*t = Type{Type: typeMap[s]}

	return nil
}
