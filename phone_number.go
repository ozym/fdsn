package fdsn

type PhoneNumber struct {
	Description *string `xml:"description,attr,omitempty"`

	ContryCode *int32 `xml:",omitempty"`
	AreaCode   int32

	// Pattern "[0-9]+-[0-9]+"
	PhoneNumber string
}
