package seedhelpers

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

var RawJSON jsonFiles

type jsonFiles struct {
	AeonCommands          []seedtypes.AeonCommand
	AeonStats             []seedtypes.AeonStat
	Aeons                 []seedtypes.Aeon
	AgilityTiers          []seedtypes.AgilityTier
	AutoAbilities         []seedtypes.AutoAbility
	BlitzballPositions    []seedtypes.BlitzballPosition
	CelestialWeapons      []seedtypes.CelestialWeapon
	CharacterClasses      []seedtypes.CharacterClass
	Characters            []seedtypes.Character
	DefaultAbilities      []seedtypes.DefaultAbilitiesEntry
	Elements              []seedtypes.Element
	EnemyAbilities        []seedtypes.EnemyAbility
	Equipment             []seedtypes.EquipmentTable
	FMVs                  []seedtypes.FMV
	Items                 []seedtypes.Item
	KeyItems              []seedtypes.KeyItem
	Locations             []seedtypes.Location
	Mixes                 []seedtypes.Mix
	MiscAbilities         []seedtypes.MiscAbility
	Modifiers             []seedtypes.Modifier
	MonsterArenaCreations []seedtypes.ArenaCreation
	MonsterFormations     []seedtypes.MonsterFormation
	Monsters              []seedtypes.Monster
	OverdriveAbilities    []seedtypes.OverdriveAbility
	OverdriveCommands     []seedtypes.OverdriveCommand
	OverdriveModes        []seedtypes.OverdriveMode
	Overdrives            []seedtypes.Overdrive
	PlayerAbilities       []seedtypes.PlayerAbility
	Primers               []seedtypes.Primer
	Properties            []seedtypes.Property
	Shops                 []seedtypes.Shop
	Sidequests            []seedtypes.Sidequest
	Songs                 []seedtypes.Song
	Spheres               []seedtypes.Sphere
	Stats                 []seedtypes.Stat
	StatusConditions      []seedtypes.StatusCondition
	Submenus              []seedtypes.Submenu
	Topmenus              []seedtypes.Topmenu
	TreasureLists         []seedtypes.TreasureList
	TriggerCommands       []seedtypes.TriggerCommand
}

func (rawJSON *jsonFiles) LoadFiles() error {
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(h.LoadJSONFile("data/aeon_commands.json", &rawJSON.AeonCommands))
	checkErr(h.LoadJSONFile("data/aeon_stats.json", &rawJSON.AeonStats))
	checkErr(h.LoadJSONFile("data/aeons.json", &rawJSON.Aeons))
	checkErr(h.LoadJSONFile("data/agility_tiers.json", &rawJSON.AgilityTiers))
	checkErr(h.LoadJSONFile("data/auto_abilities.json", &rawJSON.AutoAbilities))
	checkErr(h.LoadJSONFile("data/blitzball_items.json", &rawJSON.BlitzballPositions))
	checkErr(h.LoadJSONFile("data/celestial_weapons.json", &rawJSON.CelestialWeapons))
	checkErr(h.LoadJSONFile("data/character_classes.json", &rawJSON.CharacterClasses))
	checkErr(h.LoadJSONFile("data/characters.json", &rawJSON.Characters))
	checkErr(h.LoadJSONFile("data/default_abilities.json", &rawJSON.DefaultAbilities))
	checkErr(h.LoadJSONFile("data/elements.json", &rawJSON.Elements))
	checkErr(h.LoadJSONFile("data/enemy_abilities.json", &rawJSON.EnemyAbilities))
	checkErr(h.LoadJSONFile("data/equipment.json", &rawJSON.Equipment))
	checkErr(h.LoadJSONFile("data/fmvs.json", &rawJSON.FMVs))
	checkErr(h.LoadJSONFile("data/items.json", &rawJSON.Items))
	checkErr(h.LoadJSONFile("data/key_items.json", &rawJSON.KeyItems))
	checkErr(h.LoadJSONFile("data/locations.json", &rawJSON.Locations))
	checkErr(h.LoadJSONFile("data/misc_abilities.json", &rawJSON.MiscAbilities))
	checkErr(h.LoadJSONFile("data/mixes.json", &rawJSON.Mixes))
	checkErr(h.LoadJSONFile("data/modifiers.json", &rawJSON.Modifiers))
	checkErr(h.LoadJSONFile("data/monster_arena_creations.json", &rawJSON.MonsterArenaCreations))
	checkErr(h.LoadJSONFile("data/monster_formations.json", &rawJSON.MonsterFormations))
	checkErr(h.LoadJSONFile("data/monsters.json", &rawJSON.Monsters))
	checkErr(h.LoadJSONFile("data/overdrive_abilities.json", &rawJSON.OverdriveAbilities))
	checkErr(h.LoadJSONFile("data/overdrive_commands.json", &rawJSON.OverdriveCommands))
	checkErr(h.LoadJSONFile("data/overdrive_modes.json", &rawJSON.OverdriveModes))
	checkErr(h.LoadJSONFile("data/overdrives.json", &rawJSON.Overdrives))
	checkErr(h.LoadJSONFile("data/player_abilities.json", &rawJSON.PlayerAbilities))
	checkErr(h.LoadJSONFile("data/primers.json", &rawJSON.Primers))
	checkErr(h.LoadJSONFile("data/properties.json", &rawJSON.Properties))
	checkErr(h.LoadJSONFile("data/shops.json", &rawJSON.Shops))
	checkErr(h.LoadJSONFile("data/sidequests.json", &rawJSON.Sidequests))
	checkErr(h.LoadJSONFile("data/songs.json", &rawJSON.Songs))
	checkErr(h.LoadJSONFile("data/spheres.json", &rawJSON.Spheres))
	checkErr(h.LoadJSONFile("data/stats.json", &rawJSON.Stats))
	checkErr(h.LoadJSONFile("data/status_conditions.json", &rawJSON.StatusConditions))
	checkErr(h.LoadJSONFile("data/submenus.json", &rawJSON.Submenus))
	checkErr(h.LoadJSONFile("data/topmenus.json", &rawJSON.Topmenus))
	checkErr(h.LoadJSONFile("data/treasures.json", &rawJSON.TreasureLists))
	checkErr(h.LoadJSONFile("data/trigger_commands.json", &rawJSON.TriggerCommands))

	return err
}
