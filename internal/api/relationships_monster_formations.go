package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getFormationLoot(cfg *Config, formation seedtypes.MonsterFormation) FormationLoot {
	var loot FormationLoot

	for _, monAmt := range formation.Monsters {
		mon, _ := seeding.GetResourceByID(monAmt.MonsterID, cfg.L.MonstersID)

		loot.AP += mon.AP * monAmt.Amount
		loot.ApOverkill += mon.APOverkill * monAmt.Amount
		loot.Gil += mon.Gil * monAmt.Amount
	}

	return loot
}
