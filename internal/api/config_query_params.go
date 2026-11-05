package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type ParamUseType string

const (
	puSingle ParamUseType = "single"
	puList   ParamUseType = "list"
	puSimple ParamUseType = "simple"
)

type QueryParam struct {
	ID                 int                 `json:"-"`
	Name               ts.QueryParamName   `json:"name"`
	Type               ts.QueryParamType   `json:"param_type"`
	Endpoint           ts.EndpointName     `json:"-"`
	Description        string              `json:"description"`
	ExampleVals        []string            `json:"-"`
	Usage              string              `json:"usage"`
	ExampleUses        []string            `json:"example_uses"`
	IsExclusive        bool                `json:"only_use_alone,omitempty"`
	ParamUse           ParamUseType        `json:"param_use_type"`
	ForSegment         *ts.SectionName     `json:"for_segment,omitempty"`
	EnumVals           []enums.EnumVal     `json:"-"`
	RequiresAll        []ts.QueryParamName `json:"required_params,omitempty"`
	RequiresOne        []ts.QueryParamName `json:"usable_with,omitempty"`
	ReplacedBy         []ts.QueryParamName `json:"replaced_by,omitempty"`
	ConflictsWith      []ts.QueryParamName `json:"conflicts_with,omitempty"`
	ReferencesInt      []ts.EndpointName   `json:"-"`
	ReferencesEnumsInt []ts.EnumName       `json:"-"`
	References         []string            `json:"references,omitempty"`
	ReferencesEnums    []string            `json:"references_enums,omitempty"`
	AllowedIDs         []int32             `json:"-"`
	AllowedResources   []string            `json:"allowed_resources,omitempty"`
	AllowedValues      []ts.QueryValue     `json:"allowed_values,omitempty"`
	AllowedIntRange    []int               `json:"allowed_int_range,omitempty"`
	AllowedResTypes    []string            `json:"allowed_res_types,omitempty"`
	DefaultVal         *int                `json:"default_value,omitempty"`
	SpecialInputs      []SpecialQueryInput `json:"special_inputs,omitempty"`
}

type SpecialQueryInput struct {
	Key ts.QuerySpecialVal `json:"key"`
	Val int                `json:"value"`
}

func initLocationsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays a location's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays a location's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for locations with the given availabilities. Can be combined with other parameters that filter locations by resource/resource-type. In that case, this parameter searches for locations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same location, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for locations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for locations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the location. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresOne: []ts.QueryParamName{ts.QpnMonster, ts.QpnItem, ts.QpnKeyItem, ts.QpnAutoAbility},
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for locations where the specified monster can be found. If combined with 'availability', the location must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnItem,
			Description:   "Searches for locations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []ts.QueryParamName{ts.QpnItem},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnKeyItem,
			Description:   "Searches for locations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the location whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpKeyItems},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for locations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:        ts.QpnCharacters,
			Description: "Searches for locations where a character permanently joins the party.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnAeons,
			Description: "Searches for locations where a new aeon is acquired.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMonsters,
			Description: "Searches for locations that have monsters. If combined with 'availability', the location must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnBossFights,
			Description: "Searches for locations that have bosses. If combined with 'availability', the location must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnShops,
			Description: "Searches for locations that have shops. If combined with 'availability', the location must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnTreasures,
			Description: "Searches for locations that have treasures. If combined with 'availability', the location must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnSidequests,
			Description: "Searchces for locations that feature sidequests. If combined with 'availability', the location must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnFMVs,
			Description: "Searches for locations that feature fmv sequences.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.locations = E.Locations.completeQueryParamsInit(params)
}

func initSublocationsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays a sublocation's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays a sublocation's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for sublocations that are located within the specified location.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for sublocations with the given availabilities. Can be combined with other parameters that filter sublocations by resource/resource-type. In that case, this parameter searches for sublocations where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same sublocation, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the sublocation won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for sublocations where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for sublocations where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the sublocation. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresOne: []ts.QueryParamName{ts.QpnMonster, ts.QpnItem, ts.QpnKeyItem, ts.QpnAutoAbility},
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for sublocations where the specified monster can be found. If combined with 'availability', the sublocation must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnItem,
			Description:   "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []ts.QueryParamName{ts.QpnItem},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnKeyItem,
			Description:   "Searches for sublocations where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpKeyItems},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for sublocations where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:        ts.QpnCharacters,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnAeons,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMonsters,
			Description: "Searches for sublocations that have monsters. If combined with 'availability', the sublocation must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnBossFights,
			Description: "Searches for sublocations that have bosses. If combined with 'availability', the sublocation must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnShops,
			Description: "Searches for sublocations that have shops. If combined with 'availability', the sublocation must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnTreasures,
			Description: "Searches for sublocations that have treasures. If combined with 'availability', the sublocation must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnSidequests,
			Description: "Searchces for sublocations that feature sidequests. If combined with 'availability', the sublocation must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnFMVs,
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.sublocations = E.Sublocations.completeQueryParamsInit(params)
}

func initAreasParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays an area's related resources with the given availabilities. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays an area's related resources that can be farmed. This affects shops, treasures, quests, monsters, and monster-formations.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for areas that are located within the specified location.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for areas that are located within the specified sublocation.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for areas with the given availabilities. Can be combined with other parameters that filter areas by resource/resource-type. In that case, this parameter searches for areas where all the requested resources/resource-types are present with the given availabilities. If a resource (like an item or a monster) has multiple availabilities in the same area, because there are multiple ways of receiving/encountering it, this filter defines the most accessible version of it as its actual availability. In that case, the area won't show up for the other availability types, even if the resource technically can have that availability, since it can be received/encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for areas where the searched resources can be farmed. Must be combined with a parameter that looks up a resource ('monster', 'item', 'key_item'). This parameter looks for areas where all the requested resources are farmable. Is combinable with 'availability'. In that case, the most accessible availability where all resources are farmable is chosen and used for the area. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresOne: []ts.QueryParamName{ts.QpnMonster, ts.QpnItem, ts.QpnKeyItem, ts.QpnAutoAbility},
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for areas where the specified monster can be found. If combined with 'availability', the area must contain a monster-formation with the monster and whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', the monster must have a monster-formation whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnItem,
			Description:   "Searches for areas where the specified item can be acquired. Can be specified further with the 'methods' parameter. If combined with 'availability', the item must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the item must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []ts.QueryParamName{ts.QpnItem},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnKeyItem,
			Description:   "Searches for areas where the specified key-item can be acquired. If combined with 'availability', the key-item must have a source inside the area whose most accessible availability matches one of the specified availabilities. Key-items are never farmable, so combining this parameter with 'repeatable' will either yield 0 results (true) or the results won't be affected (false).",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpKeyItems},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for areas where the specified auto-ability can be acquired. If combined with 'availability', the auto-ability must have a source inside the area whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', the auto-ability must have a source whose farmability matches the given value.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:        ts.QpnSaveSphere,
			Description: "Searches for areas that have a save sphere.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnCompSphere,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnAirship,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnChocobo,
			Description: "Searches for areas where you can ride a chocobo.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnCharacters,
			Description: "Searches for areas where a character permanently joins the party.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnAeons,
			Description: "Searches for areas where a new aeon is acquired.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMonsters,
			Description: "Searches for areas that have monsters. If combined with 'availability', the area must inhabit at least one monster-formation whose most accessible availability matches one of the specified availabilities (based on its encounter areas).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnBossFights,
			Description: "Searches for areas that have bosses. If combined with 'availability', the area must inhabit at least one boss fight whose most accessible availability matches one of the specified availabilities (based on its monster-formation's encounter areas).",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnShops,
			Description: "Searches for areas that have shops. If combined with 'availability', the area must contain at least one shop whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnTreasures,
			Description: "Searches for areas that have treasures. If combined with 'availability', the area must contain at least one treasure whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnSidequests,
			Description: "Searchces for areas that feature sidequests. If combined with 'availability', the area must contain at least one quest whose availability matches one of the specified availabilities.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			ReplacedBy:  []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:        ts.QpnFMVs,
			Description: "Searches for areas that feature fmv sequences.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.areas = E.Areas.completeQueryParamsInit(params)
}

func initMonsterFormationsParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for monster-formations that feature the specified monster.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for monster-formations with the specified monster-formation-categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterFormationCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameMonsterFormationCategory},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for monster-formations with the given availabilities. If combined with the 'area' parameter, the availability of the monster-formation in this specific area is used. If a monster-formation has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster-formation won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with the 'area' parameter, it takes the repeatability directly from the monster-formations that occur in the specified area. Is combinable with 'availability'. In that case, the search looks for the monster-formation that is the most accessible while also being farmable and checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for monster-formations that appear within the specified location. If combined with 'availability', this parameter searches for monster-formations within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this location whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for monster-formations that appear within the specified sublocation. If combined with 'availability', this parameter searches for monster-formations within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monster-formations within this sublocation whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for monster-formations that appear within the specified area. If combined with 'availability', this parameter searches for monster-formations within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monster-formations within this area whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
		{
			Name:        ts.QpnAmbush,
			Description: "Searches for monster-formations that are forced ambushes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnEscape,
			Description: "Searches for monster-formations that can be escaped.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.monsterFormations = E.MonsterFormations.completeQueryParamsInit(params)
}

func initShopsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for shops with the specified shop categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ShopCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameShopCategory},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for shops with the given availabilities. By default, this parameter checks, if a shop simply is available at the given availability. If combined with a filter that refers to the shop's inventory, it takes the availability directly from there ('pre-story' for the inventory before acquiring the airship, and 'post' after). In that case, this filter looks, if the requested resources are available in a shop at that point in the game.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for shops that appear at the specified location. If combined with 'availability', this parameter searches for shops within this location whose availability matches one of the specified availabilities.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnAvailability},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for shops that appear at the specified sublocation. If combined with 'availability', this parameter searches for shops within this sublocation whose availability matches one of the specified availabilities.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for shops that offer equipment with the specified auto-ability. Can be combined with 'empty_slots' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          ts.QptID,
			ParamUse:      puList,
			ConflictsWith: []ts.QueryParamName{ts.QpnItems, ts.QpnEquipment},
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability},
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:            ts.QpnEmptySlots,
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots. Can be combined with 'auto_ability' and 'character' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:            ts.QptIntList,
			ParamUse:        puList,
			ConflictsWith:   []ts.QueryParamName{ts.QpnItems, ts.QpnEquipment},
			ReplacedBy:      []ts.QueryParamName{ts.QpnAvailability},
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          ts.QpnCharacter,
			Description:   "Searches for shops that offer equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ConflictsWith: []ts.QueryParamName{ts.QpnItems, ts.QpnEquipment},
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability},
			ReferencesInt: []ts.EndpointName{ts.EpCharacters},
		},
		{
			Name:          ts.QpnItems,
			Description:   "Searches for shops that offer items. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          ts.QptBool,
			ParamUse:      puList,
			ConflictsWith: []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnCharacter, ts.QpnEmptySlots},
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:          ts.QpnEquipment,
			Description:   "Searches for shops that offer equipment. If this query param is combined with 'availability', the availability of the shop's inventory is used ('pre-story' for pre-airship inventory and 'post' for post-airship inventory).",
			Type:          ts.QptBool,
			ParamUse:      puList,
			ConflictsWith: []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnCharacter, ts.QpnEmptySlots},
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability},
		},
	}

	Q.shops = E.Shops.completeQueryParamsInit(params)
}

func initTreasuresParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for treasures that appear within the specified location.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for treasures that appear within the specified sublocation.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for treasures that appear within the specified area.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for treasures that contain equipment with the specified auto-ability. Can be combined with 'empty_slots', 'character', and 'availability' for more specific searches.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:            ts.QpnEmptySlots,
			Description:     "Searches for treasures that contain equipment with the specified amounts of empty slots. Can be combined with 'auto_ability', 'character', and 'availability' for more specific searches.",
			Type:            ts.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:          ts.QpnCharacter,
			Description:   "Searches for treasures that contain equipment for the specified character. Can be combined with 'auto_ability', 'empty_slots', and 'availability' for more specific searches.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"wakka"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacters},
		},
		{
			Name:          ts.QpnItem,
			Description:   "Searches for treasures that contain the specified item.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:               ts.QpnLootType,
			Description:        "Searches for treasures with the specified loot type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.LootType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameLootType},
		},
		{
			Name:               ts.QpnTreasureType,
			Description:        "Searches for treasures with the specified treasure type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.TreasureType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTreasureType},
		},
		{
			Name:        ts.QpnAnima,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for treasures with the given availabilities.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
	}

	Q.treasures = E.Treasures.completeQueryParamsInit(params)
}

func initQuestsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnType,
			Description:        "Searches for quests that are of the specified quest type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.QuestType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameQuestType},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for quests with the given availabilities.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for quests that can be completed more than once.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.quests = E.Quests.completeQueryParamsInit(params)
}

func initSidequestsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for sidequests with the given availabilities.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
	}

	Q.sidequests = E.Sidequests.completeQueryParamsInit(params)
}

func initSubquestsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for subquests with the given availabilities.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for subquests that can be completed more than once.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.subquests = E.Subquests.completeQueryParamsInit(params)
}

func initArenaCreationsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for monster-formations with the specified arena-creation-categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ArenaCreationCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameArenaCreationCategory},
		},
	}

	Q.arenaCreations = E.ArenaCreations.completeQueryParamsInit(params)
}

func initBlitzballPrizesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.BlitzballTournamentCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameBlitzballTournamentCategory},
		},
	}

	Q.blitzballPrizes = E.BlitzballPrizes.completeQueryParamsInit(params)
}

func initSongsParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
		{
			Name:        ts.QpnFMVs,
			Description: "Searches for songs that are played in fmvs.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnSpecialUse,
			Description: "Searches for songs with a special use case.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnComposer,
			Description:        "Searches for songs that were composed by the stated composer.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.Composer.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameComposer},
		},
		{
			Name:               ts.QpnArranger,
			Description:        "Searches for songs that were arranged by the stated arranger.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.Arranger.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameArranger},
		},
	}

	Q.songs = E.Songs.completeQueryParamsInit(params)
}

func initFMVsParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for fmvs that are played within the specified location.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
	}

	Q.fmvs = E.FMVs.completeQueryParamsInit(params)
}

func initPlayerUnitsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnType,
			Description:        "Searches for player units that are of the specified unit type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.UnitType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameUnitType},
		},
	}

	Q.playerUnits = E.PlayerUnits.completeQueryParamsInit(params)
}

func initCharactersParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnOsgStats,
			Description:        "Adds all stat gains within the character's stated sphere grid to their base stats. This includes stat nodes behind sphere locks.",
			Type:               ts.QptEnum,
			ParamUse:           puSingle,
			EnumVals:           enums.T.SphereGridType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameSphereGridType},
		},
		{
			Name:        ts.QpnStoryBased,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnUnderwater,
			Description: "Searches for characters that can fight underwater.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.characters = E.Characters.completeQueryParamsInit(params)
}

