package assertions

import "assertions/dictionary"

var translates = map[string]string{
	dictionary.StringLengthLessCode: "Actual string length :a: more that expectation :e:",
	dictionary.StringLengthMoreCode: "Actual string length :a: less that expectation :e:",
	dictionary.StringIsEmptyCode:    "String (:e:) not empty",
	dictionary.StringIsNotEmptyCode: "String is empty",
	dictionary.StringsEqualCode:     "String (:a:) not equal with string (:e:)",
}
