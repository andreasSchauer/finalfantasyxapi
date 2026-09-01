package seeding

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



func (l *Lookup) saveCompletedLookups() error {
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(writeLookupFile("abilities.json", l.AbilitiesID))
	checkErr(writeLookupFile("misc_abilities.json", l.MiscAbilitiesID))
	checkErr(writeLookupFile("enemy_abilities.json", l.EnemyAbilitiesID))
	checkErr(writeLookupFile("item_abilities.json", l.ItemAbilitiesID))
	checkErr(writeLookupFile("overdrive_abilities.json", l.OverdriveAbilitiesID))
	checkErr(writeLookupFile("player_abilities.json", l.PlayerAbilitiesID))
	checkErr(writeLookupFile("trigger_commands.json", l.TriggerCommandsID))
	checkErr(writeLookupFile("player_units.json", l.PlayerUnitsID))
	checkErr(writeLookupFile("aeon_stats.json", l.AeonsID))
	checkErr(writeLookupFile("aeon_commands.json", l.AeonCommandsID))
	checkErr(writeLookupFile("agility_tiers.json", l.AgilityTiersID))
	checkErr(writeLookupFile("monster_arena_creations.json", l.ArenaCreationsID))
	checkErr(writeLookupFile("locations.json", l.LocationsID))
	checkErr(writeLookupFile("sublocations.json", l.SublocationsID))
	checkErr(writeLookupFile("areas.json", l.AreasID))
	checkErr(writeLookupFile("auto_abilities.json", l.AutoAbilitiesID))
	checkErr(writeLookupFile("celestial_weapons.json", l.CelestialWeaponsID))
	checkErr(writeLookupFile("characters.json", l.CharactersID))
	checkErr(writeLookupFile("character_classes.json", l.CharClassesID))
	checkErr(writeLookupFile("elements.json", l.ElementsID))
	checkErr(writeLookupFile("elemental_resists.json", l.ElementalResistsID))
	checkErr(writeLookupFile("equipment_names.json", l.EquipmentNamesID))
	checkErr(writeLookupFile("equipment_tables.json", l.EquipmentTablesID))
	checkErr(writeLookupFile("encounter_areas.json", l.EncounterAreasID))
	checkErr(writeLookupFile("fmvs.json", l.FMVsID))
	checkErr(writeLookupFile("items.json", l.ItemsID))
	checkErr(writeLookupFile("key_items.json", l.KeyItemsID))
	checkErr(writeLookupFile("master_items.json", l.MasterItemsID))
	checkErr(writeLookupFile("mixes.json", l.MixesID))
	checkErr(writeLookupFile("modifiers.json", l.ModifiersID))
	checkErr(writeLookupFile("monsters.json", l.MonstersID))
	checkErr(writeLookupFile("monster_formations.json", l.MonsterFormationsID))
	checkErr(writeLookupFile("overdrive_modes.json", l.OverdriveModesID))
	checkErr(writeLookupFile("overdrive_commands.json", l.OverdriveCommandsID))
	checkErr(writeLookupFile("overdrives.json", l.OverdrivesID))
	checkErr(writeLookupFile("blitzball_positions.json", l.PositionsID))
	checkErr(writeLookupFile("primers.json", l.PrimersID))
	checkErr(writeLookupFile("properties.json", l.PropertiesID))
	checkErr(writeLookupFile("quests.json", l.QuestsID))
	checkErr(writeLookupFile("ronso_rages.json", l.RonsoRagesID))
	checkErr(writeLookupFile("sidequests.json", l.SidequestsID))
	checkErr(writeLookupFile("subquests.json", l.SubquestsID))
	checkErr(writeLookupFile("shops.json", l.ShopsID))
	checkErr(writeLookupFile("songs.json", l.SongsID))
	checkErr(writeLookupFile("spheres.json", l.SpheresID))
	checkErr(writeLookupFile("stats.json", l.StatsID))
	checkErr(writeLookupFile("status_conditions.json", l.StatusConditionsID))
	checkErr(writeLookupFile("submenus.json", l.SubmenusID))
	checkErr(writeLookupFile("topmenus.json", l.TopmenusID))
	checkErr(writeLookupFile("treasures.json", l.TreasuresID))
	checkErr(l.saveHashMap())

	return err
}

