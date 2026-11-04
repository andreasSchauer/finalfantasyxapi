package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type Endpoints struct {
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

var E Endpoints

func (E *Endpoints) EndpointsInit(cfg *Config) {
	E.Locations = handlerInput[seedtypes.Location, Location, NamedAPIResource]{
		Endpoint:      ts.EpLocations,
		resTypeSingle: ts.RTSLocations,
		resTypePlural: ts.RTPLocations,
		objLookup:     lookup.L.Locations,
		objLookupID:   lookup.L.LocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Location, Location, NamedAPIResource],
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnSublocations: {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.Sublocations = handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]{
		Endpoint:      ts.EpSublocations,
		resTypeSingle: ts.RTSSublocations,
		resTypePlural: ts.RTPSublocations,
		objLookup:     lookup.L.Sublocations,
		objLookupID:   lookup.L.SublocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sublocation, Sublocation, NamedAPIResource],
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			ts.SnMonsterFormations: {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.Areas = handlerInput[seedtypes.Area, Area, AreaAPIResource]{
		Endpoint:      ts.EpAreas,
		resTypeSingle: ts.RTSAreas,
		resTypePlural: ts.RTPAreas,
		objLookup:     lookup.L.Areas,
		objLookupID:   lookup.L.AreasID,
		idToResFunc:   idToAreaAPIResource,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnSongs: {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.MonsterFormations = handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource]{
		Endpoint:      ts.EpMonsterFormations,
		resTypeSingle: ts.RTSMonsterFormations,
		resTypePlural: ts.RTPMonsterFormations,
		objLookup:     lookup.L.MonsterFormations,
		objLookupID:   lookup.L.MonsterFormationsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Shops = handlerInput[seedtypes.Shop, Shop, UnnamedAPIResource]{
		Endpoint:      ts.EpShops,
		resTypeSingle: ts.RTSShops,
		resTypePlural: ts.RTPShops,
		objLookup:     lookup.L.Shops,
		objLookupID:   lookup.L.ShopsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Shop, Shop, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		avlFunc:       filterAvlShops,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createShopSimple,
			},
		},
	}

	E.Treasures = handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]{
		Endpoint:      ts.EpTreasures,
		resTypeSingle: ts.RTSTreasures,
		resTypePlural: ts.RTPTreasures,
		objLookup:     lookup.L.Treasures,
		objLookupID:   lookup.L.TreasuresID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Treasure, Treasure, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	E.Quests = handlerInput[seedtypes.Quest, Quest, QuestAPIResource]{
		Endpoint:      ts.EpQuests,
		resTypeSingle: ts.RTSQuests,
		resTypePlural: ts.RTPQuests,
		objLookup:     lookup.L.Quests,
		objLookupID:   lookup.L.QuestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Quest, Quest, QuestAPIResource],
		retrieveQuery: cfg.db.GetQuestIDs,
		getSingleFunc: cfg.getQuest,
		retrieveFunc:  cfg.retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	E.Sidequests = handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]{
		Endpoint:      ts.EpSidequests,
		resTypeSingle: ts.RTSSidequests,
		resTypePlural: ts.RTPSidequests,
		objLookup:     lookup.L.Sidequests,
		objLookupID:   lookup.L.SidequestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Sidequest, Sidequest, QuestAPIResource],
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSubquests: {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	E.Subquests = handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]{
		Endpoint:      ts.EpSubquests,
		resTypeSingle: ts.RTSSubquests,
		resTypePlural: ts.RTPSubquests,
		objLookup:     lookup.L.Subquests,
		objLookupID:   lookup.L.SubquestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Subquest, Subquest, QuestAPIResource],
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	E.ArenaCreations = handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]{
		Endpoint:      ts.EpArenaCreations,
		resTypeSingle: ts.RTSArenaCreations,
		resTypePlural: ts.RTPArenaCreations,
		objLookup:     lookup.L.ArenaCreations,
		objLookupID:   lookup.L.ArenaCreationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource],
		retrieveQuery: cfg.db.GetArenaCreationIDs,
		getSingleFunc: cfg.getArenaCreation,
		retrieveFunc:  cfg.retrieveArenaCreations,
	}

	E.BlitzballPrizes = handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource]{
		Endpoint:      ts.EpBlitzballPrizes,
		resTypeSingle: ts.RTSBlitzballPrizes,
		resTypePlural: ts.RTPBlitzballPrizes,
		objLookup:     lookup.L.Positions,
		objLookupID:   lookup.L.PositionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource],
		retrieveQuery: cfg.db.GetBlitzballPrizeIDs,
		getSingleFunc: cfg.getBlitzballPrize,
		retrieveFunc:  cfg.retrieveBlitzballPrizes,
	}

	E.Songs = handlerInput[seedtypes.Song, Song, NamedAPIResource]{
		Endpoint:      ts.EpSongs,
		resTypeSingle: ts.RTSSongs,
		resTypePlural: ts.RTPSongs,
		objLookup:     lookup.L.Songs,
		objLookupID:   lookup.L.SongsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Song, Song, NamedAPIResource],
		retrieveQuery: cfg.db.GetSongIDs,
		getSingleFunc: cfg.getSong,
		retrieveFunc:  cfg.retrieveSongs,
	}

	E.FMVs = handlerInput[seedtypes.FMV, FMV, NamedAPIResource]{
		Endpoint:      ts.EpFMVs,
		resTypeSingle: ts.RTSFMVs,
		resTypePlural: ts.RTPFMVs,
		objLookup:     lookup.L.FMVs,
		objLookupID:   lookup.L.FMVsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.FMV, FMV, NamedAPIResource],
		retrieveQuery: cfg.db.GetFmvIDs,
		getSingleFunc: cfg.getFMV,
		retrieveFunc:  cfg.retrieveFMVs,
	}

	E.PlayerUnits = handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource]{
		Endpoint:      ts.EpPlayerUnits,
		resTypeSingle: ts.RTSPlayerUnits,
		resTypePlural: ts.RTPPlayerUnits,
		objLookup:     lookup.L.PlayerUnits,
		objLookupID:   lookup.L.PlayerUnitsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource],
		retrieveQuery: cfg.db.GetPlayerUnitIDs,
		getSingleFunc: cfg.getPlayerUnit,
		retrieveFunc:  cfg.retrievePlayerUnits,
	}

	E.Characters = handlerInput[seedtypes.Character, Character, NamedAPIResource]{
		Endpoint:      ts.EpCharacters,
		resTypeSingle: ts.RTSCharacters,
		resTypePlural: ts.RTPCharacters,
		objLookup:     lookup.L.Characters,
		objLookupID:   lookup.L.CharactersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Character, Character, NamedAPIResource],
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnStdSgAbilities: {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnExpSgAbilities: {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Aeons = handlerInput[seedtypes.Aeon, Aeon, NamedAPIResource]{
		Endpoint:      ts.EpAeons,
		resTypeSingle: ts.RTSAeons,
		resTypePlural: ts.RTPAeons,
		objLookup:     lookup.L.Aeons,
		objLookupID:   lookup.L.AeonsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Aeon, Aeon, NamedAPIResource],
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			ts.SnStats: {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	E.CharacterClasses = handlerInput[seedtypes.CharacterClass, CharacterClass, NamedAPIResource]{
		Endpoint:      ts.EpCharacterClasses,
		resTypeSingle: ts.RTSCharacterClasses,
		resTypePlural: ts.RTPCharacterClasses,
		objLookup:     lookup.L.CharClasses,
		objLookupID:   lookup.L.CharClassesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CharacterClass, CharacterClass, NamedAPIResource],
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnLearnableAbilities: {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnDefaultOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			ts.SnLearnableOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Monsters = handlerInput[seedtypes.Monster, Monster, NamedAPIResource]{
		Endpoint:         ts.EpMonsters,
		resTypeSingle:    ts.RTSMonsters,
		resTypePlural:    ts.RTPMonsters,
		objLookup:        lookup.L.Monsters,
		objLookupID:      lookup.L.MonstersID,
		idToResFunc:      idToNamedAPIResource[seedtypes.Monster, Monster, NamedAPIResource],
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnAbilities: {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnAreas: {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	E.Abilities = handlerInput[seedtypes.Ability, Ability, TypedAPIResource]{
		Endpoint:      ts.EpAbilities,
		resTypeSingle: ts.RTSAbilities,
		resTypePlural: ts.RTPAbilities,
		objLookup:     lookup.L.Abilities,
		objLookupID:   lookup.L.AbilitiesID,
		idToResFunc:   idToTypedAPIResource[seedtypes.Ability, Ability, TypedAPIResource],
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.PlayerAbilities = handlerInput[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource]{
		Endpoint:         ts.EpPlayerAbilities,
		resTypeSingle:    ts.RTSPlayerAbilities,
		resTypePlural:    ts.RTPPlayerAbilities,
		objLookup:        lookup.L.PlayerAbilities,
		objLookupID:      lookup.L.PlayerAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.PlayerAbility, PlayerAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypePlayerAbility),
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.OverdriveAbilities = handlerInput[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource]{
		Endpoint:         ts.EpOverdriveAbilities,
		resTypeSingle:    ts.RTSOverdriveAbilities,
		resTypePlural:    ts.RTPOverdriveAbilities,
		objLookup:        lookup.L.OverdriveAbilities,
		objLookupID:      lookup.L.OverdriveAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.OverdriveAbility, OverdriveAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeOverdriveAbility),
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.ItemAbilities = handlerInput[seedtypes.ItemAbility, ItemAbility, NamedAPIResource]{
		Endpoint:      ts.EpItemAbilities,
		resTypeSingle: ts.RTSItemAbilities,
		resTypePlural: ts.RTPItemAbilities,
		objLookup:     lookup.L.ItemAbilities,
		objLookupID:   lookup.L.ItemAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ItemAbility, ItemAbility, NamedAPIResource],
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	E.TriggerCommands = handlerInput[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource]{
		Endpoint:         ts.EpTriggerCommands,
		resTypeSingle:    ts.RTSTriggerCommands,
		resTypePlural:    ts.RTPTriggerCommands,
		objLookup:        lookup.L.TriggerCommands,
		objLookupID:      lookup.L.TriggerCommandsID,
		idToResFunc:      idToNamedAPIResource[seedtypes.TriggerCommand, TriggerCommand, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeTriggerCommand),
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	E.MiscAbilities = handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]{
		Endpoint:         ts.EpMiscAbilities,
		resTypeSingle:    ts.RTSMiscAbilities,
		resTypePlural:    ts.RTPMiscAbilities,
		objLookup:        lookup.L.MiscAbilities,
		objLookupID:      lookup.L.MiscAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.MiscAbility, MiscAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeMiscAbility),
		retrieveQuery:    cfg.db.GetMiscAbilityIDs,
		getSingleFunc:    cfg.getMiscAbility,
		retrieveFunc:     cfg.retrieveMiscAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMiscAbilitySimple,
			},
		},
	}

	E.EnemyAbilities = handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]{
		Endpoint:         ts.EpEnemyAbilities,
		resTypeSingle:    ts.RTSEnemyAbilities,
		resTypePlural:    ts.RTPEnemyAbilities,
		objLookup:        lookup.L.EnemyAbilities,
		objLookupID:      lookup.L.EnemyAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeEnemyAbility),
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.AeonCommands = handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource]{
		Endpoint:      ts.EpAeonCommands,
		resTypeSingle: ts.RTSAeonCommands,
		resTypePlural: ts.RTPAeonCommands,
		objLookup:     lookup.L.AeonCommands,
		objLookupID:   lookup.L.AeonCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AeonCommand, AeonCommand, NamedAPIResource],
		retrieveQuery: cfg.db.GetAeonCommandIDs,
		getSingleFunc: cfg.getAeonCommand,
		retrieveFunc:  cfg.retrieveAeonCommands,
	}

	E.OverdriveCommands = handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]{
		Endpoint:      ts.EpOverdriveCommands,
		resTypeSingle: ts.RTSOverdriveCommands,
		resTypePlural: ts.RTPOverdriveCommands,
		objLookup:     lookup.L.OverdriveCommands,
		objLookupID:   lookup.L.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[ts.SectionName]Subsection{
			ts.SnOverdriveAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs),
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Overdrives = handlerInput[seedtypes.Overdrive, Overdrive, NamedAPIResource]{
		Endpoint:      ts.EpOverdrives,
		resTypeSingle: ts.RTSOverdrives,
		resTypePlural: ts.RTPOverdrives,
		objLookup:     lookup.L.Overdrives,
		objLookupID:   lookup.L.OverdrivesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Overdrive, Overdrive, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createOverdriveSimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.RonsoRages = handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]{
		Endpoint:      ts.EpRonsoRages,
		resTypeSingle: ts.RTSRonsoRages,
		resTypePlural: ts.RTPRonsoRages,
		objLookup:     lookup.L.RonsoRages,
		objLookupID:   lookup.L.RonsoRagesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.RonsoRage, RonsoRage, NamedAPIResource],
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMonsters: {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Submenus = handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource]{
		Endpoint:      ts.EpSubmenus,
		resTypeSingle: ts.RTSSubmenus,
		resTypePlural: ts.RTPSubmenus,
		objLookup:     lookup.L.Submenus,
		objLookupID:   lookup.L.SubmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Submenu, Submenu, NamedAPIResource],
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.Topmenus = handlerInput[seedtypes.Topmenu, Topmenu, NamedAPIResource]{
		Endpoint:      ts.EpTopmenus,
		resTypeSingle: ts.RTSTopmenus,
		resTypePlural: ts.RTPTopmenus,
		objLookup:     lookup.L.Topmenus,
		objLookupID:   lookup.L.TopmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Topmenu, Topmenu, NamedAPIResource],
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.AllItems = handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource]{
		Endpoint:      ts.EpAllItems,
		resTypeSingle: ts.RTSAllItems,
		resTypePlural: ts.RTPAllItems,
		objLookup:     lookup.L.MasterItems,
		objLookupID:   lookup.L.MasterItemsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.MasterItem, MasterItem, TypedAPIResource],
		retrieveQuery: cfg.db.GetMasterItemIDs,
		getSingleFunc: cfg.getMasterItem,
		retrieveFunc:  cfg.retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	E.Items = handlerInput[seedtypes.Item, Item, NamedAPIResource]{
		Endpoint:      ts.EpItems,
		resTypeSingle: ts.RTSItems,
		resTypePlural: ts.RTPItems,
		objLookup:     lookup.L.Items,
		objLookupID:   lookup.L.ItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Item, Item, NamedAPIResource],
		retrieveQuery: cfg.db.GetItemIDs,
		getSingleFunc: cfg.getItem,
		retrieveFunc:  cfg.retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMixes: {
				dbQuery:     cfg.db.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	E.Spheres = handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource]{
		Endpoint:      ts.EpSpheres,
		resTypeSingle: ts.RTSSpheres,
		resTypePlural: ts.RTPSpheres,
		objLookup:     lookup.L.Spheres,
		objLookupID:   lookup.L.SpheresID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sphere, Sphere, NamedAPIResource],
		retrieveQuery: cfg.db.GetSphereIDs,
		getSingleFunc: cfg.getSphere,
		retrieveFunc:  cfg.retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	E.KeyItems = handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]{
		Endpoint:      ts.EpKeyItems,
		resTypeSingle: ts.RTSKeyItems,
		resTypePlural: ts.RTPKeyItems,
		objLookup:     lookup.L.KeyItems,
		objLookupID:   lookup.L.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.KeyItem, KeyItem, NamedAPIResource],
		retrieveQuery: cfg.db.GetKeyItemIDs,
		getSingleFunc: cfg.getKeyItem,
		retrieveFunc:  cfg.retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	E.Primers = handlerInput[seedtypes.Primer, Primer, NamedAPIResource]{
		Endpoint:      ts.EpPrimers,
		resTypeSingle: ts.RTSPrimers,
		resTypePlural: ts.RTPPrimers,
		objLookup:     lookup.L.Primers,
		objLookupID:   lookup.L.PrimersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Primer, Primer, NamedAPIResource],
		retrieveQuery: cfg.db.GetPrimerIDs,
		getSingleFunc: cfg.getPrimer,
		retrieveFunc:  cfg.retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	E.Mixes = handlerInput[seedtypes.Mix, Mix, NamedAPIResource]{
		Endpoint:      ts.EpMixes,
		resTypeSingle: ts.RTSMixes,
		resTypePlural: ts.RTPMixes,
		objLookup:     lookup.L.Mixes,
		objLookupID:   lookup.L.MixesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Mix, Mix, NamedAPIResource],
		retrieveQuery: cfg.db.GetMixIDs,
		getSingleFunc: cfg.getMix,
		retrieveFunc:  cfg.retrieveMixes,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMixSimple,
			},
		},
	}

	E.AutoAbilities = handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]{
		Endpoint:      ts.EpAutoAbilities,
		resTypeSingle: ts.RTSAutoAbilities,
		resTypePlural: ts.RTPAutoAbilities,
		objLookup:     lookup.L.AutoAbilities,
		objLookupID:   lookup.L.AutoAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AutoAbility, AutoAbility, NamedAPIResource],
		retrieveQuery: cfg.db.GetAutoAbilityIDs,
		getSingleFunc: cfg.getAutoAbility,
		retrieveFunc:  cfg.retrieveAutoAbilities,
		avlFunc:       filterAvlAutoAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.EquipmentTables = handlerInput[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource]{
		Endpoint:      ts.EpEquipmentTables,
		resTypeSingle: ts.RTSEquipmentTables,
		resTypePlural: ts.RTPEquipmentTables,
		objLookup:     lookup.L.EquipmentTables,
		objLookupID:   lookup.L.EquipmentTablesID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.EquipmentTable, EquipmentTable, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetEquipmentTableIDs,
		getSingleFunc: cfg.getEquipmentTable,
		retrieveFunc:  cfg.retrieveEquipmentTables,
	}

	E.Equipment = handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]{
		Endpoint:      ts.EpEquipment,
		resTypeSingle: ts.RTSEquipment,
		resTypePlural: ts.RTPEquipment,
		objLookup:     lookup.L.EquipmentNames,
		objLookupID:   lookup.L.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.EquipmentName, EquipmentName, NamedAPIResource],
		retrieveQuery: cfg.db.GetEquipmentIDs,
		getSingleFunc: cfg.getEquipment,
		retrieveFunc:  cfg.retrieveEquipment,
	}

	E.CelestialWeapons = handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource]{
		Endpoint:      ts.EpCelestialWeapons,
		resTypeSingle: ts.RTSCelestialWeapons,
		resTypePlural: ts.RTPCelestialWeapons,
		objLookup:     lookup.L.CelestialWeapons,
		objLookupID:   lookup.L.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource],
		retrieveQuery: cfg.db.GetCelestialWeaponIDs,
		getSingleFunc: cfg.getCelestialWeapon,
		retrieveFunc:  cfg.retrieveCelestialWeapons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(cfg),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.Stats = handlerInput[seedtypes.Stat, Stat, NamedAPIResource]{
		Endpoint:      ts.EpStats,
		resTypeSingle: ts.RTSStats,
		resTypePlural: ts.RTPStats,
		objLookup:     lookup.L.Stats,
		objLookupID:   lookup.L.StatsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Stat, Stat, NamedAPIResource],
		retrieveQuery: cfg.db.GetStatIDs,
		getSingleFunc: cfg.getStat,
		retrieveFunc:  cfg.retrieveStats,
	}

	E.Properties = handlerInput[seedtypes.Property, Property, NamedAPIResource]{
		Endpoint:      ts.EpProperties,
		resTypeSingle: ts.RTSProperties,
		resTypePlural: ts.RTPProperties,
		objLookup:     lookup.L.Properties,
		objLookupID:   lookup.L.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Property, Property, NamedAPIResource],
		retrieveQuery: cfg.db.GetPropertyIDs,
		getSingleFunc: cfg.getProperty,
		retrieveFunc:  cfg.retrieveProperties,
	}

	E.OverdriveModes = handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource]{
		Endpoint:      ts.EpOverdriveModes,
		resTypeSingle: ts.RTSOverdriveModes,
		resTypePlural: ts.RTPOverdriveModes,
		objLookup:     lookup.L.OverdriveModes,
		objLookupID:   lookup.L.OverdriveModesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource],
		retrieveQuery: cfg.db.GetOverdriveModeIDs,
		getSingleFunc: cfg.getOverdriveMode,
		retrieveFunc:  cfg.retrieveOverdriveModes,
	}

	E.Elements = handlerInput[seedtypes.Element, Element, NamedAPIResource]{
		Endpoint:      ts.EpElements,
		resTypeSingle: ts.RTSElements,
		resTypePlural: ts.RTPElements,
		objLookup:     lookup.L.Elements,
		objLookupID:   lookup.L.ElementsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Element, Element, NamedAPIResource],
		retrieveQuery: cfg.db.GetElementIDs,
		getSingleFunc: cfg.getElement,
		retrieveFunc:  cfg.retrieveElements,
	}

	E.StatusConditions = handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]{
		Endpoint:      ts.EpStatusConditions,
		resTypeSingle: ts.RTSStatusConditions,
		resTypePlural: ts.RTPStatusConditions,
		objLookup:     lookup.L.StatusConditions,
		objLookupID:   lookup.L.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.StatusCondition, StatusCondition, NamedAPIResource],
		retrieveQuery: cfg.db.GetStatusConditionIDs,
		getSingleFunc: cfg.getStatusCondition,
		retrieveFunc:  cfg.retrieveStatusConditions,
	}

	E.Modifiers = handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource]{
		Endpoint:      ts.EpModifiers,
		resTypeSingle: ts.RTSModifiers,
		resTypePlural: ts.RTPModifiers,
		objLookup:     lookup.L.Modifiers,
		objLookupID:   lookup.L.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Modifier, Modifier, NamedAPIResource],
		retrieveQuery: cfg.db.GetModifierIDs,
		getSingleFunc: cfg.getModifier,
		retrieveFunc:  cfg.retrieveModifiers,
	}

	E.AgilityTiers = handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]{
		Endpoint:      ts.EpAgilityTiers,
		resTypeSingle: ts.RTSAgilityTiers,
		resTypePlural: ts.RTPAgilityTiers,
		objLookup:     nil,
		objLookupID:   lookup.L.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource],
		retrieveQuery: cfg.db.GetAgilityTierIDs,
		getSingleFunc: cfg.getAgilityTier,
		retrieveFunc:  cfg.retrieveAgilityTiers,
	}

	E.Enums = handlerInputEnums{
		Endpoint: ts.EpEnums,
		usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		lookup:   enums.T.Lookup,
	}

	E.Endpoints = handlerInputEndpoints{
		Endpoint: ts.EpEndpoints,
		usage:    []string{"/api/endpoints"},
	}

	E.AlBhed = handlerInputService[AlBhedParams, AlBhedResponse]{
		Endpoint:  ts.EpAlBhed,
		verifyFn:  verifyAlBhedParams,
		executeFn: handleAlBhed,
		docsFn:    getAlBhedParamsDoc,
		samplesFn: getAlBhedParamsSamples,
	}

	E.TurnOrder = handlerInputService[TurnOrderParams, TurnOrderResponse]{
		Endpoint:  ts.EpTurnOrder,
		verifyFn:  verifyTurnOrderParams,
		executeFn: handleTurnOrder,
		docsFn:    getTurnOrderParamsDoc,
		samplesFn: getTurnOrderParamsSamples,
	}

	E.DropChance = handlerInputService[DropChanceParams, DropChanceResponse]{
		Endpoint:  ts.EpDropChance,
		verifyFn:  verifyDropChanceParams,
		executeFn: handleDropChance,
		docsFn:    getDropChanceParamsDoc,
		samplesFn: getDropChanceParamsSamples,
	}

	E.CritChance = handlerInputService[CritChanceParams, CritChanceResponse]{
		Endpoint:  ts.EpCritChance,
		verifyFn:  verifyCritChanceParams,
		executeFn: handleCritChance,
		docsFn:    getCritChanceParamsDoc,
		samplesFn: getCritChanceParamsSamples,
	}

	E.Endpoints.slice = E.InitializeEndpointSlice()
}

