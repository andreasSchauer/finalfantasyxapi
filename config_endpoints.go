package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	aeons              handlerInput[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList]
	affinities         handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]
	areas              handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList]
	autoAbilities      handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]
	characters         handlerInput[seeding.Character, any, NamedAPIResource, NamedApiResourceList]
	elements           handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]
	fmvs               handlerInput[seeding.FMV, any, NamedAPIResource, NamedApiResourceList]
	items              handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]
	keyItems           handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]
	locations          handlerInput[seeding.Location, any, NamedAPIResource, NamedApiResourceList]
	monsters           handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]
	monsterFormations  handlerInput[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList]
	overdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]
	playerAbilities    handlerInput[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList]
	enemyAbilities     handlerInput[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList]
	itemAbilities      handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities handlerInput[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList]
	triggerCommands    handlerInput[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList]
	properties         handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]
	ronsoRages         handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]
	shops              handlerInput[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList]
	sidequests         handlerInput[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList]
	songs              handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList]
	stats              handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]
	statusConditions   handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]
	sublocations       handlerInput[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList]
	treasures          handlerInput[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList]

	connectionType    handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	creationArea      handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	ctbIconType       handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterSpecies    handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	overdriveModeType handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeons = handlerInput[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "aeons",
		resourceType: "aeon",
		objLookup:    cfg.l.Aeons,
		objLookupID:  cfg.l.AeonsID,
		idToResFunc:  idToNamedAPIResource[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Aeon, any, NamedAPIResource, NamedApiResourceList],
	}

	e.affinities = handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     	"affinities",
		resourceType: 	"affinity",
		objLookup:    	cfg.l.Affinities,
		objLookupID:  	cfg.l.AffinitiesID,
		idToResFunc:  	idToNamedAPIResource[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 	newNamedAPIResourceList[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList],
	}

	e.areas = handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList]{
		endpoint:        "areas",
		resourceType:    "area",
		objLookup:       cfg.l.Areas,
		objLookupID:     cfg.l.AreasID,
		idToResFunc:  	 idToLocationAPIResource,
		resToListFunc: 	 newLocationAPIResourceList,
		queryLookup:     cfg.q.areas,
		retrieveQuery: 	 cfg.db.GetAreas,
		getSingleFunc:   cfg.getArea,
		getMultipleFunc: nil,
		retrieveFunc:    cfg.retrieveAreas,
		subsections: map[string]func(string) (APIResourceList, error){
			"treasures":          cfg.getAreaTreasuresMid,
			"shops":              cfg.getAreaShopsMid,
			"monsters":           cfg.getAreaMonstersMid,
			"monster-formations": cfg.getAreaFormationsMid,
			"connected":          cfg.getAreaConnectionsMid,
		},
	}

	e.autoAbilities = handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "auto-abilities",
		resourceType: "auto ability",
		objLookup:    cfg.l.AutoAbilities,
		objLookupID:  cfg.l.AutoAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList],
	}

	e.characters = handlerInput[seeding.Character, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "characters",
		resourceType: "character",
		objLookup:    cfg.l.Characters,
		objLookupID:  cfg.l.CharactersID,
		idToResFunc:  idToNamedAPIResource[seeding.Character, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Character, any, NamedAPIResource, NamedApiResourceList],
	}

	e.connectionType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "connection-type",
		resourceType: "connection type",
		objLookup:    cfg.t.AreaConnectionType,
	}

	e.creationArea = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "creation-area",
		resourceType: "creation area",
		objLookup:    cfg.t.CreationArea,
	}

	e.ctbIconType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "ctb-icon-type",
		resourceType: "ctb icon type",
		objLookup:    cfg.t.CTBIconType,
	}

	e.elements = handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "elements",
		resourceType: "element",
		objLookup:    cfg.l.Elements,
		objLookupID:  cfg.l.ElementsID,
		idToResFunc:  idToNamedAPIResource[seeding.Element, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Element, any, NamedAPIResource, NamedApiResourceList],
	}

	e.fmvs = handlerInput[seeding.FMV, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "fmvs",
		resourceType: "FMV",
		objLookup:    cfg.l.FMVs,
		objLookupID:  cfg.l.FMVsID,
		idToResFunc:  idToNamedAPIResource[seeding.FMV, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.FMV, any, NamedAPIResource, NamedApiResourceList],
	}

	e.items = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "items",
		resourceType: "item",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
	}

	e.keyItems = handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "key-items",
		resourceType: "key item",
		objLookup:    cfg.l.KeyItems,
		objLookupID:  cfg.l.KeyItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList],
	}

	e.locations = handlerInput[seeding.Location, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "locations",
		resourceType: "location",
		objLookup:    cfg.l.Locations,
		objLookupID:  cfg.l.LocationsID,
		idToResFunc:  idToNamedAPIResource[seeding.Location, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Location, any, NamedAPIResource, NamedApiResourceList],
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]{
		endpoint:        "monsters",
		resourceType:    "monster",
		objLookup:       cfg.l.Monsters,
		objLookupID:     cfg.l.MonstersID,
		queryLookup:     cfg.q.monsters,
		idToResFunc:     idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		getSingleFunc:   cfg.getMonster,
		getMultipleFunc: cfg.getMultipleMonsters,
		retrieveFunc:    cfg.retrieveMonsters,
		subsections: map[string]func(string) (APIResourceList, error){
			"abilities": cfg.getMonsterAbilitiesMid,
		},
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "monster-formations",
		resourceType: "monster formation",
		objLookupID:  cfg.l.MonsterFormationsID,
		idToResFunc:  idToUnnamedAPIResource[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList],
	}

	e.monsterSpecies = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "monster-species",
		resourceType: "monster species",
		objLookup:    cfg.t.MonsterSpecies,
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]{
		endpoint:        "overdrive-modes",
		resourceType:    "overdrive mode",
		objLookup:       cfg.l.OverdriveModes,
		objLookupID:     cfg.l.OverdriveModesID,
		queryLookup:     cfg.q.overdriveModes,
		idToResFunc:  	 idToNamedAPIResource[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList],
		retrieveQuery:   cfg.db.GetOverdriveModeIDs,
		getSingleFunc:   cfg.getOverdriveMode,
		getMultipleFunc: nil,
		retrieveFunc:    cfg.retrieveOverdriveModes,
	}

	e.overdriveModeType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "overdrive-mode-type",
		resourceType: "overdrive mode type",
		objLookup:    cfg.t.OverdriveModeType,
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "player-abilities",
		resourceType: "player ability",
		objLookup:    cfg.l.PlayerAbilities,
		objLookupID:  cfg.l.PlayerAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList],
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "enemy-abilities",
		resourceType: "enemy ability",
		objLookup:    cfg.l.EnemyAbilities,
		objLookupID:  cfg.l.EnemyAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList],
	}

	e.itemAbilities = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "item-abilities",
		resourceType: "item ability",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "overdrive-abilities",
		resourceType: "overdrive ability",
		objLookup:    cfg.l.OverdriveAbilities,
		objLookupID:  cfg.l.OverdriveAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList],
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "trigger-commands",
		resourceType: "trigger command",
		objLookup:    cfg.l.TriggerCommands,
		objLookupID:  cfg.l.TriggerCommandsID,
		idToResFunc:  idToNamedAPIResource[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList],
	}

	e.properties = handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "properties",
		resourceType: "property",
		objLookup:    cfg.l.Properties,
		objLookupID:  cfg.l.PropertiesID,
		idToResFunc:  idToNamedAPIResource[seeding.Property, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Property, any, NamedAPIResource, NamedApiResourceList],
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "ronso-rages",
		resourceType: "ronso rage",
		objLookup:    cfg.l.RonsoRages,
		objLookupID:  cfg.l.RonsoRagesID,
		idToResFunc:  idToNamedAPIResource[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList],
	}

	e.shops = handlerInput[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "shops",
		resourceType: "shop",
		objLookup:    cfg.l.Shops,
		objLookupID:  cfg.l.ShopsID,
		idToResFunc:  idToUnnamedAPIResource[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList],
	}

	e.sidequests = handlerInput[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "sidequests",
		resourceType: "sidequest",
		objLookup:    cfg.l.Sidequests,
		objLookupID:  cfg.l.SidequestsID,
		idToResFunc:  idToNamedAPIResource[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList],
	}

	e.songs = handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "songs",
		resourceType: "song",
		objLookup:    cfg.l.Songs,
		objLookupID:  cfg.l.SongsID,
		idToResFunc:  idToNamedAPIResource[seeding.Song, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Song, any, NamedAPIResource, NamedApiResourceList],
	}

	e.stats = handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "stats",
		resourceType: "stat",
		objLookup:    cfg.l.Stats,
		objLookupID:  cfg.l.StatsID,
		idToResFunc:  idToNamedAPIResource[seeding.Stat, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.Stat, any, NamedAPIResource, NamedApiResourceList],
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "status-conditions",
		resourceType: "status condition",
		objLookup:    cfg.l.StatusConditions,
		objLookupID:  cfg.l.StatusConditionsID,
		idToResFunc:  idToNamedAPIResource[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList],
	}

	e.sublocations = handlerInput[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "sublocations",
		resourceType: "sublocation",
		objLookup:    cfg.l.Sublocations,
		objLookupID:  cfg.l.SublocationsID,
		idToResFunc:  idToNamedAPIResource[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList],
	}

	e.treasures = handlerInput[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "treasures",
		resourceType: "treasure",
		objLookup:    cfg.l.Treasures,
		objLookupID:  cfg.l.TreasuresID,
		idToResFunc:  idToUnnamedAPIResource[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList],
	}

	cfg.e = &e
}
