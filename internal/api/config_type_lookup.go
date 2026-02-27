package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// TypeLookup holds all the enum types for the application that are either used as endpoint or query param
type TypeLookup struct {
	AreaConnectionType          EnumType[database.AreaConnectionType, any]
	ArenaCreationCategory       EnumType[database.MaCreationCategory, database.NullMaCreationCategory]
	Arranger                    EnumType[database.Arranger, database.NullArranger]
	BlitzballTournamentCategory EnumType[database.BlitzballTournamentCategory, any]
	CharacterClassCategory		EnumType[database.CharacterClassCategory, any]
	Composer                    EnumType[database.Composer, database.NullComposer]
	CreationArea                EnumType[database.MaCreationArea, database.NullMaCreationArea]
	CTBIconType                 EnumType[database.CtbIconType, any]
	ItemCategory				EnumType[database.ItemCategory, any]
	LootType                    EnumType[database.LootType, any]
	MonsterFormationCategory    EnumType[database.MonsterFormationCategory, any]
	MonsterSpecies              EnumType[database.MonsterSpecies, any]
	MonsterCategory             EnumType[database.MonsterCategory, any]
	OverdriveModeType           EnumType[database.OverdriveModeType, any]
	PlayerAbilityCategory		EnumType[database.PlayerAbilityCategory, any]
	ShopCategory                EnumType[database.ShopCategory, any]
	ShopType                    EnumType[database.ShopType, any]
	TreasureType                EnumType[database.TreasureType, any]

	AccSourceType				EnumType[database.AccSourceType, any]
	AttackType					EnumType[database.AttackType, any]
	BreakDmgLimitType			EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]
	CalculationType				EnumType[database.CalculationType, any]
	CriticalType				EnumType[database.CriticalType, database.NullCriticalType]
	CtbAttackType				EnumType[database.CtbAttackType, any]
	DamageFormula				EnumType[database.DamageFormula, any]
	DamageType					EnumType[database.DamageType, any]
	DelayType					EnumType[database.DelayType, any]
	DurationType				EnumType[database.DurationType, any]
	TargetType					EnumType[database.TargetType, database.NullTargetType]
}

func (cfg *Config) TypeLookupInit() {
	cfg.t = &TypeLookup{}

	cfg.t.initAreaConnectionType()
	cfg.t.initArenaCreationCategory()
	cfg.t.initArranger()
	cfg.t.initBlitzballTournamentCategory()
	cfg.t.initCharacterClassCategory()
	cfg.t.initComposer()
	cfg.t.initCTBIconType()
	cfg.t.initCreationArea()
	cfg.t.initItemCategory()
	cfg.t.initLootType()
	cfg.t.initMonsterFormationCategory()
	cfg.t.initMonsterSpecies()
	cfg.t.initMonsterCategory()
	cfg.t.initOverdriveModeType()
	cfg.t.initPlayerAbilityCategory()
	cfg.t.initShopCategory()
	cfg.t.initShopType()
	cfg.t.initTreasureType()
	
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

// replace Typed logic and lookup with this struct
type EnumType[T, N any] struct {
	name         string
	isEndpoint   bool
	lookup       map[string]TypedAPIResource
	convFunc     func(string) T
	nullConvFunc func(*string) N
}

func newEnumType[T, N any](name string, isEndpoint bool, typeSlice []TypedAPIResource, convFunc func(string) T, nullConvFunc func(*string) N) EnumType[T, N] {
	return EnumType[T, N]{
		name:         name,
		isEndpoint:   isEndpoint,
		lookup:       typeSliceToMap(typeSlice),
		convFunc:     convFunc,
		nullConvFunc: nullConvFunc,
	}
}

func (t *TypeLookup) initAreaConnectionType() {
	typeSlice := []TypedAPIResource{
		{
			Name:        string(database.AreaConnectionTypeBothDirections),
			Description: "The edges of two areas are directly connected with each other, and you can freely zone between those areas.",
		},
		{
			Name:        string(database.AreaConnectionTypeOneDirection),
			Description: "The edges of two areas are directly connected with each other, but you can only zone from area A to area B, and not vice versa.",
		},
		{
			Name:        string(database.AreaConnectionTypeWarp),
			Description: "A connection of two areas that doesn't require crossing their edges. Most of the time, their edges are not directly connected, but you can reach area B through other means. That might be due to a teleporter (like in Gagazet), or due to a story-based warp.",
		},
	}

	t.AreaConnectionType = newEnumType[database.AreaConnectionType, any]("area connection type", false, typeSlice, func(s string) database.AreaConnectionType {
		return database.AreaConnectionType(s)
	}, nil)
}

func (t *TypeLookup) initArenaCreationCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.MaCreationCategoryArea),
		},
		{
			Name: string(database.MaCreationCategorySpecies),
		},
		{
			Name: string(database.MaCreationCategoryOriginal),
		},
	}

	t.ArenaCreationCategory = newEnumType("arena creation category", false, typeSlice, func(s string) database.MaCreationCategory {
		return database.MaCreationCategory(s)
	}, h.NullMaCreationCategory)
}

