package main

import (
	"fmt"
	"maps"
	"net/http"
	"strings"
)

type QueryType struct {
    Name          	string		`json:"name"`
    Description   	string		`json:"description"`
    Usage			string		`json:"usage"`
	ExampleUses		[]string	`json:"example_uses"`
    ForList       	bool		`json:"for_list"`
    ForSingle     	bool		`json:"for_single"`
    RequiredWith  	[]string	`json:"required_with,omitempty"`
    AllowedIDs   	[]int32		`json:"allowed_ids,omitempty"`
    AllowedValues 	[]string	`json:"allowed_values,omitempty"`
	AllowedIntRange	[]int		`json:"allowed_int_range,omitempty"`
}


type QueryLookup struct {
	areas			map[string]QueryType
	monsters		map[string]QueryType
	overdriveModes	map[string]QueryType
}


func QueryLookupInit() QueryLookup {
	q := QueryLookup{}
	defaultParams := map[string]QueryType{
		"limit": {
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20.",
			Usage: "?limit{integer}",
			ExampleUses: []string{"?limit=50"},
			ForList: true,
			ForSingle: false,
		},
		"offset": {
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 20.",
			Usage: "?offset{integer}",
			ExampleUses: []string{"?offset=30"},
			ForList: true,
			ForSingle: false,
		},
	}

	q.initAreasParams(defaultParams)
	q.initMonstersParams(defaultParams)
	q.initOverdriveModesParams(defaultParams)

	return q
}


// if a parameter restricted by ID can be used in a list, (like kimahri-stats, or omnis-elements for /areas/monsters), I will simply not enforce the result needing to be in the current page and just leave it be
// for allowedValues and allowedIntRange, since the query needs to be parsed within the handler anyways, I will simply keep these entries as information and not do anything with them at this stage

// called after parsing, when I have the id (won't need to be called for multiple resources)
func verifyQueryParams(r *http.Request, endpoint string, id *int32, lookup map[string]QueryType) error {
	q := r.URL.Query()

	for param := range q {
		queryType, ok := lookup[param]
		if !ok {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Parameter %s does not exist for endpoint %s.", param, endpoint), nil)
		}

		if queryType.RequiredWith != nil {
			for _, reqParam := range queryType.RequiredWith {
				reqParamVal := q.Get(reqParam)
				if reqParamVal == "" {
					return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used in combination with parameter(s): %s.", param, param, strings.Join(queryType.RequiredWith, ", ")), nil)
				}
			}
		}

		if queryType.ForSingle {
			if id == nil {
				return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with single-resource-endpoints.", param, param), nil)
			}

			if queryType.AllowedIDs != nil {
				allowedIDPresent := false
				
				for _, reqID := range queryType.AllowedIDs {
					if *id == reqID {
						allowedIDPresent = true
					}
				}
				if !allowedIDPresent {
					idsString := strings.Trim(strings.Join(strings.Split(fmt.Sprint(queryType.AllowedIDs), " "), ", "), "[]")
					return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid id %d. Parameter %s can only be used with ids %s.", *id, param, idsString), nil)
				}
			}
		}

		if queryType.ForList && id != nil {
			return newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid usage of parameter %s. Parameter %s can only be used with list-endpoints.", param, param), nil)
		}
	}

	return nil
}


