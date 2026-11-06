package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getFormationLoot(formation seedtypes.MonsterFormation) apitypes.FormationLoot {
	var loot apitypes.FormationLoot

	for _, monAmt := range formation.Monsters {
		mon, _ := lookup.GetResourceByID(monAmt.MonsterID, lookup.L.MonstersID)

		loot.AP += mon.AP * monAmt.Amount
		loot.ApOverkill += mon.APOverkill * monAmt.Amount
		loot.Gil += mon.Gil * monAmt.Amount
	}

	return loot
}
