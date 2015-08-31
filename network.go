package fdsn

// This type represents the Network layer, all station metadata is contained within this element.
// The official name of the network or other descriptive information can be included in the
// Description element. The Network can contain 0 or more Stations.
type Network struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode *string `xml:"alternateCode,attr,omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode *string `xml:"historicalCode,attr,omitempty"`

	Description *string   `xml:",omitempty"`
	Comments    []Comment `xml:"comment,omitempty"`

	// The total number of stations contained in this network, including inactive or terminated stations.
	TotalNumberStations *uint32 `xml:",omitempty"`

	// The total number of stations in this network that were selected by the query that produced this document,
	// even if the stations do not appear in the document. (This might happen if the user only wants a document
	// that goes contains only information at the Network level.)
	SelectedNumberStations *uint32 `xml:",omitempty"`

	Stations []Station `xml:"Station,omitempty"`
}
