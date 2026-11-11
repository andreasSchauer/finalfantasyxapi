package apiconvert

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func BattleInteraction(bi seedtypes.BattleInteraction) apitypes.BattleInteraction {
	battleInteraction := apitypes.BattleInteraction{
		Target:                    bi.Target,
		BasedOnUserAttack:         bi.BasedOnUserAttack,
		Range:                     bi.Range,
		Damage:                    ah.ConvertObjPtr(bi.Damage, convertDamage),
		ShatterRate:               bi.ShatterRate,
		Accuracy:                  Accuracy(bi.Accuracy),
		HitAmount:                 bi.HitAmount,
		SpecialAction:             bi.SpecialAction,
		InflictedDelay:            ah.ConvertObjPtr(bi.InflictedDelay, convertInflictedDelay),
		InflictedStatusConditions: ah.ConvertObjSlice(bi.InflictedStatusConditions, InflictedStatus),
		RemovedStatusConditions:   handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, bi.RemovedStatusConditions),
		CopiedStatusConditions:    ah.ConvertObjSlice(bi.CopiedStatusConditions, InflictedStatus),
		StatChanges:               ah.ConvertObjSlice(bi.StatChanges, StatChange),
		ModifierChanges:           ah.ConvertObjSlice(bi.ModifierChanges, ModifierChange),
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
		Element:         handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Elements, d.Element, nil),
	}
}

func convertClassDamage(ad seedtypes.ClassDamage) apitypes.ClassDamage {
	return apitypes.ClassDamage{
		Condition:      ad.Condition,
		DamageFormula:  ad.DamageFormula,
		DamageConstant: ad.DamageConstant,
	}
}

func Accuracy(a seedtypes.Accuracy) apitypes.Accuracy {
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

func InflictedStatus(status seedtypes.InflictedStatus) apitypes.InflictedStatus {
	return apitypes.InflictedStatus{
		StatusCondition: handlerinput.NameToNamedAPIResource(apilookups.E.StatusConditions, status.StatusCondition, nil),
		Probability:     status.Probability,
		DurationType:    status.DurationType,
		Amount:          status.Amount,
	}
}

func StatChange(sc seedtypes.StatChange) apitypes.StatChange {
	return apitypes.StatChange{
		Stat:            handlerinput.NameToNamedAPIResource(apilookups.E.Stats, sc.StatName, nil),
		CalculationType: sc.CalculationType,
		Value:           sc.Value,
	}
}

func ModifierChange(mc seedtypes.ModifierChange) apitypes.ModifierChange {
	return apitypes.ModifierChange{
		Modifier:        handlerinput.NameToNamedAPIResource(apilookups.E.Modifiers, mc.ModifierName, nil),
		CalculationType: mc.CalculationType,
		Value:           mc.Value,
	}
}
