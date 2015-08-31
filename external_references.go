package fdsn

// This type contains a URI and description for external data that users may want to reference in StationXML.
type ExternalReference struct {
	URI         AnyURI
	Description string
}

func (e ExternalReference) IsValid() error {
	if err := e.URI.IsValid(); err != nil {
		return err
	}
	return nil
}