func writeLookupFile[T Lookupable](fileName string, lookup map[int32]T) error {
	lookupDir, err := h.GetAbsoluteFilepath("data_lookups")
	if err != nil {
		return err
	}

	err = os.MkdirAll(lookupDir, 0755)
	if err != nil {
		return err
	}

	slice := lookupToSlice(lookup)

	jsonBytes, err := json.MarshalIndent(slice, "", "    ")
	if err != nil {
		return err
	}

	destPath := filepath.Join(lookupDir, fileName)

	err = os.WriteFile(destPath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) assignLookups() error {
	defer h.MeasureTime("assigning lookups")()
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(assignLookup("abilities.json", &l.Abilities, &l.AbilitiesID))
	checkErr(assignLookup("misc_abilities.json", &l.MiscAbilities, &l.MiscAbilitiesID))
	checkErr(assignLookup("enemy_abilities.json", &l.EnemyAbilities, &l.EnemyAbilitiesID))
	checkErr(assignLookup("item_abilities.json", &l.ItemAbilities, &l.ItemAbilitiesID))
	checkErr(assignLookup("overdrive_abilities.json", &l.OverdriveAbilities, &l.OverdriveAbilitiesID))
	checkErr(assignLookup("player_abilities.json", &l.PlayerAbilities, &l.PlayerAbilitiesID))
	checkErr(assignLookup("trigger_commands.json", &l.TriggerCommands, &l.TriggerCommandsID))
	checkErr(assignLookup("player_units.json", &l.PlayerUnits, &l.PlayerUnitsID))
	checkErr(assignLookup("aeon_stats.json", &l.Aeons, &l.AeonsID))
	checkErr(assignLookup("aeon_commands.json", &l.AeonCommands, &l.AeonCommandsID))
	checkErr(assignLookup("agility_tiers.json", nil, &l.AgilityTiersID))
	checkErr(assignLookup("monster_arena_creations.json", &l.ArenaCreations, &l.ArenaCreationsID))
	checkErr(assignLookup("locations.json", &l.Locations, &l.LocationsID))
	checkErr(assignLookup("sublocations.json", &l.Sublocations, &l.SublocationsID))
	checkErr(assignLookup("areas.json", &l.Areas, &l.AreasID))
	checkErr(assignLookup("auto_abilities.json", &l.AutoAbilities, &l.AutoAbilitiesID))
	checkErr(assignLookup("celestial_weapons.json", &l.CelestialWeapons, &l.CelestialWeaponsID))
	checkErr(assignLookup("characters.json", &l.Characters, &l.CharactersID))
	checkErr(assignLookup("character_classes.json", &l.CharClasses, &l.CharClassesID))
	checkErr(assignLookup("elements.json", &l.Elements, &l.ElementsID))
	checkErr(assignLookup("elemental_resists.json", &l.ElementalResists, &l.ElementalResistsID))
	checkErr(assignLookup("equipment_names.json", &l.EquipmentNames, &l.EquipmentNamesID))
	checkErr(assignLookup("equipment_tables.json", &l.EquipmentTables, &l.EquipmentTablesID))
	checkErr(assignLookup("encounter_areas.json", &l.EncounterAreas, &l.EncounterAreasID))
	checkErr(assignLookup("fmvs.json", &l.FMVs, &l.FMVsID))
	checkErr(assignLookup("items.json", &l.Items, &l.ItemsID))
	checkErr(assignLookup("key_items.json", &l.KeyItems, &l.KeyItemsID))
	checkErr(assignLookup("master_items.json", &l.MasterItems, &l.MasterItemsID))
	checkErr(assignLookup("mixes.json", &l.Mixes, &l.MixesID))
	checkErr(assignLookup("modifiers.json", &l.Modifiers, &l.ModifiersID))
	checkErr(assignLookup("monsters.json", &l.Monsters, &l.MonstersID))
	checkErr(assignLookup("monster_formations.json", &l.MonsterFormations, &l.MonsterFormationsID))
	checkErr(assignLookup("overdrive_modes.json", &l.OverdriveModes, &l.OverdriveModesID))
	checkErr(assignLookup("overdrive_commands.json", &l.OverdriveCommands, &l.OverdriveCommandsID))
	checkErr(assignLookup("overdrives.json", &l.Overdrives, &l.OverdrivesID))
	checkErr(assignLookup("blitzball_positions.json", &l.Positions, &l.PositionsID))
	checkErr(assignLookup("primers.json", &l.Primers, &l.PrimersID))
	checkErr(assignLookup("properties.json", &l.Properties, &l.PropertiesID))
	checkErr(assignLookup("quests.json", &l.Quests, &l.QuestsID))
	checkErr(assignLookup("ronso_rages.json", &l.RonsoRages, &l.RonsoRagesID))
	checkErr(assignLookup("sidequests.json", &l.Sidequests, &l.SidequestsID))
	checkErr(assignLookup("subquests.json", &l.Subquests, &l.SubquestsID))
	checkErr(assignLookup("shops.json", &l.Shops, &l.ShopsID))
	checkErr(assignLookup("songs.json", &l.Songs, &l.SongsID))
	checkErr(assignLookup("spheres.json", &l.Spheres, &l.SpheresID))
	checkErr(assignLookup("stats.json", &l.Stats, &l.StatsID))
	checkErr(assignLookup("status_conditions.json", &l.StatusConditions, &l.StatusConditionsID))
	checkErr(assignLookup("submenus.json", &l.Submenus, &l.SubmenusID))
	checkErr(assignLookup("topmenus.json", &l.Topmenus, &l.TopmenusID))
	checkErr(assignLookup("treasures.json", &l.Treasures, &l.TreasuresID))
	checkErr(l.assignHashes())

	return err
}

func assignLookup[T Lookupable](fileName string, keyLookup *map[string]T, idLookup *map[int32]T) error {
	lookupPath := filepath.Join("data_lookups", fileName)

	var dataSlice []T
	err := loadJsonFile(lookupPath, &dataSlice)
	if err != nil {
		return err
	}

	if keyLookup != nil {
		*keyLookup = sliceToKeyLookup(dataSlice)
	}

	if idLookup != nil {
		*idLookup = sliceToIdLookup(dataSlice)
	}

	return nil
}


func lookupToSlice[T Lookupable](lookup map[int32]T) []T {
	var s []T

	for _, obj := range lookup {
		s = append(s, obj)
	}

	slices.SortStableFunc(s, func(a, b T) int {
		if a.GetID() < b.GetID() {
			return -1
		}

		if a.GetID() > b.GetID() {
			return 1
		}

		return 0
	})

	return s
}

func sliceToKeyLookup[T Lookupable](slice []T) map[string]T {
	lookup := make(map[string]T, len(slice))

	for _, obj := range slice {
		lookup[Key(obj)] = obj
	}

	return lookup
}


func sliceToIdLookup[T Lookupable](slice []T) map[int32]T {
	lookup := make(map[int32]T, len(slice))

	for _, obj := range slice {
		lookup[obj.GetID()] = obj
	}

	return lookup
}