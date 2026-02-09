package main

// checks if two standard-type values or pointers of a field are equal
func compare(test test, fieldName string, exp, got any) {
	t := test.t
	testName := test.name
	dontCheck := test.dontCheck
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	switch e := exp.(type) {

	case int:
		g, ok := got.(int)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected int, got %T", testName, fieldName, got)
		}
		compInt(test, fieldName, e, g)

	case int32:
		g, ok := got.(int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected int32, got %T", testName, fieldName, got)
		}
		compInt32(test, fieldName, e, g)

	case float32:
		g, ok := got.(float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected float32, got %T", testName, fieldName, got)
		}
		compFloat32(test, fieldName, e, g)

	case string:
		g, ok := got.(string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected string, got %T", testName, fieldName, got)
		}
		compString(test, fieldName, e, g)

	case bool:
		g, ok := got.(bool)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected bool, got %T", testName, fieldName, got)
		}
		compBool(test, fieldName, e, g)

	case *int32:
		g, ok := got.(*int32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *int32, got %T", testName, fieldName, got)
		}
		compInt32Ptr(test, fieldName, e, g)

	case *float32:
		g, ok := got.(*float32)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *float32, got %T", testName, fieldName, got)
		}
		compFloat32Ptr(test, fieldName, e, g)

	case *string:
		g, ok := got.(*string)
		if !ok {
			t.Fatalf("%s: %s type mismatch: expected *string, got %T", testName, fieldName, got)
		}
		compStringPtr(test, fieldName, e, g)

	default:
		t.Fatalf("%s: unsupported type for %s: %T", testName, fieldName, exp)
	}
}


// can be used for any direct pointer comparisons
func compPtr[T any](test test, fieldName string, exp, got *T, compFunc func(test, string, T, T)) {
	if !bothPtrsPresent(test, fieldName, exp, got) {
		return
	}
	compFunc(test, fieldName, *exp, *got)
}


func compInt(test test, fieldName string, exp, got int) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %d, got %d", test.name, fieldName, exp, got)
	}
}


func compInt32(test test, fieldName string, exp, got int32) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %d, got %d", test.name, fieldName, exp, got)
	}
}


func compFloat32(test test, fieldName string, exp, got float32) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %.2f, got %.2f", test.name, fieldName, exp, got)
	}
}


func compString(test test, fieldName, exp, got string) {
	if exp != "" && exp != got {
		test.t.Fatalf("%s: expected %s %s, got %s", test.name, fieldName, exp, got)
	}
}


func compBool(test test, fieldName string, exp, got bool) {
	if exp != got {
		test.t.Fatalf("%s: expected %s %t, got %t", test.name, fieldName, exp, got)
	}
}


func compInt32Ptr(test test, fieldName string, exp, got *int32) {
	compPtr(test, fieldName, exp, got, compInt32)
}


func compFloat32Ptr(test test, fieldName string, exp, got *float32) {
	compPtr(test, fieldName, exp, got, compFloat32)
}


func compStringPtr(test test, fieldName string, exp, got *string) {
	compPtr(test, fieldName, exp, got, compString)
}


// compares the length of the gotSlice with the length found in expLengths[fieldName], if given
func compLength(test test, fieldName string, got int) {
	exp, ok := test.expLengths[fieldName]
	if ok {
		compare(test, fieldName + " length", exp, got)
	}
}


// compares the expected pagination path with the got url
func compPageURL(test test, fieldName string, expPathPtr, gotURLPtr *string) {
	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := completeTestURL(test.cfg, expPath)
		expURLPtr = &expURL
	}

	compare(test, fieldName, expURLPtr, gotURLPtr)
}