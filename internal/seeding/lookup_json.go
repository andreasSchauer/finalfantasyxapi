package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type jsonLookup struct {
	aeonCommands          []AeonCommand
	aeonStats             []AeonStat
	aeons                 []Aeon
	agilityTiers          []AgilityTier
	autoAbilities         []AutoAbility
	blitzballPositions    []BlitzballPosition
	celestialWeapons      []CelestialWeapon
	characterClasses      []CharacterClass
	characters            []Character
	defaultAbilities      []DefaultAbilitiesEntry
	elements              []Element
	enemyAbilities        []EnemyAbility
	equipment             []EquipmentTable
	fmvs                  []FMV
	items                 []Item
	keyItems              []KeyItem
	locations             []Location
	mixes                 []Mix
	modifiers             []Modifier
	monsterArenaCreations []ArenaCreation
	monsterFormations     []MonsterFormation
	monsters              []Monster
	overdriveAbilities    []OverdriveAbility
	overdriveCommands     []OverdriveCommand
	overdriveModes        []OverdriveMode
	overdrives            []Overdrive
	playerAbilities       []PlayerAbility
	primers               []Primer
	properties            []Property
	shops                 []Shop
	sidequests            []Sidequest
	songs                 []Song
	spheres               []Sphere
	stats                 []Stat
	statusConditions      []StatusCondition
	submenus              []Submenu
	topmenus              []Topmenu
	treasureLists         []TreasureList
	triggerCommands       []TriggerCommand
	miscAbilities         []MiscAbility
}

func loadJsonFile[T any](path string, target *T) error {
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

func (l *Lookup) loadJsonFiles() error {
	l.json = jsonLookup{}
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(loadJsonFile("data/aeon_commands.json", &l.json.aeonCommands))
	checkErr(loadJsonFile("data/aeon_stats.json", &l.json.aeonStats))
	checkErr(loadJsonFile("data/aeons.json", &l.json.aeons))
	checkErr(loadJsonFile("data/agility_tiers.json", &l.json.agilityTiers))
	checkErr(loadJsonFile("data/auto_abilities.json", &l.json.autoAbilities))
	checkErr(loadJsonFile("data/blitzball_items.json", &l.json.blitzballPositions))
	checkErr(loadJsonFile("data/celestial_weapons.json", &l.json.celestialWeapons))
	checkErr(loadJsonFile("data/character_classes.json", &l.json.characterClasses))
	checkErr(loadJsonFile("data/characters.json", &l.json.characters))
	checkErr(loadJsonFile("data/default_abilities.json", &l.json.defaultAbilities))
	checkErr(loadJsonFile("data/elements.json", &l.json.elements))
	checkErr(loadJsonFile("data/enemy_abilities.json", &l.json.enemyAbilities))
	checkErr(loadJsonFile("data/equipment.json", &l.json.equipment))
	checkErr(loadJsonFile("data/fmvs.json", &l.json.fmvs))
	checkErr(loadJsonFile("data/items.json", &l.json.items))
	checkErr(loadJsonFile("data/key_items.json", &l.json.keyItems))
	checkErr(loadJsonFile("data/locations.json", &l.json.locations))
	checkErr(loadJsonFile("data/misc_abilities.json", &l.json.miscAbilities))
	checkErr(loadJsonFile("data/mixes.json", &l.json.mixes))
	checkErr(loadJsonFile("data/modifiers.json", &l.json.modifiers))
	checkErr(loadJsonFile("data/monster_arena_creations.json", &l.json.monsterArenaCreations))
	checkErr(loadJsonFile("data/monster_formations.json", &l.json.monsterFormations))
	checkErr(loadJsonFile("data/monsters.json", &l.json.monsters))
	checkErr(loadJsonFile("data/overdrive_abilities.json", &l.json.overdriveAbilities))
	checkErr(loadJsonFile("data/overdrive_commands.json", &l.json.overdriveCommands))
	checkErr(loadJsonFile("data/overdrive_modes.json", &l.json.overdriveModes))
	checkErr(loadJsonFile("data/overdrives.json", &l.json.overdrives))
	checkErr(loadJsonFile("data/player_abilities.json", &l.json.playerAbilities))
	checkErr(loadJsonFile("data/primers.json", &l.json.primers))
	checkErr(loadJsonFile("data/properties.json", &l.json.properties))
	checkErr(loadJsonFile("data/shops.json", &l.json.shops))
	checkErr(loadJsonFile("data/sidequests.json", &l.json.sidequests))
	checkErr(loadJsonFile("data/songs.json", &l.json.songs))
	checkErr(loadJsonFile("data/spheres.json", &l.json.spheres))
	checkErr(loadJsonFile("data/stats.json", &l.json.stats))
	checkErr(loadJsonFile("data/status_conditions.json", &l.json.statusConditions))
	checkErr(loadJsonFile("data/submenus.json", &l.json.submenus))
	checkErr(loadJsonFile("data/topmenus.json", &l.json.topmenus))
	checkErr(loadJsonFile("data/treasures.json", &l.json.treasureLists))
	checkErr(loadJsonFile("data/trigger_commands.json", &l.json.triggerCommands))

	return err
}
