package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Ability struct {
	testGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	abilityType        database.AbilityType
	typedAbility       string
	monsters           []int32
	battleInteractions []BattleInteraction
}

func (e Ability) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareAbilities(test test, exp Ability, got apitypes.Ability) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("ability type", string(exp.abilityType), string(got.Type))
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)

	test.compPathAPIResource("typed ability", exp.typedAbility, got.TypedAbility)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type EnemyAbility struct {
	testGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	monsters           []int32
	battleInteractions []BattleInteraction
}

func (e EnemyAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareEnemyAbilities(test test, exp EnemyAbility, got apitypes.EnemyAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type ItemAbility struct {
	testGeneral
	Unique
	rank                *int32
	appearsInHelpBar    bool
	canCopyCat          bool
	untypedAbility      int32
	item                int32
	category            database.ItemCategory
	CanUseOutsideBattle bool
	battleInteractions  []BattleInteraction
}

func (e ItemAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareItemAbilities(test test, exp ItemAbility, got apitypes.ItemAbility) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, exp.item, got.Item)
	test.compare("category", string(exp.category), got.Category)
	test.compare("can use outside battle", exp.CanUseOutsideBattle, got.CanUseOutsideBattle)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type PlayerAbility struct {
	testGeneral
	NameVer
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
	battleInteractions []BattleInteraction
}

func (e PlayerAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func comparePlayerAbilities(test test, exp PlayerAbility, got apitypes.PlayerAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compIDAPIResourcePtrs("std sg char", apilookups.E.Characters.Endpoint, exp.stdChar, got.StandardGridCharacter)
	test.compIDAPIResourcePtrs("exp sg char", apilookups.E.Characters.Endpoint, exp.expChar, got.ExpertGridCharacter)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type OverdriveAbility struct {
	testGeneral
	NameVer
	rank               *int32
	untypedAbility     int32
	overdriveCommand   *int32
	overdrives         []int32
	battleInteractions []BattleInteraction
}

func (e OverdriveAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareOverdriveAbilities(test test, exp OverdriveAbility, got apitypes.OverdriveAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, exp.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, exp.overdrives, got.Overdrives)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type MiscAbility struct {
	testGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	topmenu            *int32
	submenu            *int32
	openSubmenu        *int32
	battleInteractions []BattleInteraction
}

func (e MiscAbility) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareMiscAbilities(test test, exp MiscAbility, got apitypes.MiscAbility) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, exp.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, exp.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type TriggerCommand struct {
	testGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	topmenu            *int32
	usedBy             []int32
	monsterFormations  []int32
	battleInteractions []BattleInteraction
}

func (e TriggerCommand) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareTriggerCommands(test test, exp TriggerCommand, got apitypes.TriggerCommand) {
	test.t.Helper()
	test.compareExpNameVer(exp.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", exp.rank, got.Rank)
	test.compare("appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", exp.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, exp.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, exp.topmenu, got.Topmenu)
	test.checkResIDsInSlice("used by", apilookups.E.CharacterClasses.Endpoint, exp.usedBy, got.UsedBy)
	test.checkResIDsInSlice("monster-formations", apilookups.E.MonsterFormations.Endpoint, exp.monsterFormations, got.MonsterFormations)
	test.compTestStructSlices("battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}
