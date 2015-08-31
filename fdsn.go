package fdsn

import (
	"encoding/xml"
	"fmt"
	"strings"
)

const (
	FDSN_XML_HEADER     = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	FDSN_NAME_SPACE     = "http://www.fdsn.org/xml/station/1"
	FDSN_SCHEMA_VERSION = "1.0"
)

// FDSN StationXML schema. Designed as an XML representation of SEED metadata, the schema maps to
// the most important and commonly used structures of SEED 2.4. When definitions and usage are
// underdefined the SEED manual should be referred to for clarification.

// Top-level type for Station XML. Required field are Source (network ID of the institution sending
// the message) and one or more Network containers or one or more Station containers.
type FDSNStationXML struct {
	NameSpace string `xml:"xmlns,attr"`

	// The schema version compatible with the document.
	SchemaVersion string `xml:"schemaVersion,attr"`

	// Network ID of the institution sending the message.
	Source string

	// Name of the institution sending this message.
	Sender string `xml:",omitempty"`

	//Name of the software module that generated this document.
	Module string `xml:",omitempty"`

	// This is the address of the query that generated the document, or,
	// if applicable, the address of the software that generated this document.
	ModuleURI string `xml:",omitempty"`

	Created DateTime

	Networks []Network `xml:"Network,omitempty"`
}

func (x *FDSNStationXML) Marshal() ([]byte, error) {
	h := []byte(FDSN_XML_HEADER)
	s, err := xml.Marshal(x)
	if err != nil {
		return nil, err
	}
	return append(h, append(s, '\n')...), nil
}

// standard interface implementations

func (x FDSNStationXML) String() string {
	var parts []string

	parts = append(parts, fmt.Sprintf("<FDSNStationXML>"))
	parts = append(parts, fmt.Sprintf("SchemaVersion: %s", x.SchemaVersion))
	parts = append(parts, fmt.Sprintf("Source: \"%s\"", x.Source))
	if x.Sender != "" {
		parts = append(parts, fmt.Sprintf("Sender: \"%s\"", x.Sender))
	}
	if x.Module != "" {
		parts = append(parts, fmt.Sprintf("Module: \"%s\"", x.Module))
	}
	if x.ModuleURI != "" {
		parts = append(parts, fmt.Sprintf("ModuleURI: \"%s\"", x.ModuleURI))
	}
	parts = append(parts, fmt.Sprintf("Created: \"%s\"", x.Created))
	parts = append(parts, fmt.Sprintf("Networks: [%d]...", len(x.Networks)))

	return "<" + strings.Join(parts, "; ") + ">"
}

func (x FDSNStationXML) IsValid() error {

	if x.NameSpace != FDSN_NAME_SPACE {
		return fmt.Errorf("wrong name space: %s", x.NameSpace)
	}
	if x.SchemaVersion != FDSN_SCHEMA_VERSION {
		return fmt.Errorf("wrong schema version: %s", x.SchemaVersion)
	}
	if !(len(x.Source) > 0) {
		return fmt.Errorf("empty source element")
	}

	if err := x.Created.IsValid(); err != nil {
		return err
	}

	for _, n := range x.Networks {
		if err := Validate(n); err != nil {
			return err
		}
	}

	return nil
}
