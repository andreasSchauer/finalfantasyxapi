package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/serviceverify"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func InitEndpoints() {
	apilookups.E.Locations = handlerinput.Main[seedtypes.Location, apitypes.Location, apiresource.Named]{
		Endpoint:      ts.EpLocations,
		ResTypeSingle: ts.RTSLocations,
		ResTypePlural: ts.RTPLocations,
		ObjLookup:     lookup.L.Locations,
		ObjLookupID:   lookup.L.LocationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Location, apitypes.Location, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetLocationIDs,
		GetSingleFunc: getLocation,
		RetrieveFunc:  retrieveLocations,
		AvlFunc:       filterAvlLocations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateLocationSimple,
				RelationsFn: getLocationSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetConnectedLocationIDs,
				CreateSubFn: CreateLocationSimple,
				RelationsFn: getLocationSectionRelations,
			},
			ts.SnSublocations: {
				DBQuery:     config.Cfg.DB.GetLocationSublocationIDs,
				CreateSubFn: CreateSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetLocationAreaIDs,
				CreateSubFn: CreateAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetLocationMonsterFormationIDs,
				CreateSubFn: CreateMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetLocationMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				DBQuery:     config.Cfg.DB.GetLocationShopIDs,
				CreateSubFn: CreateShopSimple,
			},
			ts.SnSongs: {
				DBQuery:     getLocationSongIDs,
				CreateSubFn: CreateSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetLocationTreasureIDs,
				CreateSubFn: CreateTreasureSimple,
			},
		},
	}

	apilookups.E.Sublocations = handlerinput.Main[seedtypes.Sublocation, apitypes.Sublocation, apiresource.Named]{
		Endpoint:      ts.EpSublocations,
		ResTypeSingle: ts.RTSSublocations,
		ResTypePlural: ts.RTPSublocations,
		ObjLookup:     lookup.L.Sublocations,
		ObjLookupID:   lookup.L.SublocationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Sublocation, apitypes.Sublocation, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetSublocationIDs,
		GetSingleFunc: getSublocation,
		RetrieveFunc:  retrieveSublocations,
		AvlFunc:       filterAvlSublocations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetConnectedSublocationIDs,
				CreateSubFn: CreateSublocationSimple,
				RelationsFn: getSublocationSectionRelations,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetSublocationAreaIDs,
				CreateSubFn: CreateAreaSimple,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetSublocationMonsterFormationIDs,
				CreateSubFn: CreateMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetSublocationMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnShops: {
				DBQuery:     config.Cfg.DB.GetSublocationShopIDs,
				CreateSubFn: CreateShopSimple,
			},
			ts.SnSongs: {
				DBQuery:     getSublocationSongIDs,
				CreateSubFn: CreateSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetSublocationTreasureIDs,
				CreateSubFn: CreateTreasureSimple,
			},
		},
	}

	apilookups.E.Areas = handlerinput.Main[seedtypes.Area, apitypes.Area, apiresource.Area]{
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
				CreateSubFn: CreateAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnConnected: {
				DBQuery:     config.Cfg.DB.GetAreaConnectionIDs,
				CreateSubFn: CreateAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetAreaMonsterFormationIDs,
				CreateSubFn: CreateMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetAreaMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnSongs: {
				DBQuery:     getAreaSongIDs,
				CreateSubFn: CreateSongSimple,
			},
			ts.SnTreasures: {
				DBQuery:     config.Cfg.DB.GetAreaTreasureIDs,
				CreateSubFn: CreateTreasureSimple,
			},
		},
	}

	apilookups.E.MonsterFormations = handlerinput.Main[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed]{
		Endpoint:      ts.EpMonsterFormations,
		ResTypeSingle: ts.RTSMonsterFormations,
		ResTypePlural: ts.RTPMonsterFormations,
		ObjLookup:     lookup.L.MonsterFormations,
		ObjLookupID:   lookup.L.MonsterFormationsID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed],
		RetrieveQuery: config.Cfg.DB.GetMonsterFormationIDs,
		GetSingleFunc: getMonsterFormation,
		RetrieveFunc:  retrieveMonsterFormations,
		AvlFunc:       filterAvlMonsterFormations,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateMonsterFormationSimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetMonsterFormationMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	apilookups.E.Shops = handlerinput.Main[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed]{
		Endpoint:      ts.EpShops,
		ResTypeSingle: ts.RTSShops,
		ResTypePlural: ts.RTPShops,
		ObjLookup:     lookup.L.Shops,
		ObjLookupID:   lookup.L.ShopsID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed],
		RetrieveQuery: config.Cfg.DB.GetShopIDs,
		GetSingleFunc: getShop,
		RetrieveFunc:  retrieveShops,
		AvlFunc:       filterAvlShops,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateShopSimple,
			},
		},
	}

	apilookups.E.Treasures = handlerinput.Main[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed]{
		Endpoint:      ts.EpTreasures,
		ResTypeSingle: ts.RTSTreasures,
		ResTypePlural: ts.RTPTreasures,
		ObjLookup:     lookup.L.Treasures,
		ObjLookupID:   lookup.L.TreasuresID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed],
		RetrieveQuery: config.Cfg.DB.GetTreasureIDs,
		GetSingleFunc: getTreasure,
		AvlFunc:       filterAvlTreasures,
		RetrieveFunc:  retrieveTreasures,
	}

	apilookups.E.Quests = handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest]{
		Endpoint:      ts.EpQuests,
		ResTypeSingle: ts.RTSQuests,
		ResTypePlural: ts.RTPQuests,
		ObjLookup:     lookup.L.Quests,
		ObjLookupID:   lookup.L.QuestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Quest, apitypes.Quest, apiresource.Quest],
		RetrieveQuery: config.Cfg.DB.GetQuestIDs,
		GetSingleFunc: getQuest,
		RetrieveFunc:  retrieveQuests,
		AvlFunc:       filterAvlQuests,
	}

	apilookups.E.Sidequests = handlerinput.Main[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest]{
		Endpoint:      ts.EpSidequests,
		ResTypeSingle: ts.RTSSidequests,
		ResTypePlural: ts.RTPSidequests,
		ObjLookup:     lookup.L.Sidequests,
		ObjLookupID:   lookup.L.SidequestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest],
		RetrieveQuery: config.Cfg.DB.GetSidequestIDs,
		GetSingleFunc: getSidequest,
		RetrieveFunc:  retrieveSidequests,
		AvlFunc:       filterAvlSidequests,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSubquests: {
				DBQuery:     config.Cfg.DB.GetSidequestSubquestIDs,
				CreateSubFn: CreateSubquestSimple,
			},
		},
	}

	apilookups.E.Subquests = handlerinput.Main[seedtypes.Subquest, apitypes.Subquest, apiresource.Quest]{
		Endpoint:      ts.EpSubquests,
		ResTypeSingle: ts.RTSSubquests,
		ResTypePlural: ts.RTPSubquests,
		ObjLookup:     lookup.L.Subquests,
		ObjLookupID:   lookup.L.SubquestsID,
		IDToResFunc:   handlerinput.IDToQuestAPIResource[seedtypes.Subquest, apitypes.Subquest, apiresource.Quest],
		RetrieveQuery: config.Cfg.DB.GetSubquestIDs,
		GetSingleFunc: getSubquest,
		RetrieveFunc:  retrieveSubquests,
		AvlFunc:       filterAvlSubquests,
	}

	apilookups.E.ArenaCreations = handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named]{
		Endpoint:      ts.EpArenaCreations,
		ResTypeSingle: ts.RTSArenaCreations,
		ResTypePlural: ts.RTPArenaCreations,
		ObjLookup:     lookup.L.ArenaCreations,
		ObjLookupID:   lookup.L.ArenaCreationsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetArenaCreationIDs,
		GetSingleFunc: getArenaCreation,
		RetrieveFunc:  retrieveArenaCreations,
	}

	apilookups.E.BlitzballPrizes = handlerinput.Main[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.Named]{
		Endpoint:      ts.EpBlitzballPrizes,
		ResTypeSingle: ts.RTSBlitzballPrizes,
		ResTypePlural: ts.RTPBlitzballPrizes,
		ObjLookup:     lookup.L.Positions,
		ObjLookupID:   lookup.L.PositionsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetBlitzballPrizeIDs,
		GetSingleFunc: getBlitzballPrize,
		RetrieveFunc:  retrieveBlitzballPrizes,
	}

	apilookups.E.Songs = handlerinput.Main[seedtypes.Song, apitypes.Song, apiresource.Named]{
		Endpoint:      ts.EpSongs,
		ResTypeSingle: ts.RTSSongs,
		ResTypePlural: ts.RTPSongs,
		ObjLookup:     lookup.L.Songs,
		ObjLookupID:   lookup.L.SongsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Song, apitypes.Song, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetSongIDs,
		GetSingleFunc: getSong,
		RetrieveFunc:  retrieveSongs,
	}

	apilookups.E.FMVs = handlerinput.Main[seedtypes.FMV, apitypes.FMV, apiresource.Named]{
		Endpoint:      ts.EpFMVs,
		ResTypeSingle: ts.RTSFMVs,
		ResTypePlural: ts.RTPFMVs,
		ObjLookup:     lookup.L.FMVs,
		ObjLookupID:   lookup.L.FMVsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.FMV, apitypes.FMV, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetFmvIDs,
		GetSingleFunc: getFMV,
		RetrieveFunc:  retrieveFMVs,
	}

	apilookups.E.PlayerUnits = handlerinput.Main[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.Typed]{
		Endpoint:      ts.EpPlayerUnits,
		ResTypeSingle: ts.RTSPlayerUnits,
		ResTypePlural: ts.RTPPlayerUnits,
		ObjLookup:     lookup.L.PlayerUnits,
		ObjLookupID:   lookup.L.PlayerUnitsID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.Typed],
		RetrieveQuery: config.Cfg.DB.GetPlayerUnitIDs,
		GetSingleFunc: getPlayerUnit,
		RetrieveFunc:  retrievePlayerUnits,
	}

	apilookups.E.Characters = handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named]{
		Endpoint:      ts.EpCharacters,
		ResTypeSingle: ts.RTSCharacters,
		ResTypePlural: ts.RTPCharacters,
		ObjLookup:     lookup.L.Characters,
		ObjLookupID:   lookup.L.CharactersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Character, apitypes.Character, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetCharacterIDs,
		GetSingleFunc: getCharacter,
		RetrieveFunc:  retrieveCharacters,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterDefaultAbilityIDs,
				CreateSubFn: CreatePlayerAbilitySimple,
			},
			ts.SnStdSgAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterSgAbilityIDs,
				CreateSubFn: CreatePlayerAbilitySimple,
			},
			ts.SnExpSgAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterEgAbilityIDs,
				CreateSubFn: CreatePlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterOverdriveAbilityIDs,
				CreateSubFn: CreateOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterOverdriveIDs,
				CreateSubFn: CreateOverdriveSimple,
			},
		},
	}

	apilookups.E.Aeons = handlerinput.Main[seedtypes.Aeon, apitypes.Aeon, apiresource.Named]{
		Endpoint:      ts.EpAeons,
		ResTypeSingle: ts.RTSAeons,
		ResTypePlural: ts.RTPAeons,
		ObjLookup:     lookup.L.Aeons,
		ObjLookupID:   lookup.L.AeonsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Aeon, apitypes.Aeon, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetAeonIDs,
		GetSingleFunc: getAeon,
		RetrieveFunc:  retrieveAeons,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetAeonDefaultAbilityIDs,
				CreateSubFn: CreatePlayerAbilitySimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetAeonOverdriveAbilityIDs,
				CreateSubFn: CreateOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     config.Cfg.DB.GetAeonOverdriveIDs,
				CreateSubFn: CreateOverdriveSimple,
			},
			ts.SnStats: {
				CreateSubFn: CreateAeonStatSimple,
			},
		},
	}

	apilookups.E.CharacterClasses = handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named]{
		Endpoint:      ts.EpCharacterClasses,
		ResTypeSingle: ts.RTSCharacterClasses,
		ResTypePlural: ts.RTPCharacterClasses,
		ObjLookup:     lookup.L.CharClasses,
		ObjLookupID:   lookup.L.CharClassesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetCharacterClassesIDs,
		GetSingleFunc: getCharacterClass,
		RetrieveFunc:  retrieveCharacterClasses,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnDefaultAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterClassDefaultAbilityIDs,
				CreateSubFn: CreateAbilitySimple,
			},
			ts.SnLearnableAbilities: {
				DBQuery:     config.Cfg.DB.GetCharacterClassLearnableAbilityIDs,
				CreateSubFn: CreateAbilitySimple,
			},
			ts.SnDefaultOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs,
				CreateSubFn: CreateOverdriveSimple,
			},
			ts.SnLearnableOverdrives: {
				DBQuery:     config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs,
				CreateSubFn: CreateOverdriveSimple,
			},
		},
	}

	apilookups.E.Monsters = handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named]{
		Endpoint:         ts.EpMonsters,
		ResTypeSingle:    ts.RTSMonsters,
		ResTypePlural:    ts.RTPMonsters,
		ObjLookup:        lookup.L.Monsters,
		ObjLookupID:      lookup.L.MonstersID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.Monster, apitypes.Monster, apiresource.Named],
		GetMultipleQuery: config.Cfg.DB.GetMonsterIDsByName,
		RetrieveQuery:    config.Cfg.DB.GetMonsterIDs,
		GetSingleFunc:    getMonster,
		RetrieveFunc:     retrieveMonsters,
		AvlFunc:          filterAvlMonsters,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
			ts.SnAbilities: {
				DBQuery:     config.Cfg.DB.GetMonsterAbilityIDs,
				CreateSubFn: CreateAbilitySimple,
			},
			ts.SnAreas: {
				DBQuery:     config.Cfg.DB.GetMonsterAreaIDs,
				CreateSubFn: CreateAreaSimple,
				RelationsFn: getAreaSectionRelations,
			},
			ts.SnMonsterFormations: {
				DBQuery:     config.Cfg.DB.GetMonsterMonsterFormationIDs,
				CreateSubFn: CreateMonsterFormationSimple,
			},
		},
	}

	apilookups.E.Abilities = handlerinput.Main[seedtypes.Ability, apitypes.Ability, apiresource.Typed]{
		Endpoint:      ts.EpAbilities,
		ResTypeSingle: ts.RTSAbilities,
		ResTypePlural: ts.RTPAbilities,
		ObjLookup:     lookup.L.Abilities,
		ObjLookupID:   lookup.L.AbilitiesID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.Ability, apitypes.Ability, apiresource.Typed],
		RetrieveQuery: config.Cfg.DB.GetAbilityIDs,
		GetSingleFunc: getAbility,
		RetrieveFunc:  retrieveAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateAbilitySimple,
			},
		},
	}

	apilookups.E.PlayerAbilities = handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named]{
		Endpoint:         ts.EpPlayerAbilities,
		ResTypeSingle:    ts.RTSPlayerAbilities,
		ResTypePlural:    ts.RTPPlayerAbilities,
		ObjLookup:        lookup.L.PlayerAbilities,
		ObjLookupID:      lookup.L.PlayerAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypePlayerAbility),
		RetrieveQuery:    config.Cfg.DB.GetPlayerAbilityIDs,
		GetSingleFunc:    getPlayerAbility,
		RetrieveFunc:     retrievePlayerAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreatePlayerAbilitySimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetPlayerAbilityMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	apilookups.E.OverdriveAbilities = handlerinput.Main[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named]{
		Endpoint:         ts.EpOverdriveAbilities,
		ResTypeSingle:    ts.RTSOverdriveAbilities,
		ResTypePlural:    ts.RTPOverdriveAbilities,
		ObjLookup:        lookup.L.OverdriveAbilities,
		ObjLookupID:      lookup.L.OverdriveAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.OverdriveAbility, apitypes.OverdriveAbility, apiresource.Named],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeOverdriveAbility),
		RetrieveQuery:    config.Cfg.DB.GetOverdriveAbilityIDs,
		GetSingleFunc:    getOverdriveAbility,
		RetrieveFunc:     retrieveOverdriveAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateOverdriveAbilitySimple,
			},
		},
	}

	apilookups.E.ItemAbilities = handlerinput.Main[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.Named]{
		Endpoint:      ts.EpItemAbilities,
		ResTypeSingle: ts.RTSItemAbilities,
		ResTypePlural: ts.RTPItemAbilities,
		ObjLookup:     lookup.L.ItemAbilities,
		ObjLookupID:   lookup.L.ItemAbilitiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetItemAbilityIDs,
		GetSingleFunc: getItemAbility,
		RetrieveFunc:  retrieveItemAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateItemAbilitySimple,
			},
		},
	}

	apilookups.E.TriggerCommands = handlerinput.Main[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named]{
		Endpoint:         ts.EpTriggerCommands,
		ResTypeSingle:    ts.RTSTriggerCommands,
		ResTypePlural:    ts.RTPTriggerCommands,
		ObjLookup:        lookup.L.TriggerCommands,
		ObjLookupID:      lookup.L.TriggerCommandsID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.TriggerCommand, apitypes.TriggerCommand, apiresource.Named],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeTriggerCommand),
		RetrieveQuery:    config.Cfg.DB.GetTriggerCommandIDs,
		GetSingleFunc:    getTriggerCommand,
		RetrieveFunc:     retrieveTriggerCommands,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateTriggerCommandSimple,
			},
		},
	}

	apilookups.E.MiscAbilities = handlerinput.Main[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named]{
		Endpoint:         ts.EpMiscAbilities,
		ResTypeSingle:    ts.RTSMiscAbilities,
		ResTypePlural:    ts.RTPMiscAbilities,
		ObjLookup:        lookup.L.MiscAbilities,
		ObjLookupID:      lookup.L.MiscAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.MiscAbility, apitypes.MiscAbility, apiresource.Named],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeMiscAbility),
		RetrieveQuery:    config.Cfg.DB.GetMiscAbilityIDs,
		GetSingleFunc:    getMiscAbility,
		RetrieveFunc:     retrieveMiscAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateMiscAbilitySimple,
			},
		},
	}

	apilookups.E.EnemyAbilities = handlerinput.Main[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named]{
		Endpoint:         ts.EpEnemyAbilities,
		ResTypeSingle:    ts.RTSEnemyAbilities,
		ResTypePlural:    ts.RTPEnemyAbilities,
		ObjLookup:        lookup.L.EnemyAbilities,
		ObjLookupID:      lookup.L.EnemyAbilitiesID,
		IDToResFunc:      handlerinput.IDToNamedAPIResource[seedtypes.EnemyAbility, apitypes.EnemyAbility, apiresource.Named],
		GetMultipleQuery: getTypedAbilityIDsByName(database.AbilityTypeEnemyAbility),
		RetrieveQuery:    config.Cfg.DB.GetEnemyAbilityIDs,
		GetSingleFunc:    getEnemyAbility,
		RetrieveFunc:     retrieveEnemyAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateEnemyAbilitySimple,
			},
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetEnemyAbilityMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	apilookups.E.AeonCommands = handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named]{
		Endpoint:      ts.EpAeonCommands,
		ResTypeSingle: ts.RTSAeonCommands,
		ResTypePlural: ts.RTPAeonCommands,
		ObjLookup:     lookup.L.AeonCommands,
		ObjLookupID:   lookup.L.AeonCommandsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetAeonCommandIDs,
		GetSingleFunc: getAeonCommand,
		RetrieveFunc:  retrieveAeonCommands,
	}

	apilookups.E.OverdriveCommands = handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named]{
		Endpoint:      ts.EpOverdriveCommands,
		ResTypeSingle: ts.RTSOverdriveCommands,
		ResTypePlural: ts.RTPOverdriveCommands,
		ObjLookup:     lookup.L.OverdriveCommands,
		ObjLookupID:   lookup.L.OverdriveCommandsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetOverdriveCommandIDs,
		GetSingleFunc: getOverdriveCommand,
		RetrieveFunc:  retrieveOverdriveCommands,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnOverdriveAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveAbilityIDs),
				CreateSubFn: CreateOverdriveAbilitySimple,
			},
			ts.SnOverdrives: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs),
				CreateSubFn: CreateOverdriveSimple,
			},
		},
	}

	apilookups.E.Overdrives = handlerinput.Main[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named]{
		Endpoint:      ts.EpOverdrives,
		ResTypeSingle: ts.RTSOverdrives,
		ResTypePlural: ts.RTPOverdrives,
		ObjLookup:     lookup.L.Overdrives,
		ObjLookupID:   lookup.L.OverdrivesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Overdrive, apitypes.Overdrive, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetOverdriveIDs,
		GetSingleFunc: getOverdrive,
		RetrieveFunc:  retrieveOverdrives,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateOverdriveSimple,
			},
			ts.SnOverdriveAbilities: {
				DBQuery:     config.Cfg.DB.GetOverdriveOverdriveAbilityIDs,
				CreateSubFn: CreateOverdriveAbilitySimple,
			},
		},
	}

	apilookups.E.RonsoRages = handlerinput.Main[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named]{
		Endpoint:      ts.EpRonsoRages,
		ResTypeSingle: ts.RTSRonsoRages,
		ResTypePlural: ts.RTPRonsoRages,
		ObjLookup:     lookup.L.RonsoRages,
		ObjLookupID:   lookup.L.RonsoRagesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetRonsoRageIDs,
		GetSingleFunc: getRonsoRage,
		RetrieveFunc:  retrieveRonsoRages,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnMonsters: {
				DBQuery:     config.Cfg.DB.GetRonsoRageMonsterIDs,
				CreateSubFn: CreateMonsterSimple,
				RelationsFn: getMonsterSectionRelations,
			},
		},
	}

	apilookups.E.Submenus = handlerinput.Main[seedtypes.Submenu, apitypes.Submenu, apiresource.Named]{
		Endpoint:      ts.EpSubmenus,
		ResTypeSingle: ts.RTSSubmenus,
		ResTypePlural: ts.RTPSubmenus,
		ObjLookup:     lookup.L.Submenus,
		ObjLookupID:   lookup.L.SubmenusID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Submenu, apitypes.Submenu, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetSubmenuIDs,
		GetSingleFunc: getSubmenu,
		RetrieveFunc:  retrieveSubmenus,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetSubmenuAbilityIDs),
				CreateSubFn: CreateAbilitySimple,
			},
		},
	}

	apilookups.E.Topmenus = handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named]{
		Endpoint:      ts.EpTopmenus,
		ResTypeSingle: ts.RTSTopmenus,
		ResTypePlural: ts.RTPTopmenus,
		ObjLookup:     lookup.L.Topmenus,
		ObjLookupID:   lookup.L.TopmenusID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetTopmenuIDs,
		GetSingleFunc: getTopmenu,
		RetrieveFunc:  retrieveTopmenus,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAbilities: {
				DBQuery:     handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs),
				CreateSubFn: CreateAbilitySimple,
			},
		},
	}

	apilookups.E.AllItems = handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed]{
		Endpoint:      ts.EpAllItems,
		ResTypeSingle: ts.RTSAllItems,
		ResTypePlural: ts.RTPAllItems,
		ObjLookup:     lookup.L.MasterItems,
		ObjLookupID:   lookup.L.MasterItemsID,
		IDToResFunc:   handlerinput.IDToTypedAPIResource[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed],
		RetrieveQuery: config.Cfg.DB.GetMasterItemIDs,
		GetSingleFunc: getMasterItem,
		RetrieveFunc:  retrieveMasterItems,
		AvlFunc:       filterAvlMasterItems,
	}

	apilookups.E.Items = handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named]{
		Endpoint:      ts.EpItems,
		ResTypeSingle: ts.RTSItems,
		ResTypePlural: ts.RTPItems,
		ObjLookup:     lookup.L.Items,
		ObjLookupID:   lookup.L.ItemsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Item, apitypes.Item, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetItemIDs,
		GetSingleFunc: getItem,
		RetrieveFunc:  retrieveItems,
		AvlFunc:       filterAvlItems,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnMixes: {
				DBQuery:     config.Cfg.DB.GetItemMixIDs,
				CreateSubFn: CreateMixSimple,
			},
		},
	}

	apilookups.E.Spheres = handlerinput.Main[seedtypes.Sphere, apitypes.Sphere, apiresource.Named]{
		Endpoint:      ts.EpSpheres,
		ResTypeSingle: ts.RTSSpheres,
		ResTypePlural: ts.RTPSpheres,
		ObjLookup:     lookup.L.Spheres,
		ObjLookupID:   lookup.L.SpheresID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Sphere, apitypes.Sphere, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetSphereIDs,
		GetSingleFunc: getSphere,
		RetrieveFunc:  retrieveSpheres,
		AvlFunc:       filterAvlSpheres,
	}

	apilookups.E.KeyItems = handlerinput.Main[seedtypes.KeyItem, apitypes.KeyItem, apiresource.Named]{
		Endpoint:      ts.EpKeyItems,
		ResTypeSingle: ts.RTSKeyItems,
		ResTypePlural: ts.RTPKeyItems,
		ObjLookup:     lookup.L.KeyItems,
		ObjLookupID:   lookup.L.KeyItemsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.KeyItem, apitypes.KeyItem, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetKeyItemIDs,
		GetSingleFunc: getKeyItem,
		RetrieveFunc:  retrieveKeyItems,
		AvlFunc:       filterAvlKeyItems,
	}

	apilookups.E.Primers = handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named]{
		Endpoint:      ts.EpPrimers,
		ResTypeSingle: ts.RTSPrimers,
		ResTypePlural: ts.RTPPrimers,
		ObjLookup:     lookup.L.Primers,
		ObjLookupID:   lookup.L.PrimersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Primer, apitypes.Primer, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetPrimerIDs,
		GetSingleFunc: getPrimer,
		RetrieveFunc:  retrievePrimers,
		AvlFunc:       filterAvlPrimers,
	}

	apilookups.E.Mixes = handlerinput.Main[seedtypes.Mix, apitypes.Mix, apiresource.Named]{
		Endpoint:      ts.EpMixes,
		ResTypeSingle: ts.RTSMixes,
		ResTypePlural: ts.RTPMixes,
		ObjLookup:     lookup.L.Mixes,
		ObjLookupID:   lookup.L.MixesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Mix, apitypes.Mix, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetMixIDs,
		GetSingleFunc: getMix,
		RetrieveFunc:  retrieveMixes,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateMixSimple,
			},
		},
	}

	apilookups.E.AutoAbilities = handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named]{
		Endpoint:      ts.EpAutoAbilities,
		ResTypeSingle: ts.RTSAutoAbilities,
		ResTypePlural: ts.RTPAutoAbilities,
		ObjLookup:     lookup.L.AutoAbilities,
		ObjLookupID:   lookup.L.AutoAbilitiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetAutoAbilityIDs,
		GetSingleFunc: getAutoAbility,
		RetrieveFunc:  retrieveAutoAbilities,
		AvlFunc:       filterAvlAutoAbilities,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnSimple: {
				CreateSubFn: CreateAutoAbilitySimple,
			},
		},
	}

	apilookups.E.EquipmentTables = handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed]{
		Endpoint:      ts.EpEquipmentTables,
		ResTypeSingle: ts.RTSEquipmentTables,
		ResTypePlural: ts.RTPEquipmentTables,
		ObjLookup:     lookup.L.EquipmentTables,
		ObjLookupID:   lookup.L.EquipmentTablesID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed],
		RetrieveQuery: config.Cfg.DB.GetEquipmentTableIDs,
		GetSingleFunc: getEquipmentTable,
		RetrieveFunc:  retrieveEquipmentTables,
	}

	apilookups.E.Equipment = handlerinput.Main[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named]{
		Endpoint:      ts.EpEquipment,
		ResTypeSingle: ts.RTSEquipment,
		ResTypePlural: ts.RTPEquipment,
		ObjLookup:     lookup.L.EquipmentNames,
		ObjLookupID:   lookup.L.EquipmentNamesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.EquipmentName, apitypes.EquipmentName, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetEquipmentIDs,
		GetSingleFunc: getEquipment,
		RetrieveFunc:  retrieveEquipment,
	}

	apilookups.E.CelestialWeapons = handlerinput.Main[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.Named]{
		Endpoint:      ts.EpCelestialWeapons,
		ResTypeSingle: ts.RTSCelestialWeapons,
		ResTypePlural: ts.RTPCelestialWeapons,
		ObjLookup:     lookup.L.CelestialWeapons,
		ObjLookupID:   lookup.L.CelestialWeaponsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetCelestialWeaponIDs,
		GetSingleFunc: getCelestialWeapon,
		RetrieveFunc:  retrieveCelestialWeapons,
		Subsections: map[ts.SectionName]handlerinput.Subsection{
			ts.SnAutoAbilities: {
				DBQuery:     GetCelestialWeaponAutoAbilityIDs(),
				CreateSubFn: CreateAutoAbilitySimple,
			},
		},
	}

	apilookups.E.Stats = handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named]{
		Endpoint:      ts.EpStats,
		ResTypeSingle: ts.RTSStats,
		ResTypePlural: ts.RTPStats,
		ObjLookup:     lookup.L.Stats,
		ObjLookupID:   lookup.L.StatsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Stat, apitypes.Stat, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetStatIDs,
		GetSingleFunc: getStat,
		RetrieveFunc:  retrieveStats,
	}

	apilookups.E.Properties = handlerinput.Main[seedtypes.Property, apitypes.Property, apiresource.Named]{
		Endpoint:      ts.EpProperties,
		ResTypeSingle: ts.RTSProperties,
		ResTypePlural: ts.RTPProperties,
		ObjLookup:     lookup.L.Properties,
		ObjLookupID:   lookup.L.PropertiesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Property, apitypes.Property, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetPropertyIDs,
		GetSingleFunc: getProperty,
		RetrieveFunc:  retrieveProperties,
	}

	apilookups.E.OverdriveModes = handlerinput.Main[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named]{
		Endpoint:      ts.EpOverdriveModes,
		ResTypeSingle: ts.RTSOverdriveModes,
		ResTypePlural: ts.RTPOverdriveModes,
		ObjLookup:     lookup.L.OverdriveModes,
		ObjLookupID:   lookup.L.OverdriveModesID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetOverdriveModeIDs,
		GetSingleFunc: getOverdriveMode,
		RetrieveFunc:  retrieveOverdriveModes,
	}

	apilookups.E.Elements = handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named]{
		Endpoint:      ts.EpElements,
		ResTypeSingle: ts.RTSElements,
		ResTypePlural: ts.RTPElements,
		ObjLookup:     lookup.L.Elements,
		ObjLookupID:   lookup.L.ElementsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Element, apitypes.Element, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetElementIDs,
		GetSingleFunc: getElement,
		RetrieveFunc:  retrieveElements,
	}

	apilookups.E.StatusConditions = handlerinput.Main[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named]{
		Endpoint:      ts.EpStatusConditions,
		ResTypeSingle: ts.RTSStatusConditions,
		ResTypePlural: ts.RTPStatusConditions,
		ObjLookup:     lookup.L.StatusConditions,
		ObjLookupID:   lookup.L.StatusConditionsID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetStatusConditionIDs,
		GetSingleFunc: getStatusCondition,
		RetrieveFunc:  retrieveStatusConditions,
	}

	apilookups.E.Modifiers = handlerinput.Main[seedtypes.Modifier, apitypes.Modifier, apiresource.Named]{
		Endpoint:      ts.EpModifiers,
		ResTypeSingle: ts.RTSModifiers,
		ResTypePlural: ts.RTPModifiers,
		ObjLookup:     lookup.L.Modifiers,
		ObjLookupID:   lookup.L.ModifiersID,
		IDToResFunc:   handlerinput.IDToNamedAPIResource[seedtypes.Modifier, apitypes.Modifier, apiresource.Named],
		RetrieveQuery: config.Cfg.DB.GetModifierIDs,
		GetSingleFunc: getModifier,
		RetrieveFunc:  retrieveModifiers,
	}

	apilookups.E.AgilityTiers = handlerinput.Main[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.Unnamed]{
		Endpoint:      ts.EpAgilityTiers,
		ResTypeSingle: ts.RTSAgilityTiers,
		ResTypePlural: ts.RTPAgilityTiers,
		ObjLookup:     nil,
		ObjLookupID:   lookup.L.AgilityTiersID,
		IDToResFunc:   handlerinput.IDToUnnamedAPIResource[seedtypes.AgilityTier, apitypes.AgilityTier, apiresource.Unnamed],
		RetrieveQuery: config.Cfg.DB.GetAgilityTierIDs,
		GetSingleFunc: getAgilityTier,
		RetrieveFunc:  retrieveAgilityTiers,
	}

	apilookups.E.Enums = handlerinput.Enums{
		Endpoint: ts.EpEnums,
		Usage:    []string{"/api/enums", "/api/enums/{enum_name}"},
		Lookup:   enums.T.Lookup,
	}

	apilookups.E.Endpoints = handlerinput.Endpoints{
		Endpoint: ts.EpEndpoints,
		Usage:    []string{"/api/endpoints"},
	}

	apilookups.E.AlBhed = handlerinput.Service[servicetypes.AlBhedParams, servicetypes.AlBhedResponse]{
		Endpoint:  ts.EpAlBhed,
		VerifyFn:  serviceverify.VerifyAlBhedParams,
		ExecuteFn: handleAlBhed,
		DocsFn:    servicetypes.GetAlBhedParamsDoc,
		Samples:   servicetypes.AlBhedParamsSamples,
	}

	apilookups.E.TurnOrder = handlerinput.Service[servicetypes.TurnOrderParams, servicetypes.TurnOrderResponse]{
		Endpoint:  ts.EpTurnOrder,
		VerifyFn:  serviceverify.VerifyTurnOrderParams,
		ExecuteFn: handleTurnOrder,
		DocsFn:    servicetypes.GetTurnOrderParamsDoc,
		Samples:   servicetypes.TurnOrderParamsSamples,
	}

	apilookups.E.DropChance = handlerinput.Service[servicetypes.DropChanceParams, servicetypes.DropChanceResponse]{
		Endpoint:  ts.EpDropChance,
		VerifyFn:  serviceverify.VerifyDropChanceParams,
		ExecuteFn: handleDropChance,
		DocsFn:    servicetypes.GetDropChanceParamsDoc,
		Samples:   servicetypes.DropChanceParamsSamples,
	}

	apilookups.E.CritChance = handlerinput.Service[servicetypes.CritChanceParams, servicetypes.CritChanceResponse]{
		Endpoint:  ts.EpCritChance,
		VerifyFn:  serviceverify.VerifyCritChanceParams,
		ExecuteFn: handleCritChance,
		DocsFn:    servicetypes.GetCritChanceParamsDoc,
		Samples:   servicetypes.CritChanceParamsSamples,
	}

	apilookups.E.InitializeEndpointSlice()
}
