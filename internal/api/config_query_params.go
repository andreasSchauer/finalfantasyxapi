package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ParamUseType string

const (
	puSingle  ParamUseType = "single"
	puList    ParamUseType = "list"
	puSimple  ParamUseType = "simple"
	puService ParamUseType = "service"
)

type QueryParam struct {
	ID                 int                 `json:"-"`
	Name               QueryParamName      `json:"name"`
	Type               QueryParamType      `json:"param_type"`
	Description        string              `json:"description"`
	ExampleVals        []string            `json:"-"`
	Usage              string              `json:"usage"`
	ExampleUses        []string            `json:"example_uses"`
	IsExclusive        bool                `json:"only_use_alone,omitempty"`
	IsRequired         bool                `json:"is_required,omitempty"`
	ParamUse           ParamUseType        `json:"param_use_type"`
	ForSegment         *SectionName        `json:"for_segment,omitempty"`
	EnumLookup         map[string]EnumVal  `json:"-"`
	RequiredParams     []QueryParamName    `json:"required_params,omitempty"`
	UsableWith         []QueryParamName    `json:"usable_with,omitempty"`
	ReplacedBy         []QueryParamName    `json:"replaced_by,omitempty"`
	ConflictsWith      []QueryParamName    `json:"conflicts_with,omitempty"`
	ReferencesInt      []EndpointName      `json:"-"`
	ReferencesEnumsInt []EnumName          `json:"-"`
	References         []string            `json:"references,omitempty"`
	ReferencesEnums    []string            `json:"references_enums,omitempty"`
	AllowedIDs         []int32             `json:"-"`
	AllowedResources   []string            `json:"allowed_resources,omitempty"`
	AllowedValues      []QueryValue        `json:"allowed_values,omitempty"`
	AllowedIntRange    []int               `json:"allowed_int_range,omitempty"`
	AllowedResTypes    []string            `json:"allowed_res_types,omitempty"`
	DefaultVal         *int                `json:"default_value,omitempty"`
	SpecialInputs      []SpecialQueryInput `json:"special_inputs,omitempty"`
}

type SpecialQueryInput struct {
	Key QuerySpecialVal `json:"key"`
	Val int             `json:"value"`
}

func (cfg *Config) initLocationsParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays a location's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a location's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for locations with the given availabilities. Can be combined with other parameters that filter locations by resource/resource-type. In that case, this parameter searches for locations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same location, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for locations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for locations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the location. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for locations where the specified monster can be found. If combined with 'availability', the location must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for locations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for locations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the location whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for locations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for locations where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for locations where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for locations that have monsters. If combined with 'availability', the location must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for locations that have bosses. If combined with 'availability', the location must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for locations that have shops. If combined with 'availability', the location must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for locations that have treasures. If combined with 'availability', the location must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for locations that feature sidequests. If combined with 'availability', the location must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for locations that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.locations.subsections)
	cfg.q.locations = paramsMap
	cfg.e.locations.queryLookup = paramsMap
}

func (cfg *Config) initSublocationsParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays a sublocation's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a sublocation's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for sublocations that are located within the specified location.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for sublocations with the given availabilities. Can be combined with other parameters that filter sublocations by resource/resource-type. In that case, this parameter searches for sublocations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same sublocation, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the sublocation won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for sublocations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for sublocations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the sublocation. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for sublocations where the specified monster can be found. If combined with 'availability', the sublocation must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for sublocations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for sublocations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for sublocations that have monsters. If combined with 'availability', the sublocation must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for sublocations that have bosses. If combined with 'availability', the sublocation must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for sublocations that have shops. If combined with 'availability', the sublocation must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for sublocations that have treasures. If combined with 'availability', the sublocation must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for sublocations that feature sidequests. If combined with 'availability', the sublocation must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.sublocations.subsections)
	cfg.q.sublocations = paramsMap
	cfg.e.sublocations.queryLookup = paramsMap
}

