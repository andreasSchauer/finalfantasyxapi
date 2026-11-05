package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

var Q = &QueryLookup{}

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	DefaultParamSlice []QueryParam
	DefaultParams     map[ts.QueryParamName]QueryParam

	Locations    map[ts.QueryParamName]QueryParam
	Sublocations map[ts.QueryParamName]QueryParam
	Areas        map[ts.QueryParamName]QueryParam

	MonsterFormations map[ts.QueryParamName]QueryParam
	Shops             map[ts.QueryParamName]QueryParam
	Treasures         map[ts.QueryParamName]QueryParam
	Quests            map[ts.QueryParamName]QueryParam
	Sidequests        map[ts.QueryParamName]QueryParam
	Subquests         map[ts.QueryParamName]QueryParam
	ArenaCreations    map[ts.QueryParamName]QueryParam
	BlitzballPrizes   map[ts.QueryParamName]QueryParam
	Songs             map[ts.QueryParamName]QueryParam
	FMVs              map[ts.QueryParamName]QueryParam

	PlayerUnits      map[ts.QueryParamName]QueryParam
	Characters       map[ts.QueryParamName]QueryParam
	Aeons            map[ts.QueryParamName]QueryParam
	CharacterClasses map[ts.QueryParamName]QueryParam
	Monsters         map[ts.QueryParamName]QueryParam

	Abilities          map[ts.QueryParamName]QueryParam
	PlayerAbilities    map[ts.QueryParamName]QueryParam
	OverdriveAbilities map[ts.QueryParamName]QueryParam
	ItemAbilities      map[ts.QueryParamName]QueryParam
	TriggerCommands    map[ts.QueryParamName]QueryParam
	MiscAbilities      map[ts.QueryParamName]QueryParam
	EnemyAbilities     map[ts.QueryParamName]QueryParam

	AeonCommands      map[ts.QueryParamName]QueryParam
	OverdriveCommands map[ts.QueryParamName]QueryParam
	Overdrives        map[ts.QueryParamName]QueryParam
	RonsoRages        map[ts.QueryParamName]QueryParam
	Submenus          map[ts.QueryParamName]QueryParam
	Topmenus          map[ts.QueryParamName]QueryParam

	AllItems map[ts.QueryParamName]QueryParam
	Items    map[ts.QueryParamName]QueryParam
	KeyItems map[ts.QueryParamName]QueryParam
	Spheres  map[ts.QueryParamName]QueryParam
	Primers  map[ts.QueryParamName]QueryParam
	Mixes    map[ts.QueryParamName]QueryParam

	AutoAbilities    map[ts.QueryParamName]QueryParam
	EquipmentTables  map[ts.QueryParamName]QueryParam
	Equipment        map[ts.QueryParamName]QueryParam
	CelestialWeapons map[ts.QueryParamName]QueryParam

	Stats            map[ts.QueryParamName]QueryParam
	Properties       map[ts.QueryParamName]QueryParam
	OverdriveModes   map[ts.QueryParamName]QueryParam
	Elements         map[ts.QueryParamName]QueryParam
	StatusConditions map[ts.QueryParamName]QueryParam
	Modifiers        map[ts.QueryParamName]QueryParam
	AgilityTiers     map[ts.QueryParamName]QueryParam

	Enums     map[ts.QueryParamName]QueryParam
	Endpoints map[ts.QueryParamName]QueryParam
}

