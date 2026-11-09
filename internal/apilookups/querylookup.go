package apilookups

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

var Q = &QueryLookup{}

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	DefaultParamSlice []queryparam.QueryParam
	DefaultParams     map[ts.QueryParamName]queryparam.QueryParam

	Locations    map[ts.QueryParamName]queryparam.QueryParam
	Sublocations map[ts.QueryParamName]queryparam.QueryParam
	Areas        map[ts.QueryParamName]queryparam.QueryParam

	MonsterFormations map[ts.QueryParamName]queryparam.QueryParam
	Shops             map[ts.QueryParamName]queryparam.QueryParam
	Treasures         map[ts.QueryParamName]queryparam.QueryParam
	Quests            map[ts.QueryParamName]queryparam.QueryParam
	Sidequests        map[ts.QueryParamName]queryparam.QueryParam
	Subquests         map[ts.QueryParamName]queryparam.QueryParam
	ArenaCreations    map[ts.QueryParamName]queryparam.QueryParam
	BlitzballPrizes   map[ts.QueryParamName]queryparam.QueryParam
	Songs             map[ts.QueryParamName]queryparam.QueryParam
	FMVs              map[ts.QueryParamName]queryparam.QueryParam

	PlayerUnits      map[ts.QueryParamName]queryparam.QueryParam
	Characters       map[ts.QueryParamName]queryparam.QueryParam
	Aeons            map[ts.QueryParamName]queryparam.QueryParam
	CharacterClasses map[ts.QueryParamName]queryparam.QueryParam
	Monsters         map[ts.QueryParamName]queryparam.QueryParam

	Abilities          map[ts.QueryParamName]queryparam.QueryParam
	PlayerAbilities    map[ts.QueryParamName]queryparam.QueryParam
	OverdriveAbilities map[ts.QueryParamName]queryparam.QueryParam
	ItemAbilities      map[ts.QueryParamName]queryparam.QueryParam
	TriggerCommands    map[ts.QueryParamName]queryparam.QueryParam
	MiscAbilities      map[ts.QueryParamName]queryparam.QueryParam
	EnemyAbilities     map[ts.QueryParamName]queryparam.QueryParam

	AeonCommands      map[ts.QueryParamName]queryparam.QueryParam
	OverdriveCommands map[ts.QueryParamName]queryparam.QueryParam
	Overdrives        map[ts.QueryParamName]queryparam.QueryParam
	RonsoRages        map[ts.QueryParamName]queryparam.QueryParam
	Submenus          map[ts.QueryParamName]queryparam.QueryParam
	Topmenus          map[ts.QueryParamName]queryparam.QueryParam

	AllItems map[ts.QueryParamName]queryparam.QueryParam
	Items    map[ts.QueryParamName]queryparam.QueryParam
	KeyItems map[ts.QueryParamName]queryparam.QueryParam
	Spheres  map[ts.QueryParamName]queryparam.QueryParam
	Primers  map[ts.QueryParamName]queryparam.QueryParam
	Mixes    map[ts.QueryParamName]queryparam.QueryParam

	AutoAbilities    map[ts.QueryParamName]queryparam.QueryParam
	EquipmentTables  map[ts.QueryParamName]queryparam.QueryParam
	Equipment        map[ts.QueryParamName]queryparam.QueryParam
	CelestialWeapons map[ts.QueryParamName]queryparam.QueryParam

	Stats            map[ts.QueryParamName]queryparam.QueryParam
	Properties       map[ts.QueryParamName]queryparam.QueryParam
	OverdriveModes   map[ts.QueryParamName]queryparam.QueryParam
	Elements         map[ts.QueryParamName]queryparam.QueryParam
	StatusConditions map[ts.QueryParamName]queryparam.QueryParam
	Modifiers        map[ts.QueryParamName]queryparam.QueryParam
	AgilityTiers     map[ts.QueryParamName]queryparam.QueryParam

	Enums     map[ts.QueryParamName]queryparam.QueryParam
	Endpoints map[ts.QueryParamName]queryparam.QueryParam
}