package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type ParamUseType string

const (
	puSingle ParamUseType = "single"
	puList   ParamUseType = "list"
	puSimple ParamUseType = "simple"
)

type QueryParam struct {
	ID                 int                          `json:"-"`
	Name               typestrings.QueryParamName   `json:"name"`
	Type               typestrings.QueryParamType   `json:"param_type"`
	Endpoint           typestrings.EndpointName     `json:"-"`
	Description        string                       `json:"description"`
	ExampleVals        []string                     `json:"-"`
	Usage              string                       `json:"usage"`
	ExampleUses        []string                     `json:"example_uses"`
	IsExclusive        bool                         `json:"only_use_alone,omitempty"`
	ParamUse           ParamUseType                 `json:"param_use_type"`
	ForSegment         *typestrings.SectionName     `json:"for_segment,omitempty"`
	EnumVals           []enums.EnumVal              `json:"-"`
	RequiresAll        []typestrings.QueryParamName `json:"required_params,omitempty"`
	RequiresOne        []typestrings.QueryParamName `json:"usable_with,omitempty"`
	ReplacedBy         []typestrings.QueryParamName `json:"replaced_by,omitempty"`
	ConflictsWith      []typestrings.QueryParamName `json:"conflicts_with,omitempty"`
	ReferencesInt      []typestrings.EndpointName   `json:"-"`
	ReferencesEnumsInt []typestrings.EnumName       `json:"-"`
	References         []string                     `json:"references,omitempty"`
	ReferencesEnums    []string                     `json:"references_enums,omitempty"`
	AllowedIDs         []int32                      `json:"-"`
	AllowedResources   []string                     `json:"allowed_resources,omitempty"`
	AllowedValues      []typestrings.QueryValue     `json:"allowed_values,omitempty"`
	AllowedIntRange    []int                        `json:"allowed_int_range,omitempty"`
	AllowedResTypes    []string                     `json:"allowed_res_types,omitempty"`
	DefaultVal         *int                         `json:"default_value,omitempty"`
	SpecialInputs      []SpecialQueryInput          `json:"special_inputs,omitempty"`
}

type SpecialQueryInput struct {
	Key typestrings.QuerySpecialVal `json:"key"`
	Val int                         `json:"value"`
}

func (cfg *Config) initLocationsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays a location's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays a location's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for locations with the given availabilities. Can be combined with other parameters that filter locations by resource/resource-type. In that case, this parameter searches for locations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same location, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for locations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for locations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the location. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresOne: []typestrings.QueryParamName{typestrings.QpnMonster, typestrings.QpnItem, typestrings.QpnKeyItem, typestrings.QpnAutoAbility},
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for locations where the specified monster can be found. If combined with 'availability', the location must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnItem,
			Description:   "Searches for locations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []typestrings.QueryParamName{typestrings.QpnItem},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnKeyItem,
			Description:   "Searches for locations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the location whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpKeyItems},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for locations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:        typestrings.QpnCharacters,
			Description: "Searches for locations where a character permanently joins the party.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnAeons,
			Description: "Searches for locations where a new aeon is acquired.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMonsters,
			Description: "Searches for locations that have monsters. If combined with 'availability', the location must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnBossFights,
			Description: "Searches for locations that have bosses. If combined with 'availability', the location must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnShops,
			Description: "Searches for locations that have shops. If combined with 'availability', the location must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnTreasures,
			Description: "Searches for locations that have treasures. If combined with 'availability', the location must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnSidequests,
			Description: "Searchces for locations that feature sidequests. If combined with 'availability', the location must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnFMVs,
			Description: "Searches for locations that feature fmv sequences.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Locations.completeQueryParamsInit(cfg, params)
	cfg.Q.locations = paramsMap
	cfg.E.Locations.queryLookup = paramsMap
}

func (cfg *Config) initSublocationsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays a sublocation's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays a sublocation's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for sublocations that are located within the specified location.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for sublocations with the given availabilities. Can be combined with other parameters that filter sublocations by resource/resource-type. In that case, this parameter searches for sublocations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same sublocation, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the sublocation won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for sublocations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for sublocations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the sublocation. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresOne: []typestrings.QueryParamName{typestrings.QpnMonster, typestrings.QpnItem, typestrings.QpnKeyItem, typestrings.QpnAutoAbility},
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for sublocations where the specified monster can be found. If combined with 'availability', the sublocation must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnItem,
			Description:   "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []typestrings.QueryParamName{typestrings.QpnItem},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnKeyItem,
			Description:   "Searches for sublocations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpKeyItems},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for sublocations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:        typestrings.QpnCharacters,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnAeons,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMonsters,
			Description: "Searches for sublocations that have monsters. If combined with 'availability', the sublocation must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnBossFights,
			Description: "Searches for sublocations that have bosses. If combined with 'availability', the sublocation must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnShops,
			Description: "Searches for sublocations that have shops. If combined with 'availability', the sublocation must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnTreasures,
			Description: "Searches for sublocations that have treasures. If combined with 'availability', the sublocation must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnSidequests,
			Description: "Searchces for sublocations that feature sidequests. If combined with 'availability', the sublocation must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnFMVs,
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Sublocations.completeQueryParamsInit(cfg, params)
	cfg.Q.sublocations = paramsMap
	cfg.E.Sublocations.queryLookup = paramsMap
}