func (Q *QueryLookup) Init() {
	defaultParams := []QueryParam{
		{
			Name:        ts.QpnLimit,
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Type:        ts.QptInt,
			ParamUse:    puList,
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvMax,
					Val: 9999,
				},
			},
			DefaultVal: h.GetIntPtr(20),
		},
		{
			Name:        ts.QpnOffset,
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Type:        ts.QptInt,
			ParamUse:    puList,
			DefaultVal:  h.GetIntPtr(0),
		},
	}

	Q.DefaultParamSlice = defaultParams
	Q.DefaultParams = defaultParamSliceToMap(defaultParams)
	AssignQueryParams(GetLocationsParams(), &E.Locations, &Q.Locations)
	AssignQueryParams(GetSublocationsParams(), &E.Sublocations, &Q.Sublocations)
	AssignQueryParams(GetAreasParams(), &E.Areas, &Q.Areas)

	AssignQueryParams(GetMonsterFormationsParams(), &E.MonsterFormations, &Q.MonsterFormations)
	AssignQueryParams(GetShopsParams(), &E.Shops, &Q.Shops)
	AssignQueryParams(GetTreasuresParams(), &E.Treasures, &Q.Treasures)
	AssignQueryParams(GetQuestsParams(), &E.Quests, &Q.Quests)
	AssignQueryParams(GetSidequestsParams(), &E.Sidequests, &Q.Sidequests)
	AssignQueryParams(GetSubquestsParams(), &E.Subquests, &Q.Subquests)
	AssignQueryParams(GetArenaCreationsParams(), &E.ArenaCreations, &Q.ArenaCreations)
	AssignQueryParams(GetBlitzballPrizesParams(), &E.BlitzballPrizes, &Q.BlitzballPrizes)
	AssignQueryParams(GetSongsParams(), &E.Songs, &Q.Songs)
	AssignQueryParams(GetFMVsParams(), &E.FMVs, &Q.FMVs)

	AssignQueryParams(GetPlayerUnitsParams(), &E.PlayerUnits, &Q.PlayerUnits)
	AssignQueryParams(GetCharactersParams(), &E.Characters, &Q.Characters)
	AssignQueryParams(GetAeonsParams(), &E.Aeons, &Q.Aeons)
	AssignQueryParams(GetCharacterClassesParams(), &E.CharacterClasses, &Q.CharacterClasses)
	AssignQueryParams(GetMonstersParams(), &E.Monsters, &Q.Monsters)

	AssignQueryParams(GetAbilitiesParams(), &E.Abilities, &Q.Abilities)
	AssignQueryParams(GetPlayerAbilitiesParams(), &E.PlayerAbilities, &Q.PlayerAbilities)
	AssignQueryParams(GetOverdriveAbilitiesParams(), &E.OverdriveAbilities, &Q.OverdriveAbilities)
	AssignQueryParams(GetItemAbilitiesParams(), &E.ItemAbilities, &Q.ItemAbilities)
	AssignQueryParams(GetTriggerCommandsParams(), &E.TriggerCommands, &Q.TriggerCommands)
	AssignQueryParams(GetMiscAbilitiesParams(), &E.MiscAbilities, &Q.MiscAbilities)
	AssignQueryParams(GetEnemyAbilitiesParams(), &E.EnemyAbilities, &Q.EnemyAbilities)

	AssignDefaultParams(E.AeonCommands.Endpoint, &E.AeonCommands.queryLookup, &Q.AeonCommands)
	AssignDefaultParams(E.OverdriveCommands.Endpoint, &E.OverdriveCommands.queryLookup, &Q.OverdriveCommands)
	AssignQueryParams(GetOverdrivesParams(), &E.Overdrives, &Q.Overdrives)
	AssignDefaultParams(E.RonsoRages.Endpoint, &E.RonsoRages.queryLookup, &Q.RonsoRages)
	AssignQueryParams(GetSubmenusParams(), &E.Submenus, &Q.Submenus)
	AssignDefaultParams(E.Topmenus.Endpoint, &E.Topmenus.queryLookup, &Q.Topmenus)

	AssignQueryParams(GetAllItemsParams(), &E.AllItems, &Q.AllItems)
	AssignQueryParams(GetItemsParams(), &E.Items, &Q.Items)
	AssignQueryParams(GetKeyItemsParams(), &E.KeyItems, &Q.KeyItems)
	AssignQueryParams(GetSpheresParams(), &E.Spheres, &Q.Spheres)
	AssignQueryParams(GetPrimersParams(), &E.Primers, &Q.Primers)
	AssignQueryParams(GetMixesParams(), &E.Mixes, &Q.Mixes)

	AssignQueryParams(GetAutoAbilitiesParams(), &E.AutoAbilities, &Q.AutoAbilities)
	AssignQueryParams(GetEquipmentTablesParams(), &E.EquipmentTables, &Q.EquipmentTables)
	AssignQueryParams(GetEquipmentParams(), &E.Equipment, &Q.Equipment)
	AssignQueryParams(GetCelestialWeaponsParams(), &E.CelestialWeapons, &Q.CelestialWeapons)

	AssignQueryParams(GetStatsParams(), &E.Stats, &Q.Stats)
	AssignDefaultParams(E.Properties.Endpoint, &E.Properties.queryLookup, &Q.Properties)
	AssignQueryParams(GetOverdriveModesParams(), &E.OverdriveModes, &Q.OverdriveModes)
	AssignDefaultParams(E.Elements.Endpoint, &E.Elements.queryLookup, &Q.Elements)
	AssignQueryParams(GetStatusConditionsParams(), &E.StatusConditions, &Q.StatusConditions)
	AssignQueryParams(GetModifiersParams(), &E.Modifiers, &Q.Modifiers)
	AssignQueryParams(GetAgilityTierParams(), &E.AgilityTiers, &Q.AgilityTiers)

	AssignDefaultParams(E.Enums.Endpoint, &E.Enums.queryLookup, &Q.Enums)
	AssignDefaultParams(E.Endpoints.Endpoint, &E.Endpoints.queryLookup, &Q.Endpoints)
}
