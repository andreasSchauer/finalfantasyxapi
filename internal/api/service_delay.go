package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



type DelayResponse struct {
	TickSpeed		int32		`json:"tick_speed"`
	RemainingTicks	*int32		`json:"remaining_ticks,omitempty"`
	DelayTicks		int32		`json:"delay_ticks"`
	DelayTurns		float64		`json:"delay_turns"`
}

type Delay struct {
	DelayType 		string
	AttackType		string
	DelayConstant	int32
}

/*
	code is correct, but still ugly
*/

func handleDelay(cfg *Config, params DelayParams) (DelayResponse, error) {
	delay := assembleDelay(params)
	
	tickspeed, err := assembleDelayTarget(cfg, params)
	if err != nil {
		return DelayResponse{}, err
	}

	if params.Delay.DelayType == string(database.DelayTypeCtbBased) && params.Target.RemainingTicks == nil {
		return DelayResponse{}, newHTTPError(http.StatusBadRequest, "in order to calculate ctb-based delay, the target's remaining ticks need to be given.", nil)
	}

	response := calcDelay(delay, tickspeed, params.Target.RemainingTicks)

	return response, nil
}

func calcDelay(delay Delay, tickspeed int32, remainingTicks *int32) DelayResponse {
	var attackType int32 = 1
	
	if delay.AttackType == string(database.CtbAttackTypeHeal) {
		attackType = -1
	}

	var usedVal int32
	
	switch delay.DelayType {
	case string(database.DelayTypeTickSpeedBased):
		usedVal = tickspeed

	case string(database.DelayTypeCtbBased):
		usedVal = *remainingTicks
	}

	delayTicks := (usedVal * delay.DelayConstant) / 16
	turnTicks := tickspeed * 3
	delayTurns := float64(delayTicks) / float64(turnTicks)

	return DelayResponse{
		DelayTicks: 	delayTicks * attackType,
		DelayTurns: 	h.FloatRound(delayTurns, 4),
		TickSpeed: 		tickspeed,
		RemainingTicks: remainingTicks,
	}
}


func assembleDelay(params DelayParams) Delay {
	d := params.Delay
	delay := Delay{
		DelayType: d.DelayType,
		AttackType: d.AttackType,
	}

	if d.DelayConstant != nil || d.Strength == nil {
		delay.DelayConstant = *d.DelayConstant
		return delay
	}

	switch {
	case d.DelayType == string(database.DelayTypeTickSpeedBased):
		switch {
		case *d.Strength == string(database.DelayStrengthStrong):
			delay.DelayConstant = 48

		case *d.Strength == string(database.DelayStrengthWeak):
			delay.DelayConstant = 24
		}

	case d.DelayType == string(database.DelayTypeCtbBased):
		switch {
		case *d.Strength == string(database.DelayStrengthStrong):
			delay.DelayConstant = 16

		case *d.Strength == string(database.DelayStrengthWeak):
			delay.DelayConstant = 8
		}
	}

	return delay
}

func assembleDelayTarget(cfg *Config, params DelayParams) (int32, error) {
	status, agility, err := fetchDelayMonster(cfg, params)
	if err != nil {
		return 0, err
	}

	if params.Target.Agility != nil {
		agility = *params.Target.Agility
	}

	if params.Target.Status != nil {
		status = params.Target.Status
	}

	agilityTier := getAgilityTier(cfg, agility)
	tickspeed := calcTickSpeed(agilityTier.TickSpeed, status)

	return tickspeed, nil
}


func fetchDelayMonster(cfg *Config, params DelayParams) (*string, int32, error) {
	target := params.Target
	
	if target.MonsterID == nil {
		return nil, 0, nil
	}

	monster, err := quickAssembleMon(cfg, *target.MonsterID, target.AltState)
	if err != nil {
		return nil, 0, err
	}

	err = enforceMonImmunity("delay", monster, params.IgnImmunities)
	if err != nil {
		return nil, 0, err
	}

	agility := getBaseStatVal(cfg, "agility", monster.BaseStats)

	status, err := fetchMonsterHasteStatus(monster, target.Status, params.IgnImmunities)
	if err != nil {
		return nil, 0, err
	}

	return status, agility, nil
}