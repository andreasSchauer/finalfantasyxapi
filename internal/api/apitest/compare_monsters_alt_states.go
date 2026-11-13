package apitest

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareMonsterAppliedStates(test Test, _ string, exp TestAppliedState, got apitypes.AppliedState) {
	test.compare("applied state condition", exp.condition, got.Condition)
	test.compare("applied state isTemporary", exp.isTemporary, got.IsTemporary)
	test.compIDAPIResourcePtrs("applied status", apilookups.E.StatusConditions.Endpoint, exp.appliedStatus, got.AppliedStatus)
}

func checkMonsterDefaultState(test Test, exp *TestDefaultState, gotStates []apitypes.AlteredState) {
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

func compareAltStateChanges(test Test, fieldName string, exp TestAltStateChange, got apitypes.AltStateChange) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.compare(desc+"distance", exp.Distance, got.Distance)
	test.checkResAmtTypes(desc+"base stats", exp.BaseStats, got.BaseStats)
	test.compTestStructSlices(desc+"elemental resists", exp.ElemResists, got.ElemResists, compareElemResists)
}

func compareAltStateGains(test Test, fieldName string, exp TestAltStateGain, got apitypes.AltStateGain) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, exp.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, exp.StatusImmunities, got.StatusImmunities)
	test.checkResAmtTypes(desc+"status resists", exp.StatusResists, got.StatusResists)
	test.compStructPtrs(desc+"added status", exp.Status, got.Status)
}

func compareAltStateLosses(test Test, fieldName string, exp TestAltStateLoss, got apitypes.AltStateLoss) {
	desc := fmt.Sprintf("%s -", fieldName)

	test.checkResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, exp.Properties, got.Properties)
	test.checkResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, exp.StatusImmunities, got.StatusImmunities)
	test.compIDAPIResourcePtrs(desc+"removed status", apilookups.E.StatusConditions.Endpoint, exp.Status, got.Status)
}
