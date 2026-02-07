package main


// checks if stated ResourceAmount entries are in slices (used for baseStats, itemAmounts, monsterAmounts) and if their amount values match
// currently bases the key for the map on GetName() method.
// a simple solution might be to do a resourceAmountTestMap function that either uses GetURL() or a GetRaMapKey() method instead of relying on GetName()
func checkResAmtsInSlice[T ResourceAmount](test test, fieldName string, expAmounts map[string]int32, gotAmounts []T) {
	expLen, ok := test.expLengths[fieldName]
	if !ok {
		return
	}
	compare(test, fieldName+" length", expLen, len(gotAmounts))

	gotMap := getResourceAmountMap(gotAmounts)

	for key, exp := range expAmounts {
		got, ok := gotMap[key]
		if !ok {
			test.t.Fatalf("%s: %s doesn't contain resource %s", test.name, fieldName, key)
		}
		compare(test, key, exp, got)
	}
}