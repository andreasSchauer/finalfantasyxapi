package apitest

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type testAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

func compareMonsterAppliedStates(test test, _ string, exp testAppliedState, got apitypes.AppliedState) {
	test.compare("applied state condition", exp.condition, got.Condition)
	test.compare("applied state isTemporary", exp.isTemporary, got.IsTemporary)
	test.compIDAPIResourcePtrs("applied status", apilookups.E.StatusConditions.Endpoint, exp.appliedStatus, got.AppliedStatus)
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
	test.compTestStructPtrs("alt state change", exp.Change, got.Change, compareAltStateChanges)
	test.compTestStructPtrs("alt state gain", exp.Gain, got.Gain, compareAltStateGains)
	test.compTestStructPtrs("alt state loss", exp.Loss, got.Loss, compareAltStateLosses)
}

type testAltStateChange struct {
	Distance    *int32
	BaseStats   map[string]int32
	ElemResists []testElemResist
}

func compareAltStateChanges(test test, fieldName string, exp testAltStateChange, got apitypes.AltStateChange) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.compare(desc+"distance", exp.Distance, got.Distance)
	test.checkResAmtTypes(desc+"base stats", exp.BaseStats, got.BaseStats)
	test.compTestStructSlices(desc+"elemental resists", exp.ElemResists, got.ElemResists, compareElemResists)
}

type testAltStateGain struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	StatusResists    map[string]int32
	Status           *apitypes.InflictedStatus
}

func compareAltStateGains(test test, fieldName string, exp testAltStateGain, got apitypes.AltStateGain) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, exp.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, exp.StatusImmunities, got.StatusImmunities)
	test.checkResAmtTypes(desc+"status resists", exp.StatusResists, got.StatusResists)
	test.compStructPtrs(desc+"added status", exp.Status, got.Status)
}

type testAltStateLoss struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	Status           *int32
}

func compareAltStateLosses(test test, fieldName string, exp testAltStateLoss, got apitypes.AltStateLoss) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, exp.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, exp.StatusImmunities, got.StatusImmunities)
	test.compIDAPIResourcePtrs(desc+"removed status", apilookups.E.StatusConditions.Endpoint, exp.Status, got.Status)
}