func (cfg *Config) initAreasParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays an area's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays an area's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for areas that are located within the specified location.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for areas that are located within the specified sublocation.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for areas with the given availabilities. Can be combined with other parameters that filter areas by resource/resource-type. In that case, this parameter searches for areas where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same area, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for areas where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for areas where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the area. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresOne: []typestrings.QueryParamName{typestrings.QpnMonster, typestrings.QpnItem, typestrings.QpnKeyItem, typestrings.QpnAutoAbility},
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for areas where the specified monster can be found. If combined with 'availability', the area must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnItem,
			Description:   "Searches for areas where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []typestrings.QueryParamName{typestrings.QpnItem},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnKeyItem,
			Description:   "Searches for areas where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the area whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpKeyItems},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for areas where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:        typestrings.QpnSaveSphere,
			Description: "Searches for areas that have a save sphere.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnCompSphere,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnAirship,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnChocobo,
			Description: "Searches for areas where you can ride a chocobo.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnCharacters,
			Description: "Searches for areas where a character permanently joins the party.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnAeons,
			Description: "Searches for areas where a new aeon is acquired.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMonsters,
			Description: "Searches for areas that have monsters. If combined with 'availability', the area must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities (based on its encounter areas).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnBossFights,
			Description: "Searches for areas that have bosses. If combined with 'availability', the area must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation's encounter areas).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnShops,
			Description: "Searches for areas that have shops. If combined with 'availability', the area must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnTreasures,
			Description: "Searches for areas that have treasures. If combined with 'availability', the area must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnSidequests,
			Description: "Searchces for areas that feature sidequests. If combined with 'availability', the area must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:        typestrings.QpnFMVs,
			Description: "Searches for areas that feature fmv sequences.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Areas.completeQueryParamsInit(cfg, params)
	cfg.Q.areas = paramsMap
	cfg.E.Areas.queryLookup = paramsMap
}

func (cfg *Config) initMonsterFormationsParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for monster-formations that feature the specified monster.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for monster-formations with the specified monster-formation-categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterFormationCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameMonsterFormationCategory},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for monster-formations with the given availabilities. If combined with the 'area' parameter, the availability of the monster-formation in this specific area is used. If a monster-formation has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster-formation won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with the 'area' parameter, it takes the repeatability directly from the monster-formations that occur in the specified area. Is combinable with 'availability'. In that case, the search looks for the monster-formation that is the most accessible while also being farmable and checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for monster-formations that appear within the specified location. If combined with 'availability', this parameter searches for monster-formations within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this location whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for monster-formations that appear within the specified sublocation. If combined with 'availability', this parameter searches for monster-formations within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this sublocation whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for monster-formations that appear within the specified area. If combined with 'availability', this parameter searches for monster-formations within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monster-formations within this area whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
		{
			Name:        typestrings.QpnAmbush,
			Description: "Searches for monster-formations that are forced ambushes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnEscape,
			Description: "Searches for monster-formations that can be escaped.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.MonsterFormations.completeQueryParamsInit(cfg, params)
	cfg.Q.monsterFormations = paramsMap
	cfg.E.MonsterFormations.queryLookup = paramsMap
}

func (cfg *Config) initShopsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for shops with the specified shop categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ShopCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameShopCategory},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for shops with the given availabilities. By default, this parameter checks, if a shop simply is available at the given availability. If combined with a filter that refers to the shop's inventory, it takes the availability directly from there ('pre-story' for the inventory before acquiring the airship, and 'post' after). In that case, this filter looks, if the requested resources are available in a shop at that point in the game.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for shops that appear at the specified location. If combined with 'availability', this parameter searches for shops within this location whose availability matches one of the specified availabilities.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnAvailability},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for shops that appear at the specified sublocation. If combined with 'availability', this parameter searches for shops within this sublocation whose availability matches one of the specified availabilities.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for shops that offer equipment with the specified auto-ability. Can be combined with 'empty_slots' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnItems, typestrings.QpnEquipment},
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:            typestrings.QpnEmptySlots,
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots. Can be combined with 'auto_ability' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            typestrings.QptIntList,
			ParamUse:        puList,
			ConflictsWith:   []typestrings.QueryParamName{typestrings.QpnItems, typestrings.QpnEquipment},
			ReplacedBy:      []typestrings.QueryParamName{typestrings.QpnAvailability},
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          typestrings.QpnCharacter,
			Description:   "Searches for shops that offer equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnItems, typestrings.QpnEquipment},
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacters},
		},
		{
			Name:          typestrings.QpnItems,
			Description:   "Searches for shops that offer items. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          typestrings.QptBool,
			ParamUse:      puList,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnAutoAbility, typestrings.QpnCharacter, typestrings.QpnEmptySlots},
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:          typestrings.QpnEquipment,
			Description:   "Searches for shops that offer equipment. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          typestrings.QptBool,
			ParamUse:      puList,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnAutoAbility, typestrings.QpnCharacter, typestrings.QpnEmptySlots},
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
	}

	paramsMap := cfg.E.Shops.completeQueryParamsInit(cfg, params)
	cfg.Q.shops = paramsMap
	cfg.E.Shops.queryLookup = paramsMap
}

