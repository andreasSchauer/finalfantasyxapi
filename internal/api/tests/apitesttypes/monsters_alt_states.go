package apitesttypes

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestAppliedState struct {
	Condition     string
	IsTemporary   bool
	AppliedStatus *int32
}

func (e TestAppliedState) Compare(test Test, _ string, got apitypes.AppliedState) {
	test.compare("applied state condition", e.Condition, got.Condition)
	test.compare("applied state isTemporary", e.IsTemporary, got.IsTemporary)
	test.compIDAPIResourcePtrs("applied status", apilookups.E.StatusConditions.Endpoint, e.AppliedStatus, got.AppliedStatus)
}

type TestDefaultState struct {
	IsTemporary bool
	Change      *TestAltStateChange
	Gain        *TestAltStateGain
	Loss        *TestAltStateLoss
}

func testMonsterDefaultState(test Test, exp *TestDefaultState, gotStates []apitypes.AlteredState) {
	if test.DontCheck != nil && test.DontCheck["default state"] {
		return
	}

	if !test.presenceDefaultAltStates(exp, gotStates) {
		return
	}

	test.compLength("altered states", len(gotStates))

	got := gotStates[0]
	if got.Condition != "default" {
		test.T.Errorf("first altered state must be default when another is applied, got: %s", got.Condition)
	}

	test.compare("default state is temporary", exp.IsTemporary, got.IsTemporary)
	test.compTestStructPtrs("alt state change", exp.Change, got.Change)
	test.compTestStructPtrs("alt state gain", exp.Gain, got.Gain)
	test.compTestStructPtrs("alt state loss", exp.Loss, got.Loss)
}

type TestAltStateChange struct {
	Distance    *int32
	BaseStats   map[string]int32
	ElemResists []TestElemResist
}

func (e TestAltStateChange) Compare(test Test, fn string, got apitypes.AltStateChange) {
	desc := fmt.Sprintf("%s -", fn)

	test.compare(desc+"distance", e.Distance, got.Distance)
	test.checkResAmtTypes(desc+"base stats", e.BaseStats, got.BaseStats)
	test.compTestStructSlices(desc+"elemental resists", e.ElemResists, got.ElemResists)
}

type TestAltStateGain struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	StatusResists    map[string]int32
	Status           *apitypes.InflictedStatus
}

func (e TestAltStateGain) Compare(test Test, fn string, got apitypes.AltStateGain) {
	desc := fmt.Sprintf("%s -", fn)

	test.CheckResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
	test.CheckResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, e.StatusImmunities, got.StatusImmunities)
	test.checkResAmtTypes(desc+"status resists", e.StatusResists, got.StatusResists)
	test.compStructPtrs(desc+"added status", e.Status, got.Status)
}

type TestAltStateLoss struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	Status           *int32
}

func (e TestAltStateLoss) Compare(test Test, fn string, got apitypes.AltStateLoss) {
	desc := fmt.Sprintf("%s -", fn)

	test.CheckResIDsInSlice(desc+"properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
	test.CheckResIDsInSlice(desc+"auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice(desc+"status immunities", apilookups.E.StatusConditions.Endpoint, e.StatusImmunities, got.StatusImmunities)
	test.compIDAPIResourcePtrs(desc+"removed status", apilookups.E.StatusConditions.Endpoint, e.Status, got.Status)
}