func (q *QueryLookup) initAreasParams(defaultParams map[string]QueryType) {
	q.areas = map[string]QueryType{
		"location": {
			Description: "Searches for areas that are located within the specified location.",
			Usage: "?location={location_name/id}",
			ExampleUses: []string{"?location=17", "?location=macalania"},
			ForList: true,
			ForSingle: false,
		},
		"sublocation": {
			Description: "Searches for areas that are located within the specified sublocation.",
			Usage: "?sublocation={sublocation_name/id}",
			ExampleUses: []string{"?sublocation=13", "?sublocation=sanubia-desert"},
			ForList: true,
			ForSingle: false,
		},
		"item": {
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage: "?item={item_name/id}",
			ExampleUses: []string{"?item=45", "?item=mega-potion"},
			ForList: true,
			ForSingle: false,
		},
		"method": {
			Description: "Specifies the method of acquisition for the item parameter.",
			Usage: "?item={item_name/id}&method={value}",
			ExampleUses: []string{"?item=45&method=treasure", "?item=hi-potion&method=shop"},
			ForList: true,
			ForSingle: false,
			RequiredWith: []string{"item"},
			AllowedValues: []string{"monster", "treasure", "shop", "quest"},
		},
		"key-item": {
			Description: "Searches for areas where the specified key item can be acquired.",
			Usage: "?key-item={key_item_name/id}",
			ExampleUses: []string{"?key-item=mars-crest", "key-item=22"},
			ForList: true,
			ForSingle: false,
		},
		"story-based": {
			Description: "Searches for areas that can only be visited during story events.",
			Usage: "?story-based={boolean}",
			ExampleUses: []string{"?story-based=true", "?story-based=false"},
			ForList: true,
			ForSingle: false,
		},
		"save-sphere": {
			Description: "Searches for areas that have a save sphere.",
			Usage: "?save-sphere={boolean}",
			ExampleUses: []string{"?save-sphere=true", "?save-sphere=false"},
			ForList: true,
			ForSingle: false,
		},
		"comp-sphere": {
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Usage: "?comp-sphere={boolean}",
			ExampleUses: []string{"?comp-sphere=true", "?comp-sphere=false"},
			ForList: true,
			ForSingle: false,
		},
		"airship": {
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Usage: "?airship={boolean}",
			ExampleUses: []string{"?airship=true", "?airship=false"},
			ForList: true,
			ForSingle: false,
		},
		"chocobo": {
			Description: "Searches for areas where you can ride a chocobo.",
			Usage: "?chocobo={boolean}",
			ExampleUses: []string{"?chocobo=true", "?chocobo=false"},
			ForList: true,
			ForSingle: false,
		},
		"characters": {
			Description: "Searches for areas where a character permanently joins the party.",
			Usage: "?characters={boolean}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList: true,
			ForSingle: false,
		},
		"aeons": {
			Description: "Searches for areas where a new aeon is acquired.",
			Usage: "?aeons={boolean}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList: true,
			ForSingle: false,
		},
		"monsters": {
			Description: "Searches for areas that have monsters.",
			Usage: "?monsters={boolean}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList: true,
			ForSingle: false,
		},
		"boss-fights": {
			Description: "Searches for areas that have bosses.",
			Usage: "?boss-fights={boolean}",
			ExampleUses: []string{"?boss-fights=true", "?boss-fights=false"},
			ForList: true,
			ForSingle: false,
		},
		"shops": {
			Description: "Searches for areas that have shops.",
			Usage: "?shops={boolean}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList: true,
			ForSingle: false,
		},
		"treasures": {
			Description: "Searches for areas that have treasures.",
			Usage: "?treasures={boolean}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList: true,
			ForSingle: false,
		},
		"sidequests": {
			Description: "Searchces for areas that feature sidequests.",
			Usage: "?sidequests={boolean}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList: true,
			ForSingle: false,
		},
		"fmvs": {
			Description: "Searches for areas that feature fmv sequences.",
			Usage: "?fmvs={boolean}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList: true,
			ForSingle: false,
		},
	}

	maps.Copy(q.areas, defaultParams)

	for key := range q.areas {
		entry := q.areas[key]
		entry.Name = key
		q.areas[key] = entry
	}
}


