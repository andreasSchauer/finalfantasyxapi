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

func (e Ability) Compare(test test, got apitypes.Ability) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("ability type", string(e.abilityType), string(got.Type))
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)

	test.compPathAPIResource("typed ability", e.typedAbility, got.TypedAbility)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e EnemyAbility) Compare(test test, got apitypes.EnemyAbility) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e ItemAbility) Compare(test test, got apitypes.ItemAbility) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, e.item, got.Item)
	test.compare("category", string(e.category), got.Category)
	test.compare("can use outside battle", e.CanUseOutsideBattle, got.CanUseOutsideBattle)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e PlayerAbility) Compare(test test, got apitypes.PlayerAbility) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, e.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.openSubmenu, got.OpenSubmenu)
	test.compIDAPIResourcePtrs("std sg char", apilookups.E.Characters.Endpoint, e.stdChar, got.StandardGridCharacter)
	test.compIDAPIResourcePtrs("exp sg char", apilookups.E.Characters.Endpoint, e.expChar, got.ExpertGridCharacter)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e OverdriveAbility) Compare(test test, got apitypes.OverdriveAbility) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.rank, got.Rank)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, e.overdriveCommand, got.OverdriveCommand)
	test.checkResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.overdrives, got.Overdrives)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e MiscAbility) Compare(test test, got apitypes.MiscAbility) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, e.submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.openSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
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

func (e TriggerCommand) Compare(test test, got apitypes.TriggerCommand) {
	test.t.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.rank, got.Rank)
	test.compare("appears in help bar", e.appearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.canCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.untypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.topmenu, got.Topmenu)
	test.checkResIDsInSlice("used by", apilookups.E.CharacterClasses.Endpoint, e.usedBy, got.UsedBy)
	test.checkResIDsInSlice("monster-formations", apilookups.E.MonsterFormations.Endpoint, e.monsterFormations, got.MonsterFormations)
	test.compTestStructSlices("battle interactions", e.battleInteractions, got.BattleInteractions)
}
