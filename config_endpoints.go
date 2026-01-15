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
	elements           handlerInput[seeding.Element, any, NamedApiResourceList]
	fmvs               handlerInput[seeding.FMV, any, NamedApiResourceList]
	items              handlerInput[seeding.Item, any, NamedApiResourceList]
	keyItems           handlerInput[seeding.KeyItem, any, NamedApiResourceList]
	locations          handlerInput[seeding.Location, any, NamedApiResourceList]
	monsters           handlerInput[seeding.Monster, Monster, NamedApiResourceList]
	monsterFormations  handlerInput[seeding.MonsterFormation, any, UnnamedApiResourceList]
	overdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedApiResourceList]
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

	connectionType    handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	creationArea      handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	ctbIconType       handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	monsterSpecies    handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
	overdriveModeType handlerInput[TypedAPIResource, TypedAPIResource, TypedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.aeons = handlerInput[seeding.Aeon, any, NamedApiResourceList]{
		endpoint:     "aeons",
		resourceType: "aeon",
		objLookup:    cfg.l.Aeons,
		objLookupID:  cfg.l.AeonsID,
	}

	e.affinities = handlerInput[seeding.Affinity, any, NamedApiResourceList]{
		endpoint:     "affinities",
		resourceType: "affinity",
		objLookup:    cfg.l.Affinities,
		objLookupID:  cfg.l.AffinitiesID,
	}

	e.areas = handlerInput[seeding.Area, Area, LocationApiResourceList]{
		endpoint:        "areas",
		resourceType:    "area",
		objLookup:       cfg.l.Areas,
		objLookupID:     cfg.l.AreasID,
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

	e.autoAbilities = handlerInput[seeding.AutoAbility, any, NamedApiResourceList]{
		endpoint:     "auto-abilities",
		resourceType: "auto ability",
		objLookup:    cfg.l.AutoAbilities,
		objLookupID:  cfg.l.AutoAbilitiesID,
	}

	e.characters = handlerInput[seeding.Character, any, NamedApiResourceList]{
		endpoint:     "characters",
		resourceType: "character",
		objLookup:    cfg.l.Characters,
		objLookupID:  cfg.l.CharactersID,
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
		objLookupID:  cfg.l.ElementsID,
	}

	e.fmvs = handlerInput[seeding.FMV, any, NamedApiResourceList]{
		endpoint:     "fmvs",
		resourceType: "FMV",
		objLookup:    cfg.l.FMVs,
		objLookupID:  cfg.l.FMVsID,
	}

	e.items = handlerInput[seeding.Item, any, NamedApiResourceList]{
		endpoint:     "items",
		resourceType: "item",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
	}

	e.keyItems = handlerInput[seeding.KeyItem, any, NamedApiResourceList]{
		endpoint:     "key-items",
		resourceType: "key item",
		objLookup:    cfg.l.KeyItems,
		objLookupID:  cfg.l.KeyItemsID,
	}

	e.locations = handlerInput[seeding.Location, any, NamedApiResourceList]{
		endpoint:     "locations",
		resourceType: "location",
		objLookup:    cfg.l.Locations,
		objLookupID:  cfg.l.LocationsID,
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedApiResourceList]{
		endpoint:        "monsters",
		resourceType:    "monster",
		objLookup:       cfg.l.Monsters,
		objLookupID:  cfg.l.MonstersID,
		queryLookup:     cfg.q.monsters,
		getSingleFunc:   cfg.getMonster,
		getMultipleFunc: cfg.getMultipleMonsters,
		retrieveFunc:    cfg.retrieveMonsters,
		subsections: map[string]func(string) (APIResourceList, error){
			"abilities": cfg.getMonsterAbilitiesMid,
		},
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, any, UnnamedApiResourceList]{
		endpoint:     "monster-formations",
		resourceType: "monster formation",
		objLookupID:  cfg.l.MonsterFormationsID,
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
		retrieveQuery:   cfg.db.GetOverdriveModeIDs,
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
		objLookupID:  cfg.l.PlayerAbilitiesID,
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, any, NamedApiResourceList]{
		endpoint:     "enemy-abilities",
		resourceType: "enemy ability",
		objLookup:    cfg.l.EnemyAbilities,
		objLookupID:  cfg.l.EnemyAbilitiesID,
	}

	e.itemAbilities = handlerInput[seeding.Item, any, NamedApiResourceList]{
		endpoint:     "item-abilities",
		resourceType: "item ability",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, any, NamedApiResourceList]{
		endpoint:     "overdrive-abilities",
		resourceType: "overdrive ability",
		objLookup:    cfg.l.OverdriveAbilities,
		objLookupID:  cfg.l.OverdriveAbilitiesID,
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, any, NamedApiResourceList]{
		endpoint:     "trigger-commands",
		resourceType: "trigger command",
		objLookup:    cfg.l.TriggerCommands,
		objLookupID:  cfg.l.TriggerCommandsID,
	}

	e.properties = handlerInput[seeding.Property, any, NamedApiResourceList]{
		endpoint:     "properties",
		resourceType: "property",
		objLookup:    cfg.l.Properties,
		objLookupID:  cfg.l.PropertiesID,
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, any, NamedApiResourceList]{
		endpoint:     "ronso-rages",
		resourceType: "ronso rage",
		objLookup:    cfg.l.RonsoRages,
		objLookupID:  cfg.l.RonsoRagesID,
	}

	e.shops = handlerInput[seeding.Shop, any, UnnamedApiResourceList]{
		endpoint:     "shops",
		resourceType: "shop",
		objLookup:    cfg.l.Shops,
		objLookupID:  cfg.l.ShopsID,
	}

	e.sidequests = handlerInput[seeding.Sidequest, any, NamedApiResourceList]{
		endpoint:     "sidequests",
		resourceType: "sidequest",
		objLookup:    cfg.l.Sidequests,
		objLookupID:  cfg.l.SidequestsID,
	}

	e.songs = handlerInput[seeding.Song, any, NamedApiResourceList]{
		endpoint:     "songs",
		resourceType: "song",
		objLookup:    cfg.l.Songs,
		objLookupID:  cfg.l.SongsID,
	}

	e.stats = handlerInput[seeding.Stat, any, NamedApiResourceList]{
		endpoint:     "stats",
		resourceType: "stat",
		objLookup:    cfg.l.Stats,
		objLookupID:  cfg.l.StatsID,
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, any, NamedApiResourceList]{
		endpoint:     "status-conditions",
		resourceType: "status condition",
		objLookup:    cfg.l.StatusConditions,
		objLookupID:  cfg.l.StatusConditionsID,
	}

	e.sublocations = handlerInput[seeding.SubLocation, any, NamedApiResourceList]{
		endpoint:     "sublocations",
		resourceType: "sublocation",
		objLookup:    cfg.l.Sublocations,
		objLookupID:  cfg.l.SublocationsID,
	}

	e.treasures = handlerInput[seeding.Treasure, any, UnnamedApiResourceList]{
		endpoint:     "treasures",
		resourceType: "treasure",
		objLookup:    cfg.l.Treasures,
		objLookupID:  cfg.l.TreasuresID,
	}

	cfg.e = &e
}
