package fdsn

type Equipment struct {
	// This field contains a string that should serve as a unique resource identifier.
	// This identifier can be interpreted differently depending on the datacenter/software
	// that generated the document. Also, we recommend to use something like
	// GENERATOR:Meaningful ID. As a common behaviour equipment with the same ID should
	// contains the same information/be derived from the same base instruments.
	ResourceId *string `xml:"resourceId,attr,omitempty"`

	Type             *string    `xml:",omitempty"`
	Description      *string    `xml:",omitempty"`
	Manufacturer     *string    `xml:",omitempty"`
	Vendor           *string    `xml:",omitempty"`
	Model            *string    `xml:",omitempty"`
	SerialNumber     *string    `xml:",omitempty"`
	InstallationDate *DateTime  `xml:",omitempty"`
	RemovalDate      *DateTime  `xml:",omitempty"`
	CalibrationDates []DateTime `xml:"CalibrationDate,omitempty"`
}