func (t *TypeLookup) initArranger() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.ArrangerNobuouematsu),
		},
		{
			Name: string(database.ArrangerJunyanakano),
		},
		{
			Name: string(database.ArrangerMasashihamauzu),
		},
		{
			Name: string(database.ArrangerShirohamaguchi),
		},
	}

	t.Arranger = newEnumType("arranger", false, typeSlice, func(s string) database.Arranger {
		return database.Arranger(s)
	}, h.NullArranger)
}

func (t *TypeLookup) initBlitzballTournamentCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.BlitzballTournamentCategoryLeague),
		},
		{
			Name: string(database.BlitzballTournamentCategoryTournament),
		},
	}

	t.BlitzballTournamentCategory = newEnumType[database.BlitzballTournamentCategory, any]("blitzball tournament category", false, typeSlice, func(s string) database.BlitzballTournamentCategory {
		return database.BlitzballTournamentCategory(s)
	}, nil)
}

func (t *TypeLookup) initCharacterClassCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.CharacterClassCategoryGroup),
		},
		{
			Name: string(database.CharacterClassCategoryCharacter),
		},
		{
			Name: string(database.CharacterClassCategoryAeon),
		},
	}

	t.CharacterClassCategory = newEnumType[database.CharacterClassCategory, any]("character class category", false, typeSlice, func(s string) database.CharacterClassCategory {
		return database.CharacterClassCategory(s)
	}, nil)
}

func (t *TypeLookup) initComposer() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.ComposerNobuouematsu),
		},
		{
			Name: string(database.ComposerJunyanakano),
		},
		{
			Name: string(database.ComposerMasashihamauzu),
		},
	}

	t.Composer = newEnumType("composer", false, typeSlice, func(s string) database.Composer {
		return database.Composer(s)
	}, h.NullComposer)
}

func (t *TypeLookup) initCTBIconType() {
	typeSlice := []TypedAPIResource{
		{
			Name:        string(database.CtbIconTypeMonster),
			Description: "Used for regular monsters",
		},
		{
			Name:        string(database.CtbIconTypeBoss),
			Description: "Used for bosses",
		},
		{
			Name:        string(database.CtbIconTypeBossNumbered),
			Description: "Used for multiple bosses, or subparts of a boss",
		},
		{
			Name:        string(database.CtbIconTypeSummon),
			Description: "Used for aeons, except dark aeons",
		},
		{
			Name:        string(database.CtbIconTypeCid),
			Description: "Used for Cid during the Evrae fight",
		},
	}

	t.CTBIconType = newEnumType[database.CtbIconType, any]("ctb icon type", false, typeSlice, func(s string) database.CtbIconType {
		return database.CtbIconType(s)
	}, nil)
}

func (t *TypeLookup) initCreationArea() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.MaCreationAreaBesaid),
		},
		{
			Name: string(database.MaCreationAreaKilika),
		},
		{
			Name: string(database.MaCreationAreaMiihenHighroad),
		},
		{
			Name: string(database.MaCreationAreaMushroomRockRoad),
		},
		{
			Name: string(database.MaCreationAreaDjose),
		},
		{
			Name: string(database.MaCreationAreaThunderPlains),
		},
		{
			Name: string(database.MaCreationAreaMacalania),
		},
		{
			Name: string(database.MaCreationAreaBikanel),
		},
		{
			Name: string(database.MaCreationAreaCalmLands),
		},
		{
			Name: string(database.MaCreationAreaCavernOfTheStolenFayth),
		},
		{
			Name: string(database.MaCreationAreaMountGagazet),
		},
		{
			Name: string(database.MaCreationAreaSin),
		},
		{
			Name: string(database.MaCreationAreaOmegaRuins),
		},
	}

	t.CreationArea = newEnumType("creation area", false, typeSlice, func(s string) database.MaCreationArea {
		return database.MaCreationArea(s)
	}, h.NullMaCreationArea)
}

