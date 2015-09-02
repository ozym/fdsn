package fdsn

import (
	"fmt"
)

type Operator struct {
	Agencies []string `xml:"Agency,omitempty"`
	Contacts []Person `xml:"Contact,omitempty"`
	WebSites []string `xml:"WebSite,omitempty"`
}

func (o Operator) IsValid() error {
	for _, a := range o.Agencies {
		if !(len(a) > 0) {
			return fmt.Errorf("empty operator agency")
		}
	}
	for _, c := range o.Contacts {
		if err := c.IsValid(); err != nil {
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
