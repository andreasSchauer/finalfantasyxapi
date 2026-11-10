package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetMonster(r *http.Request, i handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named], id int32) (apitypes.Monster, error) {
	monster, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Monster{}, err
	}

	rel, err := getMonsterRelationships(r, monster)
	if err != nil {
		return apitypes.Monster{}, err
	}

	response := apitypes.Monster{
		ID:                   monster.ID,
		Name:                 monster.Name,
		Version:              monster.Version,
		Specification:        monster.Specification,
		Notes:                monster.Notes,
		Species:              monster.Species,
		Availability:         monster.Availability,
		IsRepeatable:         monster.IsRepeatable,
		CanBeCaptured:        monster.CanBeCaptured,
		AreaConquestLocation: monster.AreaConquestLocation,
		Category:             monster.Category,
		CTBIconType:          monster.CTBIconType,
		HasOverdrive:         monster.HasOverdrive,
		IsUnderwater:         monster.IsUnderwater,
		IsZombie:             monster.IsZombie,
		Distance:             monster.Distance,
		Properties:           handlerinput.NamesToNamedAPIResources(apilookups.E.Properties, monster.Properties),
		AutoAbilities:        handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           handlerinput.NamesToNamedAPIResources(apilookups.E.RonsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            apilookups.E.Stats.ToResAmtType(monster.BaseStats, th.NewBaseStat),
		Items:                ah.ConvertObjPtr(monster.Items, apiconv.ConvertMonsterItems),
		Equipment:            ah.ConvertObjPtr(monster.Equipment, apiconv.ConvertMonsterEquipment),
		ElemResists:          rel.ElemResists,
		StatusImmunities:     handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, monster.StatusImmunities),
		StatusResists:        apilookups.E.StatusConditions.ToResAmtType(monster.StatusResists, th.NewStatusResist),
		Abilities:            ah.ConvertObjSlice(monster.Abilities, apiconv.ConvertMonsterAbility),
		AlteredStates:        rel.AlteredStates,
	}

	return CompleteMonsterResponse(r, response)
}

func getMonsterRelationships(r *http.Request, monster seedtypes.Monster) (apitypes.Monster, error) {
	rel := apitypes.Monster{
		ElemResists:   GetMonsterElemResists(monster.ElemResists),
		AlteredStates: GetMonsterAlteredStates(r, monster),
	}

	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Monsters.QueryAvlRelGetParams(r, monster.ID)
	if err != nil {
		return apitypes.Monster{}, err
	}

	g.Go(func() error {
		var err error
		rel.Areas, err = apilookups.E.Areas.QueryAvlRelRun(ctx, monster, availabilityParams, getMonsterAreaIDs())
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, monster, availabilityParams, getMonsterMonsterFormationIDs())
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.Monster{}, err
	}

	return rel, nil
}

func CompleteMonsterResponse(r *http.Request, mon apitypes.Monster) (apitypes.Monster, error) {
	mon, err := applyAlteredStateFromQuery(r, mon, ts.QpnAlteredState)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.BaseStats, err = applyPossessedAeonStats(r, mon)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.BaseStats, err = applyRonsoStats(r, mon, ts.QpnKimahriStats)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.RegenAmounts = calcRegenAmountsMonster(mon)

	mon.ElemResists, err = applyOmnisElements(r, mon, ts.QpnOmnisElements)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.BribeChances, err = getMonsterBribeChances(mon)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.PoisonDamage, err = getMonsterPoisonDamage(mon)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.AgilityParameters, err = getMonsterAgilityParams(r, mon)
	if err != nil {
		return apitypes.Monster{}, err
	}

	mon.AlteredStates = getAlteredStatesAlterations(mon.AlteredStates)

	return mon, nil
}

func GetMonsterElemResists(resists []seedtypes.ElementalResist) []apitypes.ElementalResist {
	elemResists := th.NamesToElemResists(resists)
	elemResistMap := handlerinput.GetResourceMap(elemResists)

	for key := range lookup.L.ElementsID {
		_, ok := elemResistMap[key]
		if !ok {
			element := lookup.L.ElementsID[key]
			elemResistMap[element.ID] = th.NewElemResist(element.Name, "neutral")
		}
	}

	return handlerinput.ResourceMapToSlice(elemResistMap)
}

func getMonsterPoisonDamage(mon apitypes.Monster) (*int32, error) {
	if mon.PoisonRate == nil {
		return nil, nil
	}

	hpStat := th.GetBaseStat("hp", mon.BaseStats)

	poisonDamageFloat := float32(hpStat.Value) * *mon.PoisonRate
	poisonDamage := int32(poisonDamageFloat)

	return &poisonDamage, nil
}

func getMonsterAgilityParams(r *http.Request, mon apitypes.Monster) (*apitypes.AgilityParams, error) {
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
		AgilityTier: handlerinput.IDToUnnamedAPIResource(apilookups.E.AgilityTiers, agilityTier.ID),
		TickSpeed:   tickspeed,
		MinICV:      agilityTier.MonsterMinICV,
		MaxICV:      agilityTier.MonsterMaxICV,
	}

	fs := handlerinput.NameToNamedAPIResource(apilookups.E.AutoAbilities, "first strike", nil)
	if handlerinput.ResourcesContain(mon.AutoAbilities, fs) {
		agilityParams.MinICV, agilityParams.MaxICV = th.GetEqualICVs(-1)
	}

	return &agilityParams, nil
}

func appliedStateHaste(mon apitypes.Monster) bool {
	if mon.AppliedState == nil {
		return false
	}

	if mon.AppliedState.AppliedStatus == nil {
		return false
	}

	return mon.AppliedState.AppliedStatus.StatusCondition.NameIs("haste", nil)
}

// HP x10 = 25%, HP x15 = 50%, HP x20 = 75%, HP x25 = 100%
func getMonsterBribeChances(mon apitypes.Monster) ([]apitypes.BribeChance, error) {
	bribe := handlerinput.NameToNamedAPIResource(apilookups.E.StatusConditions, "bribe", nil)
	if handlerinput.ResourcesContain(mon.StatusImmunities, bribe) || mon.Items == nil || mon.Items.Bribe == nil {
		return nil, nil
	}

	hpStat := th.GetBaseStat("hp", mon.BaseStats)
	hp := hpStat.Value

	bribeChances := []apitypes.BribeChance{}
	var multiplier int32 = 10
	var chance int32 = 25

	for multiplier <= 25 {
		bribeChance := apitypes.BribeChance{
			Gil:    hp * multiplier,
			Chance: chance,
		}
		bribeChances = append(bribeChances, bribeChance)
		multiplier += 5
		chance += 25
	}

	return bribeChances, nil
}