func (t *TypeLookup) initItemCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.ItemCategoryHealing),
			Description: 	"Items that are used for recovery of HP and MP, or for curing negative status ailments.",
		},
		{
			Name: 			string(database.ItemCategoryOffensive),
			Description: 	"Items that deal damage to other enemies or inflict status ailments.",
		},
		{
			Name: 			string(database.ItemCategorySupport),
			Description: 	"Items that grant positive statusses or other supportive effects.",
		},
		{
			Name: 			string(database.ItemCategorySphere),
			Description: 	"Items that can only be used within the sphere grid.",
		},
		{
			Name: 			string(database.ItemCategoryDistiller),
			Description: 	"Items that cause enemies to drop spheres.",
		},
		{
			Name: 			string(database.ItemCategoryOther),
			Description: 	"Uncategorized items, that are mostly used for mixes.",
		},
	}

	t.ItemCategory = newEnumType[database.ItemCategory, any]("item category", true, typeSlice, func(s string) database.ItemCategory {
		return database.ItemCategory(s)
	}, nil)
}

func (t *TypeLookup) initLootType() {

	typeSlice := []TypedAPIResource{
		{
			Name: string(database.LootTypeItem),
		},
		{
			Name: string(database.LootTypeEquipment),
		},
		{
			Name: string(database.LootTypeGil),
		},
	}

	t.LootType = newEnumType[database.LootType, any]("loot type", true, typeSlice, func(s string) database.LootType {
		return database.LootType(s)
	}, nil)
}

func (t *TypeLookup) initMonsterFormationCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name:        string(database.MonsterFormationCategoryRandomEncounter),
			Description: "A typical random encounter which can effectively be triggered indefinitely.",
		},
		{
			Name:        string(database.MonsterFormationCategoryBossFight),
			Description: "A boss encounter. Can only be triggered once, usually during the events of the story.",
		},
		{
			Name:        string(database.MonsterFormationCategoryStoryFight),
			Description: "A story-based, non-boss-encounter. Is triggered during the events of the story. Usually once, unless stated otherwise.",
		},
		{
			Name:        string(database.MonsterFormationCategoryStaticEncounter),
			Description: "A non-boss-encounter that is triggered by interacting with it. This only applies to the Sandragoras in Bikanel.",
		},
		{
			Name:        string(database.MonsterFormationCategoryTutorial),
			Description: "A unique tutorial fight. Can only be triggered once.",
		},
		{
			Name:        string(database.MonsterFormationCategoryOnDemandFight),
			Description: "An encounter that can be triggered indefinitely via the Monster Arena.",
		},
	}

	t.MonsterFormationCategory = newEnumType[database.MonsterFormationCategory, any]("monster formation category", true, typeSlice, func(s string) database.MonsterFormationCategory {
		return database.MonsterFormationCategory(s)
	}, nil)
}