func (E Endpoints) InitializeEndpointSlice() []ts.EndpointName {
	endpoints := []ts.EndpointName{
		E.Locations.Endpoint,
		E.Sublocations.Endpoint,
		E.Areas.Endpoint,
		E.MonsterFormations.Endpoint,
		E.Shops.Endpoint,
		E.Treasures.Endpoint,
		E.Quests.Endpoint,
		E.Sidequests.Endpoint,
		E.Subquests.Endpoint,
		E.ArenaCreations.Endpoint,
		E.BlitzballPrizes.Endpoint,
		E.Songs.Endpoint,
		E.FMVs.Endpoint,
		E.PlayerUnits.Endpoint,
		E.Characters.Endpoint,
		E.Aeons.Endpoint,
		E.CharacterClasses.Endpoint,
		E.Monsters.Endpoint,
		E.Abilities.Endpoint,
		E.PlayerAbilities.Endpoint,
		E.OverdriveAbilities.Endpoint,
		E.ItemAbilities.Endpoint,
		E.TriggerCommands.Endpoint,
		E.MiscAbilities.Endpoint,
		E.EnemyAbilities.Endpoint,
		E.AeonCommands.Endpoint,
		E.OverdriveCommands.Endpoint,
		E.Overdrives.Endpoint,
		E.RonsoRages.Endpoint,
		E.Submenus.Endpoint,
		E.Topmenus.Endpoint,
		E.AllItems.Endpoint,
		E.Items.Endpoint,
		E.Spheres.Endpoint,
		E.KeyItems.Endpoint,
		E.Primers.Endpoint,
		E.Mixes.Endpoint,
		E.AutoAbilities.Endpoint,
		E.EquipmentTables.Endpoint,
		E.Equipment.Endpoint,
		E.CelestialWeapons.Endpoint,
		E.Stats.Endpoint,
		E.Properties.Endpoint,
		E.OverdriveModes.Endpoint,
		E.Elements.Endpoint,
		E.StatusConditions.Endpoint,
		E.Modifiers.Endpoint,
		E.AgilityTiers.Endpoint,
		E.Enums.Endpoint,
		E.Endpoints.Endpoint,
		E.AlBhed.Endpoint,
		E.TurnOrder.Endpoint,
		E.DropChance.Endpoint,
		E.CritChance.Endpoint,
	}

	slices.Sort(endpoints)
	return endpoints
}