func (cfg *Config) initTreasuresParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for treasures that appear within the specified location.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for treasures that appear within the specified sublocation.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for treasures that appear within the specified area.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for treasures that contain equipment with the specified auto-ability. Can be combined with 'empty_slots', 'character', and 'availability' for more specific searches.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:            typestrings.QpnEmptySlots,
			Description:     "Searches for treasures that contain equipment with the specified amounts of empty slots. Can be combined with 'auto_ability', 'character', and 'availability' for more specific searches.",
			Type:            typestrings.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          typestrings.QpnCharacter,
			Description:   "Searches for treasures that contain equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacters},
		},
		{
			Name:          typestrings.QpnItem,
			Description:   "Searches for treasures that contain the specified item.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:               typestrings.QpnLootType,
			Description:        "Searches for treasures with the specified loot type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.LootType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameLootType},
		},
		{
			Name:               typestrings.QpnTreasureType,
			Description:        "Searches for treasures with the specified treasure type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.TreasureType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTreasureType},
		},
		{
			Name:        typestrings.QpnAnima,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for treasures with the given availabilities.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
	}

	paramsMap := cfg.E.Treasures.completeQueryParamsInit(cfg, params)
	cfg.Q.treasures = paramsMap
	cfg.E.Treasures.queryLookup = paramsMap
}

func (cfg *Config) initQuestsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for quests that are of the specified quest type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.QuestType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameQuestType},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for quests with the given availabilities.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for quests that can be completed more than once.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Quests.completeQueryParamsInit(cfg, params)
	cfg.Q.quests = paramsMap
	cfg.E.Quests.queryLookup = paramsMap
}

func (cfg *Config) initSidequestsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for sidequests with the given availabilities.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
	}

	paramsMap := cfg.E.Sidequests.completeQueryParamsInit(cfg, params)
	cfg.Q.sidequests = paramsMap
	cfg.E.Sidequests.queryLookup = paramsMap
}

func (cfg *Config) initSubquestsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for subquests with the given availabilities.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for subquests that can be completed more than once.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Subquests.completeQueryParamsInit(cfg, params)
	cfg.Q.subquests = paramsMap
	cfg.E.Subquests.queryLookup = paramsMap
}

func (cfg *Config) initArenaCreationsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for monster-formations with the specified arena-creation-categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ArenaCreationCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameArenaCreationCategory},
		},
	}

	paramsMap := cfg.E.ArenaCreations.completeQueryParamsInit(cfg, params)
	cfg.Q.arenaCreations = paramsMap
	cfg.E.ArenaCreations.queryLookup = paramsMap
}

func (cfg *Config) initBlitzballPrizesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.BlitzballTournamentCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameBlitzballTournamentCategory},
		},
	}

	paramsMap := cfg.E.BlitzballPrizes.completeQueryParamsInit(cfg, params)
	cfg.Q.blitzballPrizes = paramsMap
	cfg.E.BlitzballPrizes.queryLookup = paramsMap
}

func (cfg *Config) initSongsParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
		{
			Name:        typestrings.QpnFMVs,
			Description: "Searches for songs that are played in fmvs.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnSpecialUse,
			Description: "Searches for songs with a special use case.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnComposer,
			Description:        "Searches for songs that were composed by the stated composer.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.Composer.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameComposer},
		},
		{
			Name:               typestrings.QpnArranger,
			Description:        "Searches for songs that were arranged by the stated arranger.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.Arranger.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameArranger},
		},
	}

	paramsMap := cfg.E.Songs.completeQueryParamsInit(cfg, params)
	cfg.Q.songs = paramsMap
	cfg.E.Songs.queryLookup = paramsMap
}

func (cfg *Config) initFMVsParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for fmvs that are played within the specified location.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
	}

	paramsMap := cfg.E.FMVs.completeQueryParamsInit(cfg, params)
	cfg.Q.fmvs = paramsMap
	cfg.E.FMVs.queryLookup = paramsMap
}

func (cfg *Config) initPlayerUnitsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for player units that are of the specified unit type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.UnitType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameUnitType},
		},
	}

	paramsMap := cfg.E.PlayerUnits.completeQueryParamsInit(cfg, params)
	cfg.Q.playerUnits = paramsMap
	cfg.E.PlayerUnits.queryLookup = paramsMap
}

func (cfg *Config) initCharactersParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnOsgStats,
			Description:        "Adds all stat gains within the character's stated sphere grid to their base stats. This includes stat nodes behind sphere locks.",
			Type:               typestrings.QptEnum,
			ParamUse:           puSingle,
			EnumVals:           enums.T.SphereGridType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameSphereGridType},
		},
		{
			Name:        typestrings.QpnStoryBased,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnUnderwater,
			Description: "Searches for characters that can fight underwater.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Characters.completeQueryParamsInit(cfg, params)
	cfg.Q.characters = paramsMap
	cfg.E.Characters.queryLookup = paramsMap
}

