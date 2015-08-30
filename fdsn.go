package fdsn

import (
	"encoding/xml"
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
	Sender *string `xml:",omitempty"`

	//Name of the software module that generated this document.
	Module *string `xml:",omitempty"`

	// This is the address of the query that generated the document, or,
	// if applicable, the address of the software that generated this document.
	ModuleURI *string `xml:",omitempty"`

	Created DateTime

	Networks []Network `xml:"Network,omitempty"`
}

func (x *FDSNStationXML) Marshal() ([]byte, error) {
	h := []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
	s, err := xml.Marshal(x)
	if err != nil {
		return nil, err
	}
	return append(h, append(s, '\n')...), nil
}
