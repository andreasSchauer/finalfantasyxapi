package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	aeons              handlerInput[seeding.Aeon, any, NamedApiResourceList]
	affinities         handlerInput[seeding.Affinity, any, NamedApiResourceList]
	areas              handlerInput[seeding.Area, Area, LocationApiResourceList]
	autoAbilities      handlerInput[seeding.AutoAbility, any, NamedApiResourceList]
	characters         handlerInput[seeding.Character, any, NamedApiResourceList]
	connectionType     handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	creationArea       handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	ctbIconType        handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	elements           handlerInput[seeding.Element, any, NamedApiResourceList]
	fmvs               handlerInput[seeding.FMV, any, NamedApiResourceList]
	items              handlerInput[seeding.Item, any, NamedApiResourceList]
	keyItems           handlerInput[seeding.KeyItem, any, NamedApiResourceList]
	locations          handlerInput[seeding.Location, any, NamedApiResourceList]
	monsters           handlerInput[seeding.Monster, Monster, NamedApiResourceList]
	monsterFormations  handlerInput[seeding.EncounterLocation, any, UnnamedApiResourceList]
	monsterSpecies     handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	overdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedApiResourceList]
	overdriveModeType  handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	playerAbilities    handlerInput[seeding.PlayerAbility, any, NamedApiResourceList]
	enemyAbilities     handlerInput[seeding.EnemyAbility, any, NamedApiResourceList]
	itemAbilities      handlerInput[seeding.Item, any, NamedApiResourceList]
	overdriveAbilities handlerInput[seeding.OverdriveAbility, any, NamedApiResourceList]
	triggerCommands    handlerInput[seeding.TriggerCommand, any, NamedApiResourceList]
	properties         handlerInput[seeding.Property, any, NamedApiResourceList]
	ronsoRages         handlerInput[seeding.RonsoRage, any, NamedApiResourceList]
	shops              handlerInput[seeding.Shop, any, UnnamedApiResourceList]
	sidequests         handlerInput[seeding.Sidequest, any, NamedApiResourceList]
	songs              handlerInput[seeding.Song, any, NamedApiResourceList]
	stats              handlerInput[seeding.Stat, any, NamedApiResourceList]
	statusConditions   handlerInput[seeding.StatusCondition, any, NamedApiResourceList]
	sublocations       handlerInput[seeding.SubLocation, any, NamedApiResourceList]
	treasures          handlerInput[seeding.Treasure, any, UnnamedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeons = handlerInput[seeding.Aeon, any, NamedApiResourceList]{
		endpoint:     "aeons",
		resourceType: "aeon",
		objLookup:    cfg.l.Aeons,
	}

	e.affinities = handlerInput[seeding.Affinity, any, NamedApiResourceList]{
		endpoint:     "affinities",
		resourceType: "affinity",
		objLookup:    cfg.l.Affinities,
	}

	e.areas = handlerInput[seeding.Area, Area, LocationApiResourceList]{
		endpoint:        "areas",
		resourceType:    "area",
		objLookup:       cfg.l.Areas,
		objLookupID:     cfg.l.AreasID,
		queryLookup:     cfg.q.areas,
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

	e.autoAbilities = handlerInput[seeding.AutoAbility, any, NamedApiResourceList]{
		endpoint:     "auto-abilities",
		resourceType: "auto ability",
		objLookup:    cfg.l.AutoAbilities,
	}

	e.characters = handlerInput[seeding.Character, any, NamedApiResourceList]{
		endpoint:     "characters",
		resourceType: "character",
		objLookup:    cfg.l.Characters,
	}

	e.connectionType = handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "connection-type",
		resourceType: "connection type",
		objLookup:    cfg.t.AreaConnectionType,
	}

	e.creationArea = handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "creation-area",
		resourceType: "creation area",
		objLookup:    cfg.t.CreationArea,
	}

	e.ctbIconType = handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "ctb-icon-type",
		resourceType: "ctb icon type",
		objLookup:    cfg.t.CTBIconType,
	}

	e.elements = handlerInput[seeding.Element, any, NamedApiResourceList]{
		endpoint:     "elements",
		resourceType: "element",
		objLookup:    cfg.l.Elements,
	}

	e.fmvs = handlerInput[seeding.FMV, any, NamedApiResourceList]{
		endpoint:     "fmvs",
		resourceType: "FMV",
		objLookup:    cfg.l.FMVs,
	}

	e.items = handlerInput[seeding.Item, any, NamedApiResourceList]{
		endpoint:     "items",
		resourceType: "item",
		objLookup:    cfg.l.Items,
	}

	e.keyItems = handlerInput[seeding.KeyItem, any, NamedApiResourceList]{
		endpoint:     "key-items",
		resourceType: "key item",
		objLookup:    cfg.l.KeyItems,
	}

	e.locations = handlerInput[seeding.Location, any, NamedApiResourceList]{
		endpoint:     "locations",
		resourceType: "location",
		objLookup:    cfg.l.Locations,
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedApiResourceList]{
		endpoint:        "monsters",
		resourceType:    "monster",
		objLookup:       cfg.l.Monsters,
		queryLookup:     cfg.q.monsters,
		getSingleFunc:   cfg.getMonster,
		getMultipleFunc: cfg.getMultipleMonsters,
		retrieveFunc:    cfg.retrieveMonsters,
		subsections: map[string]func(string) (APIResourceList, error){
			"abilities": cfg.getMonsterAbilitiesMid,
		},
	}

	e.monsterFormations = handlerInput[seeding.EncounterLocation, any, UnnamedApiResourceList]{
		endpoint:     "monster-formations",
		resourceType: "monster formation",
		objLookup:    cfg.l.EncounterLocations,
	}

	e.monsterSpecies = handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "monster-species",
		resourceType: "monster species",
		objLookup:    cfg.t.MonsterSpecies,
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedApiResourceList]{
		endpoint:        "overdrive-modes",
		resourceType:    "overdrive mode",
		objLookup:       cfg.l.OverdriveModes,
		objLookupID:     cfg.l.OverdriveModesID,
		queryLookup:     cfg.q.overdriveModes,
		getSingleFunc:   cfg.getOverdriveMode,
		getMultipleFunc: nil,
		retrieveFunc:    cfg.retrieveOverdriveModes,
	}

	e.overdriveModeType = handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "overdrive-mode-type",
		resourceType: "overdrive mode type",
		objLookup:    cfg.t.OverdriveModeType,
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, any, NamedApiResourceList]{
		endpoint:     "player-abilities",
		resourceType: "player ability",
		objLookup:    cfg.l.PlayerAbilities,
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, any, NamedApiResourceList]{
		endpoint:     "enemy-abilities",
		resourceType: "enemy ability",
		objLookup:    cfg.l.EnemyAbilities,
	}

	e.itemAbilities = handlerInput[seeding.Item, any, NamedApiResourceList]{
		endpoint:     "item-abilities",
		resourceType: "item ability",
		objLookup:    cfg.l.Items,
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, any, NamedApiResourceList]{
		endpoint:     "overdrive-abilities",
		resourceType: "overdrive ability",
		objLookup:    cfg.l.OverdriveAbilities,
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, any, NamedApiResourceList]{
		endpoint:     "trigger-commands",
		resourceType: "trigger command",
		objLookup:    cfg.l.TriggerCommands,
	}

	e.properties = handlerInput[seeding.Property, any, NamedApiResourceList]{
		endpoint:     "properties",
		resourceType: "property",
		objLookup:    cfg.l.Properties,
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, any, NamedApiResourceList]{
		endpoint:     "ronso-rages",
		resourceType: "ronso rage",
		objLookup:    cfg.l.RonsoRages,
	}

	e.shops = handlerInput[seeding.Shop, any, UnnamedApiResourceList]{
		endpoint:     "shops",
		resourceType: "shop",
		objLookup:    cfg.l.Shops,
	}

	e.sidequests = handlerInput[seeding.Sidequest, any, NamedApiResourceList]{
		endpoint:     "sidequests",
		resourceType: "sidequest",
		objLookup:    cfg.l.Sidequests,
	}

	e.songs = handlerInput[seeding.Song, any, NamedApiResourceList]{
		endpoint:     "songs",
		resourceType: "song",
		objLookup:    cfg.l.Songs,
	}

	e.stats = handlerInput[seeding.Stat, any, NamedApiResourceList]{
		endpoint:     "stats",
		resourceType: "stat",
		objLookup:    cfg.l.Stats,
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, any, NamedApiResourceList]{
		endpoint:     "status-conditions",
		resourceType: "status condition",
		objLookup:    cfg.l.StatusConditions,
	}

	e.sublocations = handlerInput[seeding.SubLocation, any, NamedApiResourceList]{
		endpoint:     "sublocations",
		resourceType: "sublocation",
		objLookup:    cfg.l.SubLocations,
	}

	e.treasures = handlerInput[seeding.Treasure, any, UnnamedApiResourceList]{
		endpoint:     "treasures",
		resourceType: "treasure",
		objLookup:    cfg.l.Treasures,
	}

	cfg.e = &e
}
