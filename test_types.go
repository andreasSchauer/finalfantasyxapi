package main


type testInOut struct {
	requestURL     string
	expectedStatus int
	expectedErr    string
}

type expectedSingle struct {
	id              int32
	name            string
	version         *int32
	expectedLen 	map[string]int
}

type expectedList struct {
	count		int32
	results		[]string
}

type testCheck struct {
    name     string
    got      []HasAPIResource
    expected []string
}

type expectedResourcesAreas struct {
	parentLocation    	string
	parentSublocation 	string
	sidequest      		*string
	connectedAreas 		[]string
	characters     		[]string
	aeons          		[]string
	shops          		[]string
	treasures      		[]string
	monsters       		[]string
	formations     		[]string
	bgMusic        		[]string
	cuesMusic      		[]string
	fmvsMusic      		[]string
	bossMusic      		[]string
	fmvs           		[]string
}


