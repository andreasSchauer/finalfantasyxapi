package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DropChanceResponse struct {
	DropChanceFinBlow     float64               `json:"drop_chance_fin_blow"`
	DropChanceNoFinBlow   float64               `json:"drop_chance_no_fin_blow"`
	WantedEquipment		  WantedEquipment		`json:"wanted_equipment"`
	EquipmentMatchChances EquipmentMatchChances `json:"equipment_match_chances"`
	CharacterChances      CharacterChances      `json:"character_chances"`
	MonsterChances        MonsterDropChances    `json:"monster_chances"`
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow: 		h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow: 	h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment: 		r.WantedEquipment,
		EquipmentMatchChances: 	r.EquipmentMatchChances.Percent(),
		CharacterChances: 		r.CharacterChances.Percent(),
		MonsterChances: 		r.MonsterChances.Percent(),
	}
}

func (r DropChanceResponse) Round(n int32) DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow: 		h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow: 	h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment: 		r.WantedEquipment,
		EquipmentMatchChances: 	r.EquipmentMatchChances.Round(n),
		CharacterChances: 		r.CharacterChances.Round(n),
		MonsterChances: 		r.MonsterChances.Round(n),
	}
}


func handleDropChance(cfg *Config, params DropChanceParams) (DropChanceResponse, error) {
	mon, _ := seeding.GetResourceByID(params.Monster, cfg.l.MonstersID)
	monMaxSlots := getMonMaxSlots(mon)

	wantedAbilities, monAbilities, err := vfAutoAbilities(cfg, params, mon)
	if err != nil {
		return DropChanceResponse{}, err
	}

	monDropChances := getMonsterDropChances(cfg, mon, monAbilities)
	characterChances := calcCharChances(cfg, params, wantedAbilities, monAbilities, monMaxSlots)
	matchChanceParams := createMatchChanceParams(params, mon, wantedAbilities, monAbilities, characterChances)
	equipmentMatchChances := calcEquipmentMatchChances(cfg, params, matchChanceParams)
	wantedEquipment := getWantedEquipment(params, wantedAbilities, matchChanceParams.slotsTable)

	response := DropChanceResponse{
		DropChanceFinBlow:     monDropChances.EquipmentDropChance * equipmentMatchChances.FinBlow,
		DropChanceNoFinBlow:   monDropChances.EquipmentDropChance * equipmentMatchChances.NoFinBlow,
		WantedEquipment: 	   wantedEquipment,
		EquipmentMatchChances: equipmentMatchChances,
		CharacterChances:      characterChances,
		MonsterChances:        monDropChances,
	}

	if params.Decimals {
		return response.Round(6), nil
	}

	return response.Percent(), nil
}