func initAeonsParams() {
	params := []QueryParam{
		{
			Name:            ts.QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            ts.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:        ts.QpnYunaStats,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter.",
			Type:        ts.QptStat,
			ExampleUses: []string{exampleUse(ts.QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(ts.QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
		},
		{
			Name:        ts.QpnOptional,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.aeons = E.Aeons.completeQueryParamsInit(params)
}

func initCharacterClassesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for character classes with the specified categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.CharacterClassCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameCharacterClassCategory},
		},
	}

	Q.characterClasses = E.CharacterClasses.completeQueryParamsInit(params)
}

func initMonstersParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays a monster's related resources with the given availabilities. This affects areas and monster-formations.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays a monster's related resources that can be farmed. This affects areas and monster-formations.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        ts.QpnKimahriStats,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type:        ts.QptStat,
			ExampleUses: []string{exampleUse(ts.QpnKimahriStats, "hp:3000,strength:25,magic:30,agility:40"), exampleUse(ts.QpnKimahriStats, "hp:15000,agility:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:            ts.QpnBattles,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating a possessed aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type:            ts.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
			AllowedIDs:      []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        ts.QpnYunaStats,
			Description: "Calculate a possessed aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used for the calculation instead. Can be used in combination with the 'battles' parameter. All stats are replaceable, except for MP and Luck. While Yuna's MP and Luck still play a factor in the calculation of the other stats, a posssessed aeon's own MP and Luck will always be set to 1. If unused, the aeon's stats are based on your own aeon's default stats.",
			Type:        ts.QptStat,
			ExampleUses: []string{exampleUse(ts.QpnYunaStats, "hp:3000,strength:75,defense:50,magic:30,agility:20"), exampleUse(ts.QpnYunaStats, "accuracy:150,magic_defense:255")},
			ParamUse:    puSingle,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:        ts.QpnAlteredState,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:        ts.QptID,
			ParamUse:    puSingle,
		},
		{
			Name:          ts.QpnOmnisElements,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type:          ts.QptOther,
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{exampleUse(ts.QpnOmnisElements, "ifil"), exampleUse(ts.QpnOmnisElements, "wilf")},
			ParamUse:      puSingle,
			AllowedIDs:    []int32{211},
			AllowedValues: []ts.QueryValue{ts.QvF, ts.QvL, ts.QvW, ts.QvI},
		},
		{
			Name:               ts.QpnElementalResists,
			Description:        "Searches for monsters that have the specified elemental affinities.",
			Type:               ts.QptOther,
			Usage:              "?elemental_resists={element|id}:{affinity|id},...",
			ExampleUses:        []string{exampleUse(ts.QpnElementalResists, "fire:weak,water:absorb"), exampleUse(ts.QpnElementalResists, "1:3,2:4")},
			ParamUse:           puList,
			ReferencesInt:      []ts.EndpointName{ts.EpElements},
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameElementalAffinity},
		},
		{
			Name:          ts.QpnStatusResists,
			Description:   "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type:          ts.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:            ts.QpnResistance,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            ts.QptInt,
			ParamUse:        puList,
			RequiresAll:     []ts.QueryParamName{ts.QpnStatusResists},
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:          ts.QpnItem,
			Description:   "Searches for monsters that have the specified item as loot. Can be specified further with the 'methods' parameter.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Specifies the methods of acquisition for the 'item' parameter.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			RequiresAll:   []ts.QueryParamName{ts.QpnItem},
			AllowedValues: []ts.QueryValue{ts.QvSteal, ts.QvDrop, ts.QvBribe, ts.QvOther},
		},
		{
			Name:          ts.QpnAutoAbility,
			Description:   "Searches for monsters that drop the specified auto-ability.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:        ts.QpnIsForced,
			Description: "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAutoAbility},
		},
		{
			Name:            ts.QpnEmptySlots,
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type:            ts.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:          ts.QpnRonsoRage,
			Description:   "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpRonsoRages},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for monsters with the given availabilities. If combined with a geographical filter, it takes the availability directly from the monster-formations that occur in the specified location, sublocation, or area. If a monster has multiple availabilities, because there are multiple ways of encountering it (like via always-accessible random encounter and via scripted story-fight), this filter defines the most accessible version of it as its actual availability. In that case, the monster won't show up for the other availability types, even if it technically can have that availability, since it can be encountered easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for monsters that can be farmed. If this parameter is combined with a geographical filter, it takes the repeatability directly from the monster-formations that occur in the specified location, sublocation, or area. Is combinable with 'availability'. The availability assigned to the monster is from the monster-formation that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for monsters that appear within the specified location. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this location whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this location whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for monsters that appear within the specified sublocation. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this sublocation whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for monsters that appear within the specified area. If combined with 'availability', this parameter searches for monsters that are part of at least one monster-formation within this area whose most accessible availability matches one of the specified availabilities (based on the formation's encounter areas). If combined with 'repeatable', this parameter searches for monsters that are part of at least one monster-formation within this area whose farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
		{
			Name:            ts.QpnDistance,
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type:            ts.QptIntList,
			ParamUse:        puList,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:        ts.QpnCapture,
			Description: "Searches for monsters that can be captured.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnHasOverdrive,
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnUnderwater,
			Description: "Searches for monsters that are fought underwater.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnZombie,
			Description: "Searches for monsters that are zombies.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnSpecies,
			Description:        "Searches for monsters of the specified species.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterSpecies.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameMonsterSpecies},
		},
		{
			Name:               ts.QpnCreationArea,
			Description:        "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.CreationArea.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameCreationArea},
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for monsters that are of the specified monster-categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MonsterCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameMonsterCategory},
		},
	}

	Q.monsters = E.Monsters.completeQueryParamsInit(params)
}

func initAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnType,
			Description:        "Searches for abilities that are of the specified ability types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AbilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAbilityType},
		},
		{
			Name:        ts.QpnRank,
			Description: "Searches for abilities with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnCopycat,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnHelpBar,
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for abilities that can be used by the specified monster.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:               ts.QpnTargetType,
			Description:        "Searches for abilities with the specified target types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTargetType},
		},
		{
			Name:        ts.QpnUserAtk,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnDarkable,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnSilenceable,
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnReflectable,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnAttackType,
			Description:        "Searches for abilities with battle interactions of the specified attack types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAttackType},
		},
		{
			Name:               ts.QpnDamageType,
			Description:        "Searches for abilities that deal the specified types of damage.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageType},
		},
		{
			Name:               ts.QpnDamageFormula,
			Description:        "Searches for abilities that use the specified formula to calculate their damage.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageFormula},
		},
		{
			Name:        ts.QpnCanCrit,
			Description: "Searches for abilities that can land critical hits.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnBDL,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnElement,
			Description:   "Searches for abilities that deal elemental damage based on the specified element.",
			Type:          ts.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpElements},
		},
		{
			Name:        ts.QpnDelay,
			Description: "Searches for abilities that deal delay.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnStatusInflict,
			Description:   "Searches for abilities that can inflict the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:          ts.QpnStatusRemove,
			Description:   "Searches for abilities that can remove the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:        ts.QpnStatChanges,
			Description: "Searches for abilities that cause stat changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnModChanges,
			Description: "Searches for abilities that cause modifier changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.abilities = E.Abilities.completeQueryParamsInit(params)
}

func initPlayerAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnAbilityUser,
			Description:   "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []ts.EndpointName{ts.EpPlayerUnits},
		},
		{
			Name:        ts.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the player ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:        ts.QpnBombWpn,
			Description: "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:        ts.QpnRank,
			Description: "Searches for player abilities with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnCopycat,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnHelpBar,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for player abilities that are of the specified player ability categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.PlayerAbilityCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNamePlayerAbilityCategory},
		},
		{
			Name:        ts.QpnOutsideBattle,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMp,
			Description: "Searches for player abilities with the specified mp costs.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMpMin,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type:        ts.QptInt,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnMpMax,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type:        ts.QptInt,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnRelatedStat,
			Description:   "Searches for player abilities that are related to the specified stat.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStats},
		},
		{
			Name:          ts.QpnUser,
			Description:   "Searches for player abilities that are learned by the specified character class.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacterClasses},
		},
		{
			Name:          ts.QpnStdSg,
			Description:   "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacters},
		},
		{
			Name:          ts.QpnExpSg,
			Description:   "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"tidus", "wakka"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacters},
		},
		{
			Name:          ts.QpnLearnItem,
			Description:   "Searches for player abilities an aeon can learn via the specified item.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:               ts.QpnTargetType,
			Description:        "Searches for player abilities with the specified target types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTargetType},
		},
		{
			Name:        ts.QpnUserAtk,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnDarkable,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnSilenceable,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnReflectable,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnAttackType,
			Description:        "Searches for player abilities with battle interactions of the specified attack types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAttackType},
		},
		{
			Name:               ts.QpnDamageType,
			Description:        "Searches for player abilities that deal the specified types of damage.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageType},
		},
		{
			Name:               ts.QpnDamageFormula,
			Description:        "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageFormula},
		},
		{
			Name:          ts.QpnElement,
			Description:   "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:          ts.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpElements},
		},
		{
			Name:        ts.QpnDelay,
			Description: "Searches for player abilities that deal delay.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnStatusInflict,
			Description:   "Searches for player abilities that can inflict the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:          ts.QpnStatusRemove,
			Description:   "Searches for player abilities that can remove the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:        ts.QpnStatChanges,
			Description: "Searches for player abilities that cause stat changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnModChanges,
			Description: "Searches for player abilities that cause modifier changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.playerAbilities = E.PlayerAbilities.completeQueryParamsInit(params)
}

func initOverdriveAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        ts.QpnRank,
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnUser,
			Description:   "Searches for overdrive abilities that are learned by the specified character class.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacterClasses},
		},
		{
			Name:          ts.QpnRelatedStat,
			Description:   "Searches for overdrive abilities that are related to the specified stat.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStats},
		},
		{
			Name:               ts.QpnTargetType,
			Description:        "Searches for overdrive abilities with the specified target types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTargetType},
		},
		{
			Name:               ts.QpnAttackType,
			Description:        "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAttackType},
		},
		{
			Name:               ts.QpnDamageFormula,
			Description:        "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageFormula},
		},
		{
			Name:        ts.QpnCanCrit,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnElement,
			Description:   "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:          ts.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpElements},
		},
		{
			Name:        ts.QpnDelay,
			Description: "Searches for overdrive abilities that deal delay.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnStatusInflict,
			Description:   "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:          ts.QpnStatusRemove,
			Description:   "Searches for overdrive abilities that can remove the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:        ts.QpnStatChanges,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnModChanges,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.overdriveAbilities = E.OverdriveAbilities.completeQueryParamsInit(params)
}

func initItemAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for item abilities that are of the specified item categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameItemCategory},
		},
		{
			Name:        ts.QpnOutsideBattle,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnRelatedStat,
			Description:   "Searches for item abilities that are related to the specified stat.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStats},
		},
		{
			Name:               ts.QpnTargetType,
			Description:        "Searches for item abilities with the specified target types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTargetType},
		},
		{
			Name:               ts.QpnAttackType,
			Description:        "Searches for item abilities with battle interactions of the specified attack types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAttackType},
		},
		{
			Name:               ts.QpnDamageFormula,
			Description:        "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageFormula},
		},
		{
			Name:          ts.QpnElement,
			Description:   "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:          ts.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpElements},
		},
		{
			Name:        ts.QpnDelay,
			Description: "Searches for item abilities that deal delay.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnStatusInflict,
			Description:   "Searches for item abilities that can inflict the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:          ts.QpnStatusRemove,
			Description:   "Searches for item abilities that can remove the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:        ts.QpnStatChanges,
			Description: "Searches for item abilities that cause stat changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnModChanges,
			Description: "Searches for item abilities that cause modifier changes.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.itemAbilities = E.ItemAbilities.completeQueryParamsInit(params)
}

func initTriggerCommandsParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnAbilityUser,
			Description:   "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []ts.EndpointName{ts.EpPlayerUnits},
		},
		{
			Name:        ts.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the trigger command's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:        ts.QpnBombWpn,
			Description: "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:          ts.QpnRelatedStat,
			Description:   "Searches for trigger commands that are related to the specified stat.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStats},
		},
		{
			Name:          ts.QpnUser,
			Description:   "Searches for trigger commands that are learned by the specified character class.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacterClasses},
		},
	}

	Q.triggerCommands = E.TriggerCommands.completeQueryParamsInit(params)
}

func initMiscAbilitiesParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnAbilityUser,
			Description:   "If a misc ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"wakka", "valefor"},
			ParamUse:      puSingle,
			ReferencesInt: []ts.EndpointName{ts.EpPlayerUnits},
		},
		{
			Name:        ts.QpnCelestialWeapon,
			Description: "Can only be used in combination with the 'ability_user' parameter. If this parameter is true, the misc ability's damage formula will be replaced by the character's celestial weapon damage formula.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:        ts.QpnBombWpn,
			Description: "If a misc ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
			RequiresAll: []ts.QueryParamName{ts.QpnAbilityUser},
		},
		{
			Name:        ts.QpnRank,
			Description: "Searches for misc abilities with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnCopycat,
			Description: "Searches for misc abilities that can be copied by 'copycat'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnHelpBar,
			Description: "Searches for misc abilities whose names appear in the help bar.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnUser,
			Description:   "Searches for misc abilities that are learned by the specified character class.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacterClasses},
		},
		{
			Name:        ts.QpnUserAtk,
			Description: "Searches for misc abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.miscAbilities = E.MiscAbilities.completeQueryParamsInit(params)
}

func initEnemyAbilitiesParams() {
	params := []QueryParam{
		{
			Name:        ts.QpnRank,
			Description: "Searches for enemy abilities with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnHelpBar,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for enemy abilities that can be used by the specified monster.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:               ts.QpnTargetType,
			Description:        "Searches for enemy abilities with the specified target types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.TargetType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameTargetType},
		},
		{
			Name:        ts.QpnDarkable,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnSilenceable,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnReflectable,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnAttackType,
			Description:        "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AttackType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAttackType},
		},
		{
			Name:               ts.QpnDamageType,
			Description:        "Searches for enemy abilities that deal the specified types of damage.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageType},
		},
		{
			Name:               ts.QpnDamageFormula,
			Description:        "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.DamageFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameDamageFormula},
		},
		{
			Name:        ts.QpnCanCrit,
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:        ts.QpnBDL,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnElement,
			Description:   "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:          ts.QptNameIDListNul,
			ExampleVals:   []string{"fire", "ice"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpElements},
		},
		{
			Name:        ts.QpnDelay,
			Description: "Searches for enemy abilities that deal delay.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnStatusInflict,
			Description:   "Searches for enemy abilities that can inflict the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
		{
			Name:          ts.QpnStatusRemove,
			Description:   "Searches for enemy abilities that can remove the specified status condition.",
			Type:          ts.QptIDNul,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStatusConditions},
		},
	}

	Q.enemyAbilities = E.EnemyAbilities.completeQueryParamsInit(params)
}

func initOverdrivesParams() {
	params := []QueryParam{
		{
			Name:        ts.QpnRank,
			Description: "Searches for overdrives with the specified ranks.",
			Type:        ts.QptIntList,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnUser,
			Description:   "Searches for overdrives that are learned by the specified character class.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"characters", "tidus"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacterClasses},
		},
	}

	Q.overdrives = E.Overdrives.completeQueryParamsInit(params)
}

func initSubmenusParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnTopmenu,
			Description:   "Searches for submenus that are found within the specified topmenu.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"main", "left"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpTopmenus},
		},
	}

	Q.submenus = E.Submenus.completeQueryParamsInit(params)
}

func initAllItemsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only considers an item's related resources with the given availabilities when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only considers an item's related resources that can be farmed when calculating the boolean fields. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               ts.QpnType,
			Description:        "Searches for items that are of the specified item-types.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameItemType},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
	}

	Q.allItems = E.AllItems.completeQueryParamsInit(params)
}

func initItemsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays an item's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays an item's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:        ts.QpnHasAbility,
			Description: "Searches for items that can be used in battle.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnRelatedStat,
			Description:   "Searches for items that are related to the specified stat.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"hp", "strength"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpStats},
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for items that are from one of the specified item categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ItemCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameItemCategory},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for items with the given availabilities. The availability of an item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the item. The item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for items that can be farmed. Is combinable with 'availability'. The availability assigned to the item is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Searches for items that can be obtained via at least one of the given methods.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for items that can be obtained at the specified location. If combined with 'availability', this parameter searches for items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this location whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for items that can be obtained in the specified area. If combined with 'availability', this parameter searches for items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for items within this area whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
	}

	Q.items = E.Items.completeQueryParamsInit(params)
}

func initKeyItemsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays a key-item's related resources with the given availabilities. This affects areas, treasures and quests.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for key-items with the given availabilities. The availability of a key-item is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the key-item. The key-item won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for key-items that are of the specified key-item categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.KeyItemCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameKeyItemCategory},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Searches for key-items that can be obtained via at least one of the given methods.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			AllowedValues: []ts.QueryValue{ts.QvTreasure, ts.QvQuest},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for key-items that can be obtained at the specified location. If combined with 'availability', this parameter searches for key-items within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this location whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for key-items that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for key-items within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for key-items that can be obtained in the specified area. If combined with 'availability', this parameter searches for key-items within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for key-items within this area whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
	}

	Q.keyItems = E.KeyItems.completeQueryParamsInit(params)
}

func initSpheresParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays a sphere's related resources with the given availabilities. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays a sphere's related resources that can be farmed. This affects monsters, treasures, shops, quests, and blitzball prizes.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               ts.QpnColor,
			Description:        "Searches for spheres with any of the given colors.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.SphereColor.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameSphereColor},
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for spheres with the given availabilities. The availability of a sphere is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the sphere. The sphere won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for spheres that can be farmed. Is combinable with 'availability'. The availability assigned to the sphere is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Searches for spheres that can be obtained via at least one of the given methods.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop, ts.QvQuest, ts.QvBlitzball},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for spheres that can be obtained at the specified location. If combined with 'availability', this parameter searches for spheres within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this location whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for spheres that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for spheres within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for spheres that can be obtained in the specified area. If combined with 'availability', this parameter searches for spheres within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for spheres within this area whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
	}

	Q.spheres = E.Spheres.completeQueryParamsInit(params)
}

func initPrimersParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for primers with the given availabilities. The availability of a primer is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the primer. The primer won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
	}

	Q.primers = E.Primers.completeQueryParamsInit(params)
}

func initMixesParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnContainsItem,
			Description:   "Modifies combinations to only display item combinations that include the specified item.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puSingle,
			ConflictsWith: []ts.QueryParamName{ts.QpnBest},
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnBest,
			Description:   "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:          ts.QptBool,
			ParamUse:      puSingle,
			ConflictsWith: []ts.QueryParamName{ts.QpnContainsItem},
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for mixes that are of the specified mix categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.MixCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameMixCategory},
		},
		{
			Name:          ts.QpnReqItem,
			Description:   "Searches for mixes that can be built with the specified item.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
		{
			Name:          ts.QpnSecondItem,
			Description:   "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"grenade", "power_sphere"},
			ParamUse:      puList,
			RequiresAll:   []ts.QueryParamName{ts.QpnReqItem},
			ReferencesInt: []ts.EndpointName{ts.EpItems},
		},
	}

	Q.mixes = E.Mixes.completeQueryParamsInit(params)
}

func initAutoAbilitiesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays an auto-ability's related resources with the given availabilities. This affects shops, treasures, and monsters.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnRelRepeatable,
			Description: "Only displays an auto-ability's related resources that can be farmed. This affects shops, treasures, and monsters.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
		{
			Name:               ts.QpnAvailability,
			Description:        "Searches for auto-abilities with the given availabilities. The availability of an auto-ability is always taken from its sources. The most accessible availability among those sources is the one that is assigned to the auto-ability. The auto-ability won't show up for the other availability types, even if it technically can have that availability, since it can be received easier. It is recommended to use the joined availability values ('story', 'post-game', 'pre-airship', 'post-airship') to get a full picture of your options.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:        ts.QpnPreAirship,
			Description: "Makes the 'availability' filter view availability 'pre-story' as more accessible than availability 'post', if both types are present for one resource. This is useful, when you're not in the post-game yet and therefore can't access 'post' resources.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresAll: []ts.QueryParamName{ts.QpnAvailability},
		},
		{
			Name:        ts.QpnRepeatable,
			Description: "Searches for auto-abilities that can be farmed. Is combinable with 'availability'. The availability assigned to the auto-ability is from the source that is the most accessible while also being farmable. The query then checks, if this availability matches the given availabilities. It can be that more results show up at less accessible availability values than without using 'repeatable', because the more accessible sources aren't farmable (like an always accessible equipment treasure vs. a story-exclusive monster encounter).",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.AutoAbilityCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAutoAbilityCategory},
		},
		{
			Name:               ts.QpnType,
			Description:        "Searches for auto-abilities that are of the specified equip type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameEquipType},
		},
		{
			Name:          ts.QpnMonster,
			Description:   "Searches for auto-abilities that are dropped by the specified monster.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnMonsterItems,
			Description:   "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnShop,
			Description:   "Searches for auto-abilities that can be obtained from the specified shop.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnCharacter,
			Description:   "Restricts the search for 'availability', 'monster' and 'shop' to only include auto-abilities that can be obtained by the specified character. This includes auto-abilities with no character restriction like regular monster equipment drop slots.",
			Type:          ts.QptNameID,
			ExampleVals:   []string{"kimahri"},
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			RequiresOne:   []ts.QueryParamName{ts.QpnAvailability, ts.QpnMonster, ts.QpnShop},
			ReferencesInt: []ts.EndpointName{ts.EpMonsters},
		},
		{
			Name:          ts.QpnMethods,
			Description:   "Searches for auto-abilities that can be obtained via at least one of the given methods.",
			Type:          ts.QptValueList,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			AllowedValues: []ts.QueryValue{ts.QvMonster, ts.QvTreasure, ts.QvShop},
		},
		{
			Name:        ts.QpnCustomize,
			Description: "Converts the 'availability' and 'repeatable' parameters to search for auto-abilities based on their required item's availability and/or farmability.",
			Type:        ts.QptBool,
			ParamUse:    puList,
			RequiresOne: []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
		},
		{
			Name:          ts.QpnLocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified location. If combined with 'availability', this parameter searches for auto-abilities within this location whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this location whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnSublocation, ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpLocations},
		},
		{
			Name:          ts.QpnSublocation,
			Description:   "Searches for auto-abilities that can be obtained at the specified sublocation. If combined with 'availability', this parameter searches for auto-abilities within this sublocation whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this sublocation whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnArea, ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpSublocations},
		},
		{
			Name:          ts.QpnArea,
			Description:   "Searches for auto-abilities that can be obtained in the specified area. If combined with 'availability', this parameter searches for auto-abilities within this area whose sources' most accessible availability matches one of the specified availabilities. If combined with 'repeatable', this parameter searches for auto-abilities within this area whose sources' farmability matches the given value based on its category.",
			Type:          ts.QptID,
			ParamUse:      puList,
			ReplacedBy:    []ts.QueryParamName{ts.QpnAvailability, ts.QpnRepeatable},
			ReferencesInt: []ts.EndpointName{ts.EpAreas},
		},
	}

	Q.autoAbilities = E.AutoAbilities.completeQueryParamsInit(params)
}

