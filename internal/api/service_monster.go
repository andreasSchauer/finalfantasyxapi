package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func quickAssembleMon(cfg *Config, monID int32, altStatePtr *int32) (Monster, error) {
	monsterLookup, _ := seeding.GetResourceByID(monID, cfg.l.MonstersID)

	monster := Monster{
		ID:               monsterLookup.ID,
		Name:             monsterLookup.Name,
		Version:          monsterLookup.Version,
		Specification:    monsterLookup.Specification,
		HasOverdrive:     monsterLookup.HasOverdrive,
		IsUnderwater:     monsterLookup.IsUnderwater,
		IsZombie:         monsterLookup.IsZombie,
		Distance:         monsterLookup.Distance,
		Properties:       namesToNamedAPIResources(cfg, cfg.e.properties, monsterLookup.Properties),
		AutoAbilities:    namesToNamedAPIResources(cfg, cfg.e.autoAbilities, monsterLookup.AutoAbilities),
		StealGil:         monsterLookup.StealGil,
		DoomCountdown:    monsterLookup.DoomCountdown,
		PoisonRate:       monsterLookup.PoisonRate,
		ThreatenChance:   monsterLookup.ThreatenChance,
		ZanmatoLevel:     monsterLookup.ZanmatoLevel,
		BaseStats:        toResAmtType(cfg, cfg.e.stats, monsterLookup.BaseStats, newBaseStat),
		ElemResists:      getMonsterElemResists(cfg, monsterLookup.ElemResists),
		StatusImmunities: namesToNamedAPIResources(cfg, cfg.e.statusConditions, monsterLookup.StatusImmunities),
		StatusResists:    toResAmtType(cfg, cfg.e.statusConditions, monsterLookup.StatusResists, newStatusResist),
		Abilities:        convertObjSlice(cfg, monsterLookup.Abilities, convertMonsterAbility),
		AlteredStates:    getMonsterAlteredStates(cfg, nil, monsterLookup),
	}

	monster, err := applyAlteredStateFromJson(cfg, monster, altStatePtr)
	if err != nil {
		return Monster{}, err
	}

	return monster, nil
}

func monsterHasFirstStrike(mon Monster) bool {
	for _, aa := range mon.AutoAbilities {
		if aa.Name == "first strike" {
			return true
		}
	}

	return false
}

func monHasAppliedStatus(mon Monster) bool {
	return mon.AppliedState != nil && mon.AppliedState.AppliedStatus != nil
}

func monImmuneToHaste(mon Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == string(database.HasteStatusHaste) {
			return true
		}
	}

	return false
}

func monImmuneToSlow(mon Monster) bool {
	for _, condition := range mon.StatusImmunities {
		if condition.Name == string(database.HasteStatusSlow) {
			return true
		}
	}

	return false
}