func (t *TypeLookup) initMonsterSpecies() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.MonsterSpeciesAdamantoise),
		},
		{
			Name: string(database.MonsterSpeciesAeon),
		},
		{
			Name: string(database.MonsterSpeciesArmor),
		},
		{
			Name: string(database.MonsterSpeciesBasilisk),
		},
		{
			Name: string(database.MonsterSpeciesBlade),
		},
		{
			Name: string(database.MonsterSpeciesBehemoth),
		},
		{
			Name: string(database.MonsterSpeciesBird),
		},
		{
			Name: string(database.MonsterSpeciesBomb),
		},
		{
			Name: string(database.MonsterSpeciesCactuar),
		},
		{
			Name: string(database.MonsterSpeciesCephalopod),
		},
		{
			Name: string(database.MonsterSpeciesChest),
		},
		{
			Name: string(database.MonsterSpeciesChimera),
		},
		{
			Name: string(database.MonsterSpeciesCoeurl),
		},
		{
			Name: string(database.MonsterSpeciesDefender),
		},
		{
			Name: string(database.MonsterSpeciesDinofish),
		},
		{
			Name: string(database.MonsterSpeciesDoomstone),
		},
		{
			Name: string(database.MonsterSpeciesDrake),
		},
		{
			Name: string(database.MonsterSpeciesEater),
		},
		{
			Name: string(database.MonsterSpeciesElemental),
		},
		{
			Name: string(database.MonsterSpeciesEvilEye),
		},
		{
			Name: string(database.MonsterSpeciesFlan),
		},
		{
			Name: string(database.MonsterSpeciesFungus),
		},
		{
			Name: string(database.MonsterSpeciesGel),
		},
		{
			Name: string(database.MonsterSpeciesGeo),
		},
		{
			Name: string(database.MonsterSpeciesHaizhe),
		},
		{
			Name: string(database.MonsterSpeciesHelm),
		},
		{
			Name: string(database.MonsterSpeciesHermit),
		},
		{
			Name: string(database.MonsterSpeciesHumanoid),
		},
		{
			Name: string(database.MonsterSpeciesImp),
		},
		{
			Name: string(database.MonsterSpeciesIronGiant),
		},
		{
			Name: string(database.MonsterSpeciesLarva),
		},
		{
			Name: string(database.MonsterSpeciesLupine),
		},
		{
			Name: string(database.MonsterSpeciesMachina),
		},
		{
			Name: string(database.MonsterSpeciesMalboro),
		},
		{
			Name: string(database.MonsterSpeciesMech),
		},
		{
			Name: string(database.MonsterSpeciesMimic),
		},
		{
			Name: string(database.MonsterSpeciesOchu),
		},
		{
			Name: string(database.MonsterSpeciesOgre),
		},
		{
			Name: string(database.MonsterSpeciesPhantom),
		},
		{
			Name: string(database.MonsterSpeciesPiranha),
		},
		{
			Name: string(database.MonsterSpeciesPlant),
		},
		{
			Name: string(database.MonsterSpeciesReptile),
		},
		{
			Name: string(database.MonsterSpeciesRoc),
		},
		{
			Name: string(database.MonsterSpeciesRuminant),
		},
		{
			Name: string(database.MonsterSpeciesSacredBeast),
		},
		{
			Name: string(database.MonsterSpeciesSahagin),
		},
		{
			Name: string(database.MonsterSpeciesSin),
		},
		{
			Name: string(database.MonsterSpeciesSinspawn),
		},
		{
			Name: string(database.MonsterSpeciesSpellspinner),
		},
		{
			Name: string(database.MonsterSpeciesSpiritBeast),
		},
		{
			Name: string(database.MonsterSpeciesTonberry),
		},
		{
			Name: string(database.MonsterSpeciesUnspecified),
		},
		{
			Name: string(database.MonsterSpeciesWasp),
		},
		{
			Name: string(database.MonsterSpeciesWeapon),
		},
		{
			Name: string(database.MonsterSpeciesWorm),
		},
		{
			Name: string(database.MonsterSpeciesWyrm),
		},
	}

	t.MonsterSpecies = newEnumType[database.MonsterSpecies, any]("monster species", true, typeSlice, func(s string) database.MonsterSpecies {
		return database.MonsterSpecies(s)
	}, nil)
}

func (t *TypeLookup) initMonsterCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.MonsterCategoryMonster),
		},
		{
			Name: string(database.MonsterCategoryBoss),
		},
		{
			Name: string(database.MonsterCategorySummon),
		},
	}

	t.MonsterCategory = newEnumType[database.MonsterCategory, any]("monster category", true, typeSlice, func(s string) database.MonsterCategory {
		return database.MonsterCategory(s)
	}, nil)
}

func (t *TypeLookup) initOverdriveModeType() {
	typeSlice := []TypedAPIResource{
		{
			Name:        string(database.OverdriveModeTypeFormula),
			Description: "The fill-amount of the overdrive gauge is determined by a formula.",
		},
		{
			Name:        string(database.OverdriveModeTypePerAction),
			Description: "The overdrive gauge fills by a fixed amount every time the specified action is performed.",
		},
	}

	t.OverdriveModeType = newEnumType[database.OverdriveModeType, any]("overdrive mode type", false, typeSlice, func(s string) database.OverdriveModeType {
		return database.OverdriveModeType(s)
	}, nil)
}

