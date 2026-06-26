package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getMonsterRelationships(cfg *Config, r *http.Request, monster seeding.Monster) (Monster, error) {
	var rel Monster
	g, ctx := errgroup.WithContext(r.Context())
	
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.monsters, monster.ID)
	if err != nil {
		return Monster{}, err
	}

	g.Go(func() error {
		var err error
		rel.Areas, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.areas, monster, availabilityParams, getMonsterAreaIDs(cfg))
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.Formations, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsterFormations, monster, availabilityParams, getMonsterMonsterFormationIDs(cfg))
		return err
	})
	
	err = g.Wait()
	if err != nil {
		return Monster{}, err
	}

	return rel, nil
}

func completeMonsterResponse(cfg *Config, r *http.Request, mon Monster) (Monster, error) {
	mon, err := applyAlteredState(cfg, r, mon, qpnAlteredState)
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyAeonStatsMonsters(cfg, r, mon, qpnAeonStats)
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyRonsoStats(cfg, r, mon, qpnKimahriStats)
	if err != nil {
		return Monster{}, err
	}

	mon.ElemResists, err = applyOmnisElements(cfg, r, mon, qpnOmnisElements)
	if err != nil {
		return Monster{}, err
	}

	mon.BribeChances, err = getMonsterBribeChances(cfg, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.PoisonDamage, err = getMonsterPoisonDamage(cfg, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.AgilityParameters, err = getMonsterAgilityParams(cfg, r, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.Stats = createStats(mon.BaseStats)
	mon.AlteredStates = getAlteredStatesAlterations(mon.AlteredStates)

	return mon, nil
}

func getMonsterElemResists(cfg *Config, resists []seeding.ElementalResist) []ElementalResist {
	elemResists := namesToElemResists(cfg, resists)
	elemResistMap := getResourceMap(elemResists)

	for key := range cfg.l.ElementsID {
		_, ok := elemResistMap[key]
		if !ok {
			element := cfg.l.ElementsID[key]
			elemResistMap[element.ID] = newElemResist(cfg, element.Name, "neutral")
		}
	}

	return resourceMapToSlice(elemResistMap)
}

func getMonsterPoisonDamage(cfg *Config, mon Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hpStat := getBaseStat(cfg, "hp", mon.BaseStats)

	poisonDamageFloat := float32(hpStat.Value) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}

func getMonsterAgilityParams(cfg *Config, r *http.Request, mon Monster) (*AgilityParams, error) {
	agilityTier, err := getAgilityTier(cfg, r, mon.BaseStats)
	if err != nil {
		return nil, err
	}
	if agilityTier.MinAgility == 0 {
		return nil, nil
	}

	agilityParams := AgilityParams{
		AgilityTier: idToUnnamedAPIResource(cfg, cfg.e.agilityTiers, agilityTier.ID),
		TickSpeed:   agilityTier.TickSpeed,
		MinICV:      agilityTier.MonsterMinICV,
		MaxICV:      agilityTier.MonsterMaxICV,
	}

	fs := nameToNamedAPIResource(cfg, cfg.e.autoAbilities, "first strike", nil)
	if resourcesContain(mon.AutoAbilities, fs) {
		var fsICV int32 = -1
		agilityParams.MinICV = &fsICV
		agilityParams.MaxICV = &fsICV
	}

	return &agilityParams, nil
}

// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func getMonsterBribeChances(cfg *Config, mon Monster) ([]BribeChance, error) {
	bribe := nameToNamedAPIResource(cfg, cfg.e.statusConditions, "bribe", nil)
	if resourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hpStat := getBaseStat(cfg, "hp", mon.BaseStats)
	hp := hpStat.Value

	bribeChances := []BribeChance{}
	var multiplier int32 = 10
	var chance int32 = 25

	for multiplier <= 25 {
		bribeChance := BribeChance{
			Gil:    hp * multiplier,
			Chance: chance,
		}
		bribeChances = append(bribeChances, bribeChance)
		multiplier += 5
		chance += 25
	}

	return bribeChances, nil
}
