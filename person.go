package fdsn

// Representation of a person's contact information.
// A person can belong to multiple agencies and have multiple email addresses and phone numbers.
type Person struct {
	Names        []string      `xml:"Name,omitempty"`
	Agencies     []string      `xml:"Agency,omitempty"`
	Email        []Email       `xml:"Email,omitempty"`
	PhoneNumbers []PhoneNumber `xml:"Phone,omitempty"`
}
