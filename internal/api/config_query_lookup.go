package api

import (
	"fmt"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QueryParam struct {
	ID               int                        `json:"-"`
	Name             string                     `json:"name"`
	Type			 string						`json:"param_type"`
	Description      string                     `json:"description"`
	ExampleVals		 []string					`json:"-"`
	Usage            string                     `json:"usage"`
	ExampleUses      []string                   `json:"example_uses"`
	DefaultOnly      bool                       `json:"only_use_alone"`
	ForSingle        bool                       `json:"for_single"`
	ForList          bool                       `json:"for_list"`
	ForSegment       *string                    `json:"for_segment"`
	IsRequired       bool                       `json:"is_required"`
	TypeLookup       map[string]EnumAPIResource `json:"-"`
	RequiredParams   []string                   `json:"required_params,omitempty"`
	UsableWith       []string                   `json:"usable_with,omitempty"`
	ForbiddenParams	 []string					`json:"forbidden_params,omitempty"`
	References       []string                   `json:"references,omitempty"`
	AllowedIDs       []int32                    `json:"-"`
	AllowedResources []string                   `json:"allowed_resources,omitempty"`
	AllowedValues    []string                   `json:"allowed_values,omitempty"`
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
	defaultParams map[string]QueryParam
	
	locations    map[string]QueryParam
	sublocations map[string]QueryParam
	areas        map[string]QueryParam

	monsterFormations map[string]QueryParam
	shops             map[string]QueryParam
	treasures         map[string]QueryParam
	quests            map[string]QueryParam
	sidequests        map[string]QueryParam
	subquests         map[string]QueryParam
	arenaCreations    map[string]QueryParam
	blitzballPrizes   map[string]QueryParam
	songs             map[string]QueryParam
	fmvs              map[string]QueryParam

	playerUnits      map[string]QueryParam
	characters       map[string]QueryParam
	aeons            map[string]QueryParam
	characterClasses map[string]QueryParam
	monsters         map[string]QueryParam

	abilities            map[string]QueryParam
	playerAbilities      map[string]QueryParam
	overdriveAbilities   map[string]QueryParam
	itemAbilities        map[string]QueryParam
	triggerCommands      map[string]QueryParam
	unspecifiedAbilities map[string]QueryParam
	enemyAbilities       map[string]QueryParam

	aeonCommands      map[string]QueryParam
	overdriveCommands map[string]QueryParam
	overdrives        map[string]QueryParam
	ronsoRages        map[string]QueryParam
	submenus          map[string]QueryParam
	topmenus          map[string]QueryParam

	allItems		map[string]QueryParam
	items 			map[string]QueryParam
	keyItems 		map[string]QueryParam
	spheres 		map[string]QueryParam
	primers			map[string]QueryParam
	mixes			map[string]QueryParam

	autoAbilities	map[string]QueryParam
	equipmentTables	map[string]QueryParam
	equipment		map[string]QueryParam

	overdriveModes 	map[string]QueryParam
}

func (cfg *Config) QueryLookupInit() {
	cfg.q = &QueryLookup{}

	defaultParams := []QueryParam{
		{
			Name:		 "limit",
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Type:		 "int",
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
			Name:		 "offset",
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Type: 		 "int",
			ForList:     true,
			ForSingle:   false,
			DefaultVal:  h.GetIntPtr(0),
		},
		{
			Name:		 "flip",
			Description: "Flips the filtered results in a list response and returns the negative.",
			Type: 		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	cfg.q.defaultParams = querySliceToMap(cfg, defaultParams)
	cfg.initLocationsParams(defaultParams)
	cfg.initSublocationsParams(defaultParams)
	cfg.initAreasParams(defaultParams)

	cfg.initMonsterFormationsParams(defaultParams)
	cfg.initShopsParams(defaultParams)
	cfg.initTreasuresParams(defaultParams)
	cfg.initQuestsParams(defaultParams)
	cfg.initSidequestsParams(defaultParams)
	cfg.initSubquestsParams(defaultParams)
	cfg.initArenaCreationsParams(defaultParams)
	cfg.initBlitzballPrizesParams(defaultParams)
	cfg.initSongsParams(defaultParams)
	cfg.initFMVsParams(defaultParams)

	cfg.initPlayerUnitsParams(defaultParams)
	cfg.initCharactersParams(defaultParams)
	cfg.initAeonsParams(defaultParams)
	cfg.initCharacterClassesParams(defaultParams)
	cfg.initMonstersParams(defaultParams)

	cfg.initAbilitiesParams(defaultParams)
	cfg.initPlayerAbilitiesParams(defaultParams)
	cfg.initOverdriveAbilitiesParams(defaultParams)
	cfg.initItemAbilitiesParams(defaultParams)
	cfg.initTriggerCommandsParams(defaultParams)
	cfg.initUnspecifiedAbilitiesParams(defaultParams)
	cfg.initEnemyAbilitiesParams(defaultParams)

	cfg.q.aeonCommands = cfg.assignDefaultParams(defaultParams)
	cfg.q.overdriveCommands = cfg.assignDefaultParams(defaultParams)
	cfg.initOverdrivesParams(defaultParams)
	cfg.q.ronsoRages = cfg.assignDefaultParams(defaultParams)
	cfg.initSubmenusParams(defaultParams)
	cfg.q.topmenus = cfg.assignDefaultParams(defaultParams)

	cfg.initAllItemsParams(defaultParams)
	cfg.initItemsParams(defaultParams)
	cfg.initKeyItemsParams(defaultParams)
	cfg.initSpheresParams(defaultParams)
	cfg.initPrimersParams(defaultParams)
	cfg.initMixesParams(defaultParams)

	cfg.initAutoAbilitiesParams(defaultParams)
	cfg.initEquipmentTablesParams(defaultParams)
	cfg.initEquipmentParams(defaultParams)

	cfg.initOverdriveModesParams(defaultParams)

}

func (cfg *Config) assignDefaultParams(defaultParams []QueryParam) map[string]QueryParam {
	return cfg.completeQueryParamsInit([]QueryParam{}, defaultParams, false)
}


func (cfg *Config) completeQueryParamsInit(params, defaultParams []QueryParam, hasSimpleView bool) map[string]QueryParam {
	params = slices.Concat(params, defaultParams)

	if hasSimpleView {
		queryParamIDs := QueryParam{
			Name: 		 "ids",
			Description: "Used to input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Type: 		 "id-list",
			DefaultOnly: true,
			ForList:     false,
			ForSingle:   false,
			ForSegment:  h.GetStrPtr("simple"),
		}
		params = append(params, queryParamIDs)
	}

	return querySliceToMap(cfg, params)
}


func (cfg *Config) assignParamUsage(p QueryParam) QueryParam {
	s := fmt.Sprintf("?%s=", p.Name)

	switch p.Type {
	case "bool":
		p.Usage = s + "{bool}"
		p.ExampleUses = []string{s + "true", s + "false"}

	case "enum":
		enums := createEnumResourceSlice(cfg, "", p.TypeLookup)
		e := enums[0].Name
		p.Usage = s + "{value|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case "enum-list":
		enums := createEnumResourceSlice(cfg, "", p.TypeLookup)
		e1 := enums[0].Name
		e2 := enums[1].Name
		p.Usage = s + "{value|id},..."
		p.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case "id":
		p.Usage = s + "{id}"
		p.ExampleUses = []string{s + "1"}

	case "id-nul":
		p.Usage = s + "{id|'none'}"
		p.ExampleUses = []string{s + "1", s + "none"}

	case "id-list":
		p.Usage = s + "{id},..."
		p.ExampleUses = []string{s + "1", s + "1,2"}

	case "int":
		p.Usage = s + "{int}"
		p.ExampleUses = []string{s + "1"}

	case "int-list":
		p.Usage = s + "{int},...|{int}-{int}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}

	case "name/id":
		e := p.ExampleVals[0]
		p.Usage = s + "{name|id}"
		p.ExampleUses = []string{s + "1", s + e}

	case "name/id-list":
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},..."
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}

	case "name/id-list-nul":
		e1 := p.ExampleVals[0]
		e2 := p.ExampleVals[1]
		p.Usage = s + "{name|id},...|{'none'}"
		p.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}

	case "value-list":
		e1 := p.AllowedValues[0]
		e2 := p.AllowedValues[1]
		p.Usage = s + "{val}"
		p.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}

	case "stat":
		p.Usage = s + "{stat}={int},..."

	default:
		return p
	}

	if p.SpecialInputs != nil {
		for _, input := range p.SpecialInputs {
			usageTrimmed := strings.TrimSuffix(p.Usage, "}")
			p.Usage = fmt.Sprintf("%s|'%s'}", usageTrimmed, input.Key)
			p.ExampleUses = append(p.ExampleUses, s + input.Key)
		}
	}

	return p
}


