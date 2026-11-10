// Package errs ajsidoajd
package errs

import "errors"

var ErrNotAnID = errors.New("not an id")
var ErrIDNotFound = errors.New("id doesn't exit")
var ErrContinue = errors.New("loop that called the function should be continued")
var ErrNoResource = errors.New("no resources found")
var ErrEmptyQuery = errors.New("query parameter is empty")
var ErrNullArray = errors.New("array field is empty")
var ErrQueryNone = errors.New("'none' input in query")
var ErrQueryRedirect = errors.New("query parameter is not empty, but will be dealt with else where")
var ErrNoDefaultVal = errors.New("query parameter doesn't have a default value, or default value is unused")
var ErrNoSpecialInput = errors.New("query parameter doesn't have special inputs, or no special input was found")
var ErrNoIntRange = errors.New("query parameter doesn't have integer range")
var ErrNoVal = errors.New("param has no value")
