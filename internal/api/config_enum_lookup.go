package api

import (
	"cmp"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

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

func (cfg *Config) EnumLookupInit() {
	cfg.T = &Enums{
		Lookup: make(map[string]EnumResponse),
	}

	cfg.T.initAbilityType()
	cfg.T.initUnitType()
	cfg.T.initItemType()
	cfg.T.initQuestType()

	cfg.T.initAaActivationCondition()
	//cfg.t.initAlterationType()
	cfg.T.initAreaConnectionType()
	cfg.T.initArenaCreationCategory()
	cfg.T.initArmorType()
	cfg.T.initArranger()
	cfg.T.initAutoAbilityCategory()
	cfg.T.initAvailabilityType()
	cfg.T.initBgReplacementType()
	cfg.T.initBlitzballPositionSlot()
	cfg.T.initBlitzballTournamentCategory()
	cfg.T.initCelestialFormula()
	cfg.T.initCharacterClassCategory()
	cfg.T.initComposer()
	cfg.T.initCounterType()
	cfg.T.initCreationArea()
	cfg.T.initCreationsUnlockedCategory()
	cfg.T.initCTBIconType()
	cfg.T.initElementalAffinity()
	cfg.T.initEquipClass()
	cfg.T.initEquipType()
	cfg.T.initItemCategory()
	cfg.T.initItemUsability()
	cfg.T.initKeyItemCategory()
	cfg.T.initLootType()
	cfg.T.initMixCategory()
	cfg.T.initModifierCategory()
	cfg.T.initMonsterCategory()
	cfg.T.initMonsterFormationCategory()
	cfg.T.initMonsterSpecies()
	cfg.T.initMusicUseCase()
	cfg.T.initNodePosition()
	cfg.T.initNodeState()
	cfg.T.initNodeType()
	cfg.T.initNullifyArmored()
	cfg.T.initOverdriveModeType()
	cfg.T.initPlayerAbilityCategory()
	cfg.T.initShopCategory()
	//cfg.t.initShopType()
	cfg.T.initSpecialActionType()
	cfg.T.initSphereColor()
	cfg.T.initSphereEffect()
	cfg.T.initSphereGridType()
	cfg.T.initStatusConditionCategory()
	cfg.T.initTreasureType()
	cfg.T.initWeaponType()

	cfg.T.initAccSourceType()
	cfg.T.initAttackType()
	cfg.T.initBreakDmgLimitType()
	cfg.T.initCalculationType()
	cfg.T.initCriticalType()
	cfg.T.initDamageFormula()
	cfg.T.initDamageType()
	cfg.T.initDelayStrength()
	cfg.T.initDurationType()
	cfg.T.initTargetType()

	cfg.T.initBattleStart()
	cfg.T.initHasteStatus()
	cfg.T.initTranslationDirection()
	cfg.T.initTurnOrderRNG()
}

func enumLookupToSlice(lookup map[string]EnumResponse) []EnumResponse {
	enums := make([]EnumResponse, 0, len(lookup))

	for key := range lookup {
		enums = append(enums, lookup[key])
	}

	slices.SortFunc(enums, func(a, b EnumResponse) int {
		return cmp.Compare(a.Name, b.Name)
	})
	return enums
}
