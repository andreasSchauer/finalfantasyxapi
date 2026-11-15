package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

// checks presence of two pointers
func (test Test) presencePtrs[E, G any](fn string, exp *E, got *G) bool {
	test.T.Helper()

	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.T.Errorf("expected %s nil, got %v\n\n", fn, *got)
		return false

	case exp != nil && got == nil:
		test.T.Errorf("expected %s %v, got nil\n\n", fn, *exp)
		return false

	default:
		return true
	}
}

// checks presence of two pointers, where one of them is or has an api resource
func (test Test) presenceResPtrs[T apiresource.HasAPIResource](fn string, expPathPtr *string, gotResPtr *T) bool {
	test.T.Helper()

	switch {
	case expPathPtr == nil && gotResPtr == nil:
		return false

	case expPathPtr == nil && gotResPtr != nil:
		res := *gotResPtr
		gotURL := res.GetAPIResource().GetURL()
		test.T.Errorf("expected nil for %s, but got %s\n\n", fn, gotURL)
		return false

	case expPathPtr != nil && gotResPtr == nil:
		test.T.Errorf("expected %s %v, got nil\n\n", fn, *expPathPtr)
		return false

	default:
		return true
	}
}

// checks presence of two struct slices
func (test Test) presenceSlices[E, G any](fn string, exp []E, got []G) bool {
	test.T.Helper()

	switch {
	case exp == nil && got == nil:
		return false

	case exp == nil && got != nil:
		test.T.Errorf("expected %s nil, got %v\n\n", fn, got)
		return false

	case exp != nil && got == nil:
		test.T.Errorf("expected %s %v, got nil\n\n", fn, exp)
		return false

	default:
		return true
	}
}

func (test Test) presenceDefaultAltStates(exp *TestDefaultState, gotStates []apitypes.AlteredState) bool {
	test.T.Helper()
	switch {
	case exp == nil && len(gotStates) == 0:
		return false

	case exp == nil && len(gotStates) != 0:
		test.T.Errorf("expected default state to be nil, but got alt states\n\n")
		return false

	case exp != nil && len(gotStates) == 0:
		test.T.Errorf("expected default state to be not nil, but got no alt states\n\n")
		return false

	default:
		return true
	}
}
