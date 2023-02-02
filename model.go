package assertions

import (
	"assertions/assert"
	"assertions/dictionary"
	"assertions/internal/context"
	"errors"
)

type Assert struct {
	dictionary dictionary.Dictionary
}

type assertionResult struct {
	IsOk bool
	Ctx  *context.Context
	Code string
}

func New(translates map[string]string) (*Assert, error) {
	messages, err := dictionary.NewMessages(translates)
	if err != nil {
		return nil, err
	}
	return &Assert{dictionary: *messages}, nil
}

func (a *Assert) StringLengthLess(value string, limit int) error {
	return a.execAssertion(func() assertionResult {
		isOk, actualLength := assert.StringLengthLess(value, limit)
		return assertionResult{
			IsOk: isOk,
			Ctx:  context.NewContext(actualLength, limit),
			Code: dictionary.StringLengthLessCode,
		}
	})
}

func (a *Assert) StringLengthMore(value string, limit int) error {
	return a.execAssertion(func() assertionResult {
		isOk, actualLength := assert.StringLengthMore(value, limit)
		return assertionResult{
			IsOk: isOk,
			Ctx:  context.NewContext(actualLength, limit),
			Code: dictionary.StringLengthMoreCode,
		}
	})
}

func (a *Assert) StringIsEmpty(value string) error {
	return a.execAssertion(func() assertionResult {
		isOk := assert.StringIsEmpty(value)
		return assertionResult{
			IsOk: isOk,
			Ctx:  context.NewContext(nil, value),
			Code: dictionary.StringIsEmptyCode,
		}
	})
}

func (a *Assert) StringIsNotEmpty(value string) error {
	return a.execAssertion(func() assertionResult {
		isOk := assert.StringIsNotEmpty(value)
		return assertionResult{
			IsOk: isOk,
			Ctx:  context.NewContext(nil, nil),
			Code: dictionary.StringIsNotEmptyCode,
		}
	})
}

func (a *Assert) StringsEqual(v1 string, v2 string) error {
	return a.execAssertion(func() assertionResult {
		isOk := assert.StringsEqual(v1, v2)
		return assertionResult{
			IsOk: isOk,
			Ctx:  context.NewContext(v1, v2),
			Code: dictionary.StringsEqualCode,
		}
	})
}

func (a *Assert) execAssertion(assertion func() assertionResult) error {
	result := assertion()
	if !result.IsOk {
		return a.createError(result.Code, result.Ctx)
	}
	return nil
}

func (a *Assert) createError(code string, ctx *context.Context) error {
	message := a.dictionary.Get(code)
	message = ctx.MakeContextMessage(message)
	return errors.New(message)
}
