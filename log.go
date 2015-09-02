package fdsn

// Container for log entries.
type Log struct {
	Entries []Comment `xml:"Entry,omitempty" json:",omitempty"`
}
