package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	TestGeneral
	NameVer
	AppliedState     *TestAppliedState
	Agility          *TestAgilityParams
	Species          database.MonsterSpecies
	CtbIconType      database.CtbIconType
	Distance         int32
	Properties       []int32
	AutoAbilities    []int32
	RonsoRages       []int32
	Areas            []int32
	Formations       []int32
	BaseStats        map[string]int32
	RegenAmounts     []TestRegenAmount
	Items            *TestMonItems
	BribeChances     []apitypes.BribeChance
	Equipment        *TestMonEquipment
	ElemResists      []TestElemResist
	StatusImmunities []int32
	StatusResists    map[string]int32
	DefaultState     *TestDefaultState
	Abilities        []int32
}

func (e Monster) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Monster) Compare(test Test, got apitypes.Monster) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("species", string(e.Species), got.Species)
	test.compare("ctb icon type", string(e.CtbIconType), got.CTBIconType)
	test.compare("distance", e.Distance, got.Distance)
	test.checkResAmtTypes("base stats", e.BaseStats, got.BaseStats)
	test.checkTestStructsInSlice("regen amounts", e.RegenAmounts, got.RegenAmounts)
	test.checkResAmtTypes("status resists", e.StatusResists, got.StatusResists)
	test.compTestStructPtrs("agility params", e.Agility, got.AgilityParameters)
	test.compStructSlices("bribe chances", e.BribeChances, got.BribeChances)
	test.compTestStructSlices("elemental resists", e.ElemResists, got.ElemResists)
	test.compTestStructPtrs("applied state", e.AppliedState, got.AppliedState)
	testMonsterDefaultState(test, e.DefaultState, got.AlteredStates)
	test.compTestStructPtrs("items", e.Items, got.Items)
	test.compTestStructPtrs("equipment", e.Equipment, got.Equipment)
	test.CheckResIDsInSlice("properties", apilookups.E.Properties.Endpoint, e.Properties, got.Properties)
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice("ronso rages", apilookups.E.RonsoRages.Endpoint, e.RonsoRages, got.RonsoRages)
	test.CheckResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
	test.CheckResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, e.Formations, got.Formations)
	test.CheckResIDsInSlice("status immunities", apilookups.E.StatusConditions.Endpoint, e.StatusImmunities, got.StatusImmunities)
	test.CheckResIDsInSlice("abilities", apilookups.E.Abilities.Endpoint, e.Abilities, got.Abilities)
}

type TestMonItems struct {
	ItemDropChance int32
	Items          map[string]*TestResAmount[apiresource.Typed]
	OtherItems     []TestPossibleItem
}

func (e TestMonItems) Compare(test Test, fieldName string, got apitypes.MonsterItems) {
	test.T.Helper()
	itemMap := e.Items

	test.compare("item drop chance", e.ItemDropChance, got.DropChance)
	test.compTestStructPtrs("steal common", itemMap["steal common"], got.StealCommon)
	test.compTestStructPtrs("steal rare", itemMap["steal rare"], got.StealRare)
	test.compTestStructPtrs("drop common", itemMap["drop common"], got.DropCommon)
	test.compTestStructPtrs("drop rare", itemMap["drop rare"], got.DropRare)
	test.compTestStructPtrs("sec drop common", itemMap["sec drop common"], got.SecondaryDropCommon)
	test.compTestStructPtrs("sec drop rare", itemMap["sec drop rare"], got.SecondaryDropRare)
	test.compTestStructPtrs("bribe", itemMap["bribe"], got.Bribe)
	test.compTestStructSlices("other items", e.OtherItems, got.OtherItems)
}

type TestMonEquipment struct {
	AbilitySlots      apitypes.MonsterEquipmentSlots
	AttachedAbilities apitypes.MonsterEquipmentSlots
	WeaponAbilities   []int32
	ArmorAbilities    []int32
}

func (e TestMonEquipment) Compare(test Test, _ string, got apitypes.MonsterEquipment) {
	test.T.Helper()
	test.compStructs("ability slots", e.AbilitySlots, got.AbilitySlots)
	test.compStructs("attached abilities", e.AttachedAbilities, got.AttachedAbilities)
	test.CheckResIDsInSlice("weapon abilities", apilookups.E.AutoAbilities.Endpoint, e.WeaponAbilities, got.WeaponAbilities)
	test.CheckResIDsInSlice("armor abilities", apilookups.E.AutoAbilities.Endpoint, e.ArmorAbilities, got.ArmorAbilities)
}
