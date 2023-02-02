package assert

import "net/mail"

func StringLengthLess(value string, limit int) (bool, int) {
	actual := len(value)
	return actual < limit, actual
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

func StringsEqual(v1 string, v2 string) bool {
	return v1 == v2
}
