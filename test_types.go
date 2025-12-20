package main

type testInOut struct {
	requestURL     string
	expectedStatus int
	expectedErr    string
}

type expectedNameVer struct {
	id      int32
	name    string
	version *int32
	lenMap  map[string]int
}

type expectedUnique struct {
	id     int32
	name   string
	lenMap map[string]int
}

type expectedList struct {
	count    int
	next     *string
	previous *string
	results  []string
}

type testCheck struct {
	name     string
	got      []HasAPIResource
	expected []string
}

type expResAreas struct {
	parentLocation    string
	parentSublocation string
	locBasedExpect
}

type locBasedExpect struct {
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
	description		string
	effect			string
	modeType		string
	fillRate		*float32
	actionsAmount	map[string]int32
}