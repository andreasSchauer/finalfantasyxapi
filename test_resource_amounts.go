package main

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

// checks if stated ResourceAmount entries are in slices (used for baseStats, itemAmounts, monsterAmounts) and if their amount values match
func checkResAmtsStructs[T ResourceAmount](test test, fieldName string, exp map[string]int32, got []T) {
	compLength(test, fieldName, len(got))

	gotMap := getResourceAmountTestMap(got)

	for path, expVal := range exp {
		key := completeTestURL(test.cfg, path)
		gotVal, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource %s", test.name, fieldName, key)
		}
		compare(test, key, expVal, gotVal)
	}
}

func getResourceAmountTestMap[T ResourceAmount](items []T) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.GetAPIResource().GetURL()
		amountMap[key] = item.GetVal()
	}

	return amountMap
}
