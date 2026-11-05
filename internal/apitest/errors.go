// Package apitest includes all the integration tests for the api
package apitest

import "errors"

var errCorrect = errors.New("test got the expected error")
var errIgnoredField = errors.New("this field of the test-struct is ignored")
