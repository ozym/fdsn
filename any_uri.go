package fdsn

type AnyURI string

func (u AnyURI) IsValid() error {
	return nil
}
