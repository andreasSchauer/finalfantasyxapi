package api

// checks presence of two pointers
func bothPtrsPresent[E, G any](test test, fieldName string, exp *E, got *G) bool {
	test.t.Helper()
	
	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.t.Errorf("expected %s nil, got %v\n\n", fieldName, *got)
		return false

	case exp != nil && got == nil:
		test.t.Errorf("expected %s %v, got nil\n\n", fieldName, *exp)
		return false

	default:
		return true
	}
}

// checks presence of two pointers, where one of them is or has an api resource
func bothResourcePtrsPresent[T HasAPIResource](test test, fieldName string, expPathPtr *string, gotResPtr *T) bool {
	test.t.Helper()

	switch {
	case expPathPtr == nil && gotResPtr == nil:
		return false

	case expPathPtr == nil && gotResPtr != nil:
		res := *gotResPtr
		gotURL := res.GetAPIResource().GetURL()
		test.t.Errorf("expected nil for %s, but got %s\n\n", fieldName, gotURL)
		return false

	case expPathPtr != nil && gotResPtr == nil:
		test.t.Errorf("expected %s %v, got nil\n\n", fieldName, *expPathPtr)
		return false

	default:
		return true
	}
}

// checks presence of two struct slices
func bothSlicesPresent[E, G any](test test, fieldName string, exp []E, got []G) bool {
	test.t.Helper()

	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.t.Errorf("expected %s nil, got %v\n\n", fieldName, got)
		return false

	case exp != nil && got == nil:
		test.t.Errorf("expected %s %v, got nil\n\n", fieldName, exp)
		return false

	default:
		return true
	}
}

func defaultAndAltStatesPresent(test test, exp *testDefaultState, gotStates []AlteredState) bool {
	test.t.Helper()
	switch {
	case exp == nil && len(gotStates) == 0:
		return false

	case exp == nil && len(gotStates) != 0:
		test.t.Errorf("expected default state to be nil, but got alt states\n\n")
		return false

	case exp != nil && len(gotStates) == 0:
		test.t.Errorf("expected default state to be not nil, but got no alt states\n\n")
		return false

	default:
		return true
	}
}