func (t *TypeLookup) initPlayerAbilityCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.PlayerAbilityCategorySkill),
		},
		{
			Name: string(database.PlayerAbilityCategorySpecial),
		},
		{
			Name: string(database.PlayerAbilityCategoryWhiteMagic),
		},
		{
			Name: string(database.PlayerAbilityCategoryBlackMagic),
		},
		{
			Name: string(database.PlayerAbilityCategoryAeon),
		},
	}

	t.PlayerAbilityCategory = newEnumType[database.PlayerAbilityCategory, any]("player ability category", true, typeSlice, func(s string) database.PlayerAbilityCategory {
		return database.PlayerAbilityCategory(s)
	}, nil)
}

func (t *TypeLookup) initShopCategory() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.ShopCategoryStandard),
		},
		{
			Name: string(database.ShopCategoryOaka),
		},
		{
			Name: string(database.ShopCategoryTravelAgency),
		},
		{
			Name: string(database.ShopCategoryWantz),
		},
	}

	t.ShopCategory = newEnumType[database.ShopCategory, any]("shop category", true, typeSlice, func(s string) database.ShopCategory {
		return database.ShopCategory(s)
	}, nil)
}

func (t *TypeLookup) initShopType() {
	typeSlice := []TypedAPIResource{
		{
			Name: string(database.ShopTypePreAirship),
		},
		{
			Name: string(database.ShopTypePostAirship),
		},
	}

	t.ShopType = newEnumType[database.ShopType, any]("shop type", false, typeSlice, func(s string) database.ShopType {
		return database.ShopType(s)
	}, nil)
}

func (t *TypeLookup) initTreasureType() {
	typeSlice := []TypedAPIResource{
		{
			Name:        string(database.TreasureTypeChest),
			Description: "The treasure is found in a chest.",
		},
		{
			Name:        string(database.TreasureTypeGift),
			Description: "The treasure is a gift from an NPC.",
		},
		{
			Name:        string(database.TreasureTypeObject),
			Description: "The treasure is found by interacting with an in-game object. Most of the time, the treasure is the object itself (Jecht Spheres, Al Bhed Primers), other times it's not.",
		},
	}

	t.TreasureType = newEnumType[database.TreasureType, any]("treasure type", false, typeSlice, func(s string) database.TreasureType {
		return database.TreasureType(s)
	}, nil)
}

func (t *TypeLookup) initAccSourceType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.AccSourceTypeAccuracy),
			Description:	"The accuracy of the ability is calculated via the user's accuracy stat.",
		},
		{
			Name: 			string(database.AccSourceTypeRate),
			Description:	"The ability has its own accuracy.",
		},
	}

	t.AccSourceType = newEnumType[database.AccSourceType, any]("accuracy source type", false, typeSlice, func(s string) database.AccSourceType {
		return database.AccSourceType(s)
	}, nil)
}

func (t *TypeLookup) initAttackType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.AttackTypeAttack),
		},
		{
			Name: 			string(database.AttackTypeHeal),
		},
		{
			Name: 			string(database.AttackTypeAbsorb),
		},
	}

	t.AttackType = newEnumType[database.AttackType, any]("attack type", true, typeSlice, func(s string) database.AttackType {
		return database.AttackType(s)
	}, nil)
}

func (t *TypeLookup) initBreakDmgLimitType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.BreakDmgLmtTypeAlways),
			Description:	"The ability always breaks the damage limit.",
		},
		{
			Name: 			string(database.BreakDmgLmtTypeAutoAbility),
			Description:	"The ability can only break the damage limit, if the user has the auto-ability 'Break Damage Limit' equipped.",
		},
	}

	t.BreakDmgLimitType = newEnumType("break damage limit type", false, typeSlice, func(s string) database.BreakDmgLmtType {
		return database.BreakDmgLmtType(s)
	}, h.NullBreakDmgLmtType)
}

