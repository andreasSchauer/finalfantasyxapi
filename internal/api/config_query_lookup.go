package api

import (
	"fmt"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QueryParam struct {
	ID               int                        `json:"-"`
	Name             QueryParamName             `json:"name"`
	Type             QueryParamType             `json:"param_type"`
	Description      string                     `json:"description"`
	ExampleVals      []string                   `json:"-"`
	Usage            string                     `json:"usage"`
	ExampleUses      []string                   `json:"example_uses"`
	DefaultOnly      bool                       `json:"only_use_alone"`
	ForSingle        bool                       `json:"for_single"`
	ForList          bool                       `json:"for_list"`
	ForSegment       *SectionName               `json:"for_segment"`
	IsRequired       bool                       `json:"is_required"`
	TypeLookup       map[string]EnumAPIResource `json:"-"`
	RequiredParams   []QueryParamName           `json:"required_params,omitempty"`
	UsableWith       []QueryParamName           `json:"usable_with,omitempty"`
	ReplacedBy       []QueryParamName           `json:"replaced_by,omitempty"`
	ForbiddenParams  []QueryParamName           `json:"forbidden_params,omitempty"`
	ReferencesInt    []EndpointName             `json:"-"`
	References       []string                   `json:"references,omitempty"`
	AllowedIDs       []int32                    `json:"-"`
	AllowedResources []string                   `json:"allowed_resources,omitempty"`
	AllowedValues    []QueryValue               `json:"allowed_values,omitempty"`
	AllowedIntRange  []int                      `json:"allowed_int_range,omitempty"`
	AllowedResTypes  []string                   `json:"allowed_res_types,omitempty"`
	DefaultVal       *int                       `json:"default_value,omitempty"`
	SpecialInputs    []SpecialQueryInput        `json:"special_inputs,omitempty"`
}

type SpecialQueryInput struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

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
					Key: "max",
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
		{
			Name:        qpnFlip,
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
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

func (cfg *Config) assignDefaultParams() map[QueryParamName]QueryParam {
	return cfg.completeQueryParamsInit([]QueryParam{}, false)
}

func (cfg *Config) completeQueryParamsInit(params []QueryParam, hasSimpleView bool) map[QueryParamName]QueryParam {
	params = slices.Concat(params, cfg.q.defaultParamSlice)

	if hasSimpleView {
		queryParamIDs := QueryParam{
			Name:        qpnIDs,
			Description: "Used to input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type:        qptIdList,
			DefaultOnly: true,
			ForList:     false,
			ForSingle:   false,
			ForSegment:  getSnPtr(snSimple),
		}
		params = append(params, queryParamIDs)
	}

	return querySliceToMap(cfg, params)
}

func (cfg *Config) assignParamUsage(p QueryParam) QueryParam {
	s := fmt.Sprintf("?%s=", p.Name)

	switch p.Type {
	case qptBool:
		p.Usage = s + "{bool}"
		p.ExampleUses = []string{s + "true", s + "false"}

	case qptEnum:
		enums := createEnumResourceSlice(cfg, "", p.TypeLookup)
		e := enums[0].Name
		p.Usage = s + "{value|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case qptEnumList:
		enums := createEnumResourceSlice(cfg, "", p.TypeLookup)
		e1 := enums[0].Name
		e2 := enums[1].Name
		p.Usage = s + "{value|id},..."
		p.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptId:
		p.Usage = s + "{id}"
		p.ExampleUses = []string{s + "1"}

	case qptIdNul:
		p.Usage = s + "{id|'none'}"
		p.ExampleUses = []string{s + "1", s + "none"}

	case qptIdList:
		p.Usage = s + "{id},..."
		p.ExampleUses = []string{s + "1", s + "1,2"}

	case qptInt:
		p.Usage = s + "{int}"
		p.ExampleUses = []string{s + "1"}

	case qptIntList:
		p.Usage = s + "{int},...|{int}-{int}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}

	case qptNameId:
		e := p.ExampleVals[0]
		p.Usage = s + "{name|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case qptNameIdList:
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},..."
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptNameIdListNul:
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},...|{'none'}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}

	case qptValue:
		e1 := string(p.AllowedValues[0])
		e2 := string(p.AllowedValues[1])
		p.Usage = s + "{val}"
		p.ExampleUses = []string{s + e1, s + e2}

	case qptValueList:
		e1 := string(p.AllowedValues[0])
		e2 := string(p.AllowedValues[1])
		p.Usage = s + "{val}"
		p.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}

	case qptStat:
		p.Usage = s + "{stat}={int},..."

	default:
		return p
	}

	if p.SpecialInputs != nil {
		for _, input := range p.SpecialInputs {
			usageTrimmed := strings.TrimSuffix(p.Usage, "}")
			p.Usage = fmt.Sprintf("%s|'%s'}", usageTrimmed, input.Key)
			p.ExampleUses = append(p.ExampleUses, s+input.Key)
		}
	}

	return p
}

