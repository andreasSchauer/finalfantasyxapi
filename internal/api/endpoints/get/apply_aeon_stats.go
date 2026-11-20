package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func applyAeonStats(r *http.Request, aeon *apitypes.Aeon, allowedStatIDs []int32) error {
	err := applyAeonStatsBattles(r, aeon, ts.QpnBattles)
	if err != nil {
		return err
	}

	err = applyYunaStats(r, aeon, ts.QpnYunaStats, allowedStatIDs)
	if err != nil {
		return err
	}

	return nil
}

func applyAeonStatsBattles(r *http.Request, aeon *apitypes.Aeon, queryName ts.QueryParamName) error {
	queryParam := apilookups.Q.Aeons[queryName]
	battles, err := queryParam.ParseInt(r)
	if errs.QueryIsEmpty(err) {
		return nil
	}
	if err != nil {
		return err
	}

	i := battles/30 - 1

	if battles < 60 {
		i = 0
	}

	seedAeon, _ := lookup.GetResourceByID(aeon.ID, lookup.L.AeonsID)
	newBaseStats := seedAeon.BaseStats.XVals[i].BaseStats
	aeon.BaseStats = apilookups.E.Stats.ToResAmtType(newBaseStats, th.NewBaseStat)

	return nil
}

func applyYunaStats(r *http.Request, aeon *apitypes.Aeon, queryName ts.QueryParamName, allowedStatIDs []int32) error {
	queryParam := apilookups.Q.Aeons[queryName]

	yuna := lookup.L.Characters["yuna"]
	yunaBS := apilookups.E.Stats.ToResAmtType(yuna.BaseStats, th.NewBaseStat)

	yunaStatCalcMap, err := getStatCalcMap(r, queryParam, yunaBS, allowedStatIDs)
	if errs.QueryIsEmpty(err) {
		return nil
	}
	if err != nil {
		return err
	}

	yunaStatCalcMap["hp"] = min(yunaStatCalcMap["hp"], 9999)
	yunaStatCalcMap["mp"] = min(yunaStatCalcMap["mp"], 999)

	calcAeonStats(aeon, yunaStatCalcMap)

	return nil
}

func calcAeonStats(aeon *apitypes.Aeon, yuna map[string]int32) {
	aVals, bVals := getAeonStatTables(aeon)
	yunaFloat := apiresource.ResAmtMapToFloat(yuna)
	yParameter := yParamCalc(yuna)

	for i := range aeon.BaseStats {
		baseStat := &aeon.BaseStats[i]
		yStat := yStatCalc(baseStat.GetName(), yParameter, aVals, bVals, yunaFloat)
		baseStat.Value = max(baseStat.Value, yStat)
	}
}

func getAeonStatTables(aeon *apitypes.Aeon) (map[string]float64, map[string]float64) {
	aeonLookup, _ := lookup.GetResourceByID(aeon.ID, lookup.L.AeonsID)

	aValsSlice := apilookups.E.Stats.ToResAmtType(aeonLookup.BaseStats.AVals, th.NewBaseStat)
	aValsInt := apiresource.GetResAmtTypeMap(aValsSlice)
	aVals := apiresource.ResAmtMapToFloat(aValsInt)

	bValsSlice := apilookups.E.Stats.ToResAmtType(aeonLookup.BaseStats.BVals, th.NewBaseStat)
	bValsInt := apiresource.GetResAmtTypeMap(bValsSlice)
	bVals := apiresource.ResAmtMapToFloat(bValsInt)

	return aVals, bVals
}

func yParamCalc(yuna map[string]int32) float64 {
	return float64(yuna["hp"]/100 + yuna["mp"]/10 + yuna["strength"] + yuna["defense"] + yuna["magic"] + yuna["magic defense"] + yuna["agility"] + yuna["evasion"] + yuna["accuracy"])
}

func yStatCalc(stat string, yParameter float64, aVals, bVals, yuna map[string]float64) int32 {
	var yFloat float64
	var yStat int32

	switch stat {
	case "hp":
		yFloat = aVals[stat]*yParameter + bVals[stat]*(yuna[stat]/100)
		yStat = int32(yFloat)
		return min(yStat, 99999)

	case "mp":
		yFloat = aVals[stat]*(yParameter/10) + bVals[stat]*(yuna[stat]/100)
		yStat = int32(yFloat)
		return min(yStat, 9999)

	case "luck":
		return int32(yuna[stat])

	default:
		yFloat := yParameter/aVals[stat] + bVals[stat]*(yuna[stat]/10)
		yStat = int32(yFloat)
		return min(yStat, 255)
	}
}
