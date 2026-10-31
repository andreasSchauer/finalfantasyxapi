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
	ParamUse           ParamUseType        `json:"param_use_type"`
	ForSegment         *SectionName        `json:"for_segment,omitempty"`
	EnumVals           []EnumVal           `json:"-"`
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
			Name:               QpnRelAvailability,
			Description:        "Only displays a location's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays a location's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for locations with the given availabilities. Can be combined with other parameters that filter locations by resource/resource-type. In that case, this parameter searches for locations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same location, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for locations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for locations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the location. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{QpnMonster, QpnItem, QpnKeyItem, QpnAutoAbility},
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for locations where the specified monster can be found. If combined with 'availability', the location must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnItem,
			Description:   "Searches for locations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           QpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnKeyItem,
			Description:   "Searches for locations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the location whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for locations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        QpnCharacters,
			Description: "Searches for locations where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnAeons,
			Description: "Searches for locations where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnMonsters,
			Description: "Searches for locations that have monsters. If combined with 'availability', the location must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnBossFights,
			Description: "Searches for locations that have bosses. If combined with 'availability', the location must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnShops,
			Description: "Searches for locations that have shops. If combined with 'availability', the location must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnTreasures,
			Description: "Searches for locations that have treasures. If combined with 'availability', the location must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnSidequests,
			Description: "Searchces for locations that feature sidequests. If combined with 'availability', the location must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnFMVs,
			Description: "Searches for locations that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Locations.subsections)
	cfg.Q.locations = paramsMap
	cfg.E.Locations.queryLookup = paramsMap
}

func (cfg *Config) initSublocationsParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays a sublocation's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays a sublocation's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for sublocations that are located within the specified location.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for sublocations with the given availabilities. Can be combined with other parameters that filter sublocations by resource/resource-type. In that case, this parameter searches for sublocations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same sublocation, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the sublocation won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for sublocations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for sublocations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the sublocation. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{QpnMonster, QpnItem, QpnKeyItem, QpnAutoAbility},
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for sublocations where the specified monster can be found. If combined with 'availability', the sublocation must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnItem,
			Description:   "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           QpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnKeyItem,
			Description:   "Searches for sublocations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for sublocations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        QpnCharacters,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnAeons,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnMonsters,
			Description: "Searches for sublocations that have monsters. If combined with 'availability', the sublocation must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnBossFights,
			Description: "Searches for sublocations that have bosses. If combined with 'availability', the sublocation must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnShops,
			Description: "Searches for sublocations that have shops. If combined with 'availability', the sublocation must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnTreasures,
			Description: "Searches for sublocations that have treasures. If combined with 'availability', the sublocation must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnSidequests,
			Description: "Searchces for sublocations that feature sidequests. If combined with 'availability', the sublocation must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnFMVs,
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Sublocations.subsections)
	cfg.Q.sublocations = paramsMap
	cfg.E.Sublocations.queryLookup = paramsMap
}

func (cfg *Config) initAreasParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays an area's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays an area's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for areas that are located within the specified location.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for areas that are located within the specified sublocation.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for areas with the given availabilities. Can be combined with other parameters that filter areas by resource/resource-type. In that case, this parameter searches for areas where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same area, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for areas where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for areas where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the area. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{QpnMonster, QpnItem, QpnKeyItem, QpnAutoAbility},
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for areas where the specified monster can be found. If combined with 'availability', the area must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnItem,
			Description:   "Searches for areas where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           QpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnItem},
			AllowedValues:  []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnKeyItem,
			Description:   "Searches for areas where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the area whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epKeyItems},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for areas where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:        QpnSaveSphere,
			Description: "Searches for areas that have a save sphere.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnCompSphere,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnAirship,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnChocobo,
			Description: "Searches for areas where you can ride a chocobo.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnCharacters,
			Description: "Searches for areas where a character permanently joins the party.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnAeons,
			Description: "Searches for areas where a new aeon is acquired.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnMonsters,
			Description: "Searches for areas that have monsters. If combined with 'availability', the area must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities (based on its encounter areas).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnBossFights,
			Description: "Searches for areas that have bosses. If combined with 'availability', the area must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation's encounter areas).",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnShops,
			Description: "Searches for areas that have shops. If combined with 'availability', the area must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnTreasures,
			Description: "Searches for areas that have treasures. If combined with 'availability', the area must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnSidequests,
			Description: "Searchces for areas that feature sidequests. If combined with 'availability', the area must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        qptBool,
			ParamUse:    puList,
			ReplacedBy:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:        QpnFMVs,
			Description: "Searches for areas that feature fmv sequences.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Areas.subsections)
	cfg.Q.areas = paramsMap
	cfg.E.Areas.queryLookup = paramsMap
}