func (cfg *Config) initLocationsParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays a location's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a location's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for locations with the given availabilities. Can be combined with other parameters that filter locations by resource/resource-type. In that case, this parameter searches for locations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same location, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for locations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for locations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the location. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for locations where the specified monster can be found. If combined with 'availability', the location must contain a monster formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for locations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for locations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the location whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for locations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for locations where a character permanently joins the party.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for locations where a new aeon is acquired.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for locations that have monsters. If combined with 'availability', the location must inhabit at least one monster formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for locations that have bosses. If combined with 'availability', the location must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster formation).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for locations that have shops. If combined with 'availability', the location must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for locations that have treasures. If combined with 'availability', the location must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for locations that feature sidequests. If combined with 'availability', the location must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for locations that feature fmv sequences.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.locations = paramsMap
}

func (cfg *Config) initSublocationsParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays a sublocation's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a sublocation's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for sublocations that are located within the specified location.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for sublocations with the given availabilities. Can be combined with other parameters that filter sublocations by resource/resource-type. In that case, this parameter searches for sublocations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same sublocation, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the sublocation won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for sublocations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for sublocations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the sublocation. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for sublocations where the specified monster can be found. If combined with 'availability', the sublocation must contain a monster formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for sublocations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for sublocations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for sublocations that have monsters. If combined with 'availability', the sublocation must inhabit at least one monster formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for sublocations that have bosses. If combined with 'availability', the sublocation must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster formation).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for sublocations that have shops. If combined with 'availability', the sublocation must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for sublocations that have treasures. If combined with 'availability', the sublocation must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for sublocations that feature sidequests. If combined with 'availability', the sublocation must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.sublocations = paramsMap
}

func (cfg *Config) initAreasParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays an area's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an area's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for areas that are located within the specified location.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for areas that are located within the specified sublocation.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for areas with the given availabilities. Can be combined with other parameters that filter areas by resource/resource-type. In that case, this parameter searches for areas where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same area, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for areas where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for areas where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the area. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for areas where the specified monster can be found. If combined with 'availability', the area must contain a monster formation with the monster and whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', the monster must have a monster formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for areas where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for areas where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the area whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for areas where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnSaveSphere,
			Description: "Searches for areas that have a save sphere.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnCompSphere,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnAirship,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnChocobo,
			Description: "Searches for areas where you can ride a chocobo.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for areas where a character permanently joins the party.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for areas where a new aeon is acquired.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for areas that have monsters. If combined with 'availability', the area must inhabit at least one monster formation whose most accessible availability matches one of the specified availabilities (based on its encounter areas).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for areas that have bosses. If combined with 'availability', the area must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster formation's encounter areas).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for areas that have shops. If combined with 'availability', the area must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for areas that have treasures. If combined with 'availability', the area must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for areas that feature sidequests. If combined with 'availability', the area must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for areas that feature fmv sequences.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.areas = paramsMap
}

