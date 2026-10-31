package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	Locations          handlerInput[seeding.Location, Location, NamedAPIResource, NamedAPIResourceList]
	Sublocations       handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedAPIResourceList]
	Areas              handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList]
	MonsterFormations  handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedAPIResourceList]
	Shops              handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedAPIResourceList]
	Treasures          handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedAPIResourceList]
	Quests             handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestAPIResourceList]
	Sidequests         handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestAPIResourceList]
	Subquests          handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestAPIResourceList]
	ArenaCreations     handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedAPIResourceList]
	BlitzballPrizes    handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedAPIResourceList]
	Songs              handlerInput[seeding.Song, Song, NamedAPIResource, NamedAPIResourceList]
	FMVs               handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedAPIResourceList]
	PlayerUnits        handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]
	Characters         handlerInput[seeding.Character, Character, NamedAPIResource, NamedAPIResourceList]
	Aeons              handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedAPIResourceList]
	CharacterClasses   handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedAPIResourceList]
	Monsters           handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedAPIResourceList]
	Abilities          handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]
	PlayerAbilities    handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedAPIResourceList]
	OverdriveAbilities handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedAPIResourceList]
	ItemAbilities      handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedAPIResourceList]
	TriggerCommands    handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedAPIResourceList]
	MiscAbilities      handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedAPIResourceList]
	EnemyAbilities     handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedAPIResourceList]
	AeonCommands       handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedAPIResourceList]
	OverdriveCommands  handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedAPIResourceList]
	Overdrives         handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedAPIResourceList]
	RonsoRages         handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedAPIResourceList]
	Submenus           handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedAPIResourceList]
	Topmenus           handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedAPIResourceList]
	AllItems           handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]
	Items              handlerInput[seeding.Item, Item, NamedAPIResource, NamedAPIResourceList]
	KeyItems           handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedAPIResourceList]
	Spheres            handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedAPIResourceList]
	Primers            handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedAPIResourceList]
	Mixes              handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedAPIResourceList]
	AutoAbilities      handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedAPIResourceList]
	EquipmentTables    handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedAPIResourceList]
	Equipment          handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedAPIResourceList]
	CelestialWeapons   handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedAPIResourceList]
	Stats              handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedAPIResourceList]
	Properties         handlerInput[seeding.Property, Property, NamedAPIResource, NamedAPIResourceList]
	OverdriveModes     handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedAPIResourceList]
	Elements           handlerInput[seeding.Element, Element, NamedAPIResource, NamedAPIResourceList]
	StatusConditions   handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedAPIResourceList]
	Modifiers          handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedAPIResourceList]
	AgilityTiers       handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedAPIResourceList]
	Enums              handlerInputEnums
	Endpoints          handlerInputEndpoints
	AlBhed             handlerInputService[AlBhedParams, AlBhedResponse]
	TurnOrder          handlerInputService[TurnOrderParams, TurnOrderResponse]
	DropChance         handlerInputService[DropChanceParams, DropChanceResponse]
	CritChance         handlerInputService[CritChanceParams, CritChanceResponse]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.Locations = handlerInput[seeding.Location, Location, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epLocations,
		resTypeSingle: rtsLocations,
		resTypePlural: rtpLocations,
		objLookup:     cfg.L.Locations,
		objLookupID:   cfg.L.LocationsID,
		idToResFunc:   idToNamedAPIResource[seeding.Location, Location, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			SnConnected: {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			SnSublocations: {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			SnAreas: {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			SnMonsterFormations: {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			SnShops: {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			SnSongs: {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			SnTreasures: {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.Sublocations = handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epSublocations,
		resTypeSingle: rtsSublocations,
		resTypePlural: rtpSublocations,
		objLookup:     cfg.L.Sublocations,
		objLookupID:   cfg.L.SublocationsID,
		idToResFunc:   idToNamedAPIResource[seeding.Sublocation, Sublocation, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			SnConnected: {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			SnAreas: {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			SnMonsterFormations: {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			SnShops: {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			SnSongs: {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			SnTreasures: {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.Areas = handlerInput[seeding.Area, Area, AreaAPIResource, AreaAPIResourceList]{
		Endpoint:      epAreas,
		resTypeSingle: rtsAreas,
		resTypePlural: rtpAreas,
		objLookup:     cfg.L.Areas,
		objLookupID:   cfg.L.AreasID,
		idToResFunc:   idToAreaAPIResource,
		resToListFunc: newAreaAPIResourceList,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			SnConnected: {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			SnMonsterFormations: {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			SnSongs: {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			SnTreasures: {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.MonsterFormations = handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedAPIResourceList]{
		Endpoint:      epMonsterFormations,
		resTypeSingle: rtsMonsterFormations,
		resTypePlural: rtpMonsterFormations,
		objLookup:     cfg.L.MonsterFormations,
		objLookupID:   cfg.L.MonsterFormationsID,
		idToResFunc:   idToUnnamedAPIResource[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedAPIResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.Shops = handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedAPIResourceList]{
		Endpoint:      epShops,
		resTypeSingle: rtsShops,
		resTypePlural: rtpShops,
		objLookup:     cfg.L.Shops,
		objLookupID:   cfg.L.ShopsID,
		idToResFunc:   idToUnnamedAPIResource[seeding.Shop, Shop, UnnamedAPIResource, UnnamedAPIResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		avlFunc:       filterAvlShops,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createShopSimple,
			},
		},
	}

	e.Treasures = handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedAPIResourceList]{
		Endpoint:      epTreasures,
		resTypeSingle: rtsTreasures,
		resTypePlural: rtpTreasures,
		objLookup:     cfg.L.Treasures,
		objLookupID:   cfg.L.TreasuresID,
		idToResFunc:   idToUnnamedAPIResource[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedAPIResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	e.Quests = handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestAPIResourceList]{
		Endpoint:      epQuests,
		resTypeSingle: rtsQuests,
		resTypePlural: rtpQuests,
		objLookup:     cfg.L.Quests,
		objLookupID:   cfg.L.QuestsID,
		idToResFunc:   idToQuestAPIResource[seeding.Quest, Quest, QuestAPIResource, QuestAPIResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetQuestIDs,
		getSingleFunc: cfg.getQuest,
		retrieveFunc:  cfg.retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	e.Sidequests = handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestAPIResourceList]{
		Endpoint:      epSidequests,
		resTypeSingle: rtsSidequests,
		resTypePlural: rtpSidequests,
		objLookup:     cfg.L.Sidequests,
		objLookupID:   cfg.L.SidequestsID,
		idToResFunc:   idToQuestAPIResource[seeding.Sidequest, Sidequest, QuestAPIResource, QuestAPIResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[SectionName]Subsection{
			SnSubquests: {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	e.Subquests = handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestAPIResourceList]{
		Endpoint:      epSubquests,
		resTypeSingle: rtsSubquests,
		resTypePlural: rtpSubquests,
		objLookup:     cfg.L.Subquests,
		objLookupID:   cfg.L.SubquestsID,
		idToResFunc:   idToQuestAPIResource[seeding.Subquest, Subquest, QuestAPIResource, QuestAPIResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	e.ArenaCreations = handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epArenaCreations,
		resTypeSingle: rtsArenaCreations,
		resTypePlural: rtpArenaCreations,
		objLookup:     cfg.L.ArenaCreations,
		objLookupID:   cfg.L.ArenaCreationsID,
		idToResFunc:   idToNamedAPIResource[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetArenaCreationIDs,
		getSingleFunc: cfg.getArenaCreation,
		retrieveFunc:  cfg.retrieveArenaCreations,
	}

	e.BlitzballPrizes = handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epBlitzballPrizes,
		resTypeSingle: rtsBlitzballPrizes,
		resTypePlural: rtpBlitzballPrizes,
		objLookup:     cfg.L.Positions,
		objLookupID:   cfg.L.PositionsID,
		idToResFunc:   idToNamedAPIResource[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetBlitzballPrizeIDs,
		getSingleFunc: cfg.getBlitzballPrize,
		retrieveFunc:  cfg.retrieveBlitzballPrizes,
	}

	e.Songs = handlerInput[seeding.Song, Song, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epSongs,
		resTypeSingle: rtsSongs,
		resTypePlural: rtpSongs,
		objLookup:     cfg.L.Songs,
		objLookupID:   cfg.L.SongsID,
		idToResFunc:   idToNamedAPIResource[seeding.Song, Song, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSongIDs,
		getSingleFunc: cfg.getSong,
		retrieveFunc:  cfg.retrieveSongs,
	}

	e.FMVs = handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epFMVs,
		resTypeSingle: rtsFMVs,
		resTypePlural: rtpFMVs,
		objLookup:     cfg.L.FMVs,
		objLookupID:   cfg.L.FMVsID,
		idToResFunc:   idToNamedAPIResource[seeding.FMV, FMV, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetFmvIDs,
		getSingleFunc: cfg.getFMV,
		retrieveFunc:  cfg.retrieveFMVs,
	}

	e.PlayerUnits = handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]{
		Endpoint:      epPlayerUnits,
		resTypeSingle: rtsPlayerUnits,
		resTypePlural: rtpPlayerUnits,
		objLookup:     cfg.L.PlayerUnits,
		objLookupID:   cfg.L.PlayerUnitsID,
		idToResFunc:   idToTypedAPIResource[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetPlayerUnitIDs,
		getSingleFunc: cfg.getPlayerUnit,
		retrieveFunc:  cfg.retrievePlayerUnits,
	}

	e.Characters = handlerInput[seeding.Character, Character, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epCharacters,
		resTypeSingle: rtsCharacters,
		resTypePlural: rtpCharacters,
		objLookup:     cfg.L.Characters,
		objLookupID:   cfg.L.CharactersID,
		idToResFunc:   idToNamedAPIResource[seeding.Character, Character, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[SectionName]Subsection{
			SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			SnStdSgAbilities: {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			SnExpSgAbilities: {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			SnOverdrives: {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Aeons = handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epAeons,
		resTypeSingle: rtsAeons,
		resTypePlural: rtpAeons,
		objLookup:     cfg.L.Aeons,
		objLookupID:   cfg.L.AeonsID,
		idToResFunc:   idToNamedAPIResource[seeding.Aeon, Aeon, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[SectionName]Subsection{
			SnDefaultAbilities: {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			SnOverdrives: {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			SnStats: {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	e.CharacterClasses = handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epCharacterClasses,
		resTypeSingle: rtsCharacterClasses,
		resTypePlural: rtpCharacterClasses,
		objLookup:     cfg.L.CharClasses,
		objLookupID:   cfg.L.CharClassesID,
		idToResFunc:   idToNamedAPIResource[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[SectionName]Subsection{
			SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			SnLearnableAbilities: {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			SnDefaultOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			SnLearnableOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epMonsters,
		resTypeSingle:    rtsMonsters,
		resTypePlural:    rtpMonsters,
		objLookup:        cfg.L.Monsters,
		objLookupID:      cfg.L.MonstersID,
		idToResFunc:      idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			SnAbilities: {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			SnAreas: {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			SnMonsterFormations: {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	e.Abilities = handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]{
		Endpoint:      epAbilities,
		resTypeSingle: rtsAbilities,
		resTypePlural: rtpAbilities,
		objLookup:     cfg.L.Abilities,
		objLookupID:   cfg.L.AbilitiesID,
		idToResFunc:   idToTypedAPIResource[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.PlayerAbilities = handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epPlayerAbilities,
		resTypeSingle:    rtsPlayerAbilities,
		resTypePlural:    rtpPlayerAbilities,
		objLookup:        cfg.L.PlayerAbilities,
		objLookupID:      cfg.L.PlayerAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypePlayerAbility),
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.OverdriveAbilities = handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epOverdriveAbilities,
		resTypeSingle:    rtsOverdriveAbilities,
		resTypePlural:    rtpOverdriveAbilities,
		objLookup:        cfg.L.OverdriveAbilities,
		objLookupID:      cfg.L.OverdriveAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeOverdriveAbility),
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.ItemAbilities = handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epItemAbilities,
		resTypeSingle: rtsItemAbilities,
		resTypePlural: rtpItemAbilities,
		objLookup:     cfg.L.ItemAbilities,
		objLookupID:   cfg.L.ItemAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	e.TriggerCommands = handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epTriggerCommands,
		resTypeSingle:    rtsTriggerCommands,
		resTypePlural:    rtpTriggerCommands,
		objLookup:        cfg.L.TriggerCommands,
		objLookupID:      cfg.L.TriggerCommandsID,
		idToResFunc:      idToNamedAPIResource[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeTriggerCommand),
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	e.MiscAbilities = handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epMiscAbilities,
		resTypeSingle:    rtsMiscAbilities,
		resTypePlural:    rtpMiscAbilities,
		objLookup:        cfg.L.MiscAbilities,
		objLookupID:      cfg.L.MiscAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeMiscAbility),
		retrieveQuery:    cfg.db.GetMiscAbilityIDs,
		getSingleFunc:    cfg.getMiscAbility,
		retrieveFunc:     cfg.retrieveMiscAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createMiscAbilitySimple,
			},
		},
	}

	e.EnemyAbilities = handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:         epEnemyAbilities,
		resTypeSingle:    rtsEnemyAbilities,
		resTypePlural:    rtpEnemyAbilities,
		objLookup:        cfg.L.EnemyAbilities,
		objLookupID:      cfg.L.EnemyAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeEnemyAbility),
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			SnMonsters: {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.AeonCommands = handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epAeonCommands,
		resTypeSingle: rtsAeonCommands,
		resTypePlural: rtpAeonCommands,
		objLookup:     cfg.L.AeonCommands,
		objLookupID:   cfg.L.AeonCommandsID,
		idToResFunc:   idToNamedAPIResource[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonCommandIDs,
		getSingleFunc: cfg.getAeonCommand,
		retrieveFunc:  cfg.retrieveAeonCommands,
	}

	e.OverdriveCommands = handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epOverdriveCommands,
		resTypeSingle: rtsOverdriveCommands,
		resTypePlural: rtpOverdriveCommands,
		objLookup:     cfg.L.OverdriveCommands,
		objLookupID:   cfg.L.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[SectionName]Subsection{
			SnOverdriveAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			SnOverdrives: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs),
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Overdrives = handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epOverdrives,
		resTypeSingle: rtsOverdrives,
		resTypePlural: rtpOverdrives,
		objLookup:     cfg.L.Overdrives,
		objLookupID:   cfg.L.OverdrivesID,
		idToResFunc:   idToNamedAPIResource[seeding.Overdrive, Overdrive, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createOverdriveSimple,
			},
			SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.RonsoRages = handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epRonsoRages,
		resTypeSingle: rtsRonsoRages,
		resTypePlural: rtpRonsoRages,
		objLookup:     cfg.L.RonsoRages,
		objLookupID:   cfg.L.RonsoRagesID,
		idToResFunc:   idToNamedAPIResource[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[SectionName]Subsection{
			SnMonsters: {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.Submenus = handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epSubmenus,
		resTypeSingle: rtsSubmenus,
		resTypePlural: rtpSubmenus,
		objLookup:     cfg.L.Submenus,
		objLookupID:   cfg.L.SubmenusID,
		idToResFunc:   idToNamedAPIResource[seeding.Submenu, Submenu, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[SectionName]Subsection{
			SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.Topmenus = handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epTopmenus,
		resTypeSingle: rtsTopmenus,
		resTypePlural: rtpTopmenus,
		objLookup:     cfg.L.Topmenus,
		objLookupID:   cfg.L.TopmenusID,
		idToResFunc:   idToNamedAPIResource[seeding.Topmenu, Topmenu, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[SectionName]Subsection{
			SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.AllItems = handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]{
		Endpoint:      epAllItems,
		resTypeSingle: rtsAllItems,
		resTypePlural: rtpAllItems,
		objLookup:     cfg.L.MasterItems,
		objLookupID:   cfg.L.MasterItemsID,
		idToResFunc:   idToTypedAPIResource[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetMasterItemIDs,
		getSingleFunc: cfg.getMasterItem,
		retrieveFunc:  cfg.retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	e.Items = handlerInput[seeding.Item, Item, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epItems,
		resTypeSingle: rtsItems,
		resTypePlural: rtpItems,
		objLookup:     cfg.L.Items,
		objLookupID:   cfg.L.ItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.Item, Item, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemIDs,
		getSingleFunc: cfg.getItem,
		retrieveFunc:  cfg.retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[SectionName]Subsection{
			SnMixes: {
				dbQuery:     cfg.db.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	e.Spheres = handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epSpheres,
		resTypeSingle: rtsSpheres,
		resTypePlural: rtpSpheres,
		objLookup:     cfg.L.Spheres,
		objLookupID:   cfg.L.SpheresID,
		idToResFunc:   idToNamedAPIResource[seeding.Sphere, Sphere, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSphereIDs,
		getSingleFunc: cfg.getSphere,
		retrieveFunc:  cfg.retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	e.KeyItems = handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epKeyItems,
		resTypeSingle: rtsKeyItems,
		resTypePlural: rtpKeyItems,
		objLookup:     cfg.L.KeyItems,
		objLookupID:   cfg.L.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seeding.KeyItem, KeyItem, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetKeyItemIDs,
		getSingleFunc: cfg.getKeyItem,
		retrieveFunc:  cfg.retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	e.Primers = handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epPrimers,
		resTypeSingle: rtsPrimers,
		resTypePlural: rtpPrimers,
		objLookup:     cfg.L.Primers,
		objLookupID:   cfg.L.PrimersID,
		idToResFunc:   idToNamedAPIResource[seeding.Primer, Primer, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetPrimerIDs,
		getSingleFunc: cfg.getPrimer,
		retrieveFunc:  cfg.retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	e.Mixes = handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epMixes,
		resTypeSingle: rtsMixes,
		resTypePlural: rtpMixes,
		objLookup:     cfg.L.Mixes,
		objLookupID:   cfg.L.MixesID,
		idToResFunc:   idToNamedAPIResource[seeding.Mix, Mix, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetMixIDs,
		getSingleFunc: cfg.getMix,
		retrieveFunc:  cfg.retrieveMixes,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createMixSimple,
			},
		},
	}

	e.AutoAbilities = handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epAutoAbilities,
		resTypeSingle: rtsAutoAbilities,
		resTypePlural: rtpAutoAbilities,
		objLookup:     cfg.L.AutoAbilities,
		objLookupID:   cfg.L.AutoAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAutoAbilityIDs,
		getSingleFunc: cfg.getAutoAbility,
		retrieveFunc:  cfg.retrieveAutoAbilities,
		avlFunc:       filterAvlAutoAbilities,
		subsections: map[SectionName]Subsection{
			SnSimple: {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.EquipmentTables = handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedAPIResourceList]{
		Endpoint:      epEquipmentTables,
		resTypeSingle: rtsEquipmentTables,
		resTypePlural: rtpEquipmentTables,
		objLookup:     cfg.L.EquipmentTables,
		objLookupID:   cfg.L.EquipmentTablesID,
		idToResFunc:   idToUnnamedAPIResource[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedAPIResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetEquipmentTableIDs,
		getSingleFunc: cfg.getEquipmentTable,
		retrieveFunc:  cfg.retrieveEquipmentTables,
	}

	e.Equipment = handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epEquipment,
		resTypeSingle: rtsEquipment,
		resTypePlural: rtpEquipment,
		objLookup:     cfg.L.EquipmentNames,
		objLookupID:   cfg.L.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetEquipmentIDs,
		getSingleFunc: cfg.getEquipment,
		retrieveFunc:  cfg.retrieveEquipment,
	}

	e.CelestialWeapons = handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epCelestialWeapons,
		resTypeSingle: rtsCelestialWeapons,
		resTypePlural: rtpCelestialWeapons,
		objLookup:     cfg.L.CelestialWeapons,
		objLookupID:   cfg.L.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCelestialWeaponIDs,
		getSingleFunc: cfg.getCelestialWeapon,
		retrieveFunc:  cfg.retrieveCelestialWeapons,
		subsections: map[SectionName]Subsection{
			SnAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(cfg),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.Stats = handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epStats,
		resTypeSingle: rtsStats,
		resTypePlural: rtpStats,
		objLookup:     cfg.L.Stats,
		objLookupID:   cfg.L.StatsID,
		idToResFunc:   idToNamedAPIResource[seeding.Stat, Stat, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetStatIDs,
		getSingleFunc: cfg.getStat,
		retrieveFunc:  cfg.retrieveStats,
	}

	e.Properties = handlerInput[seeding.Property, Property, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epProperties,
		resTypeSingle: rtsProperties,
		resTypePlural: rtpProperties,
		objLookup:     cfg.L.Properties,
		objLookupID:   cfg.L.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seeding.Property, Property, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetPropertyIDs,
		getSingleFunc: cfg.getProperty,
		retrieveFunc:  cfg.retrieveProperties,
	}

	e.OverdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epOverdriveModes,
		resTypeSingle: rtsOverdriveModes,
		resTypePlural: rtpOverdriveModes,
		objLookup:     cfg.L.OverdriveModes,
		objLookupID:   cfg.L.OverdriveModesID,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveModeIDs,
		getSingleFunc: cfg.getOverdriveMode,
		retrieveFunc:  cfg.retrieveOverdriveModes,
	}

	e.Elements = handlerInput[seeding.Element, Element, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epElements,
		resTypeSingle: rtsElements,
		resTypePlural: rtpElements,
		objLookup:     cfg.L.Elements,
		objLookupID:   cfg.L.ElementsID,
		idToResFunc:   idToNamedAPIResource[seeding.Element, Element, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetElementIDs,
		getSingleFunc: cfg.getElement,
		retrieveFunc:  cfg.retrieveElements,
	}

	e.StatusConditions = handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epStatusConditions,
		resTypeSingle: rtsStatusConditions,
		resTypePlural: rtpStatusConditions,
		objLookup:     cfg.L.StatusConditions,
		objLookupID:   cfg.L.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetStatusConditionIDs,
		getSingleFunc: cfg.getStatusCondition,
		retrieveFunc:  cfg.retrieveStatusConditions,
	}

	e.Modifiers = handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedAPIResourceList]{
		Endpoint:      epModifiers,
		resTypeSingle: rtsModifiers,
		resTypePlural: rtpModifiers,
		objLookup:     cfg.L.Modifiers,
		objLookupID:   cfg.L.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seeding.Modifier, Modifier, NamedAPIResource, NamedAPIResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetModifierIDs,
		getSingleFunc: cfg.getModifier,
		retrieveFunc:  cfg.retrieveModifiers,
	}

	e.AgilityTiers = handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedAPIResourceList]{
		Endpoint:      epAgilityTiers,
		resTypeSingle: rtsAgilityTiers,
		resTypePlural: rtpAgilityTiers,
		objLookup:     nil,
		objLookupID:   cfg.L.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedAPIResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetAgilityTierIDs,
		getSingleFunc: cfg.getAgilityTier,
		retrieveFunc:  cfg.retrieveAgilityTiers,
	}

	e.Enums = handlerInputEnums{
		Endpoint: epEnums,
		usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		lookup:   cfg.T.Lookup,
	}

	e.Endpoints = handlerInputEndpoints{
		Endpoint: epEndpoints,
		usage:    []string{"/api/endpoints"},
	}

	e.AlBhed = handlerInputService[AlBhedParams, AlBhedResponse]{
		Endpoint:  epAlBhed,
		verifyFn:  verifyAlBhedParams,
		executeFn: handleAlBhed,
		docsFn:    getAlBhedParamsDoc,
		samplesFn: getAlBhedParamsSamples,
	}

	e.TurnOrder = handlerInputService[TurnOrderParams, TurnOrderResponse]{
		Endpoint:  epTurnOrder,
		verifyFn:  verifyTurnOrderParams,
		executeFn: handleTurnOrder,
		docsFn:    getTurnOrderParamsDoc,
		samplesFn: getTurnOrderParamsSamples,
	}

	e.DropChance = handlerInputService[DropChanceParams, DropChanceResponse]{
		Endpoint:  epDropChance,
		verifyFn:  verifyDropChanceParams,
		executeFn: handleDropChance,
		docsFn:    getDropChanceParamsDoc,
		samplesFn: getDropChanceParamsSamples,
	}

	e.CritChance = handlerInputService[CritChanceParams, CritChanceResponse]{
		Endpoint:  epCritChance,
		verifyFn:  verifyCritChanceParams,
		executeFn: handleCritChance,
		docsFn:    getCritChanceParamsDoc,
		samplesFn: getCritChanceParamsSamples,
	}

	cfg.E = &e
	e.Endpoints.slice = e.InitializeEndpointSlice()

}

func (e endpoints) InitializeEndpointSlice() []EndpointName {
	endpoints := []EndpointName{
		e.Locations.Endpoint,
		e.Sublocations.Endpoint,
		e.Areas.Endpoint,
		e.MonsterFormations.Endpoint,
		e.Shops.Endpoint,
		e.Treasures.Endpoint,
		e.Quests.Endpoint,
		e.Sidequests.Endpoint,
		e.Subquests.Endpoint,
		e.ArenaCreations.Endpoint,
		e.BlitzballPrizes.Endpoint,
		e.Songs.Endpoint,
		e.FMVs.Endpoint,
		e.PlayerUnits.Endpoint,
		e.Characters.Endpoint,
		e.Aeons.Endpoint,
		e.CharacterClasses.Endpoint,
		e.Monsters.Endpoint,
		e.Abilities.Endpoint,
		e.PlayerAbilities.Endpoint,
		e.OverdriveAbilities.Endpoint,
		e.ItemAbilities.Endpoint,
		e.TriggerCommands.Endpoint,
		e.MiscAbilities.Endpoint,
		e.EnemyAbilities.Endpoint,
		e.AeonCommands.Endpoint,
		e.OverdriveCommands.Endpoint,
		e.Overdrives.Endpoint,
		e.RonsoRages.Endpoint,
		e.Submenus.Endpoint,
		e.Topmenus.Endpoint,
		e.AllItems.Endpoint,
		e.Items.Endpoint,
		e.Spheres.Endpoint,
		e.KeyItems.Endpoint,
		e.Primers.Endpoint,
		e.Mixes.Endpoint,
		e.AutoAbilities.Endpoint,
		e.EquipmentTables.Endpoint,
		e.Equipment.Endpoint,
		e.CelestialWeapons.Endpoint,
		e.Stats.Endpoint,
		e.Properties.Endpoint,
		e.OverdriveModes.Endpoint,
		e.Elements.Endpoint,
		e.StatusConditions.Endpoint,
		e.Modifiers.Endpoint,
		e.AgilityTiers.Endpoint,
		e.Enums.Endpoint,
		e.Endpoints.Endpoint,
		e.AlBhed.Endpoint,
		e.TurnOrder.Endpoint,
		e.DropChance.Endpoint,
		e.CritChance.Endpoint,
	}

	slices.Sort(endpoints)
	return endpoints
}
