package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
)

func handleDropChance(params servicetypes.DropChanceParams) (servicetypes.DropChanceResponse, error) {
	mon, _ := lookup.GetResourceByID(params.Monster, lookup.L.MonstersID)

	wantedAbilities, monAbilities, err := vfAutoAbilities(params, mon)
	if err != nil {
		return servicetypes.DropChanceResponse{}, err
	}

	slotsTable := extractAbilitySlots(params, mon)
	monDropChances := getMonsterDropChances(mon, monAbilities)

	ccp := CharChanceParams{
		monMaxSlots:     getEquipMaxSlots(slotsTable),
		wantedAbilities: wantedAbilities,
		monAbilities:    monAbilities,
	}
	characterChances := ccp.calcChances(params)

	mcp := MatchChanceParams{
		mon:              mon,
		wantedAbilities:  wantedAbilities,
		monAbilities:     monAbilities,
		slotsTable:       slotsTable,
		shotsTable:       mon.Equipment.AttachedAbilities.Chances,
		characterChances: characterChances,
	}
	equipmentMatchChances := mcp.calcChances(params)

	response := servicetypes.DropChanceResponse{
		DropChanceFinBlow:     monDropChances.EquipmentDropChance * equipmentMatchChances.FinBlow,
		DropChanceNoFinBlow:   monDropChances.EquipmentDropChance * equipmentMatchChances.NoFinBlow,
		WantedEquipment:       getWantedEquipment(params, wantedAbilities, slotsTable),
		EquipmentMatchChances: equipmentMatchChances,
		CharacterChances:      characterChances,
		MonsterChances:        monDropChances,
	}

	if params.Decimals {
		return response.Round(6), nil
	}

	return response.Percent(), nil
}