func (cfg *Config) initMonsterFormationsParams() {
	params := []QueryParam{
		{
			Name:          QpnMonster,
			Description:   "Searches for monster-formations that feature the specified monster.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for monster-formations with the specified monster-formation-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.MonsterFormationCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameMonsterFormationCategory},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for monster-formations with the given availabilities. If combined with the 'area' parameter, the availability of the monster-formation in this specific area is used. If a monster-formation has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster-formation won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with the 'area' parameter, it takes the repeatability directly from the monster-formations that occur in the specified area. Is combinable with 'availability'. In that case, the search looks for the monster-formation that is the most accessible while also being farmable and checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for monster-formations that appear within the specified location. If combined with 'availability', this parameter searches for monster-formations within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this location whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for monster-formations that appear within the specified sublocation. If combined with 'availability', this parameter searches for monster-formations within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for monster-formations that appear within the specified area. If combined with 'availability', this parameter searches for monster-formations within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monster-formations within this area whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        QpnAmbush,
			Description: "Searches for monster-formations that are forced ambushes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnEscape,
			Description: "Searches for monster-formations that can be escaped.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.MonsterFormations.subsections)
	cfg.Q.monsterFormations = paramsMap
	cfg.E.MonsterFormations.queryLookup = paramsMap
}

func (cfg *Config) initShopsParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for shops with the specified shop categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ShopCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameShopCategory},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for shops with the given availabilities. By default, this parameter checks, if a shop simply is available at the given availability. If combined with a filter that refers to the shop's inventory, it takes the availability directly from there ('pre-story' for the inventory before acquiring the airship, and 'post' after). In that case, this filter looks, if the requested resources are available in a shop at that point in the game.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for shops that appear at the specified location. If combined with 'availability', this parameter searches for shops within this location whose availability matches one of the specified availabilities.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnAvailability},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for shops that appear at the specified sublocation. If combined with 'availability', this parameter searches for shops within this sublocation whose availability matches one of the specified availabilities.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for shops that offer equipment with the specified auto-ability. Can be combined with 'empty_slots' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptID,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{QpnItems, QpnEquipment},
			ReplacedBy:    []QueryParamName{QpnAvailability},
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:            QpnEmptySlots,
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots. Can be combined with 'auto_ability' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            qptIntList,
			ParamUse:        puList,
			ConflictsWith:   []QueryParamName{QpnItems, QpnEquipment},
			ReplacedBy:      []QueryParamName{QpnAvailability},
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          QpnCharacter,
			Description:   "Searches for shops that offer equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{QpnItems, QpnEquipment},
			ReplacedBy:    []QueryParamName{QpnAvailability},
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          QpnItems,
			Description:   "Searches for shops that offer items. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptBool,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{QpnAutoAbility, QpnCharacter, QpnEmptySlots},
			ReplacedBy:    []QueryParamName{QpnAvailability},
		},
		{
			Name:          QpnEquipment,
			Description:   "Searches for shops that offer equipment. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          qptBool,
			ParamUse:      puList,
			ConflictsWith: []QueryParamName{QpnAutoAbility, QpnCharacter, QpnEmptySlots},
			ReplacedBy:    []QueryParamName{QpnAvailability},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Shops.subsections)
	cfg.Q.shops = paramsMap
	cfg.E.Shops.queryLookup = paramsMap
}

