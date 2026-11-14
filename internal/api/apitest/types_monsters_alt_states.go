package apitest

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type testAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

func (e testAppliedState) Compare(test test, _ string, got apitypes.AppliedState) {
	test.compare("applied state condition", e.condition, got.Condition)
	test.compare("applied state isTemporary", e.isTemporary, got.IsTemporary)
	test.compIDAPIResourcePtrs("applied status", apilookups.E.StatusConditions.Endpoint, e.appliedStatus, got.AppliedStatus)
}

type testDefaultState struct {
	IsTemporary bool
	Change      *testAltStateChange
	Gain        *testAltStateGain
	Loss        *testAltStateLoss
}

func testMonsterDefaultState(test test, exp *testDefaultState, gotStates []apitypes.AlteredState) {
	if test.dontCheck != nil && test.dontCheck["default state"] {
		return
	}

	if !test.presenceDefaultAltStates(exp, gotStates) {
		return
	}

	test.compLength("altered states", len(gotStates))

	got := gotStates[0]
	if got.Condition != "default" {
		test.t.Errorf("first altered state must be default when another is applied, got: %s", got.Condition)
	}

	test.compare("default state is temporary", exp.IsTemporary, got.IsTemporary)
	test.compTestStructPtrs("alt state change", exp.Change, got.Change)
	test.compTestStructPtrs("alt state gain", exp.Gain, got.Gain)
	test.compTestStructPtrs("alt state loss", exp.Loss, got.Loss)
}

type testAltStateChange struct {
	Distance    *int32
	BaseStats   map[string]int32
	ElemResists []testElemResist
}

func ( e testAltStateChange) Compare(test test, fieldName string, got apitypes.AltStateChange) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.compare(desc+"distance", e.Distance, got.Distance)
	test.checkResAmtTypes(desc+"base stats", e.BaseStats, got.BaseStats)
	test.compTestStructSlices(desc+"elemental resists", e.ElemResists, got.ElemResists)
}

type testAltStateGain struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	StatusResists    map[string]int32
	Status           *apitypes.InflictedStatus
}

func (e testAltStateGain) Compare(test test, fieldName string, got apitypes.AltStateGain) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, e.StatusImmunities, got.StatusImmunities)
	test.checkResAmtTypes(desc+"status resists", e.StatusResists, got.StatusResists)
	test.compStructPtrs(desc+"added status", e.Status, got.Status)
}

type testAltStateLoss struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	Status           *int32
}

func (e testAltStateLoss) Compare(test test, fieldName string, got apitypes.AltStateLoss) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, e.StatusImmunities, got.StatusImmunities)
	test.compIDAPIResourcePtrs(desc+"removed status", apilookups.E.StatusConditions.Endpoint, e.Status, got.Status)
}
