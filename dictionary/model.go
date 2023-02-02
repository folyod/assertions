package dictionary

import (
	"errors"
	"fmt"
	"strings"
)

var StringLengthLessCode = "string_length_less"
var StringLengthMoreCode = "string_length_more"
var StringIsEmptyCode = "string_empty"
var StringIsNotEmptyCode = "string_not_empty"
var IsEmailCode = "is_email"
var StringsEqualCode = "strings_equal"

var requiredKeys = []string{
	StringLengthLessCode,
	StringLengthMoreCode,
	StringIsEmptyCode,
	StringIsNotEmptyCode,
	IsEmailCode,
	StringsEqualCode,
}

type Dictionary struct {
	translates map[string]string
}

func NewMessages(translates map[string]string) (*Dictionary, error) {
	missedKeys := checkRequiredKeys(translates)
	if len(missedKeys) == 0 {
		missedKeysString := strings.Join(missedKeys, ",")
		errorMessage := fmt.Sprintf("Required keys missing: %s", missedKeysString)
		return nil, errors.New(errorMessage)
	}

	return &Dictionary{translates: translates}, nil
}

func checkRequiredKeys(dictionary map[string]string) []string {
	var missedKeys []string
	for _, key := range requiredKeys {
		_, isSet := dictionary[key]
		if !isSet {
			missedKeys = append(missedKeys, key)
		}
	}
	return missedKeys
}

func (m *Dictionary) Get(code string) string {
	message, isSet := m.translates[code]
	if !isSet {
		panic(fmt.Sprintf("Code %s not found in dictionary", code))
	}
	return message
}