func (cfg *Config) initLocationsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "item",
			Description: "Searches for locations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 	"method",
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:		 	"value-list",
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		{
			Name:		 "key_item",
			Description: "Searches for locations where the specified key item can be acquired.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		{
			Name:		 "characters",
			Description: "Searches for locations where a character permanently joins the party.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "aeons",
			Description: "Searches for locations where a new aeon is acquired.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "monsters",
			Description: "Searches for locations that have monsters.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "boss_fights",
			Description: "Searches for locations that have bosses.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "shops",
			Description: "Searches for locations that have shops.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "treasures",
			Description: "Searches for locations that have treasures.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "sidequests",
			Description: "Searchces for locations that feature sidequests.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "fmvs",
			Description: "Searches for locations that feature fmv sequences.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.locations = paramsMap
}

func (cfg *Config) initSublocationsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "location",
			Description: "Searches for sublocations that are located within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "item",
			Description: "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 	"method",
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:		 	"value-list",
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		{
			Name:		 "key_item",
			Description: "Searches for sublocations where the specified key item can be acquired.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		{
			Name:		 "characters",
			Description: "Searches for sublocations where a character permanently joins the party.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "aeons",
			Description: "Searches for sublocations where a new aeon is acquired.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "monsters",
			Description: "Searches for sublocations that have monsters.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "boss_fights",
			Description: "Searches for sublocations that have bosses.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "shops",
			Description: "Searches for sublocations that have shops.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "treasures",
			Description: "Searches for sublocations that have treasures.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "sidequests",
			Description: "Searchces for sublocations that feature sidequests.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "fmvs",
			Description: "Searches for sublocations that feature fmv sequences.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.sublocations = paramsMap
}

