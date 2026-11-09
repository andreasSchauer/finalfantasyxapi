package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getAeonRelationships(r *http.Request, aeon seedtypes.Aeon) (apitypes.Aeon, error) {
	var rel apitypes.Aeon
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = apilookups.E.CelestialWeapons.DBGetRelationPtr(ctx, aeon, config.Cfg.DB.GetAeonCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = apilookups.E.CharacterClasses.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = apilookups.E.AeonCommands.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonAeonCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, aeon, config.Cfg.DB.GetAeonDefaultAbilityIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return rel, nil
}

func getAeonBaseStats(aeon seedtypes.Aeon) []apitypes.BaseStat {
	baseStats := aeon.BaseStats.XVals[0].BaseStats
	return apilookups.E.Stats.ToResAmtType(baseStats, newBaseStat)
}

func applyAeonStats(r *http.Request, aeon apitypes.Aeon, allowedStatIDs []int32) (apitypes.Aeon, error) {
	var err error

	aeon.BaseStats, err = applyAeonStatsBattles(r, aeon, ts.QpnBattles)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	aeon.BaseStats, err = applyYunaStats(r, aeon, ts.QpnYunaStats, allowedStatIDs)
	if err != nil {
		return apitypes.Aeon{}, err
	}

	return aeon, nil
}

func applyAeonStatsBattles(r *http.Request, aeon apitypes.Aeon, queryName ts.QueryParamName) ([]apitypes.BaseStat, error) {
	queryParam := apilookups.Q.Aeons[queryName]
	battles, err := queryParam.ParseInt(r)
	if errs.QueryIsEmpty(err) {
		return aeon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	i := battles/30 - 1

	if battles < 60 {
		i = 0
	}

	seedAeon, _ := lookup.GetResourceByID(aeon.ID, lookup.L.AeonsID)
	newBaseStats := seedAeon.BaseStats.XVals[i].BaseStats
	baseStats := apilookups.E.Stats.ToResAmtType(newBaseStats, newBaseStat)

	return baseStats, nil
}

func applyYunaStats(r *http.Request, aeon apitypes.Aeon, queryName ts.QueryParamName, allowedStatIDs []int32) ([]apitypes.BaseStat, error) {
	queryParam := apilookups.Q.Aeons[queryName]

	yuna := lookup.L.Characters["yuna"]
	yunaBS := apilookups.E.Stats.ToResAmtType(yuna.BaseStats, newBaseStat)

	yunaStatCalcMap, err := getStatCalcMap(r, queryParam, yunaBS, allowedStatIDs)
	if errs.QueryIsEmpty(err) {
		return aeon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	yunaStatCalcMap["hp"] = min(yunaStatCalcMap["hp"], 9999)
	yunaStatCalcMap["mp"] = min(yunaStatCalcMap["mp"], 999)

	aeonStats := calcAeonStats(aeon, yunaStatCalcMap)

	return aeonStats, nil
}

func calcAeonStats(aeon apitypes.Aeon, yuna map[string]int32) []apitypes.BaseStat {
	baseStats := aeon.BaseStats
	aVals, bVals := getAeonStatTables(aeon)
	yunaFloat := handlerinput.ResAmtMapToFloat(yuna)
	yParameter := yParamCalc(yuna)

	for i, baseStat := range baseStats {
		yStat := yStatCalc(baseStat.GetName(), yParameter, aVals, bVals, yunaFloat)
		baseStat.Value = max(baseStat.Value, yStat)
		baseStats[i] = baseStat
	}

	return baseStats
}

func getAeonStatTables(aeon apitypes.Aeon) (map[string]float64, map[string]float64) {
	aeonLookup, _ := lookup.GetResourceByID(aeon.ID, lookup.L.AeonsID)

	aValsSlice := apilookups.E.Stats.ToResAmtType(aeonLookup.BaseStats.AVals, newBaseStat)
	aValsInt := handlerinput.GetResAmtTypeMap(aValsSlice)
	aVals := handlerinput.ResAmtMapToFloat(aValsInt)

	bValsSlice := apilookups.E.Stats.ToResAmtType(aeonLookup.BaseStats.BVals, newBaseStat)
	bValsInt := handlerinput.GetResAmtTypeMap(bValsSlice)
	bVals := handlerinput.ResAmtMapToFloat(bValsInt)

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

func getUnitAgilityParams(r *http.Request, baseStats []apitypes.BaseStat) (apitypes.AgilityParams, error) {
	agilityTier, err := getAgilityTierDB(r, baseStats)
	if err != nil {
		return apitypes.AgilityParams{}, err
	}
	agilityStat := getBaseStat("agility", baseStats)
	agility := agilityStat.Value

	var minICV *int32

	for _, subTier := range agilityTier.CharacterMinICVs {
		if agility >= subTier.MinAgility && agility <= subTier.MaxAgility {
			minICV = subTier.CharacterMinICV
			break
		}
	}

	agilityParams := apitypes.AgilityParams{
		AgilityTier: handlerinput.IDToUnnamedAPIResource(apilookups.E.AgilityTiers, agilityTier.ID),
		TickSpeed:   agilityTier.TickSpeed,
		MinICV:      minICV,
		MaxICV:      agilityTier.CharacterMaxICV,
	}

	return agilityParams, nil
}
