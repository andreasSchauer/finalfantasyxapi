package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareCharacterClasses(test Test, exp CharacterClass, got apitypes.CharacterClass) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("category", exp.category, got.Category)
	test.checkResIDsInSlice("members", apilookups.E.PlayerUnits.Endpoint, exp.members, got.Members)
	test.checkResIDsInSlice("default abilities", apilookups.E.Abilities.Endpoint, exp.defaultAbilities, got.DefaultAbilities)
	test.checkResIDsInSlice("learnable abilities", apilookups.E.Abilities.Endpoint, exp.learnableAbilities, got.LearnableAbilities)
	test.checkResIDsInSlice("default overdrives", apilookups.E.Overdrives.Endpoint, exp.defaultOverdrives, got.DefaultOverdrives)
	test.checkResIDsInSlice("learnable overdrives", apilookups.E.Overdrives.Endpoint, exp.learnableOverdrives, got.LearnableOverdrives)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, exp.submenus, got.Submenus)
}