func initEquipmentTablesParams() {
	params := []QueryParam{
		{
			Name:          ts.QpnAutoAbilities,
			Description:   "Searches for equipment tables with all of the given auto-abilities.",
			Type:          ts.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:               ts.QpnType,
			Description:        "Searches for equipment tables that are of the specified equip type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameEquipType},
		},
		{
			Name:        ts.QpnCelestialWeapon,
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.equipmentTables = E.EquipmentTables.completeQueryParamsInit(params)
}

func initEquipmentParams() {
	params := []QueryParam{
		{
			Name:        ts.QpnTable,
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type:        ts.QptInt,
			ParamUse:    puSingle,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:               ts.QpnRelAvailability,
			Description:        "Only displays an equipment's related resources with the given availabilities. This affects treasures and shops.",
			Type:               ts.QptEnumList,
			ParamUse:           puSingle,
			EnumVals:           enums.T.AvailabilityType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameAvailabilityType},
		},
		{
			Name:          ts.QpnAutoAbilities,
			Description:   "Searches for equipment with all of the given auto-abilities.",
			Type:          ts.QptIDList,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpAutoAbilities},
		},
		{
			Name:          ts.QpnCharacter,
			Description:   "Searches for equipment of the specified character.",
			ExampleVals:   []string{"yuna"},
			Type:          ts.QptNameID,
			ParamUse:      puList,
			ReferencesInt: []ts.EndpointName{ts.EpCharacters},
		},
		{
			Name:               ts.QpnType,
			Description:        "Searches for equipment that is of the specified equip type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.EquipType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameEquipType},
		},
		{
			Name:        ts.QpnCelestialWeapon,
			Description: "Searches for the celestial weapons.",
			Type:        ts.QptBool,
			ParamUse:    puList,
		},
	}

	Q.equipment = E.Equipment.completeQueryParamsInit(params)
}

func initCelestialWeaponsParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnFormula,
			Description:        "Searches for celestial-weapons that are of the specified celestial formula.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.CelestialFormula.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameCelestialFormula},
		},
	}

	Q.celestialWeapons = E.CelestialWeapons.completeQueryParamsInit(params)
}

func initStatsParams() {
	params := []QueryParam{
		{
			Name:        ts.QpnChangesOnly,
			Description: "Only includes a stat's related auto-abilities, abilities, status conditions, and properties that cause stat changes.",
			Type:        ts.QptBool,
			ParamUse:    puSingle,
		},
	}

	Q.stats = E.Stats.completeQueryParamsInit(params)
}

func initOverdriveModesParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnType,
			Description:        "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:               ts.QptEnum,
			ParamUse:           puList,
			EnumVals:           enums.T.OverdriveModeType.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameOverdriveModeType},
		},
	}

	Q.overdriveModes = E.OverdriveModes.completeQueryParamsInit(params)
}

func initStatusConditionsParams() {
	params := []QueryParam{
		{
			Name:            ts.QpnInflictMin,
			Description:     "Only shows a status condition's related abilities with an infliction rate higher than or equal to the given amount. The default value is '1'. Can be combined with 'inflict_max', but can't be higher. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            ts.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvInfinite,
					Val: 254,
				},
				{
					Key: ts.QsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:            ts.QpnInflictMax,
			Description:     "Only shows a status condition's related abilities with an infliction rate lower than or equal to the given amount. The default value is '25'. Can be combined with 'inflict_min', but can't be lower. Special values are 'infinite' (=254) and 'always' (=255).",
			Type:            ts.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 255},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvInfinite,
					Val: 254,
				},
				{
					Key: ts.QsvAlways,
					Val: 255,
				},
			},
			DefaultVal: h.GetIntPtr(255),
		},
		{
			Name:            ts.QpnResistance,
			Description:     "Only shows a status condition's related monsters with a resistance higher than or equal to the given amount. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type:            ts.QptInt,
			ParamUse:        puSingle,
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialQueryInput{
				{
					Key: ts.QsvImmune,
					Val: 254,
				},
			},
			DefaultVal: h.GetIntPtr(1),
		},
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for status conditions that are of the specified status condition categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.StatusConditionCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameStatusConditionCategory},
		},
	}

	Q.statusConditions = E.StatusConditions.completeQueryParamsInit(params)
}

func initModifiersParams() {
	params := []QueryParam{
		{
			Name:               ts.QpnCategory,
			Description:        "Searches for modifiers that are of the specified modifier categories.",
			Type:               ts.QptEnumList,
			ParamUse:           puList,
			EnumVals:           enums.T.ModifierCategory.Vals(),
			ReferencesEnumsInt: []ts.EnumName{ts.EnumNameModifierCategory},
		},
	}

	Q.modifiers = E.Modifiers.completeQueryParamsInit(params)
}

func initAgilityTierParams() {
	params := []QueryParam{
		{
			Name:            ts.QpnAgility,
			Description:     "Searches for the agility tier that the given agility value belongs to.",
			Type:            ts.QptInt,
			ParamUse:        puList,
			AllowedIntRange: []int{0, 255},
		},
	}

	Q.agilityTiers = E.AgilityTiers.completeQueryParamsInit(params)
}
