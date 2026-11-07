package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertBattleInteraction(bi seedtypes.BattleInteraction) apitypes.BattleInteraction {
	battleInteraction := apitypes.BattleInteraction{
		Target:                    bi.Target,
		BasedOnUserAttack:         bi.BasedOnUserAttack,
		Range:                     bi.Range,
		Damage:                    ah.ConvertObjPtr(bi.Damage, convertDamage),
		ShatterRate:               bi.ShatterRate,
		Accuracy:                  convertAccuracy(bi.Accuracy),
		HitAmount:                 bi.HitAmount,
		SpecialAction:             bi.SpecialAction,
		InflictedDelay:            ah.ConvertObjPtr(bi.InflictedDelay, convertInflictedDelay),
		InflictedStatusConditions: ah.ConvertObjSlice(bi.InflictedStatusConditions, convertInflictedStatus),
		RemovedStatusConditions:   handlerinput.NamesToNamedAPIResources(E.StatusConditions, bi.RemovedStatusConditions),
		CopiedStatusConditions:    ah.ConvertObjSlice(bi.CopiedStatusConditions, convertInflictedStatus),
		StatChanges:               ah.ConvertObjSlice(bi.StatChanges, convertStatChange),
		ModifierChanges:           ah.ConvertObjSlice(bi.ModifierChanges, convertModifierChange),
	}

	for _, status := range bi.AffectedBy {
		switch status {
		case "darkness":
			battleInteraction.Darkable = true

		case "silence":
			battleInteraction.Silenceable = true

		case "reflect":
			battleInteraction.Reflectable = true
		}
	}

	return battleInteraction
}

func convertDamage(d seedtypes.Damage) apitypes.Damage {
	return apitypes.Damage{
		AttackType:      d.AttackType,
		DamageType:      d.DamageType,
		HP:              ah.ConvertObjPtr(d.HP, convertClassDamage),
		MP:              ah.ConvertObjPtr(d.MP, convertClassDamage),
		CTB:             ah.ConvertObjPtr(d.CTB, convertClassDamage),
		Critical:        d.Critical,
		CriticalPlusVal: d.CriticalPlusVal,
		IsPiercing:      d.IsPiercing,
		BreakDmgLimit:   d.BreakDmgLimit,
		Element:         handlerinput.NamePtrToNamedAPIResPtr(E.Elements, d.Element, nil),
	}
}

func convertClassDamage(ad seedtypes.ClassDamage) apitypes.ClassDamage {
	return apitypes.ClassDamage{
		Condition:      ad.Condition,
		DamageFormula:  ad.DamageFormula,
		DamageConstant: ad.DamageConstant,
	}
}

func convertAccuracy(a seedtypes.Accuracy) apitypes.Accuracy {
	return apitypes.Accuracy{
		AccSource:   a.AccSource,
		HitChance:   a.HitChance,
		AccModifier: a.AccModifier,
	}
}

func convertInflictedDelay(id seedtypes.InflictedDelay) apitypes.InflictedDelay {
	return apitypes.InflictedDelay{
		Condition:     id.Condition,
		DelayStrength: id.DelayStrength,
	}
}

func convertInflictedStatus(status seedtypes.InflictedStatus) apitypes.InflictedStatus {
	return apitypes.InflictedStatus{
		StatusCondition: handlerinput.NameToNamedAPIResource(E.StatusConditions, status.StatusCondition, nil),
		Probability:     status.Probability,
		DurationType:    status.DurationType,
		Amount:          status.Amount,
	}
}

func convertStatChange(sc seedtypes.StatChange) apitypes.StatChange {
	return apitypes.StatChange{
		Stat:            handlerinput.NameToNamedAPIResource(E.Stats, sc.StatName, nil),
		CalculationType: sc.CalculationType,
		Value:           sc.Value,
	}
}

func convertModifierChange(mc seedtypes.ModifierChange) apitypes.ModifierChange {
	return apitypes.ModifierChange{
		Modifier:        handlerinput.NameToNamedAPIResource(E.Modifiers, mc.ModifierName, nil),
		CalculationType: mc.CalculationType,
		Value:           mc.Value,
	}
}
