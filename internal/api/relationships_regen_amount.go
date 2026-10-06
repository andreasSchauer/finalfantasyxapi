package api

type RegenAmount struct {
	Ticks		int32		`json:"ticks"`
	HealedHP	int32		`json:"healed_hp"`
}

func calcRegenAmountsMonster(cfg *Config, mon Monster) []RegenAmount {
	regenRes := nameToNamedAPIResource(cfg, cfg.e.statusConditions, "regen", nil)
	if resourcesContain(mon.StatusImmunities, regenRes) {
		return nil
	}

	return calcRegenAmounts(cfg, mon.BaseStats)
}

func calcRegenAmounts(cfg *Config, baseStats []BaseStat) []RegenAmount {
	ticks := []int32{0, 1, 2, 3, 4, 5, 10, 15, 20, 30, 50}
	regenAmounts := make([]RegenAmount, len(ticks))
	hp := getBaseStatVal(cfg, "hp", baseStats)

	for i, tickAmt := range ticks {
		regenAmounts[i] = getRegenAmount(hp, tickAmt)
	}

	return regenAmounts
}

func getRegenAmount(hp, ticks int32) RegenAmount {
	return RegenAmount{
		Ticks: 	  ticks,
		HealedHP: calcRegenAmount(hp, ticks),
	}
}

func calcRegenAmount(hp, ticks int32) int32 {
	hpFloat := float64(hp)
	ticksFloat := float64(ticks)
	healedFloat := ticksFloat * hpFloat / 256 + 100
	return int32(healedFloat)
}