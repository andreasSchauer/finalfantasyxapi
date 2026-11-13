package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Ability struct {
	TestGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	abilityType        database.AbilityType
	typedAbility       string
	monsters           []int32
	battleInteractions []BattleInteraction
}

func (e Ability) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type EnemyAbility struct {
	TestGeneral
	NameVer
	rank               *int32
	appearsInHelpBar   bool
	canCopyCat         bool
	untypedAbility     int32
	monsters           []int32
	battleInteractions []BattleInteraction
}

func (e EnemyAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type ItemAbility struct {
	TestGeneral
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

func (e ItemAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type PlayerAbility struct {
	TestGeneral
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

func (e PlayerAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type OverdriveAbility struct {
	TestGeneral
	NameVer
	rank               *int32
	untypedAbility     int32
	overdriveCommand   *int32
	overdrives         []int32
	battleInteractions []BattleInteraction
}

func (e OverdriveAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type MiscAbility struct {
	TestGeneral
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

func (e MiscAbility) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TriggerCommand struct {
	TestGeneral
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

func (e TriggerCommand) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}
