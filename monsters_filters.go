package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// all boolean queries can be generalized
// all type based queries can be generalized. monster type is special, as it groups two types together
// if a filter isn't tied to a second parameter, these functions can be generalized
// if more of them appear, I can generalize basic id-list queries (without second parameters) like auto-abilities
func (cfg *Config) getMonstersElemResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["elemental-affinities"]
	query := r.URL.Query().Get(queryParam.Name)

	if query == "" {
		return inputMons, nil
	}

	ids, err := cfg.getElementalAffinityIDs(query, queryParam)
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByElemResistIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by elemental affinities.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getElementalAffinityIDs(query string, queryParam QueryType) ([]int32, error) {
	eaPairs := strings.Split(query, ",")
	var ids []int32
	elemMap := make(map[int32]bool)

	for _, pair := range eaPairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input. usage: '%s'.", queryParam.Usage), nil)
		}

		elementID, err := parseQueryNamedVal(parts[0], cfg.e.elements.resourceType, queryParam, cfg.l.Elements)
		if err != nil {
			return nil, err
		}
		if elemMap[elementID] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' in '%s'. each element can only be used once.", elementID, queryParam.Name), nil)
		}
		elemMap[elementID] = true

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
			return nil, nil
		}

		ids = append(ids, elemResistLookup.ID)
	}

	return ids, nil
}

func (cfg *Config) getMonstersStatusResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["status-resists"]

	ids, err := parseIdListQuery(r, queryParam, len(cfg.l.AutoAbilities))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	resistance, err := cfg.verifyMonsterResistance(r)
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByStatusResists(r.Context(), database.GetMonstersByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance:      resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) verifyMonsterResistance(r *http.Request) (int, error) {
	queryParam := cfg.q.monsters["resistance"]

	resistance, err := parseIntQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return resistance, nil
}

func (cfg *Config) getMonstersItem(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParamItem := cfg.q.monsters["item"]

	id, err := parseIDOnlyQuery(r, queryParamItem, len(cfg.l.Items))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.queryMonstersByItemMethod(r, id)
	if err != nil {
		return nil, err
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) queryMonstersByItemMethod(r *http.Request, id int32) ([]database.Monster, error) {
	queryParam := cfg.q.monsters["method"]
	query := r.URL.Query().Get(queryParam.Name)
	var dbMons []database.Monster
	var err error

	switch query {
	case "":
		dbMons, err = cfg.db.GetMonstersByItem(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by item.", err)
		}
	case "steal":
		dbMons, err = cfg.db.GetMonstersByItemSteal(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by steal-item.", err)
		}
	case "drop":
		dbMons, err = cfg.db.GetMonstersByItemDrop(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by drop-item.", err)
		}
	case "bribe":
		dbMons, err = cfg.db.GetMonstersByItemBribe(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by bribe-item.", err)
		}
	case "other":
		dbMons, err = cfg.db.GetMonstersByItemOther(r.Context(), id)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by other items.", err)
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value: '%s'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return dbMons, nil
}

func (cfg *Config) getMonstersAutoAbility(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["auto-abilities"]

	ids, err := parseIdListQuery(r, queryParam, len(cfg.l.AutoAbilities))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByAutoAbilityIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by auto-ability.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersRonsoRage(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["ronso-rage"]

	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.RonsoRages))
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByRonsoRageID(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by ronso rage.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by sublocation.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by area.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersDistance(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["distance"]

	distance, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	dbMons, err := cfg.db.GetMonstersByDistance(r.Context(), distance)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by distance.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based monsters.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve repeatable monsters.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters that can be captured.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters with overdrive gauge.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve underwater monsters.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve zombie monsters.", err)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by species.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersCreationArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.monsters["creation-area"]

	enum, err := parseTypeQuery(r, queryParam, cfg.t.CreationArea)
	if errors.Is(err, errEmptyQuery) {
		return inputMons, nil
	}
	if err != nil {
		return nil, err
	}

	area := h.NullMaCreationArea(&enum.Name)

	dbMons, err := cfg.db.GetMonstersByMaCreationArea(r.Context(), area)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by creation area.", err)
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
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type.", err)
		}
	default:
		dbMons, err = cfg.db.GetMonstersByCTBIconType(r.Context(), species)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type.", err)
		}
	}

	resources := createNamedAPIResources(cfg, dbMons, cfg.e.monsters.endpoint, func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}
