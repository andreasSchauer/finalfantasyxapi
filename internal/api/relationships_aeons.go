package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getAeonRelationships(server *Server, r *http.Request, aeon seedtypes.Aeon) (Aeon, error) {
	var rel Aeon
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = E.CelestialWeapons.dbGetRelationPtr(server, ctx, aeon, server.Cfg.DB.GetAeonCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = E.CharacterClasses.dbGetRelations(server, ctx, aeon, server.Cfg.DB.GetAeonCharClassIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = E.AeonCommands.dbGetRelations(server, ctx, aeon, server.Cfg.DB.GetAeonAeonCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = E.Overdrives.dbGetRelations(server, ctx, aeon, server.Cfg.DB.GetAeonOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = E.PlayerAbilities.dbGetRelations(server, ctx, aeon, server.Cfg.DB.GetAeonDefaultAbilityIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Aeon{}, err
	}

	return rel, nil
}

func getAeonBaseStats(server *Server, aeon seedtypes.Aeon) []BaseStat {
	baseStats := aeon.BaseStats.XVals[0].BaseStats
	return E.Stats.toResAmtType(server, baseStats, newBaseStat)
}

func applyAeonStats(server *Server, r *http.Request, aeon Aeon, allowedStatIDs []int32) (Aeon, error) {
	var err error

	aeon.BaseStats, err = applyAeonStatsBattles(server, r, aeon, ts.QpnBattles)
	if err != nil {
		return Aeon{}, err
	}

	aeon.BaseStats, err = applyYunaStats(server, r, aeon, ts.QpnYunaStats, allowedStatIDs)
	if err != nil {
		return Aeon{}, err
	}

	return aeon, nil
}

func applyAeonStatsBattles(server *Server, r *http.Request, aeon Aeon, queryName ts.QueryParamName) ([]BaseStat, error) {
	queryParam := Q.aeons[queryName]
	battles, err := queryParam.parseInt(r)
	if queryIsEmpty(err) {
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
	baseStats := E.Stats.toResAmtType(server, newBaseStats, newBaseStat)

	return baseStats, nil
}

func applyYunaStats(server *Server, r *http.Request, aeon Aeon, queryName ts.QueryParamName, allowedStatIDs []int32) ([]BaseStat, error) {
	queryParam := Q.aeons[queryName]

	yuna := lookup.L.Characters["yuna"]
	yunaBS := E.Stats.toResAmtType(server, yuna.BaseStats, newBaseStat)

	yunaStatCalcMap, err := getStatCalcMap(r, queryParam, yunaBS, allowedStatIDs)
	if queryIsEmpty(err) {
		return aeon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	yunaStatCalcMap["hp"] = min(yunaStatCalcMap["hp"], 9999)
	yunaStatCalcMap["mp"] = min(yunaStatCalcMap["mp"], 999)

	aeonStats := calcAeonStats(server, aeon, yunaStatCalcMap)

	return aeonStats, nil
}

func calcAeonStats(server *Server, aeon Aeon, yuna map[string]int32) []BaseStat {
	baseStats := aeon.BaseStats
	aVals, bVals := getAeonStatTables(server, aeon)
	yunaFloat := resAmtMapToFloat(yuna)
	yParameter := yParamCalc(yuna)

	for i, baseStat := range baseStats {
		yStat := yStatCalc(baseStat.GetName(), yParameter, aVals, bVals, yunaFloat)
		baseStat.Value = max(baseStat.Value, yStat)
		baseStats[i] = baseStat
	}

	return baseStats
}

func getAeonStatTables(server *Server, aeon Aeon) (map[string]float64, map[string]float64) {
	aeonLookup, _ := lookup.GetResourceByID(aeon.ID, lookup.L.AeonsID)

	aValsSlice := E.Stats.toResAmtType(server, aeonLookup.BaseStats.AVals, newBaseStat)
	aValsInt := GetResAmtTypeMap(aValsSlice)
	aVals := resAmtMapToFloat(aValsInt)

	bValsSlice := E.Stats.toResAmtType(server, aeonLookup.BaseStats.BVals, newBaseStat)
	bValsInt := GetResAmtTypeMap(bValsSlice)
	bVals := resAmtMapToFloat(bValsInt)

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

func getUnitAgilityParams(server *Server, r *http.Request, baseStats []BaseStat) (AgilityParams, error) {
	agilityTier, err := getAgilityTierDB(server, r, baseStats)
	if err != nil {
		return AgilityParams{}, err
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

	agilityParams := AgilityParams{
		AgilityTier: idToUnnamedAPIResource(server, E.AgilityTiers, agilityTier.ID),
		TickSpeed:   agilityTier.TickSpeed,
		MinICV:      minICV,
		MaxICV:      agilityTier.CharacterMaxICV,
	}

	return agilityParams, nil
}
