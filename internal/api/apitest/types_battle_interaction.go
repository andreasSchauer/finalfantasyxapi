package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type BattleInteraction struct {
	target                    string
	abilityRange              *int32
	hitAmount                 int32
	shatterRate               int32
	specialAction             *string
	basedOnUserAtk            bool
	darkable                  bool
	silenceable               bool
	reflectable               bool
	accuracy                  Accuracy
	damage                    *Damage
	inflictedDelay            *InflictedDelay
	inflictedStatusConditions []int32
	removedStatusConditions   []int32
	copiedStatusConditions    []int32
	statChanges               []StatChange
	modifierChanges           []ModChange
}

func compareBattleInteractions(test test, fieldName string, exp BattleInteraction, got apitypes.BattleInteraction) {
	test.t.Helper()
	test.compare(fieldName+" target", exp.target, got.Target)
	test.compare(fieldName+" range", exp.abilityRange, got.Range)
	test.compare(fieldName+" hit amount", exp.hitAmount, got.HitAmount)
	test.compare(fieldName+" shatter rate", exp.shatterRate, got.ShatterRate)
	test.compare(fieldName+" special action", exp.specialAction, got.SpecialAction)
	test.compare(fieldName+" based on phys atk", exp.basedOnUserAtk, got.BasedOnUserAttack)
	test.compare(fieldName+" darkable", exp.darkable, got.Darkable)
	test.compare(fieldName+" silenceable", exp.silenceable, got.Silenceable)
	test.compare(fieldName+" reflectable", exp.reflectable, got.Reflectable)
	compareAccuracies(test, fieldName+" accuracy", exp.accuracy, got.Accuracy)
	test.compTestStructPtrs(fieldName+" damage", exp.damage, got.Damage, compareDamages)
	test.compTestStructPtrs(fieldName+" inflicted delay", exp.inflictedDelay, got.InflictedDelay, compareInflictedDelays)
	test.checkResIDsInSlice(fieldName+" inflicted status conditions", apilookups.E.StatusConditions.Endpoint, exp.inflictedStatusConditions, got.InflictedStatusConditions)
	test.checkResIDsInSlice(fieldName+" removed status conditions", apilookups.E.StatusConditions.Endpoint, exp.removedStatusConditions, got.RemovedStatusConditions)
	test.checkResIDsInSlice(fieldName+" copied status conditions", apilookups.E.StatusConditions.Endpoint, exp.copiedStatusConditions, got.CopiedStatusConditions)
	test.compTestStructSlices(fieldName+" stat changes", exp.statChanges, got.StatChanges, compareStatChanges)
	test.compTestStructSlices(fieldName+" mod changes", exp.modifierChanges, got.ModifierChanges, compareModChanges)
}

type Accuracy struct {
	accSource   string
	hitChance   *int32
	accModifier *float32
}

func compareAccuracies(test test, fieldName string, exp Accuracy, got apitypes.Accuracy) {
	test.compare(fieldName+" acc source", exp.accSource, got.AccSource)
	test.compare(fieldName+" hit chance", exp.hitChance, got.HitChance)
	test.compare(fieldName+" acc modifier", exp.accModifier, got.AccModifier)
}

type Damage struct {
	attackType      database.AttackType
	damageType      database.DamageType
	hp              *ClassDamage
	mp              *ClassDamage
	ctb             *ClassDamage
	critical        *string
	criticalPlusVal *int32
	isPiercing      bool
	breakDmgLmt     *string
	element         *int32
}

func compareDamages(test test, fieldName string, exp Damage, got apitypes.Damage) {
	test.compare(fieldName+" attack type", string(exp.attackType), got.AttackType)
	test.compare(fieldName+" damage type", string(exp.damageType), got.DamageType)
	test.compTestStructPtrs(fieldName+" hp", exp.hp, got.HP, compareClassDamages)
	test.compTestStructPtrs(fieldName+" mp", exp.mp, got.MP, compareClassDamages)
	test.compTestStructPtrs(fieldName+" ctb", exp.ctb, got.CTB, compareClassDamages)
	test.compare(fieldName+" critical", exp.critical, got.Critical)
	test.compare(fieldName+" criticalPlusVal", exp.criticalPlusVal, got.CriticalPlusVal)
	test.compare(fieldName+" is piercing", exp.isPiercing, got.IsPiercing)
	test.compare(fieldName+" break dmg lmt", exp.breakDmgLmt, got.BreakDmgLimit)
	test.compIDAPIResourcePtrs(fieldName+" element", apilookups.E.Elements.Endpoint, exp.element, got.Element)
}

type ClassDamage struct {
	damageFormula  database.DamageFormula
	damageConstant int32
}

func compareClassDamages(test test, fieldName string, exp ClassDamage, got apitypes.ClassDamage) {
	test.compare(fieldName+" - damage formula", string(exp.damageFormula), got.DamageFormula)
	test.compare(fieldName+" - damage constant", exp.damageConstant, got.DamageConstant)
}

type InflictedDelay struct {
	delayStrength string
}

func compareInflictedDelays(test test, fieldName string, exp InflictedDelay, got apitypes.InflictedDelay) {
	test.compare(fieldName+" delay strength", exp.delayStrength, got.DelayStrength)
}

type StatChange struct {
	stat            int32
	calculationType string
	value           float32
}

func compareStatChanges(test test, fieldName string, exp StatChange, got apitypes.StatChange) {
	test.compIDAPIResource(fieldName+" - sc stat", apilookups.E.Stats.Endpoint, exp.stat, got.Stat)
	test.compare(fieldName+" - calculation type", exp.calculationType, got.CalculationType)
	test.compare(fieldName+" - value", exp.value, got.Value)
}

type ModChange struct {
	modifier        int32
	calculationType string
	value           float32
}

func compareModChanges(test test, fieldName string, exp ModChange, got apitypes.ModifierChange) {
	test.compIDAPIResource(fieldName+" - mc modifier", apilookups.E.Modifiers.Endpoint, exp.modifier, got.Modifier)
	test.compare(fieldName+" - calculation type", exp.calculationType, got.CalculationType)
	test.compare(fieldName+" - value", exp.value, got.Value)
}
