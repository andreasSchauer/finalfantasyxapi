package turnorder

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func fetchFormationMons(params servicetypes.TurnOrderParams) servicetypes.TurnOrderParams {
	if params.FormationID == nil {
		return params
	}

	formation, _ := lookup.GetResourceByID(*params.FormationID, lookup.L.MonsterFormationsID)

	if formation.FormationData.IsForcedAmbush {
		params.BattleStart = string(database.BattleStartAmbush)
	}

	for _, monAmt := range formation.Monsters {
		monID := monAmt.MonsterID

		mon := servicetypes.TurnOrderMon{
			MonsterID: &monID,
		}

		params.Mons = append(params.Mons, mon)
	}

	return params
}

func assembleMonTurnOrder(monID int32, altState *int32) (apitypes.Monster, *seedtypes.AgilityTier) {
	penanceArmIDs := []int32{306, 307}
	firstTurnAglIDs := []int32{167, 168, 215, 306, 307}
	var firstTurnAglTier *seedtypes.AgilityTier

	if slices.Contains(penanceArmIDs, monID) {
		altState = h.GetInt32Ptr(2)
	}

	monster, _ := sh.QuickAssembleMon(monID, altState)

	if slices.Contains(firstTurnAglIDs, monID) {
		firstTurnAglTier = getFirstTurnAglTier(monID)
	}

	return monster, firstTurnAglTier
}

func getFirstTurnAglTier(monID int32) *seedtypes.AgilityTier {
	mon, _ := sh.QuickAssembleMon(monID, nil)
	altState := mon.AlteredStates[0]

	if len(altState.Alts) == 0 {
		return nil
	}

	baseStats := altState.Alts[0].BaseStats

	aglBS := th.GetBaseStat("agility", baseStats)
	tier := getTierByAgility(aglBS.Value)

	return &tier
}

func getTurnOrderMonAgility(monID int32, monster apitypes.Monster) int32 {
	agility := th.GetBaseStatVal("agility", monster.BaseStats)

	return handleNoTurnMons(monID, agility)
}

func handleNoTurnMons(monID, agility int32) int32 {
	noTurnIDs := []int32{33, 166, 194, 212, 228, 300}

	if slices.Contains(noTurnIDs, monID) {
		return 0
	}

	return agility
}
