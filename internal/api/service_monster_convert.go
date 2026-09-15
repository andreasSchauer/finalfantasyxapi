package api


import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func getConvertedMon(cfg *Config, monID int32) Monster {
	monsterLookup, _ := seeding.GetResourceByID(monID, cfg.l.MonstersID)

	return Monster{
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
}