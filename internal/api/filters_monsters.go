package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
	eaPairs := querySplit(query, ",")
	var ids []int32
	elemMap := make(map[int32]bool)

	for _, pair := range eaPairs {
		element, affinity, found := strings.Cut(pair, "=")
		if !found {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, element, queryParam.Usage), nil)
		}

		elementID, err := checkQueryNameID(element, cfg.e.elements.resourceType, queryParam, cfg.l.Elements)
		if err != nil {
			return nil, err
		}
		if elemMap[elementID] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' for parameter '%s'. each element can only be used once.", elementID, queryParam.Name), nil)
		}
		elemMap[elementID] = true

		affinityID, err := checkQueryNameID(affinity, cfg.e.affinities.resourceType, queryParam, cfg.l.Affinities)
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

func getMonstersByAutoAbility(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.monsters
	resourceType := i.resourceType
	queryParam := i.queryLookup["is_forced"]

	query, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetMonsterIDsByAutoAbility)
	}

	dbIDs, err := cfg.db.GetMonsterIDsByAutoAbilityIsForced(r.Context(), database.GetMonsterIDsByAutoAbilityIsForcedParams{
		ID:       id,
		IsForced: query,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by auto-ability id '%d'.", i.resourceType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func getMonstersByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.monsters, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"steal": cfg.db.GetMonsterIDsByItemSteal,
		"drop":  cfg.db.GetMonsterIDsByItemDrop,
		"bribe": cfg.db.GetMonsterIDsByItemBribe,
		"other": cfg.db.GetMonsterIDsByItemOther,
	})
}
