package main

import (
	"context"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type handlerInput[T h.HasID, R any, A APIResource, L APIResourceList] struct {
	endpoint        	string
	resourceType    	string
	objLookup       	map[string]T
	objLookupID     	map[int32]T
	queryLookup     	map[string]QueryType
	getMultipleQuery 	func(context.Context, string) ([]int32, error)
	retrieveQuery   	func(context.Context) ([]int32, error)
	idToResFunc     	func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc		func(*Config, *http.Request, []A) (L, error)
	getSingleFunc   	func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc    	func(*http.Request, handlerInput[T, R, A, L]) (L, error)
	subsections     	map[string]SubSectionFns
}

type SubSectionFns struct {
	dbQuery			func(context.Context, int32) ([]int32, error)
	getResultsFn	func(*Config, []int32) []SubResource
}

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
		resToListFunc: newNamedAPIResourceList,
	}

	e.affinities = handlerInput[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     	"affinities",
		resourceType: 	"affinity",
		objLookup:    	cfg.l.Affinities,
		objLookupID:  	cfg.l.AffinitiesID,
		idToResFunc:  	idToNamedAPIResource[seeding.Affinity, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 	newNamedAPIResourceList,
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
		retrieveFunc:    cfg.retrieveAreas,
		subsections: 	 map[string]SubSectionFns{
			"monsters": {
				dbQuery: 		cfg.db.GetAreaMonsterIDs,
				getResultsFn: 	getSubMonsters,
			},
		},
	}

	e.autoAbilities = handlerInput[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "auto-abilities",
		resourceType: "auto ability",
		objLookup:    cfg.l.AutoAbilities,
		objLookupID:  cfg.l.AutoAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.AutoAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.characters = handlerInput[seeding.Character, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "characters",
		resourceType: "character",
		objLookup:    cfg.l.Characters,
		objLookupID:  cfg.l.CharactersID,
		idToResFunc:  idToNamedAPIResource[seeding.Character, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.connectionType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "connection-type",
		resourceType: "connection type",
		objLookup:    cfg.t.AreaConnectionType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.creationArea = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "creation-area",
		resourceType: "creation area",
		objLookup:    cfg.t.CreationArea.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.ctbIconType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "ctb-icon-type",
		resourceType: "ctb icon type",
		objLookup:    cfg.t.CTBIconType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.elements = handlerInput[seeding.Element, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "elements",
		resourceType: "element",
		objLookup:    cfg.l.Elements,
		objLookupID:  cfg.l.ElementsID,
		idToResFunc:  idToNamedAPIResource[seeding.Element, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.fmvs = handlerInput[seeding.FMV, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "fmvs",
		resourceType: "FMV",
		objLookup:    cfg.l.FMVs,
		objLookupID:  cfg.l.FMVsID,
		idToResFunc:  idToNamedAPIResource[seeding.FMV, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.items = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "items",
		resourceType: "item",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.keyItems = handlerInput[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "key-items",
		resourceType: "key item",
		objLookup:    cfg.l.KeyItems,
		objLookupID:  cfg.l.KeyItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.KeyItem, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.locations = handlerInput[seeding.Location, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "locations",
		resourceType: "location",
		objLookup:    cfg.l.Locations,
		objLookupID:  cfg.l.LocationsID,
		idToResFunc:  idToNamedAPIResource[seeding.Location, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]{
		endpoint:        	"monsters",
		resourceType:    	"monster",
		objLookup:       	cfg.l.Monsters,
		objLookupID:     	cfg.l.MonstersID,
		queryLookup:     	cfg.q.monsters,
		idToResFunc:     	idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		resToListFunc: 	 	newNamedAPIResourceList,
		getMultipleQuery: 	cfg.db.GetMonsterIDsByName,
		retrieveQuery:	 	cfg.db.GetMonsterIDs,
		getSingleFunc:   	cfg.getMonster,
		retrieveFunc:    	cfg.retrieveMonsters,
		
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "monster-formations",
		resourceType: "monster formation",
		objLookupID:  cfg.l.MonsterFormationsID,
		idToResFunc:  idToUnnamedAPIResource[seeding.MonsterFormation, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
	}

	e.monsterSpecies = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "monster-species",
		resourceType: "monster species",
		objLookup:    cfg.t.MonsterSpecies.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]{
		endpoint:        "overdrive-modes",
		resourceType:    "overdrive mode",
		objLookup:       cfg.l.OverdriveModes,
		objLookupID:     cfg.l.OverdriveModesID,
		queryLookup:     cfg.q.overdriveModes,
		idToResFunc:  	 idToNamedAPIResource[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery:   cfg.db.GetOverdriveModeIDs,
		getSingleFunc:   cfg.getOverdriveMode,
		retrieveFunc:    cfg.retrieveOverdriveModes,
	}

	e.overdriveModeType = handlerInput[TypedAPIResource, TypedAPIResource, TypedAPIResource, TypedApiResourceList]{
		endpoint:     "overdrive-mode-type",
		resourceType: "overdrive mode type",
		objLookup:    cfg.t.OverdriveModeType.lookup,
		resToListFunc: newTypedAPIResourceList,
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "player-abilities",
		resourceType: "player ability",
		objLookup:    cfg.l.PlayerAbilities,
		objLookupID:  cfg.l.PlayerAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.PlayerAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "enemy-abilities",
		resourceType: "enemy ability",
		objLookup:    cfg.l.EnemyAbilities,
		objLookupID:  cfg.l.EnemyAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.EnemyAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.itemAbilities = handlerInput[seeding.Item, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "item-abilities",
		resourceType: "item ability",
		objLookup:    cfg.l.Items,
		objLookupID:  cfg.l.ItemsID,
		idToResFunc:  idToNamedAPIResource[seeding.Item, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "overdrive-abilities",
		resourceType: "overdrive ability",
		objLookup:    cfg.l.OverdriveAbilities,
		objLookupID:  cfg.l.OverdriveAbilitiesID,
		idToResFunc:  idToNamedAPIResource[seeding.OverdriveAbility, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "trigger-commands",
		resourceType: "trigger command",
		objLookup:    cfg.l.TriggerCommands,
		objLookupID:  cfg.l.TriggerCommandsID,
		idToResFunc:  idToNamedAPIResource[seeding.TriggerCommand, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.properties = handlerInput[seeding.Property, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "properties",
		resourceType: "property",
		objLookup:    cfg.l.Properties,
		objLookupID:  cfg.l.PropertiesID,
		idToResFunc:  idToNamedAPIResource[seeding.Property, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "ronso-rages",
		resourceType: "ronso rage",
		objLookup:    cfg.l.RonsoRages,
		objLookupID:  cfg.l.RonsoRagesID,
		idToResFunc:  idToNamedAPIResource[seeding.RonsoRage, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.shops = handlerInput[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "shops",
		resourceType: "shop",
		objLookup:    cfg.l.Shops,
		objLookupID:  cfg.l.ShopsID,
		idToResFunc:  idToUnnamedAPIResource[seeding.Shop, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
	}

	e.sidequests = handlerInput[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "sidequests",
		resourceType: "sidequest",
		objLookup:    cfg.l.Sidequests,
		objLookupID:  cfg.l.SidequestsID,
		idToResFunc:  idToNamedAPIResource[seeding.Sidequest, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.songs = handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "songs",
		resourceType: "song",
		objLookup:    cfg.l.Songs,
		objLookupID:  cfg.l.SongsID,
		idToResFunc:  idToNamedAPIResource[seeding.Song, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.stats = handlerInput[seeding.Stat, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "stats",
		resourceType: "stat",
		objLookup:    cfg.l.Stats,
		objLookupID:  cfg.l.StatsID,
		idToResFunc:  idToNamedAPIResource[seeding.Stat, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "status-conditions",
		resourceType: "status condition",
		objLookup:    cfg.l.StatusConditions,
		objLookupID:  cfg.l.StatusConditionsID,
		idToResFunc:  idToNamedAPIResource[seeding.StatusCondition, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.sublocations = handlerInput[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList]{
		endpoint:     "sublocations",
		resourceType: "sublocation",
		objLookup:    cfg.l.Sublocations,
		objLookupID:  cfg.l.SublocationsID,
		idToResFunc:  idToNamedAPIResource[seeding.SubLocation, any, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.treasures = handlerInput[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:     "treasures",
		resourceType: "treasure",
		objLookup:    cfg.l.Treasures,
		objLookupID:  cfg.l.TreasuresID,
		idToResFunc:  idToUnnamedAPIResource[seeding.Treasure, any, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
	}

	cfg.e = &e
}
