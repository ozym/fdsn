package fdsn

type Validator interface {
	IsValid() error
}

func Validate(v Validator) error {
	return v.IsValid()
}
