package seeding

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

var rawJSON jsonFiles

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


func (rawJSON *jsonFiles) loadFiles() error {
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(h.LoadJSONFile("data/aeon_commands.json", &rawJSON.aeonCommands))
	checkErr(h.LoadJSONFile("data/aeon_stats.json", &rawJSON.aeonStats))
	checkErr(h.LoadJSONFile("data/aeons.json", &rawJSON.aeons))
	checkErr(h.LoadJSONFile("data/agility_tiers.json", &rawJSON.agilityTiers))
	checkErr(h.LoadJSONFile("data/auto_abilities.json", &rawJSON.autoAbilities))
	checkErr(h.LoadJSONFile("data/blitzball_items.json", &rawJSON.blitzballPositions))
	checkErr(h.LoadJSONFile("data/celestial_weapons.json", &rawJSON.celestialWeapons))
	checkErr(h.LoadJSONFile("data/character_classes.json", &rawJSON.characterClasses))
	checkErr(h.LoadJSONFile("data/characters.json", &rawJSON.characters))
	checkErr(h.LoadJSONFile("data/default_abilities.json", &rawJSON.defaultAbilities))
	checkErr(h.LoadJSONFile("data/elements.json", &rawJSON.elements))
	checkErr(h.LoadJSONFile("data/enemy_abilities.json", &rawJSON.enemyAbilities))
	checkErr(h.LoadJSONFile("data/equipment.json", &rawJSON.equipment))
	checkErr(h.LoadJSONFile("data/fmvs.json", &rawJSON.fmvs))
	checkErr(h.LoadJSONFile("data/items.json", &rawJSON.items))
	checkErr(h.LoadJSONFile("data/key_items.json", &rawJSON.keyItems))
	checkErr(h.LoadJSONFile("data/locations.json", &rawJSON.locations))
	checkErr(h.LoadJSONFile("data/misc_abilities.json", &rawJSON.miscAbilities))
	checkErr(h.LoadJSONFile("data/mixes.json", &rawJSON.mixes))
	checkErr(h.LoadJSONFile("data/modifiers.json", &rawJSON.modifiers))
	checkErr(h.LoadJSONFile("data/monster_arena_creations.json", &rawJSON.monsterArenaCreations))
	checkErr(h.LoadJSONFile("data/monster_formations.json", &rawJSON.monsterFormations))
	checkErr(h.LoadJSONFile("data/monsters.json", &rawJSON.monsters))
	checkErr(h.LoadJSONFile("data/overdrive_abilities.json", &rawJSON.overdriveAbilities))
	checkErr(h.LoadJSONFile("data/overdrive_commands.json", &rawJSON.overdriveCommands))
	checkErr(h.LoadJSONFile("data/overdrive_modes.json", &rawJSON.overdriveModes))
	checkErr(h.LoadJSONFile("data/overdrives.json", &rawJSON.overdrives))
	checkErr(h.LoadJSONFile("data/player_abilities.json", &rawJSON.playerAbilities))
	checkErr(h.LoadJSONFile("data/primers.json", &rawJSON.primers))
	checkErr(h.LoadJSONFile("data/properties.json", &rawJSON.properties))
	checkErr(h.LoadJSONFile("data/shops.json", &rawJSON.shops))
	checkErr(h.LoadJSONFile("data/sidequests.json", &rawJSON.sidequests))
	checkErr(h.LoadJSONFile("data/songs.json", &rawJSON.songs))
	checkErr(h.LoadJSONFile("data/spheres.json", &rawJSON.spheres))
	checkErr(h.LoadJSONFile("data/stats.json", &rawJSON.stats))
	checkErr(h.LoadJSONFile("data/status_conditions.json", &rawJSON.statusConditions))
	checkErr(h.LoadJSONFile("data/submenus.json", &rawJSON.submenus))
	checkErr(h.LoadJSONFile("data/topmenus.json", &rawJSON.topmenus))
	checkErr(h.LoadJSONFile("data/treasures.json", &rawJSON.treasureLists))
	checkErr(h.LoadJSONFile("data/trigger_commands.json", &rawJSON.triggerCommands))

	return err
}
