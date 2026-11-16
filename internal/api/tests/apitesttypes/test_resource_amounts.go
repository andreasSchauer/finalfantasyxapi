package apitesttypes

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for pure ResourceAmount fields.
func (test Test) checkResAmts[A apiresource.APIResource](fn string, exp map[string]int32, got []apiresource.ResourceAmount[A]) {
	test.T.Helper()
	test.compLength(fn, len(got))

	gotMap := apiresource.GetResAmtMap(got)

	for key, expVal := range exp {
		gotVal, ok := gotMap[key]
		if !ok {
			test.T.Errorf("%s doesn't contain resource with key '%s'.\n\n", fn, key)
		}
		test.compare(key, expVal, gotVal)
	}
}

func (test Test) checkResAmtsID[A apiresource.APIResource](fn string, exp map[int32]int32, got []apiresource.ResourceAmount[A]) {
	test.T.Helper()
	test.compLength(fn, len(got))

	gotMap := apiresource.GetResAmtIDMap(got)

	for id, expVal := range exp {
		gotVal, ok := gotMap[id]
		if !ok {
			test.T.Errorf("%s doesn't contain resource with id '%d'.\n\n", fn, id)
		}
		test.compare(fmt.Sprintf("id: '%d'", id), expVal, gotVal)
	}
}

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for ResourceAmount-ish types with different field names.
func (test Test) checkResAmtTypes[T apiresource.ResourceAmountType[A], A apiresource.APIResource](fn string, exp map[string]int32, got []T) {
	test.T.Helper()
	gotResAmts := apiresource.ResAmtTypesToStructs(got)
	test.checkResAmts(fn, exp, gotResAmts)
}
