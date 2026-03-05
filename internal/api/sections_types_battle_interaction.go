package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type BattleInteractionSimple struct {
	Target						string				`json:"target"`
	Range						*int32				`json:"range"`
	Damage						[]DamageCalcSimple	`json:"damage"`
	CanCritical 				bool				`json:"can_critical"`
	CanBreakDmgLimit			bool				`json:"can_break_dmg_limit"`
	Element						*string				`json:"element"`
	AffectedBy 					[]string			`json:"affected_by"`
	HitAmount					int32				`json:"hit_amount"`
	Delay						*string				`json:"delay"`
	InflictedStatusConditions	[]string			`json:"inflicted_status_conditions"`
	RemovedStatusConditions 	[]string			`json:"removed_status_conditions"`
	CopiedStatusConditions 		[]string			`json:"copied_status_conditions"`
	StatChanges 				[]string			`json:"stat_changes"`
	ModChanges					[]string			`json:"modifier_changes"`
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
		ModChanges: 				convertObjSlice(cfg, bi.ModifierChanges, convertModChangeSimple),
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

type DamageCalcSimple struct {
	TargetStat		string			`json:"target_stat"`
	DamageType		string			`json:"damage_type"`
	DamageFormula	string			`json:"damage_formula"`
	DamageConstant	int32			`json:"damage_constant"`
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

func convertInflictedDelaySimple(_ *Config, id seeding.InflictedDelay) string {
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


func convertInflictedStatusSimple(_ *Config, is seeding.InflictedStatus) string {
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

func convertStatChangeSimple(_ *Config, sc seeding.StatChange) string {
	var operatorStr string

	switch sc.CalculationType {
	case string(database.CalculationTypeAddedValue):
		operatorStr = "+"

	case string(database.CalculationTypeMultiplyHighest):
		operatorStr = "x"
	}

	return fmt.Sprintf("%s %s%d", sc.StatName, operatorStr, int32(sc.Value))
}

func convertModChangeSimple(_ *Config, mc seeding.ModifierChange) string {
	formattedVal := fmt.Sprintf("%.2f", mc.Value)

	switch mc.CalculationType {
	case string(database.CalculationTypeAddedValue):
		if strings.HasPrefix(formattedVal, "-") {
			return fmt.Sprintf("%s %s", mc.ModifierName, formattedVal)
		}
		return fmt.Sprintf("%s +%s", mc.ModifierName, formattedVal)

	case string(database.CalculationTypeAddedPercentage):
		if strings.HasPrefix(formattedVal, "-") {
			return fmt.Sprintf("%s %s", mc.ModifierName, formattedVal)
		}
		return fmt.Sprintf("%s +%s", mc.ModifierName, formattedVal) + "%"

	case string(database.CalculationTypeMultiplyHighest), string(database.CalculationTypeMultiply):
		return fmt.Sprintf("%s x%s", mc.ModifierName, formattedVal)

	case string(database.CalculationTypeSetValue):
		return fmt.Sprintf("%s = %d", mc.ModifierName, int32(mc.Value))
	}

	return ""
}