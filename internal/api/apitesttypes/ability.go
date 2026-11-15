package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Ability struct {
	TestGeneral
	NameVer
	Rank               *int32
	AppearsInHelpBar   bool
	CanCopyCat         bool
	AbilityType        database.AbilityType
	TypedAbility       string
	Monsters           []int32
	BattleInteractions []BattleInteraction
}

func (e Ability) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Ability) Compare(test Test, got apitypes.Ability) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("ability type", string(e.AbilityType), string(got.Type))
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)

	test.compPathAPIResource("typed ability", e.TypedAbility, got.TypedAbility)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type EnemyAbility struct {
	TestGeneral
	NameVer
	Rank               *int32
	AppearsInHelpBar   bool
	CanCopyCat         bool
	UntypedAbility     int32
	Monsters           []int32
	BattleInteractions []BattleInteraction
}

func (e EnemyAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e EnemyAbility) Compare(test Test, got apitypes.EnemyAbility) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type ItemAbility struct {
	TestGeneral
	Unique
	Rank                *int32
	AppearsInHelpBar    bool
	CanCopyCat          bool
	UntypedAbility      int32
	Item                int32
	Category            database.ItemCategory
	CanUseOutsideBattle bool
	BattleInteractions  []BattleInteraction
}

func (e ItemAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e ItemAbility) Compare(test Test, got apitypes.ItemAbility) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.compIDAPIResource("item", apilookups.E.Items.Endpoint, e.Item, got.Item)
	test.compare("category", string(e.Category), got.Category)
	test.compare("can use outside battle", e.CanUseOutsideBattle, got.CanUseOutsideBattle)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type PlayerAbility struct {
	TestGeneral
	NameVer
	Rank               *int32
	AppearsInHelpBar   bool
	CanCopyCat         bool
	UntypedAbility     int32
	Topmenu            *int32
	Submenu            *int32
	OpenSubmenu        *int32
	StdChar            *int32
	ExpChar            *int32
	Monsters           []int32
	BattleInteractions []BattleInteraction
}

func (e PlayerAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e PlayerAbility) Compare(test Test, got apitypes.PlayerAbility) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, e.Submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.OpenSubmenu, got.OpenSubmenu)
	test.compIDAPIResourcePtrs("std sg char", apilookups.E.Characters.Endpoint, e.StdChar, got.StandardGridCharacter)
	test.compIDAPIResourcePtrs("exp sg char", apilookups.E.Characters.Endpoint, e.ExpChar, got.ExpertGridCharacter)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type OverdriveAbility struct {
	TestGeneral
	NameVer
	Rank               *int32
	UntypedAbility     int32
	OverdriveCommand   *int32
	Overdrives         []int32
	BattleInteractions []BattleInteraction
}

func (e OverdriveAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e OverdriveAbility) Compare(test Test, got apitypes.OverdriveAbility) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.Rank, got.Rank)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("overdrive command", apilookups.E.OverdriveCommands.Endpoint, e.OverdriveCommand, got.OverdriveCommand)
	test.CheckResIDsInSlice("overdrives", apilookups.E.Overdrives.Endpoint, e.Overdrives, got.Overdrives)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type MiscAbility struct {
	TestGeneral
	NameVer
	Rank               *int32
	AppearsInHelpBar   bool
	CanCopyCat         bool
	UntypedAbility     int32
	Topmenu            *int32
	Submenu            *int32
	OpenSubmenu        *int32
	BattleInteractions []BattleInteraction
}

func (e MiscAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e MiscAbility) Compare(test Test, got apitypes.MiscAbility) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.compIDAPIResourcePtrs("submenu", apilookups.E.Submenus.Endpoint, e.Submenu, got.Submenu)
	test.compIDAPIResourcePtrs("open submenu", apilookups.E.Submenus.Endpoint, e.OpenSubmenu, got.OpenSubmenu)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}

type TriggerCommand struct {
	TestGeneral
	NameVer
	Rank               *int32
	AppearsInHelpBar   bool
	CanCopyCat         bool
	UntypedAbility     int32
	Topmenu            *int32
	UsedBy             []int32
	MonsterFormations  []int32
	BattleInteractions []BattleInteraction
}

func (e TriggerCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e TriggerCommand) Compare(test Test, got apitypes.TriggerCommand) {
	test.T.Helper()
	test.compareExpNameVer(e.NameVer, got.ID, got.Name, got.Version)
	test.compare("rank", e.Rank, got.Rank)
	test.compare("appears in help bar", e.AppearsInHelpBar, got.AppearsInHelpBar)
	test.compare("can copycat", e.CanCopyCat, got.CanCopycat)
	test.compIDAPIResource("untyped ability", apilookups.E.Abilities.Endpoint, e.UntypedAbility, got.UntypedAbility)
	test.compIDAPIResourcePtrs("topmenu", apilookups.E.Topmenus.Endpoint, e.Topmenu, got.Topmenu)
	test.CheckResIDsInSlice("used by", apilookups.E.CharacterClasses.Endpoint, e.UsedBy, got.UsedBy)
	test.CheckResIDsInSlice("monster-formations", apilookups.E.MonsterFormations.Endpoint, e.MonsterFormations, got.MonsterFormations)
	test.compTestStructSlices("battle interactions", e.BattleInteractions, got.BattleInteractions)
}