func (cfg *Config) initAreasParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays an area's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an area's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for areas that are located within the specified location.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for areas that are located within the specified sublocation.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for areas with the given availabilities. Can be combined with other parameters that filter areas by resource/resource-type. In that case, this parameter searches for areas where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same area, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for areas where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for areas where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the area. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{qpnMonster, qpnItem, qpnKeyItem, qpnAutoAbility},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for areas where the specified monster can be found. If combined with 'availability', the area must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for areas where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnKeyItem,
			Description:   "Searches for areas where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the area whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for areas where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        qpnSaveSphere,
			Description: "Searches for areas that have a save sphere.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnCompSphere,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnAirship,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnChocobo,
			Description: "Searches for areas where you can ride a chocobo.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnCharacters,
			Description: "Searches for areas where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnAeons,
			Description: "Searches for areas where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnMonsters,
			Description: "Searches for areas that have monsters. If combined with 'availability', the area must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities (based on its encounter areas).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnBossFights,
			Description: "Searches for areas that have bosses. If combined with 'availability', the area must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation's encounter areas).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnShops,
			Description: "Searches for areas that have shops. If combined with 'availability', the area must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnTreasures,
			Description: "Searches for areas that have treasures. If combined with 'availability', the area must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnSidequests,
			Description: "Searchces for areas that feature sidequests. If combined with 'availability', the area must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for areas that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.areas.subsections)
	cfg.q.areas = paramsMap
	cfg.e.areas.queryLookup = paramsMap
}

func (cfg *Config) initMonsterFormationsParams() {
	params := []QueryParam{
		{
			Name:          qpnMonster,
			Description:   "Searches for monster-formations that feature the specified monster.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for monster-formations with the specified monster-formation-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.MonsterFormationCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameMonsterFormationCategory},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for monster-formations with the given availabilities. If combined with the 'area' parameter, the availability of the monster-formation in this specific area is used. If a monster-formation has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster-formation won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with the 'area' parameter, it takes the repeatability directly from the monster-formations that occur in the specified area. Is combinable with 'availability'. In that case, the search looks for the monster-formation that is the most accessible while also being farmable and checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for monster-formations that appear within the specified location. If combined with 'availability', this parameter searches for monster-formations within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this location whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for monster-formations that appear within the specified sublocation. If combined with 'availability', this parameter searches for monster-formations within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for monster-formations that appear within the specified area. If combined with 'availability', this parameter searches for monster-formations within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monster-formations within this area whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        qpnAmbush,
			Description: "Searches for monster-formations that are forced ambushes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnEscape,
			Description: "Searches for monster-formations that can be escaped.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.monsterFormations.subsections)
	cfg.q.monsterFormations = paramsMap
	cfg.e.monsterFormations.queryLookup = paramsMap
}

func (cfg *Config) initShopsParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for shops with the specified shop categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ShopCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameShopCategory},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for shops with the given availabilities. By default, this parameter checks, if a shop simply is available at the given availability. If combined with a filter that refers to the shop's inventory, it takes the availability directly from there ('pre-story' for the inventory before acquiring the airship, and 'post' after). In that case, this filter looks, if the requested resources are available in a shop at that point in the game.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for shops that appear at the specified location. If combined with 'availability', this parameter searches for shops within this location whose availability matches one of the specified availabilities.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnAvailability},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for shops that appear at the specified sublocation. If combined with 'availability', this parameter searches for shops within this sublocation whose availability matches one of the specified availabilities.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for shops that offer equipment with the specified auto-ability. Can be combined with 'empty_slots' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptId,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:    []QueryParamName{qpnAvailability},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots. Can be combined with 'auto_ability' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptIntList,
			ParamUse:        puList,
			ConflictsWith:   []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:      []QueryParamName{qpnAvailability},
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          qpnCharacter,
			Description:   "Searches for shops that offer equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{qpnItems, qpnEquipment},
			ReplacedBy:    []QueryParamName{qpnAvailability},
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnItems,
			Description:   "Searches for shops that offer items. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptBool,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{qpnAutoAbility, qpnCharacter, qpnEmptySlots},
			ReplacedBy:    []QueryParamName{qpnAvailability},
		},
		{
			Name:          qpnEquipment,
			Description:   "Searches for shops that offer equipment. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptBool,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{qpnAutoAbility, qpnCharacter, qpnEmptySlots},
			ReplacedBy:    []QueryParamName{qpnAvailability},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.shops.subsections)
	cfg.q.shops = paramsMap
	cfg.e.shops.queryLookup = paramsMap
}