func (t *TypeLookup) initCalculationType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.CalculationTypeAddedPercentage),
			Description:	"The given value is added (or subtracted, if negative) to a final percentage-based factor which is applied at the end of the calculation. Example: If the value is 3 (like with Auto-Ability 'Strength +3%'), then the result of the calculation will be multiplied by 1.03.",
		},
		{
			Name: 			string(database.CalculationTypeAddedValue),
			Description:	"The given value is added directly to the destination. This type is either used directly on stats or on factors within the calculation and is most prominently seen on abilities like 'Cheer' and its equivalents.",
		},
		{
			Name: 			string(database.CalculationTypeMultiply),
			Description:	"The result of the calculation will be multiplied by the given value. Values with calculation type 'multiply' can stack on the same destination. Example: If Rikku uses 'Hot Spurs' (overdrive-charge x1.5) and then 'Eccentrick' (overdrive-charge x2), the gauge will charge 3 times as fast.",
		},
		{
			Name: 			string(database.CalculationTypeMultiplyHighest),
			Description:	"The result of the calculation will be multiplied by the given value. If more than one modification with calculation type 'multiply-highest' reach the same destination, only the highest factor is applied. Example: Auto-Abilities 'Double AP' and 'Triple AP' both use 'multiply-highest'. Factor 3 of 'Triple AP' will override factor 2 of 'Double AP', since it's higher.",
		},
		{
			Name: 			string(database.CalculationTypeSetValue),
			Description:	"The destination becomes the given value. Example: Auto-Ability 'One MP Cost' sets the MP cost every spell to 1.",
		},
		{
			Name: 			string(database.CalculationTypeSetMinValue),
			Description:	"The destination can't be lower than the given value at the end of all calculations. Example: 'Trio of 9999' sets the minimum damage a character can deal to 9999.",
		},
	}

	t.CalculationType = newEnumType[database.CalculationType, any]("calculation type", false, typeSlice, func(s string) database.CalculationType {
		return database.CalculationType(s)
	}, nil)
}

func (t *TypeLookup) initCriticalType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.CriticalTypeCrit),
			Description:	"The ability uses the normal critical hit formula.",
		},
		{
			Name: 			string(database.CriticalTypeCritweapon),
			Description:	"The critical plus values of the user's equipment are added toward the critical hit chance.",
		},
		{
			Name: 			string(database.CriticalTypeCritability),
			Description:	"The critical plus value of the used ability is added toward the critical hit chance.",
		},
	}

	t.CriticalType = newEnumType("critical type", false, typeSlice, func(s string) database.CriticalType {
		return database.CriticalType(s)
	}, h.NullCriticalType)
}

func (t *TypeLookup) initCtbAttackType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.CtbAttackTypeAttack),
		},
		{
			Name: 			string(database.CtbAttackTypeHeal),
		},
	}

	t.CtbAttackType = newEnumType[database.CtbAttackType, any]("ctb attack type", false, typeSlice, func(s string) database.CtbAttackType {
		return database.CtbAttackType(s)
	}, nil)
}

func (t *TypeLookup) initDamageFormula() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.DamageFormulaStrVsDef),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaStrIgnDef),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaMagVsMdf),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaMagIgnMdf),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaPercentageCurrent),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaPercentageMax),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaHealing),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecialNoVar),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecialVar),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecialMagic),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecialGil),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecialKills),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSpecial9999),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaFixed9999),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaUserMaxHp),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSwallowedA),
			Description: 	"",
		},
		{
			Name: 			string(database.DamageFormulaSwallowedB),
			Description: 	"",
		},
	}

	t.DamageFormula = newEnumType[database.DamageFormula, any]("damage formula", true, typeSlice, func(s string) database.DamageFormula {
		return database.DamageFormula(s)
	}, nil)
}

func (t *TypeLookup) initDamageType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.DamageTypePhysical),
			Description: 	"The damage can be reduced by 'Protect', 'Defend' 'Power Break', 'Sentinel', 'Shield', and 'Cheer', as well as 'Defense +X%' Auto-Abilities.",
		},
		{
			Name: 			string(database.DamageTypeMagical),
			Description: 	"The damage can be reduced by 'Shell', 'Magic Break', 'Shield', and 'Focus', as well as 'Magic Def +X%' Auto-Abilities. It can be increased by 'Magic +X%' Auto-Abilities.",
		},
		{
			Name: 			string(database.DamageTypeSpecial),
			Description: 	"The damage can only be reduced by 'Shield'.",
		},
	}

	t.DamageType = newEnumType[database.DamageType, any]("damage type", true, typeSlice, func(s string) database.DamageType {
		return database.DamageType(s)
	}, nil)
}

