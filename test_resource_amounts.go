package main


// checks if stated ResourceAmount entries are in slices (used for baseStats, itemAmounts, monsterAmounts) and if their amount values match
func checkResAmtsInSlice[T ResourceAmount](test test, fieldName string, exp map[string]int32, got []T) {
	compLength(test, fieldName, len(got))

	gotMap := getResourceAmountTestMap(got)

	for path, exp := range exp {
		key := completeTestURL(test.cfg, path)
		got, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource %s", test.name, fieldName, key)
		}
		compare(test, key, exp, got)
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