package fdsn

import (
	"fmt"
	"regexp"
)

type Email string

func (e Email) IsValid() error {
	if !(len(e) > 0) {
		return fmt.Errorf("empty email")
	}

	if !(regexp.MustCompile(`^[\w\.\-_]+@[\w\.\-_]+$`).MatchString(string(e))) {
		return fmt.Errorf("bad email address: %s", e)
	}

	return nil
}
