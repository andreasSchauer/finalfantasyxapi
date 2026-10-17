package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DropChanceResponse struct {
	DropChanceFinBlow     float64               `json:"drop_chance_fin_blow"`
	DropChanceNoFinBlow   float64               `json:"drop_chance_no_fin_blow"`
	EquipmentMatchChances EquipmentMatchChances `json:"equipment_match_chances"`
	CharacterChances      CharacterChances      `json:"character_chances"`
	MonsterChances        MonsterDropChances    `json:"monster_chances"`
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	r.DropChanceFinBlow = h.DecimalToPercent(r.DropChanceFinBlow)
	r.DropChanceNoFinBlow = h.DecimalToPercent(r.DropChanceNoFinBlow)
	r.EquipmentMatchChances = r.EquipmentMatchChances.Percent()
	r.CharacterChances = r.CharacterChances.Percent()
	r.MonsterChances = r.MonsterChances.Percent()

	return r
}

func (r DropChanceResponse) Round(n int32) DropChanceResponse {
	r.DropChanceFinBlow = h.DecimalToPercent(r.DropChanceFinBlow)
	r.DropChanceNoFinBlow = h.DecimalToPercent(r.DropChanceNoFinBlow)
	r.EquipmentMatchChances = r.EquipmentMatchChances.Round(n)
	r.CharacterChances = r.CharacterChances.Round(n)
	r.MonsterChances = r.MonsterChances.Round(n)

	return r
}

type EquipmentMatchChances struct {
	FinBlow   float64 `json:"fin_blow"`
	NoFinBlow float64 `json:"no_fin_blow"`
}

func (c EquipmentMatchChances) Percent() EquipmentMatchChances {
	c.FinBlow = h.DecimalToPercent(c.FinBlow)
	c.NoFinBlow = h.DecimalToPercent(c.NoFinBlow)

	return c
}

func (c EquipmentMatchChances) Round(n int32) EquipmentMatchChances {
	c.FinBlow = h.FloatRound(c.FinBlow, n)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, n)

	return c
}

func handleDropChance(cfg *Config, params DropChanceParams) (DropChanceResponse, error) {
	mon, _ := seeding.GetResourceByID(params.Monster, cfg.l.MonstersID)

	wantedAbilities, monAbilities, err := vfAutoAbilities(cfg, params, mon)
	if err != nil {
		return DropChanceResponse{}, err
	}

	characterChances := calcCharChances(cfg, params, wantedAbilities, monAbilities)
	monDropChances := getMonsterDropChances(cfg, mon, monAbilities)
	equipmentMatchChances := calcEquipmentMatchChances(cfg, params, mon, wantedAbilities, monAbilities, characterChances)

	response := DropChanceResponse{
		DropChanceFinBlow:     monDropChances.EquipmentDropChance * equipmentMatchChances.FinBlow,
		DropChanceNoFinBlow:   monDropChances.EquipmentDropChance * equipmentMatchChances.NoFinBlow,
		EquipmentMatchChances: equipmentMatchChances,
		CharacterChances:      characterChances,
		MonsterChances:        monDropChances,
	}

	if params.Decimals {
		return response.Round(6), nil
	}

	return response.Percent(), nil
}
