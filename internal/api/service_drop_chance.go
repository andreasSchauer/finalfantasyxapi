package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DropChanceResponse struct {
	TotalDropChances   TotalDropChances   `json:"total_drop_chances"`
	EquipmentChances   EquipmentChances   `json:"equipment_chances"`
	CharacterChances   CharacterChances   `json:"character_chances"`
	MonsterDropChances MonsterDropChances `json:"monster_drop_chances"`
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	r.TotalDropChances = r.TotalDropChances.Percent()
	r.EquipmentChances = r.EquipmentChances.Percent()
	r.CharacterChances = r.CharacterChances.Percent()
	r.MonsterDropChances = r.MonsterDropChances.Percent()

	return r
}

func (r DropChanceResponse) Round(n int32) DropChanceResponse {
	r.TotalDropChances = r.TotalDropChances.Round(n)
	r.EquipmentChances = r.EquipmentChances.Round(n)
	r.CharacterChances = r.CharacterChances.Round(n)
	r.MonsterDropChances = r.MonsterDropChances.Round(n)

	return r
}

type TotalDropChances struct {
	WithFinBlow float64 `json:"with_fin_blow"`
	NoFinBlow   float64 `json:"no_fin_blow"`
}

func (c TotalDropChances) Percent() TotalDropChances {
	c.WithFinBlow = h.DecimalToPercent(c.WithFinBlow)
	c.NoFinBlow = h.DecimalToPercent(c.NoFinBlow)

	return c
}

func (c TotalDropChances) Round(n int32) TotalDropChances {
	c.WithFinBlow = h.FloatRound(c.WithFinBlow, n)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, n)

	return c
}

type EquipmentChances struct {
	MatchFinBlow   float64 `json:"match_fin_blow"`
	MatchNoFinBlow float64 `json:"match_no_fin_blow"`
}

func (c EquipmentChances) Percent() EquipmentChances {
	c.MatchFinBlow = h.DecimalToPercent(c.MatchFinBlow)
	c.MatchNoFinBlow = h.DecimalToPercent(c.MatchNoFinBlow)

	return c
}

func (c EquipmentChances) Round(n int32) EquipmentChances {
	c.MatchFinBlow = h.FloatRound(c.MatchFinBlow, n)
	c.MatchNoFinBlow = h.FloatRound(c.MatchNoFinBlow, n)

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
	matchFinBlow, matchNoFinBlow := calcEquipmentMatchChances(cfg, params, mon, wantedAbilities, monAbilities, characterChances)

	response := DropChanceResponse{
		TotalDropChances: TotalDropChances{
			WithFinBlow: monDropChances.EquipmentDropChance * matchFinBlow,
			NoFinBlow:   monDropChances.EquipmentDropChance * matchNoFinBlow,
		},
		EquipmentChances: EquipmentChances{
			MatchFinBlow:   matchFinBlow,
			MatchNoFinBlow: matchNoFinBlow,
		},
		CharacterChances:   characterChances,
		MonsterDropChances: getMonsterDropChances(cfg, mon, monAbilities),
	}

	if params.Decimals {
		return response.Round(6), nil
	}

	return response.Percent(), nil
}