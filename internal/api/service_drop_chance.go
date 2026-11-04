package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type DropChanceResponse struct {
	DropChanceFinBlow     float64               `json:"drop_chance_fin_blow"`
	DropChanceNoFinBlow   float64               `json:"drop_chance_no_fin_blow"`
	WantedEquipment       WantedEquipment       `json:"wanted_equipment"`
	EquipmentMatchChances EquipmentMatchChances `json:"equipment_match_chances"`
	CharacterChances      CharacterChances      `json:"character_chances"`
	MonsterChances        MonsterDropChances    `json:"monster_chances"`
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow:     h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow:   h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment:       r.WantedEquipment,
		EquipmentMatchChances: r.EquipmentMatchChances.Percent(),
		CharacterChances:      r.CharacterChances.Percent(),
		MonsterChances:        r.MonsterChances.Percent(),
	}
}

func (r DropChanceResponse) Round(n int32) DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow:     h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow:   h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment:       r.WantedEquipment,
		EquipmentMatchChances: r.EquipmentMatchChances.Round(n),
		CharacterChances:      r.CharacterChances.Round(n),
		MonsterChances:        r.MonsterChances.Round(n),
	}
}

func handleDropChance(srv *Server, params DropChanceParams) (DropChanceResponse, error) {
	mon, _ := lookup.GetResourceByID(params.Monster, lookup.L.MonstersID)

	wantedAbilities, monAbilities, err := vfAutoAbilities(params, mon)
	if err != nil {
		return DropChanceResponse{}, err
	}

	slotsTable := extractAbilitySlots(params, mon)
	monDropChances := getMonsterDropChances(srv, mon, monAbilities)

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

	response := DropChanceResponse{
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
