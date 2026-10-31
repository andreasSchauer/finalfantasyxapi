package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

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

func compareAbilities(test test, exp expAbility, got api.Ability) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "ability type", string(exp.abilityType), string(got.Type))
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)

	compPathAPIResource(test, "typed ability", exp.typedAbility, got.TypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
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

func compareEnemyAbilities(test test, exp expEnemyAbility, got api.EnemyAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
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

func compareItemAbilities(test test, exp expItemAbility, got api.ItemAbility) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	compIDAPIResource(test, "item", test.cfg.E.Items.Endpoint, exp.item, got.Item)
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

func comparePlayerAbilities(test test, exp expPlayerAbility, got api.PlayerAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	compIDAPIResourcePtrs(test, "topmenu", test.cfg.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	compIDAPIResourcePtrs(test, "submenu", test.cfg.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	compIDAPIResourcePtrs(test, "open submenu", test.cfg.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	compIDAPIResourcePtrs(test, "std sg char", test.cfg.E.Characters.Endpoint, exp.stdChar, got.StandardGridCharacter)
	compIDAPIResourcePtrs(test, "exp sg char", test.cfg.E.Characters.Endpoint, exp.expChar, got.ExpertGridCharacter)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
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

func compareOverdriveAbilities(test test, exp expOverdriveAbility, got api.OverdriveAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	compIDAPIResourcePtrs(test, "overdrive command", test.cfg.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "overdrives", test.cfg.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
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

func compareMiscAbilities(test test, exp expMiscAbility, got api.MiscAbility) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	compIDAPIResourcePtrs(test, "topmenu", test.cfg.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	compIDAPIResourcePtrs(test, "submenu", test.cfg.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	compIDAPIResourcePtrs(test, "open submenu", test.cfg.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
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

func compareTriggerCommands(test test, exp expTriggerCommand, got api.TriggerCommand) {
	test.t.Helper()
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIDAPIResource(test, "untyped ability", test.cfg.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	compIDAPIResourcePtrs(test, "topmenu", test.cfg.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	checkResIDsInSlice(test, "used by", test.cfg.E.CharacterClasses.Endpoint, exp.usedBy, got.UsedBy)
	checkResIDsInSlice(test, "monster-formations", test.cfg.E.MonsterFormations.Endpoint, exp.monsterFormations, got.MonsterFormations)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}
