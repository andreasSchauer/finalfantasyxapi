package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/serviceinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/serviceverify"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

var E = &Endpoints{}

type Endpoints struct {
	Locations          handlerinput.HandlerInput[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource]
	Sublocations       handlerinput.HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource]
	Areas              handlerinput.HandlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource]
	MonsterFormations  handlerinput.HandlerInput[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource]
	Shops              handlerinput.HandlerInput[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource]
	Treasures          handlerinput.HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource]
	Quests             handlerinput.HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource]
	Sidequests         handlerinput.HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource]
	Subquests          handlerinput.HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource]
	ArenaCreations     handlerinput.HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource]
	BlitzballPrizes    handlerinput.HandlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource]
	Songs              handlerinput.HandlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource]
	FMVs               handlerinput.HandlerInput[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource]
	PlayerUnits        handlerinput.HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource]
	Characters         handlerinput.HandlerInput[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource]
	Aeons              handlerinput.HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource]
	CharacterClasses   handlerinput.HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource]
	Monsters           handlerinput.HandlerInput[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource]
	Abilities          handlerinput.HandlerInput[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource]
	PlayerAbilities    handlerinput.HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource]
	OverdriveAbilities handlerinput.HandlerInput[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource]
	ItemAbilities      handlerinput.HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource]
	TriggerCommands    handlerinput.HandlerInput[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource]
	MiscAbilities      handlerinput.HandlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource]
	EnemyAbilities     handlerinput.HandlerInput[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource]
	AeonCommands       handlerinput.HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource]
	OverdriveCommands  handlerinput.HandlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource]
	Overdrives         handlerinput.HandlerInput[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource]
	RonsoRages         handlerinput.HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]
	Submenus           handlerinput.HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource]
	Topmenus           handlerinput.HandlerInput[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource]
	AllItems           handlerinput.HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource]
	Items              handlerinput.HandlerInput[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource]
	KeyItems           handlerinput.HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource]
	Spheres            handlerinput.HandlerInput[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource]
	Primers            handlerinput.HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource]
	Mixes              handlerinput.HandlerInput[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource]
	AutoAbilities      handlerinput.HandlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource]
	EquipmentTables    handlerinput.HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource]
	Equipment          handlerinput.HandlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource]
	CelestialWeapons   handlerinput.HandlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource]
	Stats              handlerinput.HandlerInput[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource]
	Properties         handlerinput.HandlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource]
	OverdriveModes     handlerinput.HandlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource]
	Elements           handlerinput.HandlerInput[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource]
	StatusConditions   handlerinput.HandlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource]
	Modifiers          handlerinput.HandlerInput[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource]
	AgilityTiers       handlerinput.HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource]
	Enums              handlerinput.HandlerInputEnums
	Endpoints          handlerinput.HandlerInputEndpoints
	AlBhed             serviceinput.ServiceInput[servicetypes.AlBhedParams, servicetypes.AlBhedResponse]
	TurnOrder          serviceinput.ServiceInput[servicetypes.TurnOrderParams, servicetypes.TurnOrderResponse]
	DropChance         serviceinput.ServiceInput[servicetypes.DropChanceParams, servicetypes.DropChanceResponse]
	CritChance         serviceinput.ServiceInput[servicetypes.CritChanceParams, servicetypes.CritChanceResponse]
}

