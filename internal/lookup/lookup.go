// Package lookup is fresh
package lookup

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

var L Lookup

type Lookup struct {
	Abilities            map[string]seedtypes.Ability
	AbilitiesID          map[int32]seedtypes.Ability
	EnemyAbilities       map[string]seedtypes.EnemyAbility
	EnemyAbilitiesID     map[int32]seedtypes.EnemyAbility
	ItemAbilities        map[string]seedtypes.ItemAbility
	ItemAbilitiesID      map[int32]seedtypes.ItemAbility
	MiscAbilities        map[string]seedtypes.MiscAbility
	MiscAbilitiesID      map[int32]seedtypes.MiscAbility
	OverdriveAbilities   map[string]seedtypes.OverdriveAbility
	OverdriveAbilitiesID map[int32]seedtypes.OverdriveAbility
	PlayerAbilities      map[string]seedtypes.PlayerAbility
	PlayerAbilitiesID    map[int32]seedtypes.PlayerAbility
	TriggerCommands      map[string]seedtypes.TriggerCommand
	TriggerCommandsID    map[int32]seedtypes.TriggerCommand
	PlayerUnits          map[string]seedtypes.PlayerUnit
	PlayerUnitsID        map[int32]seedtypes.PlayerUnit
	Aeons                map[string]seedtypes.Aeon
	AeonsID              map[int32]seedtypes.Aeon
	AeonCommands         map[string]seedtypes.AeonCommand
	AeonCommandsID       map[int32]seedtypes.AeonCommand
	AgilityTiersID       map[int32]seedtypes.AgilityTier
	ArenaCreations       map[string]seedtypes.ArenaCreation
	ArenaCreationsID     map[int32]seedtypes.ArenaCreation
	Locations            map[string]seedtypes.Location
	LocationsID          map[int32]seedtypes.Location
	Sublocations         map[string]seedtypes.Sublocation
	SublocationsID       map[int32]seedtypes.Sublocation
	Areas                map[string]seedtypes.Area
	AreasID              map[int32]seedtypes.Area
	AutoAbilities        map[string]seedtypes.AutoAbility
	AutoAbilitiesID      map[int32]seedtypes.AutoAbility
	CelestialWeapons     map[string]seedtypes.CelestialWeapon
	CelestialWeaponsID   map[int32]seedtypes.CelestialWeapon
	Characters           map[string]seedtypes.Character
	CharactersID         map[int32]seedtypes.Character
	CharClasses          map[string]seedtypes.CharacterClass
	CharClassesID        map[int32]seedtypes.CharacterClass
	Elements             map[string]seedtypes.Element
	ElementsID           map[int32]seedtypes.Element
	ElementalResists     map[string]seedtypes.ElementalResist
	ElementalResistsID   map[int32]seedtypes.ElementalResist
	EquipmentNames       map[string]seedtypes.EquipmentName
	EquipmentNamesID     map[int32]seedtypes.EquipmentName
	EquipmentTables      map[string]seedtypes.EquipmentTable
	EquipmentTablesID    map[int32]seedtypes.EquipmentTable
	EncounterAreas       map[string]seedtypes.EncounterArea
	EncounterAreasID     map[int32]seedtypes.EncounterArea
	FMVs                 map[string]seedtypes.FMV
	FMVsID               map[int32]seedtypes.FMV
	Items                map[string]seedtypes.Item
	ItemsID              map[int32]seedtypes.Item
	KeyItems             map[string]seedtypes.KeyItem
	KeyItemsID           map[int32]seedtypes.KeyItem
	MasterItems          map[string]seedtypes.MasterItem
	MasterItemsID        map[int32]seedtypes.MasterItem
	Mixes                map[string]seedtypes.Mix
	MixesID              map[int32]seedtypes.Mix
	Modifiers            map[string]seedtypes.Modifier
	ModifiersID          map[int32]seedtypes.Modifier
	Monsters             map[string]seedtypes.Monster
	MonstersID           map[int32]seedtypes.Monster
	MonsterFormations    map[string]seedtypes.MonsterFormation
	MonsterFormationsID  map[int32]seedtypes.MonsterFormation
	OverdriveModes       map[string]seedtypes.OverdriveMode
	OverdriveModesID     map[int32]seedtypes.OverdriveMode
	OverdriveCommands    map[string]seedtypes.OverdriveCommand
	OverdriveCommandsID  map[int32]seedtypes.OverdriveCommand
	Overdrives           map[string]seedtypes.Overdrive
	OverdrivesID         map[int32]seedtypes.Overdrive
	Positions            map[string]seedtypes.BlitzballPosition
	PositionsID          map[int32]seedtypes.BlitzballPosition
	Primers              map[string]seedtypes.Primer
	PrimersID            map[int32]seedtypes.Primer
	Properties           map[string]seedtypes.Property
	PropertiesID         map[int32]seedtypes.Property
	Quests               map[string]seedtypes.Quest
	QuestsID             map[int32]seedtypes.Quest
	RonsoRages           map[string]seedtypes.RonsoRage
	RonsoRagesID         map[int32]seedtypes.RonsoRage
	Sidequests           map[string]seedtypes.Sidequest
	SidequestsID         map[int32]seedtypes.Sidequest
	Subquests            map[string]seedtypes.Subquest
	SubquestsID          map[int32]seedtypes.Subquest
	Shops                map[string]seedtypes.Shop
	ShopsID              map[int32]seedtypes.Shop
	Songs                map[string]seedtypes.Song
	SongsID              map[int32]seedtypes.Song
	Spheres              map[string]seedtypes.Sphere
	SpheresID            map[int32]seedtypes.Sphere
	Stats                map[string]seedtypes.Stat
	StatsID              map[int32]seedtypes.Stat
	StatusConditions     map[string]seedtypes.StatusCondition
	StatusConditionsID   map[int32]seedtypes.StatusCondition
	Submenus             map[string]seedtypes.Submenu
	SubmenusID           map[int32]seedtypes.Submenu
	Topmenus             map[string]seedtypes.Topmenu
	TopmenusID           map[int32]seedtypes.Topmenu
	Treasures            map[string]seedtypes.Treasure
	TreasuresID          map[int32]seedtypes.Treasure
}

