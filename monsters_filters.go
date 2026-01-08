package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// can put the elemental affinity id for loop into its own function
func (cfg *Config) getMonstersElemResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["elemental-affinities"]
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return inputMons, nil
	}

	eaPairs := strings.Split(query, ",")
	var ids []int32

	for _, pair := range eaPairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input. usage: %s", queryParam.Usage), nil)
		}

		elementID, err := parseQueryNamedVal(parts[0], cfg.e.elements.resourceType, queryParam, cfg.l.Elements)
		if err != nil {
			return nil, err
		}

		affinityID, err := parseQueryNamedVal(parts[1], cfg.e.affinities.resourceType, queryParam, cfg.l.Affinities)
		if err != nil {
			return nil, err
		}

		elemResist := seeding.ElementalResist{
			ElementID:  elementID,
			AffinityID: affinityID,
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

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersStatusResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParamStatusResists := cfg.q.monsters["status-resists"]
	queryStatusses := r.URL.Query().Get(queryParamStatusResists.Name)

	if queryStatusses == "" {
		return inputMons, nil
	}

	resistance, err := cfg.verifyMonsterResistance(r)
	if err != nil {
		return nil, err
	}

	statusIDs := strings.Split(queryStatusses, ",")
	var ids []int32

	for _, idStr := range statusIDs {
		id, err := parseQueryNamedVal(idStr, cfg.e.statusConditions.resourceType, queryParamStatusResists, cfg.l.StatusConditions)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	dbMons, err := cfg.db.GetMonstersByStatusResists(r.Context(), database.GetMonstersByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance:      resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}


// can generalize the logic within the default statement. distance has this as well, but without a default value
func (cfg *Config) verifyMonsterResistance(r *http.Request) (int, error) {
	queryParam := cfg.q.monsters["resistance"]
	query := r.URL.Query().Get(queryParam.Name)
	minResistance := queryParam.AllowedIntRange[0]
	maxResistance := queryParam.AllowedIntRange[1]
	var resistance int

	switch query {
	case "":
		resistance = minResistance
	case "immune":
		resistance = maxResistance
	default:
		resistance, err := strconv.Atoi(query)
		if err != nil || resistance > maxResistance || resistance < minResistance {
			return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resistance: %s. resistance must be an integer ranging from %d to %d, with %d being the default value, if no resistance was provided.", query, minResistance, maxResistance, minResistance), nil)
		}
	}

	return resistance, nil
}

func (cfg *Config) getMonstersItem(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParamItem := cfg.q.monsters["item"]
	queryParamMethod := cfg.q.monsters["method"]
	queryMethod := r.URL.Query().Get(queryParamMethod.Name)

	id, err := parseIDOnlyQuery(r, queryParamItem, len(cfg.l.Items))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	var dbMons []database.Monster

	switch queryMethod {
	case "":
		dbMons, err = cfg.db.GetMonstersByItem(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by item", err)
		}
	case "steal":
		dbMons, err = cfg.db.GetMonstersByItemSteal(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by steal item", err)
		}
	case "drop":
		dbMons, err = cfg.db.GetMonstersByItemDrop(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by drop item", err)
		}
	case "bribe":
		dbMons, err = cfg.db.GetMonstersByItemBribe(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by bribe item", err)
		}
	case "other":
		dbMons, err = cfg.db.GetMonstersByItemOther(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by other items", err)
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value: %s. allowed values: %s.", queryMethod, strings.Join(queryParamMethod.AllowedValues, ", ")), err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

// generalize multiple idStr inputs to ids []int32 (also used in status-resists)
func (cfg *Config) getMonstersAutoAbility(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["auto-abilities"]
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return inputMons, nil
	}

	abilityIDs := strings.Split(query, ",")
	var ids []int32

	for _, idStr := range abilityIDs {
		id, err := parseQueryIdValStrict(idStr, queryParam, len(cfg.l.AutoAbilities))
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	dbMons, err := cfg.db.GetMonstersByAutoAbilityIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by auto ability", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

// could be a bit less complicated with a ronso rage lookup
// then again, for the query I need to add the offset anyways
// another idea would be to save the save the ronso rage id in the db overdrives table
func (cfg *Config) getMonstersRonsoRage(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["ronso-rage"]
	// don't know how to come to that naturally.
	// maybe overdrive id - ronso rage id of one resource, or enforce it in seeding
	// maybe I won't even need it since I will have both ids in the lookup and can just use the od id for queries
	const ronsoRageOffset int32 = 35
	const ronsoRageCount int32 = 12 // will be replaced by the count of the ronso rage lookup

	id, err := parseIDOnlyQuery(r, queryParam, int(ronsoRageCount))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	ronsoRageID := id + ronsoRageOffset // or just use overdrive id of ronsorage lookup

	dbMons, err := cfg.db.GetMonstersByRonsoRageID(r.Context(), ronsoRageID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by ronso rage", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["location"]
	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.Locations))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetLocationMonsters(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.GetLocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersSubLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["sublocation"]
	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.SubLocations))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetSublocationMonsters(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by sublocation", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.GetSublocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["area"]
	areaID, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.Areas))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetAreaMonsters(r.Context(), areaID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}


func (cfg *Config) getMonstersDistance(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["distance"]
	query := r.URL.Query().Get(queryParam.Name)
	minDistance := queryParam.AllowedIntRange[0]
	maxDistance := queryParam.AllowedIntRange[1]

	if query == "" {
		return inputMons, nil
	}

	// here
	distance, err := strconv.Atoi(query)
	if err != nil || distance > maxDistance || distance < minDistance {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s. distance must be an integer ranging from %d to %d.", query, minDistance, maxDistance), err)
	}

	dbMons, err := cfg.db.GetMonstersByDistance(r.Context(), distance)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by distance", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersStoryBased(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["story-based"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByIsStoryBased(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersRepeatable(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["repeatable"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByIsRepeatable(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve repeatable monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersCanBeCaptured(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["capture"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByCanBeCaptured(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters that can be captured", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersHasOverdrive(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["has-overdrive"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByHasOverdrive(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters with overdrive gauge", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersUnderwater(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["underwater"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByIsUnderwater(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve underwater monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersZombie(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["zombie"]
	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByIsZombie(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve zombie monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersSpecies(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["species"]
	enum, err := parseTypeQuery(r, queryParam, cfg.t.MonsterSpecies)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	species := database.MonsterSpecies(enum.Name)

	dbMons, err := cfg.db.GetMonstersBySpecies(r.Context(), species)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by species", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

// should I make an endpoint for this, since I'm referencing it? I don't need descriptions, or just use a copy/paste one
func (cfg *Config) getMonstersCreationArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["creation-area"]
	enum, err := parseTypeQuery(r, queryParam, cfg.t.MaCreationArea)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	area := h.NullMaCreationArea(&enum.Name)

	dbMons, err := cfg.db.GetMonstersByMaCreationArea(r.Context(), area)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by creation area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersType(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["type"]
	enum, err := parseTypeQuery(r, queryParam, cfg.t.CTBIconType)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	species := database.CtbIconType(enum.Name)
	var dbMons []database.Monster

	switch species {
	case database.CtbIconTypeBoss, database.CtbIconTypeBossNumbered:
		dbMons, err = cfg.db.GetMonstersByCTBIconTypeBoss(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	default:
		dbMons, err = cfg.db.GetMonstersByCTBIconType(r.Context(), species)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}
