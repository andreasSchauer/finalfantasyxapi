package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type BattleInteractionSimple struct {
	Target                    string        `json:"target"`
	Range                     *int32        `json:"range,omitempty"`
	HitAmount                 int32         `json:"hit_amount"`
	Accuracy                  string        `json:"accuracy"`
	AffectedBy                []string      `json:"affected_by,omitempty"`
	Damage                    *DamageSimple `json:"damage,omitempty"`
	Delay                     *string       `json:"delay,omitempty"`
	InflictedStatusConditions []string      `json:"inflicted_status_conditions,omitempty"`
	RemovedStatusConditions   []string      `json:"removed_status_conditions,omitempty"`
	CopiedStatusConditions    []string      `json:"copied_status_conditions,omitempty"`
	StatChanges               []string      `json:"stat_changes,omitempty"`
	ModChanges                []string      `json:"modifier_changes,omitempty"`
}

func convertBattleInteractionSimple(cfg *Config, bi seeding.BattleInteraction) BattleInteractionSimple {
	biSimple := BattleInteractionSimple{
		Target:                    bi.Target,
		Range:                     bi.Range,
		HitAmount:                 bi.HitAmount,
		Accuracy:                  convertAccuracySimple(cfg, bi.Accuracy),
		AffectedBy:                h.SliceOrNil(bi.AffectedBy),
		Damage:                    convertObjPtr(cfg, bi.Damage, convertDamageSimple),
		Delay:                     convertObjPtr(cfg, bi.InflictedDelay, convertInflictedDelaySimple),
		InflictedStatusConditions: convertObjSliceOrNil(cfg, bi.InflictedStatusConditions, convertInflictedStatusSimple),
		RemovedStatusConditions:   h.SliceOrNil(bi.RemovedStatusConditions),
		CopiedStatusConditions:    convertObjSliceOrNil(cfg, bi.CopiedStatusConditions, convertInflictedStatusSimple),
		StatChanges:               convertObjSliceOrNil(cfg, bi.StatChanges, convertStatChangeSimple),
		ModChanges:                convertObjSliceOrNil(cfg, bi.ModifierChanges, convertModChangeSimple),
	}

	return biSimple
}

type DamageSimple struct {
	CanCritical   bool     `json:"can_critical"`
	BreakDmgLimit *string  `json:"break_dmg_limit,omitempty"`
	Element       *string  `json:"element,omitempty"`
	DamageCalc    []string `json:"damage_calc"`
}

func convertDamageSimple(cfg *Config, d seeding.Damage) DamageSimple {
	return DamageSimple{
		CanCritical:   h.PtrIsNotNil(d.Critical),
		BreakDmgLimit: convertObjPtr(cfg, d.BreakDmgLimit, convertBreakDmgLimitSimple),
		Element:       d.Element,
		DamageCalc:    convertObjSlice(cfg, d.DamageCalc, convertDamageCalcSimple),
	}
}

func convertDamageCalcSimple(cfg *Config, dc seeding.AbilityDamage) string {
	return fmt.Sprintf("%s %s (%s), formula: %s, power: %d", dc.AttackType, dc.TargetStat, dc.DamageType, dc.DamageFormula, dc.DamageConstant)
}

func convertBreakDmgLimitSimple(_ *Config, breakDmgLimit string) string {
	switch breakDmgLimit {
	case string(database.BreakDmgLmtTypeAlways):
		return "always"

	case string(database.BreakDmgLmtTypeAutoAbility):
		return "with auto-ability"
	}

	return ""
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
		probabilityStr = strconv.Itoa(int(is.Probability)) + "%"
	}

	return fmt.Sprintf("%s (%s)", is.StatusCondition, probabilityStr)
}

func convertStatChangeSimple(_ *Config, sc seeding.StatChange) string {
	return formatChange(sc.StatName, sc.CalculationType, sc.Value)
}

func convertModChangeSimple(cfg *Config, mc seeding.ModifierChange) string {
	formatted := formatChange(mc.ModifierName, mc.CalculationType, mc.Value)

	modifier, _ := seeding.GetResource(mc.ModifierName, cfg.l.Modifiers)

	if modifier.Type == string(database.ModifierTypePercentage) {
		formatted += "%"
	}

	return formatted
}


func formatChange(name, calcType string, val float32) string {
	formattedVal := strconv.FormatFloat(float64(val), 'f', -1, 32)

	switch calcType {
	case string(database.CalculationTypeAddedValue):
		if strings.HasPrefix(formattedVal, "-") {
			return fmt.Sprintf("%s %s", name, formattedVal)
		}
		return fmt.Sprintf("%s +%s", name, formattedVal)

	case string(database.CalculationTypeAddedPercentage):
		if strings.HasPrefix(formattedVal, "-") {
			return fmt.Sprintf("%s %s", name, formattedVal) + "%"
		}
		return fmt.Sprintf("%s +%s", name, formattedVal) + "%"

	case string(database.CalculationTypeMultiplyHighest), string(database.CalculationTypeMultiply):
		return fmt.Sprintf("%s x%s", name, formattedVal)

	case string(database.CalculationTypeSetValue):
		return fmt.Sprintf("%s = %d", name, int32(val))
	}

	return ""
}