func (cfg *Config) initAeonsParams() {
	params := []QueryParam{
		{
			Name:            typestrings.QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            typestrings.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:        typestrings.QpnYunaStats,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter.",
			Type:        typestrings.QptStat,
			ExampleUses: []string{exampleUse(typestrings.QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(typestrings.QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
		},
		{
			Name:        typestrings.QpnOptional,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Aeons.completeQueryParamsInit(cfg, params)
	cfg.Q.aeons = paramsMap
	cfg.E.Aeons.queryLookup = paramsMap
}

func (cfg *Config) initCharacterClassesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for character classes with the specified categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.CharacterClassCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameCharacterClassCategory},
		},
	}

	paramsMap := cfg.E.CharacterClasses.completeQueryParamsInit(cfg, params)
	cfg.Q.characterClasses = paramsMap
	cfg.E.CharacterClasses.queryLookup = paramsMap
}

func (cfg *Config) initMonstersParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays a monster's related resources with the given availabilities. This affects areas and monster-formations.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays a monster's related resources that can be farmed. This affects areas and monster-formations.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        typestrings.QpnKimahriStats,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type:        typestrings.QptStat,
			ExampleUses: []string{exampleUse(typestrings.QpnKimahriStats, "hp:3000,strength:25,magic:30,agility:40"), exampleUse(typestrings.QpnKimahriStats, "hp:15000,agility:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:            typestrings.QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating a possessed aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            typestrings.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
			AllowedIDs:      []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        typestrings.QpnYunaStats,
			Description: "Calculate a possessed aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter. All stats are replaceable, except for MP and Luck. While Yuna's MP and Luck still play a factor in the calculation of the other stats, a posssessed aeon's own MP and Luck will always be set to 1. If unused, the aeon's stats are based on your own aeon's default stats.",
			Type:        typestrings.QptStat,
			ExampleUses: []string{exampleUse(typestrings.QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(typestrings.QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        typestrings.QpnAlteredState,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:        typestrings.QptID,
			ParamUse:    puSingle,
		},
		{
			Name:          typestrings.QpnOmnisElements,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type:          typestrings.QptOther,
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{exampleUse(typestrings.QpnOmnisElements, "ifil"), exampleUse(typestrings.QpnOmnisElements, "wilf")},
			ParamUse:      puSingle,
			AllowedIDs:    []int32{211},
			AllowedValues: []typestrings.QueryValue{typestrings.QvF, typestrings.QvL, typestrings.QvW, typestrings.QvI},
		},
		{
			Name:               typestrings.QpnElementalResists,
			Description:        "Searches for monsters that have the specified elemental affinities.",
			Type:               typestrings.QptOther,
			Usage:              "?elemental_resists={element|id}:{affinity|id},...",
			ExampleUses:        []string{exampleUse(typestrings.QpnElementalResists, "fire:weak,water:absorb"), exampleUse(typestrings.QpnElementalResists, "1:3,2:4")},
			ParamUse:           puList,
			ReferencesInt:      []typestrings.EndpointName{typestrings.EpElements},
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameElementalAffinity},
		},
		{
			Name:          typestrings.QpnStatusResists,
			Description:   "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type:          typestrings.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:            typestrings.QpnResistance,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            typestrings.QptInt,
			ParamUse:        puList,
			RequiresAll:     []typestrings.QueryParamName{typestrings.QpnStatusResists},
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: typestrings.QsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:          typestrings.QpnItem,
			Description:   "Searches for monsters that have the specified item as loot. Can be specified further with the 'methods' parameter.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []typestrings.QueryParamName{typestrings.QpnItem},
			AllowedValues: []typestrings.QueryValue{typestrings.QvSteal, typestrings.QvDrop, typestrings.QvBribe, typestrings.QvOther},
		},
		{
			Name:          typestrings.QpnAutoAbility,
			Description:   "Searches for monsters that drop the specified auto-ability.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:        typestrings.QpnIsForced,
			Description: "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAutoAbility},
		},
		{
			Name:            typestrings.QpnEmptySlots,
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type:            typestrings.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:          typestrings.QpnRonsoRage,
			Description:   "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpRonsoRages},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for monsters with the given availabilities. If combined with a geographical filter, it takes the availability directly from the monster-formations that occur in the specified location, sublocation, or area. If a monster has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with a geographical filter, it takes the repeatability directly from the monster-formations that occur in the specified location, sublocation, or area. Is combinable with 'availability'. The availability assigned to the monster is from the monster-formation that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for monsters that appear within the specified location. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this location whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for monsters that appear within the specified sublocation. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for monsters that appear within the specified area. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this area whose farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
		{
			Name:            typestrings.QpnDistance,
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type:            typestrings.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:        typestrings.QpnCapture,
			Description: "Searches for monsters that can be captured.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnHasOverdrive,
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnUnderwater,
			Description: "Searches for monsters that are fought underwater.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnZombie,
			Description: "Searches for monsters that are zombies.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnSpecies,
			Description:        "Searches for monsters of the specified species.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterSpecies.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameMonsterSpecies},
		},
		{
			Name:               typestrings.QpnCreationArea,
			Description:        "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.CreationArea.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameCreationArea},
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for monsters that are of the specified monster-categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameMonsterCategory},
		},
	}

	paramsMap := cfg.E.Monsters.completeQueryParamsInit(cfg, params)
	cfg.Q.monsters = paramsMap
	cfg.E.Monsters.queryLookup = paramsMap
}

func (cfg *Config) initAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for abilities that are of the specified ability types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AbilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAbilityType},
		},
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for abilities with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnCopycat,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnHelpBar,
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for abilities that can be used by the specified monster.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:               typestrings.QpnTargetType,
			Description:        "Searches for abilities with the specified target types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTargetType},
		},
		{
			Name:        typestrings.QpnUserAtk,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnDarkable,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnSilenceable,
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnReflectable,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnAttackType,
			Description:        "Searches for abilities with battle interactions of the specified attack types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAttackType},
		},
		{
			Name:               typestrings.QpnDamageType,
			Description:        "Searches for abilities that deal the specified types of damage.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageType},
		},
		{
			Name:               typestrings.QpnDamageFormula,
			Description:        "Searches for abilities that use the specified formula to calculate their damage.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageFormula},
		},
		{
			Name:        typestrings.QpnCanCrit,
			Description: "Searches for abilities that can land critical hits.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnBDL,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnElement,
			Description:   "Searches for abilities that deal elemental damage based on the specified element.",
			Type:          typestrings.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpElements},
		},
		{
			Name:        typestrings.QpnDelay,
			Description: "Searches for abilities that deal delay.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnStatusInflict,
			Description:   "Searches for abilities that can inflict the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:          typestrings.QpnStatusRemove,
			Description:   "Searches for abilities that can remove the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:        typestrings.QpnStatChanges,
			Description: "Searches for abilities that cause stat changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnModChanges,
			Description: "Searches for abilities that cause modifier changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Abilities.completeQueryParamsInit(cfg, params)
	cfg.Q.abilities = paramsMap
	cfg.E.Abilities.queryLookup = paramsMap
}

