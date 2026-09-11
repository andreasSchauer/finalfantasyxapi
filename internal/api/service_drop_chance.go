package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type DropChanceResponse struct {
	URL              string           `json:"url"`
	TotalDropChances TotalDropChances `json:"total_drop_chances"`
	EquipmentChances EquipmentChances `json:"equipment_chances"`
	CharacterChances CharacterChances `json:"character_chances"`
}

func (r DropChanceResponse) GetURL() string {
	return r.URL
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	r.TotalDropChances = r.TotalDropChances.Percent()
	r.EquipmentChances = r.EquipmentChances.Percent()
	r.CharacterChances = r.CharacterChances.Percent()

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

type EquipmentChances struct {
	MonsterDrop    float64 `json:"monster_drop"`
	MatchFinBlow   float64 `json:"match_fin_blow"`
	MatchNoFinBlow float64 `json:"match_no_fin_blow"`
}

func (c EquipmentChances) Percent() EquipmentChances {
	c.MonsterDrop = h.DecimalToPercent(c.MonsterDrop)
	c.MatchFinBlow = h.DecimalToPercent(c.MatchFinBlow)
	c.MatchNoFinBlow = h.DecimalToPercent(c.MatchNoFinBlow)

	return c
}

func calcDropChance(cfg *Config, params DropChanceParams, url string) (DropChanceResponse, error) {
	mon, _ := seeding.GetResourceByID(params.Monster, cfg.l.MonstersID)
	charPtr := getCharPtr(cfg, params)

	wantedAbilities, monAbilities, err := vfAutoAbilities(cfg, params, charPtr, mon)
	if err != nil {
		return DropChanceResponse{}, err
	}

	characterChances := calcCharRandomChances(cfg, params, wantedAbilities, monAbilities)
	monDropChance := calcMonDropChance(mon)
	matchFinBlow, matchNoFinBlow := calcEquipmentMatchChances(cfg, params, mon, wantedAbilities, monAbilities, characterChances)

	response := DropChanceResponse{
		URL: url,
		TotalDropChances: TotalDropChances{
			WithFinBlow: calcTotalChance(monDropChance, matchFinBlow),
			NoFinBlow:   calcTotalChance(monDropChance, matchNoFinBlow),
		},
		EquipmentChances: EquipmentChances{
			MonsterDrop:    monDropChance,
			MatchFinBlow:   matchFinBlow,
			MatchNoFinBlow: matchNoFinBlow,
		},
		CharacterChances: characterChances,
	}

	return response.Percent(), nil
}

func calcTotalChance(monDropChance, matchChance float64) float64 {
	equipTypeChance := 0.5
	return monDropChance * equipTypeChance * matchChance
}

func calcMonDropChance(mon seeding.Monster) float64 {
	return float64(mon.Equipment.DropChance) / 255
}

func getRequiredSlots(params DropChanceParams) int32 {
	var minEmptySlots int32 = 0

	if params.MinEmptySlots != nil {
		minEmptySlots = *params.MinEmptySlots
	}

	return int32(len(params.AutoAbilities)) + minEmptySlots
}
