package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	defaultParamSlice []QueryParam
	defaultParams     map[ts.QueryParamName]QueryParam

	locations    map[ts.QueryParamName]QueryParam
	sublocations map[ts.QueryParamName]QueryParam
	areas        map[ts.QueryParamName]QueryParam

	monsterFormations map[ts.QueryParamName]QueryParam
	shops             map[ts.QueryParamName]QueryParam
	treasures         map[ts.QueryParamName]QueryParam
	quests            map[ts.QueryParamName]QueryParam
	sidequests        map[ts.QueryParamName]QueryParam
	subquests         map[ts.QueryParamName]QueryParam
	arenaCreations    map[ts.QueryParamName]QueryParam
	blitzballPrizes   map[ts.QueryParamName]QueryParam
	songs             map[ts.QueryParamName]QueryParam
	fmvs              map[ts.QueryParamName]QueryParam

	playerUnits      map[ts.QueryParamName]QueryParam
	characters       map[ts.QueryParamName]QueryParam
	aeons            map[ts.QueryParamName]QueryParam
	characterClasses map[ts.QueryParamName]QueryParam
	monsters         map[ts.QueryParamName]QueryParam

	abilities          map[ts.QueryParamName]QueryParam
	playerAbilities    map[ts.QueryParamName]QueryParam
	overdriveAbilities map[ts.QueryParamName]QueryParam
	itemAbilities      map[ts.QueryParamName]QueryParam
	triggerCommands    map[ts.QueryParamName]QueryParam
	miscAbilities      map[ts.QueryParamName]QueryParam
	enemyAbilities     map[ts.QueryParamName]QueryParam

	aeonCommands      map[ts.QueryParamName]QueryParam
	overdriveCommands map[ts.QueryParamName]QueryParam
	overdrives        map[ts.QueryParamName]QueryParam
	ronsoRages        map[ts.QueryParamName]QueryParam
	submenus          map[ts.QueryParamName]QueryParam
	topmenus          map[ts.QueryParamName]QueryParam

	allItems map[ts.QueryParamName]QueryParam
	items    map[ts.QueryParamName]QueryParam
	keyItems map[ts.QueryParamName]QueryParam
	spheres  map[ts.QueryParamName]QueryParam
	primers  map[ts.QueryParamName]QueryParam
	mixes    map[ts.QueryParamName]QueryParam

	autoAbilities    map[ts.QueryParamName]QueryParam
	equipmentTables  map[ts.QueryParamName]QueryParam
	equipment        map[ts.QueryParamName]QueryParam
	celestialWeapons map[ts.QueryParamName]QueryParam

	stats            map[ts.QueryParamName]QueryParam
	properties       map[ts.QueryParamName]QueryParam
	overdriveModes   map[ts.QueryParamName]QueryParam
	elements         map[ts.QueryParamName]QueryParam
	statusConditions map[ts.QueryParamName]QueryParam
	modifiers        map[ts.QueryParamName]QueryParam
	agilityTiers     map[ts.QueryParamName]QueryParam

	enums     map[ts.QueryParamName]QueryParam
	endpoints map[ts.QueryParamName]QueryParam
}

func (cfg *Config) QueryLookupInit() {
	cfg.Q = &QueryLookup{}

	defaultParams := []QueryParam{
		{
			Name:        ts.QpnLimit,
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Type:        ts.QptInt,
			ParamUse:    puList,
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvMax,
					Val: 9999,
				},
			},
			DefaultVal: h.GetIntPtr(20),
		},
		{
			Name:        ts.QpnOffset,
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Type:        ts.QptInt,
			ParamUse:    puList,
			DefaultVal:  h.GetIntPtr(0),
		},
	}

	cfg.Q.defaultParamSlice = defaultParams
	cfg.Q.defaultParams = paramSliceToMap(defaultParams, nil)
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

	cfg.Q.aeonCommands = cfg.E.AeonCommands.assignDefaultParams(cfg)
	cfg.Q.overdriveCommands = cfg.E.OverdriveCommands.assignDefaultParams(cfg)
	cfg.initOverdrivesParams()
	cfg.Q.ronsoRages = cfg.E.RonsoRages.assignDefaultParams(cfg)
	cfg.initSubmenusParams()
	cfg.Q.topmenus = cfg.E.Topmenus.assignDefaultParams(cfg)

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
	cfg.Q.properties = cfg.E.Properties.assignDefaultParams(cfg)
	cfg.initOverdriveModesParams()
	cfg.Q.elements = cfg.E.Elements.assignDefaultParams(cfg)
	cfg.initStatusConditionsParams()
	cfg.initModifiersParams()
	cfg.initAgilityTierParams()

	cfg.Q.enums = cfg.E.Enums.assignDefaultParams(cfg)
	cfg.Q.endpoints = cfg.E.Endpoints.assignDefaultParams(cfg)

	cfg.E.AeonCommands.queryLookup = cfg.Q.aeonCommands
	cfg.E.OverdriveCommands.queryLookup = cfg.Q.overdriveCommands
	cfg.E.RonsoRages.queryLookup = cfg.Q.ronsoRages
	cfg.E.Topmenus.queryLookup = cfg.Q.topmenus
	cfg.E.Properties.queryLookup = cfg.Q.properties
	cfg.E.Elements.queryLookup = cfg.Q.elements

	cfg.E.Enums.queryLookup = cfg.Q.enums
	cfg.E.Endpoints.queryLookup = cfg.Q.endpoints
}
