package fdsn

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Email struct {
	Address string `xml:",chardata"`
}

func (e *Email) String() string {

	j, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(j)
}

func (e *Email) IsValid() error {

	if e == nil {
		return nil
	}

	if !(len(e.Address) > 0) {
		return fmt.Errorf("empty email")
	}

	if !(regexp.MustCompile(`^[\w\.\-_]+@[\w\.\-_]+$`).MatchString(e.Address)) {
		return fmt.Errorf("bad email address: %s", e)
	}

	return nil
}
