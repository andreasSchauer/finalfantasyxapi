package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type expAbility struct {
	testGeneral
	expNameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	abilityType        database.AbilityType
	typedAbility       string
	monsters           []int32
	battleInteractions []expBattleInteraction
}

func (e expAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAbilities(test test, exp expAbility, got Ability) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "ability type", string(exp.abilityType), string(got.Type))
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)

	compPathApiResource(test, "typed ability", exp.typedAbility, got.TypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expEnemyAbility struct {
	testGeneral
	expNameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	monsters           []int32
	battleInteractions []expBattleInteraction
}

func (e expEnemyAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareEnemyAbilities(test test, exp expEnemyAbility, got EnemyAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expItemAbility struct {
	testGeneral
	expUnique
	rank                *int32
	appearsInHelpBar    bool
	canCopyCat          bool
	untypedAbility      int32
	item                int32
	category            database.ItemCategory
	CanUseOutsideBattle bool
	battleInteractions  []expBattleInteraction
}

func (e expItemAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareItemAbilities(test test, exp expItemAbility, got ItemAbility) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResource(test, "item", test.cfg.e.items.endpoint, exp.item, got.Item)
	compare(test, "category", string(exp.category), got.Category)
	compare(test, "can use outside battle", exp.CanUseOutsideBattle, got.CanUseOutsideBattle)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expPlayerAbility struct {
	testGeneral
	expNameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	topmenu            *int32
	submenu            *int32
	openSubmenu        *int32
	stdChar            *int32
	expChar            *int32
	monsters           []int32
	battleInteractions []expBattleInteraction
}

func (e expPlayerAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePlayerAbilities(test test, exp expPlayerAbility, got PlayerAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResourcePtrs(test, "submenu", test.cfg.e.submenus.endpoint, exp.submenu, got.Submenu)
	compIdApiResourcePtrs(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	compIdApiResourcePtrs(test, "std sg char", test.cfg.e.characters.endpoint, exp.stdChar, got.StandardGridCharacter)
	compIdApiResourcePtrs(test, "exp sg char", test.cfg.e.characters.endpoint, exp.expChar, got.ExpertGridCharacter)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expOverdriveAbility struct {
	testGeneral
	expNameVer
	rank               *int32
	untypedAbility     int32
	overdriveCommand   *int32
	overdrives         []int32
	battleInteractions []expBattleInteraction
}

func (e expOverdriveAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveAbilities(test test, exp expOverdriveAbility, got OverdriveAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "overdrive command", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "overdrives", test.cfg.e.overdrives.endpoint, exp.overdrives, got.Overdrives)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expMiscAbility struct {
	testGeneral
	expNameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	topmenu            *int32
	submenu            *int32
	openSubmenu        *int32
	battleInteractions []expBattleInteraction
}

func (e expMiscAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMiscAbilities(test test, exp expMiscAbility, got MiscAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResourcePtrs(test, "submenu", test.cfg.e.submenus.endpoint, exp.submenu, got.Submenu)
	compIdApiResourcePtrs(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expTriggerCommand struct {
	testGeneral
	expNameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	topmenu            *int32
	usedBy             []int32
	monsterFormations  []int32
	battleInteractions []expBattleInteraction
}

func (e expTriggerCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTriggerCommands(test test, exp expTriggerCommand, got TriggerCommand) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	checkResIDsInSlice(test, "used by", test.cfg.e.characterClasses.endpoint, exp.usedBy, got.UsedBy)
	checkResIDsInSlice(test, "monster-formations", test.cfg.e.monsterFormations.endpoint, exp.monsterFormations, got.MonsterFormations)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}