func (cfg *Config) initPlayerAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnAbilityUser,
			Description:   "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpPlayerUnits},
		},
		{
			Name:        typestrings.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the player ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:        typestrings.QpnBombWpn,
			Description: "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for player abilities with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnCopycat,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnHelpBar,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for player abilities that are of the specified player ability categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.PlayerAbilityCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNamePlayerAbilityCategory},
		},
		{
			Name:        typestrings.QpnOutsideBattle,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMp,
			Description: "Searches for player abilities with the specified mp costs.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMpMin,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type:        typestrings.QptInt,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnMpMax,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type:        typestrings.QptInt,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnRelatedStat,
			Description:   "Searches for player abilities that are related to the specified stat.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStats},
		},
		{
			Name:          typestrings.QpnUser,
			Description:   "Searches for player abilities that are learned by the specified character class.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacterClasses},
		},
		{
			Name:          typestrings.QpnStdSg,
			Description:   "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacters},
		},
		{
			Name:          typestrings.QpnExpSg,
			Description:   "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacters},
		},
		{
			Name:          typestrings.QpnLearnItem,
			Description:   "Searches for player abilities an aeon can learn via the specified item.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:               typestrings.QpnTargetType,
			Description:        "Searches for player abilities with the specified target types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTargetType},
		},
		{
			Name:        typestrings.QpnUserAtk,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnDarkable,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnSilenceable,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnReflectable,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnAttackType,
			Description:        "Searches for player abilities with battle interactions of the specified attack types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAttackType},
		},
		{
			Name:               typestrings.QpnDamageType,
			Description:        "Searches for player abilities that deal the specified types of damage.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageType},
		},
		{
			Name:               typestrings.QpnDamageFormula,
			Description:        "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageFormula},
		},
		{
			Name:          typestrings.QpnElement,
			Description:   "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:          typestrings.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpElements},
		},
		{
			Name:        typestrings.QpnDelay,
			Description: "Searches for player abilities that deal delay.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnStatusInflict,
			Description:   "Searches for player abilities that can inflict the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:          typestrings.QpnStatusRemove,
			Description:   "Searches for player abilities that can remove the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:        typestrings.QpnStatChanges,
			Description: "Searches for player abilities that cause stat changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnModChanges,
			Description: "Searches for player abilities that cause modifier changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.PlayerAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.playerAbilities = paramsMap
	cfg.E.PlayerAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnUser,
			Description:   "Searches for overdrive abilities that are learned by the specified character class.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacterClasses},
		},
		{
			Name:          typestrings.QpnRelatedStat,
			Description:   "Searches for overdrive abilities that are related to the specified stat.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStats},
		},
		{
			Name:               typestrings.QpnTargetType,
			Description:        "Searches for overdrive abilities with the specified target types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTargetType},
		},
		{
			Name:               typestrings.QpnAttackType,
			Description:        "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAttackType},
		},
		{
			Name:               typestrings.QpnDamageFormula,
			Description:        "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageFormula},
		},
		{
			Name:        typestrings.QpnCanCrit,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnElement,
			Description:   "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:          typestrings.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpElements},
		},
		{
			Name:        typestrings.QpnDelay,
			Description: "Searches for overdrive abilities that deal delay.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnStatusInflict,
			Description:   "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:          typestrings.QpnStatusRemove,
			Description:   "Searches for overdrive abilities that can remove the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:        typestrings.QpnStatChanges,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnModChanges,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.OverdriveAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.overdriveAbilities = paramsMap
	cfg.E.OverdriveAbilities.queryLookup = paramsMap
}

func (cfg *Config) initItemAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for item abilities that are of the specified item categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameItemCategory},
		},
		{
			Name:        typestrings.QpnOutsideBattle,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnRelatedStat,
			Description:   "Searches for item abilities that are related to the specified stat.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStats},
		},
		{
			Name:               typestrings.QpnTargetType,
			Description:        "Searches for item abilities with the specified target types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTargetType},
		},
		{
			Name:               typestrings.QpnAttackType,
			Description:        "Searches for item abilities with battle interactions of the specified attack types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAttackType},
		},
		{
			Name:               typestrings.QpnDamageFormula,
			Description:        "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageFormula},
		},
		{
			Name:          typestrings.QpnElement,
			Description:   "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:          typestrings.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpElements},
		},
		{
			Name:        typestrings.QpnDelay,
			Description: "Searches for item abilities that deal delay.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnStatusInflict,
			Description:   "Searches for item abilities that can inflict the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:          typestrings.QpnStatusRemove,
			Description:   "Searches for item abilities that can remove the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:        typestrings.QpnStatChanges,
			Description: "Searches for item abilities that cause stat changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnModChanges,
			Description: "Searches for item abilities that cause modifier changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.ItemAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.itemAbilities = paramsMap
	cfg.E.ItemAbilities.queryLookup = paramsMap
}

