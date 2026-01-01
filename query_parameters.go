package main

import (
	"maps"
)

// might add some functionality for sections/subsections
// just a slice of sections/subsections and a (global?) parameter that filters them out when called
// global parameters will have them empty, meaning they are available from everywhere

type QueryType struct {
	ID					int				`json:"-"`
    Name          		string			`json:"name"`
    Description   		string			`json:"description"`
    Usage				string			`json:"usage"`
	ExampleUses			[]string		`json:"example_uses"`
    ForList       		bool			`json:"for_list"`
    ForSingle     		bool			`json:"for_single"`
	ForSections			[]string		`json:"for_sections"`
    RequiredWith  		[]string		`json:"required_with,omitempty"`
	References			[]string		`json:"references,omitempty"`
    AllowedIDs   		[]int32			`json:"-"`
	AllowedResources	[]IsAPIResource `json:"allowed_resources,omitempty"`
    AllowedValues 		[]string		`json:"allowed_values,omitempty"`
	AllowedIntRange		[]int			`json:"allowed_int_range,omitempty"`
}


type QueryLookup struct {
	areas			map[string]QueryType
	monsters		map[string]QueryType
	overdriveModes	map[string]QueryType
}


func (cfg *Config) QueryLookupInit() {
	cfg.q = &QueryLookup{}
	defaultParams := map[string]QueryType{
		"limit": {
			ID: -3,
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20.",
			Usage: "?limit{integer}",
			ExampleUses: []string{"?limit=50"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{},
		},
		"offset": {
			ID: -2,
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Usage: "?offset{integer}",
			ExampleUses: []string{"?offset=30"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{},
		},
		"section": {
			ID: -1,
			Description: "Filters query parameters by the section they can be used in within their endpoint.",
			Usage: "?section={name}",
			ExampleUses: []string{"?section=monsters"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{},
		},
	}
	cfg.defaultParams = defaultParams

	cfg.initAreasParams()
	cfg.initMonstersParams()
	cfg.initOverdriveModesParams()
}

func (cfg *Config) completeQueryTypeInit(params map[string]QueryType) map[string]QueryType {
	maps.Copy(params, cfg.defaultParams)

	for key, entry := range params {
		entry.Name = key
		params[key] = entry
	}

	return params
}

func (cfg *Config) initAreasParams() {
	params := map[string]QueryType{
		"location": {
			ID: 1,
			Description: "Searches for areas that are located within the specified location.",
			Usage: "?location={location_name/id}",
			ExampleUses: []string{"?location=17", "?location=macalania"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
			References: []string{cfg.createListURL("locations")},
		},
		"sublocation": {
			ID: 2,
			Description: "Searches for areas that are located within the specified sublocation.",
			Usage: "?sublocation={sublocation_name/id}",
			ExampleUses: []string{"?sublocation=13", "?sublocation=sanubia-desert"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
			References: []string{cfg.createListURL("sublocations")},
		},
		"item": {
			ID: 3,
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage: "?item={item_name/id}",
			ExampleUses: []string{"?item=45", "?item=mega-potion"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
			References: []string{cfg.createListURL("items")},
		},
		"method": {
			ID: 4,
			Description: "Specifies the method of acquisition for the item parameter.",
			Usage: "?item={item_name/id}&method={value}",
			ExampleUses: []string{"?item=45&method=treasure", "?item=hi-potion&method=shop"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
			RequiredWith: []string{"item"},
			AllowedValues: []string{"monster", "treasure", "shop", "quest"},
		},
		"key-item": {
			ID: 5,
			Description: "Searches for areas where the specified key item can be acquired.",
			Usage: "?key-item={key_item_name/id}",
			ExampleUses: []string{"?key-item=mars-crest", "?key-item=22"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
			References: []string{cfg.createListURL("key-items")},
		},
		"story-based": {
			ID: 6,
			Description: "Searches for areas that can only be visited during story events.",
			Usage: "?story-based={boolean}",
			ExampleUses: []string{"?story-based=true", "?story-based=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"save-sphere": {
			ID: 7,
			Description: "Searches for areas that have a save sphere.",
			Usage: "?save-sphere={boolean}",
			ExampleUses: []string{"?save-sphere=true", "?save-sphere=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"comp-sphere": {
			ID: 8,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Usage: "?comp-sphere={boolean}",
			ExampleUses: []string{"?comp-sphere=true", "?comp-sphere=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"airship": {
			ID: 9,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Usage: "?airship={boolean}",
			ExampleUses: []string{"?airship=true", "?airship=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"chocobo": {
			ID: 10,
			Description: "Searches for areas where you can ride a chocobo.",
			Usage: "?chocobo={boolean}",
			ExampleUses: []string{"?chocobo=true", "?chocobo=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"characters": {
			ID: 11,
			Description: "Searches for areas where a character permanently joins the party.",
			Usage: "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"aeons": {
			ID: 12,
			Description: "Searches for areas where a new aeon is acquired.",
			Usage: "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"monsters": {
			ID: 13,
			Description: "Searches for areas that have monsters.",
			Usage: "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"boss-fights": {
			ID: 14,
			Description: "Searches for areas that have bosses.",
			Usage: "?boss-fights={boolean}",
			ExampleUses: []string{"?boss-fights=true", "?boss-fights=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"shops": {
			ID: 15,
			Description: "Searches for areas that have shops.",
			Usage: "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"treasures": {
			ID: 16,
			Description: "Searches for areas that have treasures.",
			Usage: "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"sidequests": {
			ID: 17,
			Description: "Searchces for areas that feature sidequests.",
			Usage: "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
		"fmvs": {
			ID: 18,
			Description: "Searches for areas that feature fmv sequences.",
			Usage: "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"areas"},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.areas = params
}


func (cfg *Config) initMonstersParams() {
	params := map[string]QueryType{
		"kimahri-stats": {
			ID: 1,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. If unused, their stats are based on Kimahri's base stats.",
			Usage: "?kimahri-stats={stat}-{value},{stat}-{value}",
			ExampleUses: []string{"?kimahri-stats=hp-3000,strength-25,magic-30,agility-40", "?kimahri-stats=hp15000,agility-255"},
			ForList: false,
			ForSingle: true,
			ForSections: []string{"monsters"},
			AllowedIDs: []int32{167, 168},
			AllowedResources: []IsAPIResource{
				cfg.newNamedAPIResourceSimple("monsters", 167, "biran ronso"),
				cfg.newNamedAPIResourceSimple("monsters", 168, "yenke ronso"),
			},
		},
		"altered-state": {
			ID: 2,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Usage: "?altered-state={alt_state_id}",
			ExampleUses: []string{"?altered-state=1"},
			ForList: false,
			ForSingle: true,
			ForSections: []string{"monsters"},
		},
		"omnis-elements": {
			ID: 3,
			Description: "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Usage: "?omnis-elements={letter}x4",
			ExampleUses: []string{"?omnis-elements=ifil", "?omnis-elements=llll", "?omnis-elements=wfwf"},
			ForList: false,
			ForSingle: true,
			ForSections: []string{"monsters"},
			AllowedIDs: []int32{211},
			AllowedResources: []IsAPIResource{
				cfg.newNamedAPIResourceSimple("monsters", 211, "seymour omnis"),
			},
			AllowedValues: []string{"f", "l", "w", "i"},
		},
		"elemental-affinities": {
			ID: 4,
			Description: "Searches for monsters that have the specified elemental affinity/affinities.",
			Usage: "?elemental-affinities={element_name/id}-{affinity_name/id},{element_name/id}-{affinity_name/id},...",
			ExampleUses: []string{"?elemental-affinities={fire}-{weak},{water}-{absorb}", "?elemental-affinities={lightning}-{neutral}"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("elements"), cfg.createListURL("affinities")},
		},
		"status-resists": {
			ID: 5,
			Description: "Searches for monsters that resist or are immune to the specified status condition(s). You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Usage: "?status-resists={status_condition_name/id},{status_condition_name/id},...",
			ExampleUses: []string{"status-resists=darkness,silence", "status-resists=poison&resistance=50"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("status-conditions")},
		},
		"resistance": {
			ID: 6,
			Description: "Specifies the minimum resistance for the status-resists parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used.",
			Usage: "status-resists={status_condition_name/id},{status_condition_name/id},...&resistance={1-254 or immune}",
			ExampleUses: []string{"status-resists=poison,death&resistance=50", "status-resists=power-break,death&resistance=30", "status-resists=sleep&resistance=immune"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			RequiredWith: []string{"status-resists"},
			AllowedValues: []string{"immune"},
			AllowedIntRange: []int{1, 254},
		},
		"item": {
			ID: 7,
			Description: "Searches for monsters that have the specified item as loot. Can be specified further with the 'method' parameter.",
			Usage: "?item={item_name/id}",
			ExampleUses: []string{"?item=45", "?item=mega-potion"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("item")},
		},
		"method": {
			ID: 8,
			Description: "Specifies the method of acquisition for the item parameter.",
			Usage: "?item={item_name/id}&method={value}",
			ExampleUses: []string{"?item=45&method=steal", "?item=hi-potion&method=bribe"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			RequiredWith: []string{"item"},
			AllowedValues: []string{"steal", "drop", "bribe", "other"},
		},
		"auto-abilities": {
			ID: 9,
			Description: "Searches for monsters that drop one of the specified auto-abilities.",
			Usage: "?auto-abilities={auto_ability_name/id},{auto_ability_name/id},...",
			ExampleUses: []string{"?auto-abilities=16,28", "?auto-abilities=sos-haste,auto-haste", "?auto-abilities=fireproof"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("auto-abilities")},
		},
		"ronso-rage": {
			ID: 10,
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Usage: "?ronso-rage={ronso_rage_name/id}",
			ExampleUses: []string{"?ronso-rage=5", "?ronso-rage=nova"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("ronso-rages")},
		},
		"location": {
			ID: 11,
			Description: "Searches for monsters that appear within the specified location.",
			Usage: "?location={location_name/id}",
			ExampleUses: []string{"?location=17", "?location=macalania"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("locations")},
		},
		"sublocation": {
			ID: 12,
			Description: "Searches for monsters that appear within the specified sublocation.",
			Usage: "?sublocation={sublocation_name/id}",
			ExampleUses: []string{"?sublocation=13", "?sublocation=sanubia-desert"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("sublocations")},
		},
		"area": {
			ID: 13,
			Description: "Searches for monsters that appear within the specified area.",
			Usage: "?area={area_id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("areas")},
		},
		"distance": {
			ID: 14,
			Description: "Searches for monsters with the specified distance. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Usage: "?distance={value}",
			ExampleUses: []string{"?distance=3"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			AllowedIntRange: []int{1, 4},
		},
		"story-based": {
			ID: 15,
			Description: "Searches for monsters that only appear during story events.",
			Usage: "?story-based={boolean}",
			ExampleUses: []string{"?story-based=true", "?story-based=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"repeatable": {
			ID: 16,
			Description: "Searches for monsters that can be farmed.",
			Usage: "?repeatable={boolean}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"capture": {
			ID: 17,
			Description: "Searches for monsters that can be captured.",
			Usage: "?capture={boolean}",
			ExampleUses: []string{"?capture=true", "?capture=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"has-overdrive": {
			ID: 18,
			Description: "Searches for monsters that have an overdrive gauge.",
			Usage: "?has-overdrive={boolean}",
			ExampleUses: []string{"?has-overdrive=true", "?has-overdrive=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"underwater": {
			ID: 19,
			Description: "Searches for monsters that are fought underwater.",
			Usage: "?underwater={boolean}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"zombie": {
			ID: 20,
			Description: "Searches for monsters that are zombies.",
			Usage: "?zombie={boolean}",
			ExampleUses: []string{"?zombie=true", "?zombie=false"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"species": {
			ID: 21,
			Description: "Searches for monsters of the specified species.",
			Usage: "?species={species_name/id}",
			ExampleUses: []string{"?species=wyrm", "?species=5"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("monster-species")},
		},
		"creation-area": {
			ID: 22,
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Usage: "?creation-area={creation_area_name/id}",
			ExampleUses: []string{"?creation-area=thunder-plains", "?creation-area=5"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
		},
		"type": {
			ID: 23,
			Description: "Searches for monsters that are of the specified ctb-icon-type. 'boss' and 'boss-numbered' are combined and will yield the same results.",
			Usage: "?type={ctb_icon_type_name/id}",
			ExampleUses: []string{"?type=boss", "?type=2"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"monsters"},
			References: []string{cfg.createListURL("ctb-icon-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.monsters = params
}


func (cfg *Config) initOverdriveModesParams() {
	params := map[string]QueryType{
		"type": {
			ID: 1,
			Description: "Searches for overdrive-modes that are of the specified overdrive-mode-type.",
			Usage: "?type={overdrive_mode_type_name/id}",
			ExampleUses: []string{"?type=per-action", "?type=2"},
			ForList: true,
			ForSingle: false,
			ForSections: []string{"overdrive-modes"},
			References: []string{cfg.createListURL("overdrive-mode-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params)
	cfg.q.overdriveModes = params
}
