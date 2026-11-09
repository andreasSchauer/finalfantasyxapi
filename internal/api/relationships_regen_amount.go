package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
)

func calcRegenAmountsMonster(mon apitypes.Monster) []apitypes.RegenAmount {
	regenRes := handlerinput.NameToNamedAPIResource(apilookups.E.StatusConditions, "regen", nil)
	if handlerinput.ResourcesContain(mon.StatusImmunities, regenRes) {
		return nil
	}

	return calcRegenAmounts(mon.BaseStats)
}

func calcRegenAmounts(baseStats []apitypes.BaseStat) []apitypes.RegenAmount {
	ticks := []int32{0, 1, 2, 3, 4, 5, 10, 15, 20, 30, 50}
	regenAmounts := make([]apitypes.RegenAmount, len(ticks))
	hp := getBaseStatVal("hp", baseStats)

	for i, tickAmt := range ticks {
		regenAmounts[i] = getRegenAmount(hp, tickAmt)
	}

	return regenAmounts
}

func getRegenAmount(hp, ticks int32) apitypes.RegenAmount {
	return apitypes.RegenAmount{
		Ticks:    ticks,
		HealedHP: calcRegenAmount(hp, ticks),
	}
}

func calcRegenAmount(hp, ticks int32) int32 {
	hpFloat := float64(hp)
	ticksFloat := float64(ticks)
	healedFloat := ticksFloat*hpFloat/256 + 100
	return int32(healedFloat)
}
