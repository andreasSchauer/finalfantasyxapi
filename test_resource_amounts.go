package main


// expects a map[res.GetName()]Amount and checks, if all stated resources are present and their respective amounts match
func checkResAmtsNameVals[T ResourceAmount](test test, fieldName string, exp map[string]int32, got []T) {
	compLength(test, fieldName, len(got))

	gotMap := getResourceAmountMap(got)

	for key, expVal := range exp {
		gotVal, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource %s", test.name, fieldName, key)
		}
		compare(test, key, expVal, gotVal)
	}
}