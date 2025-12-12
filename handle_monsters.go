package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



type Monster struct {
	ID                   int32              	`json:"id"`
	Name                 string             	`json:"name"`
	Version              *int32             	`json:"version,omitempty"`
	Specification        *string            	`json:"specification,omitempty"`
	AppliedState         *AppliedState      	`json:"applied_state,omitempty"`
	AgilityParameters    AgilityParams      	`json:"agility_parameters"`
	Notes                *string            	`json:"notes,omitempty"`
	Species              NamedAPIResource       `json:"species"`
	IsStoryBased         bool               	`json:"is_story_based"`
	IsRepeatable			 bool					`json:"is_repeatable"`
	CanBeCaptured        bool               	`json:"can_be_captured"`
	AreaConquestLocation *string            	`json:"area_conquest_location,omitempty"`
	CTBIconType          NamedAPIResource       `json:"ctb_icon_type"`
	HasOverdrive         bool               	`json:"has_overdrive"`
	IsUnderwater         bool               	`json:"is_underwater"`
	IsZombie             bool               	`json:"is_zombie"`
	Distance             int32              	`json:"distance"`
	Properties           []NamedAPIResource 	`json:"properties"`
	AutoAbilities        []NamedAPIResource 	`json:"auto_abilities"`
	AP                   int32              	`json:"ap"`
	APOverkill           int32              	`json:"ap_overkill"`
	OverkillDamage       int32              	`json:"overkill_damage"`
	Gil                  int32              	`json:"gil"`
	StealGil             *int32             	`json:"steal_gil"`
	RonsoRages           []NamedAPIResource 	`json:"ronso_rages"`
	DoomCountdown        *int32             	`json:"doom_countdown"`
	PoisonRate           *float32           	`json:"poison_rate"`
	PoisonDamage         *int32             	`json:"poison_damage,omitempty"`
	ThreatenChance       *int32             	`json:"threaten_chance"`
	ZanmatoLevel         int32              	`json:"zanmato_level"`
	MonsterArenaPrice    *int32             	`json:"monster_arena_price,omitempty"`
	SensorText           *string            	`json:"sensor_text"`
	ScanText             *string            	`json:"scan_text"`
	Locations            []LocationAPIResource	`json:"locations"`
	Formations			 []UnnamedAPIResource	`json:"formations"`
	BaseStats            []BaseStat         	`json:"base_stats"`
	Items                *MonsterItems      	`json:"items"`
	BribeChances         []BribeChance      	`json:"bribe_chances,omitempty"`
	Equipment            *MonsterEquipment  	`json:"equipment"`
	ElemResists          []ElementalResist  	`json:"elem_resists"`
	StatusImmunities     []NamedAPIResource 	`json:"status_immunities"`
	StatusResists        []StatusResist     	`json:"status_resists"`
	AlteredStates        []AlteredState     	`json:"altered_states"`
	Abilities            []MonsterAbility   	`json:"abilities"`
}

