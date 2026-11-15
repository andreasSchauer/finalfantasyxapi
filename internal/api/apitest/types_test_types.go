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
	Compare(Test, string, G)
}

type Test struct {
	T          *testing.T
	Srv        *api.Server
	Name       string
	ExpLengths map[string]int
	DontCheck  map[string]bool
}

type TestGeneral struct {
	RequestURL     string
	ExpectedStatus int
	ExpectedErr    string
	DontCheck      map[string]bool
	ExpLengths     map[string]int
	Endpoint       ts.EndpointName
	Handler        func(http.ResponseWriter, *http.Request)
	Method         string
	RequestBody    any
}

func (t TestGeneral) GetTestGeneral() TestGeneral {
	return t
}

func (t TestGeneral) getTestName(name string, caseNum int) string {
	context := fmt.Sprintf("url: %s", t.RequestURL)

	if t.RequestBody != nil {
		json, _ := json.Marshal(t.RequestBody)
		json = h.CleanUpJSON(json)

		context = fmt.Sprintf("%s | json body: %s", context, string(json))
	}

	return fmt.Sprintf("%s-%d: %s", name, caseNum, context)
}

type NameVer struct {
	ID      int32
	Name    string
	Version *int32
}

func NewExpNameVer(id int32, name string, version int32) NameVer {
	vPtr := new(version)
	if version == 0 {
		vPtr = nil
	}

	return NameVer{
		ID:      id,
		Name:    name,
		Version: vPtr,
	}
}

func (test Test) compareExpNameVer(e NameVer, gotID int32, gotName string, gotVer *int32) {
	test.T.Helper()
	test.compare("id", e.ID, gotID)
	test.compare("name", e.Name, gotName)
	test.compare("version", e.Version, gotVer)
}

type Unique struct {
	ID   int32
	Name string
}

func NewExpUnique(id int32, name string) Unique {
	return Unique{
		ID:   id,
		Name: name,
	}
}

func (test Test) compareExpUnique(e Unique, gotID int32, gotName string) {
	test.T.Helper()
	test.compare("id", e.ID, gotID)
	test.compare("name", e.Name, gotName)
}

type IDOnly struct {
	ID int32
}

func NewExpIDOnly(id int32) IDOnly {
	return IDOnly{
		ID: id,
	}
}

func (test Test) compareExpIDOnly(e IDOnly, gotID int32) {
	test.T.Helper()
	test.compare("id", e.ID, gotID)
}

type ListIDs struct {
	TestGeneral
	apiresource.ListParams
	ParentResource *string
	Results        []int32
}

func (l ListIDs) GetListParams() apiresource.ListParams {
	return l.ListParams
}

type ListPaths struct {
	TestGeneral
	apiresource.ListParams
	Results []string
}

func (l ListPaths) GetListParams() apiresource.ListParams {
	return l.ListParams
}

type ListNames struct {
	TestGeneral
	apiresource.ListParams
	Results []ts.NamedParam
}

func (l ListNames) GetListParams() apiresource.ListParams {
	return l.ListParams
}

func (test Test) compareListParams(exp, got apiresource.ListParams) {
	test.T.Helper()
	test.compare("count", exp.Count, got.Count)
	test.compPageURL("previous", exp.Previous, got.Previous)
	test.compPageURL("next", exp.Next, got.Next)
}