func (L *Lookup) Init() {
	defer h.MeasureTime("creating lookups")()

	L.Abilities = make(map[string]seedtypes.Ability)
	L.AbilitiesID = make(map[int32]seedtypes.Ability)
	L.MiscAbilities = make(map[string]seedtypes.MiscAbility)
	L.MiscAbilitiesID = make(map[int32]seedtypes.MiscAbility)
	L.EnemyAbilities = make(map[string]seedtypes.EnemyAbility)
	L.EnemyAbilitiesID = make(map[int32]seedtypes.EnemyAbility)
	L.ItemAbilities = make(map[string]seedtypes.ItemAbility)
	L.ItemAbilitiesID = make(map[int32]seedtypes.ItemAbility)
	L.OverdriveAbilities = make(map[string]seedtypes.OverdriveAbility)
	L.OverdriveAbilitiesID = make(map[int32]seedtypes.OverdriveAbility)
	L.PlayerAbilities = make(map[string]seedtypes.PlayerAbility)
	L.PlayerAbilitiesID = make(map[int32]seedtypes.PlayerAbility)
	L.TriggerCommands = make(map[string]seedtypes.TriggerCommand)
	L.TriggerCommandsID = make(map[int32]seedtypes.TriggerCommand)
	L.PlayerUnits = make(map[string]seedtypes.PlayerUnit)
	L.PlayerUnitsID = make(map[int32]seedtypes.PlayerUnit)
	L.Aeons = make(map[string]seedtypes.Aeon)
	L.AeonsID = make(map[int32]seedtypes.Aeon)
	L.AeonCommands = make(map[string]seedtypes.AeonCommand)
	L.AeonCommandsID = make(map[int32]seedtypes.AeonCommand)
	L.AgilityTiersID = make(map[int32]seedtypes.AgilityTier)
	L.ArenaCreations = make(map[string]seedtypes.ArenaCreation)
	L.ArenaCreationsID = make(map[int32]seedtypes.ArenaCreation)
	L.Locations = make(map[string]seedtypes.Location)
	L.LocationsID = make(map[int32]seedtypes.Location)
	L.Sublocations = make(map[string]seedtypes.Sublocation)
	L.SublocationsID = make(map[int32]seedtypes.Sublocation)
	L.Areas = make(map[string]seedtypes.Area)
	L.AreasID = make(map[int32]seedtypes.Area)
	L.AutoAbilities = make(map[string]seedtypes.AutoAbility)
	L.AutoAbilitiesID = make(map[int32]seedtypes.AutoAbility)
	L.CelestialWeapons = make(map[string]seedtypes.CelestialWeapon)
	L.CelestialWeaponsID = make(map[int32]seedtypes.CelestialWeapon)
	L.Characters = make(map[string]seedtypes.Character)
	L.CharactersID = make(map[int32]seedtypes.Character)
	L.CharClasses = make(map[string]seedtypes.CharacterClass)
	L.CharClassesID = make(map[int32]seedtypes.CharacterClass)
	L.Elements = make(map[string]seedtypes.Element)
	L.ElementsID = make(map[int32]seedtypes.Element)
	L.ElementalResists = make(map[string]seedtypes.ElementalResist)
	L.ElementalResistsID = make(map[int32]seedtypes.ElementalResist)
	L.EquipmentNames = make(map[string]seedtypes.EquipmentName)
	L.EquipmentNamesID = make(map[int32]seedtypes.EquipmentName)
	L.EquipmentTables = make(map[string]seedtypes.EquipmentTable)
	L.EquipmentTablesID = make(map[int32]seedtypes.EquipmentTable)
	L.EncounterAreas = make(map[string]seedtypes.EncounterArea)
	L.EncounterAreasID = make(map[int32]seedtypes.EncounterArea)
	L.FMVs = make(map[string]seedtypes.FMV)
	L.FMVsID = make(map[int32]seedtypes.FMV)
	L.Items = make(map[string]seedtypes.Item)
	L.ItemsID = make(map[int32]seedtypes.Item)
	L.KeyItems = make(map[string]seedtypes.KeyItem)
	L.KeyItemsID = make(map[int32]seedtypes.KeyItem)
	L.MasterItems = make(map[string]seedtypes.MasterItem)
	L.MasterItemsID = make(map[int32]seedtypes.MasterItem)
	L.Mixes = make(map[string]seedtypes.Mix)
	L.MixesID = make(map[int32]seedtypes.Mix)
	L.Modifiers = make(map[string]seedtypes.Modifier)
	L.ModifiersID = make(map[int32]seedtypes.Modifier)
	L.Monsters = make(map[string]seedtypes.Monster)
	L.MonstersID = make(map[int32]seedtypes.Monster)
	L.MonsterFormations = make(map[string]seedtypes.MonsterFormation)
	L.MonsterFormationsID = make(map[int32]seedtypes.MonsterFormation)
	L.OverdriveModes = make(map[string]seedtypes.OverdriveMode)
	L.OverdriveModesID = make(map[int32]seedtypes.OverdriveMode)
	L.OverdriveCommands = make(map[string]seedtypes.OverdriveCommand)
	L.OverdriveCommandsID = make(map[int32]seedtypes.OverdriveCommand)
	L.Overdrives = make(map[string]seedtypes.Overdrive)
	L.OverdrivesID = make(map[int32]seedtypes.Overdrive)
	L.Positions = make(map[string]seedtypes.BlitzballPosition)
	L.PositionsID = make(map[int32]seedtypes.BlitzballPosition)
	L.Primers = make(map[string]seedtypes.Primer)
	L.PrimersID = make(map[int32]seedtypes.Primer)
	L.Properties = make(map[string]seedtypes.Property)
	L.PropertiesID = make(map[int32]seedtypes.Property)
	L.Quests = make(map[string]seedtypes.Quest)
	L.QuestsID = make(map[int32]seedtypes.Quest)
	L.RonsoRages = make(map[string]seedtypes.RonsoRage)
	L.RonsoRagesID = make(map[int32]seedtypes.RonsoRage)
	L.Sidequests = make(map[string]seedtypes.Sidequest)
	L.SidequestsID = make(map[int32]seedtypes.Sidequest)
	L.Subquests = make(map[string]seedtypes.Subquest)
	L.SubquestsID = make(map[int32]seedtypes.Subquest)
	L.Shops = make(map[string]seedtypes.Shop)
	L.ShopsID = make(map[int32]seedtypes.Shop)
	L.Songs = make(map[string]seedtypes.Song)
	L.SongsID = make(map[int32]seedtypes.Song)
	L.Spheres = make(map[string]seedtypes.Sphere)
	L.SpheresID = make(map[int32]seedtypes.Sphere)
	L.Stats = make(map[string]seedtypes.Stat)
	L.StatsID = make(map[int32]seedtypes.Stat)
	L.StatusConditions = make(map[string]seedtypes.StatusCondition)
	L.StatusConditionsID = make(map[int32]seedtypes.StatusCondition)
	L.Submenus = make(map[string]seedtypes.Submenu)
	L.SubmenusID = make(map[int32]seedtypes.Submenu)
	L.Topmenus = make(map[string]seedtypes.Topmenu)
	L.TopmenusID = make(map[int32]seedtypes.Topmenu)
	L.Treasures = make(map[string]seedtypes.Treasure)
	L.TreasuresID = make(map[int32]seedtypes.Treasure)
}