func (cfg *apiConfig) handleMonsters(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/monsters/")
	segments := strings.Split(path, "/")

	if path == "" {
		cfg.handleMonstersRetrieve(w, r)
		return
	}
	// this whole thing can probably be generalized
	switch len(segments) {
	case 1:
		// /api/monsters/{name or id}
		segment := segments[0]

		input, err := parseSingleSegmentResource(segment, cfg.l.Monsters)
		if handleHTTPError(w, err) {
			return
		}

		if input.Name != "" {
			resources, err := cfg.getMultipleMonsters(r, input.Name)
			if handleHTTPError(w, err) {
				return
			}
			respondWithJSON(w, http.StatusMultipleChoices, resources)
			return
		}

		monster, err := cfg.getMonster(r, input.ID)
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusOK, monster)
		return

	case 2:
		// /api/monsters/{name}/{version}

		name := segments[0]
		versionStr := segments[1]
		var input parseResponse
		var err error

		// /api/monsters/{name}/
		if versionStr == "" {
			input, err = parseSingleSegmentResource(name, cfg.l.Monsters)
			if handleHTTPError(w, err) {
				return
			}
			if input.Name != "" {
				resources, err := cfg.getMultipleMonsters(r, input.Name)
				if handleHTTPError(w, err) {
					return
				}
				respondWithJSON(w, http.StatusMultipleChoices, resources)
				return
			}
		} else {
			input, err = parseNameVersionResource(name, versionStr, cfg.l.Monsters)
			if handleHTTPError(w, err) {
				return
			}
		}

		monster, err := cfg.getMonster(r, input.ID)
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusOK, monster)
		return

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/monsters/{name or id}, or /api/monsters/{name}/{version}`, nil)
		return
	}
}



func (cfg *apiConfig) getMonster(r *http.Request, id int32) (Monster, error) {
	dbMonster, err := cfg.db.GetMonster(r.Context(), id)
	if err != nil {
		return Monster{}, newHTTPError(http.StatusNotFound, "Couldn't get Monster. Monster with this ID doesn't exist.", err)
	}

	monsterItems, err := cfg.getMonsterItems(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	monsterEquipment, err := cfg.getMonsterEquipment(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	rel, err := cfg.getMonsterRelationships(r, dbMonster)
	if err != nil {
		return Monster{}, err
	}

	species, err := cfg.newNamedAPIResourceFromType("monster-species", string(dbMonster.Species), cfg.t.MonsterSpecies)
	if err != nil {
		return Monster{}, err
	}

	ctbIconType, err := cfg.newNamedAPIResourceFromType("ctb-icon-type", string(dbMonster.CtbIconType), cfg.t.CTBIconType)
	if err != nil {
		return Monster{}, err
	}

	monster := Monster{
		ID:                   dbMonster.ID,
		Name:                 dbMonster.Name,
		Version:              h.NullInt32ToPtr(dbMonster.Version),
		Specification:        h.NullStringToPtr(dbMonster.Specification),
		Notes:                h.NullStringToPtr(dbMonster.Notes),
		Species:              species,
		IsStoryBased:         dbMonster.IsStoryBased,
		IsRepeatable: 		  dbMonster.IsRepeatable,
		CanBeCaptured:        dbMonster.CanBeCaptured,
		AreaConquestLocation: h.ConvertNullMaCreationArea(dbMonster.AreaConquestLocation),
		CTBIconType:          ctbIconType,
		HasOverdrive:         dbMonster.HasOverdrive,
		IsUnderwater:         dbMonster.IsUnderwater,
		IsZombie:             dbMonster.IsZombie,
		Distance:             anyToInt32(dbMonster.Distance),
		Properties:           rel.Properties,
		AutoAbilities:        rel.AutoAbilities,
		AP:                   dbMonster.Ap,
		APOverkill:           dbMonster.ApOverkill,
		OverkillDamage:       dbMonster.OverkillDamage,
		Gil:                  dbMonster.Gil,
		StealGil:             h.NullInt32ToPtr(dbMonster.StealGil),
		RonsoRages:           rel.RonsoRages,
		DoomCountdown:        anyToInt32Ptr(dbMonster.DoomCountdown),
		PoisonRate:           anyToFloat32Ptr(dbMonster.PoisonRate),
		ThreatenChance:       anyToInt32Ptr(dbMonster.ThreatenChance),
		ZanmatoLevel:         anyToInt32(dbMonster.ZanmatoLevel),
		MonsterArenaPrice:    h.NullInt32ToPtr(dbMonster.MonsterArenaPrice),
		SensorText:           h.NullStringToPtr(dbMonster.SensorText),
		ScanText:             h.NullStringToPtr(dbMonster.ScanText),
		Locations: 			  rel.Locations,
		Formations: 		  rel.Formations,
		BaseStats:            rel.BaseStats,
		Items:                h.NilOrPtr(monsterItems),
		Equipment:            h.NilOrPtr(monsterEquipment),
		ElemResists:          rel.ElemResists,
		StatusImmunities:     rel.StatusImmunities,
		StatusResists:        rel.StatusResists,
		AlteredStates:        rel.AlteredStates,
		Abilities:            rel.Abilities,
	}

	monster, err = cfg.applyAlteredState(r, monster)
	if err != nil {
		return Monster{}, err
	}

	monster.BribeChances, err = cfg.getMonsterBribeChances(monster)
	if err != nil {
		return Monster{}, err
	}

	monster.PoisonDamage, err = cfg.getMonsterPoisonDamage(monster)
	if err != nil {
		return Monster{}, err
	}

	monster.AgilityParameters, err = cfg.getMonsterAgilityVals(r, monster)
	if err != nil {
		return Monster{}, err
	}

	monster.BaseStats, err = cfg.applyRonsoStats(r, monster)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}



func (cfg *apiConfig) getMultipleMonsters(r *http.Request, monsterName string) (NamedApiResourceList, error) {
	dbMons, err := cfg.db.GetMonstersByName(r.Context(), monsterName)
	if err != nil {
		return NamedApiResourceList{}, newHTTPError(http.StatusNotFound, "Couldn't get multiple Monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, &mon.Version.Int32, &mon.Specification.String
	})

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return resourceList, nil
}



func (cfg *apiConfig) handleMonstersRetrieve(w http.ResponseWriter, r *http.Request) {
	dbMons, err := cfg.db.GetMonsters(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve monsters", err)
		return
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	resources, err = cfg.getMonstersElemResist(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersStatusResist(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersItem(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersAutoAbility(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersRonsoRage(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersLocation(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersSubLocation(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersArea(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersDistance(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersSpecies(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersCreationArea(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersType(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersStoryBased(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersRepeatable(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersCanBeCaptured(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersHasOverdrive(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersUnderwater(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resources, err = cfg.getMonstersZombie(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}


func (cfg *apiConfig) getMonstersElemResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("elemental-affinities")
	
	if query == "" {
		return inputMons, nil
	}

	eaPairs := strings.Split(query, ",")
	var ids []int32

	for _, pair := range eaPairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. usage: elemental-affinities={element}-{affinity},{element}-{affinity}", nil)
		}

		element, err := parseSingleSegmentResource(parts[0], cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid element: %s. element doesn't exist. use /api/elements to see existing elements.", parts[0]), err)
		}

		affinity, err := parseSingleSegmentResource(parts[1], cfg.l.Affinities)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid affinity: %s. affinity doesn't exist. use /api/affinities to see existing affinities.", parts[1]), err)
		}

		elemResist := seeding.ElementalResist{
			ElementID: element.ID,
			AffinityID: affinity.ID,
		}

		elemResistLookup, err := seeding.GetResource(elemResist, cfg.l.ElementalResists)
		if err != nil {
			return []NamedAPIResource{}, nil
		}

		ids = append(ids, elemResistLookup.ID)
	}

	dbMons, err := cfg.db.GetMonstersByElemResistIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by elemental affinities", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersStatusResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryStatusses := r.URL.Query().Get("status-resists")
	queryResistance := r.URL.Query().Get("resistance")
	defaultResistance := 1
	var resistance int

	if queryStatusses == "" && queryResistance != "" {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. resistance parameter must be paired with status-resists parameter. usage: status-resists={status},{status},...&resistance={int from 1-254}", nil)
	}
	
	if queryStatusses == "" {
		return inputMons, nil
	}
	
	switch queryResistance {
	case "":
		resistance = defaultResistance
	case "immune":
		resistance = 254
	default:
		var err error
		resistance, err = strconv.Atoi(queryResistance)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, "invalid resistance", err)
		}
	}
	
	if resistance > 254 || resistance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resistance. resistance must be a positive integer between 1 or 254, with %d being the default value, if no resistance was provided.", defaultResistance), nil)
	}
	
	statusses := strings.Split(queryStatusses, ",")
	var ids []int32
	
	for _, qStatus := range statusses {
		status, err := parseSingleSegmentResource(qStatus, cfg.l.StatusConditions)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid status condition: %s. status condition doesn't exist. use /api/status-conditions to see existing status conditions.", qStatus), err)
		}

		ids = append(ids, status.ID)
	}

	dbMons, err := cfg.db.GetMonstersByStatusResists(r.Context(), database.GetMonstersByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance: 		resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersItem(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryItem := r.URL.Query().Get("item")
	queryMethod := r.URL.Query().Get("method")

	if queryItem == "" && queryMethod != "" {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. method parameter must be paired with item parameter. usage: item={item},&method={steal/drop/bribe/other}", nil)
	}
	
	if queryItem == "" {
		return inputMons, nil
	}

	item, err := parseSingleSegmentResource(queryItem, cfg.l.MasterItems)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid item: %s. item doesn't exist. use /api/items to see existing items.", queryItem), err)
	}

	var dbMons []database.Monster

	switch queryMethod {
	case "":
		dbMons, err = cfg.db.GetMonstersByItem(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by item", err)
		}
	case "steal":
		dbMons, err = cfg.db.GetMonstersByItemSteal(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by steal item", err)
		}
	case "drop":
		dbMons, err = cfg.db.GetMonstersByItemDrop(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by drop item", err)
		}
	case "bribe":
		dbMons, err = cfg.db.GetMonstersByItemBribe(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by bribe item", err)
		}
	case "other":
		dbMons, err = cfg.db.GetMonstersByItemOther(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by other items", err)
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, "invalid method. allowed methods: steal, drop, bribe, other.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersAutoAbility(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("auto-abilities")
	
	if query == "" {
		return inputMons, nil
	}
	
	abilities := strings.Split(query, ",")
	var ids []int32

	for _, ability := range abilities {
		autoAbility, err := parseSingleSegmentResource(ability, cfg.l.AutoAbilities)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid auto ability: %s. auto ability doesn't exist. use /api/auto-abilities to see existing auto-abilities.", ability), err)
		}

		ids = append(ids, autoAbility.ID)
	}

	dbMons, err := cfg.db.GetMonstersByAutoAbilityIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by auto ability", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersRonsoRage(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	const ronsoRageOffset int = 35
	query := r.URL.Query().Get("ronso-rage")
	var ronsoRageID int32
	
	if query == "" {
		return inputMons, nil
	}

	id, err := strconv.Atoi(query)
	if err == nil {
		id += ronsoRageOffset
		ronsoRageID = int32(id)
	}

	if ronsoRageID == 0 {
		ronsoRage, err := parseNameVersionResource(query, "", cfg.l.Overdrives)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid ronso rage: %s. ronso rage doesn't exist. use /api/ronso-rages to see existing ronso-rages.", query), err)
		}
		ronsoRageID = ronsoRage.ID
	}

	if ronsoRageID > 47 || ronsoRageID <= 35 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid ronso rage id: %d. provided id must be between 1 and 12", ronsoRageID-35), err)
	}

	dbMons, err := cfg.db.GetMonstersByRonsoRageID(r.Context(), ronsoRageID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by ronso rage", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("location")
	
	if query == "" {
		return inputMons, nil
	}

	location, err := parseSingleSegmentResource(query, cfg.l.Locations)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid location: %s. location doesn't exist. use /api/locations to see existing locations", query), err)
	}

	dbMons, err := cfg.db.GetMonstersByLocation(r.Context(), location.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersSubLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("sublocation")
	
	if query == "" {
		return inputMons, nil
	}

	sublocation, err := parseSingleSegmentResource(query, cfg.l.SubLocations)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid sublocation: %s. sublocation doesn't exist. use /api/sublocations to see existing sublocations.", query), err)
	}

	dbMons, err := cfg.db.GetMonstersBySubLocation(r.Context(), sublocation.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by sublocation", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryArea := r.URL.Query().Get("area")
	
	if queryArea == "" {
		return inputMons, nil
	}

	area_id, err := strconv.Atoi(queryArea)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid area id", err)
	}

	dbMons, err := cfg.db.GetMonstersByArea(r.Context(), int32(area_id))
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location", err)
	}

	/* don't have the functionality to parse the area yet
	if err == nil for atoi, I can directly call the query
	dbMons needs to be declared in advance
	else {
		queryLocation := r.URL.Query().Get("location")
		querySubLocation := r.URL.Query().Get("sublocation")
		queryVersion := r.URL.Query().Get("version")
		var version int
		var versionPtr *int32
	
		if queryLocation == "" || querySubLocation == "" {
			return nil, newHTTPError(http.StatusBadRequest, "area either needs a valid id, or be paired with a location, sublocation and a version, like this: location={name or id}&sublocation={name or id}&area={name}&version={id (optional)}", nil)
		}

		locationParsed, err := parseSingleSegmentResource(queryLocation, cfg.l.Locations)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid location: %s. location doesn't exist. use /api/locations to see existing locations.", queryArea), err)
		}

		sublocationParsed, err := parseSingleSegmentResource(querySubLocation, cfg.l.SubLocations)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid sublocation: %s. sublocation doesn't exist. use /api/sublocations to see existing sublocations.", queryArea), err)
		}

		if queryVersion == "" {
			versionPtr = nil
		} else {
			version, err = strconv.Atoi(queryVersion)
			if err != nil {
				return nil, newHTTPError(http.StatusBadRequest, "invalid version. version must be an integer", err)
			}

			versionInt32 := int32(version)
			versionPtr = &versionInt32
		}

		areaParsed, err := parseNameVersionResource()

		dbLocationArea, err := cfg.db.GetLocationAreaByAreaName(r.Context(), database.GetLocationAreaByAreaNameParams{
			ID: locationParsed.ID,
			ID_2: sublocationParsed.ID,
			Name: ,
		})


	}
	*/

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersDistance(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("distance")

	if query == "" {
		return inputMons, nil
	}

	distance, err := strconv.Atoi(query)
	if err != nil || distance > 4 || distance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: distance={int from 1-4}", err)
	}

	dbMons, err := cfg.db.GetMonstersByDistance(r.Context(), distance)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by distance", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersStoryBased(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-story-based")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-story-based={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsStoryBased(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersRepeatable(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-repeatable")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-repeatable={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsRepeatable(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve repeatable monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersCanBeCaptured(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("can-capture")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: can-capture={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByCanBeCaptured(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters that can be captured", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersHasOverdrive(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("has-overdrive")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: has-overdrive={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByHasOverdrive(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters with overdrive gauge", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersUnderwater(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-underwater")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-underwater={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsUnderwater(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve underwater monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersZombie(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-zombie")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-zombie={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsZombie(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve zombie monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersSpecies(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("species")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.MonsterSpecies)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/monster-species to see valid values", query), err)
	}

	species := database.MonsterSpecies(enum.Name)

	dbMons, err := cfg.db.GetMonstersBySpecies(r.Context(), species)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by species", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}



func (cfg *apiConfig) getMonstersCreationArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("creation-area")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.MaCreationArea)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/creation-areas to see valid values", query), err)
	}

	area := h.NullMaCreationArea(&enum.Name)

	dbMons, err := cfg.db.GetMonstersByMaCreationArea(r.Context(), area)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by creation area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersType(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("type")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.CTBIconType)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/ctb-icon-type to see valid values. 'boss' and 'boss-numbered' will both yield all boss monsters.", query), err)
	}

	species := database.CtbIconType(enum.Name)
	var dbMons []database.Monster

	switch species {
	case database.CtbIconTypeMonster:
		dbMons, err = cfg.db.GetMonstersByCTBIconTypeMonster(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	case database.CtbIconTypeSummon:
		dbMons, err = cfg.db.GetMonstersByCTBIconTypeSummon(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	case database.CtbIconTypeBoss, database.CtbIconTypeBossNumbered:
		dbMons, err = cfg.db.GetMonstersByCTBIconTypeBoss(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}