func (cfg *Config) initMonsterFormationsParams() {
	params := []QueryParam{
		{
			Name:          qpnMonster,
			Description:   "Searches for monster formations that feature the specified monster.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for monster formations with the specified monster-formation-categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.MonsterFormationCategory.lookup,
			ReferencesInt: []EndpointName{epMonsterFormationCategory},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for monster formations with the given availabilities. If combined with the 'area' parameter, the availability of the monster formation in this specific area is used. If a monster-formation has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster formation won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with the 'area' parameter, it takes the repeatability directly from the monster formations that occur in the specified area. Is combinable with 'availability'. In that case, the search looks for the monster formation that is the most accessible while also being farmable and checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for monster formations that appear within the specified location. If combined with 'availability', this parameter searches for monster formations within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster formations within this location whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for monster formations that appear within the specified sublocation. If combined with 'availability', this parameter searches for monster formations within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster formations within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for monster formations that appear within the specified area. If combined with 'availability', this parameter searches for monster formations within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monster formations within this area whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        qpnAmbush,
			Description: "Searches for monster formations that are forced ambushes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.monsterFormations = paramsMap
}

func (cfg *Config) initShopsParams() {
	params := []QueryParam{
		{
			Name:          qpnCategory,
			Description:   "Searches for shops with the specified shop categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.ShopCategory.lookup,
			ReferencesInt: []EndpointName{epShopCategory},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for shops with the given availabilities. By default, this parameter checks, if a shop simply is available at the given availability. If combined with a filter that refers to the shop's inventory, it takes the availability directly from there ('pre-story' for the inventory before acquiring the airship, and 'post' after). In that case, this filter looks, if the requested resources are available in a shop at that point in the game.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for shops that appear at the specified location. If combined with 'availability', this parameter searches for shops within this location whose availability matches one of the specified availabilities.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnAvailability},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for shops that appear at the specified sublocation. If combined with 'availability', this parameter searches for shops within this sublocation whose availability matches one of the specified availabilities.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:            qpnAutoAbility,
			Description:     "Searches for shops that offer equipment with the specified auto-ability. Can be combined with 'empty_slots' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptId,
			ForList:         true,
			ForSingle:       false,
			ForbiddenParams: []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:      []QueryParamName{qpnAvailability},
			ReferencesInt:   []EndpointName{epAutoAbilities},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots. Can be combined with 'auto_ability' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptIntList,
			ForList:         true,
			ForSingle:       false,
			ForbiddenParams: []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:      []QueryParamName{qpnAvailability},
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:            qpnCharacter,
			Description:     "Searches for shops that offer equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptNameId,
			ExampleVals:     []string{"wakka"},
			ForList:         true,
			ForSingle:       false,
			ForbiddenParams: []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:      []QueryParamName{qpnAvailability},
			ReferencesInt:   []EndpointName{epCharacters},
		},
		{
			Name:            qpnItems,
			Description:     "Searches for shops that offer items. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptBool,
			ForList:         true,
			ForSingle:       false,
			ForbiddenParams: []QueryParamName{qpnAutoAbility, qpnCharacter, qpnEmptySlots},
			ReplacedBy:      []QueryParamName{qpnAvailability},
		},
		{
			Name:            qpnEquipment,
			Description:     "Searches for shops that offer equipment. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptBool,
			ForList:         true,
			ForSingle:       false,
			ForbiddenParams: []QueryParamName{qpnAutoAbility, qpnCharacter, qpnEmptySlots},
			ReplacedBy:      []QueryParamName{qpnAvailability},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.shops = paramsMap
}

func (cfg *Config) initTreasuresParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for treasures that appear within the specified location.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for treasures that appear within the specified sublocation.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for treasures that appear within the specified area.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for treasures that contain equipment with the specified auto-ability. Can be combined with 'empty_slots', 'character', and 'availability' for more specific searches.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for treasures that contain equipment with the specified amounts of empty slots. Can be combined with 'auto_ability', 'character', and 'availability' for more specific searches.",
			Type:            qptIntList,
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          qpnCharacter,
			Description:   "Searches for treasures that contain equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for treasures that contain the specified item.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:          qpnLootType,
			Description:   "Searches for treasures with the specified loot type.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.LootType.lookup,
			ReferencesInt: []EndpointName{epLootType},
		},
		{
			Name:        qpnTreasureType,
			Description: "Searches for treasures with the specified treasure type.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TreasureType.lookup,
		},
		{
			Name:        qpnAnima,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for treasures with the given availabilities.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.treasures = paramsMap
}

func (cfg *Config) initQuestsParams() {
	params := []QueryParam{
		{
			Name:          qpnType,
			Description:   "Searches for quests that are of the specified quest type.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.QuestType.lookup,
			ReferencesInt: []EndpointName{epQuestType},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for quests with the given availabilities.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for quests that can be completed more than once.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.quests = paramsMap
}

func (cfg *Config) initSidequestsParams() {
	params := []QueryParam{
		{
			Name:          qpnAvailability,
			Description:   "Searches for sidequests with the given availabilities.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.sidequests = paramsMap
}

func (cfg *Config) initSubquestsParams() {
	params := []QueryParam{
		{
			Name:          qpnAvailability,
			Description:   "Searches for subquests with the given availabilities.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for subquests that can be completed more than once.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.subquests = paramsMap
}

func (cfg *Config) initArenaCreationsParams() {
	params := []QueryParam{
		{
			Name:        qpnCategory,
			Description: "Searches for monster formations with the specified arena-creation-categories.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ArenaCreationCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.arenaCreations = paramsMap
}

func (cfg *Config) initBlitzballPrizesParams() {
	params := []QueryParam{
		{
			Name:        qpnCategory,
			Description: "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.BlitzballTournamentCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.blitzballPrizes = paramsMap
}

func (cfg *Config) initSongsParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for songs that are played in fmvs.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnSpecialUse,
			Description: "Searches for songs with a special use case.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnComposer,
			Description: "Searches for songs that were composed by the stated composer.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Composer.lookup,
		},
		{
			Name:        qpnArranger,
			Description: "Searches for songs that were arranged by the stated arranger.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Arranger.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.songs = paramsMap
}

func (cfg *Config) initFMVsParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for fmvs that are played within the specified location.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epLocations},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.fmvs = paramsMap
}

func (cfg *Config) initPlayerUnitsParams() {
	params := []QueryParam{
		{
			Name:          qpnType,
			Description:   "Searches for player units that are of the specified unit type.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.UnitType.lookup,
			ReferencesInt: []EndpointName{epUnitType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.playerUnits = paramsMap
}

func (cfg *Config) initCharactersParams() {
	params := []QueryParam{
		{
			Name:        qpnStoryBased,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnUnderwater,
			Description: "Searches for characters that can fight underwater.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.characters = paramsMap
}

func (cfg *Config) initAeonsParams() {
	params := []QueryParam{
		{
			Name:            qpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            qptInt,
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:        qpnYunaStats,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used instead. Every stat instead of luck is available, since an aeon simply copies Yuna's luck stat. Can be used in combination with the 'battles' parameter.",
			Type:        qptStat,
			ExampleUses: []string{"?yuna_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?yuna_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:        qpnOptional,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.aeons = paramsMap
}

func (cfg *Config) initCharacterClassesParams() {
	params := []QueryParam{
		{
			Name:        qpnCategory,
			Description: "Searches for character classes with the specified categories.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CharacterClassCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.characterClasses = paramsMap
}

func (cfg *Config) initMonstersParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays a monster's related resources with the given availabilities. This affects areas and monster-formations.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a monster's related resources that can be farmed. This affects areas and monster-formations.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:        qpnKimahriStats,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type:        qptStat,
			ExampleUses: []string{"?kimahri_stats=hp=3000,strength=25,magic=30,agility=40", "?kimahri_stats=hp=15000,agility=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:        qpnAeonStats,
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Type:        qptStat,
			ExampleUses: []string{"?aeon_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?aeon_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        qpnAlteredState,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:        qptId,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnOmnisElements,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type:          "other",
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wilf"},
			ForList:       false,
			ForSingle:     true,
			AllowedIDs:    []int32{211},
			AllowedValues: []QueryValue{qvF, qvL, qvW, qvI},
		},
		{
			Name:          qpnElementalResists,
			Description:   "Searches for monsters that have the specified elemental affinities.",
			Type:          "other",
			Usage:         "?elemental_resists={element|id}={affinity|id},...",
			ExampleUses:   []string{"?elemental_resists=fire=weak,water=absorb", "?elemental_resists=1=3,2=4"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements, epElementalAffinity},
		},
		{
			Name:          qpnStatusResists,
			Description:   "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type:          qptIdList,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:            qpnResistance,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            qptInt,
			ForList:         true,
			ForSingle:       false,
			RequiredParams:  []QueryParamName{qpnStatusResists},
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: "immune",
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:          qpnItem,
			Description:   "Searches for monsters that have the specified item as loot. Can be specified further with the 'methods' parameter.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvSteal, qvDrop, qvBribe, qvOther},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for monsters that drop the specified auto-ability.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:           qpnIsForced,
			Description:    "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAutoAbility},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type:            qptIntList,
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:          qpnRonsoRage,
			Description:   "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epRonsoRages},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for monsters with the given availabilities. If combined with a geographical filter, it takes the availability directly from the monster formations that occur in the specified location, sublocation, or area. If a monster has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with a geographical filter, it takes the repeatability directly from the monster formations that occur in the specified location, sublocation, or area. Is combinable with 'availability'. The availability assigned to the monster is from the monster formation that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for monsters that appear within the specified location. If combined with 'availability', this parameter searches for monsters that are part of at least one monster formation within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster formation within this location whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for monsters that appear within the specified sublocation. If combined with 'availability', this parameter searches for monsters that are part of at least one monster formation within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster formation within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for monsters that appear within the specified area. If combined with 'availability', this parameter searches for monsters that are part of at least one monster formation within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster formation within this area whose farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:            qpnDistance,
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type:            qptIntList,
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:        qpnCapture,
			Description: "Searches for monsters that can be captured.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnHasOverdrive,
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnUnderwater,
			Description: "Searches for monsters that are fought underwater.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnZombie,
			Description: "Searches for monsters that are zombies.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnSpecies,
			Description:   "Searches for monsters of the specified species.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.MonsterSpecies.lookup,
			ReferencesInt: []EndpointName{epMonsterSpecies},
		},
		{
			Name:        qpnCreationArea,
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CreationArea.lookup,
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for monsters that are of the specified monster-categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.MonsterCategory.lookup,
			ReferencesInt: []EndpointName{epMonsterCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.monsters = paramsMap
}

func (cfg *Config) initAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnType,
			Description:   "Searches for abilities that are of the specified ability types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AbilityType.lookup,
			ReferencesInt: []EndpointName{epAbilityType},
		},
		{
			Name:        qpnRank,
			Description: "Searches for abilities with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for abilities that can be used by the specified monster.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:        qpnTargetType,
			Description: "Searches for abilities with the specified target types.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnAttackType,
			Description:   "Searches for abilities with battle interactions of the specified attack types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AttackType.lookup,
			ReferencesInt: []EndpointName{epAttackType},
		},
		{
			Name:          qpnDamageType,
			Description:   "Searches for abilities that deal the specified types of damage.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageType.lookup,
			ReferencesInt: []EndpointName{epDamageType},
		},
		{
			Name:          qpnDamageFormula,
			Description:   "Searches for abilities that use the specified formula to calculate their damage.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageFormula.lookup,
			ReferencesInt: []EndpointName{epDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for abilities that can land critical hits.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnBDL,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for abilities that deal delay.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for abilities that cause stat changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for abilities that cause modifier changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.abilities = paramsMap
}

func (cfg *Config) initPlayerAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ForList:       false,
			ForSingle:     true,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:        qpnRank,
			Description: "Searches for player abilities with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for player abilities that are of the specified player ability categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.PlayerAbilityCategory.lookup,
			ReferencesInt: []EndpointName{epPlayerAbilityCategory},
		},
		{
			Name:        qpnOutsideBattle,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMp,
			Description: "Searches for player abilities with the specified mp costs.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMpMin,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type:        qptInt,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnMpMax,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type:        qptInt,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for player abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          qpnUser,
			Description:   "Searches for player abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          qpnStdSg,
			Description:   "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:          qptNameId,
			ExampleVals:   []string{"tidus", "wakka"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnExpSg,
			Description:   "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:          qptNameId,
			ExampleVals:   []string{"tidus", "wakka"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnLearnItem,
			Description:   "Searches for player abilities an aeon can learn via the specified item.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:        qpnTargetType,
			Description: "Searches for player abilities with the specified target types.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnAttackType,
			Description:   "Searches for player abilities with battle interactions of the specified attack types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AttackType.lookup,
			ReferencesInt: []EndpointName{epAttackType},
		},
		{
			Name:          qpnDamageType,
			Description:   "Searches for player abilities that deal the specified types of damage.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageType.lookup,
			ReferencesInt: []EndpointName{epDamageType},
		},
		{
			Name:          qpnDamageFormula,
			Description:   "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageFormula.lookup,
			ReferencesInt: []EndpointName{epDamageFormula},
		},
		{
			Name:          qpnElement,
			Description:   "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for player abilities that deal delay.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for player abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for player abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for player abilities that cause stat changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for player abilities that cause modifier changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.playerAbilities = paramsMap
}

func (cfg *Config) initOverdriveAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for overdrive abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for overdrive abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:        qpnTargetType,
			Description: "Searches for overdrive abilities with the specified target types.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:          qpnAttackType,
			Description:   "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AttackType.lookup,
			ReferencesInt: []EndpointName{epAttackType},
		},
		{
			Name:          qpnDamageFormula,
			Description:   "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageFormula.lookup,
			ReferencesInt: []EndpointName{epDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for overdrive abilities that deal delay.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for overdrive abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.overdriveAbilities = paramsMap
}

func (cfg *Config) initItemAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnCategory,
			Description:   "Searches for item abilities that are of the specified item categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.ItemCategory.lookup,
			ReferencesInt: []EndpointName{epItemCategory},
		},
		{
			Name:        qpnOutsideBattle,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for item abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:        qpnTargetType,
			Description: "Searches for item abilities with the specified target types.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:          qpnAttackType,
			Description:   "Searches for item abilities with battle interactions of the specified attack types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AttackType.lookup,
			ReferencesInt: []EndpointName{epAttackType},
		},
		{
			Name:          qpnDamageFormula,
			Description:   "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageFormula.lookup,
			ReferencesInt: []EndpointName{epDamageFormula},
		},
		{
			Name:          qpnElement,
			Description:   "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for item abilities that deal delay.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for item abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for item abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for item abilities that cause stat changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for item abilities that cause modifier changes.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.itemAbilities = paramsMap
}

func (cfg *Config) initTriggerCommandsParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ForList:       false,
			ForSingle:     true,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for trigger commands that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          qpnUser,
			Description:   "Searches for trigger commands that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.triggerCommands = paramsMap
}

func (cfg *Config) initMiscAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If an misc ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ForList:       false,
			ForSingle:     true,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If an misc ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:        qpnRank,
			Description: "Searches for misc abilities with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for misc abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for misc abilities whose names appear in the help bar.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for misc abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for misc abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.miscAbilities = paramsMap
}

func (cfg *Config) initEnemyAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for enemy abilities with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for enemy abilities that can be used by the specified monster.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:        qpnTargetType,
			Description: "Searches for enemy abilities with the specified target types.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnAttackType,
			Description:   "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AttackType.lookup,
			ReferencesInt: []EndpointName{epAttackType},
		},
		{
			Name:          qpnDamageType,
			Description:   "Searches for enemy abilities that deal the specified types of damage.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageType.lookup,
			ReferencesInt: []EndpointName{epDamageType},
		},
		{
			Name:          qpnDamageFormula,
			Description:   "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:          qptEnum,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.DamageFormula.lookup,
			ReferencesInt: []EndpointName{epDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:        qpnBDL,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for enemy abilities that deal delay.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for enemy abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for enemy abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.enemyAbilities = paramsMap
}

func (cfg *Config) initOverdrivesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for overdrives with the specified ranks.",
			Type:        qptIntList,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for overdrives that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.overdrives = paramsMap
}

func (cfg *Config) initSubmenusParams() {
	params := []QueryParam{
		{
			Name:          qpnTopmenu,
			Description:   "Searches for submenus that are found within the specified topmenu.",
			Type:          qptNameId,
			ExampleVals:   []string{"main", "left"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epTopmenus},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.submenus = paramsMap
}

func (cfg *Config) initAllItemsParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only considers an item's related resources with the given availabilities when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only considers an item's related resources that can be farmed when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnType,
			Description:   "Searches for items that are of the specified item-types.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.ItemType.lookup,
			ReferencesInt: []EndpointName{epItemType},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.allItems = paramsMap
}

func (cfg *Config) initItemsParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays an item's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an item's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:        qpnHasAbility,
			Description: "Searches for items that can be used in battle.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for items that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for items that are from one of the specified item categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.ItemCategory.lookup,
			ReferencesInt: []EndpointName{epItemCategory},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.items = paramsMap
}

func (cfg *Config) initKeyItemsParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays a key-item's related resources with the given availabilities. This affects areas, treasures and quests.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for key-items with the given availabilities. The availability of a key-item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the key-item. The key-item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for key-items that are of the specified key-item categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.KeyItemCategory.lookup,
			ReferencesInt: []EndpointName{epKeyItemCategory},
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for key-items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvTreasure, qvQuest},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for key-items that can be obtained at the specified location. If combined with 'availability', this parameter searches for key-items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for key-items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for key-items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for key-items that can be obtained in the specified area. If combined with 'availability', this parameter searches for key-items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.keyItems = paramsMap
}

func (cfg *Config) initSpheresParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays a sphere's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a sphere's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:        qpnColor,
			Description: "Searches for spheres with any of the given colors.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.SphereColor.lookup,
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for spheres with the given availabilities. The availability of a sphere is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the sphere. The sphere won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for spheres that can be farmed. Is combinable with 'availability'. The availability assigned to the sphere is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for spheres that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for spheres that can be obtained at the specified location. If combined with 'availability', this parameter searches for spheres within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for spheres that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for spheres within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for spheres that can be obtained in the specified area. If combined with 'availability', this parameter searches for spheres within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.spheres = paramsMap
}

func (cfg *Config) initPrimersParams() {
	params := []QueryParam{
		{
			Name:          qpnAvailability,
			Description:   "Searches for primers with the given availabilities. The availability of a primer is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the primer. The primer won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.primers = paramsMap
}

func (cfg *Config) initMixesParams() {
	params := []QueryParam{
		{
			Name:            qpnContainsItem,
			Description:     "Modifies combinations to only display item combinations that include the specified item.",
			Type:            qptNameId,
			ExampleVals:     []string{"grenade", "power_sphere"},
			ForList:         false,
			ForSingle:       true,
			ForbiddenParams: []QueryParamName{qpnBest},
			ReferencesInt:   []EndpointName{epItems},
		},
		{
			Name:            qpnBest,
			Description:     "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:            qptBool,
			ForList:         false,
			ForSingle:       true,
			ForbiddenParams: []QueryParamName{qpnContainsItem},
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for mixes that are of the specified mix categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.MixCategory.lookup,
			ReferencesInt: []EndpointName{epMixCategory},
		},
		{
			Name:          qpnReqItem,
			Description:   "Searches for mixes that can be built with the specified item.",
			Type:          qptNameId,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnSecondItem,
			Description:    "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:           qptNameId,
			ExampleVals:    []string{"grenade", "power_sphere"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnReqItem},
			ReferencesInt:  []EndpointName{epItems},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.mixes = paramsMap
}

func (cfg *Config) initAutoAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays an auto-ability's related resources with the given availabilities. This affects shops, treasures, and monsters.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an auto-ability's related resources that can be farmed. This affects shops, treasures, and monsters.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:          qpnAvailability,
			Description:   "Searches for auto-abilities with the given availabilities. The availability of an auto-ability is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the auto-ability. The auto-ability won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for auto-abilities that can be farmed. Is combinable with 'availability'. The availability assigned to the auto-ability is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:          qpnCategory,
			Description:   "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:          qptEnumList,
			ForList:       true,
			ForSingle:     false,
			TypeLookup:    cfg.t.AutoAbilityCategory.lookup,
			ReferencesInt: []EndpointName{epAutoAbilityCategory},
		},
		{
			Name:        qpnType,
			Description: "Searches for auto-abilities that are of the specified equip type.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for auto-abilities that are dropped by the specified monster.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnMonsterItems,
			Description:   "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnShop,
			Description:   "Searches for auto-abilities that can be obtained from the specified shop.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnCharacter,
			Description:   "Restricts the search for 'availability', 'monster' and 'shop' to only include auto-abilities that can be obtained by the specified character. This includes auto-abilities with no character restriction like regular monster equipment drop slots.",
			Type:          qptNameId,
			ExampleVals:   []string{"kimahri"},
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			UsableWith:    []QueryParamName{qpnAvailability, qpnMonster, qpnShop},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for auto-abilities that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop},
		},
		{
			Name:        qpnCustomize,
			Description: "Converts the 'availability' and 'repeatable' parameters to search for auto-abilities based on their required item's availability and/or farmability.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified location. If combined with 'availability', this parameter searches for auto-abilities within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for auto-abilities within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for auto-abilities that can be obtained in the specified area. If combined with 'availability', this parameter searches for auto-abilities within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ForList:       true,
			ForSingle:     false,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, true)
	cfg.q.autoAbilities = paramsMap
}

func (cfg *Config) initEquipmentTablesParams() {
	params := []QueryParam{
		{
			Name:          qpnAutoAbilities,
			Description:   "Searches for equipment tables with all of the given auto-abilities.",
			Type:          qptIdList,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnType,
			Description: "Searches for equipment tables that are of the specified equip type.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:        qpnCelestialWeapon,
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.equipmentTables = paramsMap
}

func (cfg *Config) initEquipmentParams() {
	params := []QueryParam{
		{
			Name:        qpnTable,
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type:        qptInt,
			ForSingle:   true,
			ForList:     false,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:          qpnRelAvailability,
			Description:   "Only displays an equipment's related resources with the given availabilities. This affects treasures and shops.",
			Type:          qptEnumList,
			ForList:       false,
			ForSingle:     true,
			TypeLookup:    cfg.t.AvailabilityType.lookup,
			ReferencesInt: []EndpointName{epAvailabilityType},
		},
		{
			Name:          qpnAutoAbilities,
			Description:   "Searches for equipment with all of the given auto-abilities.",
			Type:          qptIdList,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:          qpnCharacter,
			Description:   "Searches for equipment of the specified character.",
			ExampleVals:   []string{"yuna"},
			Type:          qptNameId,
			ForList:       true,
			ForSingle:     false,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:        qpnType,
			Description: "Searches for equipment that is of the specified equip type.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:        qpnCelestialWeapon,
			Description: "Searches for the celestial weapons.",
			Type:        qptBool,
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.equipment = paramsMap
}

func (cfg *Config) initCelestialWeaponsParams() {
	params := []QueryParam{
		{
			Name:        qpnFormula,
			Description: "Searches for celestial-weapons that are of the specified celestial formula.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CelestialFormula.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.celestialWeapons = paramsMap
}

func (cfg *Config) initStatsParams() {
	params := []QueryParam{
		{
			Name:        qpnChangesOnly,
			Description: "Only includes a stat's related auto-abilities, abilities, status conditions, and properties that cause stat changes.",
			Type:        qptBool,
			ForList:     false,
			ForSingle:   true,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.stats = paramsMap
}

func (cfg *Config) initOverdriveModesParams() {
	params := []QueryParam{
		{
			Name:        qpnType,
			Description: "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:        qptEnum,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.OverdriveModeType.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.overdriveModes = paramsMap
}

func (cfg *Config) initStatusConditionsParams() {
	params := []QueryParam{
		{
			Name:            qpnInflictMin,
			Description:     "Only shows a status condition's related abilities with an infliction rate higher than or equal to the given amount. The default value is '1'. Can be combined with 'inflict_max', but can't be higher. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            qptInt,
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: "infinite",
					Val: 254,
				},
				{
					Key: "always",
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:            qpnInflictMax,
			Description:     "Only shows a status condition's related abilities with an infliction rate lower than or equal to the given amount. The default value is '25'. Can be combined with 'inflict_min', but can't be lower. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            qptInt,
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: "infinite",
					Val: 254,
				},
				{
					Key: "always",
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(255),
		},
		{
			Name:            qpnResistance,
			Description:     "Only shows a status condition's related monsters with a resistance higher than or equal to the given amount. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            qptInt,
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: "immune",
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:        qpnCategory,
			Description: "Searches for status conditions that are of the specified status condition categories.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.StatusConditionCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.statusConditions = paramsMap
}

func (cfg *Config) initModifiersParams() {
	params := []QueryParam{
		{
			Name:        qpnCategory,
			Description: "Searches for modifiers that are of the specified modifier categories.",
			Type:        qptEnumList,
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ModifierCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.modifiers = paramsMap
}

func (cfg *Config) initAgilityTierParams() {
	params := []QueryParam{
		{
			Name:            qpnAgility,
			Description:     "Searches for the agility tier that the given agility value belongs to.",
			Type:            qptInt,
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{0, 255},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, false)
	cfg.q.agilityTiers = paramsMap
}
