// Package enums bla bla
package enums

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

var T Enums

// Enums holds all the enum types for the application that are either used as endpoint or query param
type Enums struct {
	Lookup                      map[string]EnumResponse
	AbilityType                 EnumType[database.AbilityType, any]
	UnitType                    EnumType[database.UnitType, any]
	ItemType                    EnumType[database.ItemType, any]
	QuestType                   EnumType[database.QuestType, any]
	AaActivationCondition       EnumType[database.AaActivationCondition, any]
	ElementalAffinity           EnumType[database.ElementalAffinity, any]
	AreaConnectionType          EnumType[database.AreaConnectionType, any]
	ArenaCreationCategory       EnumType[database.MaCreationCategory, database.NullMaCreationCategory]
	ArmorType                   EnumType[database.ArmorType, any]
	Arranger                    EnumType[database.Arranger, database.NullArranger]
	AutoAbilityCategory         EnumType[database.AutoAbilityCategory, any]
	AvailabilityType            EnumType[database.AvailabilityType, any]
	BgReplacementType           EnumType[database.BgReplacementType, database.NullBgReplacementType]
	BlitzballPositionSlot       EnumType[database.BlitzballPositionSlot, any]
	BlitzballTournamentCategory EnumType[database.BlitzballTournamentCategory, any]
	CelestialFormula            EnumType[database.CelestialFormula, any]
	CharacterClassCategory      EnumType[database.CharacterClassCategory, any]
	CounterType                 EnumType[database.CounterType, database.NullCounterType]
	Composer                    EnumType[database.Composer, database.NullComposer]
	CreationArea                EnumType[database.MaCreationArea, database.NullMaCreationArea]
	CreationsUnlockedCategory   EnumType[database.CreationsUnlockedCategory, database.NullCreationsUnlockedCategory]
	CTBIconType                 EnumType[database.CtbIconType, any]
	EquipClass                  EnumType[database.EquipClass, any]
	EquipType                   EnumType[database.EquipType, any]
	ItemCategory                EnumType[database.ItemCategory, any]
	ItemUsability               EnumType[database.ItemUsability, any]
	KeyItemCategory             EnumType[database.KeyItemCategory, any]
	LootType                    EnumType[database.LootType, any]
	MixCategory                 EnumType[database.MixCategory, any]
	ModifierCategory            EnumType[database.ModifierCategory, any]
	MonsterCategory             EnumType[database.MonsterCategory, any]
	MonsterFormationCategory    EnumType[database.MonsterFormationCategory, any]
	MonsterSpecies              EnumType[database.MonsterSpecies, any]
	MusicUseCase                EnumType[database.MusicUseCase, database.NullMusicUseCase]
	NodePosition                EnumType[database.NodePosition, any]
	NodeState                   EnumType[database.NodeState, database.NullNodeState]
	NodeType                    EnumType[database.NodeType, any]
	NullifyArmored              EnumType[database.NullifyArmored, database.NullNullifyArmored]
	OverdriveModeType           EnumType[database.OverdriveModeType, any]
	PlayerAbilityCategory       EnumType[database.PlayerAbilityCategory, any]
	ShopCategory                EnumType[database.ShopCategory, any]
	SphereColor                 EnumType[database.SphereColor, any]
	SphereEffect                EnumType[database.SphereEffect, any]
	SphereGridType              EnumType[database.SphereGridType, any]
	StatusConditionCategory     EnumType[database.StatusConditionCategory, any]
	TreasureType                EnumType[database.TreasureType, any]
	WeaponType                  EnumType[database.WeaponType, any]
	AccSourceType               EnumType[database.AccSourceType, any]
	AttackType                  EnumType[database.AttackType, any]
	BreakDmgLimitType           EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]
	CalculationType             EnumType[database.CalculationType, any]
	CriticalType                EnumType[database.CriticalType, database.NullCriticalType]
	DamageFormula               EnumType[database.DamageFormula, any]
	DamageType                  EnumType[database.DamageType, any]
	DurationType                EnumType[database.DurationType, any]
	SpecialActionType           EnumType[database.SpecialActionType, database.NullSpecialActionType]
	TargetType                  EnumType[database.TargetType, database.NullTargetType]
	BattleStart                 EnumType[database.BattleStart, any]
	DelayStrength               EnumType[database.DelayStrength, any]
	HasteStatus                 EnumType[database.HasteStatus, any]
	TranslationDirection        EnumType[database.TranslationDirection, any]
	TurnOrderRNG                EnumType[database.TurnOrderRng, any]

	// internal?
	// AlterationType              EnumType[database.AlterationType, any]
	// ShopType                    EnumType[database.ShopType, database.NullShopType]
}

func (T *Enums) EnumLookupInit() {
	T.Lookup = make(map[string]EnumResponse)

	T.initAbilityType()
	T.initUnitType()
	T.initItemType()
	T.initQuestType()

	T.initAaActivationCondition()
	//t.initAlterationType()
	T.initAreaConnectionType()
	T.initArenaCreationCategory()
	T.initArmorType()
	T.initArranger()
	T.initAutoAbilityCategory()
	T.initAvailabilityType()
	T.initBgReplacementType()
	T.initBlitzballPositionSlot()
	T.initBlitzballTournamentCategory()
	T.initCelestialFormula()
	T.initCharacterClassCategory()
	T.initComposer()
	T.initCounterType()
	T.initCreationArea()
	T.initCreationsUnlockedCategory()
	T.initCTBIconType()
	T.initElementalAffinity()
	T.initEquipClass()
	T.initEquipType()
	T.initItemCategory()
	T.initItemUsability()
	T.initKeyItemCategory()
	T.initLootType()
	T.initMixCategory()
	T.initModifierCategory()
	T.initMonsterCategory()
	T.initMonsterFormationCategory()
	T.initMonsterSpecies()
	T.initMusicUseCase()
	T.initNodePosition()
	T.initNodeState()
	T.initNodeType()
	T.initNullifyArmored()
	T.initOverdriveModeType()
	T.initPlayerAbilityCategory()
	T.initShopCategory()
	//t.initShopType()
	T.initSpecialActionType()
	T.initSphereColor()
	T.initSphereEffect()
	T.initSphereGridType()
	T.initStatusConditionCategory()
	T.initTreasureType()
	T.initWeaponType()

	T.initAccSourceType()
	T.initAttackType()
	T.initBreakDmgLimitType()
	T.initCalculationType()
	T.initCriticalType()
	T.initDamageFormula()
	T.initDamageType()
	T.initDelayStrength()
	T.initDurationType()
	T.initTargetType()

	T.initBattleStart()
	T.initHasteStatus()
	T.initTranslationDirection()
	T.initTurnOrderRNG()
}
