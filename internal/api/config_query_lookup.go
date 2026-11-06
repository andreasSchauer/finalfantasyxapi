package api

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

func (Q *QueryLookup) Init() {
	defaultParams := queryparam.GetDefaultParams()

	Q.DefaultParamSlice = defaultParams
	Q.DefaultParams = defaultParamSliceToMap(defaultParams)
	AssignQueryParams(queryparam.GetLocationsParams(), &E.Locations, &Q.Locations)
	AssignQueryParams(queryparam.GetSublocationsParams(), &E.Sublocations, &Q.Sublocations)
	AssignQueryParams(queryparam.GetAreasParams(), &E.Areas, &Q.Areas)

	AssignQueryParams(queryparam.GetMonsterFormationsParams(), &E.MonsterFormations, &Q.MonsterFormations)
	AssignQueryParams(queryparam.GetShopsParams(), &E.Shops, &Q.Shops)
	AssignQueryParams(queryparam.GetTreasuresParams(), &E.Treasures, &Q.Treasures)
	AssignQueryParams(queryparam.GetQuestsParams(), &E.Quests, &Q.Quests)
	AssignQueryParams(queryparam.GetSidequestsParams(), &E.Sidequests, &Q.Sidequests)
	AssignQueryParams(queryparam.GetSubquestsParams(), &E.Subquests, &Q.Subquests)
	AssignQueryParams(queryparam.GetArenaCreationsParams(), &E.ArenaCreations, &Q.ArenaCreations)
	AssignQueryParams(queryparam.GetBlitzballPrizesParams(), &E.BlitzballPrizes, &Q.BlitzballPrizes)
	AssignQueryParams(queryparam.GetSongsParams(), &E.Songs, &Q.Songs)
	AssignQueryParams(queryparam.GetFMVsParams(), &E.FMVs, &Q.FMVs)

	AssignQueryParams(queryparam.GetPlayerUnitsParams(), &E.PlayerUnits, &Q.PlayerUnits)
	AssignQueryParams(queryparam.GetCharactersParams(), &E.Characters, &Q.Characters)
	AssignQueryParams(queryparam.GetAeonsParams(), &E.Aeons, &Q.Aeons)
	AssignQueryParams(queryparam.GetCharacterClassesParams(), &E.CharacterClasses, &Q.CharacterClasses)
	AssignQueryParams(queryparam.GetMonstersParams(), &E.Monsters, &Q.Monsters)

	AssignQueryParams(queryparam.GetAbilitiesParams(), &E.Abilities, &Q.Abilities)
	AssignQueryParams(queryparam.GetPlayerAbilitiesParams(), &E.PlayerAbilities, &Q.PlayerAbilities)
	AssignQueryParams(queryparam.GetOverdriveAbilitiesParams(), &E.OverdriveAbilities, &Q.OverdriveAbilities)
	AssignQueryParams(queryparam.GetItemAbilitiesParams(), &E.ItemAbilities, &Q.ItemAbilities)
	AssignQueryParams(queryparam.GetTriggerCommandsParams(), &E.TriggerCommands, &Q.TriggerCommands)
	AssignQueryParams(queryparam.GetMiscAbilitiesParams(), &E.MiscAbilities, &Q.MiscAbilities)
	AssignQueryParams(queryparam.GetEnemyAbilitiesParams(), &E.EnemyAbilities, &Q.EnemyAbilities)

	AssignDefaultParams(E.AeonCommands.Endpoint, &E.AeonCommands.queryLookup, &Q.AeonCommands)
	AssignDefaultParams(E.OverdriveCommands.Endpoint, &E.OverdriveCommands.queryLookup, &Q.OverdriveCommands)
	AssignQueryParams(queryparam.GetOverdrivesParams(), &E.Overdrives, &Q.Overdrives)
	AssignDefaultParams(E.RonsoRages.Endpoint, &E.RonsoRages.queryLookup, &Q.RonsoRages)
	AssignQueryParams(queryparam.GetSubmenusParams(), &E.Submenus, &Q.Submenus)
	AssignDefaultParams(E.Topmenus.Endpoint, &E.Topmenus.queryLookup, &Q.Topmenus)

	AssignQueryParams(queryparam.GetAllItemsParams(), &E.AllItems, &Q.AllItems)
	AssignQueryParams(queryparam.GetItemsParams(), &E.Items, &Q.Items)
	AssignQueryParams(queryparam.GetKeyItemsParams(), &E.KeyItems, &Q.KeyItems)
	AssignQueryParams(queryparam.GetSpheresParams(), &E.Spheres, &Q.Spheres)
	AssignQueryParams(queryparam.GetPrimersParams(), &E.Primers, &Q.Primers)
	AssignQueryParams(queryparam.GetMixesParams(), &E.Mixes, &Q.Mixes)

	AssignQueryParams(queryparam.GetAutoAbilitiesParams(), &E.AutoAbilities, &Q.AutoAbilities)
	AssignQueryParams(queryparam.GetEquipmentTablesParams(), &E.EquipmentTables, &Q.EquipmentTables)
	AssignQueryParams(queryparam.GetEquipmentParams(), &E.Equipment, &Q.Equipment)
	AssignQueryParams(queryparam.GetCelestialWeaponsParams(), &E.CelestialWeapons, &Q.CelestialWeapons)

	AssignQueryParams(queryparam.GetStatsParams(), &E.Stats, &Q.Stats)
	AssignDefaultParams(E.Properties.Endpoint, &E.Properties.queryLookup, &Q.Properties)
	AssignQueryParams(queryparam.GetOverdriveModesParams(), &E.OverdriveModes, &Q.OverdriveModes)
	AssignDefaultParams(E.Elements.Endpoint, &E.Elements.queryLookup, &Q.Elements)
	AssignQueryParams(queryparam.GetStatusConditionsParams(), &E.StatusConditions, &Q.StatusConditions)
	AssignQueryParams(queryparam.GetModifiersParams(), &E.Modifiers, &Q.Modifiers)
	AssignQueryParams(queryparam.GetAgilityTierParams(), &E.AgilityTiers, &Q.AgilityTiers)

	AssignDefaultParams(E.Enums.Endpoint, &E.Enums.queryLookup, &Q.Enums)
	AssignDefaultParams(E.Endpoints.Endpoint, &E.Endpoints.queryLookup, &Q.Endpoints)
}