func (E *Endpoints) Init() {
	E.Locations = handlerinput.HandlerInput[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpLocations,
		ResTypeSingle: ts.RTSLocations,
		ResTypePlural: ts.RTPLocations,
		ObjLookup:     lookup.L.Locations,
		ObjLookupID:   lookup.L.LocationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Location, apitypes.Location, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetLocationIDs,
		GetSingleFunc: getLocation,
		RetrieveFunc:  retrieveLocations,
		AvlFunc:       filterAvlLocations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createLocationSimple,
				RelationsFn: getLocationSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetConnectedLocationIDs,
				CreateSubFn: createLocationSimple,
				RelationsFn: getLocationSectionRelations,
			},
			ts.SnSublocations: {
				DBQuery:     config.Cfg.DB.GetLocationSublocationIDs,
				CreateSubFn: createSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetLocationAreaIDs,
				CreateSubFn: createAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetLocationMonsterFormationIDs,
				CreateSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetLocationMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				DBQuery:     config.Cfg.DB.GetLocationShopIDs,
				CreateSubFn: createShopSimple,
			},
			ts.SnSongs: {
				DBQuery:     getLocationSongIDs,
				CreateSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetLocationTreasureIDs,
				CreateSubFn: createTreasureSimple,
			},
		},
	}

	E.Sublocations = handlerinput.HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSublocations,
		ResTypeSingle: ts.RTSSublocations,
		ResTypePlural: ts.RTPSublocations,
		ObjLookup:     lookup.L.Sublocations,
		ObjLookupID:   lookup.L.SublocationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSublocationIDs,
		GetSingleFunc: getSublocation,
		RetrieveFunc:  retrieveSublocations,
		AvlFunc:       filterAvlSublocations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetConnectedSublocationIDs,
				CreateSubFn: createSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetSublocationAreaIDs,
				CreateSubFn: createAreaSimple,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetSublocationMonsterFormationIDs,
				CreateSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetSublocationMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				DBQuery:     config.Cfg.DB.GetSublocationShopIDs,
				CreateSubFn: createShopSimple,
			},
			ts.SnSongs: {
				DBQuery:     getSublocationSongIDs,
				CreateSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetSublocationTreasureIDs,
				CreateSubFn: createTreasureSimple,
			},
		},
	}

	E.Areas = handlerinput.HandlerInput[seedtypes.Area, apitypes.Area, apiresource.AreaAPIResource]{
		Endpoint:      ts.EpAreas,
		ResTypeSingle: ts.RTSAreas,
		ResTypePlural: ts.RTPAreas,
		ObjLookup:     lookup.L.Areas,
		ObjLookupID:   lookup.L.AreasID,
		IDToResFunc:   handlerinput.IDToAreaAPIResource,
		RetrieveQuery: config.Cfg.DB.GetAreaIDs,
		GetSingleFunc: getArea,
		RetrieveFunc:  retrieveAreas,
		AvlFunc:       filterAvlAreas,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetAreaConnectionIDs,
				CreateSubFn: createAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetAreaMonsterFormationIDs,
				CreateSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetAreaMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnSongs: {
				DBQuery:     getAreaSongIDs,
				CreateSubFn: createSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetAreaTreasureIDs,
				CreateSubFn: createTreasureSimple,
			},
		},
	}

	E.MonsterFormations = handlerinput.HandlerInput[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpMonsterFormations,
		ResTypeSingle: ts.RTSMonsterFormations,
		ResTypePlural: ts.RTPMonsterFormations,
		ObjLookup:     lookup.L.MonsterFormations,
		ObjLookupID:   lookup.L.MonsterFormationsID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.UnnamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetMonsterFormationIDs,
		GetSingleFunc: getMonsterFormation,
		RetrieveFunc:  retrieveMonsterFormations,
		AvlFunc:       filterAvlMonsterFormations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetMonsterFormationMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Shops = handlerinput.HandlerInput[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpShops,
		ResTypeSingle: ts.RTSShops,
		ResTypePlural: ts.RTPShops,
		ObjLookup:     lookup.L.Shops,
		ObjLookupID:   lookup.L.ShopsID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.Shop, apitypes.Shop, apiresource.UnnamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetShopIDs,
		GetSingleFunc: getShop,
		RetrieveFunc:  retrieveShops,
		AvlFunc:       filterAvlShops,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createShopSimple,
			},
		},
	}

	E.Treasures = handlerinput.HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpTreasures,
		ResTypeSingle: ts.RTSTreasures,
		ResTypePlural: ts.RTPTreasures,
		ObjLookup:     lookup.L.Treasures,
		ObjLookupID:   lookup.L.TreasuresID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetTreasureIDs,
		GetSingleFunc: getTreasure,
		AvlFunc:       filterAvlTreasures,
		RetrieveFunc:  retrieveTreasures,
	}

	E.Quests = handlerinput.HandlerInput[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpQuests,
		ResTypeSingle: ts.RTSQuests,
		ResTypePlural: ts.RTPQuests,
		ObjLookup:     lookup.L.Quests,
		ObjLookupID:   lookup.L.QuestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Quest, apitypes.Quest, apiresource.QuestAPIResource],
		RetrieveQuery: config.Cfg.DB.GetQuestIDs,
		GetSingleFunc: getQuest,
		RetrieveFunc:  retrieveQuests,
		AvlFunc:       filterAvlQuests,
	}

	E.Sidequests = handlerinput.HandlerInput[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpSidequests,
		ResTypeSingle: ts.RTSSidequests,
		ResTypePlural: ts.RTPSidequests,
		ObjLookup:     lookup.L.Sidequests,
		ObjLookupID:   lookup.L.SidequestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Sidequest, apitypes.Sidequest, apiresource.QuestAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSidequestIDs,
		GetSingleFunc: getSidequest,
		RetrieveFunc:  retrieveSidequests,
		AvlFunc:       filterAvlSidequests,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSubquests: {
				DBQuery:     config.Cfg.DB.GetSidequestSubquestIDs,
				CreateSubFn: createSubquestSimple,
			},
		},
	}

	E.Subquests = handlerinput.HandlerInput[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource]{
		Endpoint:      ts.EpSubquests,
		ResTypeSingle: ts.RTSSubquests,
		ResTypePlural: ts.RTPSubquests,
		ObjLookup:     lookup.L.Subquests,
		ObjLookupID:   lookup.L.SubquestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Subquest, apitypes.Subquest, apiresource.QuestAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSubquestIDs,
		GetSingleFunc: getSubquest,
		RetrieveFunc:  retrieveSubquests,
		AvlFunc:       filterAvlSubquests,
	}

	E.ArenaCreations = handlerinput.HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpArenaCreations,
		ResTypeSingle: ts.RTSArenaCreations,
		ResTypePlural: ts.RTPArenaCreations,
		ObjLookup:     lookup.L.ArenaCreations,
		ObjLookupID:   lookup.L.ArenaCreationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetArenaCreationIDs,
		GetSingleFunc: getArenaCreation,
		RetrieveFunc:  retrieveArenaCreations,
	}

	E.BlitzballPrizes = handlerinput.HandlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpBlitzballPrizes,
		ResTypeSingle: ts.RTSBlitzballPrizes,
		ResTypePlural: ts.RTPBlitzballPrizes,
		ObjLookup:     lookup.L.Positions,
		ObjLookupID:   lookup.L.PositionsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetBlitzballPrizeIDs,
		GetSingleFunc: getBlitzballPrize,
		RetrieveFunc:  retrieveBlitzballPrizes,
	}

	E.Songs = handlerinput.HandlerInput[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSongs,
		ResTypeSingle: ts.RTSSongs,
		ResTypePlural: ts.RTPSongs,
		ObjLookup:     lookup.L.Songs,
		ObjLookupID:   lookup.L.SongsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Song, apitypes.Song, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSongIDs,
		GetSingleFunc: getSong,
		RetrieveFunc:  retrieveSongs,
	}

	E.FMVs = handlerinput.HandlerInput[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpFMVs,
		ResTypeSingle: ts.RTSFMVs,
		ResTypePlural: ts.RTPFMVs,
		ObjLookup:     lookup.L.FMVs,
		ObjLookupID:   lookup.L.FMVsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.FMV, apitypes.FMV, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetFmvIDs,
		GetSingleFunc: getFMV,
		RetrieveFunc:  retrieveFMVs,
	}

	E.PlayerUnits = handlerinput.HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpPlayerUnits,
		ResTypeSingle: ts.RTSPlayerUnits,
		ResTypePlural: ts.RTPPlayerUnits,
		ObjLookup:     lookup.L.PlayerUnits,
		ObjLookupID:   lookup.L.PlayerUnitsID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetPlayerUnitIDs,
		GetSingleFunc: getPlayerUnit,
		RetrieveFunc:  retrievePlayerUnits,
	}

	E.Characters = handlerinput.HandlerInput[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCharacters,
		ResTypeSingle: ts.RTSCharacters,
		ResTypePlural: ts.RTPCharacters,
		ObjLookup:     lookup.L.Characters,
		ObjLookupID:   lookup.L.CharactersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Character, apitypes.Character, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetCharacterIDs,
		GetSingleFunc: getCharacter,
		RetrieveFunc:  retrieveCharacters,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterDefaultAbilityIDs,
				CreateSubFn: createPlayerAbilitySimple,
			},
			ts.SnStdSgAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterSgAbilityIDs,
				CreateSubFn: createPlayerAbilitySimple,
			},
			ts.SnExpSgAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterEgAbilityIDs,
				CreateSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterOverdriveAbilityIDs,
				CreateSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterOverdriveIDs,
				CreateSubFn: createOverdriveSimple,
			},
		},
	}

	E.Aeons = handlerinput.HandlerInput[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAeons,
		ResTypeSingle: ts.RTSAeons,
		ResTypePlural: ts.RTPAeons,
		ObjLookup:     lookup.L.Aeons,
		ObjLookupID:   lookup.L.AeonsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Aeon, apitypes.Aeon, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetAeonIDs,
		GetSingleFunc: getAeon,
		RetrieveFunc:  retrieveAeons,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetAeonDefaultAbilityIDs,
				CreateSubFn: createPlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetAeonOverdriveAbilityIDs,
				CreateSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     config.Cfg.DB.GetAeonOverdriveIDs,
				CreateSubFn: createOverdriveSimple,
			},
			ts.SnStats: {
				CreateSubFn: createAeonStatSimple,
			},
		},
	}

	E.CharacterClasses = handlerinput.HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCharacterClasses,
		ResTypeSingle: ts.RTSCharacterClasses,
		ResTypePlural: ts.RTPCharacterClasses,
		ObjLookup:     lookup.L.CharClasses,
		ObjLookupID:   lookup.L.CharClassesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetCharacterClassesIDs,
		GetSingleFunc: getCharacterClass,
		RetrieveFunc:  retrieveCharacterClasses,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterClassDefaultAbilityIDs,
				CreateSubFn: createAbilitySimple,
			},
			ts.SnLearnableAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterClassLearnableAbilityIDs,
				CreateSubFn: createAbilitySimple,
			},
			ts.SnDefaultOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs,
				CreateSubFn: createOverdriveSimple,
			},
			ts.SnLearnableOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs,
				CreateSubFn: createOverdriveSimple,
			},
		},
	}

	E.Monsters = handlerinput.HandlerInput[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpMonsters,
		ResTypeSingle:    ts.RTSMonsters,
		ResTypePlural:    ts.RTPMonsters,
		ObjLookup:        lookup.L.Monsters,
		ObjLookupID:      lookup.L.MonstersID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.Monster, apitypes.Monster, apiresource.NamedAPIResource],
		GetMultipleQuery: config.Cfg.DB.GetMonsterIDsByName,
		RetrieveQuery:    config.Cfg.DB.GetMonsterIDs,
		GetSingleFunc:    getMonster,
		RetrieveFunc:     retrieveMonsters,
		AvlFunc:          filterAvlMonsters,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnAbilities: {
				DBQuery:     config.Cfg.DB.GetMonsterAbilityIDs,
				CreateSubFn: createAbilitySimple,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetMonsterAreaIDs,
				CreateSubFn: createAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetMonsterMonsterFormationIDs,
				CreateSubFn: createMonsterFormationSimple,
			},
		},
	}

	E.Abilities = handlerinput.HandlerInput[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpAbilities,
		ResTypeSingle: ts.RTSAbilities,
		ResTypePlural: ts.RTPAbilities,
		ObjLookup:     lookup.L.Abilities,
		ObjLookupID:   lookup.L.AbilitiesID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.Ability, apitypes.Ability, apiresource.TypedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetAbilityIDs,
		GetSingleFunc: getAbility,
		RetrieveFunc:  retrieveAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createAbilitySimple,
			},
		},
	}

	E.PlayerAbilities = handlerinput.HandlerInput[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpPlayerAbilities,
		ResTypeSingle:    ts.RTSPlayerAbilities,
		ResTypePlural:    ts.RTPPlayerAbilities,
		ObjLookup:        lookup.L.PlayerAbilities,
		ObjLookupID:      lookup.L.PlayerAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.NamedAPIResource],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypePlayerAbility),
		RetrieveQuery:    config.Cfg.DB.GetPlayerAbilityIDs,
		GetSingleFunc:    getPlayerAbility,
		RetrieveFunc:     retrievePlayerAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createPlayerAbilitySimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetPlayerAbilityMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.OverdriveAbilities = handlerinput.HandlerInput[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpOverdriveAbilities,
		ResTypeSingle:    ts.RTSOverdriveAbilities,
		ResTypePlural:    ts.RTPOverdriveAbilities,
		ObjLookup:        lookup.L.OverdriveAbilities,
		ObjLookupID:      lookup.L.OverdriveAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.NamedAPIResource],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeOverdriveAbility),
		RetrieveQuery:    config.Cfg.DB.GetOverdriveAbilityIDs,
		GetSingleFunc:    getOverdriveAbility,
		RetrieveFunc:     retrieveOverdriveAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.ItemAbilities = handlerinput.HandlerInput[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpItemAbilities,
		ResTypeSingle: ts.RTSItemAbilities,
		ResTypePlural: ts.RTPItemAbilities,
		ObjLookup:     lookup.L.ItemAbilities,
		ObjLookupID:   lookup.L.ItemAbilitiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetItemAbilityIDs,
		GetSingleFunc: getItemAbility,
		RetrieveFunc:  retrieveItemAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createItemAbilitySimple,
			},
		},
	}

	E.TriggerCommands = handlerinput.HandlerInput[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpTriggerCommands,
		ResTypeSingle:    ts.RTSTriggerCommands,
		ResTypePlural:    ts.RTPTriggerCommands,
		ObjLookup:        lookup.L.TriggerCommands,
		ObjLookupID:      lookup.L.TriggerCommandsID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.NamedAPIResource],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeTriggerCommand),
		RetrieveQuery:    config.Cfg.DB.GetTriggerCommandIDs,
		GetSingleFunc:    getTriggerCommand,
		RetrieveFunc:     retrieveTriggerCommands,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createTriggerCommandSimple,
			},
		},
	}

	E.MiscAbilities = handlerinput.HandlerInput[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpMiscAbilities,
		ResTypeSingle:    ts.RTSMiscAbilities,
		ResTypePlural:    ts.RTPMiscAbilities,
		ObjLookup:        lookup.L.MiscAbilities,
		ObjLookupID:      lookup.L.MiscAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.NamedAPIResource],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeMiscAbility),
		RetrieveQuery:    config.Cfg.DB.GetMiscAbilityIDs,
		GetSingleFunc:    getMiscAbility,
		RetrieveFunc:     retrieveMiscAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createMiscAbilitySimple,
			},
		},
	}

	E.EnemyAbilities = handlerinput.HandlerInput[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource]{
		Endpoint:         ts.EpEnemyAbilities,
		ResTypeSingle:    ts.RTSEnemyAbilities,
		ResTypePlural:    ts.RTPEnemyAbilities,
		ObjLookup:        lookup.L.EnemyAbilities,
		ObjLookupID:      lookup.L.EnemyAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.NamedAPIResource],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeEnemyAbility),
		RetrieveQuery:    config.Cfg.DB.GetEnemyAbilityIDs,
		GetSingleFunc:    getEnemyAbility,
		RetrieveFunc:     retrieveEnemyAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createEnemyAbilitySimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetEnemyAbilityMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.AeonCommands = handlerinput.HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAeonCommands,
		ResTypeSingle: ts.RTSAeonCommands,
		ResTypePlural: ts.RTPAeonCommands,
		ObjLookup:     lookup.L.AeonCommands,
		ObjLookupID:   lookup.L.AeonCommandsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetAeonCommandIDs,
		GetSingleFunc: getAeonCommand,
		RetrieveFunc:  retrieveAeonCommands,
	}

	E.OverdriveCommands = handlerinput.HandlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdriveCommands,
		ResTypeSingle: ts.RTSOverdriveCommands,
		ResTypePlural: ts.RTPOverdriveCommands,
		ObjLookup:     lookup.L.OverdriveCommands,
		ObjLookupID:   lookup.L.OverdriveCommandsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetOverdriveCommandIDs,
		GetSingleFunc: getOverdriveCommand,
		RetrieveFunc:  retrieveOverdriveCommands,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnOverdriveAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveAbilityIDs),
				CreateSubFn: createOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs),
				CreateSubFn: createOverdriveSimple,
			},
		},
	}

	E.Overdrives = handlerinput.HandlerInput[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdrives,
		ResTypeSingle: ts.RTSOverdrives,
		ResTypePlural: ts.RTPOverdrives,
		ObjLookup:     lookup.L.Overdrives,
		ObjLookupID:   lookup.L.OverdrivesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Overdrive, apitypes.Overdrive, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetOverdriveIDs,
		GetSingleFunc: getOverdrive,
		RetrieveFunc:  retrieveOverdrives,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createOverdriveSimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetOverdriveOverdriveAbilityIDs,
				CreateSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	E.RonsoRages = handlerinput.HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpRonsoRages,
		ResTypeSingle: ts.RTSRonsoRages,
		ResTypePlural: ts.RTPRonsoRages,
		ObjLookup:     lookup.L.RonsoRages,
		ObjLookupID:   lookup.L.RonsoRagesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetRonsoRageIDs,
		GetSingleFunc: getRonsoRage,
		RetrieveFunc:  retrieveRonsoRages,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetRonsoRageMonsterIDs,
				CreateSubFn: createMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	E.Submenus = handlerinput.HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSubmenus,
		ResTypeSingle: ts.RTSSubmenus,
		ResTypePlural: ts.RTPSubmenus,
		ObjLookup:     lookup.L.Submenus,
		ObjLookupID:   lookup.L.SubmenusID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSubmenuIDs,
		GetSingleFunc: getSubmenu,
		RetrieveFunc:  retrieveSubmenus,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetSubmenuAbilityIDs),
				CreateSubFn: createAbilitySimple,
			},
		},
	}

	E.Topmenus = handlerinput.HandlerInput[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpTopmenus,
		ResTypeSingle: ts.RTSTopmenus,
		ResTypePlural: ts.RTPTopmenus,
		ObjLookup:     lookup.L.Topmenus,
		ObjLookupID:   lookup.L.TopmenusID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Topmenu, apitypes.Topmenu, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetTopmenuIDs,
		GetSingleFunc: getTopmenu,
		RetrieveFunc:  retrieveTopmenus,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs),
				CreateSubFn: createAbilitySimple,
			},
		},
	}

	E.AllItems = handlerinput.HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource]{
		Endpoint:      ts.EpAllItems,
		ResTypeSingle: ts.RTSAllItems,
		ResTypePlural: ts.RTPAllItems,
		ObjLookup:     lookup.L.MasterItems,
		ObjLookupID:   lookup.L.MasterItemsID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetMasterItemIDs,
		GetSingleFunc: getMasterItem,
		RetrieveFunc:  retrieveMasterItems,
		AvlFunc:       filterAvlMasterItems,
	}

	E.Items = handlerinput.HandlerInput[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpItems,
		ResTypeSingle: ts.RTSItems,
		ResTypePlural: ts.RTPItems,
		ObjLookup:     lookup.L.Items,
		ObjLookupID:   lookup.L.ItemsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Item, apitypes.Item, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetItemIDs,
		GetSingleFunc: getItem,
		RetrieveFunc:  retrieveItems,
		AvlFunc:       filterAvlItems,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnMixes: {
				DBQuery:     config.Cfg.DB.GetItemMixIDs,
				CreateSubFn: createMixSimple,
			},
		},
	}

	E.Spheres = handlerinput.HandlerInput[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpSpheres,
		ResTypeSingle: ts.RTSSpheres,
		ResTypePlural: ts.RTPSpheres,
		ObjLookup:     lookup.L.Spheres,
		ObjLookupID:   lookup.L.SpheresID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Sphere, apitypes.Sphere, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetSphereIDs,
		GetSingleFunc: getSphere,
		RetrieveFunc:  retrieveSpheres,
		AvlFunc:       filterAvlSpheres,
	}

	E.KeyItems = handlerinput.HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpKeyItems,
		ResTypeSingle: ts.RTSKeyItems,
		ResTypePlural: ts.RTPKeyItems,
		ObjLookup:     lookup.L.KeyItems,
		ObjLookupID:   lookup.L.KeyItemsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetKeyItemIDs,
		GetSingleFunc: getKeyItem,
		RetrieveFunc:  retrieveKeyItems,
		AvlFunc:       filterAvlKeyItems,
	}

	E.Primers = handlerinput.HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpPrimers,
		ResTypeSingle: ts.RTSPrimers,
		ResTypePlural: ts.RTPPrimers,
		ObjLookup:     lookup.L.Primers,
		ObjLookupID:   lookup.L.PrimersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetPrimerIDs,
		GetSingleFunc: getPrimer,
		RetrieveFunc:  retrievePrimers,
		AvlFunc:       filterAvlPrimers,
	}

	E.Mixes = handlerinput.HandlerInput[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpMixes,
		ResTypeSingle: ts.RTSMixes,
		ResTypePlural: ts.RTPMixes,
		ObjLookup:     lookup.L.Mixes,
		ObjLookupID:   lookup.L.MixesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Mix, apitypes.Mix, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetMixIDs,
		GetSingleFunc: getMix,
		RetrieveFunc:  retrieveMixes,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createMixSimple,
			},
		},
	}

	E.AutoAbilities = handlerinput.HandlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpAutoAbilities,
		ResTypeSingle: ts.RTSAutoAbilities,
		ResTypePlural: ts.RTPAutoAbilities,
		ObjLookup:     lookup.L.AutoAbilities,
		ObjLookupID:   lookup.L.AutoAbilitiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetAutoAbilityIDs,
		GetSingleFunc: getAutoAbility,
		RetrieveFunc:  retrieveAutoAbilities,
		AvlFunc:       filterAvlAutoAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.EquipmentTables = handlerinput.HandlerInput[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpEquipmentTables,
		ResTypeSingle: ts.RTSEquipmentTables,
		ResTypePlural: ts.RTPEquipmentTables,
		ObjLookup:     lookup.L.EquipmentTables,
		ObjLookupID:   lookup.L.EquipmentTablesID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.UnnamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetEquipmentTableIDs,
		GetSingleFunc: getEquipmentTable,
		RetrieveFunc:  retrieveEquipmentTables,
	}

	E.Equipment = handlerinput.HandlerInput[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpEquipment,
		ResTypeSingle: ts.RTSEquipment,
		ResTypePlural: ts.RTPEquipment,
		ObjLookup:     lookup.L.EquipmentNames,
		ObjLookupID:   lookup.L.EquipmentNamesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetEquipmentIDs,
		GetSingleFunc: getEquipment,
		RetrieveFunc:  retrieveEquipment,
	}

	E.CelestialWeapons = handlerinput.HandlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpCelestialWeapons,
		ResTypeSingle: ts.RTSCelestialWeapons,
		ResTypePlural: ts.RTPCelestialWeapons,
		ObjLookup:     lookup.L.CelestialWeapons,
		ObjLookupID:   lookup.L.CelestialWeaponsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetCelestialWeaponIDs,
		GetSingleFunc: getCelestialWeapon,
		RetrieveFunc:  retrieveCelestialWeapons,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAutoAbilities: {
				DBQuery:     convertGetCelestialWeaponAutoAbilityIDs(),
				CreateSubFn: createAutoAbilitySimple,
			},
		},
	}

	E.Stats = handlerinput.HandlerInput[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpStats,
		ResTypeSingle: ts.RTSStats,
		ResTypePlural: ts.RTPStats,
		ObjLookup:     lookup.L.Stats,
		ObjLookupID:   lookup.L.StatsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Stat, apitypes.Stat, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetStatIDs,
		GetSingleFunc: getStat,
		RetrieveFunc:  retrieveStats,
	}

	E.Properties = handlerinput.HandlerInput[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpProperties,
		ResTypeSingle: ts.RTSProperties,
		ResTypePlural: ts.RTPProperties,
		ObjLookup:     lookup.L.Properties,
		ObjLookupID:   lookup.L.PropertiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Property, apitypes.Property, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetPropertyIDs,
		GetSingleFunc: getProperty,
		RetrieveFunc:  retrieveProperties,
	}

	E.OverdriveModes = handlerinput.HandlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpOverdriveModes,
		ResTypeSingle: ts.RTSOverdriveModes,
		ResTypePlural: ts.RTPOverdriveModes,
		ObjLookup:     lookup.L.OverdriveModes,
		ObjLookupID:   lookup.L.OverdriveModesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetOverdriveModeIDs,
		GetSingleFunc: getOverdriveMode,
		RetrieveFunc:  retrieveOverdriveModes,
	}

	E.Elements = handlerinput.HandlerInput[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpElements,
		ResTypeSingle: ts.RTSElements,
		ResTypePlural: ts.RTPElements,
		ObjLookup:     lookup.L.Elements,
		ObjLookupID:   lookup.L.ElementsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Element, apitypes.Element, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetElementIDs,
		GetSingleFunc: getElement,
		RetrieveFunc:  retrieveElements,
	}

	E.StatusConditions = handlerinput.HandlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpStatusConditions,
		ResTypeSingle: ts.RTSStatusConditions,
		ResTypePlural: ts.RTPStatusConditions,
		ObjLookup:     lookup.L.StatusConditions,
		ObjLookupID:   lookup.L.StatusConditionsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetStatusConditionIDs,
		GetSingleFunc: getStatusCondition,
		RetrieveFunc:  retrieveStatusConditions,
	}

	E.Modifiers = handlerinput.HandlerInput[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource]{
		Endpoint:      ts.EpModifiers,
		ResTypeSingle: ts.RTSModifiers,
		ResTypePlural: ts.RTPModifiers,
		ObjLookup:     lookup.L.Modifiers,
		ObjLookupID:   lookup.L.ModifiersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Modifier, apitypes.Modifier, apiresource.NamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetModifierIDs,
		GetSingleFunc: getModifier,
		RetrieveFunc:  retrieveModifiers,
	}

	E.AgilityTiers = handlerinput.HandlerInput[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource]{
		Endpoint:      ts.EpAgilityTiers,
		ResTypeSingle: ts.RTSAgilityTiers,
		ResTypePlural: ts.RTPAgilityTiers,
		ObjLookup:     nil,
		ObjLookupID:   lookup.L.AgilityTiersID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.UnnamedAPIResource],
		RetrieveQuery: config.Cfg.DB.GetAgilityTierIDs,
		GetSingleFunc: getAgilityTier,
		RetrieveFunc:  retrieveAgilityTiers,
	}

	E.Enums = handlerinput.HandlerInputEnums{
		Endpoint: ts.EpEnums,
		Usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		Lookup:   enums.T.Lookup,
	}

	E.Endpoints = handlerinput.HandlerInputEndpoints{
		Endpoint: ts.EpEndpoints,
		Usage:    []string{"/api/endpoints"},
	}

	E.AlBhed = serviceinput.ServiceInput[servicetypes.AlBhedParams, servicetypes.AlBhedResponse]{
		Endpoint:  ts.EpAlBhed,
		VerifyFn:  serviceverify.VerifyAlBhedParams,
		ExecuteFn: handleAlBhed,
		DocsFn:    servicetypes.GetAlBhedParamsDoc,
		Samples:   servicetypes.AlBhedParamsSamples,
	}

	E.TurnOrder = serviceinput.ServiceInput[servicetypes.TurnOrderParams, servicetypes.TurnOrderResponse]{
		Endpoint:  ts.EpTurnOrder,
		VerifyFn:  serviceverify.VerifyTurnOrderParams,
		ExecuteFn: handleTurnOrder,
		DocsFn:    servicetypes.GetTurnOrderParamsDoc,
		Samples:   servicetypes.TurnOrderParamsSamples,
	}

	E.DropChance = serviceinput.ServiceInput[servicetypes.DropChanceParams, servicetypes.DropChanceResponse]{
		Endpoint:  ts.EpDropChance,
		VerifyFn:  serviceverify.VerifyDropChanceParams,
		ExecuteFn: handleDropChance,
		DocsFn:    servicetypes.GetDropChanceParamsDoc,
		Samples:   servicetypes.DropChanceParamsSamples,
	}

	E.CritChance = serviceinput.ServiceInput[servicetypes.CritChanceParams, servicetypes.CritChanceResponse]{
		Endpoint:  ts.EpCritChance,
		VerifyFn:  serviceverify.VerifyCritChanceParams,
		ExecuteFn: handleCritChance,
		DocsFn:    servicetypes.GetCritChanceParamsDoc,
		Samples:   servicetypes.CritChanceParamsSamples,
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
