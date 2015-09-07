package fdsn

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

const (
	CF_FUNCTION_UNKNOWN uint = iota
	CF_FUNCTION_ANALOG_RADIANS_PER_SECOND
	CF_FUNCTION_ANALOG_HERTZ
	CF_FUNCTION_DIGITAL
)

var cfFunctionLookup = []string{
	CF_FUNCTION_UNKNOWN:                   "UNKNOWN",
	CF_FUNCTION_ANALOG_RADIANS_PER_SECOND: "ANALOG (RADIANS/SECOND)",
	CF_FUNCTION_ANALOG_HERTZ:              "ANALOG (HERTZ)",
	CF_FUNCTION_DIGITAL:                   "DIGITAL",
}

var cfFunctionMap = map[string]uint{
	"UNKNOWN":                 CF_FUNCTION_UNKNOWN,
	"ANALOG (RADIANS/SECOND)": CF_FUNCTION_ANALOG_RADIANS_PER_SECOND,
	"ANALOG (HERTZ)":          CF_FUNCTION_ANALOG_HERTZ,
	"DIGITAL":                 CF_FUNCTION_DIGITAL,
}

// The type of data this channel collects. Corresponds to
// channel flags in SEED blockette 52. The SEED volume producer could
// use the first letter of an Output value as the SEED channel flag.
type CfTransferFunctionType struct {
	Type uint
}

func (f *CfTransferFunctionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !(int(f.Type) < len(cfFunctionLookup)) {
		return fmt.Errorf("invalid function entry: %d", f.Type)
	}
	return e.EncodeElement(cfFunctionLookup[f.Type], start)
}

func (f *CfTransferFunctionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	if _, ok := cfFunctionMap[s]; !ok {
		return fmt.Errorf("invalid function: %s", s)
	}

	f.Type = cfFunctionMap[s]

	return nil
}

func (f *CfTransferFunctionType) MarshalJSON() ([]byte, error) {
	if !(int(f.Type) < len(cfFunctionLookup)) {
		return nil, fmt.Errorf("invalid type: %d", f.Type)
	}
	return json.Marshal(cfFunctionLookup[f.Type])
}

func (f *CfTransferFunctionType) UnmarshalJSON(data []byte) error {
	var b []byte
	err := json.Unmarshal(data, b)
	if err != nil {
		return err
	}
	s := string(b)

	if _, ok := cfFunctionMap[s]; !ok {
		return fmt.Errorf("invalid type: %s", s)
	}

	f.Type = cfFunctionMap[s]

	return nil
}

func (f *CfTransferFunctionType) String() string {
	j, err := json.Marshal(f)
	if err == nil {
		return string(j)
	}
	return ""
}

func (f *CfTransferFunctionType) IsValid() error {
	if f == nil {
		return nil
	}

	if !(int(f.Type) < len(cfFunctionLookup)) {
		return fmt.Errorf("invalid function entry: %d", f.Type)
	}

	return nil
}