func (cfg *Config) initTriggerCommandsParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnAbilityUser,
			Description:   "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpPlayerUnits},
		},
		{
			Name:        typestrings.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the trigger command's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:        typestrings.QpnBombWpn,
			Description: "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:          typestrings.QpnRelatedStat,
			Description:   "Searches for trigger commands that are related to the specified stat.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStats},
		},
		{
			Name:          typestrings.QpnUser,
			Description:   "Searches for trigger commands that are learned by the specified character class.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacterClasses},
		},
	}

	paramsMap := cfg.E.TriggerCommands.completeQueryParamsInit(cfg, params)
	cfg.Q.triggerCommands = paramsMap
	cfg.E.TriggerCommands.queryLookup = paramsMap
}

func (cfg *Config) initMiscAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnAbilityUser,
			Description:   "If a misc ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpPlayerUnits},
		},
		{
			Name:        typestrings.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the misc ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:        typestrings.QpnBombWpn,
			Description: "If a misc ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAbilityUser},
		},
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for misc abilities with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnCopycat,
			Description: "Searches for misc abilities that can be copied by 'copycat'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnHelpBar,
			Description: "Searches for misc abilities whose names appear in the help bar.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnUser,
			Description:   "Searches for misc abilities that are learned by the specified character class.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacterClasses},
		},
		{
			Name:        typestrings.QpnUserAtk,
			Description: "Searches for misc abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.MiscAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.miscAbilities = paramsMap
	cfg.E.MiscAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEnemyAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for enemy abilities with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnHelpBar,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for enemy abilities that can be used by the specified monster.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:               typestrings.QpnTargetType,
			Description:        "Searches for enemy abilities with the specified target types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameTargetType},
		},
		{
			Name:        typestrings.QpnDarkable,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnSilenceable,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnReflectable,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnAttackType,
			Description:        "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAttackType},
		},
		{
			Name:               typestrings.QpnDamageType,
			Description:        "Searches for enemy abilities that deal the specified types of damage.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageType},
		},
		{
			Name:               typestrings.QpnDamageFormula,
			Description:        "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameDamageFormula},
		},
		{
			Name:        typestrings.QpnCanCrit,
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        typestrings.QpnBDL,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnElement,
			Description:   "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:          typestrings.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpElements},
		},
		{
			Name:        typestrings.QpnDelay,
			Description: "Searches for enemy abilities that deal delay.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnStatusInflict,
			Description:   "Searches for enemy abilities that can inflict the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
		{
			Name:          typestrings.QpnStatusRemove,
			Description:   "Searches for enemy abilities that can remove the specified status condition.",
			Type:          typestrings.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStatusConditions},
		},
	}

	paramsMap := cfg.E.EnemyAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.enemyAbilities = paramsMap
	cfg.E.EnemyAbilities.queryLookup = paramsMap
}

func (cfg *Config) initOverdrivesParams() {
	params := []QueryParam{
		{
			Name:        typestrings.QpnRank,
			Description: "Searches for overdrives with the specified ranks.",
			Type:        typestrings.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnUser,
			Description:   "Searches for overdrives that are learned by the specified character class.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacterClasses},
		},
	}

	paramsMap := cfg.E.Overdrives.completeQueryParamsInit(cfg, params)
	cfg.Q.overdrives = paramsMap
	cfg.E.Overdrives.queryLookup = paramsMap
}

func (cfg *Config) initSubmenusParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnTopmenu,
			Description:   "Searches for submenus that are found within the specified topmenu.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"main", "left"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpTopmenus},
		},
	}

	paramsMap := cfg.E.Submenus.completeQueryParamsInit(cfg, params)
	cfg.Q.submenus = paramsMap
	cfg.E.Submenus.queryLookup = paramsMap
}

func (cfg *Config) initAllItemsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only considers an item's related resources with the given availabilities when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only considers an item's related resources that can be farmed when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for items that are of the specified item-types.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameItemType},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
	}

	paramsMap := cfg.E.AllItems.completeQueryParamsInit(cfg, params)
	cfg.Q.allItems = paramsMap
	cfg.E.AllItems.queryLookup = paramsMap
}

func (cfg *Config) initItemsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays an item's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays an item's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        typestrings.QpnHasAbility,
			Description: "Searches for items that can be used in battle.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnRelatedStat,
			Description:   "Searches for items that are related to the specified stat.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpStats},
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for items that are from one of the specified item categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameItemCategory},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
	}

	paramsMap := cfg.E.Items.completeQueryParamsInit(cfg, params)
	cfg.Q.items = paramsMap
	cfg.E.Items.queryLookup = paramsMap
}

func (cfg *Config) initKeyItemsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays a key-item's related resources with the given availabilities. This affects areas, treasures and quests.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for key-items with the given availabilities. The availability of a key-item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the key-item. The key-item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for key-items that are of the specified key-item categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.KeyItemCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameKeyItemCategory},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Searches for key-items that can be obtained via at least one of the given methods.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			AllowedValues: []typestrings.QueryValue{typestrings.QvTreasure, typestrings.QvQuest},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for key-items that can be obtained at the specified location. If combined with 'availability', this parameter searches for key-items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this location whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for key-items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for key-items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for key-items that can be obtained in the specified area. If combined with 'availability', this parameter searches for key-items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this area whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
	}

	paramsMap := cfg.E.KeyItems.completeQueryParamsInit(cfg, params)
	cfg.Q.keyItems = paramsMap
	cfg.E.KeyItems.queryLookup = paramsMap
}

