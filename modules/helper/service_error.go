package helper

type ServiceError interface {
	ErrorType() error
	Reason() string
	Cause() error
	Error() string
}
