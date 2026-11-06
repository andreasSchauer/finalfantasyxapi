package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
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

func compareAbilities(test test, exp expAbility, got apitypes.Ability) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("ability type", string(exp.abilityType), string(got.Type))
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)

	test.compPathAPIResource("typed ability", exp.typedAbility, got.TypedAbility)
	test.checkResIDsInSlice("monsters", api.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func compareEnemyAbilities(test test, exp expEnemyAbility, got apitypes.EnemyAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.checkResIDsInSlice("monsters", api.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func compareItemAbilities(test test, exp expItemAbility, got apitypes.ItemAbility) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResource("item", api.E.Items.Endpoint, exp.item, got.Item)
	test.compare("category", string(exp.category), got.Category)
	test.compare("can use outside battle", exp.CanUseOutsideBattle, got.CanUseOutsideBattle)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func comparePlayerAbilities(test test, exp expPlayerAbility, got apitypes.PlayerAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", api.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", api.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compIDAPIResourcePtrs("std sg char", api.E.Characters.Endpoint, exp.stdChar, got.StandardGridCharacter)
	test.compIDAPIResourcePtrs("exp sg char", api.E.Characters.Endpoint, exp.expChar, got.ExpertGridCharacter)
	test.checkResIDsInSlice("monsters", api.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func compareOverdriveAbilities(test test, exp expOverdriveAbility, got apitypes.OverdriveAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("overdrive command", api.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("overdrives", api.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func compareMiscAbilities(test test, exp expMiscAbility, got apitypes.MiscAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", api.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", api.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
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

func compareTriggerCommands(test test, exp expTriggerCommand, got apitypes.TriggerCommand) {
	test.t.Helper()
	test.compareExpNameVer(exp.expNameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", api.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", api.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.checkResIDsInSlice("used by", api.E.CharacterClasses.Endpoint, exp.usedBy, got.UsedBy)
	test.checkResIDsInSlice("monster-formations", api.E.MonsterFormations.Endpoint, exp.monsterFormations, got.MonsterFormations)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}
