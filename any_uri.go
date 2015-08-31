package fdsn

import (
	"fmt"
)

type AnyURI string

func (a AnyURI) IsValid() error {
	if !(len(a) > 0) {
		return fmt.Errorf("empty uri")
	}
	return nil
}
