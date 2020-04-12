package pkg

type ErrCode string

const (
	InternalError ErrCode = "InternalError"
)

type PublicError struct {
	Message       string  `json:"message"`
	Code          ErrCode `json:"code"`
	InternalError error   `json:"-"`
}

func (f *PublicError) Error() string {
	return f.Message
}

func (f *PublicError) Unwrap() error {
	return f.InternalError
}

func NewPublicError(message string, params ...interface{}) *PublicError {
	var internalError error
	code := InternalError
	for _, param := range params {
		if code == "" {
			if curCode, ok := param.(ErrCode); ok {
				code = curCode
				continue
			}
		}
		if internalError == nil {
			if curInternalError, ok := param.(error); ok && curInternalError != nil {
				internalError = curInternalError
				if ie, ok := internalError.(*PublicError); ok && ie != nil {
					message += ": " + ie.Message
				}
				continue
			}
		}
	}
	return &PublicError{Message: message, Code: code, InternalError: internalError}
}
