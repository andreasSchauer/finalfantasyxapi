package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	aeonCommands       handlerInput[seeding.AeonCommand, any, NamedAPIResource, NamedApiResourceList]
	aeons              handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]
	affinities         handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]
	arenaCreations     handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]
	areas              handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]
	autoAbilities      handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]
	blitzballPrizes    handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList]
	celestialWeapons   handlerInput[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList]
	characters         handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]
	characterClasses   handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]
	elements           handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]
	equipment          handlerInput[seeding.EquipmentName, any, NamedAPIResource, NamedApiResourceList]
	fmvs               handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]
	items              handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]
	keyItems           handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]
	locations          handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]
	modifiers		   handlerInput[seeding.Modifier, any, NamedAPIResource, NamedApiResourceList]
	monsters           handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]
	monsterFormations  handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]
	overdriveCommands  handlerInput[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList]
	overdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]
	overdrives         handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]
	otherAbilities     handlerInput[seeding.OtherAbility, OtherAbility, NamedAPIResource, NamedApiResourceList]
	playerAbilities    handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]
	enemyAbilities     handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]
	itemAbilities      handlerInput[seeding.Item, ItemAbility, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]
	triggerCommands    handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]
	properties         handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]
	ronsoRages         handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]
	shops              handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]
	sidequests         handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]
	subquests          handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]
	songs              handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]
	stats              handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]
	statusConditions   handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]
	sublocations       handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]
	submenus           handlerInput[seeding.Submenu, any, NamedAPIResource, NamedApiResourceList]
	topmenus           handlerInput[seeding.Topmenu, any, NamedAPIResource, NamedApiResourceList]
	treasures          handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]



	attackType				 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	damageFormula			 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	damageType				 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	itemCategory     	     handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterCategory          handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	lootType                 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterFormationCategory handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterSpecies           handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	playerAbilityCategory	 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	shopCategory           	 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeonCommands = handlerInput[seeding.AeonCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeon-commands",
		resourceType:  "aeon command",
		objLookup:     cfg.l.AeonCommands,
		objLookupID:   cfg.l.AeonCommandsID,
		idToResFunc:   idToNamedAPIResource[seeding.AeonCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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
				createSubFn: createShopSub,
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

	e.overdriveCommands = handlerInput[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrive-commands",
		resourceType:  "overdrive command",
		objLookup:     cfg.l.OverdriveCommands,
		objLookupID:   cfg.l.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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
		idToResFunc:   idToNamedAPIResource[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.otherAbilities = handlerInput[seeding.OtherAbility, OtherAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"other-abilities",
		resourceType:  		"other ability",
		objLookup:     		cfg.l.OtherAbilities,
		objLookupID:   		cfg.l.OtherAbilitiesID,
		queryLookup: 		cfg.q.otherAbilities,
		idToResFunc:   		idToNamedAPIResource[seeding.OtherAbility, OtherAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetOtherAbilityIDsByName,
		retrieveQuery: 		cfg.db.GetOtherAbilityIDs,
		getSingleFunc: 		cfg.getOtherAbility,
		retrieveFunc: 		cfg.retrieveOtherAbilities,
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"player-abilities",
		resourceType:  		"player ability",
		objLookup:     		cfg.l.PlayerAbilities,
		objLookupID:   		cfg.l.PlayerAbilitiesID,
		queryLookup:   		cfg.q.playerAbilities,
		idToResFunc:   		idToNamedAPIResource[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetPlayerAbilityIDsByName,
		retrieveQuery: 		cfg.db.GetPlayerAbilityIDs,
		getSingleFunc: 		cfg.getPlayerAbility,
		retrieveFunc: 		cfg.retrievePlayerAbilities,
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"enemy-abilities",
		resourceType:  		"enemy ability",
		objLookup:     		cfg.l.EnemyAbilities,
		objLookupID:   		cfg.l.EnemyAbilitiesID,
		queryLookup: 		cfg.q.enemyAbilities,
		idToResFunc:   		idToNamedAPIResource[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetEnemyAbilityIDsByName,
		retrieveQuery: 		cfg.db.GetEnemyAbilityIDs,
		getSingleFunc: 		cfg.getEnemyAbility,
		retrieveFunc: 		cfg.retrieveEnemyAbilities,
	}

	e.itemAbilities = handlerInput[seeding.Item, ItemAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"item-abilities",
		resourceType:  		"item ability",
		objLookup:     		cfg.l.Items,
		objLookupID:   		cfg.l.ItemsID,
		queryLookup: 		cfg.q.itemAbilities,
		idToResFunc:   		idToNamedAPIResource[seeding.Item, ItemAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		retrieveQuery: 		cfg.db.GetItemAbilityIDs,
		getSingleFunc: 		cfg.getItemAbility,
		retrieveFunc: 		cfg.retrieveItemAbilities,
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"overdrive-abilities",
		resourceType:  		"overdrive ability",
		objLookup:     		cfg.l.OverdriveAbilities,
		objLookupID:   		cfg.l.OverdriveAbilitiesID,
		queryLookup: 		cfg.q.overdriveAbilities,
		idToResFunc:   		idToNamedAPIResource[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetOverdriveAbilityIDsByName,
		retrieveQuery: 		cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc: 		cfg.getOverdriveAbility,
		retrieveFunc: 		cfg.retrieveOverdriveAbilities,
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      		"trigger-commands",
		resourceType:  		"trigger command",
		objLookup:     		cfg.l.TriggerCommands,
		objLookupID:   		cfg.l.TriggerCommandsID,
		queryLookup: 		cfg.q.triggerCommands,
		idToResFunc:   		idToNamedAPIResource[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 		newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetTriggerCommandIDsByName,
		retrieveQuery: 		cfg.db.GetTriggerCommandIDs,
		getSingleFunc: 		cfg.getTriggerCommand,
		retrieveFunc: 		cfg.retrieveTriggerCommands,
	}

	e.properties = handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "properties",
		resourceType:  "property",
		objLookup:     cfg.l.Properties,
		objLookupID:   cfg.l.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seeding.Property, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "ronso-rages",
		resourceType:  "ronso rage",
		objLookup:     cfg.l.RonsoRages,
		objLookupID:   cfg.l.RonsoRagesID,
		idToResFunc:   idToNamedAPIResource[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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
				createSubFn: createShopSub,
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
				createSubFn: createShopSub,
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

	e.submenus = handlerInput[seeding.Submenu, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "submenus",
		resourceType:  "submenu",
		objLookup:     cfg.l.Submenus,
		objLookupID:   cfg.l.SubmenusID,
		idToResFunc:   idToNamedAPIResource[seeding.Submenu, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.topmenus = handlerInput[seeding.Topmenu, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "topmenus",
		resourceType:  "topmenu",
		objLookup:     cfg.l.Topmenus,
		objLookupID:   cfg.l.TopmenusID,
		idToResFunc:   idToNamedAPIResource[seeding.Topmenu, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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


	e.attackType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "attack-type",
		resourceType:  "attack type",
		objLookup:     cfg.t.AttackType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.damageFormula = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "damage-formula",
		resourceType:  "damage formula",
		objLookup:     cfg.t.DamageFormula.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.damageType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "damage-type",
		resourceType:  "damage type",
		objLookup:     cfg.t.DamageType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.itemCategory = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "item-category",
		resourceType:  "item category",
		objLookup:     cfg.t.ItemCategory.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.lootType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "loot-type",
		resourceType:  "loot type",
		objLookup:     cfg.t.LootType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}
	
	e.monsterCategory = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "monster-type",
		resourceType:  "monster type",
		objLookup:     cfg.t.MonsterCategory.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.monsterFormationCategory = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "monster-formation-category",
		resourceType:  "monster formation category",
		objLookup:     cfg.t.MonsterFormationCategory.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.monsterSpecies = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "monster-species",
		resourceType:  "monster species",
		objLookup:     cfg.t.MonsterSpecies.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.playerAbilityCategory = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "player-ability-category",
		resourceType:  "player ability category",
		objLookup:     cfg.t.PlayerAbilityCategory.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.shopCategory = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "shop-category",
		resourceType:  "shop category",
		objLookup:     cfg.t.ShopCategory.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	cfg.e = &e
}
