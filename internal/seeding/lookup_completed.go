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

	jsonBytes, err := json.MarshalIndent(slice, "", "  ")
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