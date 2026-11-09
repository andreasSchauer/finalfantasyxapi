package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
)

func InitQueryLookup(Q *QueryLookup) {
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
