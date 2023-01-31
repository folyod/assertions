package assertions

import "testing"

func TestStringLengthLess(t *testing.T) {
	dataProvider := []struct {
		value             string
		limit             int
		expectation       bool
		expectationLength int
	}{
		{"somestring", 100, true, 10},
		{"somestring", 11, true, 10},
		{"somestring", 10, false, 10},
		{"somestring", 5, false, 10},
		{"somestring", 0, false, 10},
		{"somestring", -1, false, 10},
	}

	for _, testcase := range dataProvider {
		result, actualLength := StringLengthLess(testcase.value, testcase.limit)
		if testcase.expectation != result {
			t.Fatalf(
				"Actual string length (%d) more than expectation (%d) characters",
				len(testcase.value),
				testcase.limit)
		}
		if actualLength != testcase.expectationLength {
			t.Fatalf(
				"Expectaion string length (%d) characters - actual (%d)",
				testcase.expectationLength,
				len(testcase.value))
		}
	}
}

func TestStringLengthMore(t *testing.T) {
	dataProvider := []struct {
		value             string
		limit             int
		expectation       bool
		expectationLength int
	}{
		{"somestring", 100, false, 10},
		{"somestring", 11, false, 10},
		{"somestring", 10, true, 10},
		{"somestring", 5, true, 10},
		{"somestring", 0, true, 10},
		{"somestring", -1, true, 10},
	}

	for _, testcase := range dataProvider {
		result, actualLength := StringLengthMore(testcase.value, testcase.limit)
		if testcase.expectation != result {
			t.Fatalf(
				"Actual string length (%d) less than expectation (%d) characters",
				len(testcase.value),
				testcase.limit)
		}
		if actualLength != testcase.expectationLength {
			t.Fatalf(
				"Expectaion string length (%d) characters - actual (%d)",
				testcase.expectationLength,
				len(testcase.value))
		}
	}
}

func TestStringIsEmail(t *testing.T) {
	dataProvider := []struct {
		value       string
		expectation bool
	}{
		{"correct@email.test", true},
		{"incorrectemailtest", false},
	}

	for _, testcase := range dataProvider {
		result := IsEmail(testcase.value)
		if testcase.expectation != result {
			t.Fatalf(
				"Expection that value (%s) is email did not match. Excpect - (%t), actual - (%t)",
				testcase.value,
				testcase.expectation,
				result)
		}
	}
}

func TestStringIsEmpty(t *testing.T) {
	dataProvider := []struct {
		value       string
		expectation bool
	}{
		{"", true},
		{"not empty string", false},
	}

	for _, testcase := range dataProvider {
		result := StringIsEmpty(testcase.value)
		if testcase.expectation != result {
			t.Fatalf(
				"Expection that string (%s) is empty did not match. Excpect - (%t), actual - (%t)",
				testcase.value,
				testcase.expectation,
				result)
		}
	}
}

func TestStringIsNotEmpty(t *testing.T) {
	dataProvider := []struct {
		value       string
		expectation bool
	}{
		{"", false},
		{"not empty string", true},
	}

	for _, testcase := range dataProvider {
		result := StringIsNotEmpty(testcase.value)
		if testcase.expectation != result {
			t.Fatalf(
				"Expection that string (%s) is not empty did not match. Excpect - (%t), actual - (%t)",
				testcase.value,
				testcase.expectation,
				result)
		}
	}
}
