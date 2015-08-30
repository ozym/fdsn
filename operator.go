package fdsn

type Operator struct {
	Agencies []string `xml:"Agency,omitempty"`
	Contacts []Person `xml:"Contact,omitempty"`
	WebSites []AnyURI `xml:"WebSite,omitempty"`
}
