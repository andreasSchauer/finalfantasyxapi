package main

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
	IsTemporary bool                 `json:"is_temporary"`
	Changes     []testAltStateChange `json:"changes"`
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
		test.t.Fatalf("%s: first altered state must be default when another is applied, got: %s", test.name, got.Condition)
	}

	compare(test, "default state is temporary", exp.IsTemporary, got.IsTemporary)
	compare(test, "def state changes length", len(exp.Changes), len(got.Changes))
	compTestStructSlices(test, "def state change", exp.Changes, got.Changes, compareMonsterAltStateChanges)
}


type testAltStateChange struct {
	AlterationType   string
	Distance         *int32
	Properties       []int32
	AutoAbilities    []int32
	BaseStats        map[string]int32
	ElemResists      []testElemResist
	StatusImmunities []int32
	StatusResists    map[string]int32
	AddedStatus      *InflictedStatus
	RemovedStatus    *int32
}

func compareMonsterAltStateChanges(test test, fieldName string, exp testAltStateChange, got AltStateChange) {
	desc := fmt.Sprintf("%s: %s -", fieldName, exp.AlterationType)

	compare(test, desc+"type", exp.AlterationType, string(got.AlterationType))
	compare(test, desc+"distance", exp.Distance, got.Distance)
	checkResAmtsNameVals(test, desc+"base stats", exp.BaseStats, got.BaseStats)
	checkResAmtsNameVals(test, desc+"status resists", exp.StatusResists, got.StatusResists)
	compStructPtrs(test, desc+"added status", exp.AddedStatus, got.AddedStatus)
	compIdApiResourcePtrs(test, desc+"removed status", test.cfg.e.statusConditions.endpoint, exp.RemovedStatus, got.RemovedStatus)
	compTestStructSlices(test, desc+"elemental resists", exp.ElemResists, got.ElemResists, compareElemResists)

	checkResIDsInSlice(test, desc+"properties", test.cfg.e.properties.endpoint, exp.Properties, got.Properties)
	checkResIDsInSlice(test, desc+"auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.AutoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, desc+"status immunities", test.cfg.e.statusConditions.endpoint, exp.StatusImmunities, got.StatusImmunities)
}
