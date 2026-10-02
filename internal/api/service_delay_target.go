package api

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func delayGetTargetData(cfg *Config, params DelayParams) (DelayTargetData, error) {
	target, err := delayFetchMonData(cfg, params)
	if err != nil {
		return DelayTargetData{}, err
	}

	if params.Target.Agility != nil {
		target.Agility = *params.Target.Agility
	}

	if params.Target.Status != nil {
		target.Status = params.Target.Status
	}

	agilityTier := getAgilityTier(cfg, target.Agility)
	target.TickSpeed = calcTickSpeed(agilityTier.TickSpeed, target.Status)

	return target, nil
}

func delayFetchMonData(cfg *Config, params DelayParams) (DelayTargetData, error) {
	target := params.Target
	
	if target.MonsterID == nil {
		return DelayTargetData{}, nil
	}

	monster, err := quickAssembleMon(cfg, *target.MonsterID, target.AltState)
	if err != nil {
		return DelayTargetData{}, err
	}

	err = enforceMonImmunity("delay", monster, params.IgnImmunities)
	if err != nil {
		return DelayTargetData{}, err
	}

	agility := getBaseStatVal(cfg, "agility", monster.BaseStats)

	status, err := fetchMonsterHasteStatus(monster, target.Status, params.IgnImmunities)
	if err != nil {
		return DelayTargetData{}, err
	}
	monName := h.NameToString(monster.Name, monster.Version, monster.Specification)

	targetData := DelayTargetData{
		Monster: &monName,
		Agility: agility,
		Status:  status,
	}

	return targetData, nil
}