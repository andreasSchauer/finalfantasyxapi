package main

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// TypesLookup holds all the enum types for the application
type TypeLookup struct {
	AreaConnectionType map[string]TypedAPIResource
	CTBIconType        map[string]TypedAPIResource
	MaCreationArea     map[string]TypedAPIResource
	MonsterSpecies     map[string]TypedAPIResource
	OverdriveModeType  map[string]TypedAPIResource
}

func TypeLookupInit() TypeLookup {
	t := TypeLookup{}

	t.initAreaConnectionType()
	t.initCTBIconType()
	t.initMaCreationArea()
	t.initMonsterSpecies()
	t.initOverdriveModeType()

	return t
}

func GetEnumType(key string, lookup map[string]TypedAPIResource) (TypedAPIResource, error) {
	enumType, found := lookup[key]
	if !found {
		return TypedAPIResource{}, fmt.Errorf("value %s is not valid in this enum", key)
	}

	return enumType, nil
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

	t.AreaConnectionType = typeSliceToMap(typeSlice)
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
	}

	t.CTBIconType = typeSliceToMap(typeSlice)
}

func (t *TypeLookup) initMaCreationArea() {
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

	t.MaCreationArea = typeSliceToMap(typeSlice)
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

	t.MonsterSpecies = typeSliceToMap(typeSlice)
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

	t.OverdriveModeType = typeSliceToMap(typeSlice)
}