func (cfg *Config) initTreasuresParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for treasures that appear within the specified location.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for treasures that appear within the specified sublocation.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for treasures that appear within the specified area.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for treasures that contain equipment with the specified auto-ability. Can be combined with 'empty_slots', 'character', and 'availability' for more specific searches.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for treasures that contain equipment with the specified amounts of empty slots. Can be combined with 'auto_ability', 'character', and 'availability' for more specific searches.",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          qpnCharacter,
			Description:   "Searches for treasures that contain equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnItem,
			Description:   "Searches for treasures that contain the specified item.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:               qpnLootType,
			Description:        "Searches for treasures with the specified loot type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.LootType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameLootType},
		},
		{
			Name:               qpnTreasureType,
			Description:        "Searches for treasures with the specified treasure type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TreasureType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTreasureType},
		},
		{
			Name:        qpnAnima,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for treasures with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.treasures.subsections)
	cfg.q.treasures = paramsMap
	cfg.e.treasures.queryLookup = paramsMap
}

func (cfg *Config) initQuestsParams() {
	params := []QueryParam{
		{
			Name:               qpnType,
			Description:        "Searches for quests that are of the specified quest type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.QuestType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameQuestType},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for quests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for quests that can be completed more than once.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.quests.subsections)
	cfg.q.quests = paramsMap
	cfg.e.quests.queryLookup = paramsMap
}

func (cfg *Config) initSidequestsParams() {
	params := []QueryParam{
		{
			Name:               qpnAvailability,
			Description:        "Searches for sidequests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.sidequests.subsections)
	cfg.q.sidequests = paramsMap
	cfg.e.sidequests.queryLookup = paramsMap
}

func (cfg *Config) initSubquestsParams() {
	params := []QueryParam{
		{
			Name:               qpnAvailability,
			Description:        "Searches for subquests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for subquests that can be completed more than once.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.subquests.subsections)
	cfg.q.subquests = paramsMap
	cfg.e.subquests.queryLookup = paramsMap
}

func (cfg *Config) initArenaCreationsParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for monster-formations with the specified arena-creation-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ArenaCreationCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameArenaCreationCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.arenaCreations.subsections)
	cfg.q.arenaCreations = paramsMap
	cfg.e.arenaCreations.queryLookup = paramsMap
}

func (cfg *Config) initBlitzballPrizesParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.BlitzballTournamentCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameBlitzballTournamentCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.blitzballPrizes.subsections)
	cfg.q.blitzballPrizes = paramsMap
	cfg.e.blitzballPrizes.queryLookup = paramsMap
}

func (cfg *Config) initSongsParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        qpnFMVs,
			Description: "Searches for songs that are played in fmvs.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnSpecialUse,
			Description: "Searches for songs with a special use case.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnComposer,
			Description:        "Searches for songs that were composed by the stated composer.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.Composer.lookup,
			ReferencesEnumsInt: []EnumName{enumNameComposer},
		},
		{
			Name:               qpnArranger,
			Description:        "Searches for songs that were arranged by the stated arranger.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.Arranger.lookup,
			ReferencesEnumsInt: []EnumName{enumNameArranger},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.songs.subsections)
	cfg.q.songs = paramsMap
	cfg.e.songs.queryLookup = paramsMap
}

