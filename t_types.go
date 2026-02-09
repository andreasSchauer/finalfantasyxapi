package main

import (
	"testing"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

type testStructIdx interface {
	GetIndex() int
}

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

func newExpNameVer(id int32, name string, version int32) expNameVer {
	vPtr := h.GetInt32Ptr(version)
	if version == 0 {
		vPtr = nil
	}
	
	return expNameVer{
		id: 		id,
		name: 		name,
		version: 	vPtr,
	}
}

func compareExpNameVer(test test, exp expNameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
	compare(test, "name", exp.name, gotName)
	compare(test, "version", exp.version, gotVer)
}

type expUnique struct {
	id   int32
	name string
}

func newExpUnique(id int32, name string) expUnique {
	return expUnique{
		id: 	id,
		name: 	name,
	}
}

func compareExpUnique(test test, exp expUnique, gotID int32, gotName string) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
	compare(test, "name", exp.name, gotName)
}

type expIdOnly struct {
	id int32
}

func newExpIdOnly(id int32) expIdOnly {
	return expIdOnly{
		id: id,
	}
}

func compareExpIdOnly(test test, exp expIdOnly, gotID int32) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
}

type expListIDs struct {
	testGeneral
	count          int
	previous       *string
	next           *string
	parentResource *string
	results        []int32
}

func (l expListIDs) getListParams() ListParams {
	return ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

type expListNames struct {
	testGeneral
	count    int
	previous *string
	next     *string
	results  []string
}

func (l expListNames) getListParams() ListParams {
	return ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

func compareListParams(test test, exp, got ListParams) {
	compare(test, "count", exp.Count, got.Count)
	compPageURL(test, "previous", exp.Previous, got.Previous)
	compPageURL(test, "next", exp.Next, got.Next)
}