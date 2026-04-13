package api

import (
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

	abilities            handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]
	playerAbilities      handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]
	overdriveAbilities   handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]
	itemAbilities        handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]
	triggerCommands      handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]
	unspecifiedAbilities handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]
	enemyAbilities       handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]

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
	agilityTiers	 handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]

	abilityType handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	unitType    handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	itemType    handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	questType   handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]

	attackType               handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	autoAbilityCategory      handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	availabilityType         handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	damageFormula            handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	damageType               handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	elementalAffinity        handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	itemCategory             handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	keyItemCategory          handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	lootType                 handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	mixCategory              handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterCategory          handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterFormationCategory handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	monsterSpecies           handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	playerAbilityCategory    handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
	shopCategory             handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]
}

func (cfg *Config) EndpointsInit() {
	e := endpoints{}

	e.locations = handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "locations",
		resourceType:  "location",
		objLookup:     cfg.l.Locations,
		objLookupID:   cfg.l.LocationsID,
		queryLookup:   cfg.q.locations,
		idToResFunc:   idToNamedAPIResource[seeding.Location, Location, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetLocationIDs,
		getSingleFunc: cfg.getLocation,
		retrieveFunc:  cfg.retrieveLocations,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedLocationIDs,
				createSubFn: createLocationSimple,
				relationsFn: getLocationSectionRelations,
			},
			"sublocations": {
				dbQuery:     cfg.db.GetLocationSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			"areas": {
				dbQuery:     cfg.db.GetLocationAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetLocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetLocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			"shops": {
				dbQuery:     cfg.db.GetLocationShopIDs,
				createSubFn: createShopSimple,
			},
			"songs": {
				dbQuery:     cfg.getLocationSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetLocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.sublocations = handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "sublocations",
		resourceType:  "sublocation",
		objLookup:     cfg.l.Sublocations,
		objLookupID:   cfg.l.SublocationsID,
		queryLookup:   cfg.q.sublocations,
		idToResFunc:   idToNamedAPIResource[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSublocationIDs,
		getSingleFunc: cfg.getSublocation,
		retrieveFunc:  cfg.retrieveSublocations,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			"connected": {
				dbQuery:     cfg.db.GetConnectedSublocationIDs,
				createSubFn: createSublocationSimple,
				relationsFn: getSublocationSectionRelations,
			},
			"areas": {
				dbQuery:     cfg.db.GetSublocationAreaIDs,
				createSubFn: createAreaSimple,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetSublocationMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetSublocationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			"shops": {
				dbQuery:     cfg.db.GetSublocationShopIDs,
				createSubFn: createShopSimple,
			},
			"songs": {
				dbQuery:     cfg.getSublocationSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetSublocationTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.areas = handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]{
		endpoint:      "areas",
		resourceType:  "area",
		objLookup:     cfg.l.Areas,
		objLookupID:   cfg.l.AreasID,
		queryLookup:   cfg.q.areas,
		idToResFunc:   idToAreaAPIResource,
		resToListFunc: newAreaAPIResourceList,
		retrieveQuery: cfg.db.GetAreaIDs,
		getSingleFunc: cfg.getArea,
		retrieveFunc:  cfg.retrieveAreas,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			"connected": {
				dbQuery:     cfg.db.GetAreaConnectionIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetAreaMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetAreaMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			"songs": {
				dbQuery:     cfg.getAreaSongIDs,
				createSubFn: createSongSimple,
			},
			"treasures": {
				dbQuery:     cfg.db.GetAreaTreasureIDs,
				createSubFn: createTreasureSimple,
			},
		},
	}

	e.monsterFormations = handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "monster-formations",
		resourceType:  "monster formation",
		objLookup:     cfg.l.MonsterFormations,
		objLookupID:   cfg.l.MonsterFormationsID,
		queryLookup:   cfg.q.monsterFormations,
		idToResFunc:   idToUnnamedAPIResource[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetMonsterFormationIDs,
		getSingleFunc: cfg.getMonsterFormation,
		retrieveFunc:  cfg.retrieveMonsterFormations,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createMonsterFormationSimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetMonsterFormationMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.shops = handlerInput[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "shops",
		resourceType:  "shop",
		objLookup:     cfg.l.Shops,
		objLookupID:   cfg.l.ShopsID,
		queryLookup:   cfg.q.shops,
		idToResFunc:   idToUnnamedAPIResource[seeding.Shop, Shop, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetShopIDs,
		getSingleFunc: cfg.getShop,
		retrieveFunc:  cfg.retrieveShops,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createShopSimple,
			},
		},
	}

	e.treasures = handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "treasures",
		resourceType:  "treasure",
		objLookup:     cfg.l.Treasures,
		objLookupID:   cfg.l.TreasuresID,
		queryLookup:   cfg.q.treasures,
		idToResFunc:   idToUnnamedAPIResource[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
		retrieveQuery: cfg.db.GetTreasureIDs,
		getSingleFunc: cfg.getTreasure,
		retrieveFunc:  cfg.retrieveTreasures,
	}

	e.quests = handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      "quests",
		resourceType:  "quest",
		objLookup:     cfg.l.Quests,
		objLookupID:   cfg.l.QuestsID,
		queryLookup:   cfg.q.quests,
		idToResFunc:   idToQuestAPIResource[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetQuestIDs,
		getSingleFunc: cfg.getQuest,
		retrieveFunc:  cfg.retrieveQuests,
	}

	e.sidequests = handlerInput[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      "sidequests",
		resourceType:  "sidequest",
		objLookup:     cfg.l.Sidequests,
		objLookupID:   cfg.l.SidequestsID,
		queryLookup:   cfg.q.sidequests,
		idToResFunc:   idToQuestAPIResource[seeding.Sidequest, Sidequest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSidequestIDs,
		getSingleFunc: cfg.getSidequest,
		retrieveFunc:  cfg.retrieveSidequests,
		subsections: map[string]Subsection{
			"subquests": {
				dbQuery:     cfg.db.GetSidequestSubquestIDs,
				createSubFn: createSubquestSimple,
			},
		},
	}

	e.subquests = handlerInput[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList]{
		endpoint:      "subquests",
		resourceType:  "subquest",
		objLookup:     cfg.l.Subquests,
		objLookupID:   cfg.l.SubquestsID,
		queryLookup:   cfg.q.subquests,
		idToResFunc:   idToQuestAPIResource[seeding.Subquest, Subquest, QuestAPIResource, QuestApiResourceList],
		resToListFunc: newQuestAPIResourceList,
		retrieveQuery: cfg.db.GetSubquestIDs,
		getSingleFunc: cfg.getSubquest,
		retrieveFunc:  cfg.retrieveSubquests,
	}

	e.arenaCreations = handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "arena-creations",
		resourceType:  "arena creation",
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
		endpoint:      "blitzball-prizes",
		resourceType:  "blitzball prize table",
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
		endpoint:      "songs",
		resourceType:  "song",
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
		endpoint:      "fmvs",
		resourceType:  "fmv",
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
		endpoint:      "player-units",
		resourceType:  "player unit",
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
		endpoint:      "characters",
		resourceType:  "character",
		objLookup:     cfg.l.Characters,
		objLookupID:   cfg.l.CharactersID,
		queryLookup:   cfg.q.characters,
		idToResFunc:   idToNamedAPIResource[seeding.Character, Character, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterIDs,
		getSingleFunc: cfg.getCharacter,
		retrieveFunc:  cfg.retrieveCharacters,
		subsections: map[string]Subsection{
			"default-abilities": {
				dbQuery:     cfg.db.GetCharacterDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"std-sg-abilities": {
				dbQuery:     cfg.db.GetCharacterSgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"exp-sg-abilities": {
				dbQuery:     cfg.db.GetCharacterEgAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetCharacterOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
				relationsFn: getOverdriveAbilitySectionRelations,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetCharacterOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.aeons = handlerInput[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeons",
		resourceType:  "aeon",
		objLookup:     cfg.l.Aeons,
		objLookupID:   cfg.l.AeonsID,
		queryLookup:   cfg.q.aeons,
		idToResFunc:   idToNamedAPIResource[seeding.Aeon, Aeon, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAeonIDs,
		getSingleFunc: cfg.getAeon,
		retrieveFunc:  cfg.retrieveAeons,
		subsections: map[string]Subsection{
			"default-abilities": {
				dbQuery:     cfg.db.GetAeonDefaultAbilityIDs,
				createSubFn: createPlayerAbilitySimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetAeonOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
				relationsFn: getOverdriveAbilitySectionRelations,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetAeonOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			"stats": {
				createSubFn: createAeonStatSimple,
			},
		},
	}

	e.characterClasses = handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "character-classes",
		resourceType:  "character class",
		objLookup:     cfg.l.CharClasses,
		objLookupID:   cfg.l.CharClassesID,
		queryLookup:   cfg.q.characterClasses,
		idToResFunc:   idToNamedAPIResource[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCharacterClassesIDs,
		getSingleFunc: cfg.getCharacterClass,
		retrieveFunc:  cfg.retrieveCharacterClasses,
		subsections: map[string]Subsection{
			"default-abilities": {
				dbQuery:     cfg.db.GetCharacterClassDefaultAbilityIDs,
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
			"learnable-abilities": {
				dbQuery:     cfg.db.GetCharacterClassLearnableAbilityIDs,
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
			"default-overdrives": {
				dbQuery:     cfg.db.GetCharacterClassDefaultOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
			"learnable-overdrives": {
				dbQuery:     cfg.db.GetCharacterClassLearnableOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.monsters = handlerInput[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "monsters",
		resourceType:     "monster",
		objLookup:        cfg.l.Monsters,
		objLookupID:      cfg.l.MonstersID,
		queryLookup:      cfg.q.monsters,
		idToResFunc:      idToNamedAPIResource[seeding.Monster, Monster, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetMonsterIDsByName,
		retrieveQuery:    cfg.db.GetMonsterIDs,
		getSingleFunc:    cfg.getMonster,
		retrieveFunc:     cfg.retrieveMonsters,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
			"abilities": {
				dbQuery:     cfg.db.GetMonsterAbilityIDs,
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
			"areas": {
				dbQuery:     cfg.db.GetMonsterAreaIDs,
				createSubFn: createAreaSimple,
				relationsFn: getAreaSectionRelations,
			},
			"monster-formations": {
				dbQuery:     cfg.db.GetMonsterMonsterFormationIDs,
				createSubFn: createMonsterFormationSimple,
			},
		},
	}

	e.abilities = handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]{
		endpoint:      "abilities",
		resourceType:  "ability",
		objLookup:     cfg.l.Abilities,
		objLookupID:   cfg.l.AbilitiesID,
		queryLookup:   cfg.q.abilities,
		idToResFunc:   idToTypedAPIResource[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetAbilityIDs,
		getSingleFunc: cfg.getAbility,
		retrieveFunc:  cfg.retrieveAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
		},
	}

	e.playerAbilities = handlerInput[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "player-abilities",
		resourceType:     "player ability",
		objLookup:        cfg.l.PlayerAbilities,
		objLookupID:      cfg.l.PlayerAbilitiesID,
		queryLookup:      cfg.q.playerAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.PlayerAbility, PlayerAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetPlayerAbilityIDsByName,
		retrieveQuery:    cfg.db.GetPlayerAbilityIDs,
		getSingleFunc:    cfg.getPlayerAbility,
		retrieveFunc:     cfg.retrievePlayerAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createPlayerAbilitySimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetPlayerAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.overdriveAbilities = handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "overdrive-abilities",
		resourceType:     "overdrive ability",
		objLookup:        cfg.l.OverdriveAbilities,
		objLookupID:      cfg.l.OverdriveAbilitiesID,
		queryLookup:      cfg.q.overdriveAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetOverdriveAbilityIDsByName,
		retrieveQuery:    cfg.db.GetOverdriveAbilityIDs,
		getSingleFunc:    cfg.getOverdriveAbility,
		retrieveFunc:     cfg.retrieveOverdriveAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createOverdriveAbilitySimple,
				relationsFn: getOverdriveAbilitySectionRelations,
			},
		},
	}

	e.itemAbilities = handlerInput[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "item-abilities",
		resourceType:  "item ability",
		objLookup:     cfg.l.ItemAbilities,
		objLookupID:   cfg.l.ItemAbilitiesID,
		queryLookup:   cfg.q.itemAbilities,
		idToResFunc:   idToNamedAPIResource[seeding.ItemAbility, ItemAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemAbilityIDs,
		getSingleFunc: cfg.getItemAbility,
		retrieveFunc:  cfg.retrieveItemAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createItemAbilitySimple,
			},
		},
	}

	e.triggerCommands = handlerInput[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "trigger-commands",
		resourceType:     "trigger command",
		objLookup:        cfg.l.TriggerCommands,
		objLookupID:      cfg.l.TriggerCommandsID,
		queryLookup:      cfg.q.triggerCommands,
		idToResFunc:      idToNamedAPIResource[seeding.TriggerCommand, TriggerCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetTriggerCommandIDsByName,
		retrieveQuery:    cfg.db.GetTriggerCommandIDs,
		getSingleFunc:    cfg.getTriggerCommand,
		retrieveFunc:     cfg.retrieveTriggerCommands,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createTriggerCommandSimple,
			},
		},
	}

	e.unspecifiedAbilities = handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "unspecified-abilities",
		resourceType:     "unspecified ability",
		objLookup:        cfg.l.UnspecifiedAbilities,
		objLookupID:      cfg.l.UnspecifiedAbilitiesID,
		queryLookup:      cfg.q.unspecifiedAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetUnspecifiedAbilityIDsByName,
		retrieveQuery:    cfg.db.GetUnspecifiedAbilityIDs,
		getSingleFunc:    cfg.getUnspecifiedAbility,
		retrieveFunc:     cfg.retrieveUnspecifiedAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createUnspecifiedAbilitySimple,
			},
		},
	}

	e.enemyAbilities = handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:         "enemy-abilities",
		resourceType:     "enemy ability",
		objLookup:        cfg.l.EnemyAbilities,
		objLookupID:      cfg.l.EnemyAbilitiesID,
		queryLookup:      cfg.q.enemyAbilities,
		idToResFunc:      idToNamedAPIResource[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc:    newNamedAPIResourceList,
		getMultipleQuery: cfg.db.GetEnemyAbilityIDsByName,
		retrieveQuery:    cfg.db.GetEnemyAbilityIDs,
		getSingleFunc:    cfg.getEnemyAbility,
		retrieveFunc:     cfg.retrieveEnemyAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createEnemyAbilitySimple,
			},
			"monsters": {
				dbQuery:     cfg.db.GetEnemyAbilityMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.aeonCommands = handlerInput[seeding.AeonCommand, AeonCommand, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "aeon-commands",
		resourceType:  "aeon command",
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
		endpoint:      "overdrive-commands",
		resourceType:  "overdrive command",
		objLookup:     cfg.l.OverdriveCommands,
		objLookupID:   cfg.l.OverdriveCommandsID,
		queryLookup:   cfg.q.overdriveCommands,
		idToResFunc:   idToNamedAPIResource[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveCommandIDs,
		getSingleFunc: cfg.getOverdriveCommand,
		retrieveFunc:  cfg.retrieveOverdriveCommands,
		subsections: map[string]Subsection{
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetOverdriveCommandOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
				relationsFn: getOverdriveAbilitySectionRelations,
			},
			"overdrives": {
				dbQuery:     cfg.db.GetOverdriveCommandOverdriveIDs,
				createSubFn: createOverdriveSimple,
			},
		},
	}

	e.overdrives = handlerInput[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrives",
		resourceType:  "overdrive",
		objLookup:     cfg.l.Overdrives,
		objLookupID:   cfg.l.OverdrivesID,
		queryLookup:   cfg.q.overdrives,
		idToResFunc:   idToNamedAPIResource[seeding.Overdrive, Overdrive, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetOverdriveIDs,
		getSingleFunc: cfg.getOverdrive,
		retrieveFunc:  cfg.retrieveOverdrives,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createOverdriveSimple,
			},
			"overdrive-abilities": {
				dbQuery:     cfg.db.GetOverdriveOverdriveAbilityIDs,
				createSubFn: createOverdriveAbilitySimple,
				relationsFn: getOverdriveAbilitySectionRelations,
			},
		},
	}

	e.ronsoRages = handlerInput[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "ronso-rages",
		resourceType:  "ronso rage",
		objLookup:     cfg.l.RonsoRages,
		objLookupID:   cfg.l.RonsoRagesID,
		queryLookup:   cfg.q.ronsoRages,
		idToResFunc:   idToNamedAPIResource[seeding.RonsoRage, RonsoRage, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetRonsoRageIDs,
		getSingleFunc: cfg.getRonsoRage,
		retrieveFunc:  cfg.retrieveRonsoRages,
		subsections: map[string]Subsection{
			"monsters": {
				dbQuery:     cfg.db.GetRonsoRageMonsterIDs,
				createSubFn: createMonsterSimple,
				relationsFn: getMonsterSectionRelations,
			},
		},
	}

	e.submenus = handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "submenus",
		resourceType:  "submenu",
		objLookup:     cfg.l.Submenus,
		objLookupID:   cfg.l.SubmenusID,
		queryLookup:   cfg.q.submenus,
		idToResFunc:   idToNamedAPIResource[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSubmenuIDs,
		getSingleFunc: cfg.getSubmenu,
		retrieveFunc:  cfg.retrieveSubmenus,
		subsections: map[string]Subsection{
			"abilities": {
				dbQuery:     ToIntManyNull(cfg.db.GetSubmenuAbilityIDs),
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
		},
	}

	e.topmenus = handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "topmenus",
		resourceType:  "topmenu",
		objLookup:     cfg.l.Topmenus,
		objLookupID:   cfg.l.TopmenusID,
		queryLookup:   cfg.q.topmenus,
		idToResFunc:   idToNamedAPIResource[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetTopmenuIDs,
		getSingleFunc: cfg.getTopmenu,
		retrieveFunc:  cfg.retrieveTopmenus,
		subsections: map[string]Subsection{
			"abilities": {
				dbQuery:     ToIntManyNull(cfg.db.GetTopmenuAbilityIDs),
				createSubFn: createAbilitySimple,
				relationsFn: getAbilitySectionRelations,
			},
		},
	}

	e.allItems = handlerInput[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList]{
		endpoint:      "all-items",
		resourceType:  "all item",
		objLookup:     cfg.l.MasterItems,
		objLookupID:   cfg.l.MasterItemsID,
		queryLookup:   cfg.q.allItems,
		idToResFunc:   idToTypedAPIResource[seeding.MasterItem, MasterItem, TypedAPIResource, TypedAPIResourceList],
		resToListFunc: newTypedAPIResourceList,
		retrieveQuery: cfg.db.GetMasterItemIDs,
		getSingleFunc: cfg.getMasterItem,
		retrieveFunc:  cfg.retrieveMasterItems,
	}

	e.items = handlerInput[seeding.Item, Item, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "items",
		resourceType:  "item",
		objLookup:     cfg.l.Items,
		objLookupID:   cfg.l.ItemsID,
		queryLookup:   cfg.q.items,
		idToResFunc:   idToNamedAPIResource[seeding.Item, Item, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetItemIDs,
		getSingleFunc: cfg.getItem,
		retrieveFunc:  cfg.retrieveItems,
		subsections: map[string]Subsection{
			"mixes": {
				dbQuery:     cfg.db.GetItemMixIDs,
				createSubFn: createMixSimple,
			},
		},
	}

	e.spheres = handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "spheres",
		resourceType:  "sphere",
		objLookup:     cfg.l.Spheres,
		objLookupID:   cfg.l.SpheresID,
		queryLookup:   cfg.q.spheres,
		idToResFunc:   idToNamedAPIResource[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetSphereIDs,
		getSingleFunc: cfg.getSphere,
		retrieveFunc:  cfg.retrieveSpheres,
	}

	e.keyItems = handlerInput[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "key-items",
		resourceType:  "key-item",
		objLookup:     cfg.l.KeyItems,
		objLookupID:   cfg.l.KeyItemsID,
		queryLookup:   cfg.q.keyItems,
		idToResFunc:   idToNamedAPIResource[seeding.KeyItem, KeyItem, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetKeyItemIDs,
		getSingleFunc: cfg.getKeyItem,
		retrieveFunc:  cfg.retrieveKeyItems,
	}

	e.primers = handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "primers",
		resourceType:  "primer",
		objLookup:     cfg.l.Primers,
		objLookupID:   cfg.l.PrimersID,
		queryLookup:   cfg.q.primers,
		idToResFunc:   idToNamedAPIResource[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetPrimerIDs,
		getSingleFunc: cfg.getPrimer,
		retrieveFunc:  cfg.retrievePrimers,
	}

	e.mixes = handlerInput[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "mixes",
		resourceType:  "mix",
		objLookup:     cfg.l.Mixes,
		objLookupID:   cfg.l.MixesID,
		queryLookup:   cfg.q.mixes,
		idToResFunc:   idToNamedAPIResource[seeding.Mix, Mix, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetMixIDs,
		getSingleFunc: cfg.getMix,
		retrieveFunc:  cfg.retrieveMixes,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createMixSimple,
			},
		},
	}

	e.autoAbilities = handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "auto-abilities",
		resourceType:  "auto-ability",
		objLookup:     cfg.l.AutoAbilities,
		objLookupID:   cfg.l.AutoAbilitiesID,
		queryLookup:   cfg.q.autoAbilities,
		idToResFunc:   idToNamedAPIResource[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetAutoAbilityIDs,
		getSingleFunc: cfg.getAutoAbility,
		retrieveFunc:  cfg.retrieveAutoAbilities,
		subsections: map[string]Subsection{
			"simple": {
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.equipmentTables = handlerInput[seeding.EquipmentTable, EquipmentTable, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "equipment-tables",
		resourceType:  "equipment table",
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
		endpoint:      "equipment",
		resourceType:  "equipment",
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
		endpoint:      "celestial-weapons",
		resourceType:  "celestial weapon",
		objLookup:     cfg.l.CelestialWeapons,
		objLookupID:   cfg.l.CelestialWeaponsID,
		queryLookup:   cfg.q.celestialWeapons,
		idToResFunc:   idToNamedAPIResource[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
		retrieveQuery: cfg.db.GetCelestialWeaponIDs,
		getSingleFunc: cfg.getCelestialWeapon,
		retrieveFunc:  cfg.retrieveCelestialWeapons,
		subsections: map[string]Subsection{
			"auto-abilities": {
				dbQuery:     convertGetCelestialWeaponAutoAbilityIDs(cfg),
				createSubFn: createAutoAbilitySimple,
			},
		},
	}

	e.stats = handlerInput[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "stats",
		resourceType:  "stat",
		objLookup:     cfg.l.Stats,
		objLookupID:   cfg.l.StatsID,
		idToResFunc:   idToNamedAPIResource[seeding.Stat, Stat, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.properties = handlerInput[seeding.Property, Property, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "properties",
		resourceType:  "property",
		objLookup:     cfg.l.Properties,
		objLookupID:   cfg.l.PropertiesID,
		idToResFunc:   idToNamedAPIResource[seeding.Property, Property, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.overdriveModes = handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "overdrive-modes",
		resourceType:  "overdrive mode",
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
		endpoint:      "elements",
		resourceType:  "element",
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
		endpoint:      "status-conditions",
		resourceType:  "status condition",
		objLookup:     cfg.l.StatusConditions,
		objLookupID:   cfg.l.StatusConditionsID,
		idToResFunc:   idToNamedAPIResource[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.modifiers = handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList]{
		endpoint:      "modifiers",
		resourceType:  "modifier",
		objLookup:     cfg.l.Modifiers,
		objLookupID:   cfg.l.ModifiersID,
		idToResFunc:   idToNamedAPIResource[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList],
		resToListFunc: newNamedAPIResourceList,
	}

	e.agilityTiers = handlerInput[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList]{
		endpoint:      "agility-tiers",
		resourceType:  "agility tier",
		objLookup:     nil,
		objLookupID:   cfg.l.AgilityTiersID,
		idToResFunc:   idToUnnamedAPIResource[seeding.AgilityTier, AgilityTier, UnnamedAPIResource, UnnamedApiResourceList],
		resToListFunc: newUnnamedAPIResourceList,
	}

	e.abilityType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "ability-type",
		resourceType:  "ability type",
		objLookup:     cfg.t.AbilityType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.unitType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "unit-type",
		resourceType:  "unit type",
		objLookup:     cfg.t.UnitType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.itemType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "item-type",
		resourceType:  "item type",
		objLookup:     cfg.t.ItemType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.questType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "quest-type",
		resourceType:  "quest type",
		objLookup:     cfg.t.QuestType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.attackType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "attack-type",
		resourceType:  "attack type",
		objLookup:     cfg.t.AttackType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.autoAbilityCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "auto-ability-category",
		resourceType:  "auto ability category",
		objLookup:     cfg.t.AutoAbilityCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.availabilityType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "availability",
		resourceType:  "availability type",
		objLookup:     cfg.t.AvailabilityType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.damageFormula = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "damage-formula",
		resourceType:  "damage formula",
		objLookup:     cfg.t.DamageFormula.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.damageType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "damage-type",
		resourceType:  "damage type",
		objLookup:     cfg.t.DamageType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.elementalAffinity = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "elemental-affinities",
		resourceType:  "elemental affinity",
		objLookup:     cfg.t.ElementalAffinity.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.itemCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "item-category",
		resourceType:  "item category",
		objLookup:     cfg.t.ItemCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.keyItemCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "key-item-category",
		resourceType:  "key-item category",
		objLookup:     cfg.t.KeyItemCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.lootType = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "loot-type",
		resourceType:  "loot type",
		objLookup:     cfg.t.LootType.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.mixCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "mix-category",
		resourceType:  "mix category",
		objLookup:     cfg.t.MixCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-type",
		resourceType:  "monster type",
		objLookup:     cfg.t.MonsterCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterFormationCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-formation-category",
		resourceType:  "monster formation category",
		objLookup:     cfg.t.MonsterFormationCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.monsterSpecies = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "monster-species",
		resourceType:  "monster species",
		objLookup:     cfg.t.MonsterSpecies.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.playerAbilityCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "player-ability-category",
		resourceType:  "player ability category",
		objLookup:     cfg.t.PlayerAbilityCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	e.shopCategory = handlerInput[EnumAPIResource, EnumAPIResource, EnumAPIResource, EnumApiResourceList]{
		endpoint:      "shop-category",
		resourceType:  "shop category",
		objLookup:     cfg.t.ShopCategory.lookup,
		resToListFunc: newEnumAPIResourceList,
	}

	cfg.e = &e
}
