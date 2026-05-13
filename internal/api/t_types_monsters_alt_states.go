package api

import (
	"fmt"
)

type testAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

func compareMonsterAppliedStates(test test, _ string, exp testAppliedState, got AppliedState) {
	compare(test, "applied state condition", exp.condition, got.Condition)
	compare(test, "applied state isTemporary", exp.isTemporary, got.IsTemporary)
	compIdApiResourcePtrs(test, "applied status", test.cfg.e.statusConditions.endpoint, exp.appliedStatus, got.AppliedStatus)
}

type testDefaultState struct {
	IsTemporary bool
	Change      *testAltStateChange
	Gain        *testAltStateGain
	Loss        *testAltStateLoss
}

func testMonsterDefaultState(test test, exp *testDefaultState, gotStates []AlteredState) {
	if test.dontCheck != nil && test.dontCheck["default state"] {
		return
	}

	if !defaultAndAltStatesPresent(test, exp, gotStates) {
		return
	}

	compLength(test, "altered states", len(gotStates))

	got := gotStates[0]
	if got.Condition != "default" {
		test.t.Errorf("first altered state must be default when another is applied, got: %s", got.Condition)
	}

	compare(test, "default state is temporary", exp.IsTemporary, got.IsTemporary)
	compTestStructPtrs(test, "alt state change", exp.Change, got.Change, compareAltStateChanges)
	compTestStructPtrs(test, "alt state gain", exp.Gain, got.Gain, compareAltStateGains)
	compTestStructPtrs(test, "alt state loss", exp.Loss, got.Loss, compareAltStateLosses)
}

type testAltStateChange struct {
	Distance    *int32
	BaseStats   map[string]int32
	ElemResists []testElemResist
}

func compareAltStateChanges(test test, fieldName string, exp testAltStateChange, got AltStateChange) {
	desc := fmt.Sprintf("%s -", fieldName)

	compare(test, desc+"distance", exp.Distance, got.Distance)
	checkResAmtTypes(test, desc+"base stats", exp.BaseStats, got.BaseStats)
	compTestStructSlices(test, desc+"elemental resists", exp.ElemResists, got.ElemResists, compareElemResists)
}

type testAltStateGain struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	StatusResists    map[string]int32
	Status           *InflictedStatus
}

func compareAltStateGains(test test, fieldName string, exp testAltStateGain, got AltStateGain) {
	desc := fmt.Sprintf("%s -", fieldName)

	checkResIDsInSlice(test, desc+"properties", test.cfg.e.properties.endpoint, exp.Properties, got.Properties)
	checkResIDsInSlice(test, desc+"auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.AutoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, desc+"status immunities", test.cfg.e.statusConditions.endpoint, exp.StatusImmunities, got.StatusImmunities)
	checkResAmtTypes(test, desc+"status resists", exp.StatusResists, got.StatusResists)
	compStructPtrs(test, desc+"added status", exp.Status, got.Status)
}

type testAltStateLoss struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	Status           *int32
}

func compareAltStateLosses(test test, fieldName string, exp testAltStateLoss, got AltStateLoss) {
	desc := fmt.Sprintf("%s -", fieldName)

	checkResIDsInSlice(test, desc+"properties", test.cfg.e.properties.endpoint, exp.Properties, got.Properties)
	checkResIDsInSlice(test, desc+"auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.AutoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, desc+"status immunities", test.cfg.e.statusConditions.endpoint, exp.StatusImmunities, got.StatusImmunities)
	compIdApiResourcePtrs(test, desc+"removed status", test.cfg.e.statusConditions.endpoint, exp.Status, got.Status)
}
