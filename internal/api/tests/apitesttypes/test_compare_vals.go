package apitesttypes

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
)

// checks if two standard-type values or pointers of a field are equal
func (test Test) compare(fn string, exp, got any) {
	t := test.T
	dontCheck := test.DontCheck
	t.Helper()

	if dontCheck != nil && dontCheck[fn] {
		return
	}

	switch e := exp.(type) {
	case int:
		g, ok := got.(int)
		if !ok {
			t.Errorf("%s type mismatch: expected int, got %T\n\n", fn, got)
		}
		test.compInt(fn, e, g)

	case int32:
		g, ok := got.(int32)
		if !ok {
			t.Errorf("%s type mismatch: expected int32, got %T\n\n", fn, got)
		}
		test.compInt32(fn, e, g)

	case float32:
		g, ok := got.(float32)
		if !ok {
			t.Errorf("%s type mismatch: expected float32, got %T\n\n", fn, got)
		}
		test.compFloat32(fn, e, g)

	case float64:
		g, ok := got.(float64)
		if !ok {
			t.Errorf("%s type mismatch: expected float64, got %T\n\n", fn, got)
		}
		test.compFloat64(fn, e, g)

	case string:
		g, ok := got.(string)
		if !ok {
			t.Errorf("%s type mismatch: expected string, got %T\n\n", fn, got)
		}
		test.compString(fn, e, g)

	case bool:
		g, ok := got.(bool)
		if !ok {
			t.Errorf("%s type mismatch: expected bool, got %T\n\n", fn, got)
		}
		test.compBool(fn, e, g)

	case *int32:
		g, ok := got.(*int32)
		if !ok {
			t.Errorf("%s type mismatch: expected *int32, got %T\n\n", fn, got)
		}
		test.compInt32Ptr(fn, e, g)

	case *float32:
		g, ok := got.(*float32)
		if !ok {
			t.Errorf("%s type mismatch: expected *float32, got %T\n\n", fn, got)
		}
		test.compFloat32Ptr(fn, e, g)

	case *float64:
		g, ok := got.(*float64)
		if !ok {
			t.Errorf("%s type mismatch: expected *float64, got %T\n\n", fn, got)
		}
		test.compFloat64Ptr(fn, e, g)

	case *string:
		g, ok := got.(*string)
		if !ok {
			t.Errorf("%s type mismatch: expected *string, got %T\n\n", fn, got)
		}
		test.compStringPtr(fn, e, g)

	default:
		t.Errorf("unsupported type for %s: %T", fn, exp)
	}
}

func (test Test) compInt(fn string, exp, got int) {
	test.T.Helper()
	if exp != got {
		test.T.Errorf("expected %s %d, got %d\n\n", fn, exp, got)
	}
}

func (test Test) compInt32(fn string, exp, got int32) {
	test.T.Helper()
	if exp != got {
		test.T.Errorf("expected %s %d, got %d\n\n", fn, exp, got)
	}
}

func (test Test) compFloat32(fn string, exp, got float32) {
	test.T.Helper()
	if exp != got {
		test.T.Errorf("expected %s %.2f, got %.2f\n\n", fn, exp, got)
	}
}

func (test Test) compFloat64(fn string, exp, got float64) {
	test.T.Helper()
	if exp != got {
		test.T.Errorf("expected %s %f, got %f\n\n", fn, exp, got)
	}
}

func (test Test) compString(fn, exp, got string) {
	test.T.Helper()
	if exp != "" && exp != got {
		test.T.Errorf("expected %s %s, got %s\n\n", fn, exp, got)
	}
}

func (test Test) compBool(fn string, exp, got bool) {
	test.T.Helper()
	if exp != got {
		test.T.Errorf("expected %s %t, got %t\n\n", fn, exp, got)
	}
}

// can be used for any direct pointer comparisons
func (test Test) compPtr[T any](fn string, exp, got *T, compFn compFnVals[T]) {
	test.T.Helper()
	if !test.presencePtrs(fn, exp, got) {
		return
	}
	compFn(fn, *exp, *got)
}

func (test Test) compInt32Ptr(fn string, exp, got *int32) {
	test.T.Helper()
	test.compPtr(fn, exp, got, test.compInt32)
}

func (test Test) compFloat32Ptr(fn string, exp, got *float32) {
	test.T.Helper()
	test.compPtr(fn, exp, got, test.compFloat32)
}

func (test Test) compFloat64Ptr(fn string, exp, got *float64) {
	test.T.Helper()
	test.compPtr(fn, exp, got, test.compFloat64)
}

func (test Test) compStringPtr(fn string, exp, got *string) {
	test.T.Helper()
	test.compPtr(fn, exp, got, test.compString)
}

// compares the length of the gotSlice with the length found in expLengths[fn], if given
func (test Test) compLength(fn string, got int) {
	test.T.Helper()
	exp, ok := test.ExpLengths[fn]
	if ok {
		test.compare(childFn(fn, "length"), exp, got)
	}
}

// CompPageURL compares the expected pagination path with the got url
func (test Test) CompPageURL(fn string, expPathPtr, gotURLPtr *string) {
	test.T.Helper()
	if test.DontCheck != nil && test.DontCheck[fn] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := ah.CompleteTestURL(expPath)
		expURLPtr = &expURL
	}

	test.compare(fn, expURLPtr, gotURLPtr)
}
