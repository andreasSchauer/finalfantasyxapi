package seeding

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type jsonLookup struct {
	aeonCommands			[]AeonCommand
	aeonStats				[]AeonStat
	aeons					[]Aeon
	agilityTiers 			[]AgilityTier
	autoAbilities			[]AutoAbility
	blitzballPositions		[]BlitzballPosition
	celestialWeapons		[]CelestialWeapon
	characterClasses		[]CharacterClass
	characters				[]Character
	defaultAbilities		[]DefaultAbilitiesEntry
	elements	 			[]Element
	enemyAbilities			[]EnemyAbility
	equipment				[]EquipmentTable
	fmvs					[]FMV
	items					[]Item
	keyItems				[]KeyItem
	locations				[]Location
	mixes					[]Mix
	modifiers				[]Modifier
	monsterArenaCreations	[]ArenaCreation
	monsterFormations		[]MonsterFormation
	monsters				[]Monster
	overdriveAbilities		[]OverdriveAbility
	overdriveCommands		[]OverdriveCommand
	overdriveModes			[]OverdriveMode
	overdrives				[]Overdrive
	playerAbilities			[]PlayerAbility
	primers					[]Primer
	properties				[]Property
	shops					[]Shop
	sidequests				[]Sidequest
	songs					[]Song
	spheres					[]Sphere
	stats					[]Stat
	statusConditions		[]StatusCondition
	submenus				[]Submenu
	topmenus				[]Topmenu
	treasures				[]Treasure
	triggerCommands			[]TriggerCommand
	unspecifiedAbilities	[]UnspecifiedAbility
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
	checkErr(loadJSONFile("data/treasures.json", &l.json.treasures))
	checkErr(loadJSONFile("data/trigger_commands.json", &l.json.triggerCommands))
	checkErr(loadJSONFile("data/unspecified_abilities.json", &l.json.unspecifiedAbilities))

	return err
}