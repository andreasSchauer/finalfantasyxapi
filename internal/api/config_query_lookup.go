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
	Q.DefaultParams = queryparam.DefaultParamSliceToMap(defaultParams)
	assignQueryParams(queryparam.GetLocationsParams(), &E.Locations, &Q.Locations)
	assignQueryParams(queryparam.GetSublocationsParams(), &E.Sublocations, &Q.Sublocations)
	assignQueryParams(queryparam.GetAreasParams(), &E.Areas, &Q.Areas)

	assignQueryParams(queryparam.GetMonsterFormationsParams(), &E.MonsterFormations, &Q.MonsterFormations)
	assignQueryParams(queryparam.GetShopsParams(), &E.Shops, &Q.Shops)
	assignQueryParams(queryparam.GetTreasuresParams(), &E.Treasures, &Q.Treasures)
	assignQueryParams(queryparam.GetQuestsParams(), &E.Quests, &Q.Quests)
	assignQueryParams(queryparam.GetSidequestsParams(), &E.Sidequests, &Q.Sidequests)
	assignQueryParams(queryparam.GetSubquestsParams(), &E.Subquests, &Q.Subquests)
	assignQueryParams(queryparam.GetArenaCreationsParams(), &E.ArenaCreations, &Q.ArenaCreations)
	assignQueryParams(queryparam.GetBlitzballPrizesParams(), &E.BlitzballPrizes, &Q.BlitzballPrizes)
	assignQueryParams(queryparam.GetSongsParams(), &E.Songs, &Q.Songs)
	assignQueryParams(queryparam.GetFMVsParams(), &E.FMVs, &Q.FMVs)

	assignQueryParams(queryparam.GetPlayerUnitsParams(), &E.PlayerUnits, &Q.PlayerUnits)
	assignQueryParams(queryparam.GetCharactersParams(), &E.Characters, &Q.Characters)
	assignQueryParams(queryparam.GetAeonsParams(), &E.Aeons, &Q.Aeons)
	assignQueryParams(queryparam.GetCharacterClassesParams(), &E.CharacterClasses, &Q.CharacterClasses)
	assignQueryParams(queryparam.GetMonstersParams(), &E.Monsters, &Q.Monsters)

	assignQueryParams(queryparam.GetAbilitiesParams(), &E.Abilities, &Q.Abilities)
	assignQueryParams(queryparam.GetPlayerAbilitiesParams(), &E.PlayerAbilities, &Q.PlayerAbilities)
	assignQueryParams(queryparam.GetOverdriveAbilitiesParams(), &E.OverdriveAbilities, &Q.OverdriveAbilities)
	assignQueryParams(queryparam.GetItemAbilitiesParams(), &E.ItemAbilities, &Q.ItemAbilities)
	assignQueryParams(queryparam.GetTriggerCommandsParams(), &E.TriggerCommands, &Q.TriggerCommands)
	assignQueryParams(queryparam.GetMiscAbilitiesParams(), &E.MiscAbilities, &Q.MiscAbilities)
	assignQueryParams(queryparam.GetEnemyAbilitiesParams(), &E.EnemyAbilities, &Q.EnemyAbilities)

	assignDefaultParams(E.AeonCommands.Endpoint, &E.AeonCommands.QueryLookup, &Q.AeonCommands)
	assignDefaultParams(E.OverdriveCommands.Endpoint, &E.OverdriveCommands.QueryLookup, &Q.OverdriveCommands)
	assignQueryParams(queryparam.GetOverdrivesParams(), &E.Overdrives, &Q.Overdrives)
	assignDefaultParams(E.RonsoRages.Endpoint, &E.RonsoRages.QueryLookup, &Q.RonsoRages)
	assignQueryParams(queryparam.GetSubmenusParams(), &E.Submenus, &Q.Submenus)
	assignDefaultParams(E.Topmenus.Endpoint, &E.Topmenus.QueryLookup, &Q.Topmenus)

	assignQueryParams(queryparam.GetAllItemsParams(), &E.AllItems, &Q.AllItems)
	assignQueryParams(queryparam.GetItemsParams(), &E.Items, &Q.Items)
	assignQueryParams(queryparam.GetKeyItemsParams(), &E.KeyItems, &Q.KeyItems)
	assignQueryParams(queryparam.GetSpheresParams(), &E.Spheres, &Q.Spheres)
	assignQueryParams(queryparam.GetPrimersParams(), &E.Primers, &Q.Primers)
	assignQueryParams(queryparam.GetMixesParams(), &E.Mixes, &Q.Mixes)

	assignQueryParams(queryparam.GetAutoAbilitiesParams(), &E.AutoAbilities, &Q.AutoAbilities)
	assignQueryParams(queryparam.GetEquipmentTablesParams(), &E.EquipmentTables, &Q.EquipmentTables)
	assignQueryParams(queryparam.GetEquipmentParams(), &E.Equipment, &Q.Equipment)
	assignQueryParams(queryparam.GetCelestialWeaponsParams(), &E.CelestialWeapons, &Q.CelestialWeapons)

	assignQueryParams(queryparam.GetStatsParams(), &E.Stats, &Q.Stats)
	assignDefaultParams(E.Properties.Endpoint, &E.Properties.QueryLookup, &Q.Properties)
	assignQueryParams(queryparam.GetOverdriveModesParams(), &E.OverdriveModes, &Q.OverdriveModes)
	assignDefaultParams(E.Elements.Endpoint, &E.Elements.QueryLookup, &Q.Elements)
	assignQueryParams(queryparam.GetStatusConditionsParams(), &E.StatusConditions, &Q.StatusConditions)
	assignQueryParams(queryparam.GetModifiersParams(), &E.Modifiers, &Q.Modifiers)
	assignQueryParams(queryparam.GetAgilityTierParams(), &E.AgilityTiers, &Q.AgilityTiers)

	assignDefaultParams(E.Enums.Endpoint, &E.Enums.QueryLookup, &Q.Enums)
	assignDefaultParams(E.Endpoints.Endpoint, &E.Endpoints.QueryLookup, &Q.Endpoints)
}
