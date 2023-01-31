package assertions

type context struct {
	actual      *interface{}
	expectation *interface{}
}

func (c *context) hasActual() bool {
	return c.actual != nil
}

func (c *context) Actual() interface{} {
	return c.actual
}

func (c *context) hasExpectation() bool {
	return c.expectation != nil
}

func (c *context) Expectation() interface{} {
	return c.expectation
}