func (cfg *Config) initFMVsParams() {
	params := []QueryParam{
		{
			Name:          qpnLocation,
			Description:   "Searches for fmvs that are played within the specified location.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epLocations},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.fmvs.subsections)
	cfg.q.fmvs = paramsMap
	cfg.e.fmvs.queryLookup = paramsMap
}

func (cfg *Config) initPlayerUnitsParams() {
	params := []QueryParam{
		{
			Name:               qpnType,
			Description:        "Searches for player units that are of the specified unit type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.UnitType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameUnitType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.playerUnits.subsections)
	cfg.q.playerUnits = paramsMap
	cfg.e.playerUnits.queryLookup = paramsMap
}

func (cfg *Config) initCharactersParams() {
	params := []QueryParam{
		{
			Name:               qpnOsgStats,
			Description:        "Adds all stat gains within the character's stated sphere grid to their base stats. This includes stat nodes behind sphere locks.",
			Type:               qptEnum,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.SphereGridType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameSphereGridType},
		},
		{
			Name:        qpnStoryBased,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnUnderwater,
			Description: "Searches for characters that can fight underwater.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.characters.subsections)
	cfg.q.characters = paramsMap
	cfg.e.characters.queryLookup = paramsMap
}

func (cfg *Config) initAeonsParams() {
	params := []QueryParam{
		{
			Name:            qpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:        qpnYunaStats,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used instead. Every stat instead of luck is available, since an aeon simply copies Yuna's luck stat. Can be used in combination with the 'battles' parameter.",
			Type:        qptStat,
			ExampleUses: []string{"?yuna_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?yuna_stats=accuracy=150,magic_defense=255"},
			ParamUse:    puSingle,
		},
		{
			Name:        qpnOptional,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.aeons.subsections)
	cfg.q.aeons = paramsMap
	cfg.e.aeons.queryLookup = paramsMap
}

func (cfg *Config) initCharacterClassesParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for character classes with the specified categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.CharacterClassCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameCharacterClassCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.characterClasses.subsections)
	cfg.q.characterClasses = paramsMap
	cfg.e.characterClasses.queryLookup = paramsMap
}

func (cfg *Config) initMonstersParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays a monster's related resources with the given availabilities. This affects areas and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a monster's related resources that can be farmed. This affects areas and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        qpnKimahriStats,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type:        qptStat,
			ExampleUses: []string{"?kimahri_stats=hp=3000,strength=25,magic=30,agility=40", "?kimahri_stats=hp=15000,agility=255"},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:        qpnAeonStats,
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Type:        qptStat,
			ExampleUses: []string{"?aeon_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?aeon_stats=accuracy=150,magic_defense=255"},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        qpnAlteredState,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:        qptId,
			ParamUse:    puSingle,
		},
		{
			Name:          qpnOmnisElements,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type:          qptOther,
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wilf"},
			ParamUse:      puSingle,
			AllowedIDs:    []int32{211},
			AllowedValues: []QueryValue{qvF, qvL, qvW, qvI},
		},
		{
			Name:               qpnElementalResists,
			Description:        "Searches for monsters that have the specified elemental affinities.",
			Type:               qptOther,
			Usage:              "?elemental_resists={element|id}={affinity|id},...",
			ExampleUses:        []string{"?elemental_resists=fire=weak,water=absorb", "?elemental_resists=1=3,2=4"},
			ParamUse:           puList,
			ReferencesInt:      []EndpointName{epElements},
			ReferencesEnumsInt: []EnumName{enumNameElementalAffinity},
		},
		{
			Name:          qpnStatusResists,
			Description:   "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type:          qptIdList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:            qpnResistance,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            qptInt,
			ParamUse:        puList,
			RequiredParams:  []QueryParamName{qpnStatusResists},
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: qsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:          qpnItem,
			Description:   "Searches for monsters that have the specified item as loot. Can be specified further with the 'methods' parameter.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnItem},
			AllowedValues:  []QueryValue{qvSteal, qvDrop, qvBribe, qvOther},
		},
		{
			Name:          qpnAutoAbility,
			Description:   "Searches for monsters that drop the specified auto-ability.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:           qpnIsForced,
			Description:    "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAutoAbility},
		},
		{
			Name:            qpnEmptySlots,
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:          qpnRonsoRage,
			Description:   "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epRonsoRages},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for monsters with the given availabilities. If combined with a geographical filter, it takes the availability directly from the monster-formations that occur in the specified location, sublocation, or area. If a monster has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with a geographical filter, it takes the repeatability directly from the monster-formations that occur in the specified location, sublocation, or area. Is combinable with 'availability'. The availability assigned to the monster is from the monster-formation that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for monsters that appear within the specified location. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this location whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for monsters that appear within the specified sublocation. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for monsters that appear within the specified area. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this area whose farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:            qpnDistance,
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:        qpnCapture,
			Description: "Searches for monsters that can be captured.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnHasOverdrive,
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnUnderwater,
			Description: "Searches for monsters that are fought underwater.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnZombie,
			Description: "Searches for monsters that are zombies.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnSpecies,
			Description:        "Searches for monsters of the specified species.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.MonsterSpecies.lookup,
			ReferencesEnumsInt: []EnumName{enumNameMonsterSpecies},
		},
		{
			Name:               qpnCreationArea,
			Description:        "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.CreationArea.lookup,
			ReferencesEnumsInt: []EnumName{enumNameCreationArea},
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for monsters that are of the specified monster-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.MonsterCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameMonsterCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.monsters.subsections)
	cfg.q.monsters = paramsMap
	cfg.e.monsters.queryLookup = paramsMap
}

func (cfg *Config) initAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               qpnType,
			Description:        "Searches for abilities that are of the specified ability types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AbilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAbilityType},
		},
		{
			Name:        qpnRank,
			Description: "Searches for abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for abilities that can be used by the specified monster.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               qpnTargetType,
			Description:        "Searches for abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TargetType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnAttackType,
			Description:        "Searches for abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AttackType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               qpnDamageType,
			Description:        "Searches for abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               qpnDamageFormula,
			Description:        "Searches for abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnBDL,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.abilities.subsections)
	cfg.q.abilities = paramsMap
	cfg.e.abilities.queryLookup = paramsMap
}

func (cfg *Config) initPlayerAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the player ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:        qpnRank,
			Description: "Searches for player abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for player abilities that are of the specified player ability categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.PlayerAbilityCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNamePlayerAbilityCategory},
		},
		{
			Name:        qpnOutsideBattle,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnMp,
			Description: "Searches for player abilities with the specified mp costs.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        qpnMpMin,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type:        qptInt,
			ParamUse:    puList,
		},
		{
			Name:        qpnMpMax,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type:        qptInt,
			ParamUse:    puList,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for player abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          qpnUser,
			Description:   "Searches for player abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          qpnStdSg,
			Description:   "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:          qptNameId,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnExpSg,
			Description:   "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:          qptNameId,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          qpnLearnItem,
			Description:   "Searches for player abilities an aeon can learn via the specified item.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:               qpnTargetType,
			Description:        "Searches for player abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TargetType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnAttackType,
			Description:        "Searches for player abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AttackType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               qpnDamageType,
			Description:        "Searches for player abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               qpnDamageFormula,
			Description:        "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:          qpnElement,
			Description:   "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for player abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for player abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for player abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for player abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for player abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.playerAbilities.subsections)
	cfg.q.playerAbilities = paramsMap
	cfg.e.playerAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for overdrive abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for overdrive abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               qpnTargetType,
			Description:        "Searches for overdrive abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TargetType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:               qpnAttackType,
			Description:        "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AttackType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               qpnDamageFormula,
			Description:        "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for overdrive abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for overdrive abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.overdriveAbilities.subsections)
	cfg.q.overdriveAbilities = paramsMap
	cfg.e.overdriveAbilities.queryLookup = paramsMap
}

