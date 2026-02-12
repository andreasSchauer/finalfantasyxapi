package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMonstersByElemResists(cfg *Config, r *http.Request, query string, queryParam QueryType) ([]int32, error) {
	ids, err := getElemResistIDs(cfg, query, queryParam)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMonsterIDsByElemResistIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by elemental affinities.", err)
	}

	return dbIDs, nil
}

func getElemResistIDs(cfg *Config, query string, queryParam QueryType) ([]int32, error) {
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

func getMonstersByStatusResists(cfg *Config, r *http.Request, ids []int32) ([]int32, error) {
	resistance, err := verifyMonsterResistance(cfg, r)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMonsterIDsByStatusResists(r.Context(), database.GetMonsterIDsByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance:      resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions.", err)
	}

	return dbIDs, nil
}

func verifyMonsterResistance(cfg *Config, r *http.Request) (int, error) {
	queryParam := cfg.q.monsters["resistance"]

	resistance, err := parseIntQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return resistance, nil
}

func getMonstersByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.monsters
	resourceType := i.resourceType
	queryParam := i.queryLookup["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []NamedAPIResource
	var err error

	switch query {
	case "":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByItem)
		if err != nil {
			return nil, err
		}
	case "steal":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByItemSteal)
		if err != nil {
			return nil, err
		}
	case "drop":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByItemDrop)
		if err != nil {
			return nil, err
		}
	case "bribe":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByItemBribe)
		if err != nil {
			return nil, err
		}
	case "other":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByItemOther)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value: '%s'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}

func getMonstersByType(cfg *Config, r *http.Request, iconType database.CtbIconType) ([]int32, error) {
	var ids []int32
	var err error

	switch iconType {
	case database.CtbIconTypeBoss, database.CtbIconTypeBossNumbered:
		ids, err = cfg.db.GetMonsterIDsByCTBIconTypeBoss(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters of type 'boss'.", err)
		}
	default:
		ids, err = cfg.db.GetMonsterIDsByCTBIconType(r.Context(), iconType)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve monsters of type '%s'.", iconType), err)
		}
	}

	return ids, nil
}
