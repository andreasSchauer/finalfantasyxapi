package apitest

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for pure ResourceAmount fields.
func (test test) checkResAmts[A apiresource.APIResource](fieldName string, exp map[string]int32, got []apiresource.ResourceAmount[A]) {
	test.t.Helper()
	test.compLength(fieldName, len(got))

	gotMap := api.GetResAmtMap(got)

	for key, expVal := range exp {
		gotVal, ok := gotMap[key]
		if !ok {
			test.t.Errorf("%s doesn't contain resource with key '%s'.\n\n", fieldName, key)
		}
		test.compare(key, expVal, gotVal)
	}
}

func (test test) checkResAmtsID[A apiresource.APIResource](fieldName string, exp map[int32]int32, got []apiresource.ResourceAmount[A]) {
	test.t.Helper()
	test.compLength(fieldName, len(got))

	gotMap := api.GetResAmtIDMap(got)

	for id, expVal := range exp {
		gotVal, ok := gotMap[id]
		if !ok {
			test.t.Errorf("%s doesn't contain resource with id '%d'.\n\n", fieldName, id)
		}
		test.compare(fmt.Sprintf("id: '%d'", id), expVal, gotVal)
	}
}

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for ResourceAmount-ish types with different field names.
func (test test) checkResAmtTypes[T apiresource.ResourceAmountType[A], A apiresource.APIResource](fieldName string, exp map[string]int32, got []T) {
	test.t.Helper()
	gotResAmts := api.ResAmtTypesToStructs(got)
	test.checkResAmts(fieldName, exp, gotResAmts)
}