func (cfg *Config) initTreasuresParams() {
	params := []QueryParam{
		{
			Name:          QpnLocation,
			Description:   "Searches for treasures that appear within the specified location.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for treasures that appear within the specified sublocation.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for treasures that appear within the specified area.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for treasures that contain equipment with the specified auto-ability. Can be combined with 'empty_slots', 'character', and 'availability' for more specific searches.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:            QpnEmptySlots,
			Description:     "Searches for treasures that contain equipment with the specified amounts of empty slots. Can be combined with 'auto_ability', 'character', and 'availability' for more specific searches.",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          QpnCharacter,
			Description:   "Searches for treasures that contain equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches.",
			Type:          qptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          QpnItem,
			Description:   "Searches for treasures that contain the specified item.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:               QpnLootType,
			Description:        "Searches for treasures with the specified loot type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.LootType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameLootType},
		},
		{
			Name:               QpnTreasureType,
			Description:        "Searches for treasures with the specified treasure type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.TreasureType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTreasureType},
		},
		{
			Name:        QpnAnima,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for treasures with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Treasures.subsections)
	cfg.Q.treasures = paramsMap
	cfg.E.Treasures.queryLookup = paramsMap
}

func (cfg *Config) initQuestsParams() {
	params := []QueryParam{
		{
			Name:               QpnType,
			Description:        "Searches for quests that are of the specified quest type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.QuestType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameQuestType},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for quests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for quests that can be completed more than once.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Quests.subsections)
	cfg.Q.quests = paramsMap
	cfg.E.Quests.queryLookup = paramsMap
}

func (cfg *Config) initSidequestsParams() {
	params := []QueryParam{
		{
			Name:               QpnAvailability,
			Description:        "Searches for sidequests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Sidequests.subsections)
	cfg.Q.sidequests = paramsMap
	cfg.E.Sidequests.queryLookup = paramsMap
}

func (cfg *Config) initSubquestsParams() {
	params := []QueryParam{
		{
			Name:               QpnAvailability,
			Description:        "Searches for subquests with the given availabilities.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for subquests that can be completed more than once.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Subquests.subsections)
	cfg.Q.subquests = paramsMap
	cfg.E.Subquests.queryLookup = paramsMap
}

func (cfg *Config) initArenaCreationsParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for monster-formations with the specified arena-creation-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ArenaCreationCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameArenaCreationCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.ArenaCreations.subsections)
	cfg.Q.arenaCreations = paramsMap
	cfg.E.ArenaCreations.queryLookup = paramsMap
}

func (cfg *Config) initBlitzballPrizesParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.BlitzballTournamentCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameBlitzballTournamentCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.BlitzballPrizes.subsections)
	cfg.Q.blitzballPrizes = paramsMap
	cfg.E.BlitzballPrizes.queryLookup = paramsMap
}

func (cfg *Config) initSongsParams() {
	params := []QueryParam{
		{
			Name:          QpnLocation,
			Description:   "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:        QpnFMVs,
			Description: "Searches for songs that are played in fmvs.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnSpecialUse,
			Description: "Searches for songs with a special use case.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnComposer,
			Description:        "Searches for songs that were composed by the stated composer.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.Composer.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameComposer},
		},
		{
			Name:               QpnArranger,
			Description:        "Searches for songs that were arranged by the stated arranger.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.Arranger.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameArranger},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Songs.subsections)
	cfg.Q.songs = paramsMap
	cfg.E.Songs.queryLookup = paramsMap
}

func (cfg *Config) initFMVsParams() {
	params := []QueryParam{
		{
			Name:          QpnLocation,
			Description:   "Searches for fmvs that are played within the specified location.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epLocations},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.FMVs.subsections)
	cfg.Q.fmvs = paramsMap
	cfg.E.FMVs.queryLookup = paramsMap
}

func (cfg *Config) initPlayerUnitsParams() {
	params := []QueryParam{
		{
			Name:               QpnType,
			Description:        "Searches for player units that are of the specified unit type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.UnitType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameUnitType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.PlayerUnits.subsections)
	cfg.Q.playerUnits = paramsMap
	cfg.E.PlayerUnits.queryLookup = paramsMap
}

func (cfg *Config) initCharactersParams() {
	params := []QueryParam{
		{
			Name:               QpnOsgStats,
			Description:        "Adds all stat gains within the character's stated sphere grid to their base stats. This includes stat nodes behind sphere locks.",
			Type:               qptEnum,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.SphereGridType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameSphereGridType},
		},
		{
			Name:        QpnStoryBased,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnUnderwater,
			Description: "Searches for characters that can fight underwater.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Characters.subsections)
	cfg.Q.characters = paramsMap
	cfg.E.Characters.queryLookup = paramsMap
}

func (cfg *Config) initAeonsParams() {
	params := []QueryParam{
		{
			Name:            QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:        QpnYunaStats,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter.",
			Type:        qptStat,
			ExampleUses: []string{exampleUse(QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
		},
		{
			Name:        QpnOptional,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Aeons.subsections)
	cfg.Q.aeons = paramsMap
	cfg.E.Aeons.queryLookup = paramsMap
}

func (cfg *Config) initCharacterClassesParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for character classes with the specified categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.CharacterClassCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameCharacterClassCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.CharacterClasses.subsections)
	cfg.Q.characterClasses = paramsMap
	cfg.E.CharacterClasses.queryLookup = paramsMap
}

func (cfg *Config) initMonstersParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays a monster's related resources with the given availabilities. This affects areas and monster-formations.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays a monster's related resources that can be farmed. This affects areas and monster-formations.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        QpnKimahriStats,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type:        qptStat,
			ExampleUses: []string{exampleUse(QpnKimahriStats, "hp:3000,strength:25,magic:30,agility:40"), exampleUse(QpnKimahriStats, "hp:15000,agility:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:            QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating a possessed aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            qptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
			AllowedIDs:      []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        QpnYunaStats,
			Description: "Calculate a possessed aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter. All stats are replaceable, except for MP and Luck. While Yuna's MP and Luck still play a factor in the calculation of the other stats, a posssessed aeon's own MP and Luck will always be set to 1. If unused, the aeon's stats are based on your own aeon's default stats.",
			Type:        qptStat,
			ExampleUses: []string{exampleUse(QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        QpnAlteredState,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:        qptID,
			ParamUse:    puSingle,
		},
		{
			Name:          QpnOmnisElements,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type:          qptOther,
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{exampleUse(QpnOmnisElements, "ifil"), exampleUse(QpnOmnisElements, "wilf")},
			ParamUse:      puSingle,
			AllowedIDs:    []int32{211},
			AllowedValues: []QueryValue{qvF, qvL, qvW, qvI},
		},
		{
			Name:               QpnElementalResists,
			Description:        "Searches for monsters that have the specified elemental affinities.",
			Type:               qptOther,
			Usage:              "?elemental_resists={element|id}:{affinity|id},...",
			ExampleUses:        []string{exampleUse(QpnElementalResists, "fire:weak,water:absorb"), exampleUse(QpnElementalResists, "1:3,2:4")},
			ParamUse:           puList,
			ReferencesInt:      []EndpointName{epElements},
			ReferencesEnumsInt: []EnumName{enumNameElementalAffinity},
		},
		{
			Name:          QpnStatusResists,
			Description:   "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type:          qptIDList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:            QpnResistance,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            qptInt,
			ParamUse:        puList,
			RequiredParams:  []QueryParamName{QpnStatusResists},
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
			Name:          QpnItem,
			Description:   "Searches for monsters that have the specified item as loot. Can be specified further with the 'methods' parameter.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           QpnMethods,
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:           qptValueList,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnItem},
			AllowedValues:  []QueryValue{qvSteal, qvDrop, qvBribe, qvOther},
		},
		{
			Name:          QpnAutoAbility,
			Description:   "Searches for monsters that drop the specified auto-ability.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:           QpnIsForced,
			Description:    "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAutoAbility},
		},
		{
			Name:            QpnEmptySlots,
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:          QpnRonsoRage,
			Description:   "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epRonsoRages},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for monsters with the given availabilities. If combined with a geographical filter, it takes the availability directly from the monster-formations that occur in the specified location, sublocation, or area. If a monster has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with a geographical filter, it takes the repeatability directly from the monster-formations that occur in the specified location, sublocation, or area. Is combinable with 'availability'. The availability assigned to the monster is from the monster-formation that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for monsters that appear within the specified location. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this location whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for monsters that appear within the specified sublocation. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for monsters that appear within the specified area. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this area whose farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
		{
			Name:            QpnDistance,
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type:            qptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:        QpnCapture,
			Description: "Searches for monsters that can be captured.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnHasOverdrive,
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnUnderwater,
			Description: "Searches for monsters that are fought underwater.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnZombie,
			Description: "Searches for monsters that are zombies.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnSpecies,
			Description:        "Searches for monsters of the specified species.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.MonsterSpecies.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameMonsterSpecies},
		},
		{
			Name:               QpnCreationArea,
			Description:        "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.CreationArea.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameCreationArea},
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for monsters that are of the specified monster-categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.MonsterCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameMonsterCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Monsters.subsections)
	cfg.Q.monsters = paramsMap
	cfg.E.Monsters.queryLookup = paramsMap
}

func (cfg *Config) initAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               QpnType,
			Description:        "Searches for abilities that are of the specified ability types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AbilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAbilityType},
		},
		{
			Name:        QpnRank,
			Description: "Searches for abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        QpnCopycat,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnHelpBar,
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for abilities that can be used by the specified monster.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               QpnTargetType,
			Description:        "Searches for abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.TargetType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        QpnUserAtk,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnDarkable,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnSilenceable,
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnReflectable,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnAttackType,
			Description:        "Searches for abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AttackType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               QpnDamageType,
			Description:        "Searches for abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               QpnDamageFormula,
			Description:        "Searches for abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        QpnCanCrit,
			Description: "Searches for abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnBDL,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnElement,
			Description:   "Searches for abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        QpnDelay,
			Description: "Searches for abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnStatusInflict,
			Description:   "Searches for abilities that can inflict the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          QpnStatusRemove,
			Description:   "Searches for abilities that can remove the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        QpnStatChanges,
			Description: "Searches for abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnModChanges,
			Description: "Searches for abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Abilities.subsections)
	cfg.Q.abilities = paramsMap
	cfg.E.Abilities.queryLookup = paramsMap
}

func (cfg *Config) initPlayerAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          QpnAbilityUser,
			Description:   "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           QpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the player ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:           QpnBombWpn,
			Description:    "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:        QpnRank,
			Description: "Searches for player abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        QpnCopycat,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnHelpBar,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for player abilities that are of the specified player ability categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.PlayerAbilityCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNamePlayerAbilityCategory},
		},
		{
			Name:        QpnOutsideBattle,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnMp,
			Description: "Searches for player abilities with the specified mp costs.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        QpnMpMin,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type:        qptInt,
			ParamUse:    puList,
		},
		{
			Name:        QpnMpMax,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type:        qptInt,
			ParamUse:    puList,
		},
		{
			Name:          QpnRelatedStat,
			Description:   "Searches for player abilities that are related to the specified stat.",
			Type:          qptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          QpnUser,
			Description:   "Searches for player abilities that are learned by the specified character class.",
			Type:          qptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          QpnStdSg,
			Description:   "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:          qptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          QpnExpSg,
			Description:   "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:          qptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:          QpnLearnItem,
			Description:   "Searches for player abilities an aeon can learn via the specified item.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:               QpnTargetType,
			Description:        "Searches for player abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.TargetType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        QpnUserAtk,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnDarkable,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnSilenceable,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnReflectable,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnAttackType,
			Description:        "Searches for player abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AttackType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               QpnDamageType,
			Description:        "Searches for player abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               QpnDamageFormula,
			Description:        "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:          QpnElement,
			Description:   "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        QpnDelay,
			Description: "Searches for player abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnStatusInflict,
			Description:   "Searches for player abilities that can inflict the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          QpnStatusRemove,
			Description:   "Searches for player abilities that can remove the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        QpnStatChanges,
			Description: "Searches for player abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnModChanges,
			Description: "Searches for player abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.PlayerAbilities.subsections)
	cfg.Q.playerAbilities = paramsMap
	cfg.E.PlayerAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        QpnRank,
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:          QpnUser,
			Description:   "Searches for overdrive abilities that are learned by the specified character class.",
			Type:          qptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:          QpnRelatedStat,
			Description:   "Searches for overdrive abilities that are related to the specified stat.",
			Type:          qptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               QpnTargetType,
			Description:        "Searches for overdrive abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.TargetType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:               QpnAttackType,
			Description:        "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AttackType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               QpnDamageFormula,
			Description:        "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        QpnCanCrit,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnElement,
			Description:   "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        QpnDelay,
			Description: "Searches for overdrive abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnStatusInflict,
			Description:   "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          QpnStatusRemove,
			Description:   "Searches for overdrive abilities that can remove the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        QpnStatChanges,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnModChanges,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.OverdriveAbilities.subsections)
	cfg.Q.overdriveAbilities = paramsMap
	cfg.E.OverdriveAbilities.queryLookup = paramsMap
}

func (cfg *Config) initItemAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for item abilities that are of the specified item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameItemCategory},
		},
		{
			Name:        QpnOutsideBattle,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnRelatedStat,
			Description:   "Searches for item abilities that are related to the specified stat.",
			Type:          qptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               QpnTargetType,
			Description:        "Searches for item abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.TargetType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:               QpnAttackType,
			Description:        "Searches for item abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AttackType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               QpnDamageFormula,
			Description:        "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:          QpnElement,
			Description:   "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        QpnDelay,
			Description: "Searches for item abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnStatusInflict,
			Description:   "Searches for item abilities that can inflict the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          QpnStatusRemove,
			Description:   "Searches for item abilities that can remove the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:        QpnStatChanges,
			Description: "Searches for item abilities that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnModChanges,
			Description: "Searches for item abilities that cause modifier changes.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.ItemAbilities.subsections)
	cfg.Q.itemAbilities = paramsMap
	cfg.E.ItemAbilities.queryLookup = paramsMap
}

func (cfg *Config) initTriggerCommandsParams() {
	params := []QueryParam{
		{
			Name:          QpnAbilityUser,
			Description:   "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:          qptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           QpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the trigger command's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:           QpnBombWpn,
			Description:    "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:          QpnRelatedStat,
			Description:   "Searches for trigger commands that are related to the specified stat.",
			Type:          qptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:          QpnUser,
			Description:   "Searches for trigger commands that are learned by the specified character class.",
			Type:          qptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.TriggerCommands.subsections)
	cfg.Q.triggerCommands = paramsMap
	cfg.E.TriggerCommands.queryLookup = paramsMap
}

func (cfg *Config) initMiscAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          QpnAbilityUser,
			Description:   "If a misc ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          qptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []EndpointName{epPlayerUnits},
		},
		{
			Name:           QpnCelestialWeapon,
			Description:    "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the misc ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:           QpnBombWpn,
			Description:    "If a misc ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:           qptBool,
			ParamUse:       puSingle,
			RequiredParams: []QueryParamName{QpnAbilityUser},
		},
		{
			Name:        QpnRank,
			Description: "Searches for misc abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        QpnCopycat,
			Description: "Searches for misc abilities that can be copied by 'copycat'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnHelpBar,
			Description: "Searches for misc abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnUser,
			Description:   "Searches for misc abilities that are learned by the specified character class.",
			Type:          qptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
		{
			Name:        QpnUserAtk,
			Description: "Searches for misc abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.MiscAbilities.subsections)
	cfg.Q.miscAbilities = paramsMap
	cfg.E.MiscAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEnemyAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        QpnRank,
			Description: "Searches for enemy abilities with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:        QpnHelpBar,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for enemy abilities that can be used by the specified monster.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:               QpnTargetType,
			Description:        "Searches for enemy abilities with the specified target types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.TargetType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameTargetType},
		},
		{
			Name:        QpnDarkable,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnSilenceable,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnReflectable,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnAttackType,
			Description:        "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AttackType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAttackType},
		},
		{
			Name:               QpnDamageType,
			Description:        "Searches for enemy abilities that deal the specified types of damage.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageType},
		},
		{
			Name:               QpnDamageFormula,
			Description:        "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameDamageFormula},
		},
		{
			Name:        QpnCanCrit,
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:        QpnBDL,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnElement,
			Description:   "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:          qptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epElements},
		},
		{
			Name:        QpnDelay,
			Description: "Searches for enemy abilities that deal delay.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnStatusInflict,
			Description:   "Searches for enemy abilities that can inflict the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
		{
			Name:          QpnStatusRemove,
			Description:   "Searches for enemy abilities that can remove the specified status condition.",
			Type:          qptIDNul,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStatusConditions},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.EnemyAbilities.subsections)
	cfg.Q.enemyAbilities = paramsMap
	cfg.E.EnemyAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdrivesParams() {
	params := []QueryParam{
		{
			Name:        QpnRank,
			Description: "Searches for overdrives with the specified ranks.",
			Type:        qptIntList,
			ParamUse:    puList,
		},
		{
			Name:          QpnUser,
			Description:   "Searches for overdrives that are learned by the specified character class.",
			Type:          qptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacterClasses},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Overdrives.subsections)
	cfg.Q.overdrives = paramsMap
	cfg.E.Overdrives.queryLookup = paramsMap
}

func (cfg *Config) initSubmenusParams() {
	params := []QueryParam{
		{
			Name:          QpnTopmenu,
			Description:   "Searches for submenus that are found within the specified topmenu.",
			Type:          qptNameID,
			ExampleVals:   []string{"main", "left"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epTopmenus},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Submenus.subsections)
	cfg.Q.submenus = paramsMap
	cfg.E.Submenus.queryLookup = paramsMap
}

func (cfg *Config) initAllItemsParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only considers an item's related resources with the given availabilities when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only considers an item's related resources that can be farmed when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               QpnType,
			Description:        "Searches for items that are of the specified item-types.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ItemType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameItemType},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.AllItems.subsections)
	cfg.Q.allItems = paramsMap
	cfg.E.AllItems.queryLookup = paramsMap
}

func (cfg *Config) initItemsParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays an item's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays an item's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        QpnHasAbility,
			Description: "Searches for items that can be used in battle.",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnRelatedStat,
			Description:   "Searches for items that are related to the specified stat.",
			Type:          qptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epStats},
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for items that are from one of the specified item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameItemCategory},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Items.subsections)
	cfg.Q.items = paramsMap
	cfg.E.Items.queryLookup = paramsMap
}

func (cfg *Config) initKeyItemsParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays a key-item's related resources with the given availabilities. This affects areas, treasures and quests.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for key-items with the given availabilities. The availability of a key-item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the key-item. The key-item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for key-items that are of the specified key-item categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.KeyItemCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameKeyItemCategory},
		},
		{
			Name:          QpnMethods,
			Description:   "Searches for key-items that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			AllowedValues: []QueryValue{qvTreasure, qvQuest},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for key-items that can be obtained at the specified location. If combined with 'availability', this parameter searches for key-items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for key-items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for key-items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for key-items that can be obtained in the specified area. If combined with 'availability', this parameter searches for key-items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.KeyItems.subsections)
	cfg.Q.keyItems = paramsMap
	cfg.E.KeyItems.queryLookup = paramsMap
}

func (cfg *Config) initSpheresParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays a sphere's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays a sphere's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               QpnColor,
			Description:        "Searches for spheres with any of the given colors.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.SphereColor.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameSphereColor},
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for spheres with the given availabilities. The availability of a sphere is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the sphere. The sphere won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for spheres that can be farmed. Is combinable with 'availability'. The availability assigned to the sphere is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:          QpnMethods,
			Description:   "Searches for spheres that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop, qvQuest, qvBlitzball},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for spheres that can be obtained at the specified location. If combined with 'availability', this parameter searches for spheres within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for spheres that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for spheres within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for spheres that can be obtained in the specified area. If combined with 'availability', this parameter searches for spheres within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Spheres.subsections)
	cfg.Q.spheres = paramsMap
	cfg.E.Spheres.queryLookup = paramsMap
}

func (cfg *Config) initPrimersParams() {
	params := []QueryParam{
		{
			Name:               QpnAvailability,
			Description:        "Searches for primers with the given availabilities. The availability of a primer is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the primer. The primer won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Primers.subsections)
	cfg.Q.primers = paramsMap
	cfg.E.Primers.queryLookup = paramsMap
}

func (cfg *Config) initMixesParams() {
	params := []QueryParam{
		{
			Name:          QpnContainsItem,
			Description:   "Modifies combinations to only display item combinations that include the specified item.",
			Type:          qptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puSingle,
			ConflictsWith: []QueryParamName{QpnBest},
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:          QpnBest,
			Description:   "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:          qptBool,
			ParamUse:      puSingle,
			ConflictsWith: []QueryParamName{QpnContainsItem},
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for mixes that are of the specified mix categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.MixCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameMixCategory},
		},
		{
			Name:          QpnReqItem,
			Description:   "Searches for mixes that can be built with the specified item.",
			Type:          qptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epItems},
		},
		{
			Name:           QpnSecondItem,
			Description:    "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:           qptNameID,
			ExampleVals:    []string{"grenade", "power_sphere"},
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnReqItem},
			ReferencesInt:  []EndpointName{epItems},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Mixes.subsections)
	cfg.Q.mixes = paramsMap
	cfg.E.Mixes.queryLookup = paramsMap
}

func (cfg *Config) initAutoAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays an auto-ability's related resources with the given availabilities. This affects shops, treasures, and monsters.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:        QpnRelRepeatable,
			Description: "Only displays an auto-ability's related resources that can be farmed. This affects shops, treasures, and monsters.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               QpnAvailability,
			Description:        "Searches for auto-abilities with the given availabilities. The availability of an auto-ability is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the auto-ability. The auto-ability won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:           QpnPreAirship,
			Description:    "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:           qptBool,
			ParamUse:       puList,
			RequiredParams: []QueryParamName{QpnAvailability},
		},
		{
			Name:        QpnRepeatable,
			Description: "Searches for auto-abilities that can be farmed. Is combinable with 'availability'. The availability assigned to the auto-ability is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        qptBool,
			ParamUse:    puList,
		},
		{
			Name:               QpnCategory,
			Description:        "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.AutoAbilityCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAutoAbilityCategory},
		},
		{
			Name:               QpnType,
			Description:        "Searches for auto-abilities that are of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.EquipType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:          QpnMonster,
			Description:   "Searches for auto-abilities that are dropped by the specified monster.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnMonsterItems,
			Description:   "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnShop,
			Description:   "Searches for auto-abilities that can be obtained from the specified shop.",
			Type:          qptID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnCharacter,
			Description:   "Restricts the search for 'availability', 'monster' and 'shop' to only include auto-abilities that can be obtained by the specified character. This includes auto-abilities with no character restriction like regular monster equipment drop slots.",
			Type:          qptNameID,
			ExampleVals:   []string{"kimahri"},
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			UsableWith:    []QueryParamName{QpnAvailability, QpnMonster, QpnShop},
			ReferencesInt: []EndpointName{epMonsters},
		},
		{
			Name:          QpnMethods,
			Description:   "Searches for auto-abilities that can be obtained via at least one of the given methods.",
			Type:          qptValueList,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			AllowedValues: []QueryValue{qvMonster, qvTreasure, qvShop},
		},
		{
			Name:        QpnCustomize,
			Description: "Converts the 'availability' and 'repeatable' parameters to search for auto-abilities based on their required item's availability and/or farmability.",
			Type:        qptBool,
			ParamUse:    puList,
			UsableWith:  []QueryParamName{QpnAvailability, QpnRepeatable},
		},
		{
			Name:          QpnLocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified location. If combined with 'availability', this parameter searches for auto-abilities within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this location whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnSublocation, QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epLocations},
		},
		{
			Name:          QpnSublocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for auto-abilities within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnArea, QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epSublocations},
		},
		{
			Name:          QpnArea,
			Description:   "Searches for auto-abilities that can be obtained in the specified area. If combined with 'availability', this parameter searches for auto-abilities within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this area whose sources' farmability matches the given value based on its category.",
			Type:          qptID,
			ParamUse:      puList,
			ReplacedBy:    []QueryParamName{QpnAvailability, QpnRepeatable},
			ReferencesInt: []EndpointName{epAreas},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.AutoAbilities.subsections)
	cfg.Q.autoAbilities = paramsMap
	cfg.E.AutoAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentTablesParams() {
	params := []QueryParam{
		{
			Name:          QpnAutoAbilities,
			Description:   "Searches for equipment tables with all of the given auto-abilities.",
			Type:          qptIDList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:               QpnType,
			Description:        "Searches for equipment tables that are of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.EquipType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:        QpnCelestialWeapon,
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.EquipmentTables.subsections)
	cfg.Q.equipmentTables = paramsMap
	cfg.E.EquipmentTables.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentParams() {
	params := []QueryParam{
		{
			Name:        QpnTable,
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type:        qptInt,
			ParamUse:    puSingle,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:               QpnRelAvailability,
			Description:        "Only displays an equipment's related resources with the given availabilities. This affects treasures and shops.",
			Type:               qptEnumList,
			ParamUse:           puSingle,
			EnumVals:           cfg.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameAvailabilityType},
		},
		{
			Name:          QpnAutoAbilities,
			Description:   "Searches for equipment with all of the given auto-abilities.",
			Type:          qptIDList,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epAutoAbilities},
		},
		{
			Name:          QpnCharacter,
			Description:   "Searches for equipment of the specified character.",
			ExampleVals:   []string{"yuna"},
			Type:          qptNameID,
			ParamUse:      puList,
			ReferencesInt: []EndpointName{epCharacters},
		},
		{
			Name:               QpnType,
			Description:        "Searches for equipment that is of the specified equip type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.EquipType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameEquipType},
		},
		{
			Name:        QpnCelestialWeapon,
			Description: "Searches for the celestial weapons.",
			Type:        qptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Equipment.subsections)
	cfg.Q.equipment = paramsMap
	cfg.E.Equipment.queryLookup = paramsMap
}

func (cfg *Config) initCelestialWeaponsParams() {
	params := []QueryParam{
		{
			Name:               QpnFormula,
			Description:        "Searches for celestial-weapons that are of the specified celestial formula.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.CelestialFormula.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameCelestialFormula},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.CelestialWeapons.subsections)
	cfg.Q.celestialWeapons = paramsMap
	cfg.E.CelestialWeapons.queryLookup = paramsMap
}

func (cfg *Config) initStatsParams() {
	params := []QueryParam{
		{
			Name:        QpnChangesOnly,
			Description: "Only includes a stat's related auto-abilities, abilities, status conditions, and properties that cause stat changes.",
			Type:        qptBool,
			ParamUse:    puSingle,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Stats.subsections)
	cfg.Q.stats = paramsMap
	cfg.E.Stats.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveModesParams() {
	params := []QueryParam{
		{
			Name:               QpnType,
			Description:        "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:               qptEnum,
			ParamUse:           puList,
			EnumVals:           cfg.T.OverdriveModeType.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameOverdriveModeType},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.OverdriveModes.subsections)
	cfg.Q.overdriveModes = paramsMap
	cfg.E.OverdriveModes.queryLookup = paramsMap
}

func (cfg *Config) initStatusConditionsParams() {
	params := []QueryParam{
		{
			Name:            QpnInflictMin,
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
			Name:            QpnInflictMax,
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
			Name:            QpnResistance,
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
			Name:               QpnCategory,
			Description:        "Searches for status conditions that are of the specified status condition categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.StatusConditionCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameStatusConditionCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.StatusConditions.subsections)
	cfg.Q.statusConditions = paramsMap
	cfg.E.StatusConditions.queryLookup = paramsMap
}

func (cfg *Config) initModifiersParams() {
	params := []QueryParam{
		{
			Name:               QpnCategory,
			Description:        "Searches for modifiers that are of the specified modifier categories.",
			Type:               qptEnumList,
			ParamUse:           puList,
			EnumVals:           cfg.T.ModifierCategory.Vals(),
			ReferencesEnumsInt: []EnumName{enumNameModifierCategory},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.Modifiers.subsections)
	cfg.Q.modifiers = paramsMap
	cfg.E.Modifiers.queryLookup = paramsMap
}

func (cfg *Config) initAgilityTierParams() {
	params := []QueryParam{
		{
			Name:            QpnAgility,
			Description:     "Searches for the agility tier that the given agility value belongs to.",
			Type:            qptInt,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 255},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, cfg.E.AgilityTiers.subsections)
	cfg.Q.agilityTiers = paramsMap
	cfg.E.AgilityTiers.queryLookup = paramsMap
}
