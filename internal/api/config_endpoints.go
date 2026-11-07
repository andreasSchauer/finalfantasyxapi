package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

var E = &Endpoints{}

type Endpoints struct {
	Locations          HandlerInput[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource]
	Sublocations       HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource]
	Areas              HandlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource]
	MonsterFormations  HandlerInput[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource]
	Shops              HandlerInput[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource]
	Treasures          HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource]
	Quests             HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource]
	Sidequests         HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource]
	Subquests          HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource]
	ArenaCreations     HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource]
	BlitzballPrizes    HandlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource]
	Songs              HandlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource]
	FMVs               HandlerInput[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource]
	PlayerUnits        HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource]
	Characters         HandlerInput[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource]
	Aeons              HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource]
	CharacterClasses   HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource]
	Monsters           HandlerInput[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource]
	Abilities          HandlerInput[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource]
	PlayerAbilities    HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource]
	OverdriveAbilities HandlerInput[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource]
	ItemAbilities      HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource]
	TriggerCommands    HandlerInput[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource]
	MiscAbilities      HandlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource]
	EnemyAbilities     HandlerInput[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource]
	AeonCommands       HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource]
	OverdriveCommands  HandlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource]
	Overdrives         HandlerInput[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource]
	RonsoRages         HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]
	Submenus           HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource]
	Topmenus           HandlerInput[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource]
	AllItems           HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource]
	Items              HandlerInput[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource]
	KeyItems           HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource]
	Spheres            HandlerInput[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource]
	Primers            HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource]
	Mixes              HandlerInput[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource]
	AutoAbilities      HandlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource]
	EquipmentTables    HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource]
	Equipment          HandlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource]
	CelestialWeapons   HandlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource]
	Stats              HandlerInput[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource]
	Properties         HandlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource]
	OverdriveModes     HandlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource]
	Elements           HandlerInput[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource]
	StatusConditions   HandlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource]
	Modifiers          HandlerInput[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource]
	AgilityTiers       HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource]
	Enums              HandlerInputEnums
	Endpoints          HandlerInputEndpoints
	AlBhed             HandlerInputService[AlBhedParams, AlBhedResponse]
	TurnOrder          HandlerInputService[TurnOrderParams, TurnOrderResponse]
	DropChance         HandlerInputService[DropChanceParams, DropChanceResponse]
	CritChance         HandlerInputService[CritChanceParams, CritChanceResponse]
}

