package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type endpoints struct {
	Locations          handlerInput[seedtypes.Location, Location, NamedAPIResource]
	Sublocations       handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]
	Areas              handlerInput[seedtypes.Area, Area, AreaAPIResource]
	MonsterFormations  handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource]
	Shops              handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]
	Treasures          handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]
	Quests             handlerInput[seedtypes.Quest, Quest, QuestAPIResource]
	Sidequests         handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]
	Subquests          handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]
	ArenaCreations     handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]
	BlitzballPrizes    handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource]
	Songs              handlerInput[seedtypes.Song, Song, NamedAPIResource]
	FMVs               handlerInput[seedtypes.FMV, FMV, NamedAPIResource]
	PlayerUnits        handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource]
	Characters         handlerInput[seedtypes.Character, Character, NamedAPIResource]
	Aeons              handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]
	CharacterClasses   handlerInput[seedtypes.CharacterClass, CharacterClass, NamedAPIResource]
	Monsters           handlerInput[seedtypes.Monster, Monster, NamedAPIResource]
	Abilities          handlerInput[seedtypes.Ability, Ability, TypedAPIResource]
	PlayerAbilities    handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]
	OverdriveAbilities handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]
	ItemAbilities      handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]
	TriggerCommands    handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]
	MiscAbilities      handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]
	EnemyAbilities     handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]
	AeonCommands       handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource]
	OverdriveCommands  handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]
	Overdrives         handlerInput[seedtypes.Overdrive, Overdrive, NamedAPIResource]
	RonsoRages         handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]
	Submenus           handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource]
	Topmenus           handlerInput[seedtypes.Topmenu, Topmenu, NamedAPIResource]
	AllItems           handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource]
	Items              handlerInput[seedtypes.Item, Item, NamedAPIResource]
	KeyItems           handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]
	Spheres            handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource]
	Primers            handlerInput[seedtypes.Primer, Primer, NamedAPIResource]
	Mixes              handlerInput[seedtypes.Mix, Mix, NamedAPIResource]
	AutoAbilities      handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]
	EquipmentTables    handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]
	Equipment          handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]
	CelestialWeapons   handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource]
	Stats              handlerInput[seedtypes.Stat, Stat, NamedAPIResource]
	Properties         handlerInput[seedtypes.Property, Property, NamedAPIResource]
	OverdriveModes     handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource]
	Elements           handlerInput[seedtypes.Element, Element, NamedAPIResource]
	StatusConditions   handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]
	Modifiers          handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource]
	AgilityTiers       handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]
	Enums              handlerInputEnums
	Endpoints          handlerInputEndpoints
	AlBhed             handlerInputService[AlBhedParams, AlBhedResponse]
	TurnOrder          handlerInputService[TurnOrderParams, TurnOrderResponse]
	DropChance         handlerInputService[DropChanceParams, DropChanceResponse]
	CritChance         handlerInputService[CritChanceParams, CritChanceResponse]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.Locations = handlerInput[seedtypes.Location, Location, NamedAPIResource]{
		Endpoint:      typestrings.EpLocations,
		resTypeSingle: typestrings.RTSLocations,
		resTypePlural: typestrings.RTPLocations,
		objLookup:     lookup.L.Locations,
		objLookupID:   lookup.L.LocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Location, Location, NamedAPIResource],
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			typestrings.SnConnected: {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			typestrings.SnSublocations: {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			typestrings.SnAreas: {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			typestrings.SnMonsterFormations: {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			typestrings.SnShops: {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			typestrings.SnSongs: {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			typestrings.SnTreasures: {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.Sublocations = handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]{
		Endpoint:      typestrings.EpSublocations,
		resTypeSingle: typestrings.RTSSublocations,
		resTypePlural: typestrings.RTPSublocations,
		objLookup:     lookup.L.Sublocations,
		objLookupID:   lookup.L.SublocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sublocation, Sublocation, NamedAPIResource],
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			typestrings.SnConnected: {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			typestrings.SnAreas: {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			typestrings.SnMonsterFormations: {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			typestrings.SnShops: {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			typestrings.SnSongs: {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			typestrings.SnTreasures: {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.Areas = handlerInput[seedtypes.Area, Area, AreaAPIResource]{
		Endpoint:      typestrings.EpAreas,
		resTypeSingle: typestrings.RTSAreas,
		resTypePlural: typestrings.RTPAreas,
		objLookup:     lookup.L.Areas,
		objLookupID:   lookup.L.AreasID,
		idToResFunc:   idToAreaAPIResource,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			typestrings.SnConnected: {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			typestrings.SnMonsterFormations: {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			typestrings.SnSongs: {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			typestrings.SnTreasures: {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.MonsterFormations = handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource]{
		Endpoint:      typestrings.EpMonsterFormations,
		resTypeSingle: typestrings.RTSMonsterFormations,
		resTypePlural: typestrings.RTPMonsterFormations,
		objLookup:     lookup.L.MonsterFormations,
		objLookupID:   lookup.L.MonsterFormationsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.Shops = handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]{
		Endpoint:      typestrings.EpShops,
		resTypeSingle: typestrings.RTSShops,
		resTypePlural: typestrings.RTPShops,
		objLookup:     lookup.L.Shops,
		objLookupID:   lookup.L.ShopsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Shop, Shop, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		avlFunc:       filterAvlShops,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createShopSimple,
			},
		},
	}

	e.Treasures = handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]{
		Endpoint:      typestrings.EpTreasures,
		resTypeSingle: typestrings.RTSTreasures,
		resTypePlural: typestrings.RTPTreasures,
		objLookup:     lookup.L.Treasures,
		objLookupID:   lookup.L.TreasuresID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Treasure, Treasure, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	e.Quests = handlerInput[seedtypes.Quest, Quest, QuestAPIResource]{
		Endpoint:      typestrings.EpQuests,
		resTypeSingle: typestrings.RTSQuests,
		resTypePlural: typestrings.RTPQuests,
		objLookup:     lookup.L.Quests,
		objLookupID:   lookup.L.QuestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Quest, Quest, QuestAPIResource],
		retrieveQuery: cfg.db.GetQuestIDs,
		getSingleFunc: cfg.getQuest,
		retrieveFunc:  cfg.retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	e.Sidequests = handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]{
		Endpoint:      typestrings.EpSidequests,
		resTypeSingle: typestrings.RTSSidequests,
		resTypePlural: typestrings.RTPSidequests,
		objLookup:     lookup.L.Sidequests,
		objLookupID:   lookup.L.SidequestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Sidequest, Sidequest, QuestAPIResource],
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSubquests: {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	e.Subquests = handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]{
		Endpoint:      typestrings.EpSubquests,
		resTypeSingle: typestrings.RTSSubquests,
		resTypePlural: typestrings.RTPSubquests,
		objLookup:     lookup.L.Subquests,
		objLookupID:   lookup.L.SubquestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Subquest, Subquest, QuestAPIResource],
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	e.ArenaCreations = handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]{
		Endpoint:      typestrings.EpArenaCreations,
		resTypeSingle: typestrings.RTSArenaCreations,
		resTypePlural: typestrings.RTPArenaCreations,
		objLookup:     lookup.L.ArenaCreations,
		objLookupID:   lookup.L.ArenaCreationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource],
		retrieveQuery: cfg.db.GetArenaCreationIDs,
		getSingleFunc: cfg.getArenaCreation,
		retrieveFunc:  cfg.retrieveArenaCreations,
	}

	e.BlitzballPrizes = handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource]{
		Endpoint:      typestrings.EpBlitzballPrizes,
		resTypeSingle: typestrings.RTSBlitzballPrizes,
		resTypePlural: typestrings.RTPBlitzballPrizes,
		objLookup:     lookup.L.Positions,
		objLookupID:   lookup.L.PositionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource],
		retrieveQuery: cfg.db.GetBlitzballPrizeIDs,
		getSingleFunc: cfg.getBlitzballPrize,
		retrieveFunc:  cfg.retrieveBlitzballPrizes,
	}

	e.Songs = handlerInput[seedtypes.Song, Song, NamedAPIResource]{
		Endpoint:      typestrings.EpSongs,
		resTypeSingle: typestrings.RTSSongs,
		resTypePlural: typestrings.RTPSongs,
		objLookup:     lookup.L.Songs,
		objLookupID:   lookup.L.SongsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Song, Song, NamedAPIResource],
		retrieveQuery: cfg.db.GetSongIDs,
		getSingleFunc: cfg.getSong,
		retrieveFunc:  cfg.retrieveSongs,
	}

	e.FMVs = handlerInput[seedtypes.FMV, FMV, NamedAPIResource]{
		Endpoint:      typestrings.EpFMVs,
		resTypeSingle: typestrings.RTSFMVs,
		resTypePlural: typestrings.RTPFMVs,
		objLookup:     lookup.L.FMVs,
		objLookupID:   lookup.L.FMVsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.FMV, FMV, NamedAPIResource],
		retrieveQuery: cfg.db.GetFmvIDs,
		getSingleFunc: cfg.getFMV,
		retrieveFunc:  cfg.retrieveFMVs,
	}

	e.PlayerUnits = handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource]{
		Endpoint:      typestrings.EpPlayerUnits,
		resTypeSingle: typestrings.RTSPlayerUnits,
		resTypePlural: typestrings.RTPPlayerUnits,
		objLookup:     lookup.L.PlayerUnits,
		objLookupID:   lookup.L.PlayerUnitsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource],
		retrieveQuery: cfg.db.GetPlayerUnitIDs,
		getSingleFunc: cfg.getPlayerUnit,
		retrieveFunc:  cfg.retrievePlayerUnits,
	}

	e.Characters = handlerInput[seedtypes.Character, Character, NamedAPIResource]{
		Endpoint:      typestrings.EpCharacters,
		resTypeSingle: typestrings.RTSCharacters,
		resTypePlural: typestrings.RTPCharacters,
		objLookup:     lookup.L.Characters,
		objLookupID:   lookup.L.CharactersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Character, Character, NamedAPIResource],
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			typestrings.SnStdSgAbilities: {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			typestrings.SnExpSgAbilities: {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			typestrings.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			typestrings.SnOverdrives: {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Aeons = handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]{
		Endpoint:      typestrings.EpAeons,
		resTypeSingle: typestrings.RTSAeons,
		resTypePlural: typestrings.RTPAeons,
		objLookup:     lookup.L.Aeons,
		objLookupID:   lookup.L.AeonsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Aeon, Aeon, NamedAPIResource],
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			typestrings.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			typestrings.SnOverdrives: {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			typestrings.SnStats: {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	e.CharacterClasses = handlerInput[seedtypes.CharacterClass, CharacterClass, NamedAPIResource]{
		Endpoint:      typestrings.EpCharacterClasses,
		resTypeSingle: typestrings.RTSCharacterClasses,
		resTypePlural: typestrings.RTPCharacterClasses,
		objLookup:     lookup.L.CharClasses,
		objLookupID:   lookup.L.CharClassesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CharacterClass, CharacterClass, NamedAPIResource],
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			typestrings.SnLearnableAbilities: {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			typestrings.SnDefaultOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			typestrings.SnLearnableOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Monsters = handlerInput[seedtypes.Monster, Monster, NamedAPIResource]{
		Endpoint:         typestrings.EpMonsters,
		resTypeSingle:    typestrings.RTSMonsters,
		resTypePlural:    typestrings.RTPMonsters,
		objLookup:        lookup.L.Monsters,
		objLookupID:      lookup.L.MonstersID,
		idToResFunc:      idToNamedAPIResource[seedtypes.Monster, Monster, NamedAPIResource],
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			typestrings.SnAbilities: {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			typestrings.SnAreas: {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			typestrings.SnMonsterFormations: {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	e.Abilities = handlerInput[seedtypes.Ability, Ability, TypedAPIResource]{
		Endpoint:      typestrings.EpAbilities,
		resTypeSingle: typestrings.RTSAbilities,
		resTypePlural: typestrings.RTPAbilities,
		objLookup:     lookup.L.Abilities,
		objLookupID:   lookup.L.AbilitiesID,
		idToResFunc:   idToTypedAPIResource[seedtypes.Ability, Ability, TypedAPIResource],
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.PlayerAbilities = handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]{
		Endpoint:         typestrings.EpPlayerAbilities,
		resTypeSingle:    typestrings.RTSPlayerAbilities,
		resTypePlural:    typestrings.RTPPlayerAbilities,
		objLookup:        lookup.L.PlayerAbilities,
		objLookupID:      lookup.L.PlayerAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypePlayerAbility),
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.OverdriveAbilities = handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]{
		Endpoint:         typestrings.EpOverdriveAbilities,
		resTypeSingle:    typestrings.RTSOverdriveAbilities,
		resTypePlural:    typestrings.RTPOverdriveAbilities,
		objLookup:        lookup.L.OverdriveAbilities,
		objLookupID:      lookup.L.OverdriveAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeOverdriveAbility),
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.ItemAbilities = handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]{
		Endpoint:      typestrings.EpItemAbilities,
		resTypeSingle: typestrings.RTSItemAbilities,
		resTypePlural: typestrings.RTPItemAbilities,
		objLookup:     lookup.L.ItemAbilities,
		objLookupID:   lookup.L.ItemAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ItemAbility, ItemAbility, NamedAPIResource],
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	e.TriggerCommands = handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]{
		Endpoint:         typestrings.EpTriggerCommands,
		resTypeSingle:    typestrings.RTSTriggerCommands,
		resTypePlural:    typestrings.RTPTriggerCommands,
		objLookup:        lookup.L.TriggerCommands,
		objLookupID:      lookup.L.TriggerCommandsID,
		idToResFunc:      idToNamedAPIResource[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeTriggerCommand),
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	e.MiscAbilities = handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]{
		Endpoint:         typestrings.EpMiscAbilities,
		resTypeSingle:    typestrings.RTSMiscAbilities,
		resTypePlural:    typestrings.RTPMiscAbilities,
		objLookup:        lookup.L.MiscAbilities,
		objLookupID:      lookup.L.MiscAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.MiscAbility, MiscAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeMiscAbility),
		retrieveQuery:    cfg.db.GetMiscAbilityIDs,
		getSingleFunc:    cfg.getMiscAbility,
		retrieveFunc:     cfg.retrieveMiscAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createMiscAbilitySimple,
			},
		},
	}

	e.EnemyAbilities = handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]{
		Endpoint:         typestrings.EpEnemyAbilities,
		resTypeSingle:    typestrings.RTSEnemyAbilities,
		resTypePlural:    typestrings.RTPEnemyAbilities,
		objLookup:        lookup.L.EnemyAbilities,
		objLookupID:      lookup.L.EnemyAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeEnemyAbility),
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.AeonCommands = handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource]{
		Endpoint:      typestrings.EpAeonCommands,
		resTypeSingle: typestrings.RTSAeonCommands,
		resTypePlural: typestrings.RTPAeonCommands,
		objLookup:     lookup.L.AeonCommands,
		objLookupID:   lookup.L.AeonCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AeonCommand, AeonCommand, NamedAPIResource],
		retrieveQuery: cfg.db.GetAeonCommandIDs,
		getSingleFunc: cfg.getAeonCommand,
		retrieveFunc:  cfg.retrieveAeonCommands,
	}

	e.OverdriveCommands = handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]{
		Endpoint:      typestrings.EpOverdriveCommands,
		resTypeSingle: typestrings.RTSOverdriveCommands,
		resTypePlural: typestrings.RTPOverdriveCommands,
		objLookup:     lookup.L.OverdriveCommands,
		objLookupID:   lookup.L.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnOverdriveAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			typestrings.SnOverdrives: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs),
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.Overdrives = handlerInput[seedtypes.Overdrive, Overdrive, NamedAPIResource]{
		Endpoint:      typestrings.EpOverdrives,
		resTypeSingle: typestrings.RTSOverdrives,
		resTypePlural: typestrings.RTPOverdrives,
		objLookup:     lookup.L.Overdrives,
		objLookupID:   lookup.L.OverdrivesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Overdrive, Overdrive, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createOverdriveSimple,
			},
			typestrings.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.RonsoRages = handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]{
		Endpoint:      typestrings.EpRonsoRages,
		resTypeSingle: typestrings.RTSRonsoRages,
		resTypePlural: typestrings.RTPRonsoRages,
		objLookup:     lookup.L.RonsoRages,
		objLookupID:   lookup.L.RonsoRagesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.RonsoRage, RonsoRage, NamedAPIResource],
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnMonsters: {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.Submenus = handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource]{
		Endpoint:      typestrings.EpSubmenus,
		resTypeSingle: typestrings.RTSSubmenus,
		resTypePlural: typestrings.RTPSubmenus,
		objLookup:     lookup.L.Submenus,
		objLookupID:   lookup.L.SubmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Submenu, Submenu, NamedAPIResource],
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.Topmenus = handlerInput[seedtypes.Topmenu, Topmenu, NamedAPIResource]{
		Endpoint:      typestrings.EpTopmenus,
		resTypeSingle: typestrings.RTSTopmenus,
		resTypePlural: typestrings.RTPTopmenus,
		objLookup:     lookup.L.Topmenus,
		objLookupID:   lookup.L.TopmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Topmenu, Topmenu, NamedAPIResource],
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.AllItems = handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource]{
		Endpoint:      typestrings.EpAllItems,
		resTypeSingle: typestrings.RTSAllItems,
		resTypePlural: typestrings.RTPAllItems,
		objLookup:     lookup.L.MasterItems,
		objLookupID:   lookup.L.MasterItemsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.MasterItem, MasterItem, TypedAPIResource],
		retrieveQuery: cfg.db.GetMasterItemIDs,
		getSingleFunc: cfg.getMasterItem,
		retrieveFunc:  cfg.retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	e.Items = handlerInput[seedtypes.Item, Item, NamedAPIResource]{
		Endpoint:      typestrings.EpItems,
		resTypeSingle: typestrings.RTSItems,
		resTypePlural: typestrings.RTPItems,
		objLookup:     lookup.L.Items,
		objLookupID:   lookup.L.ItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Item, Item, NamedAPIResource],
		retrieveQuery: cfg.db.GetItemIDs,
		getSingleFunc: cfg.getItem,
		retrieveFunc:  cfg.retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnMixes: {
				dbQuery:     cfg.db.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	e.Spheres = handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource]{
		Endpoint:      typestrings.EpSpheres,
		resTypeSingle: typestrings.RTSSpheres,
		resTypePlural: typestrings.RTPSpheres,
		objLookup:     lookup.L.Spheres,
		objLookupID:   lookup.L.SpheresID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sphere, Sphere, NamedAPIResource],
		retrieveQuery: cfg.db.GetSphereIDs,
		getSingleFunc: cfg.getSphere,
		retrieveFunc:  cfg.retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	e.KeyItems = handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]{
		Endpoint:      typestrings.EpKeyItems,
		resTypeSingle: typestrings.RTSKeyItems,
		resTypePlural: typestrings.RTPKeyItems,
		objLookup:     lookup.L.KeyItems,
		objLookupID:   lookup.L.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.KeyItem, KeyItem, NamedAPIResource],
		retrieveQuery: cfg.db.GetKeyItemIDs,
		getSingleFunc: cfg.getKeyItem,
		retrieveFunc:  cfg.retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	e.Primers = handlerInput[seedtypes.Primer, Primer, NamedAPIResource]{
		Endpoint:      typestrings.EpPrimers,
		resTypeSingle: typestrings.RTSPrimers,
		resTypePlural: typestrings.RTPPrimers,
		objLookup:     lookup.L.Primers,
		objLookupID:   lookup.L.PrimersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Primer, Primer, NamedAPIResource],
		retrieveQuery: cfg.db.GetPrimerIDs,
		getSingleFunc: cfg.getPrimer,
		retrieveFunc:  cfg.retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	e.Mixes = handlerInput[seedtypes.Mix, Mix, NamedAPIResource]{
		Endpoint:      typestrings.EpMixes,
		resTypeSingle: typestrings.RTSMixes,
		resTypePlural: typestrings.RTPMixes,
		objLookup:     lookup.L.Mixes,
		objLookupID:   lookup.L.MixesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Mix, Mix, NamedAPIResource],
		retrieveQuery: cfg.db.GetMixIDs,
		getSingleFunc: cfg.getMix,
		retrieveFunc:  cfg.retrieveMixes,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createMixSimple,
			},
		},
	}

	e.AutoAbilities = handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]{
		Endpoint:      typestrings.EpAutoAbilities,
		resTypeSingle: typestrings.RTSAutoAbilities,
		resTypePlural: typestrings.RTPAutoAbilities,
		objLookup:     lookup.L.AutoAbilities,
		objLookupID:   lookup.L.AutoAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AutoAbility, AutoAbility, NamedAPIResource],
		retrieveQuery: cfg.db.GetAutoAbilityIDs,
		getSingleFunc: cfg.getAutoAbility,
		retrieveFunc:  cfg.retrieveAutoAbilities,
		avlFunc:       filterAvlAutoAbilities,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnSimple: {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.EquipmentTables = handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]{
		Endpoint:      typestrings.EpEquipmentTables,
		resTypeSingle: typestrings.RTSEquipmentTables,
		resTypePlural: typestrings.RTPEquipmentTables,
		objLookup:     lookup.L.EquipmentTables,
		objLookupID:   lookup.L.EquipmentTablesID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetEquipmentTableIDs,
		getSingleFunc: cfg.getEquipmentTable,
		retrieveFunc:  cfg.retrieveEquipmentTables,
	}

	e.Equipment = handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]{
		Endpoint:      typestrings.EpEquipment,
		resTypeSingle: typestrings.RTSEquipment,
		resTypePlural: typestrings.RTPEquipment,
		objLookup:     lookup.L.EquipmentNames,
		objLookupID:   lookup.L.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.EquipmentName, EquipmentName, NamedAPIResource],
		retrieveQuery: cfg.db.GetEquipmentIDs,
		getSingleFunc: cfg.getEquipment,
		retrieveFunc:  cfg.retrieveEquipment,
	}

	e.CelestialWeapons = handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource]{
		Endpoint:      typestrings.EpCelestialWeapons,
		resTypeSingle: typestrings.RTSCelestialWeapons,
		resTypePlural: typestrings.RTPCelestialWeapons,
		objLookup:     lookup.L.CelestialWeapons,
		objLookupID:   lookup.L.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource],
		retrieveQuery: cfg.db.GetCelestialWeaponIDs,
		getSingleFunc: cfg.getCelestialWeapon,
		retrieveFunc:  cfg.retrieveCelestialWeapons,
		subsections: map[typestrings.SectionName]Subsection{
			typestrings.SnAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(cfg),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.Stats = handlerInput[seedtypes.Stat, Stat, NamedAPIResource]{
		Endpoint:      typestrings.EpStats,
		resTypeSingle: typestrings.RTSStats,
		resTypePlural: typestrings.RTPStats,
		objLookup:     lookup.L.Stats,
		objLookupID:   lookup.L.StatsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Stat, Stat, NamedAPIResource],
		retrieveQuery: cfg.db.GetStatIDs,
		getSingleFunc: cfg.getStat,
		retrieveFunc:  cfg.retrieveStats,
	}

	e.Properties = handlerInput[seedtypes.Property, Property, NamedAPIResource]{
		Endpoint:      typestrings.EpProperties,
		resTypeSingle: typestrings.RTSProperties,
		resTypePlural: typestrings.RTPProperties,
		objLookup:     lookup.L.Properties,
		objLookupID:   lookup.L.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Property, Property, NamedAPIResource],
		retrieveQuery: cfg.db.GetPropertyIDs,
		getSingleFunc: cfg.getProperty,
		retrieveFunc:  cfg.retrieveProperties,
	}

	e.OverdriveModes = handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource]{
		Endpoint:      typestrings.EpOverdriveModes,
		resTypeSingle: typestrings.RTSOverdriveModes,
		resTypePlural: typestrings.RTPOverdriveModes,
		objLookup:     lookup.L.OverdriveModes,
		objLookupID:   lookup.L.OverdriveModesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveModeIDs,
		getSingleFunc: cfg.getOverdriveMode,
		retrieveFunc:  cfg.retrieveOverdriveModes,
	}

	e.Elements = handlerInput[seedtypes.Element, Element, NamedAPIResource]{
		Endpoint:      typestrings.EpElements,
		resTypeSingle: typestrings.RTSElements,
		resTypePlural: typestrings.RTPElements,
		objLookup:     lookup.L.Elements,
		objLookupID:   lookup.L.ElementsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Element, Element, NamedAPIResource],
		retrieveQuery: cfg.db.GetElementIDs,
		getSingleFunc: cfg.getElement,
		retrieveFunc:  cfg.retrieveElements,
	}

	e.StatusConditions = handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]{
		Endpoint:      typestrings.EpStatusConditions,
		resTypeSingle: typestrings.RTSStatusConditions,
		resTypePlural: typestrings.RTPStatusConditions,
		objLookup:     lookup.L.StatusConditions,
		objLookupID:   lookup.L.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.StatusCondition, StatusCondition, NamedAPIResource],
		retrieveQuery: cfg.db.GetStatusConditionIDs,
		getSingleFunc: cfg.getStatusCondition,
		retrieveFunc:  cfg.retrieveStatusConditions,
	}

	e.Modifiers = handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource]{
		Endpoint:      typestrings.EpModifiers,
		resTypeSingle: typestrings.RTSModifiers,
		resTypePlural: typestrings.RTPModifiers,
		objLookup:     lookup.L.Modifiers,
		objLookupID:   lookup.L.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Modifier, Modifier, NamedAPIResource],
		retrieveQuery: cfg.db.GetModifierIDs,
		getSingleFunc: cfg.getModifier,
		retrieveFunc:  cfg.retrieveModifiers,
	}

	e.AgilityTiers = handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]{
		Endpoint:      typestrings.EpAgilityTiers,
		resTypeSingle: typestrings.RTSAgilityTiers,
		resTypePlural: typestrings.RTPAgilityTiers,
		objLookup:     nil,
		objLookupID:   lookup.L.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetAgilityTierIDs,
		getSingleFunc: cfg.getAgilityTier,
		retrieveFunc:  cfg.retrieveAgilityTiers,
	}

	e.Enums = handlerInputEnums{
		Endpoint: typestrings.EpEnums,
		usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		lookup:   enums.T.Lookup,
	}

	e.Endpoints = handlerInputEndpoints{
		Endpoint: typestrings.EpEndpoints,
		usage:    []string{"/api/endpoints"},
	}

	e.AlBhed = handlerInputService[AlBhedParams, AlBhedResponse]{
		Endpoint:  typestrings.EpAlBhed,
		verifyFn:  verifyAlBhedParams,
		executeFn: handleAlBhed,
		docsFn:    getAlBhedParamsDoc,
		samplesFn: getAlBhedParamsSamples,
	}

	e.TurnOrder = handlerInputService[TurnOrderParams, TurnOrderResponse]{
		Endpoint:  typestrings.EpTurnOrder,
		verifyFn:  verifyTurnOrderParams,
		executeFn: handleTurnOrder,
		docsFn:    getTurnOrderParamsDoc,
		samplesFn: getTurnOrderParamsSamples,
	}

	e.DropChance = handlerInputService[DropChanceParams, DropChanceResponse]{
		Endpoint:  typestrings.EpDropChance,
		verifyFn:  verifyDropChanceParams,
		executeFn: handleDropChance,
		docsFn:    getDropChanceParamsDoc,
		samplesFn: getDropChanceParamsSamples,
	}

	e.CritChance = handlerInputService[CritChanceParams, CritChanceResponse]{
		Endpoint:  typestrings.EpCritChance,
		verifyFn:  verifyCritChanceParams,
		executeFn: handleCritChance,
		docsFn:    getCritChanceParamsDoc,
		samplesFn: getCritChanceParamsSamples,
	}

	cfg.E = &e
	e.Endpoints.slice = e.InitializeEndpointSlice()

}

func (e endpoints) InitializeEndpointSlice() []typestrings.EndpointName {
	endpoints := []typestrings.EndpointName{
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
