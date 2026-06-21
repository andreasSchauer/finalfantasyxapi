package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMonstersByElemResists(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
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

func getElemResistIDs(cfg *Config, query string, queryParam QueryParam) ([]int32, error) {
	eaPairs, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}
	var ids []int32
	elemMap := make(map[int32]bool)

	for _, pair := range eaPairs {
		elementStr, affinityStr, found := strings.Cut(pair, "=")
		if !found {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, elementStr, queryParam.Usage), nil)
		}

		elementID, err := checkQueryNameID(elementStr, cfg.e.elements.resourceType, queryParam, cfg.l.Elements)
		if err != nil {
			return nil, err
		}
		if elemMap[elementID] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' for parameter '%s'. each element can only be used once.", elementID, queryParam.Name), nil)
		}
		elemMap[elementID] = true

		affinity, err := checkQueryEnum(affinityStr, cfg.e.elementalAffinity.endpoint, queryParam, cfg.t.ElementalAffinity)
		if err != nil {
			return nil, err
		}

		id, err := cfg.l.GetHashID(seeding.ElementalResist{
			ElementID: elementID,
			Affinity:  affinity.Name,
		})
		if err != nil {
			return nil, nil
		}

		ids = append(ids, id)
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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't filter monsters by status conditions.", err)
	}

	return dbIDs, nil
}

func verifyMonsterResistance(cfg *Config, r *http.Request) (int32, error) {
	queryParam := cfg.q.monsters["resistance"]

	resistance, err := parseIntQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	return int32(resistance), nil
}

func getMonstersByAutoAbility(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.monsters

	queryParam := i.queryLookup["is_forced"]

	isForced, err := parseBooleanQuery(r, queryParam)
	if queryIsEmpty(err) {
		return cfg.db.GetMonsterIDsByAutoAbility(r.Context(), id)
	}

	dbIDs, err := cfg.db.GetMonsterIDsByAutoAbilityIsForced(r.Context(), database.GetMonsterIDsByAutoAbilityIsForcedParams{
		AutoAbilityID: id,
		IsForced:      isForced,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by auto-ability id '%d'.", i.resourceType, id), err)
	}

	return dbIDs, nil
}

func getMonstersByItem(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.monsters
	
	methods, err := getMonsterItemMethods(cfg, r)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMonsterIDsByItem(r.Context(), database.GetMonsterIDsByItemParams{
		ItemID:  id,
		Methods: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["item"], err)
	}

	return dbIDs, nil
}

func getMonsterItemMethods(cfg *Config, r *http.Request) ([]string, error) {
	i := cfg.e.monsters
	paramMethods := i.queryLookup["methods"]
	aliasses := map[string][]string{
		"steal": {"steal_common", "steal_rare"},
		"drop": {"drop_common", "drop_rare", "drop_secondary_common", "drop_secondary_rare"},
	}

	queryMethods, err := parseValueListQuery(cfg, r, paramMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		queryMethods = paramMethods.AllowedValues
	}

	methods := []string{}

	for _, method := range queryMethods {
		vals, ok := aliasses[method]
		if ok {
			methods = slices.Concat(methods, vals)
			continue
		}

		methods = append(methods, method)
	}

	return methods, nil
}