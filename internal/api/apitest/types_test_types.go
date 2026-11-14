package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
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

// compFnDifferent will soon not be necessary anymore
type compFnVals[T any] func(string, T, T)

type testStructIdx[G any] interface {
	GetIndex() int
	testStruct[G]
}

type testStruct[G any] interface {
	Compare(test, string, G)
}

type test struct {
	t          *testing.T
	srv        *api.Server
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
	endpoint       ts.EndpointName
	handler        func(http.ResponseWriter, *http.Request)
	method         string
	requestBody    any
}

func (t testGeneral) GetTestGeneral() testGeneral {
	return t
}

func (t testGeneral) getTestName(name string, caseNum int) string {
	context := fmt.Sprintf("url: %s", t.requestURL)

	if t.requestBody != nil {
		json, _ := json.Marshal(t.requestBody)
		json = h.CleanUpJSON(json)

		context = fmt.Sprintf("%s | json body: %s", context, string(json))
	}

	return fmt.Sprintf("%s-%d: %s", name, caseNum, context)
}

type NameVer struct {
	id      int32
	name    string
	version *int32
}

func newExpNameVer(id int32, name string, version int32) NameVer {
	vPtr := new(version)
	if version == 0 {
		vPtr = nil
	}

	return NameVer{
		id:      id,
		name:    name,
		version: vPtr,
	}
}

func (test test) compareExpNameVer(e NameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	test.compare("id", e.id, gotID)
	test.compare("name", e.name, gotName)
	test.compare("version", e.version, gotVer)
}

type Unique struct {
	id   int32
	name string
}

func newExpUnique(id int32, name string) Unique {
	return Unique{
		id:   id,
		name: name,
	}
}

func (test test) compareExpUnique(e Unique, gotID int32, gotName string) {
	test.t.Helper()
	test.compare("id", e.id, gotID)
	test.compare("name", e.name, gotName)
}

type IDOnly struct {
	id int32
}

func newExpIDOnly(id int32) IDOnly {
	return IDOnly{
		id: id,
	}
}

func (test test) compareExpIDOnly(e IDOnly, gotID int32) {
	test.t.Helper()
	test.compare("id", e.id, gotID)
}

type ListIDs struct {
	testGeneral
	count          int
	previous       *string
	next           *string
	parentResource *string
	results        []int32
}

func (l ListIDs) GetListParams() apiresource.ListParams {
	return apiresource.ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

type ListPaths struct {
	testGeneral
	count    int
	previous *string
	next     *string
	results  []string
}

func (l ListPaths) GetListParams() apiresource.ListParams {
	return apiresource.ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

type ListNames struct {
	testGeneral
	count    int
	previous *string
	next     *string
	results  []ts.NamedParam
}

func (l ListNames) GetListParams() apiresource.ListParams {
	return apiresource.ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

func (test test) compareListParams(exp, got apiresource.ListParams) {
	test.t.Helper()
	test.compare("count", exp.Count, got.Count)
	test.compPageURL("previous", exp.Previous, got.Previous)
	test.compPageURL("next", exp.Next, got.Next)
}
