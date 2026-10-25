package apiTest

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
)

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for pure ResourceAmount fields.
func checkResAmts[A api.APIResource](test test, fieldName string, exp map[string]int32, got []api.ResourceAmount[A]) {
	test.t.Helper()
	compLength(test, fieldName, len(got))

	gotMap := api.GetResAmtMap(got)

	for key, expVal := range exp {
		gotVal, ok := gotMap[key]
		if !ok {
			test.t.Errorf("%s doesn't contain resource with key '%s'.\n\n", fieldName, key)
		}
		compare(test, key, expVal, gotVal)
	}
}

func checkResAmtsID[A api.APIResource](test test, fieldName string, exp map[int32]int32, got []api.ResourceAmount[A]) {
	test.t.Helper()
	compLength(test, fieldName, len(got))

	gotMap := api.GetResAmtIdMap(got)

	for id, expVal := range exp {
		gotVal, ok := gotMap[id]
		if !ok {
			test.t.Errorf("%s doesn't contain resource with id '%d'.\n\n", fieldName, id)
		}
		compare(test, fmt.Sprintf("id: '%d'", id), expVal, gotVal)
	}
}

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for ResourceAmount-ish types with different field names.
func checkResAmtTypes[T api.ResourceAmountType[A], A api.APIResource](test test, fieldName string, exp map[string]int32, got []T) {
	test.t.Helper()
	gotResAmts := api.ResAmtTypesToStructs(got)
	checkResAmts(test, fieldName, exp, gotResAmts)
}
