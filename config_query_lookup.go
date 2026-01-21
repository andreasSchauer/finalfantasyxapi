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
	ForSections      []string       `json:"for_sections"`
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
	monsters       map[string]QueryType
	overdriveModes map[string]QueryType
}

func (cfg *Config) QueryLookupInit() {
	defaultLimit := 20
	defaultOffset := 0
	cfg.q = &QueryLookup{}

	cfg.q.defaultParams = map[string]QueryType{
		"limit": {
			ID:          -3,
			Name: 		 "limit",
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Usage:       "?limit{integer or 'max'}",
			ExampleUses: []string{"?limit=50"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{},
			SpecialInputs: []SpecialInput{
				{
					Key: "max",
					Val: 9999,
				},
			},
			DefaultVal: &defaultLimit,
		},
		"offset": {
			ID:          -2,
			Name: 		 "offset",
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Usage:       "?offset{integer}",
			ExampleUses: []string{"?offset=30"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{},
			DefaultVal:  &defaultOffset,
		},
		"section": {
			ID:          -1,
			Name: 		 "section",
			Description: "Filters query parameters by the section they can be used in within their endpoint. 'self' can be used to display only parameters specific to their own endpoint.",
			Usage:       "/parameters?section={section_name or 'self'}",
			ExampleUses: []string{"/parameters?section=monsters", "/parameters?section=self"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{},
		},
	}

	cfg.initAreasParams()
	cfg.initMonstersParams()
	cfg.initOverdriveModesParams()
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
			ForSections: []string{"areas"},
			References:  []string{cfg.createListURL("locations")},
		},
		"sublocation": {
			ID:          2,
			Description: "Searches for areas that are located within the specified sublocation.",
			Usage:       "?sublocation={sublocation_id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
			References:  []string{cfg.createListURL("sublocations")},
		},
		"item": {
			ID:          3,
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={item_id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
			References:  []string{cfg.createListURL("items")},
		},
		"method": {
			ID:             4,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			ForSections:    []string{"areas"},
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key-item": {
			ID:          5,
			Description: "Searches for areas where the specified key item can be acquired.",
			Usage:       "?key-item={key_item_id}",
			ExampleUses: []string{"?key-item=22"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
			References:  []string{cfg.createListURL("key-items")},
		},
		"story_based": {
			ID:          6,
			Description: "Searches for areas that can only be visited during story events.",
			Usage:       "?story_based={boolean}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"save_sphere": {
			ID:          7,
			Description: "Searches for areas that have a save sphere.",
			Usage:       "?save_sphere={boolean}",
			ExampleUses: []string{"?save_sphere=true", "?save_sphere=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"comp_sphere": {
			ID:          8,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Usage:       "?comp_sphere={boolean}",
			ExampleUses: []string{"?comp_sphere=true", "?comp_sphere=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"airship": {
			ID:          9,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Usage:       "?airship={boolean}",
			ExampleUses: []string{"?airship=true", "?airship=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"chocobo": {
			ID:          10,
			Description: "Searches for areas where you can ride a chocobo.",
			Usage:       "?chocobo={boolean}",
			ExampleUses: []string{"?chocobo=true", "?chocobo=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"characters": {
			ID:          11,
			Description: "Searches for areas where a character permanently joins the party.",
			Usage:       "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"aeons": {
			ID:          12,
			Description: "Searches for areas where a new aeon is acquired.",
			Usage:       "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"monsters": {
			ID:          13,
			Description: "Searches for areas that have monsters.",
			Usage:       "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"boss_fights": {
			ID:          14,
			Description: "Searches for areas that have bosses.",
			Usage:       "?boss_fights={boolean}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"shops": {
			ID:          15,
			Description: "Searches for areas that have shops.",
			Usage:       "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"treasures": {
			ID:          16,
			Description: "Searches for areas that have treasures.",
			Usage:       "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"sidequests": {
			ID:          17,
			Description: "Searchces for areas that feature sidequests.",
			Usage:       "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
		},
		"fmvs": {
			ID:          18,
			Description: "Searches for areas that feature fmv sequences.",
			Usage:       "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"areas"},
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
			ForSections: []string{"monsters"},
			AllowedIDs:  []int32{167, 168},
		},
		"aeon_stats": {
			ID:          2,
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Usage:       "?aeon_stats={stat}-{value},{stat}-{value}",
			ExampleUses: []string{"?aeon_stats=hp-3000,strength-75,defense-50,magic-30,agility-20", "?aeon_stats=accuracy-150,magic_defense-255"},
			ForList:     false,
			ForSingle:   true,
			ForSections: []string{"monsters"},
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		"altered_state": {
			ID:          3,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Usage:       "?altered_state={alt_state_id}",
			ExampleUses: []string{"?altered_state=1"},
			ForList:     false,
			ForSingle:   true,
			ForSections: []string{"monsters"},
		},
		"omnis_elements": {
			ID:            4,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Usage:         "?omnis_elements={letter}x4",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wfwf"},
			ForList:       false,
			ForSingle:     true,
			ForSections:   []string{"monsters"},
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
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("elements"), cfg.createListURL("affinities")},
		},
		"status_resists": {
			ID:          6,
			Description: "Searches for monsters that resist or are immune to the specified status condition(s). You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Usage:       "?status_resists={status_condition_id},{status_condition_id},...",
			ExampleUses: []string{"status_resists=1,4"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("status-conditions")},
		},
		"resistance": {
			ID:              7,
			Description:     "Specifies the minimum resistance for the status_resists parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Usage:           "status_resists={status_condition_id},{status_condition_id},...&resistance={1-254 or 'immune'}",
			ExampleUses:     []string{"status_resists=13&resistance=50", "status_resists=4,17&resistance=30", "status_resists=sleep&resistance=immune"},
			ForList:         true,
			ForSingle:       false,
			ForSections:     []string{"monsters"},
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
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("item")},
		},
		"method": {
			ID:             9,
			Description:    "Specifies the method of acquisition for the item parameter.",
			Usage:          "?item={item_id}&method={value}",
			ExampleUses:    []string{"?item=45&method=steal"},
			ForList:        true,
			ForSingle:      false,
			ForSections:    []string{"monsters"},
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
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("auto_abilities")},
		},
		"ronso_rage": {
			ID:          11,
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Usage:       "?ronso_rage={ronso_rage_id}",
			ExampleUses: []string{"?ronso_rage=5"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("ronso_rages")},
		},
		"location": {
			ID:          12,
			Description: "Searches for monsters that appear within the specified location.",
			Usage:       "?location={location_id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("locations")},
		},
		"sublocation": {
			ID:          13,
			Description: "Searches for monsters that appear within the specified sublocation.",
			Usage:       "?sublocation={sublocation_id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("sublocations")},
		},
		"area": {
			ID:          14,
			Description: "Searches for monsters that appear within the specified area.",
			Usage:       "?area={area_id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("areas")},
		},
		"distance": {
			ID:              15,
			Description:     "Searches for monsters with the specified distance. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Usage:           "?distance={value}",
			ExampleUses:     []string{"?distance=3"},
			ForList:         true,
			ForSingle:       false,
			ForSections:     []string{"monsters"},
			AllowedIntRange: []int{1, 4},
		},
		"story_based": {
			ID:          16,
			Description: "Searches for monsters that only appear during story events.",
			Usage:       "?story_based={boolean}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"repeatable": {
			ID:          17,
			Description: "Searches for monsters that can be farmed.",
			Usage:       "?repeatable={boolean}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"capture": {
			ID:          18,
			Description: "Searches for monsters that can be captured.",
			Usage:       "?capture={boolean}",
			ExampleUses: []string{"?capture=true", "?capture=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"has_overdrive": {
			ID:          19,
			Description: "Searches for monsters that have an overdrive gauge.",
			Usage:       "?has_overdrive={boolean}",
			ExampleUses: []string{"?has_overdrive=true", "?has_overdrive=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"underwater": {
			ID:          20,
			Description: "Searches for monsters that are fought underwater.",
			Usage:       "?underwater={boolean}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"zombie": {
			ID:          21,
			Description: "Searches for monsters that are zombies.",
			Usage:       "?zombie={boolean}",
			ExampleUses: []string{"?zombie=true", "?zombie=false"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"species": {
			ID:          22,
			Description: "Searches for monsters of the specified species.",
			Usage:       "?species={species_name/id}",
			ExampleUses: []string{"?species=wyrm", "?species=5"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("monster-species")},
		},
		"creation_area": {
			ID:          23,
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Usage:       "?creation_area={creation_area_name/id}",
			ExampleUses: []string{"?creation_area=thunder-plains", "?creation_area=5"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
		},
		"type": {
			ID:          24,
			Description: "Searches for monsters that are of the specified ctb-icon-type. 'boss' and 'boss-numbered' are combined and will yield the same results.",
			Usage:       "?type={ctb_icon_type_name/id}",
			ExampleUses: []string{"?type=boss", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			ForSections: []string{"monsters"},
			References:  []string{cfg.createListURL("ctb-icon-type")},
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
			ForSections: []string{"overdrive-modes"},
			References:  []string{cfg.createListURL("overdrive-mode-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.overdriveModes = params
}