func (cfg *Config) initSpheresParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays a sphere's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays a sphere's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               typestrings.QpnColor,
			Description:        "Searches for spheres with any of the given colors.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.SphereColor.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameSphereColor},
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for spheres with the given availabilities. The availability of a sphere is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the sphere. The sphere won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for spheres that can be farmed. Is combinable with 'availability'. The availability assigned to the sphere is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Searches for spheres that can be obtained via at least one of the given methods.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop, typestrings.QvQuest, typestrings.QvBlitzball},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for spheres that can be obtained at the specified location. If combined with 'availability', this parameter searches for spheres within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this location whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for spheres that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for spheres within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for spheres that can be obtained in the specified area. If combined with 'availability', this parameter searches for spheres within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this area whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
	}

	paramsMap := cfg.E.Spheres.completeQueryParamsInit(cfg, params)
	cfg.Q.spheres = paramsMap
	cfg.E.Spheres.queryLookup = paramsMap
}

func (cfg *Config) initPrimersParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for primers with the given availabilities. The availability of a primer is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the primer. The primer won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
	}

	paramsMap := cfg.E.Primers.completeQueryParamsInit(cfg, params)
	cfg.Q.primers = paramsMap
	cfg.E.Primers.queryLookup = paramsMap
}

func (cfg *Config) initMixesParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnContainsItem,
			Description:   "Modifies combinations to only display item combinations that include the specified item.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puSingle,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnBest},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnBest,
			Description:   "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:          typestrings.QptBool,
			ParamUse:      puSingle,
			ConflictsWith: []typestrings.QueryParamName{typestrings.QpnContainsItem},
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for mixes that are of the specified mix categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MixCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameMixCategory},
		},
		{
			Name:          typestrings.QpnReqItem,
			Description:   "Searches for mixes that can be built with the specified item.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
		{
			Name:          typestrings.QpnSecondItem,
			Description:   "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			RequiresAll:   []typestrings.QueryParamName{typestrings.QpnReqItem},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpItems},
		},
	}

	paramsMap := cfg.E.Mixes.completeQueryParamsInit(cfg, params)
	cfg.Q.mixes = paramsMap
	cfg.E.Mixes.queryLookup = paramsMap
}

func (cfg *Config) initAutoAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays an auto-ability's related resources with the given availabilities. This affects shops, treasures, and monsters.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnRelRepeatable,
			Description: "Only displays an auto-ability's related resources that can be farmed. This affects shops, treasures, and monsters.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               typestrings.QpnAvailability,
			Description:        "Searches for auto-abilities with the given availabilities. The availability of an auto-ability is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the auto-ability. The auto-ability won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:        typestrings.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresAll: []typestrings.QueryParamName{typestrings.QpnAvailability},
		},
		{
			Name:        typestrings.QpnRepeatable,
			Description: "Searches for auto-abilities that can be farmed. Is combinable with 'availability'. The availability assigned to the auto-ability is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AutoAbilityCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAutoAbilityCategory},
		},
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for auto-abilities that are of the specified equip type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameEquipType},
		},
		{
			Name:          typestrings.QpnMonster,
			Description:   "Searches for auto-abilities that are dropped by the specified monster.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnMonsterItems,
			Description:   "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnShop,
			Description:   "Searches for auto-abilities that can be obtained from the specified shop.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnCharacter,
			Description:   "Restricts the search for 'availability', 'monster' and 'shop' to only include auto-abilities that can be obtained by the specified character. This includes auto-abilities with no character restriction like regular monster equipment drop slots.",
			Type:          typestrings.QptNameID,
			ExampleVals:   []string{"kimahri"},
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			RequiresOne:   []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnMonster, typestrings.QpnShop},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpMonsters},
		},
		{
			Name:          typestrings.QpnMethods,
			Description:   "Searches for auto-abilities that can be obtained via at least one of the given methods.",
			Type:          typestrings.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			AllowedValues: []typestrings.QueryValue{typestrings.QvMonster, typestrings.QvTreasure, typestrings.QvShop},
		},
		{
			Name:        typestrings.QpnCustomize,
			Description: "Converts the 'availability' and 'repeatable' parameters to search for auto-abilities based on their required item's availability and/or farmability.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
			RequiresOne: []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
		},
		{
			Name:          typestrings.QpnLocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified location. If combined with 'availability', this parameter searches for auto-abilities within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this location whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnSublocation, typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpLocations},
		},
		{
			Name:          typestrings.QpnSublocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for auto-abilities within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnArea, typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpSublocations},
		},
		{
			Name:          typestrings.QpnArea,
			Description:   "Searches for auto-abilities that can be obtained in the specified area. If combined with 'availability', this parameter searches for auto-abilities within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this area whose sources' farmability matches the given value based on its category.",
			Type:          typestrings.QptID,
			ParamUse:      puList,
			ReplacedBy:    []typestrings.QueryParamName{typestrings.QpnAvailability, typestrings.QpnRepeatable},
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAreas},
		},
	}

	paramsMap := cfg.E.AutoAbilities.completeQueryParamsInit(cfg, params)
	cfg.Q.autoAbilities = paramsMap
	cfg.E.AutoAbilities.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentTablesParams() {
	params := []QueryParam{
		{
			Name:          typestrings.QpnAutoAbilities,
			Description:   "Searches for equipment tables with all of the given auto-abilities.",
			Type:          typestrings.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for equipment tables that are of the specified equip type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameEquipType},
		},
		{
			Name:        typestrings.QpnCelestialWeapon,
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.EquipmentTables.completeQueryParamsInit(cfg, params)
	cfg.Q.equipmentTables = paramsMap
	cfg.E.EquipmentTables.queryLookup = paramsMap
}

func (cfg *Config) initEquipmentParams() {
	params := []QueryParam{
		{
			Name:        typestrings.QpnTable,
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type:        typestrings.QptInt,
			ParamUse:    puSingle,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:               typestrings.QpnRelAvailability,
			Description:        "Only displays an equipment's related resources with the given availabilities. This affects treasures and shops.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameAvailabilityType},
		},
		{
			Name:          typestrings.QpnAutoAbilities,
			Description:   "Searches for equipment with all of the given auto-abilities.",
			Type:          typestrings.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpAutoAbilities},
		},
		{
			Name:          typestrings.QpnCharacter,
			Description:   "Searches for equipment of the specified character.",
			ExampleVals:   []string{"yuna"},
			Type:          typestrings.QptNameID,
			ParamUse:      puList,
			ReferencesInt: []typestrings.EndpointName{typestrings.EpCharacters},
		},
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for equipment that is of the specified equip type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameEquipType},
		},
		{
			Name:        typestrings.QpnCelestialWeapon,
			Description: "Searches for the celestial weapons.",
			Type:        typestrings.QptBool,
			ParamUse:    puList,
		},
	}

	paramsMap := cfg.E.Equipment.completeQueryParamsInit(cfg, params)
	cfg.Q.equipment = paramsMap
	cfg.E.Equipment.queryLookup = paramsMap
}

