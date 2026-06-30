package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type endpoints struct {
	locations    handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]
	sublocations handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]
	areas        handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]

	monsterFormations handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]
	shops             handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]
	treasures         handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]
	quests            handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList]
	sidequests        handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList]
	subquests         handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList]
	arenaCreations    handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]
	blitzballPrizes   handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedApiResourceList]
	songs             handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]
	fmvs              handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]

	playerUnits      handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]
	characters       handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]
	aeons            handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]
	characterClasses handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]
	monsters         handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]

	abilities          handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]
	playerAbilities    handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]
	itemAbilities      handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]
	triggerCommands    handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]
	miscAbilities      handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList]
	enemyAbilities     handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]

	aeonCommands      handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]
	overdriveCommands handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]
	overdrives        handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]
	ronsoRages        handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]
	submenus          handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]
	topmenus          handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]

	allItems handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]
	items    handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]
	keyItems handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList]
	spheres  handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList]
	primers  handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList]
	mixes    handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList]

	autoAbilities    handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList]
	equipmentTables  handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList]
	equipment        handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList]
	celestialWeapons handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList]

	stats            handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList]
	properties       handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList]
	overdriveModes   handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]
	elements         handlerInput[seeding.Element, Element, NamedAPIResource, NamedApiResourceList]
	statusConditions handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList]
	modifiers        handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList]
	agilityTiers     handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.locations = handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epLocations,
		resTypeSingle: rtsLocations,
		resTypePlural: rtpLocations,
		objLookup:     cfg.l.Locations,
		objLookupID:   cfg.l.LocationsID,
		queryLookup:   cfg.q.locations,
		idToResFunc:   idToNamedAPIResource[seeding.Location, Location, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		avlFunc:       filterAvlLocations,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			snConnected: {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			snSublocations: {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			snAreas: {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			snMonsterFormations: {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			snShops: {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			snSongs: {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			snTreasures: {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.sublocations = handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epSublocations,
		resTypeSingle: rtsSublocations,
		resTypePlural: rtpSublocations,
		objLookup:     cfg.l.Sublocations,
		objLookupID:   cfg.l.SublocationsID,
		queryLookup:   cfg.q.sublocations,
		idToResFunc:   idToNamedAPIResource[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		avlFunc:       filterAvlSublocations,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			snConnected: {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			snAreas: {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			snMonsterFormations: {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			snShops: {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			snSongs: {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			snTreasures: {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.areas = handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]{
		endpoint:      epAreas,
		resTypeSingle: rtsAreas,
		resTypePlural: rtpAreas,
		objLookup:     cfg.l.Areas,
		objLookupID:   cfg.l.AreasID,
		queryLookup:   cfg.q.areas,
		idToResFunc:   idToAreaAPIResource,
		resToListFunc: newAreaAPIResourceList,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		avlFunc:       filterAvlAreas,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			snConnected: {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			snMonsterFormations: {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			snSongs: {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			snTreasures: {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      epMonsterFormations,
		resTypeSingle: rtsMonsterFormations,
		resTypePlural: rtpMonsterFormations,
		objLookup:     cfg.l.MonsterFormations,
		objLookupID:   cfg.l.MonsterFormationsID,
		queryLookup:   cfg.q.monsterFormations,
		idToResFunc:   idToUnnamedAPIResource[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		avlFunc:       filterAvlMonsterFormations,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createMonsterFormationSimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.shops = handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      epShops,
		resTypeSingle: rtsShops,
		resTypePlural: rtpShops,
		objLookup:     cfg.l.Shops,
		objLookupID:   cfg.l.ShopsID,
		queryLookup:   cfg.q.shops,
		idToResFunc:   idToUnnamedAPIResource[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		avlFunc:       filterAvlShops,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createShopSimple,
			},
		},
	}

	e.treasures = handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      epTreasures,
		resTypeSingle: rtsTreasures,
		resTypePlural: rtpTreasures,
		objLookup:     cfg.l.Treasures,
		objLookupID:   cfg.l.TreasuresID,
		queryLookup:   cfg.q.treasures,
		idToResFunc:   idToUnnamedAPIResource[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		avlFunc:       filterAvlTreasures,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	e.quests = handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      epQuests,
		resTypeSingle: rtsQuests,
		resTypePlural: rtpQuests,
		objLookup:     cfg.l.Quests,
		objLookupID:   cfg.l.QuestsID,
		queryLookup:   cfg.q.quests,
		idToResFunc:   idToQuestAPIResource[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetQuestIDs,
		getSingleFunc: cfg.getQuest,
		retrieveFunc:  cfg.retrieveQuests,
		avlFunc:       filterAvlQuests,
	}

	e.sidequests = handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      epSidequests,
		resTypeSingle: rtsSidequests,
		resTypePlural: rtpSidequests,
		objLookup:     cfg.l.Sidequests,
		objLookupID:   cfg.l.SidequestsID,
		queryLookup:   cfg.q.sidequests,
		idToResFunc:   idToQuestAPIResource[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		avlFunc:       filterAvlSidequests,
		subsections: map[SectionName]Subsection{
			snSubquests: {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	e.subquests = handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      epSubquests,
		resTypeSingle: rtsSubquests,
		resTypePlural: rtpSubquests,
		objLookup:     cfg.l.Subquests,
		objLookupID:   cfg.l.SubquestsID,
		queryLookup:   cfg.q.subquests,
		idToResFunc:   idToQuestAPIResource[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
		avlFunc:       filterAvlSubquests,
	}

	e.arenaCreations = handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epArenaCreations,
		resTypeSingle: rtsArenaCreations,
		resTypePlural: rtpArenaCreations,
		objLookup:     cfg.l.ArenaCreations,
		objLookupID:   cfg.l.ArenaCreationsID,
		queryLookup:   cfg.q.arenaCreations,
		idToResFunc:   idToNamedAPIResource[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetArenaCreationIDs,
		getSingleFunc: cfg.getArenaCreation,
		retrieveFunc:  cfg.retrieveArenaCreations,
	}

	e.blitzballPrizes = handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epBlitzballPrizes,
		resTypeSingle: rtsBlitzballPrizes,
		resTypePlural: rtpBlitzballPrizes,
		objLookup:     cfg.l.Positions,
		objLookupID:   cfg.l.PositionsID,
		queryLookup:   cfg.q.blitzballPrizes,
		idToResFunc:   idToNamedAPIResource[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetBlitzballPrizeIDs,
		getSingleFunc: cfg.getBlitzballPrize,
		retrieveFunc:  cfg.retrieveBlitzballPrizes,
	}

	e.songs = handlerInput[seeding.Song, Song, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epSongs,
		resTypeSingle: rtsSongs,
		resTypePlural: rtpSongs,
		objLookup:     cfg.l.Songs,
		objLookupID:   cfg.l.SongsID,
		queryLookup:   cfg.q.songs,
		idToResFunc:   idToNamedAPIResource[seeding.Song, Song, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSongIDs,
		getSingleFunc: cfg.getSong,
		retrieveFunc:  cfg.retrieveSongs,
	}

	e.fmvs = handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epFMVs,
		resTypeSingle: rtsFMVs,
		resTypePlural: rtpFMVs,
		objLookup:     cfg.l.FMVs,
		objLookupID:   cfg.l.FMVsID,
		queryLookup:   cfg.q.fmvs,
		idToResFunc:   idToNamedAPIResource[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetFmvIDs,
		getSingleFunc: cfg.getFMV,
		retrieveFunc:  cfg.retrieveFMVs,
	}

	e.playerUnits = handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]{
		endpoint:      epPlayerUnits,
		resTypeSingle: rtsPlayerUnits,
		resTypePlural: rtpPlayerUnits,
		objLookup:     cfg.l.PlayerUnits,
		objLookupID:   cfg.l.PlayerUnitsID,
		queryLookup:   cfg.q.playerUnits,
		idToResFunc:   idToTypedAPIResource[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetPlayerUnitIDs,
		getSingleFunc: cfg.getPlayerUnit,
		retrieveFunc:  cfg.retrievePlayerUnits,
	}

	e.characters = handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epCharacters,
		resTypeSingle: rtsCharacters,
		resTypePlural: rtpCharacters,
		objLookup:     cfg.l.Characters,
		objLookupID:   cfg.l.CharactersID,
		queryLookup:   cfg.q.characters,
		idToResFunc:   idToNamedAPIResource[seeding.Character, Character, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[SectionName]Subsection{
			snDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			snStdSgAbilities: {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			snExpSgAbilities: {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			snOverdriveAbilities: {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			snOverdrives: {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.aeons = handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epAeons,
		resTypeSingle: rtsAeons,
		resTypePlural: rtpAeons,
		objLookup:     cfg.l.Aeons,
		objLookupID:   cfg.l.AeonsID,
		queryLookup:   cfg.q.aeons,
		idToResFunc:   idToNamedAPIResource[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[SectionName]Subsection{
			snDefaultAbilities: {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			snOverdriveAbilities: {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
			snOverdrives: {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			snStats: {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	e.characterClasses = handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epCharacterClasses,
		resTypeSingle: rtsCharacterClasses,
		resTypePlural: rtpCharacterClasses,
		objLookup:     cfg.l.CharClasses,
		objLookupID:   cfg.l.CharClassesID,
		queryLookup:   cfg.q.characterClasses,
		idToResFunc:   idToNamedAPIResource[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[SectionName]Subsection{
			snDefaultAbilities: {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			snLearnableAbilities: {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			snDefaultOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			snLearnableOverdrives: {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epMonsters,
		resTypeSingle:    rtsMonsters,
		resTypePlural:    rtpMonsters,
		objLookup:        cfg.l.Monsters,
		objLookupID:      cfg.l.MonstersID,
		queryLookup:      cfg.q.monsters,
		idToResFunc:      idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		avlFunc:          filterAvlMonsters,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			snAbilities: {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
			},
			snAreas: {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			snMonsterFormations: {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	e.abilities = handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]{
		endpoint:      epAbilities,
		resTypeSingle: rtsAbilities,
		resTypePlural: rtpAbilities,
		objLookup:     cfg.l.Abilities,
		objLookupID:   cfg.l.AbilitiesID,
		queryLookup:   cfg.q.abilities,
		idToResFunc:   idToTypedAPIResource[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epPlayerAbilities,
		resTypeSingle:    rtsPlayerAbilities,
		resTypePlural:    rtpPlayerAbilities,
		objLookup:        cfg.l.PlayerAbilities,
		objLookupID:      cfg.l.PlayerAbilitiesID,
		queryLookup:      cfg.q.playerAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypePlayerAbility),
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createPlayerAbilitySimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epOverdriveAbilities,
		resTypeSingle:    rtsOverdriveAbilities,
		resTypePlural:    rtpOverdriveAbilities,
		objLookup:        cfg.l.OverdriveAbilities,
		objLookupID:      cfg.l.OverdriveAbilitiesID,
		queryLookup:      cfg.q.overdriveAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeOverdriveAbility),
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.itemAbilities = handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epItemAbilities,
		resTypeSingle: rtsItemAbilities,
		resTypePlural: rtpItemAbilities,
		objLookup:     cfg.l.ItemAbilities,
		objLookupID:   cfg.l.ItemAbilitiesID,
		queryLookup:   cfg.q.itemAbilities,
		idToResFunc:   idToNamedAPIResource[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epTriggerCommands,
		resTypeSingle:    rtsTriggerCommands,
		resTypePlural:    rtpTriggerCommands,
		objLookup:        cfg.l.TriggerCommands,
		objLookupID:      cfg.l.TriggerCommandsID,
		queryLookup:      cfg.q.triggerCommands,
		idToResFunc:      idToNamedAPIResource[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeTriggerCommand),
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	e.miscAbilities = handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epMiscAbilities,
		resTypeSingle:    rtsMiscAbilities,
		resTypePlural:    rtpMiscAbilities,
		objLookup:        cfg.l.MiscAbilities,
		objLookupID:      cfg.l.MiscAbilitiesID,
		queryLookup:      cfg.q.miscAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeMiscAbility),
		retrieveQuery:    cfg.db.GetMiscAbilityIDs,
		getSingleFunc:    cfg.getMiscAbility,
		retrieveFunc:     cfg.retrieveMiscAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createMiscAbilitySimple,
			},
		},
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         epEnemyAbilities,
		resTypeSingle:    rtsEnemyAbilities,
		resTypePlural:    rtpEnemyAbilities,
		objLookup:        cfg.l.EnemyAbilities,
		objLookupID:      cfg.l.EnemyAbilitiesID,
		queryLookup:      cfg.q.enemyAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: getTypedAbilityIDsByName(cfg, database.AbilityTypeEnemyAbility),
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createEnemyAbilitySimple,
			},
			snMonsters: {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.aeonCommands = handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epAeonCommands,
		resTypeSingle: rtsAeonCommands,
		resTypePlural: rtpAeonCommands,
		objLookup:     cfg.l.AeonCommands,
		objLookupID:   cfg.l.AeonCommandsID,
		queryLookup:   cfg.q.aeonCommands,
		idToResFunc:   idToNamedAPIResource[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonCommandIDs,
		getSingleFunc: cfg.getAeonCommand,
		retrieveFunc:  cfg.retrieveAeonCommands,
	}

	e.overdriveCommands = handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epOverdriveCommands,
		resTypeSingle: rtsOverdriveCommands,
		resTypePlural: rtpOverdriveCommands,
		objLookup:     cfg.l.OverdriveCommands,
		objLookupID:   cfg.l.OverdriveCommandsID,
		queryLookup:   cfg.q.overdriveCommands,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[SectionName]Subsection{
			snOverdriveAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveAbilityIDs),
				createSubFn: createOverdriveAbilitySimple,
			},
			snOverdrives: {
				dbQuery:     ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs),
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.overdrives = handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epOverdrives,
		resTypeSingle: rtsOverdrives,
		resTypePlural: rtpOverdrives,
		objLookup:     cfg.l.Overdrives,
		objLookupID:   cfg.l.OverdrivesID,
		queryLookup:   cfg.q.overdrives,
		idToResFunc:   idToNamedAPIResource[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createOverdriveSimple,
			},
			snOverdriveAbilities: {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
			},
		},
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epRonsoRages,
		resTypeSingle: rtsRonsoRages,
		resTypePlural: rtpRonsoRages,
		objLookup:     cfg.l.RonsoRages,
		objLookupID:   cfg.l.RonsoRagesID,
		queryLookup:   cfg.q.ronsoRages,
		idToResFunc:   idToNamedAPIResource[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[SectionName]Subsection{
			snMonsters: {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.submenus = handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epSubmenus,
		resTypeSingle: rtsSubmenus,
		resTypePlural: rtpSubmenus,
		objLookup:     cfg.l.Submenus,
		objLookupID:   cfg.l.SubmenusID,
		queryLookup:   cfg.q.submenus,
		idToResFunc:   idToNamedAPIResource[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[SectionName]Subsection{
			snAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.topmenus = handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epTopmenus,
		resTypeSingle: rtsTopmenus,
		resTypePlural: rtpTopmenus,
		objLookup:     cfg.l.Topmenus,
		objLookupID:   cfg.l.TopmenusID,
		queryLookup:   cfg.q.topmenus,
		idToResFunc:   idToNamedAPIResource[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[SectionName]Subsection{
			snAbilities: {
				dbQuery:     ToIntManyNull(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
			},
		},
	}

	e.allItems = handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]{
		endpoint:      epAllItems,
		resTypeSingle: rtsAllItems,
		resTypePlural: rtpAllItems,
		objLookup:     cfg.l.MasterItems,
		objLookupID:   cfg.l.MasterItemsID,
		queryLookup:   cfg.q.allItems,
		idToResFunc:   idToTypedAPIResource[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetMasterItemIDs,
		getSingleFunc: cfg.getMasterItem,
		retrieveFunc:  cfg.retrieveMasterItems,
		avlFunc:       filterAvlMasterItems,
	}

	e.items = handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epItems,
		resTypeSingle: rtsItems,
		resTypePlural: rtpItems,
		objLookup:     cfg.l.Items,
		objLookupID:   cfg.l.ItemsID,
		queryLookup:   cfg.q.items,
		idToResFunc:   idToNamedAPIResource[seeding.Item, Item, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemIDs,
		getSingleFunc: cfg.getItem,
		retrieveFunc:  cfg.retrieveItems,
		avlFunc:       filterAvlItems,
		subsections: map[SectionName]Subsection{
			snMixes: {
				dbQuery:     cfg.db.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	e.spheres = handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epSpheres,
		resTypeSingle: rtsSpheres,
		resTypePlural: rtpSpheres,
		objLookup:     cfg.l.Spheres,
		objLookupID:   cfg.l.SpheresID,
		queryLookup:   cfg.q.spheres,
		idToResFunc:   idToNamedAPIResource[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSphereIDs,
		getSingleFunc: cfg.getSphere,
		retrieveFunc:  cfg.retrieveSpheres,
		avlFunc:       filterAvlSpheres,
	}

	e.keyItems = handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epKeyItems,
		resTypeSingle: rtsKeyItems,
		resTypePlural: rtpKeyItems,
		objLookup:     cfg.l.KeyItems,
		objLookupID:   cfg.l.KeyItemsID,
		queryLookup:   cfg.q.keyItems,
		idToResFunc:   idToNamedAPIResource[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetKeyItemIDs,
		getSingleFunc: cfg.getKeyItem,
		retrieveFunc:  cfg.retrieveKeyItems,
		avlFunc:       filterAvlKeyItems,
	}

	e.primers = handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epPrimers,
		resTypeSingle: rtsPrimers,
		resTypePlural: rtpPrimers,
		objLookup:     cfg.l.Primers,
		objLookupID:   cfg.l.PrimersID,
		queryLookup:   cfg.q.primers,
		idToResFunc:   idToNamedAPIResource[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetPrimerIDs,
		getSingleFunc: cfg.getPrimer,
		retrieveFunc:  cfg.retrievePrimers,
		avlFunc:       filterAvlPrimers,
	}

	e.mixes = handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epMixes,
		resTypeSingle: rtsMixes,
		resTypePlural: rtpMixes,
		objLookup:     cfg.l.Mixes,
		objLookupID:   cfg.l.MixesID,
		queryLookup:   cfg.q.mixes,
		idToResFunc:   idToNamedAPIResource[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetMixIDs,
		getSingleFunc: cfg.getMix,
		retrieveFunc:  cfg.retrieveMixes,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createMixSimple,
			},
		},
	}

	e.autoAbilities = handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epAutoAbilities,
		resTypeSingle: rtsAutoAbilities,
		resTypePlural: rtpAutoAbilities,
		objLookup:     cfg.l.AutoAbilities,
		objLookupID:   cfg.l.AutoAbilitiesID,
		queryLookup:   cfg.q.autoAbilities,
		idToResFunc:   idToNamedAPIResource[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAutoAbilityIDs,
		getSingleFunc: cfg.getAutoAbility,
		retrieveFunc:  cfg.retrieveAutoAbilities,
		avlFunc:       filterAvlAutoAbilities,
		subsections: map[SectionName]Subsection{
			snSimple: {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.equipmentTables = handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      epEquipmentTables,
		resTypeSingle: rtsEquipmentTables,
		resTypePlural: rtpEquipmentTables,
		objLookup:     cfg.l.EquipmentTables,
		objLookupID:   cfg.l.EquipmentTablesID,
		queryLookup:   cfg.q.equipmentTables,
		idToResFunc:   idToUnnamedAPIResource[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetEquipmentTableIDs,
		getSingleFunc: cfg.getEquipmentTable,
		retrieveFunc:  cfg.retrieveEquipmentTables,
	}

	e.equipment = handlerInput[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epEquipment,
		resTypeSingle: rtsEquipment,
		resTypePlural: rtpEquipment,
		objLookup:     cfg.l.EquipmentNames,
		objLookupID:   cfg.l.EquipmentNamesID,
		queryLookup:   cfg.q.equipment,
		idToResFunc:   idToNamedAPIResource[seeding.EquipmentName, EquipmentName, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetEquipmentIDs,
		getSingleFunc: cfg.getEquipment,
		retrieveFunc:  cfg.retrieveEquipment,
	}

	e.celestialWeapons = handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epCelestialWeapons,
		resTypeSingle: rtsCelestialWeapons,
		resTypePlural: rtpCelestialWeapons,
		objLookup:     cfg.l.CelestialWeapons,
		objLookupID:   cfg.l.CelestialWeaponsID,
		queryLookup:   cfg.q.celestialWeapons,
		idToResFunc:   idToNamedAPIResource[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCelestialWeaponIDs,
		getSingleFunc: cfg.getCelestialWeapon,
		retrieveFunc:  cfg.retrieveCelestialWeapons,
		subsections: map[SectionName]Subsection{
			snAutoAbilities: {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(cfg),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.stats = handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epStats,
		resTypeSingle: rtsStats,
		resTypePlural: rtpStats,
		objLookup:     cfg.l.Stats,
		objLookupID:   cfg.l.StatsID,
		queryLookup:   cfg.q.stats,
		idToResFunc:   idToNamedAPIResource[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetStatIDs,
		getSingleFunc: cfg.getStat,
		retrieveFunc:  cfg.retrieveStats,
	}

	e.properties = handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epProperties,
		resTypeSingle: rtsProperties,
		resTypePlural: rtpProperties,
		objLookup:     cfg.l.Properties,
		objLookupID:   cfg.l.PropertiesID,
		queryLookup:   cfg.q.properties,
		idToResFunc:   idToNamedAPIResource[seeding.Property, Property, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetPropertyIDs,
		getSingleFunc: cfg.getProperty,
		retrieveFunc:  cfg.retrieveProperties,
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epOverdriveModes,
		resTypeSingle: rtsOverdriveModes,
		resTypePlural: rtpOverdriveModes,
		objLookup:     cfg.l.OverdriveModes,
		objLookupID:   cfg.l.OverdriveModesID,
		queryLookup:   cfg.q.overdriveModes,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveModeIDs,
		getSingleFunc: cfg.getOverdriveMode,
		retrieveFunc:  cfg.retrieveOverdriveModes,
	}

	e.elements = handlerInput[seeding.Element, Element, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epElements,
		resTypeSingle: rtsElements,
		resTypePlural: rtpElements,
		objLookup:     cfg.l.Elements,
		objLookupID:   cfg.l.ElementsID,
		queryLookup:   cfg.q.elements,
		idToResFunc:   idToNamedAPIResource[seeding.Element, Element, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetElementIDs,
		getSingleFunc: cfg.getElement,
		retrieveFunc:  cfg.retrieveElements,
	}

	e.statusConditions = handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epStatusConditions,
		resTypeSingle: rtsStatusConditions,
		resTypePlural: rtpStatusConditions,
		objLookup:     cfg.l.StatusConditions,
		objLookupID:   cfg.l.StatusConditionsID,
		queryLookup:   cfg.q.statusConditions,
		idToResFunc:   idToNamedAPIResource[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetStatusConditionIDs,
		getSingleFunc: cfg.getStatusCondition,
		retrieveFunc:  cfg.retrieveStatusConditions,
	}

	e.modifiers = handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList]{
		endpoint:      epModifiers,
		resTypeSingle: rtsModifiers,
		resTypePlural: rtpModifiers,
		objLookup:     cfg.l.Modifiers,
		objLookupID:   cfg.l.ModifiersID,
		queryLookup:   cfg.q.modifiers,
		idToResFunc:   idToNamedAPIResource[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetModifierIDs,
		getSingleFunc: cfg.getModifier,
		retrieveFunc:  cfg.retrieveModifiers,
	}

	e.agilityTiers = handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      epAgilityTiers,
		resTypeSingle: rtsAgilityTiers,
		resTypePlural: rtpAgilityTiers,
		objLookup:     nil,
		objLookupID:   cfg.l.AgilityTiersID,
		queryLookup:   cfg.q.agilityTiers,
		idToResFunc:   idToUnnamedAPIResource[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetAgilityTierIDs,
		getSingleFunc: cfg.getAgilityTier,
		retrieveFunc:  cfg.retrieveAgilityTiers,
	}

	cfg.e = &e
}