func (cfg *Config) initAreasParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "location",
			Description: "Searches for areas that are located within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for areas that are located within the specified sublocation.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "item",
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 	"method",
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:			"value-list",
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		{
			Name:		 "key_item",
			Description: "Searches for areas where the specified key item can be acquired.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		{
			Name:		 "availability",
			Description: "Searches for areas with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "save_sphere",
			Description: "Searches for areas that have a save sphere.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "comp_sphere",
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "airship",
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "chocobo",
			Description: "Searches for areas where you can ride a chocobo.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "characters",
			Description: "Searches for areas where a character permanently joins the party.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "aeons",
			Description: "Searches for areas where a new aeon is acquired.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "monsters",
			Description: "Searches for areas that have monsters.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "boss_fights",
			Description: "Searches for areas that have bosses.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "shops",
			Description: "Searches for areas that have shops.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "treasures",
			Description: "Searches for areas that have treasures.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "sidequests",
			Description: "Searchces for areas that feature sidequests.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "fmvs",
			Description: "Searches for areas that feature fmv sequences.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.areas = paramsMap
}

func (cfg *Config) initMonsterFormationsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "monster",
			Description: "Searches for monster formations that feature the specified monster.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		{
			Name:		 "category",
			Description: "Searches for monster formations with the specified monster-formation-categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterFormationCategory.lookup,
			References:  []string{createListURL(cfg, "monster-formation-category")},
		},
		{
			Name:		 "location",
			Description: "Searches for monster formations that appear within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for monster formations that appear within the specified sublocation.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "area",
			Description: "Searches for monster formations that appear within the specified area.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		{
			Name:		 "availability",
			Description: "Searches for monster formations with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Searches for monster formations that can be triggered more than once.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "ambush",
			Description: "Searches for monster formations that are forced ambushes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.monsterFormations = paramsMap
}

func (cfg *Config) initShopsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "category",
			Description: "Searches for shops with the specified shop categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ShopCategory.lookup,
			References:  []string{createListURL(cfg, "shop-category")},
		},
		{
			Name:		 "location",
			Description: "Searches for shops that appear within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for shops that appear within the specified sublocation.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "auto_ability",
			Description: "Searches for shops that offer equipment with the specified auto-ability.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		{
			Name:		 "empty_slots",
			Description:     "Searches for shops that offer equipment with the specified amounts of empty slots.",
			Type: 		     "int-list",
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{0, 4},
		},
		{
			Name:		 "character",
			Description: "Specifies the character the offered equipment is for when searching for shops with the 'auto_ability' or 'empty_slots' parameters.",
			Type:		 "name/id",
			ExampleVals: []string{"wakka"},
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []string{"auto_ability", "empty_slots"},
			References:  []string{createListURL(cfg, "characters")},
		},
		{
			Name:		 "shop_type",
			Description: "Specifies whether the given auto-ability is sold before or after acquiring the airship when searching for shops with the 'auto_ability' or 'empty_slots' parameters.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			UsableWith:  []string{"auto_ability", "empty_slots"},
			TypeLookup:  cfg.t.ShopType.lookup,
		},
		{
			Name:		 "items",
			Description: "Searches for shops that offer items.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "equipment",
			Description: "Searches for shops that offer equipment.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "availability",
			Description: "Searches for shops with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.shops = paramsMap
}

func (cfg *Config) initTreasuresParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "location",
			Description: "Searches for treasures that appear within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for treasures that appear within the specified sublocation.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "area",
			Description: "Searches for treasures that appear within the specified area.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		{
			Name:		 "loot_type",
			Description: "Searches for treasures with the specified loot type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.LootType.lookup,
			References:  []string{createListURL(cfg, "loot-type")},
		},
		{
			Name:		 "treasure_type",
			Description: "Searches for treasures with the specified treasure type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TreasureType.lookup,
		},
		{
			Name:		 "anima",
			Description: "Searches for treasures that are necessary for getting Anima.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "availability",
			Description: "Searches for treasures with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.treasures = paramsMap
}

