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

type compFnVals[T any] func(string, T, T)
type compFnEqual[T any] func(Test, string, T, T)
type compFnDifferent[E, G any] func(Test, string, E, G)

type TestStructIdx interface {
	GetIndex() int
}

type Test struct {
	t          *testing.T
	srv        *api.Server
	name       string
	expLengths map[string]int
	dontCheck  map[string]bool
}

type TestGeneral struct {
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

func (t TestGeneral) GetTestGeneral() TestGeneral {
	return t
}

func (t TestGeneral) getTestName(name string, caseNum int) string {
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

type IDOnly struct {
	id int32
}

func newExpIDOnly(id int32) IDOnly {
	return IDOnly{
		id: id,
	}
}

type ListIDs struct {
	TestGeneral
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
	TestGeneral
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
	TestGeneral
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