func (cfg *Config) initItemAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for item abilities that are of the specified item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ItemCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameItemCategory},
		},
		{
			Name:        qpnOutsideBattle,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for item abilities that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               qpnTargetType,
			Description:        "Searches for item abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TargetType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:               qpnAttackType,
			Description:        "Searches for item abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AttackType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               qpnDamageFormula,
			Description:        "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:          qpnElement,
			Description:   "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for item abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for item abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for item abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        qpnStatChanges,
			Description: "Searches for item abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnModChanges,
			Description: "Searches for item abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.itemAbilities.subsections)
	cfg.q.itemAbilities = paramsMap
	cfg.e.itemAbilities.queryLookup = paramsMap
}

func (cfg *Config) initTriggerCommandsParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the trigger command's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for trigger commands that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          qpnUser,
			Description:   "Searches for trigger commands that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.triggerCommands.subsections)
	cfg.q.triggerCommands = paramsMap
	cfg.e.triggerCommands.queryLookup = paramsMap
}

func (cfg *Config) initMiscAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          qpnAbilityUser,
			Description:   "If a misc ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameId,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           qpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the misc ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:           qpnBombWpn,
			Description:    "If a misc ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{qpnAbilityUser},
		},
		{
			Name:        qpnRank,
			Description: "Searches for misc abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        qpnCopycat,
			Description: "Searches for misc abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for misc abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for misc abilities that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:        qpnUserAtk,
			Description: "Searches for misc abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.miscAbilities.subsections)
	cfg.q.miscAbilities = paramsMap
	cfg.e.miscAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEnemyAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for enemy abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        qpnHelpBar,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for enemy abilities that can be used by the specified monster.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               qpnTargetType,
			Description:        "Searches for enemy abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.TargetType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        qpnDarkable,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnSilenceable,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnReflectable,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnAttackType,
			Description:        "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AttackType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               qpnDamageType,
			Description:        "Searches for enemy abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               qpnDamageFormula,
			Description:        "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.DamageFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        qpnCanCrit,
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        qpnBDL,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnElement,
			Description:   "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIdListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        qpnDelay,
			Description: "Searches for enemy abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnStatusInflict,
			Description:   "Searches for enemy abilities that can inflict the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          qpnStatusRemove,
			Description:   "Searches for enemy abilities that can remove the specified status condition.",
			Type:          qptIdNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.enemyAbilities.subsections)
	cfg.q.enemyAbilities = paramsMap
	cfg.e.enemyAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdrivesParams() {
	params := []QueryParam{
		{
			Name:        qpnRank,
			Description: "Searches for overdrives with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:          qpnUser,
			Description:   "Searches for overdrives that are learned by the specified character class.",
			Type:          qptNameId,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.overdrives.subsections)
	cfg.q.overdrives = paramsMap
	cfg.e.overdrives.queryLookup = paramsMap
}

func (cfg *Config) initSubmenusParams() {
	params := []QueryParam{
		{
			Name:          qpnTopmenu,
			Description:   "Searches for submenus that are found within the specified topmenu.",
			Type:          qptNameId,
			ExampleVals:   []string{"main", "left"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epTopmenus},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.submenus.subsections)
	cfg.q.submenus = paramsMap
	cfg.e.submenus.queryLookup = paramsMap
}

func (cfg *Config) initAllItemsParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only considers an item's related resources with the given availabilities when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only considers an item's related resources that can be farmed when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               qpnType,
			Description:        "Searches for items that are of the specified item-types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ItemType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameItemType},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.allItems.subsections)
	cfg.q.allItems = paramsMap
	cfg.e.allItems.queryLookup = paramsMap
}

func (cfg *Config) initItemsParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays an item's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an item's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        qpnHasAbility,
			Description: "Searches for items that can be used in battle.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnRelatedStat,
			Description:   "Searches for items that are related to the specified stat.",
			Type:          qptNameId,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for items that are from one of the specified item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ItemCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameItemCategory},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.items.subsections)
	cfg.q.items = paramsMap
	cfg.e.items.queryLookup = paramsMap
}

