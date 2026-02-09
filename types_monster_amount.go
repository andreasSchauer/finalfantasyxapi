package main

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type MonsterAmount struct {
	Monster NamedAPIResource `json:"monster"`
	Amount  int32            `json:"amount"`
}

func (ma MonsterAmount) IsZero() bool {
	return ma.Monster.Name == ""
}

func (ma MonsterAmount) GetAPIResource() APIResource {
	return ma.Monster
}

func (ma MonsterAmount) GetName() string {
	return nameToString(ma.Monster.Name, ma.Monster.Version, nil)
}

func (ma MonsterAmount) GetVersion() *int32 {
	return ma.Monster.Version
}

func (ma MonsterAmount) GetVal() int32 {
	return ma.Amount
}

func convertMonsterAmount(cfg *Config, input seeding.MonsterAmount) MonsterAmount {
	return nameToResourceAmount(cfg, cfg.e.monsters, input, newMonsterAmount)
}

func newMonsterAmount(res NamedAPIResource, amount int32) MonsterAmount {
	return MonsterAmount{
		Monster: res,
		Amount:  amount,
	}
}
