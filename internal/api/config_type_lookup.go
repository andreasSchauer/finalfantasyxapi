package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// Types holds all the enum types for the application that are either used as endpoint or query param
type Types struct {
	Lookup		map[EndpointName]EnumResponse

	AbilityType EnumType[database.AbilityType, any]
	UnitType    EnumType[database.UnitType, any]
	ItemType    EnumType[database.ItemType, any]
	QuestType   EnumType[database.QuestType, any]

	AaActivationCondition       EnumType[database.AaActivationCondition, any]
	ElementalAffinity           EnumType[database.ElementalAffinity, any]
	AlterationType              EnumType[database.AlterationType, any]
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
	ShopType                    EnumType[database.ShopType, database.NullShopType]
	SpecialActionType           EnumType[database.SpecialActionType, database.NullSpecialActionType]
	SphereColor                 EnumType[database.SphereColor, any]
	SphereEffect                EnumType[database.SphereEffect, any]
	SphereGridType              EnumType[database.SphereGridType, any]
	StatusConditionCategory     EnumType[database.StatusConditionCategory, any]
	TreasureType                EnumType[database.TreasureType, any]
	WeaponType                  EnumType[database.WeaponType, any]

	AccSourceType     EnumType[database.AccSourceType, any]
	AttackType        EnumType[database.AttackType, any]
	BreakDmgLimitType EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]
	CalculationType   EnumType[database.CalculationType, any]
	CriticalType      EnumType[database.CriticalType, database.NullCriticalType]
	CtbAttackType     EnumType[database.CtbAttackType, any]
	DamageFormula     EnumType[database.DamageFormula, any]
	DamageType        EnumType[database.DamageType, any]
	DelayType         EnumType[database.DelayType, any]
	DurationType      EnumType[database.DurationType, any]
	TargetType        EnumType[database.TargetType, database.NullTargetType]
}

func (cfg *Config) TypeLookupInit() {
	cfg.t = &Types{
		Lookup: make(map[EndpointName]EnumResponse),
	}

	cfg.t.initAbilityType()
	cfg.t.initUnitType()
	cfg.t.initItemType()
	cfg.t.initQuestType()

	cfg.t.initAaActivationCondition()
	cfg.t.initAlterationType()
	cfg.t.initAreaConnectionType()
	cfg.t.initArenaCreationCategory()
	cfg.t.initArmorType()
	cfg.t.initArranger()
	cfg.t.initAutoAbilityCategory()
	cfg.t.initAvailabilityType()
	cfg.t.initBgReplacementType()
	cfg.t.initBlitzballPositionSlot()
	cfg.t.initBlitzballTournamentCategory()
	cfg.t.initCelestialFormula()
	cfg.t.initCharacterClassCategory()
	cfg.t.initComposer()
	cfg.t.initCounterType()
	cfg.t.initCreationArea()
	cfg.t.initCreationsUnlockedCategory()
	cfg.t.initCTBIconType()
	cfg.t.initElementalAffinity()
	cfg.t.initEquipClass()
	cfg.t.initEquipType()
	cfg.t.initItemCategory()
	cfg.t.initKeyItemCategory()
	cfg.t.initLootType()
	cfg.t.initMixCategory()
	cfg.t.initModifierCategory()
	cfg.t.initMonsterCategory()
	cfg.t.initMonsterFormationCategory()
	cfg.t.initMonsterSpecies()
	cfg.t.initMusicUseCase()
	cfg.t.initNodePosition()
	cfg.t.initNodeState()
	cfg.t.initNodeType()
	cfg.t.initNullifyArmored()
	cfg.t.initOverdriveModeType()
	cfg.t.initPlayerAbilityCategory()
	cfg.t.initShopCategory()
	cfg.t.initShopType()
	cfg.t.initSpecialActionType()
	cfg.t.initSphereColor()
	cfg.t.initSphereEffect()
	cfg.t.initSphereGridType()
	cfg.t.initStatusConditionCategory()
	cfg.t.initTreasureType()
	cfg.t.initWeaponType()

	cfg.t.initAccSourceType()
	cfg.t.initAttackType()
	cfg.t.initBreakDmgLimitType()
	cfg.t.initCalculationType()
	cfg.t.initCriticalType()
	cfg.t.initCtbAttackType()
	cfg.t.initDamageFormula()
	cfg.t.initDamageType()
	cfg.t.initDelayType()
	cfg.t.initDurationType()
	cfg.t.initTargetType()
}