func (E *Endpoints) Init() {
	E.Locations = HandlerInput[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpLocations,
		resTypeSingle: ts.RTSLocations,
		resTypePlural: ts.RTPLocations,
		objLookup:     lookup.L.Locations,
		objLookupID:   lookup.L.LocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetLocationIDs,
		getSingleFunc: getLocation,
		retrieveFunc:  retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     config.Cfg.DB.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			ts.SnSublocations: {
				dbQuery:     config.Cfg.DB.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     config.Cfg.DB.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     config.Cfg.DB.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     config.Cfg.DB.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     config.Cfg.DB.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.Sublocations = HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSublocations,
		resTypeSingle: ts.RTSSublocations,
		resTypePlural: ts.RTPSublocations,
		objLookup:     lookup.L.Sublocations,
		objLookupID:   lookup.L.SublocationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetSublocationIDs,
		getSingleFunc: getSublocation,
		retrieveFunc:  retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     config.Cfg.DB.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				dbQuery:     config.Cfg.DB.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			ts.SnMonsterFormations: {
				dbQuery:     config.Cfg.DB.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				dbQuery:     config.Cfg.DB.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			ts.SnSongs: {
				dbQuery:     getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     config.Cfg.DB.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.Areas = HandlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource]{
		Endpoint:      ts.EpAreas,
		resTypeSingle: ts.RTSAreas,
		resTypePlural: ts.RTPAreas,
		objLookup:     lookup.L.Areas,
		objLookupID:   lookup.L.AreasID,
		idToResFunc:   idToAreaAPIResource,
		retrieveQuery: config.Cfg.DB.GetAreaIDs,
		getSingleFunc: getArea,
		retrieveFunc:  retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnConnected: {
				dbQuery:     config.Cfg.DB.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     config.Cfg.DB.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnSongs: {
				dbQuery:     getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				dbQuery:     config.Cfg.DB.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	E.MonsterFormations = HandlerInput[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpMonsterFormations,
		resTypeSingle: ts.RTSMonsterFormations,
		resTypePlural: ts.RTPMonsterFormations,
		objLookup:     lookup.L.MonsterFormations,
		objLookupID:   lookup.L.MonsterFormationsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetMonsterFormationIDs,
		getSingleFunc: getMonsterFormation,
		retrieveFunc:  retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Shops = HandlerInput[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpShops,
		resTypeSingle: ts.RTSShops,
		resTypePlural: ts.RTPShops,
		objLookup:     lookup.L.Shops,
		objLookupID:   lookup.L.ShopsID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetShopIDs,
		getSingleFunc: getShop,
		retrieveFunc:  retrieveShops,
		avlFunc:       filterAvlShops,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createShopSimple,
			},
		},
	}

	E.Treasures = HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpTreasures,
		resTypeSingle: ts.RTSTreasures,
		resTypePlural: ts.RTPTreasures,
		objLookup:     lookup.L.Treasures,
		objLookupID:   lookup.L.TreasuresID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetTreasureIDs,
		getSingleFunc: getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  retrieveTreasures,
	}

	E.Quests = HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpQuests,
		resTypeSingle: ts.RTSQuests,
		resTypePlural: ts.RTPQuests,
		objLookup:     lookup.L.Quests,
		objLookupID:   lookup.L.QuestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource],
		retrieveQuery: config.Cfg.DB.GetQuestIDs,
		getSingleFunc: getQuest,
		retrieveFunc:  retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	E.Sidequests = HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpSidequests,
		resTypeSingle: ts.RTSSidequests,
		resTypePlural: ts.RTPSidequests,
		objLookup:     lookup.L.Sidequests,
		objLookupID:   lookup.L.SidequestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource],
		retrieveQuery: config.Cfg.DB.GetSidequestIDs,
		getSingleFunc: getSidequest,
		retrieveFunc:  retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSubquests: {
				dbQuery:     config.Cfg.DB.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	E.Subquests = HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpSubquests,
		resTypeSingle: ts.RTSSubquests,
		resTypePlural: ts.RTPSubquests,
		objLookup:     lookup.L.Subquests,
		objLookupID:   lookup.L.SubquestsID,
		idToResFunc:   idToQuestAPIResource[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource],
		retrieveQuery: config.Cfg.DB.GetSubquestIDs,
		getSingleFunc: getSubquest,
		retrieveFunc:  retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	E.ArenaCreations = HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpArenaCreations,
		resTypeSingle: ts.RTSArenaCreations,
		resTypePlural: ts.RTPArenaCreations,
		objLookup:     lookup.L.ArenaCreations,
		objLookupID:   lookup.L.ArenaCreationsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetArenaCreationIDs,
		getSingleFunc: getArenaCreation,
		retrieveFunc:  retrieveArenaCreations,
	}

	E.BlitzballPrizes = HandlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpBlitzballPrizes,
		resTypeSingle: ts.RTSBlitzballPrizes,
		resTypePlural: ts.RTPBlitzballPrizes,
		objLookup:     lookup.L.Positions,
		objLookupID:   lookup.L.PositionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetBlitzballPrizeIDs,
		getSingleFunc: getBlitzballPrize,
		retrieveFunc:  retrieveBlitzballPrizes,
	}

	E.Songs = HandlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSongs,
		resTypeSingle: ts.RTSSongs,
		resTypePlural: ts.RTPSongs,
		objLookup:     lookup.L.Songs,
		objLookupID:   lookup.L.SongsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetSongIDs,
		getSingleFunc: getSong,
		retrieveFunc:  retrieveSongs,
	}

	E.FMVs = HandlerInput[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpFMVs,
		resTypeSingle: ts.RTSFMVs,
		resTypePlural: ts.RTPFMVs,
		objLookup:     lookup.L.FMVs,
		objLookupID:   lookup.L.FMVsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetFmvIDs,
		getSingleFunc: getFMV,
		retrieveFunc:  retrieveFMVs,
	}

	E.PlayerUnits = HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpPlayerUnits,
		resTypeSingle: ts.RTSPlayerUnits,
		resTypePlural: ts.RTPPlayerUnits,
		objLookup:     lookup.L.PlayerUnits,
		objLookupID:   lookup.L.PlayerUnitsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource],
		retrieveQuery: config.Cfg.DB.GetPlayerUnitIDs,
		getSingleFunc: getPlayerUnit,
		retrieveFunc:  retrievePlayerUnits,
	}

	E.Characters = HandlerInput[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCharacters,
		resTypeSingle: ts.RTSCharacters,
		resTypePlural: ts.RTPCharacters,
		objLookup:     lookup.L.Characters,
		objLookupID:   lookup.L.CharactersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetCharacterIDs,
		getSingleFunc: getCharacter,
		retrieveFunc:  retrieveCharacters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnStdSgAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnExpSgAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     config.Cfg.DB.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Aeons = HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAeons,
		resTypeSingle: ts.RTSAeons,
		resTypePlural: ts.RTPAeons,
		objLookup:     lookup.L.Aeons,
		objLookupID:   lookup.L.AeonsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetAeonIDs,
		getSingleFunc: getAeon,
		retrieveFunc:  retrieveAeons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     config.Cfg.DB.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     config.Cfg.DB.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     config.Cfg.DB.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			ts.SnStats: {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	E.CharacterClasses = HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCharacterClasses,
		resTypeSingle: ts.RTSCharacterClasses,
		resTypePlural: ts.RTPCharacterClasses,
		objLookup:     lookup.L.CharClasses,
		objLookupID:   lookup.L.CharClassesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetCharacterClassesIDs,
		getSingleFunc: getCharacterClass,
		retrieveFunc:  retrieveCharacterClasses,
		subsections: map[ts.SectionName]Subsection{
			ts.SnDefaultAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnLearnableAbilities: {
				dbQuery:     config.Cfg.DB.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnDefaultOverdrives: {
				dbQuery:     config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			ts.SnLearnableOverdrives: {
				dbQuery:     config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Monsters = HandlerInput[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpMonsters,
		resTypeSingle:    ts.RTSMonsters,
		resTypePlural:    ts.RTPMonsters,
		objLookup:        lookup.L.Monsters,
		objLookupID:      lookup.L.MonstersID,
		idToResFunc:      idToNamedAPIResource[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource],
		getMultipleQuery: config.Cfg.DB.GetMonsterIDsByName,
		retrieveQuery:    config.Cfg.DB.GetMonsterIDs,
		getSingleFunc:    getMonster,
		retrieveFunc:     retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			ts.SnAbilities: {
				dbQuery:     config.Cfg.DB.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			ts.SnAreas: {
				dbQuery:     config.Cfg.DB.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				dbQuery:     config.Cfg.DB.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	E.Abilities = HandlerInput[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpAbilities,
		resTypeSingle: ts.RTSAbilities,
		resTypePlural: ts.RTPAbilities,
		objLookup:     lookup.L.Abilities,
		objLookupID:   lookup.L.AbilitiesID,
		idToResFunc:   idToTypedAPIResource[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource],
		retrieveQuery: config.Cfg.DB.GetAbilityIDs,
		getSingleFunc: getAbility,
		retrieveFunc:  retrieveAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.PlayerAbilities = HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpPlayerAbilities,
		resTypeSingle:    ts.RTSPlayerAbilities,
		resTypePlural:    ts.RTPPlayerAbilities,
		objLookup:        lookup.L.PlayerAbilities,
		objLookupID:      lookup.L.PlayerAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypePlayerAbility),
		retrieveQuery:    config.Cfg.DB.GetPlayerAbilityIDs,
		getSingleFunc:    getPlayerAbility,
		retrieveFunc:     retrievePlayerAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.OverdriveAbilities = HandlerInput[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpOverdriveAbilities,
		resTypeSingle:    ts.RTSOverdriveAbilities,
		resTypePlural:    ts.RTPOverdriveAbilities,
		objLookup:        lookup.L.OverdriveAbilities,
		objLookupID:      lookup.L.OverdriveAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeOverdriveAbility),
		retrieveQuery:    config.Cfg.DB.GetOverdriveAbilityIDs,
		getSingleFunc:    getOverdriveAbility,
		retrieveFunc:     retrieveOverdriveAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.ItemAbilities = HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpItemAbilities,
		resTypeSingle: ts.RTSItemAbilities,
		resTypePlural: ts.RTPItemAbilities,
		objLookup:     lookup.L.ItemAbilities,
		objLookupID:   lookup.L.ItemAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetItemAbilityIDs,
		getSingleFunc: getItemAbility,
		retrieveFunc:  retrieveItemAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	E.TriggerCommands = HandlerInput[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpTriggerCommands,
		resTypeSingle:    ts.RTSTriggerCommands,
		resTypePlural:    ts.RTPTriggerCommands,
		objLookup:        lookup.L.TriggerCommands,
		objLookupID:      lookup.L.TriggerCommandsID,
		idToResFunc:      idToNamedAPIResource[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeTriggerCommand),
		retrieveQuery:    config.Cfg.DB.GetTriggerCommandIDs,
		getSingleFunc:    getTriggerCommand,
		retrieveFunc:     retrieveTriggerCommands,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	E.MiscAbilities = HandlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpMiscAbilities,
		resTypeSingle:    ts.RTSMiscAbilities,
		resTypePlural:    ts.RTPMiscAbilities,
		objLookup:        lookup.L.MiscAbilities,
		objLookupID:      lookup.L.MiscAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeMiscAbility),
		retrieveQuery:    config.Cfg.DB.GetMiscAbilityIDs,
		getSingleFunc:    getMiscAbility,
		retrieveFunc:     retrieveMiscAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMiscAbilitySimple,
			},
		},
	}

	E.EnemyAbilities = HandlerInput[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpEnemyAbilities,
		resTypeSingle:    ts.RTSEnemyAbilities,
		resTypePlural:    ts.RTPEnemyAbilities,
		objLookup:        lookup.L.EnemyAbilities,
		objLookupID:      lookup.L.EnemyAbilitiesID,
		idToResFunc:      idToNamedAPIResource[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource],
		getMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeEnemyAbility),
		retrieveQuery:    config.Cfg.DB.GetEnemyAbilityIDs,
		getSingleFunc:    getEnemyAbility,
		retrieveFunc:     retrieveEnemyAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.AeonCommands = HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAeonCommands,
		resTypeSingle: ts.RTSAeonCommands,
		resTypePlural: ts.RTPAeonCommands,
		objLookup:     lookup.L.AeonCommands,
		objLookupID:   lookup.L.AeonCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetAeonCommandIDs,
		getSingleFunc: getAeonCommand,
		retrieveFunc:  retrieveAeonCommands,
	}

	E.OverdriveCommands = HandlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdriveCommands,
		resTypeSingle: ts.RTSOverdriveCommands,
		resTypePlural: ts.RTPOverdriveCommands,
		objLookup:     lookup.L.OverdriveCommands,
		objLookupID:   lookup.L.OverdriveCommandsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetOverdriveCommandIDs,
		getSingleFunc: getOverdriveCommand,
		retrieveFunc:  retrieveOverdriveCommands,
		subsections: map[ts.SectionName]Subsection{
			ts.SnOverdriveAbilities: {
				dbQuery:     ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				dbQuery:     ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs),
				createSubFn: createOverdriveSimple,
			},
		},
	}

	E.Overdrives = HandlerInput[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdrives,
		resTypeSingle: ts.RTSOverdrives,
		resTypePlural: ts.RTPOverdrives,
		objLookup:     lookup.L.Overdrives,
		objLookupID:   lookup.L.OverdrivesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetOverdriveIDs,
		getSingleFunc: getOverdrive,
		retrieveFunc:  retrieveOverdrives,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createOverdriveSimple,
			},
			ts.SnOverdriveAbilities: {
				dbQuery:     config.Cfg.DB.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.RonsoRages = HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpRonsoRages,
		resTypeSingle: ts.RTSRonsoRages,
		resTypePlural: ts.RTPRonsoRages,
		objLookup:     lookup.L.RonsoRages,
		objLookupID:   lookup.L.RonsoRagesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetRonsoRageIDs,
		getSingleFunc: getRonsoRage,
		retrieveFunc:  retrieveRonsoRages,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMonsters: {
				dbQuery:     config.Cfg.DB.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Submenus = HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSubmenus,
		resTypeSingle: ts.RTSSubmenus,
		resTypePlural: ts.RTPSubmenus,
		objLookup:     lookup.L.Submenus,
		objLookupID:   lookup.L.SubmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetSubmenuIDs,
		getSingleFunc: getSubmenu,
		retrieveFunc:  retrieveSubmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(config.Cfg.DB.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.Topmenus = HandlerInput[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpTopmenus,
		resTypeSingle: ts.RTSTopmenus,
		resTypePlural: ts.RTPTopmenus,
		objLookup:     lookup.L.Topmenus,
		objLookupID:   lookup.L.TopmenusID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetTopmenuIDs,
		getSingleFunc: getTopmenu,
		retrieveFunc:  retrieveTopmenus,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAbilities: {
				dbQuery:     ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	E.AllItems = HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpAllItems,
		resTypeSingle: ts.RTSAllItems,
		resTypePlural: ts.RTPAllItems,
		objLookup:     lookup.L.MasterItems,
		objLookupID:   lookup.L.MasterItemsID,
		idToResFunc:   idToTypedAPIResource[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource],
		retrieveQuery: config.Cfg.DB.GetMasterItemIDs,
		getSingleFunc: getMasterItem,
		retrieveFunc:  retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	E.Items = HandlerInput[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpItems,
		resTypeSingle: ts.RTSItems,
		resTypePlural: ts.RTPItems,
		objLookup:     lookup.L.Items,
		objLookupID:   lookup.L.ItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetItemIDs,
		getSingleFunc: getItem,
		retrieveFunc:  retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[ts.SectionName]Subsection{
			ts.SnMixes: {
				dbQuery:     config.Cfg.DB.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	E.Spheres = HandlerInput[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSpheres,
		resTypeSingle: ts.RTSSpheres,
		resTypePlural: ts.RTPSpheres,
		objLookup:     lookup.L.Spheres,
		objLookupID:   lookup.L.SpheresID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetSphereIDs,
		getSingleFunc: getSphere,
		retrieveFunc:  retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	E.KeyItems = HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpKeyItems,
		resTypeSingle: ts.RTSKeyItems,
		resTypePlural: ts.RTPKeyItems,
		objLookup:     lookup.L.KeyItems,
		objLookupID:   lookup.L.KeyItemsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetKeyItemIDs,
		getSingleFunc: getKeyItem,
		retrieveFunc:  retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	E.Primers = HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpPrimers,
		resTypeSingle: ts.RTSPrimers,
		resTypePlural: ts.RTPPrimers,
		objLookup:     lookup.L.Primers,
		objLookupID:   lookup.L.PrimersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetPrimerIDs,
		getSingleFunc: getPrimer,
		retrieveFunc:  retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	E.Mixes = HandlerInput[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpMixes,
		resTypeSingle: ts.RTSMixes,
		resTypePlural: ts.RTPMixes,
		objLookup:     lookup.L.Mixes,
		objLookupID:   lookup.L.MixesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetMixIDs,
		getSingleFunc: getMix,
		retrieveFunc:  retrieveMixes,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createMixSimple,
			},
		},
	}

	E.AutoAbilities = HandlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAutoAbilities,
		resTypeSingle: ts.RTSAutoAbilities,
		resTypePlural: ts.RTPAutoAbilities,
		objLookup:     lookup.L.AutoAbilities,
		objLookupID:   lookup.L.AutoAbilitiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetAutoAbilityIDs,
		getSingleFunc: getAutoAbility,
		retrieveFunc:  retrieveAutoAbilities,
		avlFunc:       filterAvlAutoAbilities,
		subsections: map[ts.SectionName]Subsection{
			ts.SnSimple: {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.EquipmentTables = HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpEquipmentTables,
		resTypeSingle: ts.RTSEquipmentTables,
		resTypePlural: ts.RTPEquipmentTables,
		objLookup:     lookup.L.EquipmentTables,
		objLookupID:   lookup.L.EquipmentTablesID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetEquipmentTableIDs,
		getSingleFunc: getEquipmentTable,
		retrieveFunc:  retrieveEquipmentTables,
	}

	E.Equipment = HandlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpEquipment,
		resTypeSingle: ts.RTSEquipment,
		resTypePlural: ts.RTPEquipment,
		objLookup:     lookup.L.EquipmentNames,
		objLookupID:   lookup.L.EquipmentNamesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetEquipmentIDs,
		getSingleFunc: getEquipment,
		retrieveFunc:  retrieveEquipment,
	}

	E.CelestialWeapons = HandlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCelestialWeapons,
		resTypeSingle: ts.RTSCelestialWeapons,
		resTypePlural: ts.RTPCelestialWeapons,
		objLookup:     lookup.L.CelestialWeapons,
		objLookupID:   lookup.L.CelestialWeaponsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetCelestialWeaponIDs,
		getSingleFunc: getCelestialWeapon,
		retrieveFunc:  retrieveCelestialWeapons,
		subsections: map[ts.SectionName]Subsection{
			ts.SnAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.Stats = HandlerInput[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpStats,
		resTypeSingle: ts.RTSStats,
		resTypePlural: ts.RTPStats,
		objLookup:     lookup.L.Stats,
		objLookupID:   lookup.L.StatsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetStatIDs,
		getSingleFunc: getStat,
		retrieveFunc:  retrieveStats,
	}

	E.Properties = HandlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpProperties,
		resTypeSingle: ts.RTSProperties,
		resTypePlural: ts.RTPProperties,
		objLookup:     lookup.L.Properties,
		objLookupID:   lookup.L.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetPropertyIDs,
		getSingleFunc: getProperty,
		retrieveFunc:  retrieveProperties,
	}

	E.OverdriveModes = HandlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdriveModes,
		resTypeSingle: ts.RTSOverdriveModes,
		resTypePlural: ts.RTPOverdriveModes,
		objLookup:     lookup.L.OverdriveModes,
		objLookupID:   lookup.L.OverdriveModesID,
		idToResFunc:   idToNamedAPIResource[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetOverdriveModeIDs,
		getSingleFunc: getOverdriveMode,
		retrieveFunc:  retrieveOverdriveModes,
	}

	E.Elements = HandlerInput[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpElements,
		resTypeSingle: ts.RTSElements,
		resTypePlural: ts.RTPElements,
		objLookup:     lookup.L.Elements,
		objLookupID:   lookup.L.ElementsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetElementIDs,
		getSingleFunc: getElement,
		retrieveFunc:  retrieveElements,
	}

	E.StatusConditions = HandlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpStatusConditions,
		resTypeSingle: ts.RTSStatusConditions,
		resTypePlural: ts.RTPStatusConditions,
		objLookup:     lookup.L.StatusConditions,
		objLookupID:   lookup.L.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetStatusConditionIDs,
		getSingleFunc: getStatusCondition,
		retrieveFunc:  retrieveStatusConditions,
	}

	E.Modifiers = HandlerInput[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpModifiers,
		resTypeSingle: ts.RTSModifiers,
		resTypePlural: ts.RTPModifiers,
		objLookup:     lookup.L.Modifiers,
		objLookupID:   lookup.L.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetModifierIDs,
		getSingleFunc: getModifier,
		retrieveFunc:  retrieveModifiers,
	}

	E.AgilityTiers = HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpAgilityTiers,
		resTypeSingle: ts.RTSAgilityTiers,
		resTypePlural: ts.RTPAgilityTiers,
		objLookup:     nil,
		objLookupID:   lookup.L.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource],
		retrieveQuery: config.Cfg.DB.GetAgilityTierIDs,
		getSingleFunc: getAgilityTier,
		retrieveFunc:  retrieveAgilityTiers,
	}

	E.Enums = HandlerInputEnums{
		Endpoint: ts.EpEnums,
		usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		lookup:   enums.T.Lookup,
	}

	E.Endpoints = HandlerInputEndpoints{
		Endpoint: ts.EpEndpoints,
		usage:    []string{"/api/endpoints"},
	}

	E.AlBhed = HandlerInputService[AlBhedParams, AlBhedResponse]{
		Endpoint:  ts.EpAlBhed,
		verifyFn:  verifyAlBhedParams,
		executeFn: handleAlBhed,
		docsFn:    getAlBhedParamsDoc,
		samplesFn: getAlBhedParamsSamples,
	}

	E.TurnOrder = HandlerInputService[TurnOrderParams, TurnOrderResponse]{
		Endpoint:  ts.EpTurnOrder,
		verifyFn:  verifyTurnOrderParams,
		executeFn: handleTurnOrder,
		docsFn:    getTurnOrderParamsDoc,
		samplesFn: getTurnOrderParamsSamples,
	}

	E.DropChance = HandlerInputService[DropChanceParams, DropChanceResponse]{
		Endpoint:  ts.EpDropChance,
		verifyFn:  verifyDropChanceParams,
		executeFn: handleDropChance,
		docsFn:    getDropChanceParamsDoc,
		samplesFn: getDropChanceParamsSamples,
	}

	E.CritChance = HandlerInputService[CritChanceParams, CritChanceResponse]{
		Endpoint:  ts.EpCritChance,
		verifyFn:  verifyCritChanceParams,
		executeFn: handleCritChance,
		docsFn:    getCritChanceParamsDoc,
		samplesFn: getCritChanceParamsSamples,
	}

	E.InitializeEndpointSlice()
}

func (E *Endpoints) InitializeEndpointSlice() {
	E.Endpoints.Slice = []ts.EndpointName{
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

	slices.Sort(E.Endpoints.Slice)
}
