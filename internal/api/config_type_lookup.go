package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// TypeLookup holds all the enum types for the application
type TypeLookup struct {
	AreaConnectionType          EnumType[database.AreaConnectionType, any]
	ArenaCreationCategory       EnumType[database.MaCreationCategory, database.NullMaCreationCategory]
	Arranger                    EnumType[database.Arranger, database.NullArranger]
	BlitzballTournamentCategory EnumType[database.BlitzballTournamentCategory, any]
	Composer                    EnumType[database.Composer, database.NullComposer]
	CreationArea                EnumType[database.MaCreationArea, database.NullMaCreationArea]
	CTBIconType                 EnumType[database.CtbIconType, any]
	LootType                    EnumType[database.LootType, any]
	MonsterFormationCategory    EnumType[database.MonsterFormationCategory, any]
	MonsterSpecies              EnumType[database.MonsterSpecies, any]
	MonsterType                 EnumType[database.CtbIconType, any]
	OverdriveModeType           EnumType[database.OverdriveModeType, any]
	ShopCategory                EnumType[database.ShopCategory, any]
	ShopType                    EnumType[database.ShopType, any]
	TreasureType                EnumType[database.TreasureType, any]
}

func (cfg *Config) TypeLookupInit() {
	cfg.t = &TypeLookup{}

	cfg.t.initAreaConnectionType()
	cfg.t.initArenaCreationCategory()
	cfg.t.initArranger()
	cfg.t.initBlitzballTournamentCategory()
	cfg.t.initComposer()
	cfg.t.initCTBIconType()
	cfg.t.initCreationArea()
	cfg.t.initLootType()
	cfg.t.initMonsterFormationCategory()
	cfg.t.initMonsterSpecies()
	cfg.t.initMonsterType()
	cfg.t.initOverdriveModeType()
	cfg.t.initShopCategory()
	cfg.t.initShopType()
	cfg.t.initTreasureType()
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

	t.AreaConnectionType = newEnumType[database.AreaConnectionType, any]("area connection type", true, typeSlice, func(s string) database.AreaConnectionType {
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

	t.CTBIconType = newEnumType[database.CtbIconType, any]("ctb icon type", true, typeSlice, func(s string) database.CtbIconType {
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

	t.LootType = newEnumType[database.LootType, any]("loot type", false, typeSlice, func(s string) database.LootType {
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

	t.MonsterSpecies = newEnumType[database.MonsterSpecies, any]("monster species", false, typeSlice, func(s string) database.MonsterSpecies {
		return database.MonsterSpecies(s)
	}, nil)
}

func (t *TypeLookup) initMonsterType() {

	typeSlice := []TypedAPIResource{
		{
			Name: string(database.CtbIconTypeMonster),
		},
		{
			Name: string(database.CtbIconTypeBoss),
		},
		{
			Name: string(database.CtbIconTypeSummon),
		},
	}

	t.MonsterType = newEnumType[database.CtbIconType, any]("monster type", false, typeSlice, func(s string) database.CtbIconType {
		return database.CtbIconType(s)
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

	t.OverdriveModeType = newEnumType[database.OverdriveModeType, any]("overdrive mode type", true, typeSlice, func(s string) database.OverdriveModeType {
		return database.OverdriveModeType(s)
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

	t.ShopCategory = newEnumType[database.ShopCategory, any]("shop category", false, typeSlice, func(s string) database.ShopCategory {
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

	t.TreasureType = newEnumType[database.TreasureType, any]("treasure type", true, typeSlice, func(s string) database.TreasureType {
		return database.TreasureType(s)
	}, nil)
}
