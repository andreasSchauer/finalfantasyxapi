package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getAeonRelationships(cfg *Config, r *http.Request, aeon seeding.Aeon) (Aeon, error) {
	var rel Aeon
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.CelestialWeapon, err = getResPtrDB(cfg, ctx, cfg.e.celestialWeapons, aeon, cfg.db.GetAeonCelestialWeaponID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.CharacterClasses, err = getResourcesDbItem(cfg, ctx, cfg.e.characterClasses, aeon, cfg.db.GetAeonCharClassIDs)
		return err
	}) 

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = getResourcesDbItem(cfg, ctx, cfg.e.aeonCommands, aeon, cfg.db.GetAeonAeonCommandIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.Overdrives, err = getResourcesDbItem(cfg, ctx, cfg.e.overdrives, aeon, cfg.db.GetAeonOverdriveIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.DefaultPlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, aeon, cfg.db.GetAeonDefaultAbilityIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Aeon{}, err
	}

	return rel, nil
}

func applyAeonStats(cfg *Config, r *http.Request, aeon Aeon) (Aeon, error) {
	var err error

	aeon.BaseStats, err = applyAeonStatsBattles(cfg, r, aeon, qpnBattles)
	if err != nil {
		return Aeon{}, err
	}

	aeon.BaseStats, err = applyYunaStats(cfg, r, aeon, qpnYunaStats)
	if err != nil {
		return Aeon{}, err
	}

	aeon.AgilityParameters, err = getAeonAgilityParams(cfg, r, aeon)
	if err != nil {
		return Aeon{}, err
	}

	return aeon, nil
}

func applyAeonStatsBattles(cfg *Config, r *http.Request, aeon Aeon, queryName QueryParamName) ([]BaseStat, error) {
	queryParam := cfg.q.aeons[queryName]
	battles, err := parseIntQuery(r, queryParam)
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

	seedAeon, _ := seeding.GetResourceByID(aeon.ID, cfg.l.AeonsID)
	newBaseStats := seedAeon.BaseStats.XVals[i].BaseStats
	baseStats := toResAmtType(cfg, cfg.e.stats, newBaseStats, newBaseStat)

	return baseStats, nil
}

func applyYunaStats(cfg *Config, r *http.Request, aeon Aeon, queryName QueryParamName) ([]BaseStat, error) {
	allowedStatIDs := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	baseStats := aeon.BaseStats
	queryParam := cfg.q.aeons[queryName]

	yuna := cfg.l.Characters["yuna"]
	yunaBS := toResAmtType(cfg, cfg.e.stats, yuna.BaseStats, newBaseStat)

	yunaStatMapInt, err := parseStatQuery(cfg, r, queryParam, yunaBS, allowedStatIDs)
	if queryIsEmpty(err) {
		return baseStats, nil
	}
	if err != nil {
		return nil, err
	}

	yunaStatMapInt["hp"] = min(yunaStatMapInt["hp"], 9999)
	yunaStatMapInt["mp"] = min(yunaStatMapInt["mp"], 999)

	aeonStats := calcAeonStats(cfg, aeon, yunaStatMapInt)

	return aeonStats, nil
}

func calcAeonStats(cfg *Config, aeon Aeon, yuna map[string]int32) []BaseStat {
	baseStats := aeon.BaseStats
	aVals, bVals := getAeonStatTables(cfg, aeon)
	yunaFloat := resAmtMapToFloat(yuna)
	yParameter := yParamCalc(yuna)

	for i, baseStat := range baseStats {
		yStat := yStatCalc(baseStat.GetName(), yParameter, aVals, bVals, yunaFloat)
		baseStat.Value = max(baseStat.Value, yStat)
		baseStats[i] = baseStat
	}

	return baseStats
}

func getAeonStatTables(cfg *Config, aeon Aeon) (map[string]float64, map[string]float64) {
	aeonLookup, _ := seeding.GetResourceByID(aeon.ID, cfg.l.AeonsID)

	aValsSlice := toResAmtType(cfg, cfg.e.stats, aeonLookup.BaseStats.AVals, newBaseStat)
	aValsInt := getResAmtTypeMap(aValsSlice)
	aVals := resAmtMapToFloat(aValsInt)

	bValsSlice := toResAmtType(cfg, cfg.e.stats, aeonLookup.BaseStats.BVals, newBaseStat)
	bValsInt := getResAmtTypeMap(bValsSlice)
	bVals := resAmtMapToFloat(bValsInt)

	return aVals, bVals
}

func resAmtMapToFloat(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64)

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}

	return floatMap
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

func getAeonAgilityParams(cfg *Config, r *http.Request, aeon Aeon) (AgilityParams, error) {
	agilityTier, err := getAgilityTier(cfg, r, aeon.BaseStats)
	if err != nil {
		return AgilityParams{}, err
	}
	agilityStat := getBaseStat(cfg, "agility", aeon.BaseStats)
	agility := agilityStat.Value

	var minICV *int32

	for _, subTier := range agilityTier.CharacterMinICVs {
		if agility >= subTier.MinAgility && agility <= subTier.MaxAgility {
			minICV = subTier.CharacterMinICV
			break
		}
	}

	agilityParams := AgilityParams{
		AgilityTier: idToUnnamedAPIResource(cfg, cfg.e.agilityTiers, agilityTier.ID),
		TickSpeed:   agilityTier.TickSpeed,
		MinICV:      minICV,
		MaxICV:      agilityTier.CharacterMaxICV,
	}

	return agilityParams, nil
}
