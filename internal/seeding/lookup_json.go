package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type jsonLookup struct {
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

func loadJSONFile[T any](path string, target *T) error {
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

func (l *Lookup) loadJSONFiles() error {
	l.json = jsonLookup{}
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(loadJSONFile("data/aeon_commands.json", &l.json.aeonCommands))
	checkErr(loadJSONFile("data/aeon_stats.json", &l.json.aeonStats))
	checkErr(loadJSONFile("data/aeons.json", &l.json.aeons))
	checkErr(loadJSONFile("data/agility_tiers.json", &l.json.agilityTiers))
	checkErr(loadJSONFile("data/auto_abilities.json", &l.json.autoAbilities))
	checkErr(loadJSONFile("data/blitzball_items.json", &l.json.blitzballPositions))
	checkErr(loadJSONFile("data/celestial_weapons.json", &l.json.celestialWeapons))
	checkErr(loadJSONFile("data/character_classes.json", &l.json.characterClasses))
	checkErr(loadJSONFile("data/characters.json", &l.json.characters))
	checkErr(loadJSONFile("data/default_abilities.json", &l.json.defaultAbilities))
	checkErr(loadJSONFile("data/elements.json", &l.json.elements))
	checkErr(loadJSONFile("data/enemy_abilities.json", &l.json.enemyAbilities))
	checkErr(loadJSONFile("data/equipment.json", &l.json.equipment))
	checkErr(loadJSONFile("data/fmvs.json", &l.json.fmvs))
	checkErr(loadJSONFile("data/items.json", &l.json.items))
	checkErr(loadJSONFile("data/key_items.json", &l.json.keyItems))
	checkErr(loadJSONFile("data/locations.json", &l.json.locations))
	checkErr(loadJSONFile("data/misc_abilities.json", &l.json.miscAbilities))
	checkErr(loadJSONFile("data/mixes.json", &l.json.mixes))
	checkErr(loadJSONFile("data/modifiers.json", &l.json.modifiers))
	checkErr(loadJSONFile("data/monster_arena_creations.json", &l.json.monsterArenaCreations))
	checkErr(loadJSONFile("data/monster_formations.json", &l.json.monsterFormations))
	checkErr(loadJSONFile("data/monsters.json", &l.json.monsters))
	checkErr(loadJSONFile("data/overdrive_abilities.json", &l.json.overdriveAbilities))
	checkErr(loadJSONFile("data/overdrive_commands.json", &l.json.overdriveCommands))
	checkErr(loadJSONFile("data/overdrive_modes.json", &l.json.overdriveModes))
	checkErr(loadJSONFile("data/overdrives.json", &l.json.overdrives))
	checkErr(loadJSONFile("data/player_abilities.json", &l.json.playerAbilities))
	checkErr(loadJSONFile("data/primers.json", &l.json.primers))
	checkErr(loadJSONFile("data/properties.json", &l.json.properties))
	checkErr(loadJSONFile("data/shops.json", &l.json.shops))
	checkErr(loadJSONFile("data/sidequests.json", &l.json.sidequests))
	checkErr(loadJSONFile("data/songs.json", &l.json.songs))
	checkErr(loadJSONFile("data/spheres.json", &l.json.spheres))
	checkErr(loadJSONFile("data/stats.json", &l.json.stats))
	checkErr(loadJSONFile("data/status_conditions.json", &l.json.statusConditions))
	checkErr(loadJSONFile("data/submenus.json", &l.json.submenus))
	checkErr(loadJSONFile("data/topmenus.json", &l.json.topmenus))
	checkErr(loadJSONFile("data/treasures.json", &l.json.treasureLists))
	checkErr(loadJSONFile("data/trigger_commands.json", &l.json.triggerCommands))

	return err
}
