package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type TestAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

type TestDefaultState struct {
	IsTemporary bool
	Change      *TestAltStateChange
	Gain        *TestAltStateGain
	Loss        *TestAltStateLoss
}

type TestAltStateChange struct {
	Distance    *int32
	BaseStats   map[string]int32
	ElemResists []TestElemResist
}

type TestAltStateGain struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	StatusResists    map[string]int32
	Status           *apitypes.InflictedStatus
}

type TestAltStateLoss struct {
	Properties       []int32
	AutoAbilities    []int32
	StatusImmunities []int32
	Status           *int32
}