func (q *QueryLookup) initMonstersParams(defaultParams map[string]QueryType) {
	q.monsters = map[string]QueryType{
		"kimahri-stats": {
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. If unused, their stats are based on Kimahri's base stats.",
			Usage: "?kimahri-stats={stat}-{value},{stat}-{value}",
			ExampleUses: []string{"?kimahri-stats=hp-3000,strength-25,magic-30,agility-40", "?kimahri-stats=hp15000,agility-255"},
			ForList: false,
			ForSingle: true,
			AllowedIDs: []int32{167, 168},
		},
		"altered-state": {
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Usage: "?altered-state={alt_state_id}",
			ExampleUses: []string{"?altered-state=1"},
			ForList: false,
			ForSingle: true,
		},
		"omnis-elements": {
			Description: "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Usage: "?omnis-elements={letter}x4",
			ExampleUses: []string{"?omnis-elements=ifil", "?omnis-elements=llll", "?omnis-elements=wfwf"},
			ForList: false,
			ForSingle: true,
			AllowedIDs: []int32{211},
			AllowedValues: []string{"f", "l", "w", "i"},
		},
		"elemental-affinities": {
			Description: "Searches for monsters that have the specified elemental affinity/affinities.",
			Usage: "?elemental-affinities={element_name/id}-{affinity_name/id},{element_name/id}-{affinity_name/id},...",
			ExampleUses: []string{"?elemental-affinities={fire}-{weak},{water}-{absorb}", "?elemental-affinities={lightning}-{neutral}"},
			ForList: true,
			ForSingle: false,
		},
		"status-resists": {
			Description: "Searches for monsters that resist or are immune to the specified status condition(s). You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Usage: "?status-resists={status_name/id},{status_name/id},...",
			ExampleUses: []string{"status-resists=darkness,silence", "status-resists=poison&resistance=50"},
			ForList: true,
			ForSingle: false,
		},
		"resistance": {
			Description: "Specifies the minimum resistance for the status-resists parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used.",
			Usage: "status-resists={status_name/id},{status_name/id},...&resistance={1-254 or immune}",
			ExampleUses: []string{"status-resists=poison,death&resistance=50", "status-resists=power-break,death&resistance=30", "status-resists=sleep&resistance=immune"},
			ForList: true,
			ForSingle: false,
			RequiredWith: []string{"status-resists"},
			AllowedValues: []string{"immune"},
			AllowedIntRange: []int{1, 254},
		},
		"item": {
			Description: "Searches for monsters that have the specified item as loot. Can be specified further with the 'method' parameter.",
			Usage: "?item={item_name/id}",
			ExampleUses: []string{"?item=45", "?item=mega-potion"},
			ForList: true,
			ForSingle: false,
		},
		"method": {
			Description: "Specifies the method of acquisition for the item parameter.",
			Usage: "?item={item_name/id}&method={value}",
			ExampleUses: []string{"?item=45&method=steal", "?item=hi-potion&method=bribe"},
			ForList: true,
			ForSingle: false,
			RequiredWith: []string{"item"},
			AllowedValues: []string{"steal", "drop", "bribe", "other"},
		},
		"auto-abilities": {
			Description: "Searches for monsters that drop one of the specified auto-abilities.",
			Usage: "?auto-abilities={auto_ability_name/id},{auto_ability_name/id},...",
			ExampleUses: []string{"?auto-abilities=16,28", "?auto-abilities=sos-haste,auto-haste", "?auto-abilities=fireproof"},
			ForList: true,
			ForSingle: false,
		},
		"ronso-rage": {
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Usage: "?ronso-rage={ronso_rage_name/id}",
			ExampleUses: []string{"?ronso-rage=5", "?ronso-rage=nova"},
			ForList: true,
			ForSingle: false,
		},
		"location": {
			Description: "Searches for monsters that appear within the specified location.",
			Usage: "?location={location_name/id}",
			ExampleUses: []string{"?location=17", "?location=macalania"},
			ForList: true,
			ForSingle: false,
		},
		"sublocation": {
			Description: "Searches for monsters that appear within the specified sublocation.",
			Usage: "?sublocation={sublocation_name/id}",
			ExampleUses: []string{"?sublocation=13", "?sublocation=sanubia-desert"},
			ForList: true,
			ForSingle: false,
		},
		"area": {
			Description: "Searches for monsters that appear within the specified area.",
			Usage: "?area={area_id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList: true,
			ForSingle: false,
		},
		"distance": {
			Description: "Searches for monsters with the specified distance. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Usage: "?distance={value}",
			ExampleUses: []string{"?distance=3"},
			ForList: true,
			ForSingle: false,
			AllowedIntRange: []int{1, 4},
		},
		"story-based": {
			Description: "Searches for monsters that only appear during story events.",
			Usage: "?story-based={boolean}",
			ExampleUses: []string{"?story-based=true", "?story-based=false"},
			ForList: true,
			ForSingle: false,
		},
		"repeatable": {
			Description: "Searches for monsters that can be farmed.",
			Usage: "?repeatable={boolean}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList: true,
			ForSingle: false,
		},
		"capture": {
			Description: "Searches for monsters that can be captured.",
			Usage: "?capture={boolean}",
			ExampleUses: []string{"?capture=true", "?capture=false"},
			ForList: true,
			ForSingle: false,
		},
		"has-overdrive": {
			Description: "Searches for monsters that have an overdrive gauge.",
			Usage: "?has-overdrive={boolean}",
			ExampleUses: []string{"?has-overdrive=true", "?has-overdrive=false"},
			ForList: true,
			ForSingle: false,
		},
		"underwater": {
			Description: "Searches for monsters that are fought underwater.",
			Usage: "?underwater={boolean}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList: true,
			ForSingle: false,
		},
		"zombie": {
			Description: "Searches for monsters that are zombies.",
			Usage: "?zombie={boolean}",
			ExampleUses: []string{"?zombie=true", "?zombie=false"},
			ForList: true,
			ForSingle: false,
		},
		"species": {
			Description: "Searches for monsters of the specified species.",
			Usage: "?species={species_name/id}",
			ExampleUses: []string{"?species=wyrm", "?species=5"},
			ForList: true,
			ForSingle: false,
		},
		"creation-area": {
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Usage: "?creation-area={creation_area_name/id}",
			ExampleUses: []string{"?creation-area=thunder-plains", "?creation-area=5"},
			ForList: true,
			ForSingle: false,
		},
		"type": {
			Description: "Searches for monsters that are of the specified ctb-icon-type. 'boss' and 'boss-numbered' are combined and will yield the same results.",
			Usage: "?type={ctb_icon_type_name/id}",
			ExampleUses: []string{"?type=boss", "?type=2"},
			ForList: true,
			ForSingle: false,
		},
	}

	maps.Copy(q.monsters, defaultParams)

	for key := range q.monsters {
		entry := q.monsters[key]
		entry.Name = key
		q.monsters[key] = entry
	}
}


func (q *QueryLookup) initOverdriveModesParams(defaultParams map[string]QueryType) {
	q.overdriveModes = map[string]QueryType{
		"type": {
			Description: "Searches for overdrive-modes that are of the specified overdrive-mode-type.",
			Usage: "?type={overdrive_mode_type_name/id}",
			ExampleUses: []string{"?type=per-action", "?type=2"},
			ForList: true,
			ForSingle: false,
		},
	}

	maps.Copy(q.overdriveModes, defaultParams)

	for key := range q.overdriveModes {
		entry := q.overdriveModes[key]
		entry.Name = key
		q.overdriveModes[key] = entry
	}
}
