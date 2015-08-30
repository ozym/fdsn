package fdsn

// Equivalent to SEED blockette 52 and parent element for the related the response blockettes.
type Channel struct {
	Code             string    `xml:"code,attr"`
	StartDate        *DateTime `xml:"startDate,attr,omitempty"`
	EndDate          *DateTime `xml:"endDate,attr,omitempty"`
	RestrictedStatus *string   `xml:"restrictedStatus,attr,omitempty"`
	LocationCode     string    `xml:"locationCode,attr"`

	// A code used for display or association, alternate to the SEED-compliant code.
	AlternateCode *string `xml:"alternateCode,attr,omitempty"`

	// A previously used code if different from the current code.
	HistoricalCode *string `xml:"historicalCode,attr,omitempty"`

	Description *string   `xml:"description,omitempty"`
	Comments    []Comment `xml:"comment,omitempty"`

	// URI of any type of external report, such as data quality reports.
	ExternalReferences []ExternalReference `xml:"ExternalReference,omitempty"`

	// Latitude coordinate of this channel's sensor.
	Latitude Latitude

	//Longitude coordinate of this channel's sensor.
	Longitude Longitude

	// Elevation of the sensor.
	Elevation Distance

	// The local depth or overburden of the instrument's location.
	// For downhole instruments, the depth of the instrument under the surface ground level.
	// For underground vaults, the distance from the instrument to the local ground level above.
	Depth Distance

	// Azimuth of the sensor in degrees from north, clockwise.
	Azimuth *Azimuth `xml:",omitempty"`

	// Dip of the instrument in degrees, down from horizontal
	Dip *Dip `xml:",omitempty"`

	// The type of data this channel collects. Corresponds to channel flags in SEED blockette 52.
	// The SEED volume producer could use the first letter of an Output value as the SEED channel flag.
	Types []Type `xml:"Type,omitempty"`

	// This is a group of elements that represent sample rate.
	// If this group is included, then SampleRate, which is the sample rate in samples per second, is required.
	// SampleRateRatio, which is expressed as a ratio of number of samples in a number of seconds, is optional.
	// If both are included, SampleRate should be considered more definitive.
	SampleRate      SampleRate
	SampleRateRatio *SampleRateRatio `xml:",omitempty"`

	// The storage format of the recorded data (e.g. SEED).
	StorageFormat string
	// A tolerance value, measured in seconds per sample, used as a threshold for time
	// error detection in data from the channel.
	ClockDrift *ClockDrift `xml:",omitempty"`

	CalibrationUnits *Units     `xml:",omitempty"`
	Sensor           *Equipment `xml:",omitempty"`
	PreAmplifier     *Equipment `xml:",omitempty"`
	DataLogger       *Equipment `xml:",omitempty"`
	Equipments       *Equipment `xml:",omitempty"`
	Response         *Response  `xml:",omitempty"`
}
