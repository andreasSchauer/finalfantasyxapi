package lookup

import (
	"path/filepath"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func (L *Lookup) Cache() error {
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(writeLookupFile("abilities.json", L.AbilitiesID))
	checkErr(writeLookupFile("misc_abilities.json", L.MiscAbilitiesID))
	checkErr(writeLookupFile("enemy_abilities.json", L.EnemyAbilitiesID))
	checkErr(writeLookupFile("item_abilities.json", L.ItemAbilitiesID))
	checkErr(writeLookupFile("overdrive_abilities.json", L.OverdriveAbilitiesID))
	checkErr(writeLookupFile("player_abilities.json", L.PlayerAbilitiesID))
	checkErr(writeLookupFile("trigger_commands.json", L.TriggerCommandsID))
	checkErr(writeLookupFile("player_units.json", L.PlayerUnitsID))
	checkErr(writeLookupFile("aeon_stats.json", L.AeonsID))
	checkErr(writeLookupFile("aeon_commands.json", L.AeonCommandsID))
	checkErr(writeLookupFile("agility_tiers.json", L.AgilityTiersID))
	checkErr(writeLookupFile("monster_arena_creations.json", L.ArenaCreationsID))
	checkErr(writeLookupFile("locations.json", L.LocationsID))
	checkErr(writeLookupFile("sublocations.json", L.SublocationsID))
	checkErr(writeLookupFile("areas.json", L.AreasID))
	checkErr(writeLookupFile("auto_abilities.json", L.AutoAbilitiesID))
	checkErr(writeLookupFile("celestial_weapons.json", L.CelestialWeaponsID))
	checkErr(writeLookupFile("characters.json", L.CharactersID))
	checkErr(writeLookupFile("character_classes.json", L.CharClassesID))
	checkErr(writeLookupFile("elements.json", L.ElementsID))
	checkErr(writeLookupFile("elemental_resists.json", L.ElementalResistsID))
	checkErr(writeLookupFile("equipment_names.json", L.EquipmentNamesID))
	checkErr(writeLookupFile("equipment_tables.json", L.EquipmentTablesID))
	checkErr(writeLookupFile("encounter_areas.json", L.EncounterAreasID))
	checkErr(writeLookupFile("fmvs.json", L.FMVsID))
	checkErr(writeLookupFile("items.json", L.ItemsID))
	checkErr(writeLookupFile("key_items.json", L.KeyItemsID))
	checkErr(writeLookupFile("master_items.json", L.MasterItemsID))
	checkErr(writeLookupFile("mixes.json", L.MixesID))
	checkErr(writeLookupFile("modifiers.json", L.ModifiersID))
	checkErr(writeLookupFile("monsters.json", L.MonstersID))
	checkErr(writeLookupFile("monster_formations.json", L.MonsterFormationsID))
	checkErr(writeLookupFile("overdrive_modes.json", L.OverdriveModesID))
	checkErr(writeLookupFile("overdrive_commands.json", L.OverdriveCommandsID))
	checkErr(writeLookupFile("overdrives.json", L.OverdrivesID))
	checkErr(writeLookupFile("blitzball_positions.json", L.PositionsID))
	checkErr(writeLookupFile("primers.json", L.PrimersID))
	checkErr(writeLookupFile("properties.json", L.PropertiesID))
	checkErr(writeLookupFile("quests.json", L.QuestsID))
	checkErr(writeLookupFile("ronso_rages.json", L.RonsoRagesID))
	checkErr(writeLookupFile("sidequests.json", L.SidequestsID))
	checkErr(writeLookupFile("subquests.json", L.SubquestsID))
	checkErr(writeLookupFile("shops.json", L.ShopsID))
	checkErr(writeLookupFile("songs.json", L.SongsID))
	checkErr(writeLookupFile("spheres.json", L.SpheresID))
	checkErr(writeLookupFile("stats.json", L.StatsID))
	checkErr(writeLookupFile("status_conditions.json", L.StatusConditionsID))
	checkErr(writeLookupFile("submenus.json", L.SubmenusID))
	checkErr(writeLookupFile("topmenus.json", L.TopmenusID))
	checkErr(writeLookupFile("treasures.json", L.TreasuresID))

	return err
}

func writeLookupFile[T seedtypes.Lookupable](fileName string, lookup map[int32]T) error {
	return h.WriteFile("data_lookups", fileName, lookupToSlice(lookup))
}

func lookupToSlice[T seedtypes.Lookupable](lookup map[int32]T) []T {
	var s []T

	for _, obj := range lookup {
		s = append(s, obj)
	}

	slices.SortStableFunc(s, h.SortOnID)

	return s
}

func (L *Lookup) InitCached() error {
	defer h.MeasureTime("assigning lookups")()
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(assignLookup("abilities.json", &L.Abilities, &L.AbilitiesID))
	checkErr(assignLookup("misc_abilities.json", &L.MiscAbilities, &L.MiscAbilitiesID))
	checkErr(assignLookup("enemy_abilities.json", &L.EnemyAbilities, &L.EnemyAbilitiesID))
	checkErr(assignLookup("item_abilities.json", &L.ItemAbilities, &L.ItemAbilitiesID))
	checkErr(assignLookup("overdrive_abilities.json", &L.OverdriveAbilities, &L.OverdriveAbilitiesID))
	checkErr(assignLookup("player_abilities.json", &L.PlayerAbilities, &L.PlayerAbilitiesID))
	checkErr(assignLookup("trigger_commands.json", &L.TriggerCommands, &L.TriggerCommandsID))
	checkErr(assignLookup("player_units.json", &L.PlayerUnits, &L.PlayerUnitsID))
	checkErr(assignLookup("aeon_stats.json", &L.Aeons, &L.AeonsID))
	checkErr(assignLookup("aeon_commands.json", &L.AeonCommands, &L.AeonCommandsID))
	checkErr(assignLookup("agility_tiers.json", nil, &L.AgilityTiersID))
	checkErr(assignLookup("monster_arena_creations.json", &L.ArenaCreations, &L.ArenaCreationsID))
	checkErr(assignLookup("locations.json", &L.Locations, &L.LocationsID))
	checkErr(assignLookup("sublocations.json", &L.Sublocations, &L.SublocationsID))
	checkErr(assignLookup("areas.json", &L.Areas, &L.AreasID))
	checkErr(assignLookup("auto_abilities.json", &L.AutoAbilities, &L.AutoAbilitiesID))
	checkErr(assignLookup("celestial_weapons.json", &L.CelestialWeapons, &L.CelestialWeaponsID))
	checkErr(assignLookup("characters.json", &L.Characters, &L.CharactersID))
	checkErr(assignLookup("character_classes.json", &L.CharClasses, &L.CharClassesID))
	checkErr(assignLookup("elements.json", &L.Elements, &L.ElementsID))
	checkErr(assignLookup("elemental_resists.json", &L.ElementalResists, &L.ElementalResistsID))
	checkErr(assignLookup("equipment_names.json", &L.EquipmentNames, &L.EquipmentNamesID))
	checkErr(assignLookup("equipment_tables.json", &L.EquipmentTables, &L.EquipmentTablesID))
	checkErr(assignLookup("encounter_areas.json", &L.EncounterAreas, &L.EncounterAreasID))
	checkErr(assignLookup("fmvs.json", &L.FMVs, &L.FMVsID))
	checkErr(assignLookup("items.json", &L.Items, &L.ItemsID))
	checkErr(assignLookup("key_items.json", &L.KeyItems, &L.KeyItemsID))
	checkErr(assignLookup("master_items.json", &L.MasterItems, &L.MasterItemsID))
	checkErr(assignLookup("mixes.json", &L.Mixes, &L.MixesID))
	checkErr(assignLookup("modifiers.json", &L.Modifiers, &L.ModifiersID))
	checkErr(assignLookup("monsters.json", &L.Monsters, &L.MonstersID))
	checkErr(assignLookup("monster_formations.json", &L.MonsterFormations, &L.MonsterFormationsID))
	checkErr(assignLookup("overdrive_modes.json", &L.OverdriveModes, &L.OverdriveModesID))
	checkErr(assignLookup("overdrive_commands.json", &L.OverdriveCommands, &L.OverdriveCommandsID))
	checkErr(assignLookup("overdrives.json", &L.Overdrives, &L.OverdrivesID))
	checkErr(assignLookup("blitzball_positions.json", &L.Positions, &L.PositionsID))
	checkErr(assignLookup("primers.json", &L.Primers, &L.PrimersID))
	checkErr(assignLookup("properties.json", &L.Properties, &L.PropertiesID))
	checkErr(assignLookup("quests.json", &L.Quests, &L.QuestsID))
	checkErr(assignLookup("ronso_rages.json", &L.RonsoRages, &L.RonsoRagesID))
	checkErr(assignLookup("sidequests.json", &L.Sidequests, &L.SidequestsID))
	checkErr(assignLookup("subquests.json", &L.Subquests, &L.SubquestsID))
	checkErr(assignLookup("shops.json", &L.Shops, &L.ShopsID))
	checkErr(assignLookup("songs.json", &L.Songs, &L.SongsID))
	checkErr(assignLookup("spheres.json", &L.Spheres, &L.SpheresID))
	checkErr(assignLookup("stats.json", &L.Stats, &L.StatsID))
	checkErr(assignLookup("status_conditions.json", &L.StatusConditions, &L.StatusConditionsID))
	checkErr(assignLookup("submenus.json", &L.Submenus, &L.SubmenusID))
	checkErr(assignLookup("topmenus.json", &L.Topmenus, &L.TopmenusID))
	checkErr(assignLookup("treasures.json", &L.Treasures, &L.TreasuresID))

	return err
}

func assignLookup[T seedtypes.Lookupable](fileName string, keyLookup *map[string]T, idLookup *map[int32]T) error {
	lookupPath := filepath.Join("data_lookups", fileName)

	var dataSlice []T
	err := h.LoadJSONFile(lookupPath, &dataSlice)
	if err != nil {
		return err
	}

	if keyLookup != nil {
		*keyLookup = sliceToKeyLookup(dataSlice)
	}

	if idLookup != nil {
		*idLookup = sliceToIDLookup(dataSlice)
	}

	return nil
}

func sliceToKeyLookup[T seedtypes.Lookupable](slice []T) map[string]T {
	lookup := make(map[string]T, len(slice))

	for _, obj := range slice {
		lookup[seedtypes.Key(obj)] = obj
	}

	return lookup
}

func sliceToIDLookup[T seedtypes.Lookupable](slice []T) map[int32]T {
	lookup := make(map[int32]T, len(slice))

	for _, obj := range slice {
		lookup[obj.GetID()] = obj
	}

	return lookup
}
