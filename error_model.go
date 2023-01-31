package assertions

type AssertionError struct {
	code    string
	context *context
}

func newError(code string) AssertionError {
	message, isExist := codemap[code]
	if !isExist {

	}

	return AssertionError{
		code:           code,
		serviceMessage: message,
		context:        nil,
	}
}

func (e *AssertionError) RewriteDefaultMessage(message string) {
	e.serviceMessage = message
}

func (e *AssertionError) WithContext(context *context) {
	e.context = context
}

func (e *AssertionError) Code() string {
	return e.code
}

func (e *AssertionError) ServiceMessage() string {
	return e.serviceMessage
}

func (e *AssertionError) HasContext() bool {
	return e.context != nil
}
