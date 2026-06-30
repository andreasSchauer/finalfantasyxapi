package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	defaultParamSlice []QueryParam
	defaultParams     map[QueryParamName]QueryParam

	locations    map[QueryParamName]QueryParam
	sublocations map[QueryParamName]QueryParam
	areas        map[QueryParamName]QueryParam

	monsterFormations map[QueryParamName]QueryParam
	shops             map[QueryParamName]QueryParam
	treasures         map[QueryParamName]QueryParam
	quests            map[QueryParamName]QueryParam
	sidequests        map[QueryParamName]QueryParam
	subquests         map[QueryParamName]QueryParam
	arenaCreations    map[QueryParamName]QueryParam
	blitzballPrizes   map[QueryParamName]QueryParam
	songs             map[QueryParamName]QueryParam
	fmvs              map[QueryParamName]QueryParam

	playerUnits      map[QueryParamName]QueryParam
	characters       map[QueryParamName]QueryParam
	aeons            map[QueryParamName]QueryParam
	characterClasses map[QueryParamName]QueryParam
	monsters         map[QueryParamName]QueryParam

	abilities          map[QueryParamName]QueryParam
	playerAbilities    map[QueryParamName]QueryParam
	overdriveAbilities map[QueryParamName]QueryParam
	itemAbilities      map[QueryParamName]QueryParam
	triggerCommands    map[QueryParamName]QueryParam
	miscAbilities      map[QueryParamName]QueryParam
	enemyAbilities     map[QueryParamName]QueryParam

	aeonCommands      map[QueryParamName]QueryParam
	overdriveCommands map[QueryParamName]QueryParam
	overdrives        map[QueryParamName]QueryParam
	ronsoRages        map[QueryParamName]QueryParam
	submenus          map[QueryParamName]QueryParam
	topmenus          map[QueryParamName]QueryParam

	allItems map[QueryParamName]QueryParam
	items    map[QueryParamName]QueryParam
	keyItems map[QueryParamName]QueryParam
	spheres  map[QueryParamName]QueryParam
	primers  map[QueryParamName]QueryParam
	mixes    map[QueryParamName]QueryParam

	autoAbilities    map[QueryParamName]QueryParam
	equipmentTables  map[QueryParamName]QueryParam
	equipment        map[QueryParamName]QueryParam
	celestialWeapons map[QueryParamName]QueryParam

	stats            map[QueryParamName]QueryParam
	properties       map[QueryParamName]QueryParam
	overdriveModes   map[QueryParamName]QueryParam
	elements         map[QueryParamName]QueryParam
	statusConditions map[QueryParamName]QueryParam
	modifiers        map[QueryParamName]QueryParam
	agilityTiers     map[QueryParamName]QueryParam
}

func (cfg *Config) QueryLookupInit() {
	cfg.q = &QueryLookup{}

	defaultParams := []QueryParam{
		{
			Name:        qpnLimit,
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Type:        qptInt,
			ForList:     true,
			ForSingle:   false,
			SpecialInputs: []SpecialQueryInput{
				{
					Key: qsvMax,
					Val: 9999,
				},
			},
			DefaultVal: h.GetIntPtr(20),
		},
		{
			Name:        qpnOffset,
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Type:        qptInt,
			ForList:     true,
			ForSingle:   false,
			DefaultVal:  h.GetIntPtr(0),
		},
	}

	cfg.q.defaultParamSlice = defaultParams
	cfg.q.defaultParams = querySliceToMap(cfg, defaultParams)
	cfg.initLocationsParams()
	cfg.initSublocationsParams()
	cfg.initAreasParams()

	cfg.initMonsterFormationsParams()
	cfg.initShopsParams()
	cfg.initTreasuresParams()
	cfg.initQuestsParams()
	cfg.initSidequestsParams()
	cfg.initSubquestsParams()
	cfg.initArenaCreationsParams()
	cfg.initBlitzballPrizesParams()
	cfg.initSongsParams()
	cfg.initFMVsParams()

	cfg.initPlayerUnitsParams()
	cfg.initCharactersParams()
	cfg.initAeonsParams()
	cfg.initCharacterClassesParams()
	cfg.initMonstersParams()

	cfg.initAbilitiesParams()
	cfg.initPlayerAbilitiesParams()
	cfg.initOverdriveAbilitiesParams()
	cfg.initItemAbilitiesParams()
	cfg.initTriggerCommandsParams()
	cfg.initMiscAbilitiesParams()
	cfg.initEnemyAbilitiesParams()

	cfg.q.aeonCommands = cfg.assignDefaultParams()
	cfg.q.overdriveCommands = cfg.assignDefaultParams()
	cfg.initOverdrivesParams()
	cfg.q.ronsoRages = cfg.assignDefaultParams()
	cfg.initSubmenusParams()
	cfg.q.topmenus = cfg.assignDefaultParams()

	cfg.initAllItemsParams()
	cfg.initItemsParams()
	cfg.initKeyItemsParams()
	cfg.initSpheresParams()
	cfg.initPrimersParams()
	cfg.initMixesParams()

	cfg.initAutoAbilitiesParams()
	cfg.initEquipmentTablesParams()
	cfg.initEquipmentParams()
	cfg.initCelestialWeaponsParams()

	cfg.initStatsParams()
	cfg.q.properties = cfg.assignDefaultParams()
	cfg.initOverdriveModesParams()
	cfg.q.elements = cfg.assignDefaultParams()
	cfg.initStatusConditionsParams()
	cfg.initModifiersParams()
	cfg.initAgilityTierParams()
}