func (cfg *Config) initQuestsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "type",
			Description: "Searches for quests that are of the specified quest type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.QuestType.lookup,
			References:  []string{createListURL(cfg, "quest-type")},
		},
		{
			Name:		 "availability",
			Description: "Searches for quests with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Searches for quests that can be completed more than once.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.quests = paramsMap
}

func (cfg *Config) initSidequestsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "availability",
			Description: "Searches for sidequests with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.sidequests = paramsMap
}

func (cfg *Config) initSubquestsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "availability",
			Description: "Searches for subquests with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Searches for subquests that can be completed more than once.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.subquests = paramsMap
}

func (cfg *Config) initArenaCreationsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "category",
			Description: "Searches for monster formations with the specified arena-creation-categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ArenaCreationCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.arenaCreations = paramsMap
}

func (cfg *Config) initBlitzballPrizesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "category",
			Description: "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.BlitzballTournamentCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.blitzballPrizes = paramsMap
}

func (cfg *Config) initSongsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "location",
			Description: "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "area",
			Description: "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		{
			Name:		 "fmvs",
			Description: "Searches for songs that are played in fmvs.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "special_use",
			Description: "Searches for songs with a special use case.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "composer",
			Description: "Searches for songs that were composed by the stated composer.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Composer.lookup,
		},
		{
			Name:		 "arranger",
			Description: "Searches for songs that were arranged by the stated arranger.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Arranger.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.songs = paramsMap
}

func (cfg *Config) initFMVsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "location",
			Description: "Searches for fmvs that are played within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.fmvs = paramsMap
}

func (cfg *Config) initPlayerUnitsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "type",
			Description: "Searches for player units that are of the specified unit type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.UnitType.lookup,
			References:  []string{createListURL(cfg, "unit-type")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.playerUnits = paramsMap
}

func (cfg *Config) initCharactersParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "story_based",
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "underwater",
			Description: "Searches for characters that can fight underwater.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.characters = paramsMap
}

func (cfg *Config) initAeonsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "battles",
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Type: 		     "int",
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		{
			Name:		 "yuna_stats",
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used instead. Every stat instead of luck is available, since an aeon simply copies Yuna's luck stat. Can be used in combination with the 'battles' parameter.",
			Type: 		 "stat",
			ExampleUses: []string{"?yuna_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?yuna_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:		 "optional",
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.aeons = paramsMap
}

func (cfg *Config) initCharacterClassesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "category",
			Description: "Searches for character classes with the specified categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CharacterClassCategory.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.characterClasses = paramsMap
}

