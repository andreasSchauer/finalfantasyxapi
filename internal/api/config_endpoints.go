package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	aeonCommands         handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]
	aeons                handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]
	affinities           handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]
	arenaCreations       handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]
	areas                handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]
	autoAbilities        handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]
	blitzballPrizes      handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList]
	celestialWeapons     handlerInput[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList]
	characters           handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]
	characterClasses     handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]
	elements             handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]
	equipment            handlerInput[seeding.EquipmentName, any, NamedAPIResource, NamedApiResourceList]
	fmvs                 handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]
	items                handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]
	keyItems             handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]
	locations            handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]
	modifiers            handlerInput[seeding.Modifier, any, NamedAPIResource, NamedApiResourceList]
	monsters             handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]
	monsterFormations    handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]
	overdriveCommands    handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]
	overdriveModes       handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]
	overdrives           handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]
	abilities            handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList]
	unspecifiedAbilities handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]
	playerAbilities      handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]
	enemyAbilities       handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]
	itemAbilities        handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities   handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]
	triggerCommands      handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]
	properties           handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]
	ronsoRages           handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]
	shops                handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]
	sidequests           handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]
	subquests            handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]
	songs                handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]
	stats                handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]
	statusConditions     handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]
	sublocations         handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]
	submenus             handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]
	topmenus             handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]
	treasures            handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]

	abilityType              handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	attackType               handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	damageFormula            handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	damageType               handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	itemCategory             handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterCategory          handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	lootType                 handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterFormationCategory handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterSpecies           handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	playerAbilityCategory    handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	shopCategory             handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeonCommands = handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeon-commands",
		resourceType:  "aeon command",
		objLookup:     cfg.l.AeonCommands,
		objLookupID:   cfg.l.AeonCommandsID,
		queryLookup:   cfg.q.aeonCommands,
		idToResFunc:   idToNamedAPIResource[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonCommandIDs,
		getSingleFunc: cfg.getAeonCommand,
		retrieveFunc:  cfg.retrieveAeonCommands,
	}

	e.aeons = handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeons",
		resourceType:  "aeon",
		objLookup:     cfg.l.Aeons,
		objLookupID:   cfg.l.AeonsID,
		queryLookup:   cfg.q.aeons,
		idToResFunc:   idToNamedAPIResource[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[string]SubSectionFns{
			"default-abilities": {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			"stats": {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	e.affinities = handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "affinities",
		resourceType:  "affinity",
		objLookup:     cfg.l.Affinities,
		objLookupID:   cfg.l.AffinitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.arenaCreations = handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "arena-creations",
		resourceType:  "arena creation",
		objLookup:     cfg.l.ArenaCreations,
		objLookupID:   cfg.l.ArenaCreationsID,
		queryLookup:   cfg.q.arenaCreations,
		idToResFunc:   idToNamedAPIResource[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetArenaCreationIDs,
		getSingleFunc: cfg.getArenaCreation,
		retrieveFunc:  cfg.retrieveArenaCreations,
	}

	e.areas = handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]{
		endpoint:      "areas",
		resourceType:  "area",
		objLookup:     cfg.l.Areas,
		objLookupID:   cfg.l.AreasID,
		queryLookup:   cfg.q.areas,
		idToResFunc:   idToAreaAPIResource,
		resToListFunc: newAreaAPIResourceList,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createAreaSimple,
			},
			"connected": {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
			},
			"songs": {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.autoAbilities = handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "auto-abilities",
		resourceType:  "auto ability",
		objLookup:     cfg.l.AutoAbilities,
		objLookupID:   cfg.l.AutoAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.blitzballPrizes = handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "blitzball-prizes",
		resourceType:  "blitzball prize table",
		objLookup:     cfg.l.Positions,
		objLookupID:   cfg.l.PositionsID,
		queryLookup:   cfg.q.blitzballPrizes,
		idToResFunc:   idToUnnamedAPIResource[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetBlitzballPrizeIDs,
		getSingleFunc: cfg.getBlitzballPrize,
		retrieveFunc:  cfg.retrieveBlitzballPrizes,
	}

	e.celestialWeapons = handlerInput[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "celestial-weapons",
		resourceType:  "celestial weapon",
		objLookup:     cfg.l.CelestialWeapons,
		objLookupID:   cfg.l.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.characters = handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "characters",
		resourceType:  "character",
		objLookup:     cfg.l.Characters,
		objLookupID:   cfg.l.CharactersID,
		queryLookup:   cfg.q.characters,
		idToResFunc:   idToNamedAPIResource[seeding.Character, Character, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[string]SubSectionFns{
			"default-abilities": {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"std-sg-abilities": {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"exp-sg-abilities": {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.characterClasses = handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "character-classes",
		resourceType:  "character class",
		objLookup:     cfg.l.CharClasses,
		objLookupID:   cfg.l.CharClassesID,
		queryLookup:   cfg.q.characterClasses,
		idToResFunc:   idToNamedAPIResource[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[string]SubSectionFns{
			"default-abilities": {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			"learnable-abilities": {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			"default-overdrives": {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			"learnable-overdrives": {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.elements = handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "elements",
		resourceType:  "element",
		objLookup:     cfg.l.Elements,
		objLookupID:   cfg.l.ElementsID,
		idToResFunc:   idToNamedAPIResource[seeding.Element, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.equipment = handlerInput[seeding.EquipmentName, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "equipment",
		resourceType:  "equipment",
		objLookup:     cfg.l.EquipmentNames,
		objLookupID:   cfg.l.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seeding.EquipmentName, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.fmvs = handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "fmvs",
		resourceType:  "fmv",
		objLookup:     cfg.l.FMVs,
		objLookupID:   cfg.l.FMVsID,
		queryLookup:   cfg.q.fmvs,
		idToResFunc:   idToNamedAPIResource[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetFmvIDs,
		getSingleFunc: cfg.getFMV,
		retrieveFunc:  cfg.retrieveFMVs,
	}

	e.items = handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "items",
		resourceType:  "item",
		objLookup:     cfg.l.Items,
		objLookupID:   cfg.l.ItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.Item, Item, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.keyItems = handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "key-items",
		resourceType:  "key item",
		objLookup:     cfg.l.KeyItems,
		objLookupID:   cfg.l.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.locations = handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "locations",
		resourceType:  "location",
		objLookup:     cfg.l.Locations,
		objLookupID:   cfg.l.LocationsID,
		queryLookup:   cfg.q.locations,
		idToResFunc:   idToNamedAPIResource[seeding.Location, Location, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createLocationSimple,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
			},
			"sublocations": {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
			},
			"areas": {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
			},
			"shops": {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			"songs": {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.modifiers = handlerInput[seeding.Modifier, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "modifiers",
		resourceType:  "modifier",
		objLookup:     cfg.l.Modifiers,
		objLookupID:   cfg.l.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seeding.Modifier, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "monsters",
		resourceType:     "monster",
		objLookup:        cfg.l.Monsters,
		objLookupID:      cfg.l.MonstersID,
		queryLookup:      cfg.q.monsters,
		idToResFunc:      idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createMonsterSimple,
			},
			"abilities": {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			"areas": {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "monster-formations",
		resourceType:  "monster formation",
		objLookup:     cfg.l.MonsterFormations,
		objLookupID:   cfg.l.MonsterFormationsID,
		queryLookup:   cfg.q.monsterFormations,
		idToResFunc:   idToUnnamedAPIResource[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
			},
		},
	}

	e.overdriveCommands = handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrive-commands",
		resourceType:  "overdrive command",
		objLookup:     cfg.l.OverdriveCommands,
		objLookupID:   cfg.l.OverdriveCommandsID,
		queryLookup:   cfg.q.overdriveCommands,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[string]SubSectionFns{
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetOverdriveCommandOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetOverdriveCommandOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrive-modes",
		resourceType:  "overdrive mode",
		objLookup:     cfg.l.OverdriveModes,
		objLookupID:   cfg.l.OverdriveModesID,
		queryLookup:   cfg.q.overdriveModes,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveModeIDs,
		getSingleFunc: cfg.getOverdriveMode,
		retrieveFunc:  cfg.retrieveOverdriveModes,
	}

	e.overdrives = handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrives",
		resourceType:  "overdrive",
		objLookup:     cfg.l.Overdrives,
		objLookupID:   cfg.l.OverdrivesID,
		queryLookup:   cfg.q.overdrives,
		idToResFunc:   idToNamedAPIResource[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createOverdriveSimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.abilities = handlerInput[seeding.Ability, Ability, AbilityAPIResource, AbilityAPIResourceList]{
		endpoint:      "abilities",
		resourceType:  "ability",
		objLookup:     cfg.l.Abilities,
		objLookupID:   cfg.l.AbilitiesID,
		queryLookup:   cfg.q.abilities,
		idToResFunc:   idToAbilityAPIResource,
		resToListFunc: newAbilityAPIResourceList,
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.unspecifiedAbilities = handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "unspecified-abilities",
		resourceType:     "unspecified ability",
		objLookup:        cfg.l.UnspecifiedAbilities,
		objLookupID:      cfg.l.UnspecifiedAbilitiesID,
		queryLookup:      cfg.q.unspecifiedAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetUnspecifiedAbilityIDsByName,
		retrieveQuery:    cfg.db.GetUnspecifiedAbilityIDs,
		getSingleFunc:    cfg.getUnspecifiedAbility,
		retrieveFunc:     cfg.retrieveUnspecifiedAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createUnspecifiedAbilitySimple,
			},
		},
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "player-abilities",
		resourceType:     "player ability",
		objLookup:        cfg.l.PlayerAbilities,
		objLookupID:      cfg.l.PlayerAbilitiesID,
		queryLookup:      cfg.q.playerAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetPlayerAbilityIDsByName,
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createPlayerAbilitySimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
			},
		},
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "enemy-abilities",
		resourceType:     "enemy ability",
		objLookup:        cfg.l.EnemyAbilities,
		objLookupID:      cfg.l.EnemyAbilitiesID,
		queryLookup:      cfg.q.enemyAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetEnemyAbilityIDsByName,
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createEnemyAbilitySimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
			},
		},
	}

	e.itemAbilities = handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "item-abilities",
		resourceType:  "item ability",
		objLookup:     cfg.l.ItemAbilities,
		objLookupID:   cfg.l.ItemAbilitiesID,
		queryLookup:   cfg.q.itemAbilities,
		idToResFunc:   idToNamedAPIResource[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "overdrive-abilities",
		resourceType:     "overdrive ability",
		objLookup:        cfg.l.OverdriveAbilities,
		objLookupID:      cfg.l.OverdriveAbilitiesID,
		queryLookup:      cfg.q.overdriveAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetOverdriveAbilityIDsByName,
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "trigger-commands",
		resourceType:     "trigger command",
		objLookup:        cfg.l.TriggerCommands,
		objLookupID:      cfg.l.TriggerCommandsID,
		queryLookup:      cfg.q.triggerCommands,
		idToResFunc:      idToNamedAPIResource[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetTriggerCommandIDsByName,
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	e.properties = handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "properties",
		resourceType:  "property",
		objLookup:     cfg.l.Properties,
		objLookupID:   cfg.l.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seeding.Property, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "ronso-rages",
		resourceType:  "ronso rage",
		objLookup:     cfg.l.RonsoRages,
		objLookupID:   cfg.l.RonsoRagesID,
		queryLookup:   cfg.q.ronsoRages,
		idToResFunc:   idToNamedAPIResource[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[string]SubSectionFns{
			"monsters": {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
			},
		},
	}

	e.shops = handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "shops",
		resourceType:  "shop",
		objLookup:     cfg.l.Shops,
		objLookupID:   cfg.l.ShopsID,
		queryLookup:   cfg.q.shops,
		idToResFunc:   idToUnnamedAPIResource[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createShopSimple,
			},
		},
	}

	e.sidequests = handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "sidequests",
		resourceType:  "sidequest",
		objLookup:     cfg.l.Sidequests,
		objLookupID:   cfg.l.SidequestsID,
		queryLookup:   cfg.q.sidequests,
		idToResFunc:   idToNamedAPIResource[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		subsections: map[string]SubSectionFns{
			"subquests": {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	e.subquests = handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "subquests",
		resourceType:  "subquest",
		objLookup:     cfg.l.Subquests,
		objLookupID:   cfg.l.SubquestsID,
		queryLookup:   cfg.q.subquests,
		idToResFunc:   idToNamedAPIResource[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
	}

	e.songs = handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "songs",
		resourceType:  "song",
		objLookup:     cfg.l.Songs,
		objLookupID:   cfg.l.SongsID,
		queryLookup:   cfg.q.songs,
		idToResFunc:   idToNamedAPIResource[seeding.Song, Song, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSongIDs,
		getSingleFunc: cfg.getSong,
		retrieveFunc:  cfg.retrieveSongs,
	}

	e.stats = handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "stats",
		resourceType:  "stat",
		objLookup:     cfg.l.Stats,
		objLookupID:   cfg.l.StatsID,
		idToResFunc:   idToNamedAPIResource[seeding.Stat, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "status-conditions",
		resourceType:  "status condition",
		objLookup:     cfg.l.StatusConditions,
		objLookupID:   cfg.l.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.sublocations = handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "sublocations",
		resourceType:  "sublocation",
		objLookup:     cfg.l.Sublocations,
		objLookupID:   cfg.l.SublocationsID,
		queryLookup:   cfg.q.sublocations,
		idToResFunc:   idToNamedAPIResource[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		subsections: map[string]SubSectionFns{
			"simple": {
				createSubFn: createSublocationSimple,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
			},
			"areas": {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
			},
			"shops": {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			"songs": {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.submenus = handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "submenus",
		resourceType:  "submenu",
		objLookup:     cfg.l.Submenus,
		objLookupID:   cfg.l.SubmenusID,
		queryLookup:   cfg.q.submenus,
		idToResFunc:   idToNamedAPIResource[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[string]SubSectionFns{
			"abilities": {
				dbQuery:     queryMany(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.topmenus = handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "topmenus",
		resourceType:  "topmenu",
		objLookup:     cfg.l.Topmenus,
		objLookupID:   cfg.l.TopmenusID,
		queryLookup:   cfg.q.topmenus,
		idToResFunc:   idToNamedAPIResource[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[string]SubSectionFns{
			"abilities": {
				dbQuery:     queryMany(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.treasures = handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "treasures",
		resourceType:  "treasure",
		objLookup:     cfg.l.Treasures,
		objLookupID:   cfg.l.TreasuresID,
		queryLookup:   cfg.q.treasures,
		idToResFunc:   idToUnnamedAPIResource[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	e.abilityType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "ability-type",
		resourceType:  "ability type",
		objLookup:     cfg.t.AbilityType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.attackType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "attack-type",
		resourceType:  "attack type",
		objLookup:     cfg.t.AttackType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.damageFormula = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "damage-formula",
		resourceType:  "damage formula",
		objLookup:     cfg.t.DamageFormula.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.damageType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "damage-type",
		resourceType:  "damage type",
		objLookup:     cfg.t.DamageType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.itemCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "item-category",
		resourceType:  "item category",
		objLookup:     cfg.t.ItemCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.lootType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "loot-type",
		resourceType:  "loot type",
		objLookup:     cfg.t.LootType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-type",
		resourceType:  "monster type",
		objLookup:     cfg.t.MonsterCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterFormationCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-formation-category",
		resourceType:  "monster formation category",
		objLookup:     cfg.t.MonsterFormationCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterSpecies = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-species",
		resourceType:  "monster species",
		objLookup:     cfg.t.MonsterSpecies.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.playerAbilityCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "player-ability-category",
		resourceType:  "player ability category",
		objLookup:     cfg.t.PlayerAbilityCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.shopCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "shop-category",
		resourceType:  "shop category",
		objLookup:     cfg.t.ShopCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	cfg.e = &e
}
