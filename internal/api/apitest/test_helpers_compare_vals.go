package apitest

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
)

// checks if two standard-type values or pointers of a field are equal
func (test test) compare(fieldName string, exp, got any) {
	t := test.t
	dontCheck := test.dontCheck
	t.Helper()

	if dontCheck != nil && dontCheck[fieldName] {
		return
	}

	switch e := exp.(type) {
	case int:
		g, ok := got.(int)
		if !ok {
			t.Errorf("%s type mismatch: expected int, got %T\n\n", fieldName, got)
		}
		test.compInt(fieldName, e, g)

	case int32:
		g, ok := got.(int32)
		if !ok {
			t.Errorf("%s type mismatch: expected int32, got %T\n\n", fieldName, got)
		}
		test.compInt32(fieldName, e, g)

	case float32:
		g, ok := got.(float32)
		if !ok {
			t.Errorf("%s type mismatch: expected float32, got %T\n\n", fieldName, got)
		}
		test.compFloat32(fieldName, e, g)

	case float64:
		g, ok := got.(float64)
		if !ok {
			t.Errorf("%s type mismatch: expected float64, got %T\n\n", fieldName, got)
		}
		test.compFloat64(fieldName, e, g)

	case string:
		g, ok := got.(string)
		if !ok {
			t.Errorf("%s type mismatch: expected string, got %T\n\n", fieldName, got)
		}
		test.compString(fieldName, e, g)

	case bool:
		g, ok := got.(bool)
		if !ok {
			t.Errorf("%s type mismatch: expected bool, got %T\n\n", fieldName, got)
		}
		test.compBool(fieldName, e, g)

	case *int32:
		g, ok := got.(*int32)
		if !ok {
			t.Errorf("%s type mismatch: expected *int32, got %T\n\n", fieldName, got)
		}
		test.compInt32Ptr(fieldName, e, g)

	case *float32:
		g, ok := got.(*float32)
		if !ok {
			t.Errorf("%s type mismatch: expected *float32, got %T\n\n", fieldName, got)
		}
		test.compFloat32Ptr(fieldName, e, g)

	case *float64:
		g, ok := got.(*float64)
		if !ok {
			t.Errorf("%s type mismatch: expected *float64, got %T\n\n", fieldName, got)
		}
		test.compFloat64Ptr(fieldName, e, g)

	case *string:
		g, ok := got.(*string)
		if !ok {
			t.Errorf("%s type mismatch: expected *string, got %T\n\n", fieldName, got)
		}
		test.compStringPtr(fieldName, e, g)

	default:
		t.Errorf("unsupported type for %s: %T", fieldName, exp)
	}
}

func (test test) compInt(fieldName string, exp, got int) {
	test.t.Helper()
	if exp != got {
		test.t.Errorf("expected %s %d, got %d\n\n", fieldName, exp, got)
	}
}

func (test test) compInt32(fieldName string, exp, got int32) {
	test.t.Helper()
	if exp != got {
		test.t.Errorf("expected %s %d, got %d\n\n", fieldName, exp, got)
	}
}

func (test test) compFloat32(fieldName string, exp, got float32) {
	test.t.Helper()
	if exp != got {
		test.t.Errorf("expected %s %.2f, got %.2f\n\n", fieldName, exp, got)
	}
}

func (test test) compFloat64(fieldName string, exp, got float64) {
	test.t.Helper()
	if exp != got {
		test.t.Errorf("expected %s %f, got %f\n\n", fieldName, exp, got)
	}
}

func (test test) compString(fieldName, exp, got string) {
	test.t.Helper()
	if exp != "" && exp != got {
		test.t.Errorf("expected %s %s, got %s\n\n", fieldName, exp, got)
	}
}

func (test test) compBool(fieldName string, exp, got bool) {
	test.t.Helper()
	if exp != got {
		test.t.Errorf("expected %s %t, got %t\n\n", fieldName, exp, got)
	}
}

// can be used for any direct pointer comparisons
func (test test) compPtr[T any](fieldName string, exp, got *T, compFn compFnVals[T]) {
	test.t.Helper()
	if !test.presencePtrs(fieldName, exp, got) {
		return
	}
	compFn(fieldName, *exp, *got)
}

func (test test) compInt32Ptr(fieldName string, exp, got *int32) {
	test.t.Helper()
	test.compPtr(fieldName, exp, got, test.compInt32)
}

func (test test) compFloat32Ptr(fieldName string, exp, got *float32) {
	test.t.Helper()
	test.compPtr(fieldName, exp, got, test.compFloat32)
}

func (test test) compFloat64Ptr(fieldName string, exp, got *float64) {
	test.t.Helper()
	test.compPtr(fieldName, exp, got, test.compFloat64)
}

func (test test) compStringPtr(fieldName string, exp, got *string) {
	test.t.Helper()
	test.compPtr(fieldName, exp, got, test.compString)
}

// compares the length of the gotSlice with the length found in expLengths[fieldName], if given
func (test test) compLength(fieldName string, got int) {
	test.t.Helper()
	exp, ok := test.expLengths[fieldName]
	if ok {
		test.compare(fieldName+" length", exp, got)
	}
}

// compares the expected pagination path with the got url
func (test test) compPageURL(fieldName string, expPathPtr, gotURLPtr *string) {
	test.t.Helper()
	if test.dontCheck != nil && test.dontCheck[fieldName] {
		return
	}

	var expURLPtr *string

	if expPathPtr != nil {
		expPath := *expPathPtr
		expURL := ah.CompleteTestURL(expPath)
		expURLPtr = &expURL
	}

	test.compare(fieldName, expURLPtr, gotURLPtr)
}