func (t *TypeLookup) initDelayType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.DelayTypeCtbBased),
			Description: 	"Delay is based on current ticks. CTB damage/heal is only applied, if 'Slow'/'Haste' is succcessful or if the status was successfully removed.",
		},
		{
			Name: 			string(database.DelayTypeTickSpeedBased),
			Description: 	"Delay is based on tick speed. CTB damage is applied via an attack. Example: 'Delay Attack'.",
		},
	}

	t.DelayType = newEnumType[database.DelayType, any]("delay type", false, typeSlice, func(s string) database.DelayType {
		return database.DelayType(s)
	}, nil)
}

func (t *TypeLookup) initDurationType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.DurationTypeTurns),
			Description: 	"The status condition wears off after a set amount of turns.",
		},
		{
			Name: 			string(database.DurationTypeInflictorNextTurn),
			Description: 	"The status condition wears off on the inflictor's next turn. This is only used for 'Threaten'.",
		},
		{
			Name: 			string(database.DurationTypeBlocks),
			Description: 	"The status condition is present as long as it has blocks left. Used only for 'Nul-' status conditions.",
		},
		{
			Name: 			string(database.DurationTypeEndless),
			Description: 	"The status condition won't wear off. It is present until it is removed.",
		},
		{
			Name: 			string(database.DurationTypeInstant),
			Description: 	"The status condition wears off instantly. Most commonly seen on 'Death' and 'Life', but there are exceptions like Sinspawn Gui and Ultima Buster gaining 'Defend' while blocking, or Penance's Arms gaining 'Haste' while taking an action.",
		},
		{
			Name: 			string(database.DurationTypeAuto),
			Description: 	"The status condition is present forever and can't be removed. Only used on Biran Ronso's 'Mighty Guard'.",
		},
	}

	t.DurationType = newEnumType[database.DurationType, any]("duration type", false, typeSlice, func(s string) database.DurationType {
		return database.DurationType(s)
	}, nil)
}

func (t *TypeLookup) initTargetType() {
	typeSlice := []TypedAPIResource{
		{
			Name: 			string(database.TargetTypeSelf),
			Description:	"The action targets its user.",
		},
		{
			Name: 			string(database.TargetTypeSingleCharacter),
			Description:	"The action targets one unit of the user's party.",
		},
		{
			Name: 			string(database.TargetTypeSingleEnemy),
			Description:	"The action targets one unit of the user's opposing party.",
		},
		{
			Name: 			string(database.TargetTypeSingleTarget),
			Description:	"The action targets the selected unit.",
		},
		{
			Name: 			string(database.TargetTypeRandomCharacter),
			Description:	"The action targets a random unit of the user's party.",
		},
		{
			Name: 			string(database.TargetTypeRandomEnemy),
			Description:	"The action targets a random unit of the user's opposing party.",
		},
		{
			Name: 			string(database.TargetTypeAllCharacters),
			Description:	"The action targets all units of the user's party.",
		},
		{
			Name: 			string(database.TargetTypeAllEnemies),
			Description:	"The action targets all units of the user's opposing party.",
		},
		{
			Name: 			string(database.TargetTypeTargetParty),
			Description:	"The action targets all units of the selected party.",
		},
		{
			Name: 			string(database.TargetTypeNTargets),
			Description:	"The action targets N amount of units (N is stated via the ability's hit_amount). The action can also target KO'd characters and inanimate objects. Only Seymour Natus' multi-spells and Spectral Keeper's counter attack, as well as its glyph mine activation use this target type.",
		},
		{
			Name: 			string(database.TargetTypeEveryone),
			Description:	"The action targets every unit on the field.",
		},
	}

	t.TargetType = newEnumType("target type", false, typeSlice, func(s string) database.TargetType {
		return database.TargetType(s)
	}, h.NullTargetType)
}