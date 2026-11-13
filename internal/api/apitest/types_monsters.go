package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	TestGeneral
	NameVer
	appliedState     *TestAppliedState
	agility          *TestAgilityParams
	species          database.MonsterSpecies
	ctbIconType      database.CtbIconType
	distance         int32
	properties       []int32
	autoAbilities    []int32
	ronsoRages       []int32
	areas            []int32
	formations       []int32
	baseStats        map[string]int32
	regenAmounts     []TestRegenAmount
	items            *TestMonItems
	bribeChances     []apitypes.BribeChance
	equipment        *TestMonEquipment
	elemResists      []TestElemResist
	statusImmunities []int32
	statusResists    map[string]int32
	defaultState     *TestDefaultState
	abilities        []int32
}

func (e Monster) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestMonItems struct {
	itemDropChance int32
	items          map[string]*TestResAmount[apiresource.Typed]
	otherItems     []TestPossibleItem
}

type TestMonEquipment struct {
	abilitySlots      apitypes.MonsterEquipmentSlots
	attachedAbilities apitypes.MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}