func (cfg *Config) initKeyItemsParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays a key-item's related resources with the given availabilities. This affects areas, treasures and quests.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for key-items with the given availabilities. The availability of a key-item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the key-item. The key-item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for key-items that are of the specified key-item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.KeyItemCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameKeyItemCategory},
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for key-items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvTreasure, qvQuest},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for key-items that can be obtained at the specified location. If combined with 'availability', this parameter searches for key-items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for key-items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for key-items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for key-items that can be obtained in the specified area. If combined with 'availability', this parameter searches for key-items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.keyItems.subsections)
	cfg.q.keyItems = paramsMap
	cfg.e.keyItems.queryLookup = paramsMap
}

func (cfg *Config) initSpheresParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays a sphere's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays a sphere's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               qpnColor,
			Description:        "Searches for spheres with any of the given colors.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.SphereColor.lookup,
			ReferencesEnumsInt: []EnumName{enumNameSphereColor},
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for spheres with the given availabilities. The availability of a sphere is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the sphere. The sphere won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for spheres that can be farmed. Is combinable with 'availability'. The availability assigned to the sphere is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for spheres that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for spheres that can be obtained at the specified location. If combined with 'availability', this parameter searches for spheres within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for spheres that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for spheres within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for spheres that can be obtained in the specified area. If combined with 'availability', this parameter searches for spheres within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.spheres.subsections)
	cfg.q.spheres = paramsMap
	cfg.e.spheres.queryLookup = paramsMap
}

