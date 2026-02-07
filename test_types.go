package main

import (
	"testing"
)

/*
fields that need to be explicitly ignored with dontCheck:
- standard types and pointers (through compare),
- direct apiResource references or pointers
- structs and pointers to structs that are borrowed from the result
- basically anything that isn't a slice or a map and doesn't have nil checks in the test function body


fields that can be explicitly ignored with dontCheck:
- any field that is referenced explicitly as part of dont check in the function body


fields that are implicitly ignored by leaving them blank:
- slices of api resource references
- resourceAmount map references
- slices, structs, pointers to structs of some kind that have nil checks in the function body
*/

type test struct {
	t          *testing.T
	cfg        *Config
	name       string
	expLengths map[string]int
	dontCheck  map[string]bool
}

type testGeneral struct {
	requestURL     string
	expectedStatus int
	expectedErr    string
	dontCheck      map[string]bool
	expLengths     map[string]int
}

type expNameVer struct {
	id      int32
	name    string
	version *int32
}

func testExpectedNameVer(test test, tc expNameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	compare(test, "id", tc.id, gotID)
	compare(test, "name", tc.name, gotName)
	compare(test, "version", tc.version, gotVer)
}

type expUnique struct {
	id   int32
	name string
}

func testExpectedUnique(test test, tc expUnique, gotID int32, gotName string) {
	test.t.Helper()
	compare(test, "id", tc.id, gotID)
	compare(test, "name", tc.name, gotName)
}

type expIdOnly struct {
	id int32
}

func testExpectedIdOnly(test test, tc expIdOnly, gotID int32) {
	test.t.Helper()
	compare(test, "id", tc.id, gotID)
}

type expListIDs struct {
	testGeneral
	count          int
	previous       *string
	next           *string
	parentResource *string
	results        []int32
}

type expListNames struct {
	testGeneral
	count    int
	previous *string
	next     *string
	results  []string
}


type expLocRel struct {
	sidequests []int32
	characters []int32
	aeons      []int32
	shops      []int32
	treasures  []int32
	monsters   []int32
	formations []int32
	bgMusic    []int32
	cuesMusic  []int32
	fmvsMusic  []int32
	bossMusic  []int32
	fmvs       []int32
}
