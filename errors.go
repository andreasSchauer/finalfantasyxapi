package main

import "errors"

var errNotAnID = errors.New("not an id")
var errNoResource = errors.New("no resources found")
var errEmptyQuery = errors.New("query parameter is empty")