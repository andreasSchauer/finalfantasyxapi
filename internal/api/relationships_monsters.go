package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getMonsterRelationships(r *http.Request, monster seedtypes.Monster) (Monster, error) {
	var rel Monster
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Monsters.queryAvlRelGetParams(r, monster.ID)
	if err != nil {
		return Monster{}, err
	}

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.queryAvlRelRun(ctx, monster, availabilityParams, getMonsterAreaIDs())
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = E.MonsterFormations.queryAvlRelRun(ctx, monster, availabilityParams, getMonsterMonsterFormationIDs())
		return err
	})

	err = g.Wait()
	if err != nil {
		return Monster{}, err
	}

	return rel, nil
}

func completeMonsterResponse(r *http.Request, mon Monster) (Monster, error) {
	mon, err := applyAlteredStateFromQuery(r, mon, ts.QpnAlteredState)
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyPossessedAeonStats(r, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.BaseStats, err = applyRonsoStats(r, mon, ts.QpnKimahriStats)
	if err != nil {
		return Monster{}, err
	}

	mon.RegenAmounts = calcRegenAmountsMonster(mon)

	mon.ElemResists, err = applyOmnisElements(r, mon, ts.QpnOmnisElements)
	if err != nil {
		return Monster{}, err
	}

	mon.BribeChances, err = getMonsterBribeChances(mon)
	if err != nil {
		return Monster{}, err
	}

	mon.PoisonDamage, err = getMonsterPoisonDamage(mon)
	if err != nil {
		return Monster{}, err
	}

	mon.AgilityParameters, err = getMonsterAgilityParams(r, mon)
	if err != nil {
		return Monster{}, err
	}

	mon.AlteredStates = getAlteredStatesAlterations(mon.AlteredStates)

	return mon, nil
}

func getMonsterElemResists(resists []seedtypes.ElementalResist) []apitypes.ElementalResist {
	elemResists := namesToElemResists(resists)
	elemResistMap := getResourceMap(elemResists)

	for key := range lookup.L.ElementsID {
		_, ok := elemResistMap[key]
		if !ok {
			element := lookup.L.ElementsID[key]
			elemResistMap[element.ID] = newElemResist(element.Name, "neutral")
		}
	}

	return resourceMapToSlice(elemResistMap)
}

func getMonsterPoisonDamage(mon Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hpStat := getBaseStat("hp", mon.BaseStats)

	poisonDamageFloat := float32(hpStat.Value) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}

func getMonsterAgilityParams(r *http.Request, mon Monster) (*apitypes.AgilityParams, error) {
	agilityTier, err := getAgilityTierDB(r, mon.BaseStats)
	if err != nil {
		return nil, err
	}
	if agilityTier.MinAgility == 0 {
		return nil, nil
	}

	tickspeed := agilityTier.TickSpeed

	if appliedStateHaste(mon) {
		tickspeed /= 2
	}

	agilityParams := apitypes.AgilityParams{
		AgilityTier: idToUnnamedAPIResource(E.AgilityTiers, agilityTier.ID),
		TickSpeed:   tickspeed,
		MinICV:      agilityTier.MonsterMinICV,
		MaxICV:      agilityTier.MonsterMaxICV,
	}

	fs := nameToNamedAPIResource(E.AutoAbilities, "first strike", nil)
	if resourcesContain(mon.AutoAbilities, fs) {
		agilityParams.MinICV, agilityParams.MaxICV = getEqualICVs(-1)
	}

	return &agilityParams, nil
}

func appliedStateHaste(mon Monster) bool {
	if mon.AppliedState == nil {
		return false
	}

	if mon.AppliedState.AppliedStatus == nil {
		return false
	}

	return mon.AppliedState.AppliedStatus.StatusCondition.NameIs("haste", nil)
}

// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func getMonsterBribeChances(mon Monster) ([]BribeChance, error) {
	bribe := nameToNamedAPIResource(E.StatusConditions, "bribe", nil)
	if resourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hpStat := getBaseStat("hp", mon.BaseStats)
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
