package interfaces

type ServiceValidation interface {
	any
	Validator() error
}
