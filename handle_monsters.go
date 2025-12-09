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

		element, err := seeding.GetResource(parts[0], cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid element: %s. element doesn't exist", parts[0]), err)
		}

		affinity, err := seeding.GetResource(parts[1], cfg.l.Affinities)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid affinity: %s. affinity doesn't exist", parts[1]), err)
		}

		elemResist := seeding.ElementalResist{
			Element: element.Name,
			Affinity: affinity.Name,
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

	statusses := strings.Split(queryStatusses, ",")

	if queryResistance == "" {
		resistance = defaultResistance
	} else {
		var err error
		resistance, err = strconv.Atoi(queryResistance)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, "invalid resistance", err)
		}
	}
	
	if resistance > 254 || resistance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resistance. resistance must be a positive integer between 1 or 254, with %d being the default value, if no resistance was provided.", defaultResistance), nil)
	}

	var ids []int32

	for _, qStatus := range statusses {
		status, err := seeding.GetResource(qStatus, cfg.l.StatusConditions)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid status condition: %s. status condition doesn't exist", qStatus), err)
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