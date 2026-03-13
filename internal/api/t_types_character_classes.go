package api

type expCharacterClass struct {
	testGeneral
	expUnique
	category            string
	members             []string
	defaultAbilities    []int32
	learnableAbilities  []int32
	defaultOverdrives   []int32
	learnableOverdrives []int32
	submenus            []int32
}

func (e expCharacterClass) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCharacterClasses(test test, exp expCharacterClass, got CharacterClass) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "category", exp.category, got.Category)
	checkResPathsInSlice(test, "members", exp.members, got.Members)
	checkResIDsInSlice(test, "default abilities", test.cfg.e.abilities.endpoint, exp.defaultAbilities, got.DefaultAbilities)
	checkResIDsInSlice(test, "learnable abilities", test.cfg.e.abilities.endpoint, exp.learnableAbilities, got.LearnableAbilities)
	checkResIDsInSlice(test, "default overdrives", test.cfg.e.overdrives.endpoint, exp.defaultOverdrives, got.DefaultOverdrives)
	checkResIDsInSlice(test, "learnable overdrives", test.cfg.e.overdrives.endpoint, exp.learnableOverdrives, got.LearnableOverdrives)
	checkResIDsInSlice(test, "submenus", test.cfg.e.submenus.endpoint, exp.submenus, got.Submenus)
}
