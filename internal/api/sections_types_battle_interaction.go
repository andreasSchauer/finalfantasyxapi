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
	HitAmount					int32				`json:"hit_amount"`
	Accuracy					string				`json:"accuracy"`
	AffectedBy 					[]string			`json:"affected_by,omitempty"`
	Damage						*DamageSimple		`json:"damage,omitempty"`
	Delay						*string				`json:"delay,omitempty"`
	InflictedStatusConditions	[]string			`json:"inflicted_status_conditions,omitempty"`
	RemovedStatusConditions 	[]string			`json:"removed_status_conditions,omitempty"`
	CopiedStatusConditions 		[]string			`json:"copied_status_conditions,omitempty"`
	StatChanges 				[]string			`json:"stat_changes,omitempty"`
	ModChanges					[]string			`json:"modifier_changes,omitempty"`
}


func convertBattleInteractionSimple(cfg *Config, bi seeding.BattleInteraction) BattleInteractionSimple {
	biSimple := BattleInteractionSimple{
		Target: 					bi.Target,
		Range: 						bi.Range,
		AffectedBy: 				sliceOrNil(bi.AffectedBy),
		HitAmount: 					bi.HitAmount,
		Accuracy: 					convertAccuracySimple(cfg, bi.Accuracy),
		Damage: 					convertObjPtr(cfg, bi.Damage, convertDamageSimple),
		Delay: 						convertObjPtr(cfg, bi.InflictedDelay, convertInflictedDelaySimple),
		InflictedStatusConditions: 	convertObjSliceNullable(cfg, bi.InflictedStatusConditions, convertInflictedStatusSimple),
		RemovedStatusConditions: 	sliceOrNil(bi.RemovedStatusConditions),
		CopiedStatusConditions: 	convertObjSliceNullable(cfg, bi.CopiedStatusConditions, convertInflictedStatusSimple),
		StatChanges: 				convertObjSliceNullable(cfg, bi.StatChanges, convertStatChangeSimple),
		ModChanges: 				convertObjSliceNullable(cfg, bi.ModifierChanges, convertModChangeSimple),
	}

	return biSimple
}

type DamageSimple struct {
	CanCritical 		bool		`json:"can_critical"`
	CanBreakDmgLimit	bool		`json:"can_break_dmg_limit"`
	Element				*string		`json:"element,omitempty"`
	DamageCalc			[]string	`json:"damage_calc"`
}

func convertDamageSimple(cfg *Config, d seeding.Damage) DamageSimple {
	return DamageSimple{
		CanCritical: 		ptrIsNotNil(d.Critical),
		CanBreakDmgLimit: 	ptrIsNotNil(d.BreakDmgLimit),
		Element: 			d.Element,
		DamageCalc: 		convertObjSlice(cfg, d.DamageCalc, convertDamageCalcSimple),
	}
}

func convertDamageCalcSimple(cfg *Config, dc seeding.AbilityDamage) string {
	return fmt.Sprintf("%s %s (%s), formula: %s, power: %d", dc.AttackType, dc.TargetStat, dc.DamageType, dc.DamageFormula, dc.DamageConstant)
}


func convertAccuracySimple(_ *Config, acc seeding.Accuracy) string {
	switch acc.AccSource {
	case string(database.AccSourceTypeRate):
		if *acc.HitChance == 255 {
			return "always hits"
		}
		return fmt.Sprintf("%d", *acc.HitChance) + "% base chance of hitting"

	case string(database.AccSourceTypeAccuracy):
		return fmt.Sprintf("based on accuracy with %.1f modifier", *acc.AccModifier)
	}

	return ""
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