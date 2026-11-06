package api

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMonstersByElemResists(r *http.Request, ctx context.Context, query string, queryParam queryparam.QueryParam) ([]int32, error) {
	ids, err := getElemResistIDs(query, queryParam)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMonsterIDsByElemResistIDs(ctx, ids)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by elemental affinities", err)
	}

	return dbIDs, nil
}

func getElemResistIDs(query string, queryParam queryparam.QueryParam) ([]int32, error) {
	eaPairs, err := queryparam.QueryListSplit(query)
	if err != nil {
		return nil, err
	}
	var ids []int32
	elemMap := make(map[int32]bool)

	for _, pair := range eaPairs {
		elementStr, affinityStr, found := strings.Cut(pair, ":")
		if !found {
			return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'", queryParam.Name, elementStr, queryParam.Usage), nil)
		}

		elementID, err := queryParam.CheckNameID(elementStr, E.Elements.resTypeSingle, lookup.L.Elements)
		if err != nil {
			return nil, err
		}
		if elemMap[elementID] {
			return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of id '%d' for parameter '%s'. each element can only be used once", elementID, queryParam.Name), nil)
		}
		elemMap[elementID] = true

		affinity, err := queryParam.CheckEnum(affinityStr, enums.T.ElementalAffinity)
		if err != nil {
			return nil, err
		}

		elemResist := seedtypes.ElementalResist{
			ElementID: elementID,
			Affinity:  affinity.Name,
		}

		erLookup, err := lookup.GetResource(seedtypes.Key(elemResist), lookup.L.ElementalResists)
		if err != nil {
			return nil, nil
		}

		ids = append(ids, erLookup.ID)
	}

	return ids, nil
}

func getMonstersByStatusResists(r *http.Request, ctx context.Context, ids []int32) ([]int32, error) {
	resistance, err := verifyMonsterResistance(r)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMonsterIDsByStatusResists(ctx, database.GetMonsterIDsByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance:      resistance,
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, "couldn't filter monsters by status conditions", err)
	}

	return dbIDs, nil
}

func verifyMonsterResistance(r *http.Request) (int32, error) {
	queryParam := Q.Monsters[ts.QpnResistance]

	resistance, err := queryParam.ParseInt(r)
	if err != nil {
		return 0, err
	}

	return int32(resistance), nil
}

func getMonstersByAutoAbility(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Monsters

	queryParam := i.queryLookup[ts.QpnIsForced]

	isForced, err := queryParam.ParseBool(r)
	if errs.QueryIsEmpty(err) {
		return config.Cfg.DB.GetMonsterIDsByAutoAbility(ctx, id)
	}

	dbIDs, err := config.Cfg.DB.GetMonsterIDsByAutoAbilityIsForced(ctx, database.GetMonsterIDsByAutoAbilityIsForcedParams{
		AutoAbilityID: id,
		IsForced:      isForced,
	})
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %s by auto-ability id '%d'", i.resTypePlural, id), err)
	}

	return dbIDs, nil
}

func getMonstersByItem(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Monsters

	methods, err := getMonsterItemMethods(r)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMonsterIDsByItem(ctx, database.GetMonsterIDsByItemParams{
		ItemID:  id,
		Methods: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, i.queryLookup[ts.QpnItem], err)
	}

	return dbIDs, nil
}

func getMonsterItemMethods(r *http.Request) ([]string, error) {
	i := E.Monsters
	paramMethods := i.queryLookup[ts.QpnMethods]
	aliasses := map[string][]string{
		"steal": {"steal_common", "steal_rare"},
		"drop":  {"drop_common", "drop_rare", "drop_secondary_common", "drop_secondary_rare"},
	}

	queryMethods, err := paramMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		queryMethods = ts.QVsToStrings(paramMethods.AllowedValues)
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