func (cfg *Config) initCelestialWeaponsParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnFormula,
			Description:        "Searches for celestial-weapons that are of the specified celestial formula.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.CelestialFormula.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameCelestialFormula},
		},
	}

	paramsMap := cfg.E.CelestialWeapons.completeQueryParamsInit(cfg, params)
	cfg.Q.celestialWeapons = paramsMap
	cfg.E.CelestialWeapons.queryLookup = paramsMap
}

func (cfg *Config) initStatsParams() {
	params := []QueryParam{
		{
			Name:        typestrings.QpnChangesOnly,
			Description: "Only includes a stat's related auto-abilities, abilities, status conditions, and properties that cause stat changes.",
			Type:        typestrings.QptBool,
			ParamUse:    puSingle,
		},
	}

	paramsMap := cfg.E.Stats.completeQueryParamsInit(cfg, params)
	cfg.Q.stats = paramsMap
	cfg.E.Stats.queryLookup = paramsMap
}

func (cfg *Config) initOverdriveModesParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnType,
			Description:        "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:               typestrings.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.OverdriveModeType.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameOverdriveModeType},
		},
	}

	paramsMap := cfg.E.OverdriveModes.completeQueryParamsInit(cfg, params)
	cfg.Q.overdriveModes = paramsMap
	cfg.E.OverdriveModes.queryLookup = paramsMap
}

func (cfg *Config) initStatusConditionsParams() {
	params := []QueryParam{
		{
			Name:            typestrings.QpnInflictMin,
			Description:     "Only shows a status condition's related abilities with an infliction rate higher than or equal to the given amount. The default value is '1'. Can be combined with 'inflict_max', but can't be higher. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            typestrings.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: typestrings.QsvInfinite,
					Val: 254,
				},
				{
					Key: typestrings.QsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:            typestrings.QpnInflictMax,
			Description:     "Only shows a status condition's related abilities with an infliction rate lower than or equal to the given amount. The default value is '25'. Can be combined with 'inflict_min', but can't be lower. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            typestrings.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: typestrings.QsvInfinite,
					Val: 254,
				},
				{
					Key: typestrings.QsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(255),
		},
		{
			Name:            typestrings.QpnResistance,
			Description:     "Only shows a status condition's related monsters with a resistance higher than or equal to the given amount. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            typestrings.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: typestrings.QsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for status conditions that are of the specified status condition categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.StatusConditionCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameStatusConditionCategory},
		},
	}

	paramsMap := cfg.E.StatusConditions.completeQueryParamsInit(cfg, params)
	cfg.Q.statusConditions = paramsMap
	cfg.E.StatusConditions.queryLookup = paramsMap
}

func (cfg *Config) initModifiersParams() {
	params := []QueryParam{
		{
			Name:               typestrings.QpnCategory,
			Description:        "Searches for modifiers that are of the specified modifier categories.",
			Type:               typestrings.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ModifierCategory.Vals(),
			ReferencesEnumsInt: []typestrings.EnumName{typestrings.EnumNameModifierCategory},
		},
	}

	paramsMap := cfg.E.Modifiers.completeQueryParamsInit(cfg, params)
	cfg.Q.modifiers = paramsMap
	cfg.E.Modifiers.queryLookup = paramsMap
}

func (cfg *Config) initAgilityTierParams() {
	params := []QueryParam{
		{
			Name:            typestrings.QpnAgility,
			Description:     "Searches for the agility tier that the given agility value belongs to.",
			Type:            typestrings.QptInt,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 255},
		},
	}

	paramsMap := cfg.E.AgilityTiers.completeQueryParamsInit(cfg, params)
	cfg.Q.agilityTiers = paramsMap
	cfg.E.AgilityTiers.queryLookup = paramsMap
}
