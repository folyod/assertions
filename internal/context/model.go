package context

import (
	"fmt"
	"strings"
)

type Context struct {
	actual      interface{}
	expectation interface{}
}

func NewContext(actual interface{}, expectation interface{}) *Context {
	return &Context{
		actual:      actual,
		expectation: expectation,
	}
}

func (c *Context) MakeContextMessage(message string) string {
	message = c.replaceActualContext(message)
	message = c.replaceExpectationContext(message)
	return message
}

func (c *Context) HasActual() bool {
	return c.actual != nil
}

func (c *Context) Actual() interface{} {
	return c.actual
}

func (c *Context) HasExpectation() bool {
	return c.expectation != nil
}

func (c *Context) Expectation() interface{} {
	return c.expectation
}

func (c *Context) replaceActualContext(message string) string {
	if strings.Index(message, ":a:") != -1 {
		if c.HasActual() {
			return strings.ReplaceAll(message, ":a:", fmt.Sprintf("%v", c.Actual()))
		}
	}
	return message
}

func (c *Context) replaceExpectationContext(message string) string {
	if strings.Index(message, ":e:") != -1 {
		if c.HasExpectation() {
			return strings.ReplaceAll(message, ":e:", fmt.Sprintf("%v", c.Expectation()))
		}
	}
	return message
}
