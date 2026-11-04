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

func (E *Endpoints) Init(srv *Server) {
	E.Locations = handlerInput[seedtypes.Location, Location, NamedAPIResource]{
		Endpoint:      ts.EpLocations,
		resTypeSingle: ts.RTSLocations,
		resTypePlural: ts.RTPLocations,
		objLookup:     lookup.L.Locations,
		objLookupID:   lookup.L.LocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Location, Location, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetLocationIDs,
		getSingleFunc: srv.getLocation,
		retrieveFunc:  srv.retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     srv.Cfg.DB.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnSublocations: {
				dbQuery:     srv.Cfg.DB.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     srv.Cfg.DB.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     srv.Cfg.DB.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     srv.Cfg.DB.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     srv.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     srv.Cfg.DB.GetLocationTreasureIDs,
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
		retrieveQuery: srv.Cfg.DB.GetSublocationIDs,
		getSingleFunc: srv.getSublocation,
		retrieveFunc:  srv.retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     srv.Cfg.DB.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     srv.Cfg.DB.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			ts.SnMonsterFormations: {
				dbQuery:     srv.Cfg.DB.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     srv.Cfg.DB.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     srv.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     srv.Cfg.DB.GetSublocationTreasureIDs,
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
		retrieveQuery: srv.Cfg.DB.GetAreaIDs,
		getSingleFunc: srv.getArea,
		retrieveFunc:  srv.retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     srv.Cfg.DB.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     srv.Cfg.DB.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnSongs: {
				dbQuery:     srv.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     srv.Cfg.DB.GetAreaTreasureIDs,
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
		retrieveQuery: srv.Cfg.DB.GetMonsterFormationIDs,
		getSingleFunc: srv.getMonsterFormation,
		retrieveFunc:  srv.retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetMonsterFormationMonsterIDs,
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
		retrieveQuery: srv.Cfg.DB.GetShopIDs,
		getSingleFunc: srv.getShop,
		retrieveFunc:  srv.retrieveShops,
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
		retrieveQuery: srv.Cfg.DB.GetTreasureIDs,
		getSingleFunc: srv.getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  srv.retrieveTreasures,
	}

	E.Quests = handlerInput[seedtypes.Quest, Quest, QuestAPIResource]{
		Endpoint:      ts.EpQuests,
		resTypeSingle: ts.RTSQuests,
		resTypePlural: ts.RTPQuests,
		objLookup:     lookup.L.Quests,
		objLookupID:   lookup.L.QuestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Quest, Quest, QuestAPIResource],
		retrieveQuery: srv.Cfg.DB.GetQuestIDs,
		getSingleFunc: srv.getQuest,
		retrieveFunc:  srv.retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	E.Sidequests = handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]{
		Endpoint:      ts.EpSidequests,
		resTypeSingle: ts.RTSSidequests,
		resTypePlural: ts.RTPSidequests,
		objLookup:     lookup.L.Sidequests,
		objLookupID:   lookup.L.SidequestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Sidequest, Sidequest, QuestAPIResource],
		retrieveQuery: srv.Cfg.DB.GetSidequestIDs,
		getSingleFunc: srv.getSidequest,
		retrieveFunc:  srv.retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSubquests: {
				dbQuery:     srv.Cfg.DB.GetSidequestSubquestIDs,
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
		retrieveQuery: srv.Cfg.DB.GetSubquestIDs,
		getSingleFunc: srv.getSubquest,
		retrieveFunc:  srv.retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	E.ArenaCreations = handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]{
		Endpoint:      ts.EpArenaCreations,
		resTypeSingle: ts.RTSArenaCreations,
		resTypePlural: ts.RTPArenaCreations,
		objLookup:     lookup.L.ArenaCreations,
		objLookupID:   lookup.L.ArenaCreationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetArenaCreationIDs,
		getSingleFunc: srv.getArenaCreation,
		retrieveFunc:  srv.retrieveArenaCreations,
	}

	E.BlitzballPrizes = handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource]{
		Endpoint:      ts.EpBlitzballPrizes,
		resTypeSingle: ts.RTSBlitzballPrizes,
		resTypePlural: ts.RTPBlitzballPrizes,
		objLookup:     lookup.L.Positions,
		objLookupID:   lookup.L.PositionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetBlitzballPrizeIDs,
		getSingleFunc: srv.getBlitzballPrize,
		retrieveFunc:  srv.retrieveBlitzballPrizes,
	}

	E.Songs = handlerInput[seedtypes.Song, Song, NamedAPIResource]{
		Endpoint:      ts.EpSongs,
		resTypeSingle: ts.RTSSongs,
		resTypePlural: ts.RTPSongs,
		objLookup:     lookup.L.Songs,
		objLookupID:   lookup.L.SongsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Song, Song, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetSongIDs,
		getSingleFunc: srv.getSong,
		retrieveFunc:  srv.retrieveSongs,
	}

	E.FMVs = handlerInput[seedtypes.FMV, FMV, NamedAPIResource]{
		Endpoint:      ts.EpFMVs,
		resTypeSingle: ts.RTSFMVs,
		resTypePlural: ts.RTPFMVs,
		objLookup:     lookup.L.FMVs,
		objLookupID:   lookup.L.FMVsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.FMV, FMV, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetFmvIDs,
		getSingleFunc: srv.getFMV,
		retrieveFunc:  srv.retrieveFMVs,
	}

	E.PlayerUnits = handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource]{
		Endpoint:      ts.EpPlayerUnits,
		resTypeSingle: ts.RTSPlayerUnits,
		resTypePlural: ts.RTPPlayerUnits,
		objLookup:     lookup.L.PlayerUnits,
		objLookupID:   lookup.L.PlayerUnitsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetPlayerUnitIDs,
		getSingleFunc: srv.getPlayerUnit,
		retrieveFunc:  srv.retrievePlayerUnits,
	}

	E.Characters = handlerInput[seedtypes.Character, Character, NamedAPIResource]{
		Endpoint:      ts.EpCharacters,
		resTypeSingle: ts.RTSCharacters,
		resTypePlural: ts.RTPCharacters,
		objLookup:     lookup.L.Characters,
		objLookupID:   lookup.L.CharactersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Character, Character, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetCharacterIDs,
		getSingleFunc: srv.getCharacter,
		retrieveFunc:  srv.retrieveCharacters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnStdSgAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnExpSgAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     srv.Cfg.DB.GetCharacterOverdriveIDs,
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
		retrieveQuery: srv.Cfg.DB.GetAeonIDs,
		getSingleFunc: srv.getAeon,
		retrieveFunc:  srv.retrieveAeons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     srv.Cfg.DB.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     srv.Cfg.DB.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     srv.Cfg.DB.GetAeonOverdriveIDs,
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
		retrieveQuery: srv.Cfg.DB.GetCharacterClassesIDs,
		getSingleFunc: srv.getCharacterClass,
		retrieveFunc:  srv.retrieveCharacterClasses,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnLearnableAbilities: {
				dbQuery:     srv.Cfg.DB.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnDefaultOverdrives: {
				dbQuery:     srv.Cfg.DB.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			ts.SnLearnableOverdrives: {
				dbQuery:     srv.Cfg.DB.GetCharacterClassLearnableOverdriveIDs,
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
		getMultipleQuery: srv.Cfg.DB.GetMonsterIDsByName,
		retrieveQuery:    srv.Cfg.DB.GetMonsterIDs,
		getSingleFunc:    srv.getMonster,
		retrieveFunc:     srv.retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnAbilities: {
				dbQuery:     srv.Cfg.DB.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnAreas: {
				dbQuery:     srv.Cfg.DB.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     srv.Cfg.DB.GetMonsterMonsterFormationIDs,
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
		retrieveQuery: srv.Cfg.DB.GetAbilityIDs,
		getSingleFunc: srv.getAbility,
		retrieveFunc:  srv.retrieveAbilities,
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
		getMultipleQuery: getTypedAbilityIDsByName(srv, database.AbilityTypePlayerAbility),
		retrieveQuery:    srv.Cfg.DB.GetPlayerAbilityIDs,
		getSingleFunc:    srv.getPlayerAbility,
		retrieveFunc:     srv.retrievePlayerAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetPlayerAbilityMonsterIDs,
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
		getMultipleQuery: getTypedAbilityIDsByName(srv, database.AbilityTypeOverdriveAbility),
		retrieveQuery:    srv.Cfg.DB.GetOverdriveAbilityIDs,
		getSingleFunc:    srv.getOverdriveAbility,
		retrieveFunc:     srv.retrieveOverdriveAbilities,
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
		retrieveQuery: srv.Cfg.DB.GetItemAbilityIDs,
		getSingleFunc: srv.getItemAbility,
		retrieveFunc:  srv.retrieveItemAbilities,
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
		getMultipleQuery: getTypedAbilityIDsByName(srv, database.AbilityTypeTriggerCommand),
		retrieveQuery:    srv.Cfg.DB.GetTriggerCommandIDs,
		getSingleFunc:    srv.getTriggerCommand,
		retrieveFunc:     srv.retrieveTriggerCommands,
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
		getMultipleQuery: getTypedAbilityIDsByName(srv, database.AbilityTypeMiscAbility),
		retrieveQuery:    srv.Cfg.DB.GetMiscAbilityIDs,
		getSingleFunc:    srv.getMiscAbility,
		retrieveFunc:     srv.retrieveMiscAbilities,
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
		getMultipleQuery: getTypedAbilityIDsByName(srv, database.AbilityTypeEnemyAbility),
		retrieveQuery:    srv.Cfg.DB.GetEnemyAbilityIDs,
		getSingleFunc:    srv.getEnemyAbility,
		retrieveFunc:     srv.retrieveEnemyAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetEnemyAbilityMonsterIDs,
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
		retrieveQuery: srv.Cfg.DB.GetAeonCommandIDs,
		getSingleFunc: srv.getAeonCommand,
		retrieveFunc:  srv.retrieveAeonCommands,
	}

	E.OverdriveCommands = handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]{
		Endpoint:      ts.EpOverdriveCommands,
		resTypeSingle: ts.RTSOverdriveCommands,
		resTypePlural: ts.RTPOverdriveCommands,
		objLookup:     lookup.L.OverdriveCommands,
		objLookupID:   lookup.L.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetOverdriveCommandIDs,
		getSingleFunc: srv.getOverdriveCommand,
		retrieveFunc:  srv.retrieveOverdriveCommands,
		subsections: map[ts.SectionName]Subsection{
			ts.SnOverdriveAbilities: {
				dbQuery:     ToIntManyNull(srv.Cfg.DB.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     ToIntManyNull(srv.Cfg.DB.GetOverdriveCommandOverdriveIDs),
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
		retrieveQuery: srv.Cfg.DB.GetOverdriveIDs,
		getSingleFunc: srv.getOverdrive,
		retrieveFunc:  srv.retrieveOverdrives,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createOverdriveSimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     srv.Cfg.DB.GetOverdriveOverdriveAbilityIDs,
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
		retrieveQuery: srv.Cfg.DB.GetRonsoRageIDs,
		getSingleFunc: srv.getRonsoRage,
		retrieveFunc:  srv.retrieveRonsoRages,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMonsters: {
				dbQuery:     srv.Cfg.DB.GetRonsoRageMonsterIDs,
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
		retrieveQuery: srv.Cfg.DB.GetSubmenuIDs,
		getSingleFunc: srv.getSubmenu,
		retrieveFunc:  srv.retrieveSubmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(srv.Cfg.DB.GetSubmenuAbilityIDs),
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
		retrieveQuery: srv.Cfg.DB.GetTopmenuIDs,
		getSingleFunc: srv.getTopmenu,
		retrieveFunc:  srv.retrieveTopmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(srv.Cfg.DB.GetTopmenuAbilityIDs),
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
		retrieveQuery: srv.Cfg.DB.GetMasterItemIDs,
		getSingleFunc: srv.getMasterItem,
		retrieveFunc:  srv.retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	E.Items = handlerInput[seedtypes.Item, Item, NamedAPIResource]{
		Endpoint:      ts.EpItems,
		resTypeSingle: ts.RTSItems,
		resTypePlural: ts.RTPItems,
		objLookup:     lookup.L.Items,
		objLookupID:   lookup.L.ItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Item, Item, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetItemIDs,
		getSingleFunc: srv.getItem,
		retrieveFunc:  srv.retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMixes: {
				dbQuery:     srv.Cfg.DB.GetItemMixIDs,
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
		retrieveQuery: srv.Cfg.DB.GetSphereIDs,
		getSingleFunc: srv.getSphere,
		retrieveFunc:  srv.retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	E.KeyItems = handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]{
		Endpoint:      ts.EpKeyItems,
		resTypeSingle: ts.RTSKeyItems,
		resTypePlural: ts.RTPKeyItems,
		objLookup:     lookup.L.KeyItems,
		objLookupID:   lookup.L.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.KeyItem, KeyItem, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetKeyItemIDs,
		getSingleFunc: srv.getKeyItem,
		retrieveFunc:  srv.retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	E.Primers = handlerInput[seedtypes.Primer, Primer, NamedAPIResource]{
		Endpoint:      ts.EpPrimers,
		resTypeSingle: ts.RTSPrimers,
		resTypePlural: ts.RTPPrimers,
		objLookup:     lookup.L.Primers,
		objLookupID:   lookup.L.PrimersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Primer, Primer, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetPrimerIDs,
		getSingleFunc: srv.getPrimer,
		retrieveFunc:  srv.retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	E.Mixes = handlerInput[seedtypes.Mix, Mix, NamedAPIResource]{
		Endpoint:      ts.EpMixes,
		resTypeSingle: ts.RTSMixes,
		resTypePlural: ts.RTPMixes,
		objLookup:     lookup.L.Mixes,
		objLookupID:   lookup.L.MixesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Mix, Mix, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetMixIDs,
		getSingleFunc: srv.getMix,
		retrieveFunc:  srv.retrieveMixes,
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
		retrieveQuery: srv.Cfg.DB.GetAutoAbilityIDs,
		getSingleFunc: srv.getAutoAbility,
		retrieveFunc:  srv.retrieveAutoAbilities,
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
		retrieveQuery: srv.Cfg.DB.GetEquipmentTableIDs,
		getSingleFunc: srv.getEquipmentTable,
		retrieveFunc:  srv.retrieveEquipmentTables,
	}

	E.Equipment = handlerInput[seedtypes.EquipmentName, EquipmentName, NamedAPIResource]{
		Endpoint:      ts.EpEquipment,
		resTypeSingle: ts.RTSEquipment,
		resTypePlural: ts.RTPEquipment,
		objLookup:     lookup.L.EquipmentNames,
		objLookupID:   lookup.L.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.EquipmentName, EquipmentName, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetEquipmentIDs,
		getSingleFunc: srv.getEquipment,
		retrieveFunc:  srv.retrieveEquipment,
	}

	E.CelestialWeapons = handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource]{
		Endpoint:      ts.EpCelestialWeapons,
		resTypeSingle: ts.RTSCelestialWeapons,
		resTypePlural: ts.RTPCelestialWeapons,
		objLookup:     lookup.L.CelestialWeapons,
		objLookupID:   lookup.L.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetCelestialWeaponIDs,
		getSingleFunc: srv.getCelestialWeapon,
		retrieveFunc:  srv.retrieveCelestialWeapons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(srv),
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
		retrieveQuery: srv.Cfg.DB.GetStatIDs,
		getSingleFunc: srv.getStat,
		retrieveFunc:  srv.retrieveStats,
	}

	E.Properties = handlerInput[seedtypes.Property, Property, NamedAPIResource]{
		Endpoint:      ts.EpProperties,
		resTypeSingle: ts.RTSProperties,
		resTypePlural: ts.RTPProperties,
		objLookup:     lookup.L.Properties,
		objLookupID:   lookup.L.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Property, Property, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetPropertyIDs,
		getSingleFunc: srv.getProperty,
		retrieveFunc:  srv.retrieveProperties,
	}

	E.OverdriveModes = handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource]{
		Endpoint:      ts.EpOverdriveModes,
		resTypeSingle: ts.RTSOverdriveModes,
		resTypePlural: ts.RTPOverdriveModes,
		objLookup:     lookup.L.OverdriveModes,
		objLookupID:   lookup.L.OverdriveModesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetOverdriveModeIDs,
		getSingleFunc: srv.getOverdriveMode,
		retrieveFunc:  srv.retrieveOverdriveModes,
	}

	E.Elements = handlerInput[seedtypes.Element, Element, NamedAPIResource]{
		Endpoint:      ts.EpElements,
		resTypeSingle: ts.RTSElements,
		resTypePlural: ts.RTPElements,
		objLookup:     lookup.L.Elements,
		objLookupID:   lookup.L.ElementsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Element, Element, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetElementIDs,
		getSingleFunc: srv.getElement,
		retrieveFunc:  srv.retrieveElements,
	}

	E.StatusConditions = handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]{
		Endpoint:      ts.EpStatusConditions,
		resTypeSingle: ts.RTSStatusConditions,
		resTypePlural: ts.RTPStatusConditions,
		objLookup:     lookup.L.StatusConditions,
		objLookupID:   lookup.L.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.StatusCondition, StatusCondition, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetStatusConditionIDs,
		getSingleFunc: srv.getStatusCondition,
		retrieveFunc:  srv.retrieveStatusConditions,
	}

	E.Modifiers = handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource]{
		Endpoint:      ts.EpModifiers,
		resTypeSingle: ts.RTSModifiers,
		resTypePlural: ts.RTPModifiers,
		objLookup:     lookup.L.Modifiers,
		objLookupID:   lookup.L.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Modifier, Modifier, NamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetModifierIDs,
		getSingleFunc: srv.getModifier,
		retrieveFunc:  srv.retrieveModifiers,
	}

	E.AgilityTiers = handlerInput[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource]{
		Endpoint:      ts.EpAgilityTiers,
		resTypeSingle: ts.RTSAgilityTiers,
		resTypePlural: ts.RTPAgilityTiers,
		objLookup:     nil,
		objLookupID:   lookup.L.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.AgilityTier, AgilityTier, UnnamedAPIResource],
		retrieveQuery: srv.Cfg.DB.GetAgilityTierIDs,
		getSingleFunc: srv.getAgilityTier,
		retrieveFunc:  srv.retrieveAgilityTiers,
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
