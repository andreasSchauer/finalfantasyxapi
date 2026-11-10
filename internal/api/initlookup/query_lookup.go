package initlookup

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
)

func Queries() {
	defaultParams := queryparam.GetDefaultParams()

	apilookups.Q.DefaultParamSlice = defaultParams
	apilookups.Q.DefaultParams = queryparam.DefaultParamSliceToMap(defaultParams)
	assignQueryParams(queryparam.GetLocationsParams(), &apilookups.E.Locations, &apilookups.Q.Locations)
	assignQueryParams(queryparam.GetSublocationsParams(), &apilookups.E.Sublocations, &apilookups.Q.Sublocations)
	assignQueryParams(queryparam.GetAreasParams(), &apilookups.E.Areas, &apilookups.Q.Areas)

	assignQueryParams(queryparam.GetMonsterFormationsParams(), &apilookups.E.MonsterFormations, &apilookups.Q.MonsterFormations)
	assignQueryParams(queryparam.GetShopsParams(), &apilookups.E.Shops, &apilookups.Q.Shops)
	assignQueryParams(queryparam.GetTreasuresParams(), &apilookups.E.Treasures, &apilookups.Q.Treasures)
	assignQueryParams(queryparam.GetQuestsParams(), &apilookups.E.Quests, &apilookups.Q.Quests)
	assignQueryParams(queryparam.GetSidequestsParams(), &apilookups.E.Sidequests, &apilookups.Q.Sidequests)
	assignQueryParams(queryparam.GetSubquestsParams(), &apilookups.E.Subquests, &apilookups.Q.Subquests)
	assignQueryParams(queryparam.GetArenaCreationsParams(), &apilookups.E.ArenaCreations, &apilookups.Q.ArenaCreations)
	assignQueryParams(queryparam.GetBlitzballPrizesParams(), &apilookups.E.BlitzballPrizes, &apilookups.Q.BlitzballPrizes)
	assignQueryParams(queryparam.GetSongsParams(), &apilookups.E.Songs, &apilookups.Q.Songs)
	assignQueryParams(queryparam.GetFMVsParams(), &apilookups.E.FMVs, &apilookups.Q.FMVs)

	assignQueryParams(queryparam.GetPlayerUnitsParams(), &apilookups.E.PlayerUnits, &apilookups.Q.PlayerUnits)
	assignQueryParams(queryparam.GetCharactersParams(), &apilookups.E.Characters, &apilookups.Q.Characters)
	assignQueryParams(queryparam.GetAeonsParams(), &apilookups.E.Aeons, &apilookups.Q.Aeons)
	assignQueryParams(queryparam.GetCharacterClassesParams(), &apilookups.E.CharacterClasses, &apilookups.Q.CharacterClasses)
	assignQueryParams(queryparam.GetMonstersParams(), &apilookups.E.Monsters, &apilookups.Q.Monsters)

	assignQueryParams(queryparam.GetAbilitiesParams(), &apilookups.E.Abilities, &apilookups.Q.Abilities)
	assignQueryParams(queryparam.GetPlayerAbilitiesParams(), &apilookups.E.PlayerAbilities, &apilookups.Q.PlayerAbilities)
	assignQueryParams(queryparam.GetOverdriveAbilitiesParams(), &apilookups.E.OverdriveAbilities, &apilookups.Q.OverdriveAbilities)
	assignQueryParams(queryparam.GetItemAbilitiesParams(), &apilookups.E.ItemAbilities, &apilookups.Q.ItemAbilities)
	assignQueryParams(queryparam.GetTriggerCommandsParams(), &apilookups.E.TriggerCommands, &apilookups.Q.TriggerCommands)
	assignQueryParams(queryparam.GetMiscAbilitiesParams(), &apilookups.E.MiscAbilities, &apilookups.Q.MiscAbilities)
	assignQueryParams(queryparam.GetEnemyAbilitiesParams(), &apilookups.E.EnemyAbilities, &apilookups.Q.EnemyAbilities)

	assignDefaultParams(apilookups.E.AeonCommands.Endpoint, &apilookups.E.AeonCommands.QueryLookup, &apilookups.Q.AeonCommands)
	assignDefaultParams(apilookups.E.OverdriveCommands.Endpoint, &apilookups.E.OverdriveCommands.QueryLookup, &apilookups.Q.OverdriveCommands)
	assignQueryParams(queryparam.GetOverdrivesParams(), &apilookups.E.Overdrives, &apilookups.Q.Overdrives)
	assignDefaultParams(apilookups.E.RonsoRages.Endpoint, &apilookups.E.RonsoRages.QueryLookup, &apilookups.Q.RonsoRages)
	assignQueryParams(queryparam.GetSubmenusParams(), &apilookups.E.Submenus, &apilookups.Q.Submenus)
	assignDefaultParams(apilookups.E.Topmenus.Endpoint, &apilookups.E.Topmenus.QueryLookup, &apilookups.Q.Topmenus)

	assignQueryParams(queryparam.GetAllItemsParams(), &apilookups.E.AllItems, &apilookups.Q.AllItems)
	assignQueryParams(queryparam.GetItemsParams(), &apilookups.E.Items, &apilookups.Q.Items)
	assignQueryParams(queryparam.GetKeyItemsParams(), &apilookups.E.KeyItems, &apilookups.Q.KeyItems)
	assignQueryParams(queryparam.GetSpheresParams(), &apilookups.E.Spheres, &apilookups.Q.Spheres)
	assignQueryParams(queryparam.GetPrimersParams(), &apilookups.E.Primers, &apilookups.Q.Primers)
	assignQueryParams(queryparam.GetMixesParams(), &apilookups.E.Mixes, &apilookups.Q.Mixes)

	assignQueryParams(queryparam.GetAutoAbilitiesParams(), &apilookups.E.AutoAbilities, &apilookups.Q.AutoAbilities)
	assignQueryParams(queryparam.GetEquipmentTablesParams(), &apilookups.E.EquipmentTables, &apilookups.Q.EquipmentTables)
	assignQueryParams(queryparam.GetEquipmentParams(), &apilookups.E.Equipment, &apilookups.Q.Equipment)
	assignQueryParams(queryparam.GetCelestialWeaponsParams(), &apilookups.E.CelestialWeapons, &apilookups.Q.CelestialWeapons)

	assignQueryParams(queryparam.GetStatsParams(), &apilookups.E.Stats, &apilookups.Q.Stats)
	assignDefaultParams(apilookups.E.Properties.Endpoint, &apilookups.E.Properties.QueryLookup, &apilookups.Q.Properties)
	assignQueryParams(queryparam.GetOverdriveModesParams(), &apilookups.E.OverdriveModes, &apilookups.Q.OverdriveModes)
	assignDefaultParams(apilookups.E.Elements.Endpoint, &apilookups.E.Elements.QueryLookup, &apilookups.Q.Elements)
	assignQueryParams(queryparam.GetStatusConditionsParams(), &apilookups.E.StatusConditions, &apilookups.Q.StatusConditions)
	assignQueryParams(queryparam.GetModifiersParams(), &apilookups.E.Modifiers, &apilookups.Q.Modifiers)
	assignQueryParams(queryparam.GetAgilityTierParams(), &apilookups.E.AgilityTiers, &apilookups.Q.AgilityTiers)

	assignDefaultParams(apilookups.E.Enums.Endpoint, &apilookups.E.Enums.QueryLookup, &apilookups.Q.Enums)
	assignDefaultParams(apilookups.E.Endpoints.Endpoint, &apilookups.E.Endpoints.QueryLookup, &apilookups.Q.Endpoints)
}
