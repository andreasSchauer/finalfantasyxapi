package api

import (
	"fmt"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type BattleInteractionSimple struct {
	Target	string
	Range	*int32
	Damage	[]DamageCalcSimple
	CanCritical bool
	CanBreakDmgLimit	bool
	Element		*string
	AffectedBy []string
	HitAmount	int32
	Delay		*string	 // weak/strong. delay type doesn't matter
	InflictedStatusConditions	[]string // condition (N%, infinite %, always)
	RemovedStatusConditions []string
	CopiedStatusConditions []string // like inflicted
	StatChanges []string // stat +N, or stat xN
}

type DamageCalcSimple struct {
	TargetStat	string // combine attack type + target stat (heal hp, attack hp, absorb mp)
	DamageType	string
	DamageFormula	string
	DamageConstant	int32
}

func convertDamageCalcSimple(cfg *Config, dc seeding.AbilityDamage) DamageCalcSimple {
	return DamageCalcSimple{
		TargetStat: 	convertTargetStatSimple(cfg, dc),
		DamageType: 	dc.DamageType,
		DamageFormula: 	dc.DamageFormula,
		DamageConstant: dc.DamageConstant,
	}
}

func convertTargetStatSimple(_ *Config, dc seeding.AbilityDamage) string {
	return fmt.Sprintf("%s %s", dc.AttackType, dc.TargetStat)
}

func convertInflictedDelaySimple(cfg *Config, id seeding.InflictedDelay) string {
	var delayStr string

	switch id.DelayType {
	case string(database.DelayTypeCtbBased):
		if id.DamageConstant == 8 {
			delayStr = "weak"
		}
		
		if id.DamageConstant == 16 {
			delayStr = "strong"
		}

	case string(database.DelayTypeTickSpeedBased):
		if id.DamageConstant == 24 {
			delayStr = "weak"
		}
		
		if id.DamageConstant == 48 {
			delayStr = "strong"
		}
	}

	return delayStr
}


func convertInflictedStatusSimple(cfg *Config, is seeding.InflictedStatus) string {
	var probabilityStr string

	switch is.Probability {
	case 255:
		probabilityStr = "always"

	case 254:
		probabilityStr = "infinite %"

	default:
		probabilityStr = strconv.Itoa(int(is.Probability)) + " %"
	}

	return fmt.Sprintf("%s (%s)", is.StatusCondition, probabilityStr)
}

func convertStatChangeSimple(cfg *Config, sc seeding.StatChange) string {
	var operatorStr string

	switch sc.CalculationType {
	case string(database.CalculationTypeAddedValue):
		operatorStr = "+"

	case string(database.CalculationTypeMultiplyHighest):
		operatorStr = "x"
	}

	return fmt.Sprintf("%s %s%.2f", sc.StatName, operatorStr, sc.Value)
}


func convertBattleInteractionSimple(cfg *Config, bi seeding.BattleInteraction) BattleInteractionSimple {
	biSimple := BattleInteractionSimple{
		Target: 					bi.Target,
		Range: 						bi.Range,
		AffectedBy: 				bi.AffectedBy,
		HitAmount: 					bi.HitAmount,
		Delay: 						convertObjPtr(cfg, bi.InflictedDelay, convertInflictedDelaySimple),
		InflictedStatusConditions: 	convertObjSlice(cfg, bi.InflictedStatusConditions, convertInflictedStatusSimple),
		RemovedStatusConditions: 	bi.RemovedStatusConditions,
		CopiedStatusConditions: 	convertObjSlice(cfg, bi.CopiedStatusConditions, convertInflictedStatusSimple),
		StatChanges: 				convertObjSlice(cfg, bi.StatChanges, convertStatChangeSimple),
	}

	if bi.Damage != nil {
		biSimple.Damage = convertObjSlice(cfg, bi.Damage.DamageCalc, convertDamageCalcSimple)

		if bi.Damage.Critical != nil {
			biSimple.CanCritical = true
		}

		if bi.Damage.BreakDmgLimit != nil {
			biSimple.CanBreakDmgLimit = true
		}

		biSimple.Element = bi.Damage.Element
	}

	return biSimple
}