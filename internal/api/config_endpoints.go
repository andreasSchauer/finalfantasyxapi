package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type endpoints struct {
	aeons              handlerInput[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList]
	affinities         handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]
	arenaCreations     handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]
	areas              handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]
	autoAbilities      handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]
	blitzballPrizes    handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList]
	celestialWeapons   handlerInput[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList]
	characters         handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]
	characterClasses   handlerInput[seeding.CharacterClass, any, NamedAPIResource, NamedApiResourceList]
	elements           handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]
	equipment          handlerInput[seeding.EquipmentName, any, NamedAPIResource, NamedApiResourceList]
	fmvs               handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]
	items              handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]
	keyItems           handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]
	locations          handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]
	monsters           handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]
	monsterFormations  handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]
	overdriveCommands  handlerInput[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList]
	overdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]
	genericAbilities   handlerInput[seeding.GenericAbility, any, NamedAPIResource, NamedApiResourceList]
	playerAbilities    handlerInput[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList]
	enemyAbilities     handlerInput[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList]
	itemAbilities      handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities handlerInput[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList]
	triggerCommands    handlerInput[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList]
	properties         handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]
	ronsoRages         handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]
	shops              handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]
	sidequests         handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]
	subquests          handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]
	songs              handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]
	stats              handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]
	statusConditions   handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]
	sublocations       handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]
	treasures          handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]

	connectionType           handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	creationArea             handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	ctbIconType              handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	lootType                 handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterFormationCategory handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterSpecies           handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	overdriveModeType        handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	treasureType             handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeons = handlerInput[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeons",
		resourceType:  "aeon",
		objLookup:     cfg.l.Aeons,
		objLookupID:   cfg.l.AeonsID,
		idToResFunc:   idToNamedAPIResource[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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
				createSubFn: createAreaSub,
			},
			"connected": {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSub,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSub,
			},
			"monsters": {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSub,
			},
			"songs": {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSub,
			},
			"treasures": {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSub,
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
		endpoint:      "character-classes",
		resourceType:  "character class",
		objLookup:     cfg.l.CelestialWeapons,
		objLookupID:   cfg.l.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seeding.CelestialWeapon, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.characters = handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]{
		endpoint:      	"characters",
		resourceType:  	"character",
		objLookup:     	cfg.l.Characters,
		objLookupID:   	cfg.l.CharactersID,
		queryLookup: 	cfg.q.characters,
		idToResFunc:   	idToNamedAPIResource[seeding.Character, Character, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 	newNamedAPIResourceList,
		retrieveQuery: 	cfg.db.GetCharacterIDs,
		getSingleFunc: 	cfg.getCharacter,
		retrieveFunc: 	cfg.retrieveCharacters,
	}

	e.characterClasses = handlerInput[seeding.CharacterClass, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "character-classes",
		resourceType:  "character class",
		objLookup:     cfg.l.CharClasses,
		objLookupID:   cfg.l.CharClassesID,
		idToResFunc:   idToNamedAPIResource[seeding.CharacterClass, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.connectionType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "connection-type",
		resourceType:  "connection type",
		objLookup:     cfg.t.AreaConnectionType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.creationArea = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "creation-area",
		resourceType:  "creation area",
		objLookup:     cfg.t.CreationArea.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.ctbIconType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "ctb-icon-type",
		resourceType:  "ctb icon type",
		objLookup:     cfg.t.CTBIconType.lookup,
		resToListFunc: newTypedAPIResourceList,
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

	e.items = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "items",
		resourceType:  "item",
		objLookup:     cfg.l.Items,
		objLookupID:   cfg.l.ItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
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
				createSubFn: createLocationSub,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSub,
			},
			"sublocations": {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSub,
			},
			"areas": {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSub,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSub,
			},
			"monsters": {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSub,
			},
			"shops": {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSub,
			},
			"songs": {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSub,
			},
			"treasures": {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSub,
			},
		},
	}

	e.lootType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "loot-type",
		resourceType:  "loot type",
		objLookup:     cfg.t.LootType.lookup,
		resToListFunc: newTypedAPIResourceList,
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
				createSubFn: createMonsterSub,
			},
			"areas": {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSub,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSub,
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
				createSubFn: createMonsterFormationSub,
			},
			"monsters": {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSub,
			},
		},
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

	e.overdriveCommands = handlerInput[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint: 		"overdrive-commands",
		resourceType: 	"overdrive command",
		objLookup: 		cfg.l.OverdriveCommands,
		objLookupID: 	cfg.l.OverdriveCommandsID,
		idToResFunc: 	idToNamedAPIResource[seeding.OverdriveCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 	newNamedAPIResourceList,
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

	e.overdriveModeType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "overdrive-mode-type",
		resourceType:  "overdrive mode type",
		objLookup:     cfg.t.OverdriveModeType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.genericAbilities = handlerInput[seeding.GenericAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "generic-abilities",
		resourceType:  "generic ability",
		objLookup:     cfg.l.GenericAbilities,
		objLookupID:   cfg.l.GenericAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.GenericAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "player-abilities",
		resourceType:  "player ability",
		objLookup:     cfg.l.PlayerAbilities,
		objLookupID:   cfg.l.PlayerAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "enemy-abilities",
		resourceType:  "enemy ability",
		objLookup:     cfg.l.EnemyAbilities,
		objLookupID:   cfg.l.EnemyAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.itemAbilities = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "item-abilities",
		resourceType:  "item ability",
		objLookup:     cfg.l.Items,
		objLookupID:   cfg.l.ItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrive-abilities",
		resourceType:  "overdrive ability",
		objLookup:     cfg.l.OverdriveAbilities,
		objLookupID:   cfg.l.OverdriveAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "trigger-commands",
		resourceType:  "trigger command",
		objLookup:     cfg.l.TriggerCommands,
		objLookupID:   cfg.l.TriggerCommandsID,
		idToResFunc:   idToNamedAPIResource[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
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
				createSubFn: createSubquestSub,
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
				createSubFn: createSublocationSub,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSub,
			},
			"areas": {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSub,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSub,
			},
			"monsters": {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSub,
			},
			"shops": {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSub,
			},
			"songs": {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSub,
			},
			"treasures": {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSub,
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

	e.treasureType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:      "treasure-type",
		resourceType:  "treasure type",
		objLookup:     cfg.t.TreasureType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	cfg.e = &e
}
