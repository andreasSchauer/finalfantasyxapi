package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareCharacterClasses(test test, exp expCharacterClass, got api.CharacterClass) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "category", exp.category, got.Category)
	checkResIDsInSlice(test, "members", test.cfg.E.PlayerUnits.Endpoint, exp.members, got.Members)
	checkResIDsInSlice(test, "default abilities", test.cfg.E.Abilities.Endpoint, exp.defaultAbilities, got.DefaultAbilities)
	checkResIDsInSlice(test, "learnable abilities", test.cfg.E.Abilities.Endpoint, exp.learnableAbilities, got.LearnableAbilities)
	checkResIDsInSlice(test, "default overdrives", test.cfg.E.Overdrives.Endpoint, exp.defaultOverdrives, got.DefaultOverdrives)
	checkResIDsInSlice(test, "learnable overdrives", test.cfg.E.Overdrives.Endpoint, exp.learnableOverdrives, got.LearnableOverdrives)
	checkResIDsInSlice(test, "submenus", test.cfg.E.Submenus.Endpoint, exp.submenus, got.Submenus)
}
