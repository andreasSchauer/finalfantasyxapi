package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expCharacterClass struct {
	testGeneral
	expUnique
	category            string
	members             []int32
	defaultAbilities    []int32
	learnableAbilities  []int32
	defaultOverdrives   []int32
	learnableOverdrives []int32
	submenus            []int32
}

func (e expCharacterClass) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCharacterClasses(test test, exp expCharacterClass, got apitypes.CharacterClass) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("category", exp.category, got.Category)
	test.checkResIDsInSlice("members", apilookups.E.PlayerUnits.Endpoint, exp.members, got.Members)
	test.checkResIDsInSlice("default abilities", apilookups.E.Abilities.Endpoint, exp.defaultAbilities, got.DefaultAbilities)
	test.checkResIDsInSlice("learnable abilities", apilookups.E.Abilities.Endpoint, exp.learnableAbilities, got.LearnableAbilities)
	test.checkResIDsInSlice("default overdrives", apilookups.E.Overdrives.Endpoint, exp.defaultOverdrives, got.DefaultOverdrives)
	test.checkResIDsInSlice("learnable overdrives", apilookups.E.Overdrives.Endpoint, exp.learnableOverdrives, got.LearnableOverdrives)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, exp.submenus, got.Submenus)
}
