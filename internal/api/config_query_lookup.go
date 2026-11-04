package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	defaultParamSlice []QueryParam
	defaultParams     map[typestrings.QueryParamName]QueryParam

	locations    map[typestrings.QueryParamName]QueryParam
	sublocations map[typestrings.QueryParamName]QueryParam
	areas        map[typestrings.QueryParamName]QueryParam

	monsterFormations map[typestrings.QueryParamName]QueryParam
	shops             map[typestrings.QueryParamName]QueryParam
	treasures         map[typestrings.QueryParamName]QueryParam
	quests            map[typestrings.QueryParamName]QueryParam
	sidequests        map[typestrings.QueryParamName]QueryParam
	subquests         map[typestrings.QueryParamName]QueryParam
	arenaCreations    map[typestrings.QueryParamName]QueryParam
	blitzballPrizes   map[typestrings.QueryParamName]QueryParam
	songs             map[typestrings.QueryParamName]QueryParam
	fmvs              map[typestrings.QueryParamName]QueryParam

	playerUnits      map[typestrings.QueryParamName]QueryParam
	characters       map[typestrings.QueryParamName]QueryParam
	aeons            map[typestrings.QueryParamName]QueryParam
	characterClasses map[typestrings.QueryParamName]QueryParam
	monsters         map[typestrings.QueryParamName]QueryParam

	abilities          map[typestrings.QueryParamName]QueryParam
	playerAbilities    map[typestrings.QueryParamName]QueryParam
	overdriveAbilities map[typestrings.QueryParamName]QueryParam
	itemAbilities      map[typestrings.QueryParamName]QueryParam
	triggerCommands    map[typestrings.QueryParamName]QueryParam
	miscAbilities      map[typestrings.QueryParamName]QueryParam
	enemyAbilities     map[typestrings.QueryParamName]QueryParam

	aeonCommands      map[typestrings.QueryParamName]QueryParam
	overdriveCommands map[typestrings.QueryParamName]QueryParam
	overdrives        map[typestrings.QueryParamName]QueryParam
	ronsoRages        map[typestrings.QueryParamName]QueryParam
	submenus          map[typestrings.QueryParamName]QueryParam
	topmenus          map[typestrings.QueryParamName]QueryParam

	allItems map[typestrings.QueryParamName]QueryParam
	items    map[typestrings.QueryParamName]QueryParam
	keyItems map[typestrings.QueryParamName]QueryParam
	spheres  map[typestrings.QueryParamName]QueryParam
	primers  map[typestrings.QueryParamName]QueryParam
	mixes    map[typestrings.QueryParamName]QueryParam

	autoAbilities    map[typestrings.QueryParamName]QueryParam
	equipmentTables  map[typestrings.QueryParamName]QueryParam
	equipment        map[typestrings.QueryParamName]QueryParam
	celestialWeapons map[typestrings.QueryParamName]QueryParam

	stats            map[typestrings.QueryParamName]QueryParam
	properties       map[typestrings.QueryParamName]QueryParam
	overdriveModes   map[typestrings.QueryParamName]QueryParam
	elements         map[typestrings.QueryParamName]QueryParam
	statusConditions map[typestrings.QueryParamName]QueryParam
	modifiers        map[typestrings.QueryParamName]QueryParam
	agilityTiers     map[typestrings.QueryParamName]QueryParam

	enums     map[typestrings.QueryParamName]QueryParam
	endpoints map[typestrings.QueryParamName]QueryParam
}

func (cfg *Config) QueryLookupInit() {
	cfg.Q = &QueryLookup{}

	defaultParams := []QueryParam{
		{
			Name:        typestrings.QpnLimit,
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Type:        typestrings.QptInt,
			ParamUse:    puList,
			SpecialInputs: []SpecialQueryInput{
				{
					Key: typestrings.QsvMax,
					Val: 9999,
				},
			},
			DefaultVal: h.GetIntPtr(20),
		},
		{
			Name:        typestrings.QpnOffset,
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Type:        typestrings.QptInt,
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
