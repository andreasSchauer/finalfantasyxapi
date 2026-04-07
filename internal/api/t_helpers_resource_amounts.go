package api

import (
	"fmt"
)

// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for pure ResourceAmount fields.
func checkResAmts[A APIResource](test test, fieldName string, exp map[string]int32, got []ResourceAmount[A]) {
	compLength(test, fieldName, len(got))

	gotMap := getResAmtMap(got)

	for key, expVal := range exp {
		gotVal, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource with key '%s'.", test.name, fieldName, key)
		}
		compare(test, key, expVal, gotVal)
	}
}


func checkResAmtsID[A APIResource](test test, fieldName string, exp map[int32]int32, got []ResourceAmount[A]) {
	compLength(test, fieldName, len(got))

	gotMap := getResAmtIdMap(got)

	for id, expVal := range exp {
		gotVal, ok := gotMap[id]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource with id '%d'.", test.name, fieldName, id)
		}
		compare(test, fmt.Sprintf("id: '%d'", id), expVal, gotVal)
	}
}


// expects a map[res.GetKey()]Amount and checks, if all stated resources are present and their respective amounts match. Used for ResourceAmount-ish types with different field names.
func checkResAmtTypes[T ResourceAmountType[A], A APIResource](test test, fieldName string, exp map[string]int32, got []T) {
	gotResAmts := resAmtTypesToStructs(got)
	checkResAmts(test, fieldName, exp, gotResAmts)
}