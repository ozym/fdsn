package fdsn

// Container for a comment or log entry.
// Corresponds to SEED blockettes 31, 51 and 59.
type Comment struct {
	Id uint32 `xml:"id,attr"`

	Value              string
	BeginEffectiveTime *DateTime `xml:",omitempty"`
	EndEffectiveTime   *DateTime `xml:",omitempty"`

	Authors []Person `xml:"Author,omitempty"`
}
