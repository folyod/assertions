package assertions

import (
	"testing"
)

func getDataProvider(t *testing.T) []struct {
	assertion   func() string
	expectation string
} {
	a := createAssert(t)
	return []struct {
		assertion   func() string
		expectation string
	}{
		{
			func() string { return a.StringLengthLess("example", 4).Error() },
			"Actual string length 7 more that expectation 4",
		},
		{
			func() string { return a.StringLengthMore("example", 12).Error() },
			"Actual string length 7 less that expectation 12",
		},
		{
			func() string { return a.StringIsEmpty("example").Error() },
			"String (example) not empty",
		},
		{
			func() string { return a.StringIsNotEmpty("").Error() },
			"String is empty",
		},
		{
			func() string { return a.StringsEqual("example", "exa").Error() },
			"String (example) not equal with string (exa)",
		},
	}
}

func TestErrorMessages(t *testing.T) {
	dataProvider := getDataProvider(t)
	for _, testcase := range dataProvider {
		actual := testcase.assertion()
		if actual != testcase.expectation {
			t.Fatalf("Expectation error message: %s, actual: %s", testcase.expectation, actual)
		}
	}
}

func createAssert(t *testing.T) *Assert {
	a, err := New(translates)
	if err != nil {
		t.Fatalf(err.Error())
	}

	return a
}
