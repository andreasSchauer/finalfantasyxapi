// Package apitest includes all the integration tests for the api
package apitest

import "errors"

// could also simply go to errs
var errCorrect = errors.New("test got the expected error") // chilling in exec_test.go
var errIgnoredField = errors.New("this field of the test-struct is ignored") // chilling in test_helpers files (will go to apitesttypes)
