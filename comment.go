package fdsn

// Container for a comment or log entry.
// Corresponds to SEED blockettes 31, 51 and 59.
type Comment struct {
	Id uint32 `xml:"id,attr"`

	Value              string
	BeginEffectiveTime DateTime `xml:",omitempty" json:",omitempty"`
	EndEffectiveTime   DateTime `xml:",omitempty" json:",omitempty"`

	Authors []Person `xml:"Author,omitempty" json:",omitempty"`
}

func (c Comment) IsValid() error {

	if err := Validate(c.BeginEffectiveTime); err != nil {
		return err
	}
	if err := Validate(c.EndEffectiveTime); err != nil {
		return err
	}

	for _, p := range c.Authors {
		if err := Validate(&p); err != nil {
			return err
		}
	}

	return nil
}
