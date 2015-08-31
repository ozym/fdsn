package fdsn

// This type represents a Station epoch.
// It is common to only have a single station epoch with the station's creation
// and termination dates as the epoch start and end dates.
type Station struct {
	Code             string            `xml:"code,attr"`
	StartDate        *DateTime         `xml:"startDate,attr,omitempty"`
	EndDate          *DateTime         `xml:"endDate,attr,omitempty"`
	RestrictedStatus *RestrictedStatus `xml:"restrictedStatus,attr,omitempty"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode *string `xml:"alternateCode,attr,omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode *string `xml:"historicalCode,attr,omitempty"`

	Description *string   `xml:"description,omitempty"`
	Comments    []Comment `xml:"comment,omitempty"`

	Latitude  Latitude
	Longitude Longitude
	Elevation Distance

	// These fields describe the location of the station using geopolitical
	// entities (country, city, etc.).
	Site Site

	// Type of vault, e.g. WWSSN, tunnel, transportable array, etc.
	Vault *string `xml:,"omitempty"`

	// Type of rock and/or geologic formation.
	Geology *string `xml:,"omitempty"`

	// Equipment used by all channels at a station.
	Equipments []Equipment `xml:"Equipment,omitempty"`

	// An operating agency and associated contact persons.
	// If there multiple operators, each one should be encapsulated within an Operator tag.
	// Since the Contact element is a generic type that represents any contact person,
	// it also has its own optional Agency element.
	Operators []Operator `xml:"Operator,omitempty"`

	// Date and time (UTC) when the station was first installed.
	CreationDate DateTime

	// Date and time (UTC) when the station was terminated or will be terminated.
	// A blank value should be assumed to mean that the station is still active.
	TerminationDate *DateTime `xml",omitempty"`

	// Total number of channels recorded at this station.
	TotalNumberChannels *uint32 `xml:",omitempty"`

	// Number of channels recorded at this station and selected by the query
	// that produced this document.
	SelectedNumberChannels *uint32 `xml:",omitempty"`

	// URI of any type of external report, such as IRIS data reports or dataless SEED volumes.
	ExternalReferences []ExternalReference `xml:"ExternalReference,omitempty"`

	Channels []Channel `xml:"Channel,omitempty"`
}
