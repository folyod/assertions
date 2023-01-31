package assertions

import "net/mail"

func StringLengthLess(value string, limit int) *AssertionError {
	actual := len(value)
	isOk := actual < limit
	if !isOk {
		err := newError()
	}

	return nil
}

func StringLengthMore(value string, limit int) (bool, int) {
	actual := len(value)
	return actual >= limit, actual
}

func StringIsEmpty(value string) bool {
	return value == ""
}

func StringIsNotEmpty(value string) bool {
	return value != ""
}

func IsEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}
