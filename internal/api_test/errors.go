package apiTest

import "errors"

var errCorrect = errors.New("test got the expected error.")
var errIgnoredField = errors.New("this field of the test-struct is ignored.")