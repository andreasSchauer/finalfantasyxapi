package api

import (
	"maps"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QueryType struct {
	ID               int                         `json:"-"`
	Name             string                      `json:"name"`
	Description      string                      `json:"description"`
	Usage            string                      `json:"usage"`
	ExampleUses      []string                    `json:"example_uses"`
	DefaultOnly      bool                        `json:"only_use_alone"`
	ForSingle        bool                        `json:"for_single"`
	ForList          bool                        `json:"for_list"`
	ForSegment       *string                     `json:"for_segment"`
	IsRequired       bool                        `json:"is_required"`
	TypeLookup       map[string]TypedAPIResource `json:"-"`
	RequiredParams   []string                    `json:"required_params,omitempty"`
	References       []string                    `json:"references,omitempty"`
	AllowedIDs       []int32                     `json:"-"`
	AllowedResources []string                    `json:"allowed_resources,omitempty"`
	AllowedValues    []string                    `json:"allowed_values,omitempty"`
	AllowedIntRange  []int                       `json:"allowed_int_range,omitempty"`
	AllowedResTypes  []string                    `json:"allowed_res_types"`
	DefaultVal       *int                        `json:"default_value,omitempty"`
	SpecialInputs    []SpecialInput              `json:"special_inputs,omitempty"`
}

type SpecialInput struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

// QueryLookup holds all the Query Parameters for the application
type QueryLookup struct {
	defaultParams        map[string]QueryType
	aeons                map[string]QueryType
	arenaCreations       map[string]QueryType
	areas                map[string]QueryType
	blitzballPrizes      map[string]QueryType
	characters           map[string]QueryType
	characterClasses     map[string]QueryType
	fmvs                 map[string]QueryType
	locations            map[string]QueryType
	monsters             map[string]QueryType
	monsterFormations    map[string]QueryType
	overdriveModes       map[string]QueryType
	abilities            map[string]QueryType
	enemyAbilities       map[string]QueryType
	itemAbilities        map[string]QueryType
	unspecifiedAbilities map[string]QueryType
	overdriveAbilities   map[string]QueryType
	playerAbilities      map[string]QueryType
	triggerCommands      map[string]QueryType
	overdriveCommands	 map[string]QueryType
	overdrives			 map[string]QueryType
	ronsoRages			 map[string]QueryType
	aeonCommands		 map[string]QueryType
	submenus			 map[string]QueryType
	topmenus			 map[string]QueryType
	shops                map[string]QueryType
	songs                map[string]QueryType
	sidequests           map[string]QueryType
	subquests            map[string]QueryType
	sublocations         map[string]QueryType
	treasures            map[string]QueryType
}

func (cfg *Config) QueryLookupInit() {
	cfg.q = &QueryLookup{}

	cfg.q.defaultParams = map[string]QueryType{
		"limit": {
			ID:          1001,
			Name:        "limit",
			Description: "Sets the amount of displayed entries in a list response. If not set manually, the default is 20. The value 'max' can also be used to forgo pagination of lists entirely.",
			Usage:       "?limit={int|'max'}",
			ExampleUses: []string{"?limit=50", "?limit=max"},
			ForList:     true,
			ForSingle:   false,
			SpecialInputs: []SpecialInput{
				{
					Key: "max",
					Val: 9999,
				},
			},
			DefaultVal: h.GetIntPtr(20),
		},
		"offset": {
			ID:          1002,
			Name:        "offset",
			Description: "Sets the offset from where to start the displayed entries in a list response. If not set manually, the default is 0.",
			Usage:       "?offset={int}",
			ExampleUses: []string{"?offset=30"},
			ForList:     true,
			ForSingle:   false,
			DefaultVal:  h.GetIntPtr(0),
		},
	}

	cfg.initAeonsParams()
	cfg.initAreasParams()
	cfg.initArenaCreationsParams()
	cfg.initBlitzballPrizesParams()
	cfg.initCharacterClassesParams()
	cfg.initCharactersParams()
	cfg.initFMVsParams()
	cfg.initMonstersParams()
	cfg.initMonsterFormationsParams()
	cfg.initOverdriveModesParams()
	cfg.initAbilitiesParams()
	cfg.initEnemyAbilitiesParams()
	cfg.initItemAbilitiesParams()
	cfg.initUnspecifiedAbilitiesParams()
	cfg.initOverdriveAbilitiesParams()
	cfg.initPlayerAbilitiesParams()
	cfg.initTriggerCommandsParams()
	cfg.initOverdrivesParams()
	cfg.initSubmenusParams()
	cfg.initSublocationsParams()
	cfg.initLocationsParams()
	cfg.initSidequestsParams()
	cfg.initSubquestsParams()
	cfg.initShopsParams()
	cfg.initSongsParams()
	cfg.initTreasuresParams()

	cfg.q.ronsoRages = cfg.assignDefaultParams()
	cfg.q.aeonCommands = cfg.assignDefaultParams()
	cfg.q.topmenus = cfg.assignDefaultParams()
	cfg.q.overdriveCommands = cfg.assignDefaultParams()
}

func (cfg *Config) assignDefaultParams() map[string]QueryType {
	return cfg.completeQueryTypeInit(createEmptyQueryMap(), false)
}

func createEmptyQueryMap() map[string]QueryType {
	return make(map[string]QueryType)
}

func (cfg *Config) completeQueryTypeInit(params map[string]QueryType, hasSimpleView bool) map[string]QueryType {
	maps.Copy(params, cfg.q.defaultParams)

	for key, entry := range params {
		entry.Name = key
		params[key] = entry
	}

	if hasSimpleView {
		params["ids"] = QueryType{
			ID:          1003,
			Name:        "ids",
			Description: "Used to input the ids of resources to be batch-fetched for simple display. The original order will be preserved, but duplicates will be removed.",
			Usage:       "?ids={id},...",
			ExampleUses: []string{"?ids=1,3,4"},
			DefaultOnly: true,
			ForList:     false,
			ForSingle:   false,
			ForSegment:  h.GetStrPtr("simple"),
		}
	}

	return params
}

func (cfg *Config) initAeonsParams() {
	params := map[string]QueryType{
		"battles": {
			ID:              1,
			Description:     "Specifies the amount of battles the player has taken part in and takes them into account when calculating the aeon's stats. An aeon's stats increase for the first time after 60 battles and then every 30 additional battles with the final increase being at 600. Can be used in combination with the 'yuna_stats' parameter.",
			Usage:           "?battles={int}",
			ExampleUses:     []string{"?battles=123"},
			ForList:         false,
			ForSingle:       true,
			AllowedIntRange: []int{0, 600},
			DefaultVal:      h.GetIntPtr(0),
		},
		"yuna_stats": {
			ID:          2,
			Description: "Calculate an aeon's stats based on Yuna's stats. If a stat is not given, Yuna's respective default stat is used instead. Every stat instead of luck is available, since an aeon simply copies Yuna's luck stat. Can be used in combination with the 'battles' parameter.",
			Usage:       "?yuna_stats={stat}={int},...",
			ExampleUses: []string{"?yuna_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?yuna_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
		},
		"optional": {
			ID:          3,
			Description: "Searches for aeons that are not mandatory to complete the main story.",
			Usage:       "?optional={bool}",
			ExampleUses: []string{"?optional=true", "?optional=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.aeons = params
}

func (cfg *Config) initAreasParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for areas that are located within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          2,
			Description: "Searches for areas that are located within the specified sublocation.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"item": {
			ID:          3,
			Description: "Searches for areas where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             4,
			Description:    "Specifies the method of acquisition for the 'item' parameter.",
			Usage:          "?item={id}&method={method_name}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          5,
			Description: "Searches for areas where the specified key item can be acquired.",
			Usage:       "?key_item={id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		"post_airship": {
			ID:          6,
			Description: "Searches for areas that can only be accessed after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"story_based": {
			ID:          7,
			Description: "Searches for areas that can only be accessed during certain sections of the story.",
			Usage:       "?story_based={bool}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"save_sphere": {
			ID:          8,
			Description: "Searches for areas that have a save sphere.",
			Usage:       "?save_sphere={bool}",
			ExampleUses: []string{"?save_sphere=true", "?save_sphere=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"comp_sphere": {
			ID:          9,
			Description: "Searches for areas that contain an al bhed compilation sphere.",
			Usage:       "?comp_sphere={bool}",
			ExampleUses: []string{"?comp_sphere=true", "?comp_sphere=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"airship": {
			ID:          10,
			Description: "Searches for areas where you get dropped off when visiting with the airship.",
			Usage:       "?airship={bool}",
			ExampleUses: []string{"?airship=true", "?airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"chocobo": {
			ID:          11,
			Description: "Searches for areas where you can ride a chocobo.",
			Usage:       "?chocobo={bool}",
			ExampleUses: []string{"?chocobo=true", "?chocobo=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"characters": {
			ID:          12,
			Description: "Searches for areas where a character permanently joins the party.",
			Usage:       "?characters={bool}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          13,
			Description: "Searches for areas where a new aeon is acquired.",
			Usage:       "?aeons={bool}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          14,
			Description: "Searches for areas that have monsters.",
			Usage:       "?monsters={bool}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          15,
			Description: "Searches for areas that have bosses.",
			Usage:       "?boss_fights={bool}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          16,
			Description: "Searches for areas that have shops.",
			Usage:       "?shops={bool}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          17,
			Description: "Searches for areas that have treasures.",
			Usage:       "?treasures={bool}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          18,
			Description: "Searchces for areas that feature sidequests.",
			Usage:       "?sidequests={bool}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          19,
			Description: "Searches for areas that feature fmv sequences.",
			Usage:       "?fmvs={bool}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.areas = params
}

func (cfg *Config) initArenaCreationsParams() {
	params := map[string]QueryType{
		"category": {
			ID:          1,
			Description: "Searches for monster formations with the specified arena-creation-category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=species", "?category=3"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ArenaCreationCategory.lookup,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.arenaCreations = params
}

func (cfg *Config) initBlitzballPrizesParams() {
	params := map[string]QueryType{
		"category": {
			ID:          1,
			Description: "Searches for blitzball prize tables with the specified blitzball-tournament-category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=league", "?category=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.BlitzballTournamentCategory.lookup,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.blitzballPrizes = params
}

func (cfg *Config) initCharactersParams() {
	params := map[string]QueryType{
		"story_based": {
			ID:          1,
			Description: "Searches for characters that are only playable during certain sections of the story.",
			Usage:       "?story_based={bool}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"underwater": {
			ID:          2,
			Description: "Searches for characters that can fight underwater.",
			Usage:       "?underwater={bool}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.characters = params
}

func (cfg *Config) initCharacterClassesParams() {
	params := map[string]QueryType{
		"category": {
			ID:          1,
			Description: "Searches for character classes with the specified category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=group", "?category=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CharacterClassCategory.lookup,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.characterClasses = params
}

func (cfg *Config) initFMVsParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for fmvs that are played within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.fmvs = params
}

func (cfg *Config) initMonstersParams() {
	params := map[string]QueryType{
		"kimahri_stats": {
			ID:          1,
			Description: "Calculate the stats of Biran and Yenke Ronso that are based on Kimahri's stats. These are: HP, strength, magic, agility. If unused, their stats are based on Kimahri's base stats.",
			Usage:       "?kimahri_stats={stat}={int},...",
			ExampleUses: []string{"?kimahri_stats=hp=3000,strength=25,magic=30,agility=40", "?kimahri_stats=hp=15000,agility=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{167, 168},
		},
		"aeon_stats": {
			ID:          2,
			Description: "Replace the stats of Possessed Aeons with your own. All stats are replaceable, except for MP and luck. If unused, their stats are based on your own Aeon's base stats.",
			Usage:       "?aeon_stats={stat}={int},...",
			ExampleUses: []string{"?aeon_stats=hp=3000,strength=75,defense=50,magic=30,agility=20", "?aeon_stats=accuracy=150,magic_defense=255"},
			ForList:     false,
			ForSingle:   true,
			AllowedIDs:  []int32{216, 217, 218, 219, 220, 221, 222, 223, 224, 225},
		},
		"altered_state": {
			ID:          3,
			Description: "If a monster has altered states, apply the changes of an altered state to that monster. The default state will show up as the first altered state in the new entry.",
			Usage:       "?altered_state={id}",
			ExampleUses: []string{"?altered_state=1"},
			ForList:     false,
			ForSingle:   true,
		},
		"omnis_elements": {
			ID:            4,
			Description:   "Calculate the elemental affinities of Seymour Omnis by using exactly four of the letters 'f' (fire), 'l' (lightning), 'w' (water) and 'i' (ice). The letters represent the Mortiphasms pointing at Omnis. 0 of a color = 'neutral', 1 = 'halved', 2 = 'immune', 3 = 'absorb', 4 = 'absorb' + 'weak' to opposing element. The order of the letters doesn't matter.",
			Usage:         "?omnis_elements={4xf|l|w|i}",
			ExampleUses:   []string{"?omnis_elements=ifil", "?omnis_elements=llll", "?omnis_elements=wfwf"},
			ForList:       false,
			ForSingle:     true,
			AllowedIDs:    []int32{211},
			AllowedValues: []string{"f", "l", "w", "i"},
		},
		"elemental_resists": {
			ID:          5,
			Description: "Searches for monsters that have the specified elemental affinities.",
			Usage:       "?elemental_resists={element|id}={affinity|id},...",
			ExampleUses: []string{"?elemental_resists=fire=weak,water=absorb", "?elemental_resists=1=3,2=4"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements"), createListURL(cfg, "affinities")},
		},
		"status_resists": {
			ID:          6,
			Description: "Searches for monsters that resist or are immune to the specified status condition(s). You can optionally use the 'resistance' parameter to specify the minimum resistance. By default, the minimum resistance is 1.",
			Usage:       "?status_resists={id},...",
			ExampleUses: []string{"status_resists=1,4"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"resistance": {
			ID:              7,
			Description:     "Specifies the minimum resistance for the 'status_resists' parameter. Resistance is an integer ranging from 1 to 254 (immune). The value 'immune' can also be used, which counts as 254.",
			Usage:           "status_resists={id},...&resistance={int|'immune'}",
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
			DefaultVal: h.GetIntPtr(1),
		},
		"item": {
			ID:          8,
			Description: "Searches for monsters that have the specified item as loot. Can be specified further with the 'method' parameter.",
			Usage:       "?item={id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             9,
			Description:    "Specifies the method of acquisition for the 'item' parameter.",
			Usage:          "?item={id}&method={method_name}",
			ExampleUses:    []string{"?item=45&method=steal"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"steal", "drop", "bribe", "other"},
		},
		"auto_ability": {
			ID:          10,
			Description: "Searches for monsters that drop the specified auto-ability.",
			Usage:       "?auto_ability={id}",
			ExampleUses: []string{"?auto_ability=16"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		"is_forced": {
			ID:             11,
			Description:    "Specifies whether the auto-ability a monster drops is forced or not when using the 'auto_ability' parameter.",
			Usage:          "?auto_ability={id}&is_forced={bool}",
			ExampleUses:    []string{"?auto_ability=45&is_forced=false"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"auto_ability"},
		},
		"empty_slots": {
			ID:              12,
			Description:     "Searches for monsters that can drop equipment with the specified amount of empty slots and no other auto-abilities attached to it.",
			Usage:           "?empty_slots={int}",
			ExampleUses:     []string{"?empty_slots=3"},
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		"ronso_rage": {
			ID:          13,
			Description: "Searches for monsters that can teach the specified ronso rage to Kimahri.",
			Usage:       "?ronso_rage={id}",
			ExampleUses: []string{"?ronso_rage=5"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "ronso-rages")},
		},
		"location": {
			ID:          14,
			Description: "Searches for monsters that appear within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          15,
			Description: "Searches for monsters that appear within the specified sublocation.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"area": {
			ID:          16,
			Description: "Searches for monsters that appear within the specified area.",
			Usage:       "?area={id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		"distance": {
			ID:              17,
			Description:     "Searches for monsters with the specified distance. Distance is an integer ranging from 1 (close) to 4 (very far/infinite).",
			Usage:           "?distance={int}",
			ExampleUses:     []string{"?distance=3"},
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{1, 4},
		},
		"post_airship": {
			ID:          18,
			Description: "Searches for monsters that only appear after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"story_based": {
			ID:          19,
			Description: "Searches for monsters that only appear during certain sections of the story.",
			Usage:       "?story_based={bool}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"repeatable": {
			ID:          20,
			Description: "Searches for monsters that can be farmed.",
			Usage:       "?repeatable={bool}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"capture": {
			ID:          21,
			Description: "Searches for monsters that can be captured.",
			Usage:       "?capture={bool}",
			ExampleUses: []string{"?capture=true", "?capture=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"has_overdrive": {
			ID:          22,
			Description: "Searches for monsters that have an overdrive gauge.",
			Usage:       "?has_overdrive={bool}",
			ExampleUses: []string{"?has_overdrive=true", "?has_overdrive=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"underwater": {
			ID:          23,
			Description: "Searches for monsters that are fought underwater.",
			Usage:       "?underwater={bool}",
			ExampleUses: []string{"?underwater=true", "?underwater=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"zombie": {
			ID:          24,
			Description: "Searches for monsters that are zombies.",
			Usage:       "?zombie={bool}",
			ExampleUses: []string{"?zombie=true", "?zombie=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"species": {
			ID:          25,
			Description: "Searches for monsters of the specified species.",
			Usage:       "?species={name|id}",
			ExampleUses: []string{"?species=wyrm", "?species=5"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterSpecies.lookup,
			References:  []string{createListURL(cfg, "monster-species")},
		},
		"creation_area": {
			ID:          26,
			Description: "Searches for monsters that need to be captured in the specified area to create its area creation.",
			Usage:       "?creation_area={name|id}",
			ExampleUses: []string{"?creation_area=thunder-plains", "?creation_area=5"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.CreationArea.lookup,
		},
		"type": {
			ID:          27,
			Description: "Searches for monsters that are of the specified monster-type.",
			Usage:       "?type={name|id}",
			ExampleUses: []string{"?type=boss", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterCategory.lookup,
			References:  []string{createListURL(cfg, "monster-type")},
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.monsters = params
}

func (cfg *Config) initMonsterFormationsParams() {
	params := map[string]QueryType{
		"monster": {
			ID:          1,
			Description: "Searches for monster formations that feature the specified monster.",
			Usage:       "?monster={id}",
			ExampleUses: []string{"?monster=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		"category": {
			ID:          2,
			Description: "Searches for monster formations with the specified monster-formation-category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=boss-fight", "?category=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.MonsterFormationCategory.lookup,
			References:  []string{createListURL(cfg, "monster-formation-category")},
		},
		"location": {
			ID:          3,
			Description: "Searches for monster formations that appear within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          4,
			Description: "Searches for monster formations that appear within the specified sublocation.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"area": {
			ID:          5,
			Description: "Searches for monster formations that appear within the specified area.",
			Usage:       "?area={id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		"post_airship": {
			ID:          6,
			Description: "Searches for monster formations that are only available after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"story_based": {
			ID:          7,
			Description: "Searches for monster formations that are only available during certain sections of the story.",
			Usage:       "?story_based={bool}",
			ExampleUses: []string{"?story_based=true", "?story_based=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"repeatable": {
			ID:          8,
			Description: "Searches for monster formations that can be triggered more than once.",
			Usage:       "?repeatable={bool}",
			ExampleUses: []string{"?repeatable=true", "?repeatable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"ambush": {
			ID:          9,
			Description: "Searches for monster formations that are forced ambushes.",
			Usage:       "?ambush={bool}",
			ExampleUses: []string{"?ambush=true", "?ambush=false"},
			ForList:     true,
			ForSingle:   false,
		},
		
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.monsterFormations = params
}

func (cfg *Config) initOverdriveModesParams() {
	params := map[string]QueryType{
		"type": {
			ID:          1,
			Description: "Searches for overdrive-modes that are of the specified overdrive-mode-type.",
			Usage:       "?type={name|id}",
			ExampleUses: []string{"?type=per-action", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.OverdriveModeType.lookup,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.overdriveModes = params
}

func (cfg *Config) initAbilitiesParams() {
	params := map[string]QueryType{
		"type": {
			ID:          1,
			Description: "Searches for abilities that are of the specified ability type.",
			Usage:       "?type={name|id}",
			ExampleUses: []string{"?type=player-ability", "?type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AbilityType.lookup,
			References:  []string{createListURL(cfg, "ability-type")},
		},
		"rank": {
			ID:          2,
			Description: "Searches for abilities with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"copycat": {
			ID:          3,
			Description: "Searches for abilities that can be copied by 'copycat'.",
			Usage:       "?copycat={bool}",
			ExampleUses: []string{"?copycat=true", "?can_copycat=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"help_bar": {
			ID:          4,
			Description: "Searches for abilities whose names appear in the help bar.",
			Usage:       "?help_bar={bool}",
			ExampleUses: []string{"?help_bar=true", "?help_bar=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monster": {
			ID:          5,
			Description: "Searches for abilities that can be used by the specified monster.",
			Usage:       "?monster={id}",
			ExampleUses: []string{"?monster=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		"target_type": {
			ID:          6,
			Description: "Searches for abilities with the specified target type.",
			Usage:       "?target_type={name|id}",
			ExampleUses: []string{"?target_type=3", "?target_type=single-target"},
			ForList:     true,
			ForSingle:   false,
		},
		"user_atk": {
			ID:          7,
			Description: "Searches for abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Usage:       "?user_atk={bool}",
			ExampleUses: []string{"?user_atk=true", "?user_atk=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"darkable": {
			ID:          8,
			Description: "Searches for abilities that are affected by 'darkness'.",
			Usage:       "?darkable={bool}",
			ExampleUses: []string{"?darkable=true", "?darkable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"silenceable": {
			ID:          9,
			Description: "Searches for abilities that are affected by 'silence'.",
			Usage:       "?silenceable={bool}",
			ExampleUses: []string{"?silenceable=true", "?silenceable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"reflectable": {
			ID:          10,
			Description: "Searches for abilities that are affected by 'reflect'.",
			Usage:       "?reflectable={bool}",
			ExampleUses: []string{"?reflectable=true", "?reflectable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"attack_type": {
			ID:          11,
			Description: "Searches for abilities with battle interactions of the specified attack type.",
			Usage:       "?attack_type={name|id}",
			ExampleUses: []string{"?attack_type=attack", "?attack_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		"damage_type": {
			ID:          12,
			Description: "Searches for abilities that deal the specified type of damage.",
			Usage:       "?damage_type={name|id}",
			ExampleUses: []string{"?damage_type=3", "?damage_type=physical"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		"damage_formula": {
			ID:          13,
			Description: "Searches for abilities that use the specified formula to calculate their damage.",
			Usage:       "?damage_formula={name|id}",
			ExampleUses: []string{"?damage_formula=str-vs-def", "?attack_type=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		"can_crit": {
			ID:          14,
			Description: "Searches for abilities that can land critical hits.",
			Usage:       "?can_crit={bool}",
			ExampleUses: []string{"?can_crit=true", "?can_crit=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"bdl": {
			ID:          15,
			Description: "Searches for abilities that can break the damage cap of 9999.",
			Usage:       "?bdl={bool}",
			ExampleUses: []string{"?bdl=true", "?bdl=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"element": {
			ID:          16,
			Description: "Searches for abilities that deal elemental damage based on the specified element.",
			Usage:       "?element={name|id}",
			ExampleUses: []string{"?element=3", "?element=fire"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		"delay": {
			ID:          17,
			Description: "Searches for abilities that deal delay.",
			Usage:       "?delay={bool}",
			ExampleUses: []string{"?delay=true", "?delay=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"status_inflict": {
			ID:          18,
			Description: "Searches for abilities that can inflict the specified status condition.",
			Usage:       "?status_inflict={id}",
			ExampleUses: []string{"?status_inflict=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"status_remove": {
			ID:          19,
			Description: "Searches for abilities that can remove the specified status condition.",
			Usage:       "?status_remove={id}",
			ExampleUses: []string{"?status_remove=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"stat_changes": {
			ID:          20,
			Description: "Searches for abilities that cause stat changes.",
			Usage:       "?stat_changes={bool}",
			ExampleUses: []string{"?stat_changes=true", "?stat_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"mod_changes": {
			ID:          21,
			Description: "Searches for abilities that cause modifier changes.",
			Usage:       "?mod_changes={bool}",
			ExampleUses: []string{"?mod_changes=true", "?mod_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.abilities = params
}

func (cfg *Config) initEnemyAbilitiesParams() {
	params := map[string]QueryType{
		"rank": {
			ID:          1,
			Description: "Searches for enemy abilities with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"help_bar": {
			ID:          2,
			Description: "Searches for enemy abilities whose names appear in the help bar.",
			Usage:       "?help_bar={bool}",
			ExampleUses: []string{"?help_bar=true", "?help_bar=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monster": {
			ID:          3,
			Description: "Searches for enemy abilities that can be used by the specified monster.",
			Usage:       "?monster={id}",
			ExampleUses: []string{"?monster=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "monsters")},
		},
		"target_type": {
			ID:          4,
			Description: "Searches for enemy abilities with the specified target type.",
			Usage:       "?target_type={name|id}",
			ExampleUses: []string{"?target_type=3", "?target_type=single-target"},
			ForList:     true,
			ForSingle:   false,
		},
		"darkable": {
			ID:          5,
			Description: "Searches for enemy abilities that are affected by 'darkness'.",
			Usage:       "?darkable={bool}",
			ExampleUses: []string{"?darkable=true", "?darkable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"silenceable": {
			ID:          6,
			Description: "Searches for enemy abilities that are affected by 'silence'.",
			Usage:       "?silenceable={bool}",
			ExampleUses: []string{"?silenceable=true", "?silenceable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"reflectable": {
			ID:          7,
			Description: "Searches for enemy abilities that are affected by 'reflect'.",
			Usage:       "?reflectable={bool}",
			ExampleUses: []string{"?reflectable=true", "?reflectable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"attack_type": {
			ID:          8,
			Description: "Searches for enemy abilities with battle interactions of the specified attack type.",
			Usage:       "?attack_type={name|id}",
			ExampleUses: []string{"?attack_type=attack", "?attack_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		"damage_type": {
			ID:          9,
			Description: "Searches for enemy abilities that deal the specified type of damage.",
			Usage:       "?damage_type={name|id}",
			ExampleUses: []string{"?damage_type=3", "?damage_type=physical"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		"damage_formula": {
			ID:          10,
			Description: "Searches for enemy abilities that use the specified formula to calculate their damage.",
			Usage:       "?damage_formula={name|id}",
			ExampleUses: []string{"?damage_formula=str-vs-def", "?attack_type=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		"can_crit": {
			ID:          11,
			Description: "Searches for enemy abilities that can land critical hits.",
			Usage:       "?can_crit={bool}",
			ExampleUses: []string{"?can_crit=true", "?can_crit=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"bdl": {
			ID:          12,
			Description: "Searches for enemy abilities that can break the damage cap of 9999.",
			Usage:       "?bdl={bool}",
			ExampleUses: []string{"?bdl=true", "?bdl=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"element": {
			ID:          13,
			Description: "Searches for enemy abilities that deal elemental damage based on the specified element.",
			Usage:       "?element={name|id}",
			ExampleUses: []string{"?element=3", "?element=fire"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		"delay": {
			ID:          14,
			Description: "Searches for enemy abilities that deal delay.",
			Usage:       "?delay={bool}",
			ExampleUses: []string{"?delay=true", "?delay=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"status_inflict": {
			ID:          15,
			Description: "Searches for enemy abilities that can inflict the specified status condition.",
			Usage:       "?status_inflict={id}",
			ExampleUses: []string{"?status_inflict=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"status_remove": {
			ID:          16,
			Description: "Searches for enemy abilities that can remove the specified status condition.",
			Usage:       "?status_remove={id}",
			ExampleUses: []string{"?status_remove=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.enemyAbilities = params
}

func (cfg *Config) initItemAbilitiesParams() {
	params := map[string]QueryType{
		"category": {
			ID:          1,
			Description: "Searches for item abilities that are of the specified item category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=healing", "?category=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ItemCategory.lookup,
			References:  []string{createListURL(cfg, "item-category")},
		},
		"outside_battle": {
			ID:          2,
			Description: "Searches for item abilities that can be used outside of battle, in the 'abilities' menu.",
			Usage:       "?outside_battle={bool}",
			ExampleUses: []string{"?outside_battle=true", "?outside_battle=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"related_stat": {
			ID:          3,
			Description: "Searches for item abilities that are related to the specified stat.",
			Usage:       "?related_stat={name|id}",
			ExampleUses: []string{"?related_stat=3", "?related_stat=hp"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		"target_type": {
			ID:          4,
			Description: "Searches for item abilities with the specified target type.",
			Usage:       "?target_type={name|id}",
			ExampleUses: []string{"?target_type=3", "?target_type=single-target"},
			ForList:     true,
			ForSingle:   false,
		},
		"attack_type": {
			ID:          5,
			Description: "Searches for item abilities with battle interactions of the specified attack type.",
			Usage:       "?attack_type={name|id}",
			ExampleUses: []string{"?attack_type=attack", "?attack_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		"damage_formula": {
			ID:          6,
			Description: "Searches for item abilities that use the specified formula to calculate their damage.",
			Usage:       "?damage_formula={name|id}",
			ExampleUses: []string{"?damage_formula=str-vs-def", "?attack_type=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		"element": {
			ID:          7,
			Description: "Searches for item abilities that deal elemental damage based on the specified element.",
			Usage:       "?element={name|id}",
			ExampleUses: []string{"?element=3", "?element=fire"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		"delay": {
			ID:          8,
			Description: "Searches for item abilities that deal delay.",
			Usage:       "?delay={bool}",
			ExampleUses: []string{"?delay=true", "?delay=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"status_inflict": {
			ID:          9,
			Description: "Searches for item abilities that can inflict the specified status condition.",
			Usage:       "?status_inflict={id}",
			ExampleUses: []string{"?status_inflict=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"status_remove": {
			ID:          10,
			Description: "Searches for item abilities that can remove the specified status condition.",
			Usage:       "?status_remove={id}",
			ExampleUses: []string{"?status_remove=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"stat_changes": {
			ID:          11,
			Description: "Searches for item abilities that cause stat changes.",
			Usage:       "?stat_changes={bool}",
			ExampleUses: []string{"?stat_changes=true", "?stat_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"mod_changes": {
			ID:          12,
			Description: "Searches for item abilities that cause modifier changes.",
			Usage:       "?mod_changes={bool}",
			ExampleUses: []string{"?mod_changes=true", "?mod_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.itemAbilities = params
}

func (cfg *Config) initOverdriveAbilitiesParams() {
	params := map[string]QueryType{
		"rank": {
			ID:          1,
			Description: "Searches for overdrive abilities with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"user": {
			ID:          2,
			Description: "Searches for overdrive abilities that are learned by the specified character class.",
			Usage:       "?user={name|id}",
			ExampleUses: []string{"?user=3", "?user=characters"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		"related_stat": {
			ID:          3,
			Description: "Searches for overdrive abilities that are related to the specified stat.",
			Usage:       "?related_stat={name|id}",
			ExampleUses: []string{"?related_stat=3", "?related_stat=hp"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		"target_type": {
			ID:          4,
			Description: "Searches for overdrive abilities with the specified target type.",
			Usage:       "?target_type={name|id}",
			ExampleUses: []string{"?target_type=3", "?target_type=single-target"},
			ForList:     true,
			ForSingle:   false,
		},
		"attack_type": {
			ID:          5,
			Description: "Searches for overdrive abilities with battle interactions of the specified attack type.",
			Usage:       "?attack_type={name|id}",
			ExampleUses: []string{"?attack_type=attack", "?attack_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		"damage_formula": {
			ID:          6,
			Description: "Searches for overdrive abilities that use the specified formula to calculate their damage.",
			Usage:       "?damage_formula={name|id}",
			ExampleUses: []string{"?damage_formula=str-vs-def", "?attack_type=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		"can_crit": {
			ID:          7,
			Description: "Searches for overdrive abilities that can land critical hits.",
			Usage:       "?can_crit={bool}",
			ExampleUses: []string{"?can_crit=true", "?can_crit=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"element": {
			ID:          8,
			Description: "Searches for overdrive abilities that deal elemental damage based on the specified element.",
			Usage:       "?element={name|id}",
			ExampleUses: []string{"?element=3", "?element=fire"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		"delay": {
			ID:          9,
			Description: "Searches for overdrive abilities that deal delay.",
			Usage:       "?delay={bool}",
			ExampleUses: []string{"?delay=true", "?delay=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"status_inflict": {
			ID:          10,
			Description: "Searches for overdrive abilities that can inflict the specified status condition.",
			Usage:       "?status_inflict={id}",
			ExampleUses: []string{"?status_inflict=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"status_remove": {
			ID:          11,
			Description: "Searches for overdrive abilities that can remove the specified status condition.",
			Usage:       "?status_remove={id}",
			ExampleUses: []string{"?status_remove=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"stat_changes": {
			ID:          12,
			Description: "Searches for overdrive abilities that cause stat changes.",
			Usage:       "?stat_changes={bool}",
			ExampleUses: []string{"?stat_changes=true", "?stat_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"mod_changes": {
			ID:          13,
			Description: "Searches for overdrive abilities that cause modifier changes.",
			Usage:       "?mod_changes={bool}",
			ExampleUses: []string{"?mod_changes=true", "?mod_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.overdriveAbilities = params
}

func (cfg *Config) initPlayerAbilitiesParams() {
	params := map[string]QueryType{
		"ability_user": {
			ID:              1,
			Description:     "If a player ability is based on a user's attack, this parameter modifies its accuracy, range, shatter rate and power based on the given user. User can be a character or an aeon. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Usage:           "?ability_user={type}:{name|id}",
			ExampleUses:     []string{"?ability_user=character:wakka", "?ability_user=aeon:valefor", "?ability_user=character:2"},
			ForList:         false,
			ForSingle:       true,
			AllowedResTypes: []string{"character", "aeon"},
		},
		"bomb_wpn": {
			ID:              2,
			Description:     "If a player ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Usage:           "?ability_user={type}:{name|id}&bomb_wpn={bool}",
			ExampleUses:     []string{"?ability_user=character:wakka&bomb_wpn=true"},
			ForList:         false,
			ForSingle:       true,
			RequiredParams:  []string{"ability_user"},
		},
		"rank": {
			ID:          3,
			Description: "Searches for player abilities with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"copycat": {
			ID:          4,
			Description: "Searches for player abilities that can be copied by 'copycat'.",
			Usage:       "?copycat={bool}",
			ExampleUses: []string{"?copycat=true", "?can_copycat=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"help_bar": {
			ID:          5,
			Description: "Searches for player abilities whose names appear in the help bar.",
			Usage:       "?help_bar={bool}",
			ExampleUses: []string{"?help_bar=true", "?help_bar=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"category": {
			ID:          6,
			Description: "Searches for player abilities that are of the specified player ability category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=black-magic", "?category=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.PlayerAbilityCategory.lookup,
			References:  []string{createListURL(cfg, "player-ability-category")},
		},
		"outside_battle": {
			ID:          7,
			Description: "Searches for player abilities that can be used outside of battle, in the 'abilities' menu.",
			Usage:       "?outside_battle={bool}",
			ExampleUses: []string{"?outside_battle=true", "?outside_battle=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"mp": {
			ID:          8,
			Description: "Searches for player abilities with the specified mp cost.",
			Usage:       "?mp={int}",
			ExampleUses: []string{"?mp=16"},
			ForList:     true,
			ForSingle:   false,
		},
		"mp_min": {
			ID:          9,
			Description: "Searches for player abilities with an mp cost that is equal or more than the specified amount.",
			Usage:       "?mp_min={int}",
			ExampleUses: []string{"?mp_min=16"},
			ForList:     true,
			ForSingle:   false,
		},
		"mp_max": {
			ID:          10,
			Description: "Searches for player abilities with an mp cost that is equal or less than the specified amount.",
			Usage:       "?mp_max={int}",
			ExampleUses: []string{"?mp_max=16"},
			ForList:     true,
			ForSingle:   false,
		},
		"related_stat": {
			ID:          11,
			Description: "Searches for player abilities that are related to the specified stat.",
			Usage:       "?related_stat={name|id}",
			ExampleUses: []string{"?related_stat=3", "?related_stat=hp"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		"user": {
			ID:          12,
			Description: "Searches for player abilities that are learned by the specified character class.",
			Usage:       "?user={name|id}",
			ExampleUses: []string{"?user=3", "?user=characters"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		"std_sg": {
			ID:          13,
			Description: "Searches for player abilities that are located on the specified character's standard sphere grid.",
			Usage:       "?std_sg={name|id}",
			ExampleUses: []string{"?std_sg=3", "?std_sg=tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "characters")},
		},
		"exp_sg": {
			ID:          14,
			Description: "Searches for player abilities that are located on the specified character's expert sphere grid.",
			Usage:       "?exp_sg={name|id}",
			ExampleUses: []string{"?exp_sg=3", "?exp_sg=tidus"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "characters")},
		},
		"learn_item": {
			ID:          15,
			Description: "Searches for player abilities an aeon can learn via the specified item.",
			Usage:       "?learn_item={id}",
			ExampleUses: []string{"?learn_item=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"target_type": {
			ID:          16,
			Description: "Searches for player abilities with the specified target type.",
			Usage:       "?target_type={name|id}",
			ExampleUses: []string{"?target_type=3", "?target_type=single-target"},
			ForList:     true,
			ForSingle:   false,
		},
		"user_atk": {
			ID:          17,
			Description: "Searches for player abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Usage:       "?user_atk={bool}",
			ExampleUses: []string{"?user_atk=true", "?user_atk=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"darkable": {
			ID:          18,
			Description: "Searches for player abilities that are affected by 'darkness'.",
			Usage:       "?darkable={bool}",
			ExampleUses: []string{"?darkable=true", "?darkable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"silenceable": {
			ID:          19,
			Description: "Searches for player abilities that are affected by 'silence'.",
			Usage:       "?silenceable={bool}",
			ExampleUses: []string{"?silenceable=true", "?silenceable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"reflectable": {
			ID:          20,
			Description: "Searches for player abilities that are affected by 'reflect'.",
			Usage:       "?reflectable={bool}",
			ExampleUses: []string{"?reflectable=true", "?reflectable=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"attack_type": {
			ID:          21,
			Description: "Searches for player abilities with battle interactions of the specified attack type.",
			Usage:       "?attack_type={name|id}",
			ExampleUses: []string{"?attack_type=attack", "?attack_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.AttackType.lookup,
			References:  []string{createListURL(cfg, "attack-type")},
		},
		"damage_type": {
			ID:          22,
			Description: "Searches for player abilities that deal the specified type of damage.",
			Usage:       "?damage_type={name|id}",
			ExampleUses: []string{"?damage_type=3", "?damage_type=physical"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageType.lookup,
			References:  []string{createListURL(cfg, "damage-type")},
		},
		"damage_formula": {
			ID:          23,
			Description: "Searches for player abilities that use the specified formula to calculate their damage.",
			Usage:       "?damage_formula={name|id}",
			ExampleUses: []string{"?damage_formula=str-vs-def", "?attack_type=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.DamageFormula.lookup,
			References:  []string{createListURL(cfg, "damage-formula")},
		},
		"element": {
			ID:          24,
			Description: "Searches for player abilities that deal elemental damage based on the specified element.",
			Usage:       "?element={name|id}",
			ExampleUses: []string{"?element=3", "?element=fire"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "elements")},
		},
		"delay": {
			ID:          25,
			Description: "Searches for player abilities that deal delay.",
			Usage:       "?delay={bool}",
			ExampleUses: []string{"?delay=true", "?delay=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"status_inflict": {
			ID:          26,
			Description: "Searches for player abilities that can inflict the specified status condition.",
			Usage:       "?status_inflict={id}",
			ExampleUses: []string{"?status_inflict=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"status_remove": {
			ID:          27,
			Description: "Searches for player abilities that can remove the specified status condition.",
			Usage:       "?status_remove={id}",
			ExampleUses: []string{"?status_remove=3"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "status-conditions")},
		},
		"stat_changes": {
			ID:          28,
			Description: "Searches for player abilities that cause stat changes.",
			Usage:       "?stat_changes={bool}",
			ExampleUses: []string{"?stat_changes=true", "?stat_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"mod_changes": {
			ID:          29,
			Description: "Searches for player abilities that cause modifier changes.",
			Usage:       "?mod_changes={bool}",
			ExampleUses: []string{"?mod_changes=true", "?mod_changes=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.playerAbilities = params
}

func (cfg *Config) initTriggerCommandsParams() {
	params := map[string]QueryType{
		"ability_user": {
			ID:              1,
			Description:     "If a trigger command is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. User can be a character or an aeon. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this command.",
			Usage:           "?ability_user={type}:{name|id}",
			ExampleUses:     []string{"?ability_user=character:wakka", "?ability_user=aeon:valefor", "?ability_user=character:2"},
			ForList:         false,
			ForSingle:       true,
			AllowedResTypes: []string{"character", "aeon"},
		},
		"bomb_wpn": {
			ID:              2,
			Description:     "If a trigger command is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Usage:           "?ability_user={type}:{name|id}&bomb_wpn={bool}",
			ExampleUses:     []string{"?ability_user=character:wakka&bomb_wpn=true"},
			ForList:         false,
			ForSingle:       true,
			RequiredParams:  []string{"ability_user"},
		},
		"related_stat": {
			ID:          3,
			Description: "Searches for trigger commands that are related to the specified stat.",
			Usage:       "?related_stat={name|id}",
			ExampleUses: []string{"?related_stat=3", "?related_stat=hp"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "stats")},
		},
		"user": {
			ID:          4,
			Description: "Searches for trigger commands that are learned by the specified character class.",
			Usage:       "?user={name|id}",
			ExampleUses: []string{"?user=3", "?user=characters"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.triggerCommands = params
}

func (cfg *Config) initUnspecifiedAbilitiesParams() {
	params := map[string]QueryType{
		"ability_user": {
			ID:              1,
			Description:     "If an unspecified ability is based on a user's attack, this parameter modifies the its accuracy, range, shatter rate and power based on the given user. User can be a character or an aeon. For characters, only the range is modified in the case of Wakka. Responds with an error, if the specified user can't learn this ability.",
			Usage:           "?ability_user={type}:{name|id}",
			ExampleUses:     []string{"?ability_user=character:wakka", "?ability_user=aeon:valefor", "?ability_user=character:2"},
			ForList:         false,
			ForSingle:       true,
			AllowedResTypes: []string{"character", "aeon"},
		},
		"bomb_wpn": {
			ID:              2,
			Description:     "If an unspecified ability is based on a user's attack, this parameter modifies its damage constant to be 18 instead of 16, since that is the power of weapons dropped by bombs specifically. Can only be used in combination with the 'ability_user' parameter and only takes effect, if the specified user is a character.",
			Usage:           "?ability_user={type}:{name|id}&bomb_wpn={bool}",
			ExampleUses:     []string{"?ability_user=character:wakka&bomb_wpn=true"},
			ForList:         false,
			ForSingle:       true,
			RequiredParams:  []string{"ability_user"},
		},
		"rank": {
			ID:          3,
			Description: "Searches for unspecified abilities with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"copycat": {
			ID:          4,
			Description: "Searches for unspecified abilities that can be copied by 'copycat'.",
			Usage:       "?copycat={bool}",
			ExampleUses: []string{"?copycat=true", "?can_copycat=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"help_bar": {
			ID:          5,
			Description: "Searches for unspecified abilities whose names appear in the help bar.",
			Usage:       "?help_bar={bool}",
			ExampleUses: []string{"?help_bar=true", "?help_bar=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"user": {
			ID:          6,
			Description: "Searches for unspecified abilities that are learned by the specified character class.",
			Usage:       "?user={name|id}",
			ExampleUses: []string{"?user=3", "?user=characters"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
		"user_atk": {
			ID:          7,
			Description: "Searches for unspecified abilities whose range, shatter rate, accuracy, and damage constant are based on the user's attack.",
			Usage:       "?user_atk={bool}",
			ExampleUses: []string{"?user_atk=true", "?user_atk=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.unspecifiedAbilities = params
}

func (cfg *Config) initOverdrivesParams() {
	params := map[string]QueryType{
		"rank": {
			ID:          1,
			Description: "Searches for overdrives with the specified rank.",
			Usage:       "?rank={int}",
			ExampleUses: []string{"?rank=3"},
			ForList:     true,
			ForSingle:   false,
		},
		"user": {
			ID:          2,
			Description: "Searches for overdrives that are learned by the specified character class.",
			Usage:       "?user={name|id}",
			ExampleUses: []string{"?user=3", "?user=characters"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "character-classes")},
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.overdrives = params
}

func (cfg *Config) initSubmenusParams() {
	params := map[string]QueryType{
		"topmenu": {
			ID:          1,
			Description: "Searches for submenus that are found within the specified topmenu.",
			Usage:       "?topmenu={name|id}",
			ExampleUses: []string{"?topmenu=2", "?topmenu=main"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "topmenus")},
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.submenus = params
}

func (cfg *Config) initSublocationsParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for sublocations that are located within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"item": {
			ID:          2,
			Description: "Searches for sublocations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             3,
			Description:    "Specifies the method of acquisition for the 'item' parameter.",
			Usage:          "?item={id}&method={method_name}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          4,
			Description: "Searches for sublocations where the specified key item can be acquired.",
			Usage:       "?key_item={id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		"characters": {
			ID:          5,
			Description: "Searches for sublocations where a character permanently joins the party.",
			Usage:       "?characters={bool}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          6,
			Description: "Searches for sublocations where a new aeon is acquired.",
			Usage:       "?aeons={bool}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          7,
			Description: "Searches for sublocations that have monsters.",
			Usage:       "?monsters={bool}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          8,
			Description: "Searches for sublocations that have bosses.",
			Usage:       "?boss_fights={bool}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          9,
			Description: "Searches for sublocations that have shops.",
			Usage:       "?shops={bool}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          10,
			Description: "Searches for sublocations that have treasures.",
			Usage:       "?treasures={bool}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          11,
			Description: "Searchces for sublocations that feature sidequests.",
			Usage:       "?sidequests={bool}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          12,
			Description: "Searches for sublocations that feature fmv sequences.",
			Usage:       "?fmvs={bool}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.sublocations = params
}

func (cfg *Config) initLocationsParams() {
	params := map[string]QueryType{
		"item": {
			ID:          1,
			Description: "Searches for locations where the specified item can be acquired. Can be specified further with the 'method' parameter.",
			Usage:       "?item={id}",
			ExampleUses: []string{"?item=45"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "items")},
		},
		"method": {
			ID:             2,
			Description:    "Specifies the method of acquisition for the 'item' parameter.",
			Usage:          "?item={id}&method={method_name}",
			ExampleUses:    []string{"?item=45&method=treasure"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"item"},
			AllowedValues:  []string{"monster", "treasure", "shop", "quest"},
		},
		"key_item": {
			ID:          3,
			Description: "Searches for locations where the specified key item can be acquired.",
			Usage:       "?key_item={id}",
			ExampleUses: []string{"?key_item=22"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "key-items")},
		},
		"characters": {
			ID:          4,
			Description: "Searches for locations where a character permanently joins the party.",
			Usage:       "?characters={bool}",
			ExampleUses: []string{"?characters=true", "?characters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"aeons": {
			ID:          5,
			Description: "Searches for locations where a new aeon is acquired.",
			Usage:       "?aeons={bool}",
			ExampleUses: []string{"?aeons=true", "?aeons=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"monsters": {
			ID:          6,
			Description: "Searches for locations that have monsters.",
			Usage:       "?monsters={bool}",
			ExampleUses: []string{"?monsters=true", "?monsters=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"boss_fights": {
			ID:          7,
			Description: "Searches for locations that have bosses.",
			Usage:       "?boss_fights={bool}",
			ExampleUses: []string{"?boss_fights=true", "?boss_fights=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"shops": {
			ID:          8,
			Description: "Searches for locations that have shops.",
			Usage:       "?shops={bool}",
			ExampleUses: []string{"?shops=true", "?shops=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"treasures": {
			ID:          9,
			Description: "Searches for locations that have treasures.",
			Usage:       "?treasures={bool}",
			ExampleUses: []string{"?treasures=true", "?treasures=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"sidequests": {
			ID:          10,
			Description: "Searchces for locations that feature sidequests.",
			Usage:       "?sidequests={bool}",
			ExampleUses: []string{"?sidequests=true", "?sidequests=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"fmvs": {
			ID:          11,
			Description: "Searches for locations that feature fmv sequences.",
			Usage:       "?fmvs={bool}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.locations = params
}

func (cfg *Config) initSidequestsParams() {
	params := map[string]QueryType{
		"post_airship": {
			ID:          1,
			Description: "Searches for sidequests that are only completable after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.sidequests = params
}

func (cfg *Config) initSubquestsParams() {
	params := map[string]QueryType{
		"post_airship": {
			ID:          1,
			Description: "Searches for subquests that are only completable after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.subquests = params
}

func (cfg *Config) initShopsParams() {
	params := map[string]QueryType{
		"category": {
			ID:          1,
			Description: "Searches for shops with the specified shop category.",
			Usage:       "?category={name|id}",
			ExampleUses: []string{"?category=oaka", "?category=4"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.ShopCategory.lookup,
			References:  []string{createListURL(cfg, "shop-category")},
		},
		"location": {
			ID:          2,
			Description: "Searches for shops that appear within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          3,
			Description: "Searches for shops that appear within the specified sublocation.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"auto_ability": {
			ID:          4,
			Description: "Searches for shops that offer equipment with the specified auto-ability.",
			Usage:       "?auto_ability={id}",
			ExampleUses: []string{"?auto_ability=7"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "auto-abilities")},
		},
		"empty_slots": {
			ID:              5,
			Description:     "Searches for shops that offer equipment with the specified amount of empty slots.",
			Usage:           "?empty_slots={int}",
			ExampleUses:     []string{"?empty_slots=3"},
			ForList:         true,
			ForSingle:       false,
			AllowedIntRange: []int{0, 4},
		},
		"character": {
			ID:             6,
			Description:    "Specifies the character the offered equipment is for when searching for shops with the 'auto_ability' or 'empty_slots' parameters.",
			Usage:          "?auto_ability={id}&character={id|name}",
			ExampleUses:    []string{"?auto_ability=111&character=wakka"},
			ForList:        true,
			ForSingle:      false,
			RequiredParams: []string{"auto_ability", "empty_slots"},
			References: 	[]string{createListURL(cfg, "characters")},
		},
		"shop_type": {
			ID:          7,
			Description: "Specifies whether the given auto-ability is sold before or after acquiring the airship when searching for shops with the 'auto_ability' or 'empty_slots' parameters.",
			Usage:       "?shop_type={value|id}",
			ExampleUses: []string{"?shop_type=pre-airship", "?shop_type=2"},
			ForList:     true,
			ForSingle:   false,
			RequiredParams: []string{"auto_ability", "empty_slots"},
			TypeLookup:  cfg.t.ShopType.lookup,
		},
		"items": {
			ID:          8,
			Description: "Searches for shops that offer items.",
			Usage:       "?items={bool}",
			ExampleUses: []string{"?items=true", "?items=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"equipment": {
			ID:          9,
			Description: "Searches for shops that offer equipment.",
			Usage:       "?equipment={bool}",
			ExampleUses: []string{"?equipment=true", "?equipment=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"pre_airship": {
			ID:          10,
			Description: "Searches for shops that are available before acquiring the airship.",
			Usage:       "?pre_airship={bool}",
			ExampleUses: []string{"?pre_airship=true", "?pre_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"post_airship": {
			ID:          11,
			Description: "Searches for shops that are available after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, true)
	cfg.q.shops = params
}

func (cfg *Config) initSongsParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for songs that are played within the specified location. Songs with special use cases are not included.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          2,
			Description: "Searches for songs that are played within the specified sublocation. Songs with special use cases are not included.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"area": {
			ID:          3,
			Description: "Searches for songs that are played within the specified area. Songs with special use cases are not included.",
			Usage:       "?area={id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		"fmvs": {
			ID:          4,
			Description: "Searches for songs that are played in fmvs.",
			Usage:       "?fmvs={bool}",
			ExampleUses: []string{"?fmvs=true", "?fmvs=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"special_use": {
			ID:          5,
			Description: "Searches for songs with a special use case.",
			Usage:       "?special_use={bool}",
			ExampleUses: []string{"?special_use=true", "?special_use=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"composer": {
			ID:          6,
			Description: "Searches for songs that were composed by the stated composer.",
			Usage:       "?composer={name|id}",
			ExampleUses: []string{"?composer=nobuo-uematsu", "?composer=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Composer.lookup,
		},
		"arranger": {
			ID:          7,
			Description: "Searches for songs that were arranged by the stated arranger.",
			Usage:       "?arranger={name|id}",
			ExampleUses: []string{"?arranger=nobuo-uematsu", "?arranger=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.Arranger.lookup,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.songs = params
}

func (cfg *Config) initTreasuresParams() {
	params := map[string]QueryType{
		"location": {
			ID:          1,
			Description: "Searches for treasures that appear within the specified location.",
			Usage:       "?location={id}",
			ExampleUses: []string{"?location=17"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "locations")},
		},
		"sublocation": {
			ID:          2,
			Description: "Searches for treasures that appear within the specified sublocation.",
			Usage:       "?sublocation={id}",
			ExampleUses: []string{"?sublocation=13"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "sublocations")},
		},
		"area": {
			ID:          3,
			Description: "Searches for treasures that appear within the specified area.",
			Usage:       "?area={id}",
			ExampleUses: []string{"?area=13", "?area=240"},
			ForList:     true,
			ForSingle:   false,
			References:  []string{createListURL(cfg, "areas")},
		},
		"loot_type": {
			ID:          4,
			Description: "Searches for treasures with the specified loot type.",
			Usage:       "?loot_type={name|id}",
			ExampleUses: []string{"?loot_type=item", "?loot_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.LootType.lookup,
			References:  []string{createListURL(cfg, "loot-type")},
		},
		"treasure_type": {
			ID:          5,
			Description: "Searches for treasures with the specified treasure type.",
			Usage:       "?treasure_type={name|id}",
			ExampleUses: []string{"?treasure_type=chest", "?treasure_type=2"},
			ForList:     true,
			ForSingle:   false,
			TypeLookup:  cfg.t.TreasureType.lookup,
		},
		"anima": {
			ID:          6,
			Description: "Searches for treasures that are necessary for getting Anima.",
			Usage:       "?anima={bool}",
			ExampleUses: []string{"?anima=true", "?anima=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"post_airship": {
			ID:          7,
			Description: "Searches for treasures that are only available after acquiring the airship.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
		"story_based": {
			ID:          7,
			Description: "Searches for treasures that are only available during certain sections of the story.",
			Usage:       "?post_airship={bool}",
			ExampleUses: []string{"?post_airship=true", "?post_airship=false"},
			ForList:     true,
			ForSingle:   false,
		},
	}

	params = cfg.completeQueryTypeInit(params, false)
	cfg.q.treasures = params
}
