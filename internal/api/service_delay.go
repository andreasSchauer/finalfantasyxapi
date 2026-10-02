package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type DelayResponse struct {
	Delay			Delay			`json:"delay"`
	Target			DelayTargetData	`json:"target"`
	DelayTicks		int32			`json:"delay_ticks"`
	DelayTurns		float64			`json:"delay_turns"`
}

type Delay struct {
	DelayType 		string	`json:"delay_type"`
	AttackType		string	`json:"attack_type"`
	DelayConstant	int32	`json:"delay_constant"`
}

type DelayTargetData struct {
	Monster			*string		`json:"monster,omitempty"`
	Agility			int32		`json:"agility"`
	TickSpeed		int32		`json:"tick_speed"`
	RemainingTicks	*int32		`json:"remaining_ticks,omitempty"`
	Status			*string		`json:"status,omitempty"`
}


func handleDelay(cfg *Config, params DelayParams) (DelayResponse, error) {
	delay := assembleDelay(params)
	
	target, err := delayGetTargetData(cfg, params)
	if err != nil {
		return DelayResponse{}, err
	}
	target.RemainingTicks = params.Target.RemainingTicks

	if delay.DelayType == string(database.DelayTypeCtbBased) && target.RemainingTicks == nil {
		return DelayResponse{}, newHTTPError(http.StatusBadRequest, "in order to calculate ctb-based delay, the target's remaining ticks need to be given.", nil)
	}
	
	delayTicks, delayTurns := calcDelay(delay, target)

	response := DelayResponse{
		Delay: 			delay,
		Target: 		target,
		DelayTicks: 	delayTicks,
		DelayTurns: 	delayTurns,
	}

	return response, nil
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

func calcDelay(delay Delay, target DelayTargetData) (int32, float64) {
	attackType := getDelayAttackTypeFactor(delay)
	usedVal := getDelayUsedVal(delay, target)

	delayTicks := (usedVal * delay.DelayConstant) / 16
	ticksPerTurn := target.TickSpeed * 3
	delayTurns := float64(delayTicks) / float64(ticksPerTurn)

	return delayTicks * attackType, h.FloatRound(delayTurns, 4)
}

func getDelayAttackTypeFactor(delay Delay) int32 {
	var attackType int32 = 1
	
	if delay.AttackType == string(database.CtbAttackTypeHeal) {
		attackType = -1
	}

	return attackType
}

func getDelayUsedVal(delay Delay, target DelayTargetData) int32 {
	if delay.DelayType == string(database.DelayTypeTickSpeedBased) {
		return target.TickSpeed
	}

	return *target.RemainingTicks
}