package fdsn

import (
	"encoding/json"
)

// Container for log entries.
type Log struct {
	Entries []Comment `xml:"Entry,omitempty" json:",omitempty"`
}

func (l *Log) String() string {

	j, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(j)
}

func (l *Log) IsValid() error {

	if l == nil {
		return nil
	}

	for _, c := range l.Entries {
		if err := Validate(&c); err != nil {
			return err
		}
	}

	return nil
}
