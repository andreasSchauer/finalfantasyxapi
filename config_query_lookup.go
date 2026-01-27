package main

import (
	"maps"
)

type QueryType struct {
	ID               int            `json:"-"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Usage            string         `json:"usage"`
	ExampleUses      []string       `json:"example_uses"`
	ForList          bool           `json:"for_list"`
	ForSingle        bool           `json:"for_single"`
	RequiredParams   []string       `json:"required_params,omitempty"`
	References       []string       `json:"references,omitempty"`
	AllowedIDs       []int32        `json:"-"`
	AllowedResources []APIResource  `json:"allowed_resources,omitempty"`
	AllowedValues    []string       `json:"allowed_values,omitempty"`
	AllowedIntRange  []int          `json:"allowed_int_range,omitempty"`
	DefaultVal       *int           `json:"default_value,omitempty"`
	SpecialInputs    []SpecialInput `json:"special_inputs,omitempty"`
}

type SpecialInput struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	defaultParams  map[string]QueryType
	areas          map[string]QueryType
	locations	   map[string]QueryType
	monsters       map[string]QueryType
	overdriveModes map[string]QueryType
	sublocations   map[string]QueryType
}

func (cfg *Config) QueryLookupInit() {
	defaultLimit := 20
	defaultOffset := 0
	cfg.q = &QueryLookup{}

	cfg.q.defaultParams = map[string]QueryType{
		"limit": {
			ID:          -2,
			Name:        "limit",
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Usage:       "?limit{integer or 'max'}",
			ExampleUses: []string{"?limit=50"},
			ForList:     true,
			ForSingle:   false,
			SpecialInputs: []SpecialInput{
				{
					Key: "max",
					Val: 9999,
				},
			},
			DefaultVal: &defaultLimit,
		},
		"offset": {
			ID:          -1,
			Name:        "offset",
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Usage:       "?offset{integer}",
			ExampleUses: []string{"?offset=30"},
			ForList:     true,
			ForSingle:   false,
			DefaultVal:  &defaultOffset,
		},
	}

	cfg.initAreasParams()
	cfg.initMonstersParams()
	cfg.initOverdriveModesParams()
	cfg.initSublocationsParams()
	cfg.initLocationsParams()
}

func (cfg *Config) completeQueryTypeInit(params map[string]QueryType) map[string]QueryType {
	maps.Copy(params, cfg.q.defaultParams)

	for key, entry := range params {
		entry.Name = key
		params[key] = entry
	}

	return params
}

func (cfg *Config) initAreasParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for areas that are located within the specified location.",
			Usage:       "?location={location_id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          2,
			Description: "Searches for areas that are located within the specified sublocation.",
			Usage:       "?sublocation={sublocation_id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"item": {
			ID:          3,
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={item_id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             4,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          5,
			Description: "Searches for areas where the specified key item can be acquired.",
			Usage:       "?key_item={key_item_id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key_items")},
		},
		"story_based": {
			ID:          6,
			Description: "Searches for areas that can only be visited during story events.",
			Usage:       "?story_based={boolean}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"save_sphere": {
			ID:          7,
			Description: "Searches for areas that have a save sphere.",
			Usage:       "?save_sphere={boolean}",
			ExampleUses: []string{"?save_sphere=true", "?save_sphere=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"comp_sphere": {
			ID:          8,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Usage:       "?comp_sphere={boolean}",
			ExampleUses: []string{"?comp_sphere=true", "?comp_sphere=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"airship": {
			ID:          9,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Usage:       "?airship={boolean}",
			ExampleUses: []string{"?airship=true", "?airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"chocobo": {
			ID:          10,
			Description: "Searches for areas where you can ride a chocobo.",
			Usage:       "?chocobo={boolean}",
			ExampleUses: []string{"?chocobo=true", "?chocobo=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"characters": {
			ID:          11,
			Description: "Searches for areas where a character permanently joins the party.",
			Usage:       "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          12,
			Description: "Searches for areas where a new aeon is acquired.",
			Usage:       "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          13,
			Description: "Searches for areas that have monsters.",
			Usage:       "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          14,
			Description: "Searches for areas that have bosses.",
			Usage:       "?boss_fights={boolean}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          15,
			Description: "Searches for areas that have shops.",
			Usage:       "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          16,
			Description: "Searches for areas that have treasures.",
			Usage:       "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          17,
			Description: "Searchces for areas that feature sidequests.",
			Usage:       "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          18,
			Description: "Searches for areas that feature fmv sequences.",
			Usage:       "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.areas = params
}

func (cfg *Config) initMonstersParams() {
	defaultResistance := 1

	params := map[string]QueryType{
		"kimahri_stats": {
			ID:          1,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Usage:       "?kimahri_stats={stat}-{value},{stat}-{value}",
			ExampleUses: []string{"?kimahri_stats=hp-3000,strength-25,magic-30,agility-40", "?kimahri_stats=hp15000,agility-255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{167, 168},
		},
		"aeon_stats": {
			ID:          2,
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Usage:       "?aeon_stats={stat}-{value},{stat}-{value}",
			ExampleUses: []string{"?aeon_stats=hp-3000,strength-75,defense-50,magic-30,agility-20", "?aeon_stats=accuracy-150,magic_defense-255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		"altered_state": {
			ID:          3,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Usage:       "?altered_state={alt_state_id}",
			ExampleUses: []string{"?altered_state=1"},
			ForList:     false,
			ForSingle:   true,
		},
		"omnis_elements": {
			ID:            4,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Usage:         "?omnis_elements={letter}x4",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wfwf"},
			ForList:       false,
			ForSingle:     true,
			AllowedIDs:    []int32{211},
			AllowedValues: []string{"f", "l", "w", "i"},
		},
		"elemental_resists": {
			ID:          5,
			Description: "Searches for monsters that have the specified elemental affinity/affinities.",
			Usage:       "?elemental_resists={element_name/id}-{affinity_name/id},{element_name/id}-{affinity_name/id},...",
			ExampleUses: []string{"?elemental_resists={fire}-{weak},{water}-{absorb}", "?elemental_resists={lightning}-{neutral}"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements"), createListURL(cfg, "affinities")},
		},
		"status_resists": {
			ID:          6,
			Description: "Searches for monsters that resist or are immune to the specified status condition(s). You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Usage:       "?status_resists={status_condition_id},{status_condition_id},...",
			ExampleUses: []string{"status_resists=1,4"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"resistance": {
			ID:              7,
			Description:     "Specifies the minimum resistance for the status_resists parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Usage:           "status_resists={status_condition_id},{status_condition_id},...&resistance={1-254 or 'immune'}",
			ExampleUses:     []string{"status_resists=13&resistance=50", "status_resists=4,17&resistance=30", "status_resists=sleep&resistance=immune"},
			ForList:         true,
			ForSingle:       false,
			RequiredParams:  []string{"status_resists"},
			AllowedIntRange: []int{1, 254},
			SpecialInputs: []SpecialInput{
				{
					Key: "immune",
					Val: 254,
				},
			},
			DefaultVal: &defaultResistance,
		},
		"item": {
			ID:          8,
			Description: "Searches for monsters that have the specified item as loot. Can be specified further with the 'method' parameter.",
			Usage:       "?item={item_id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "item")},
		},
		"method": {
			ID:             9,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=steal"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"steal", "drop", "bribe", "other"},
		},
		"auto_abilities": {
			ID:          10,
			Description: "Searches for monsters that drop one of the specified auto_abilities.",
			Usage:       "?auto_abilities={auto_ability_id},{auto_ability_id},...",
			ExampleUses: []string{"?auto_abilities=16", "?auto_abilities=99,100"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto_abilities")},
		},
		"ronso_rage": {
			ID:          11,
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Usage:       "?ronso_rage={ronso_rage_id}",
			ExampleUses: []string{"?ronso_rage=5"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "ronso_rages")},
		},
		"location": {
			ID:          12,
			Description: "Searches for monsters that appear within the specified location.",
			Usage:       "?location={location_id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          13,
			Description: "Searches for monsters that appear within the specified sublocation.",
			Usage:       "?sublocation={sublocation_id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"area": {
			ID:          14,
			Description: "Searches for monsters that appear within the specified area.",
			Usage:       "?area={area_id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		"distance": {
			ID:              15,
			Description:     "Searches for monsters with the specified distance. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Usage:           "?distance={value}",
			ExampleUses:     []string{"?distance=3"},
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		"story_based": {
			ID:          16,
			Description: "Searches for monsters that only appear during story events.",
			Usage:       "?story_based={boolean}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"repeatable": {
			ID:          17,
			Description: "Searches for monsters that can be farmed.",
			Usage:       "?repeatable={boolean}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"capture": {
			ID:          18,
			Description: "Searches for monsters that can be captured.",
			Usage:       "?capture={boolean}",
			ExampleUses: []string{"?capture=true", "?capture=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"has_overdrive": {
			ID:          19,
			Description: "Searches for monsters that have an overdrive gauge.",
			Usage:       "?has_overdrive={boolean}",
			ExampleUses: []string{"?has_overdrive=true", "?has_overdrive=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"underwater": {
			ID:          20,
			Description: "Searches for monsters that are fought underwater.",
			Usage:       "?underwater={boolean}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"zombie": {
			ID:          21,
			Description: "Searches for monsters that are zombies.",
			Usage:       "?zombie={boolean}",
			ExampleUses: []string{"?zombie=true", "?zombie=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"species": {
			ID:          22,
			Description: "Searches for monsters of the specified species.",
			Usage:       "?species={species_name/id}",
			ExampleUses: []string{"?species=wyrm", "?species=5"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monster-species")},
		},
		"creation_area": {
			ID:          23,
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Usage:       "?creation_area={creation_area_name/id}",
			ExampleUses: []string{"?creation_area=thunder-plains", "?creation_area=5"},
			ForList:     true,
			ForSingle:   false,
		},
		"type": {
			ID:          24,
			Description: "Searches for monsters that are of the specified ctb-icon-type. 'boss' and 'boss-numbered' are combined and will yield the same results.",
			Usage:       "?type={ctb_icon_type_name/id}",
			ExampleUses: []string{"?type=boss", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "ctb-icon-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.monsters = params
}

func (cfg *Config) initOverdriveModesParams() {
	params := map[string]QueryType{
		"type": {
			ID:          1,
			Description: "Searches for overdrive-modes that are of the specified overdrive-mode-type.",
			Usage:       "?type={overdrive_mode_type_name/id}",
			ExampleUses: []string{"?type=per-action", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "overdrive-mode-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.overdriveModes = params
}

func (cfg *Config) initSublocationsParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for sublocations that are located within the specified location.",
			Usage:       "?location={location_id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"item": {
			ID:          2,
			Description: "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={item_id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             3,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          4,
			Description: "Searches for sublocations where the specified key item can be acquired.",
			Usage:       "?key_item={key_item_id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key_items")},
		},
		"characters": {
			ID:          5,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Usage:       "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          6,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Usage:       "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          7,
			Description: "Searches for sublocations that have monsters.",
			Usage:       "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          8,
			Description: "Searches for sublocations that have bosses.",
			Usage:       "?boss_fights={boolean}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          9,
			Description: "Searches for sublocations that have shops.",
			Usage:       "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          10,
			Description: "Searches for sublocations that have treasures.",
			Usage:       "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          11,
			Description: "Searchces for sublocations that feature sidequests.",
			Usage:       "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          12,
			Description: "Searches for sublocations that feature fmv sequences.",
			Usage:       "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.sublocations = params
}

func (cfg *Config) initLocationsParams() {
	params := map[string]QueryType{
		"item": {
			ID:          1,
			Description: "Searches for locations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={item_id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             2,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          3,
			Description: "Searches for locations where the specified key item can be acquired.",
			Usage:       "?key_item={key_item_id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key_items")},
		},
		"characters": {
			ID:          4,
			Description: "Searches for locations where a character permanently joins the party.",
			Usage:       "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          5,
			Description: "Searches for locations where a new aeon is acquired.",
			Usage:       "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          6,
			Description: "Searches for locations that have monsters.",
			Usage:       "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          7,
			Description: "Searches for locations that have bosses.",
			Usage:       "?boss_fights={boolean}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          8,
			Description: "Searches for locations that have shops.",
			Usage:       "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          9,
			Description: "Searches for locations that have treasures.",
			Usage:       "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          10,
			Description: "Searchces for locations that feature sidequests.",
			Usage:       "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          11,
			Description: "Searches for locations that feature fmv sequences.",
			Usage:       "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.locations = params
}