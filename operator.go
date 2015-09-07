package fdsn

import (
	"encoding/json"
	"fmt"
)

type Operator struct {
	Agencies []string `xml:"Agency,omitempty" json:",omitempty"`
	Contacts []Person `xml:"Contact,omitempty" json:",omitempty"`
	WebSites []string `xml:"WebSite,omitempty" json:",omitempty"`
}

func (o Operator) String() string {

	j, err := json.Marshal(&o)
	if err != nil {
		return ""
	}
	return string(j)
}

func (o *Operator) IsValid() error {

	if o == nil {
		return nil
	}

	for _, a := range o.Agencies {
		if !(len(a) > 0) {
			return fmt.Errorf("empty operator agency")
		}
	}
	for _, c := range o.Contacts {
		if err := Validate(&c); err != nil {
			return err
		}
	}
	for _, w := range o.WebSites {
		if !(len(w) > 0) {
			return fmt.Errorf("empty websites uri")
		}
	}
	return nil
}
