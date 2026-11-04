package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

var Q QueryLookup

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

func (Q *QueryLookup) Init() {
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

	Q.defaultParamSlice = defaultParams
	Q.defaultParams = paramSliceToMap(defaultParams, nil)
	initLocationsParams()
	initSublocationsParams()
	initAreasParams()

	initMonsterFormationsParams()
	initShopsParams()
	initTreasuresParams()
	initQuestsParams()
	initSidequestsParams()
	initSubquestsParams()
	initArenaCreationsParams()
	initBlitzballPrizesParams()
	initSongsParams()
	initFMVsParams()

	initPlayerUnitsParams()
	initCharactersParams()
	initAeonsParams()
	initCharacterClassesParams()
	initMonstersParams()

	initAbilitiesParams()
	initPlayerAbilitiesParams()
	initOverdriveAbilitiesParams()
	initItemAbilitiesParams()
	initTriggerCommandsParams()
	initMiscAbilitiesParams()
	initEnemyAbilitiesParams()

	Q.aeonCommands = E.AeonCommands.assignDefaultParams()
	Q.overdriveCommands = E.OverdriveCommands.assignDefaultParams()
	initOverdrivesParams()
	Q.ronsoRages = E.RonsoRages.assignDefaultParams()
	initSubmenusParams()
	Q.topmenus = E.Topmenus.assignDefaultParams()

	initAllItemsParams()
	initItemsParams()
	initKeyItemsParams()
	initSpheresParams()
	initPrimersParams()
	initMixesParams()

	initAutoAbilitiesParams()
	initEquipmentTablesParams()
	initEquipmentParams()
	initCelestialWeaponsParams()

	initStatsParams()
	Q.properties = E.Properties.assignDefaultParams()
	initOverdriveModesParams()
	Q.elements = E.Elements.assignDefaultParams()
	initStatusConditionsParams()
	initModifiersParams()
	initAgilityTierParams()

	Q.enums = E.Enums.assignDefaultParams()
	Q.endpoints = E.Endpoints.assignDefaultParams()

	E.AeonCommands.queryLookup = Q.aeonCommands
	E.OverdriveCommands.queryLookup = Q.overdriveCommands
	E.RonsoRages.queryLookup = Q.ronsoRages
	E.Topmenus.queryLookup = Q.topmenus
	E.Properties.queryLookup = Q.properties
	E.Elements.queryLookup = Q.elements

	E.Enums.queryLookup = Q.enums
	E.Endpoints.queryLookup = Q.endpoints
}
