package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type jsonFiles struct {
	aeonCommands          []seedtypes.AeonCommand
	aeonStats             []seedtypes.AeonStat
	aeons                 []seedtypes.Aeon
	agilityTiers          []seedtypes.AgilityTier
	autoAbilities         []seedtypes.AutoAbility
	blitzballPositions    []seedtypes.BlitzballPosition
	celestialWeapons      []seedtypes.CelestialWeapon
	characterClasses      []seedtypes.CharacterClass
	characters            []seedtypes.Character
	defaultAbilities      []DefaultAbilitiesEntry
	elements              []seedtypes.Element
	enemyAbilities        []seedtypes.EnemyAbility
	equipment             []seedtypes.EquipmentTable
	fmvs                  []seedtypes.FMV
	items                 []seedtypes.Item
	keyItems              []seedtypes.KeyItem
	locations             []seedtypes.Location
	mixes                 []seedtypes.Mix
	modifiers             []seedtypes.Modifier
	monsterArenaCreations []seedtypes.ArenaCreation
	monsterFormations     []seedtypes.MonsterFormation
	monsters              []seedtypes.Monster
	overdriveAbilities    []seedtypes.OverdriveAbility
	overdriveCommands     []seedtypes.OverdriveCommand
	overdriveModes        []seedtypes.OverdriveMode
	overdrives            []seedtypes.Overdrive
	playerAbilities       []seedtypes.PlayerAbility
	primers               []seedtypes.Primer
	properties            []seedtypes.Property
	shops                 []seedtypes.Shop
	sidequests            []seedtypes.Sidequest
	songs                 []seedtypes.Song
	spheres               []seedtypes.Sphere
	stats                 []seedtypes.Stat
	statusConditions      []seedtypes.StatusCondition
	submenus              []seedtypes.Submenu
	topmenus              []seedtypes.Topmenu
	treasureLists         []seedtypes.TreasureList
	triggerCommands       []seedtypes.TriggerCommand
	miscAbilities         []seedtypes.MiscAbility
}

func loadJSONFile[T any](path string, target *[]T) error {
	fullPath, err := h.GetAbsoluteFilepath(path)
	if err != nil {
		return err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("couldn't parse JSON: %v", err)
	}

	return nil
}

func (rawJSON *jsonFiles) loadFiles() error {
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(loadJSONFile("data/aeon_commands.json", &rawJSON.aeonCommands))
	checkErr(loadJSONFile("data/aeon_stats.json", &rawJSON.aeonStats))
	checkErr(loadJSONFile("data/aeons.json", &rawJSON.aeons))
	checkErr(loadJSONFile("data/agility_tiers.json", &rawJSON.agilityTiers))
	checkErr(loadJSONFile("data/auto_abilities.json", &rawJSON.autoAbilities))
	checkErr(loadJSONFile("data/blitzball_items.json", &rawJSON.blitzballPositions))
	checkErr(loadJSONFile("data/celestial_weapons.json", &rawJSON.celestialWeapons))
	checkErr(loadJSONFile("data/character_classes.json", &rawJSON.characterClasses))
	checkErr(loadJSONFile("data/characters.json", &rawJSON.characters))
	checkErr(loadJSONFile("data/default_abilities.json", &rawJSON.defaultAbilities))
	checkErr(loadJSONFile("data/elements.json", &rawJSON.elements))
	checkErr(loadJSONFile("data/enemy_abilities.json", &rawJSON.enemyAbilities))
	checkErr(loadJSONFile("data/equipment.json", &rawJSON.equipment))
	checkErr(loadJSONFile("data/fmvs.json", &rawJSON.fmvs))
	checkErr(loadJSONFile("data/items.json", &rawJSON.items))
	checkErr(loadJSONFile("data/key_items.json", &rawJSON.keyItems))
	checkErr(loadJSONFile("data/locations.json", &rawJSON.locations))
	checkErr(loadJSONFile("data/misc_abilities.json", &rawJSON.miscAbilities))
	checkErr(loadJSONFile("data/mixes.json", &rawJSON.mixes))
	checkErr(loadJSONFile("data/modifiers.json", &rawJSON.modifiers))
	checkErr(loadJSONFile("data/monster_arena_creations.json", &rawJSON.monsterArenaCreations))
	checkErr(loadJSONFile("data/monster_formations.json", &rawJSON.monsterFormations))
	checkErr(loadJSONFile("data/monsters.json", &rawJSON.monsters))
	checkErr(loadJSONFile("data/overdrive_abilities.json", &rawJSON.overdriveAbilities))
	checkErr(loadJSONFile("data/overdrive_commands.json", &rawJSON.overdriveCommands))
	checkErr(loadJSONFile("data/overdrive_modes.json", &rawJSON.overdriveModes))
	checkErr(loadJSONFile("data/overdrives.json", &rawJSON.overdrives))
	checkErr(loadJSONFile("data/player_abilities.json", &rawJSON.playerAbilities))
	checkErr(loadJSONFile("data/primers.json", &rawJSON.primers))
	checkErr(loadJSONFile("data/properties.json", &rawJSON.properties))
	checkErr(loadJSONFile("data/shops.json", &rawJSON.shops))
	checkErr(loadJSONFile("data/sidequests.json", &rawJSON.sidequests))
	checkErr(loadJSONFile("data/songs.json", &rawJSON.songs))
	checkErr(loadJSONFile("data/spheres.json", &rawJSON.spheres))
	checkErr(loadJSONFile("data/stats.json", &rawJSON.stats))
	checkErr(loadJSONFile("data/status_conditions.json", &rawJSON.statusConditions))
	checkErr(loadJSONFile("data/submenus.json", &rawJSON.submenus))
	checkErr(loadJSONFile("data/topmenus.json", &rawJSON.topmenus))
	checkErr(loadJSONFile("data/treasures.json", &rawJSON.treasureLists))
	checkErr(loadJSONFile("data/trigger_commands.json", &rawJSON.triggerCommands))

	return err
}
