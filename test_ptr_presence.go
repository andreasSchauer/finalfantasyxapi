package main


// checks presence of two pointers
func bothPtrsPresent[E, G any](test test, fieldName string, exp *E, got *G) bool {
	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.t.Fatalf("%s: expected %s nil, got %v", test.name, fieldName, *got)
		return false

	case exp != nil && got == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, *exp)
		return false

	default:
		return true
	}
}

// checks presence of two pointers, where one of them is or has an api resource
func bothResourcePtrsPresent[T HasAPIResource](test test, fieldName string, expPathPtr *string, gotResPtr *T) bool {
	switch {
	case expPathPtr == nil && gotResPtr == nil:
		return false

	case expPathPtr == nil && gotResPtr != nil:
		res := *gotResPtr
		gotURL := res.GetAPIResource().GetURL()
		test.t.Fatalf("%s: expected nil for %s, but got %s", test.name, fieldName, gotURL)
		return false

	case expPathPtr != nil && gotResPtr == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, *expPathPtr)
		return false

	default:
		return true
	}
}

// checks presence of two equally typed slices
func bothStructSlicesPresent[E, G any](test test, fieldName string, exp []E, got []G) bool {
	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.t.Fatalf("%s: expected %s nil, got %v", test.name, fieldName, got)
		return false

	case exp != nil && got == nil:
		test.t.Fatalf("%s: expected %s %v, got nil", test.name, fieldName, exp)
		return false

	default:
		return true
	}
}


func defaultAndAltStatesPresent(test test, exp *testDefaultState, gotStates []AlteredState) bool {
	switch {
	case exp == nil && len(gotStates) == 0:
		return false
		
		
	case exp == nil && len(gotStates) != 0:
		test.t.Fatalf("%s: expected default state to be nil, but got alt states", test.name)
		return false
	
	case exp != nil && len(gotStates) == 0:
		test.t.Fatalf("%s: expected default state to be not nil, but got no alt states", test.name)
		return false
		
	default:
		return true
	}
}