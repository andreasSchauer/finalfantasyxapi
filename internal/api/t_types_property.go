package api

type expProperty struct {
	testGeneral
	expUnique
	autoAbilities []int32
	monsters      []int32
}

func (e expProperty) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareProperties(test test, exp expProperty, got Property) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
}

type expElement struct {
	testGeneral
	expUnique
	statusProtection 	*int32
	autoAbilities 		[]int32
	playerAbilities 	[]int32
	overdriveAbilities 	[]int32
	itemAbilities 		[]int32
	enemyAbilities 		[]int32
	monstersWeak 		[]int32
	monstersHalved 		[]int32
	monstersImmune 		[]int32
	monstersAbsorb 		[]int32
}

func (e expElement) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareElements(test test, exp expElement, got Element) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResourcePtrs(test, "status protection", test.cfg.e.statusConditions.endpoint, exp.statusProtection, got.StatusProtection)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, "monsters weak", test.cfg.e.monsters.endpoint, exp.monstersWeak, got.MonstersWeak)
	checkResIDsInSlice(test, "monsters halved", test.cfg.e.monsters.endpoint, exp.monstersHalved, got.MonstersHalved)
	checkResIDsInSlice(test, "monsters immune", test.cfg.e.monsters.endpoint, exp.monstersImmune, got.MonstersImmune)
	checkResIDsInSlice(test, "monsters absorb", test.cfg.e.monsters.endpoint, exp.monstersAbsorb, got.MonstersAbsorb)
}


type expStat struct {
	testGeneral
	expUnique
	spheres 			[]int32
	autoAbilities 		[]int32
	playerAbilities 	[]int32
	overdriveAbilities 	[]int32
	itemAbilities 		[]int32
	triggerCommands 	[]int32
	statusConditions 	[]int32
	properties 			[]int32
}

func (e expStat) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareStats(test test, exp expStat, got Stat) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "spheres", test.cfg.e.spheres.endpoint, exp.spheres, got.Spheres)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "trigger commands", test.cfg.e.triggerCommands.endpoint, exp.triggerCommands, got.TriggerCommands)
	checkResIDsInSlice(test, "status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
}


type expModifier struct {
	testGeneral
	expUnique
	autoAbilities 		[]int32
	playerAbilities 	[]int32
	overdriveAbilities 	[]int32
	itemAbilities 		[]int32
	triggerCommands 	[]int32
	enemyAbilities 		[]int32
	statusConditions 	[]int32
	properties 			[]int32
}

func (e expModifier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareModifiers(test test, exp expModifier, got Modifier) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	checkResIDsInSlice(test, "player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, "overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, "item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, "trigger commands", test.cfg.e.triggerCommands.endpoint, exp.triggerCommands, got.TriggerCommands)
	checkResIDsInSlice(test, "enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, "status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
	checkResIDsInSlice(test, "properties", test.cfg.e.properties.endpoint, exp.properties, got.Properties)
}


type expStatusCondition struct {
	testGeneral
	expUnique
	autoAbilities 		[]int32
	inflictedBy 		*testStatusInteractions
	removedBy 			*testStatusInteractions
	monstersResistance 	[]int32
}

func (e expStatusCondition) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareStatusConditions(test test, exp expStatusCondition, got StatusCondition) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	compTestStructPtrs(test, "inflicted by", exp.inflictedBy, got.InflictedBy, compareStatusInteractions)
	compTestStructPtrs(test, "removed by", exp.removedBy, got.RemovedBy, compareStatusInteractions)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monstersResistance, got.MonstersResistance)
}


type testStatusInteractions struct {
	playerAbilities 		[]int32
	overdriveAbilities 		[]int32
	itemAbilities 			[]int32
	unspecifiedAbilities 	[]int32
	enemyAbilities 			[]int32
	statusConditions 		[]int32
}

func compareStatusInteractions(test test, fieldName string, exp testStatusInteractions, got StatusInteractions) {
	checkResIDsInSlice(test, fieldName+" - player abilities", test.cfg.e.playerAbilities.endpoint, exp.playerAbilities, got.PlayerAbilities)
	checkResIDsInSlice(test, fieldName+" - overdrive abilities", test.cfg.e.overdriveAbilities.endpoint, exp.overdriveAbilities, got.OverdriveAbilities)
	checkResIDsInSlice(test, fieldName+" - item abilities", test.cfg.e.itemAbilities.endpoint, exp.itemAbilities, got.ItemAbilities)
	checkResIDsInSlice(test, fieldName+" - unspecified abilities", test.cfg.e.unspecifiedAbilities.endpoint, exp.unspecifiedAbilities, got.UnspecifiedAbilities)
	checkResIDsInSlice(test, fieldName+" - enemy abilities", test.cfg.e.enemyAbilities.endpoint, exp.enemyAbilities, got.EnemyAbilities)
	checkResIDsInSlice(test, fieldName+" - status conditions", test.cfg.e.statusConditions.endpoint, exp.statusConditions, got.StatusConditions)
}

type expAgilityTier struct {
	testGeneral
	expIdOnly
	fromAgility int32
	toAgility	int32
	tickSpeed	int32
	monMaxICV	*int32
	monMinICV	*int32
	charMaxICV	*int32
	charMinICVs	[]AgilitySubtier
}

func (e expAgilityTier) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAgilityTiers(test test, exp expAgilityTier, got AgilityTier) {
	compareExpIdOnly(test, exp.expIdOnly, got.ID)
	compare(test, "from agility", exp.fromAgility, got.FromAgility)
	compare(test, "to agility", exp.toAgility, got.ToAgility)
	compare(test, "tick speed", exp.tickSpeed, got.TickSpeed)
	compare(test, "mon max icv", exp.monMaxICV, got.MonMaxICV)
	compare(test, "mon min icv", exp.monMinICV, got.MonMinICV)
	compare(test, "char max icv", exp.charMaxICV, got.CharMaxICV)
	compStructSlices(test, "char min icvs", exp.charMinICVs, got.CharMinICVs)
}