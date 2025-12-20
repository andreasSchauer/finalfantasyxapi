package main

type testInOut struct {
	requestURL     	string
	expectedStatus 	int
	expectedErr    	string
	dontCheck 		map[string]bool
}

type expNameVer struct {
	id      int32
	name    string
	version *int32
	lenMap  map[string]int
}

type expUnique struct {
	id     int32
	name   string
	lenMap map[string]int
}

type expList struct {
	count    int
	next     *string
	previous *string
	results  []string
}

type resListTest struct {
	name 		string
	exp  		[]string
	got  		[]HasAPIResource
}

func newResListTest[T HasAPIResource](fieldName string, exp []string, got []T) resListTest {
	return resListTest{
		name: fieldName,
		exp:  exp,
		got:  toHasAPIResSlice(got),
	}
}

type expResAreas struct {
	parentLocation    string
	parentSublocation string
	expLocBased
}

type expLocBased struct {
	sidequest      *string
	connectedAreas []string
	characters     []string
	aeons          []string
	shops          []string
	treasures      []string
	monsters       []string
	formations     []string
	bgMusic        []string
	cuesMusic      []string
	fmvsMusic      []string
	bossMusic      []string
	fmvs           []string
}

type expResOverdriveModes struct {
	description   string
	effect        string
	modeType      string
	fillRate      *float32
	actionsAmount map[string]int32
}