func (cfg *Config) initMonstersParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "kimahri_stats",
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Type: 		 "stat",
			ExampleUses: []string{"?kimahri_stats=hp=3000,strength=25,magic=30,agility=40", "?kimahri_stats=hp=15000,agility=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{167, 168},
		},
		{
			Name:		 "aeon_stats",
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Type: 		 "stat",
			ExampleUses: []string{"?aeon_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?aeon_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		{
			Name:		 "altered_state",
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Type:		 "id",
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:		   "omnis_elements",
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Type: 		   "other",
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wilf"},
			ForList:       false,
			ForSingle:     true,
			AllowedIDs:    []int32{211},
			AllowedValues: []string{"f", "l", "w", "i"},
		},
		{
			Name:		 "elemental_resists",
			Description: "Searches for monsters that have the specified elemental affinities.",
			Type: 		 "other",
			Usage:       "?elemental_resists={element|id}={affinity|id},...",
			ExampleUses: []string{"?elemental_resists=fire=weak,water=absorb", "?elemental_resists=1=3,2=4"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements"), createListURL(cfg, "affinities")},
		},
		{
			Name:		 "status_resists",
			Description: "Searches for monsters that resist or are immune to the specified status conditions. You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Type: 		 "id-list",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 	 "resistance",
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Type: 		     "int",
			ForList:         true,
			ForSingle:       false,
			RequiredParams:  []string{"status_resists"},
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
			Name:		 "item",
			Description: "Searches for monsters that have the specified item as loot. Can be specified further with the 'method' parameter.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 	"method",
			Description:    "Specifies the methods of acquisition for the 'item' parameter.",
			Type:		 	"value-list",
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"steal", "drop", "bribe", "other"},
		},
		{
			Name:		 "auto_ability",
			Description: "Searches for monsters that drop the specified auto-ability.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		{
			Name:		 	"is_forced",
			Description:    "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Type:		 	"bool",
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"auto_ability"},
		},
		{
			Name:		 "empty_slots",
			Description:     "Searches for monsters that can drop equipment with the specified amounts of empty slots and no other auto-abilities attached to it.",
			Type: 		 	 "int-list",
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:		 "ronso_rage",
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "ronso-rages")},
		},
		{
			Name:		 "location",
			Description: "Searches for monsters that appear within the specified location.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		{
			Name:		 "sublocation",
			Description: "Searches for monsters that appear within the specified sublocation.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		{
			Name:		 "area",
			Description: "Searches for monsters that appear within the specified area.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		{
			Name:		 "distance",
			Description:     "Searches for monsters with the specified distances. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Type: 		 	 "int-list",
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		{
			Name:		 "availability",
			Description: "Searches for monsters with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Searches for monsters that can be farmed.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "capture",
			Description: "Searches for monsters that can be captured.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "has_overdrive",
			Description: "Searches for monsters that have an overdrive gauge.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "underwater",
			Description: "Searches for monsters that are fought underwater.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "zombie",
			Description: "Searches for monsters that are zombies.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "species",
			Description: "Searches for monsters of the specified species.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterSpecies.lookup,
			References:  []string{createListURL(cfg, "monster-species")},
		},
		{
			Name:		 "creation_area",
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CreationArea.lookup,
		},
		{
			Name:		 "category",
			Description: "Searches for monsters that are of the specified monster-categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterCategory.lookup,
			References:  []string{createListURL(cfg, "monster-category")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.monsters = paramsMap
}

func (cfg *Config) initAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "type",
			Description: "Searches for abilities that are of the specified ability types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AbilityType.lookup,
			References:  []string{createListURL(cfg, "ability-type")},
		},
		{
			Name:		 "rank",
			Description: "Searches for abilities with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "copycat",
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "help_bar",
			Description: "Searches for abilities whose names appear in the help bar.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "monster",
			Description: "Searches for abilities that can be used by the specified monster.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		{
			Name:		 "target_type",
			Description: "Searches for abilities with the specified target types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:		 "user_atk",
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "darkable",
			Description: "Searches for abilities that are affected by 'darkness'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "silenceable",
			Description: "Searches for abilities that are affected by 'silence'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "reflectable",
			Description: "Searches for abilities that are affected by 'reflect'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "attack_type",
			Description: "Searches for abilities with battle interactions of the specified attack types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		{
			Name:		 "damage_type",
			Description: "Searches for abilities that deal the specified types of damage.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		{
			Name:		 "damage_formula",
			Description: "Searches for abilities that use the specified formula to calculate their damage.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		{
			Name:		 "can_crit",
			Description: "Searches for abilities that can land critical hits.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "bdl",
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "element",
			Description: "Searches for abilities that deal elemental damage based on the specified element.",
			Type:		 "name/id-list-nul",
			ExampleVals: []string{"fire", "ice"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		{
			Name:		 "delay",
			Description: "Searches for abilities that deal delay.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "status_inflict",
			Description: "Searches for abilities that can inflict the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "status_remove",
			Description: "Searches for abilities that can remove the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "stat_changes",
			Description: "Searches for abilities that cause stat changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mod_changes",
			Description: "Searches for abilities that cause modifier changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.abilities = paramsMap
}

func (cfg *Config) initPlayerAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "ability_user",
			Description: "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:		 "name/id",
			ExampleVals: []string{"wakka", "valefor"},
			ForList:     false,
			ForSingle:   true,
			References:  []string{createListURL(cfg, "player-units")},
		},
		{
			Name:		 	"bomb_wpn",
			Description:    "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:		    "bool",
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []string{"ability_user"},
		},
		{
			Name:		 "rank",
			Description: "Searches for player abilities with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "copycat",
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "help_bar",
			Description: "Searches for player abilities whose names appear in the help bar.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "category",
			Description: "Searches for player abilities that are of the specified player ability categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.PlayerAbilityCategory.lookup,
			References:  []string{createListURL(cfg, "player-ability-category")},
		},
		{
			Name:		 "outside_battle",
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mp",
			Description: "Searches for player abilities with the specified mp costs.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mp_min",
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Type: 		 "int",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mp_max",
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Type: 		 "int",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "related_stat",
			Description: "Searches for player abilities that are related to the specified stat.",
			Type:		 "name/id",
			ExampleVals: []string{"hp", "strength"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		{
			Name:		 "user",
			Description: "Searches for player abilities that are learned by the specified character class.",
			Type:		 "name/id",
			ExampleVals: []string{"characters", "tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		{
			Name:		 "std_sg",
			Description: "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Type:		 "name/id",
			ExampleVals: []string{"tidus", "wakka"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "characters")},
		},
		{
			Name:		 "exp_sg",
			Description: "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Type:		 "name/id",
			ExampleVals: []string{"tidus", "wakka"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "characters")},
		},
		{
			Name:		 "learn_item",
			Description: "Searches for player abilities an aeon can learn via the specified item.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 "target_type",
			Description: "Searches for player abilities with the specified target types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:		 "user_atk",
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "darkable",
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "silenceable",
			Description: "Searches for player abilities that are affected by 'silence'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "reflectable",
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "attack_type",
			Description: "Searches for player abilities with battle interactions of the specified attack types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		{
			Name:		 "damage_type",
			Description: "Searches for player abilities that deal the specified types of damage.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		{
			Name:		 "damage_formula",
			Description: "Searches for player abilities that use the specified formula to calculate their damage.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		{
			Name:		 "element",
			Description: "Searches for player abilities that deal elemental damage based on the specified element.",
			Type:		 "name/id-list-nul",
			ExampleVals: []string{"fire", "ice"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		{
			Name:		 "delay",
			Description: "Searches for player abilities that deal delay.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "status_inflict",
			Description: "Searches for player abilities that can inflict the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "status_remove",
			Description: "Searches for player abilities that can remove the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "stat_changes",
			Description: "Searches for player abilities that cause stat changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mod_changes",
			Description: "Searches for player abilities that cause modifier changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.playerAbilities = paramsMap
}

func (cfg *Config) initOverdriveAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rank",
			Description: "Searches for overdrive abilities with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "user",
			Description: "Searches for overdrive abilities that are learned by the specified character class.",
			Type:		 "name/id",
			ExampleVals: []string{"characters", "tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		{
			Name:		 "related_stat",
			Description: "Searches for overdrive abilities that are related to the specified stat.",
			Type:		 "name/id",
			ExampleVals: []string{"hp", "strength"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		{
			Name:		 "target_type",
			Description: "Searches for overdrive abilities with the specified target types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:		 "attack_type",
			Description: "Searches for overdrive abilities with battle interactions of the specified attack types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		{
			Name:		 "damage_formula",
			Description: "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		{
			Name:		 "can_crit",
			Description: "Searches for overdrive abilities that can land critical hits.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "element",
			Description: "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Type:		 "name/id-list-nul",
			ExampleVals: []string{"fire", "ice"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		{
			Name:		 "delay",
			Description: "Searches for overdrive abilities that deal delay.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "status_inflict",
			Description: "Searches for overdrive abilities that can inflict the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "status_remove",
			Description: "Searches for overdrive abilities that can remove the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "stat_changes",
			Description: "Searches for overdrive abilities that cause stat changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mod_changes",
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.overdriveAbilities = paramsMap
}

func (cfg *Config) initItemAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "category",
			Description: "Searches for item abilities that are of the specified item categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ItemCategory.lookup,
			References:  []string{createListURL(cfg, "item-category")},
		},
		{
			Name:		 "outside_battle",
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "related_stat",
			Description: "Searches for item abilities that are related to the specified stat.",
			Type:		 "name/id",
			ExampleVals: []string{"hp", "strength"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		{
			Name:		 "target_type",
			Description: "Searches for item abilities with the specified target types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:		 "attack_type",
			Description: "Searches for item abilities with battle interactions of the specified attack types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		{
			Name:		 "damage_formula",
			Description: "Searches for item abilities that use the specified formula to calculate their damage.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		{
			Name:		 "element",
			Description: "Searches for item abilities that deal elemental damage based on the specified element.",
			Type:		 "name/id-list-nul",
			ExampleVals: []string{"fire", "ice"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		{
			Name:		 "delay",
			Description: "Searches for item abilities that deal delay.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "status_inflict",
			Description: "Searches for item abilities that can inflict the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "status_remove",
			Description: "Searches for item abilities that can remove the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "stat_changes",
			Description: "Searches for item abilities that cause stat changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "mod_changes",
			Description: "Searches for item abilities that cause modifier changes.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.itemAbilities = paramsMap
}

func (cfg *Config) initTriggerCommandsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "ability_user",
			Description: "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Type:		 "name/id",
			ExampleVals: []string{"wakka", "valefor"},
			ForList:     false,
			ForSingle:   true,
			References:  []string{createListURL(cfg, "player-units")},
		},
		{
			Name:		 	"bomb_wpn",
			Description:    "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:		    "bool",
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []string{"ability_user"},
		},
		{
			Name:		 "related_stat",
			Description: "Searches for trigger commands that are related to the specified stat.",
			Type:		 "name/id",
			ExampleVals: []string{"hp", "strength"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		{
			Name:		 "user",
			Description: "Searches for trigger commands that are learned by the specified character class.",
			Type:		 "name/id",
			ExampleVals: []string{"characters", "tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.triggerCommands = paramsMap
}

func (cfg *Config) initUnspecifiedAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "ability_user",
			Description: "If an unspecified ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Type:		 "name/id",
			ExampleVals: []string{"wakka", "valefor"},
			ForList:     false,
			ForSingle:   true,
			References:  []string{createListURL(cfg, "player-units")},
		},
		{
			Name:		 	"bomb_wpn",
			Description:    "If an unspecified ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Type:		    "bool",
			ForList:        false,
			ForSingle:      true,
			RequiredParams: []string{"ability_user"},
		},
		{
			Name:		 "rank",
			Description: "Searches for unspecified abilities with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "copycat",
			Description: "Searches for unspecified abilities that can be copied by 'copycat'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "help_bar",
			Description: "Searches for unspecified abilities whose names appear in the help bar.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "user",
			Description: "Searches for unspecified abilities that are learned by the specified character class.",
			Type:		 "name/id",
			ExampleVals: []string{"characters", "tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		{
			Name:		 "user_atk",
			Description: "Searches for unspecified abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.unspecifiedAbilities = paramsMap
}

func (cfg *Config) initEnemyAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rank",
			Description: "Searches for enemy abilities with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "help_bar",
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "monster",
			Description: "Searches for enemy abilities that can be used by the specified monster.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		{
			Name:		 "target_type",
			Description: "Searches for enemy abilities with the specified target types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TargetType.lookup,
		},
		{
			Name:		 "darkable",
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "silenceable",
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "reflectable",
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "attack_type",
			Description: "Searches for enemy abilities with battle interactions of the specified attack types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		{
			Name:		 "damage_type",
			Description: "Searches for enemy abilities that deal the specified types of damage.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		{
			Name:		 "damage_formula",
			Description: "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		{
			Name:		 "can_crit",
			Description: "Searches for enemy abilities that can land critical hits.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "bdl",
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "element",
			Description: "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Type:		 "name/id-list-nul",
			ExampleVals: []string{"fire", "ice"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		{
			Name:		 "delay",
			Description: "Searches for enemy abilities that deal delay.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "status_inflict",
			Description: "Searches for enemy abilities that can inflict the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		{
			Name:		 "status_remove",
			Description: "Searches for enemy abilities that can remove the specified status condition.",
			Type:		 "id-nul",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.enemyAbilities = paramsMap
}

func (cfg *Config) initOverdrivesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rank",
			Description: "Searches for overdrives with the specified ranks.",
			Type: 		 "int-list",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "user",
			Description: "Searches for overdrives that are learned by the specified character class.",
			Type:		 "name/id",
			ExampleVals: []string{"characters", "tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.overdrives = paramsMap
}

func (cfg *Config) initSubmenusParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "topmenu",
			Description: "Searches for submenus that are found within the specified topmenu.",
			Type:		 "name/id",
			ExampleVals: []string{"main", "left"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "topmenus")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.submenus = paramsMap
}

func (cfg *Config) initAllItemsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "type",
			Description: "Searches for items that are of the specified item-types.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ItemType.lookup,
			References:  []string{createListURL(cfg, "item-type")},
		},
		{
			Name:		 "method",
			Description: "Searches for items that can be obtained via all of the given methods.",
			Type:		 "value-list",
			ForList:     true,
			ForSingle:   false,
			AllowedValues: []string{"monster", "treasure", "shop", "quest"},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.allItems = paramsMap
}

func (cfg *Config) initItemsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rel_availability",
			Description: "Only displays an item's related resources with the given availabilities.",
			Type:		 "enum-list",
			ForList:     false,
			ForSingle:   true,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Only displays an item's related monsters and quests that can be farmed.",
			Type:		 "bool",
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:		 "category",
			Description: "Searches for items that are from one of the specified item categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ItemCategory.lookup,
			References:  []string{createListURL(cfg, "item-category")},
		},
		{
			Name:		 "method",
			Description: "Searches for items that can be obtained via all of the given methods.",
			Type:		 "value-list",
			ForList:     true,
			ForSingle:   false,
			AllowedValues: []string{"monster", "treasure", "shop", "quest"},
		},
		{
			Name:		 "has_ability",
			Description: "Searches for items that can be used in battle.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
		{
			Name:		 "related_stat",
			Description: "Searches for items that are related to the specified stat.",
			Type:		 "name/id",
			ExampleVals: []string{"hp", "strength"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.items = paramsMap
}

func (cfg *Config) initKeyItemsParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "availability",
			Description: "Searches for key-items with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "category",
			Description: "Searches for key-items that are of the specified key-item categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.KeyItemCategory.lookup,
			References:  []string{createListURL(cfg, "key-item-category")},
		},
		{
			Name:		 "method",
			Description: "Searches for key-items that can be obtained via all of the given methods.",
			Type:		 "value-list",
			ForList:     true,
			ForSingle:   false,
			AllowedValues: []string{"treasure", "shop"},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.keyItems = paramsMap
}

func (cfg *Config) initSpheresParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rel_availability",
			Description: "Only displays a sphere's related resources with the given availabilities.",
			Type:		 "enum-list",
			ForList:     false,
			ForSingle:   true,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Only displays an sphere's related monsters and quests that can be farmed.",
			Type:		 "bool",
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:		 "color",
			Description: "Searches for spheres with any of the given colors.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.SphereColor.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.spheres = paramsMap
}

func (cfg *Config) initPrimersParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "availability",
			Description: "Searches for primers with the given availabilities.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.primers = paramsMap
}

func (cfg *Config) initMixesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "contains_item",
			Description: "Modifies combinations to only display item combinations that include the specified item.",
			Type:		 "name/id",
			ExampleVals: []string{"grenade", "power_sphere"},
			ForList:     false,
			ForSingle:   true,
			ForbiddenParams: []string{"best"},
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 "best",
			Description: "Modifies combinations to only display the easiest item combinations to accumulate (hand-picked by the dev).",
			Type:		 "bool",
			ForList:     false,
			ForSingle:   true,
			ForbiddenParams: []string{"contains_item"},
		},
		{
			Name:		 "category",
			Description: "Searches for mixes that are of the specified mix categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MixCategory.lookup,
			References:  []string{createListURL(cfg, "mix-category")},
		},
		{
			Name:		 "req_item",
			Description: "Searches for mixes that can be built with the specified item.",
			Type:		 "name/id",
			ExampleVals: []string{"grenade", "power_sphere"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		{
			Name:		 "second_item",
			Description: "Can be used in combination with 'req_item' to get the mix the two specified items will create.",
			Type:		 "name/id",
			ExampleVals: []string{"grenade", "power_sphere"},
			ForList:     true,
			ForSingle:   false,
			RequiredParams: []string{"req_item"},
			References:  []string{createListURL(cfg, "items")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, true)
	cfg.q.mixes = paramsMap
}

func (cfg *Config) initAutoAbilitiesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "rel_availability",
			Description: "Only displays an auto-ability's related monsters with the given availabilities.",
			Type:		 "enum-list",
			ForList:     false,
			ForSingle:   true,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "repeatable",
			Description: "Only displays an auto-ability's related monsters that can be farmed.",
			Type:		 "bool",
			ForList:     false,
			ForSingle:   true,
		},
		{
			Name:		 "category",
			Description: "Searches for auto-abilities that are of the specified auto-ability categories.",
			Type:		 "enum-list",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AutoAbilityCategory.lookup,
			References:  []string{createListURL(cfg, "auto-ability-category")},
		},
		{
			Name:		 "type",
			Description: "Searches for auto-abilities that are of the specified equip type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:		 "monster",
			Description: "Searches for auto-abilities that are dropped by the specified monster.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		{
			Name:		 "monster_items",
			Description: "Searches for auto-abilities that can be crafted with the items dropped by the specified monster.",
			Type:		 "id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.autoAbilities = paramsMap
}

func (cfg *Config) initEquipmentTablesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "auto_abilities",
			Description: "Searches for equipment tables with all of the given auto-abilities.",
			Type:		 "id-list",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		{
			Name:		 "type",
			Description: "Searches for equipment tables that are of the specified equip type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:		 "celestial_weapon",
			Description: "Searches for the equipment tables of the celestial weapons.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.equipmentTables = paramsMap
}

func (cfg *Config) initEquipmentParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name: 		 "table",
			Description: "Selects the equipment table whose data should be displayed for celestial weapons and the brotherhood. The default is set to the fully-upgraded table (1). For the brotherhood, only 1 and 2 are available. For celestial weapons, 1 equals the fully-upgraded table, 2 is the table with just the crest, and 3 is the table with no upgrades.",
			Type: 		 "int",
			ForSingle:   true,
			ForList: 	 false,
			AllowedIDs:  []int32{1, 2, 3, 4, 5, 6, 7, 8},
			DefaultVal:  h.GetIntPtr(1),
		},
		{
			Name:		 "rel_availability",
			Description: "Only displays an equipment's related treasures and shops with the given availabilities.",
			Type:		 "enum-list",
			ForList:     false,
			ForSingle:   true,
			TypeLookup:  cfg.t.AvailabilityType.lookup,
			References:  []string{createListURL(cfg, "availability")},
		},
		{
			Name:		 "auto_abilities",
			Description: "Searches for equipment with all of the given auto-abilities.",
			Type:		 "id-list",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		{
			Name:		 "character",
			Description: "Searches for equipment of the specified character.",
			ExampleVals: []string{"yuna"},
			Type:		 "name/id",
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character")},
		},
		{
			Name:		 "type",
			Description: "Searches for equipment that is of the specified equip type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.EquipType.lookup,
		},
		{
			Name:		 "celestial_weapon",
			Description: "Searches for the celestial weapons.",
			Type:		 "bool",
			ForList:     true,
			ForSingle:   false,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.equipment = paramsMap
}

func (cfg *Config) initOverdriveModesParams(defaultParams []QueryParam) {
	params := []QueryParam{
		{
			Name:		 "type",
			Description: "Searches for overdrive modes that are of the specified overdrive-mode-type.",
			Type:		 "enum",
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.OverdriveModeType.lookup,
		},
	}

	paramsMap := cfg.completeQueryParamsInit(params, defaultParams, false)
	cfg.q.overdriveModes = paramsMap
}
