package main

import "errors"

var errNotAnID = errors.New("not an id.")
var errNoResource = errors.New("no resources found.")
var errEmptyQuery = errors.New("query parameter is empty.")
var errNoDefaultVal = errors.New("query parameter doesn't have a default value, or default value is unused.")
var errNoSpecialInput = errors.New("query parameter doesn't have special inputs, or no special input was found.")
var errNoIntRange = errors.New("query parameter doesn't have integer range.")
var errCorrect = errors.New("test got the expected error.")
var errIgnoredField = errors.New("this field of the test struct is ignored")