func (cfg *Config) initPrimersParams() {
	params := []QueryParam{
		{
			Name:               qpnAvailability,
			Description:        "Searches for primers with the given availabilities. The availability of a primer is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the primer. The primer won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.primers.subsections)
	cfg.q.primers = paramsMap
	cfg.e.primers.queryLookup = paramsMap
}

func (cfg *Config) initMixesParams() {
	params := []QueryParam{
		{
			Name:          qpnContainsItem,
			Description:   "Modifies combinations to only display item combinations that include the specified item.",
			Type:          qptNameId,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puSingle,
			ConflictsWith: []QueryParamName{qpnBest},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:          qpnBest,
			Description:   "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:          qptBool,
			ParamUse:      puSingle,
			ConflictsWith: []QueryParamName{qpnContainsItem},
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for mixes that are of the specified mix categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.MixCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameMixCategory},
		},
		{
			Name:          qpnReqItem,
			Description:   "Searches for mixes that can be built with the specified item.",
			Type:          qptNameId,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           qpnSecondItem,
			Description:    "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:           qptNameId,
			ExampleVals:    []string{"grenade", "power_sphere"},
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnReqItem},
			ReferencesInt:  []EndpointName{epItems},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.mixes.subsections)
	cfg.q.mixes = paramsMap
	cfg.e.mixes.queryLookup = paramsMap
}

func (cfg *Config) initAutoAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays an auto-ability's related resources with the given availabilities. This affects shops, treasures, and monsters.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        qpnRelRepeatable,
			Description: "Only displays an auto-ability's related resources that can be farmed. This affects shops, treasures, and monsters.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               qpnAvailability,
			Description:        "Searches for auto-abilities with the given availabilities. The availability of an auto-ability is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the auto-ability. The auto-ability won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           qpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{qpnAvailability},
		},
		{
			Name:        qpnRepeatable,
			Description: "Searches for auto-abilities that can be farmed. Is combinable with 'availability'. The availability assigned to the auto-ability is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.AutoAbilityCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAutoAbilityCategory},
		},
		{
			Name:               qpnType,
			Description:        "Searches for auto-abilities that are of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.EquipType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:          qpnMonster,
			Description:   "Searches for auto-abilities that are dropped by the specified monster.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnMonsterItems,
			Description:   "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnShop,
			Description:   "Searches for auto-abilities that can be obtained from the specified shop.",
			Type:          qptId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnCharacter,
			Description:   "Restricts the search for 'availability', 'monster' and 'shop' to only include auto-abilities that can be obtained by the specified character. This includes auto-abilities with no character restriction like regular monster equipment drop slots.",
			Type:          qptNameId,
			ExampleVals:   []string{"kimahri"},
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			UsableWith:    []QueryParamName{qpnAvailability, qpnMonster, qpnShop},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          qpnMethods,
			Description:   "Searches for auto-abilities that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop},
		},
		{
			Name:        qpnCustomize,
			Description: "Converts the 'availability' and 'repeatable' parameters to search for auto-abilities based on their required item's availability and/or farmability.",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{qpnAvailability, qpnRepeatable},
		},
		{
			Name:          qpnLocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified location. If combined with 'availability', this parameter searches for auto-abilities within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnSublocation, qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          qpnSublocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for auto-abilities within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnArea, qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          qpnArea,
			Description:   "Searches for auto-abilities that can be obtained in the specified area. If combined with 'availability', this parameter searches for auto-abilities within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptId,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{qpnAvailability, qpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.autoAbilities.subsections)
	cfg.q.autoAbilities = paramsMap
	cfg.e.autoAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentTablesParams() {
	params := []QueryParam{
		{
			Name:          qpnAutoAbilities,
			Description:   "Searches for equipment tables with all of the given auto-abilities.",
			Type:          qptIdList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:               qpnType,
			Description:        "Searches for equipment tables that are of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.EquipType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:        qpnCelestialWeapon,
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.equipmentTables.subsections)
	cfg.q.equipmentTables = paramsMap
	cfg.e.equipmentTables.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentParams() {
	params := []QueryParam{
		{
			Name:        qpnTable,
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type:        qptInt,
			ParamUse:    puSingle,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:               qpnRelAvailability,
			Description:        "Only displays an equipment's related resources with the given availabilities. This affects treasures and shops.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumLookup:         cfg.t.AvailabilityType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:          qpnAutoAbilities,
			Description:   "Searches for equipment with all of the given auto-abilities.",
			Type:          qptIdList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:          qpnCharacter,
			Description:   "Searches for equipment of the specified character.",
			ExampleVals:   []string{"yuna"},
			Type:          qptNameId,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:               qpnType,
			Description:        "Searches for equipment that is of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.EquipType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:        qpnCelestialWeapon,
			Description: "Searches for the celestial weapons.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.equipment.subsections)
	cfg.q.equipment = paramsMap
	cfg.e.equipment.queryLookup = paramsMap
}

func (cfg *Config) initCelestialWeaponsParams() {
	params := []QueryParam{
		{
			Name:               qpnFormula,
			Description:        "Searches for celestial-weapons that are of the specified celestial formula.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.CelestialFormula.lookup,
			ReferencesEnumsInt: []EnumName{enumNameCelestialFormula},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.celestialWeapons.subsections)
	cfg.q.celestialWeapons = paramsMap
	cfg.e.celestialWeapons.queryLookup = paramsMap
}

func (cfg *Config) initStatsParams() {
	params := []QueryParam{
		{
			Name:        qpnChangesOnly,
			Description: "Only includes a stat's related auto-abilities, abilities, status conditions, and properties that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.stats.subsections)
	cfg.q.stats = paramsMap
	cfg.e.stats.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveModesParams() {
	params := []QueryParam{
		{
			Name:               qpnType,
			Description:        "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumLookup:         cfg.t.OverdriveModeType.lookup,
			ReferencesEnumsInt: []EnumName{enumNameOverdriveModeType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.overdriveModes.subsections)
	cfg.q.overdriveModes = paramsMap
	cfg.e.overdriveModes.queryLookup = paramsMap
}

func (cfg *Config) initStatusConditionsParams() {
	params := []QueryParam{
		{
			Name:            qpnInflictMin,
			Description:     "Only shows a status condition's related abilities with an infliction rate higher than or equal to the given amount. The default value is '1'. Can be combined with 'inflict_max', but can't be higher. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: qsvInfinite,
					Val: 254,
				},
				{
					Key: qsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:            qpnInflictMax,
			Description:     "Only shows a status condition's related abilities with an infliction rate lower than or equal to the given amount. The default value is '25'. Can be combined with 'inflict_min', but can't be lower. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: qsvInfinite,
					Val: 254,
				},
				{
					Key: qsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(255),
		},
		{
			Name:            qpnResistance,
			Description:     "Only shows a status condition's related monsters with a resistance higher than or equal to the given amount. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: qsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:               qpnCategory,
			Description:        "Searches for status conditions that are of the specified status condition categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.StatusConditionCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameStatusConditionCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.statusConditions.subsections)
	cfg.q.statusConditions = paramsMap
	cfg.e.statusConditions.queryLookup = paramsMap
}

func (cfg *Config) initModifiersParams() {
	params := []QueryParam{
		{
			Name:               qpnCategory,
			Description:        "Searches for modifiers that are of the specified modifier categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumLookup:         cfg.t.ModifierCategory.lookup,
			ReferencesEnumsInt: []EnumName{enumNameModifierCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.modifiers.subsections)
	cfg.q.modifiers = paramsMap
	cfg.e.modifiers.queryLookup = paramsMap
}

func (cfg *Config) initAgilityTierParams() {
	params := []QueryParam{
		{
			Name:            qpnAgility,
			Description:     "Searches for the agility tier that the given agility value belongs to.",
			Type:            qptInt,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 255},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.e.agilityTiers.subsections)
	cfg.q.agilityTiers = paramsMap
	cfg.e.agilityTiers.queryLookup = paramsMap
}

func (cfg *Config) initAlBhedParams() {
	exampleUses := []string{}
	paramsMap := cfg.initComputeEndpointQueryParams(epAlBhed, exampleUses)
	cfg.q.alBhed = paramsMap
	cfg.e.alBhed.queryLookup = paramsMap
}
