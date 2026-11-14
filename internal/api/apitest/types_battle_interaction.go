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

func (e BattleInteraction) Compare(test test, fieldName string, got apitypes.BattleInteraction) {
	test.t.Helper()
	test.compare(fieldName+" target", e.target, got.Target)
	test.compare(fieldName+" range", e.abilityRange, got.Range)
	test.compare(fieldName+" hit amount", e.hitAmount, got.HitAmount)
	test.compare(fieldName+" shatter rate", e.shatterRate, got.ShatterRate)
	test.compare(fieldName+" special action", e.specialAction, got.SpecialAction)
	test.compare(fieldName+" based on phys atk", e.basedOnUserAtk, got.BasedOnUserAttack)
	test.compare(fieldName+" darkable", e.darkable, got.Darkable)
	test.compare(fieldName+" silenceable", e.silenceable, got.Silenceable)
	test.compare(fieldName+" reflectable", e.reflectable, got.Reflectable)
	e.accuracy.Compare(test, fieldName+" accuracy", got.Accuracy)
	test.compTestStructPtrs(fieldName+" damage", e.damage, got.Damage)
	test.compTestStructPtrs(fieldName+" inflicted delay", e.inflictedDelay, got.InflictedDelay)
	test.checkResIDsInSlice(fieldName+" inflicted status conditions", apilookups.E.StatusConditions.Endpoint, e.inflictedStatusConditions, got.InflictedStatusConditions)
	test.checkResIDsInSlice(fieldName+" removed status conditions", apilookups.E.StatusConditions.Endpoint, e.removedStatusConditions, got.RemovedStatusConditions)
	test.checkResIDsInSlice(fieldName+" copied status conditions", apilookups.E.StatusConditions.Endpoint, e.copiedStatusConditions, got.CopiedStatusConditions)
	test.compTestStructSlices(fieldName+" stat changes", e.statChanges, got.StatChanges)
	test.compTestStructSlices(fieldName+" mod changes", e.modifierChanges, got.ModifierChanges)
}

type Accuracy struct {
	accSource   string
	hitChance   *int32
	accModifier *float32
}

func (e Accuracy) Compare(test test, fieldName string, got apitypes.Accuracy) {
	test.compare(fieldName+" acc source", e.accSource, got.AccSource)
	test.compare(fieldName+" hit chance", e.hitChance, got.HitChance)
	test.compare(fieldName+" acc modifier", e.accModifier, got.AccModifier)
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

func (e Damage) Compare(test test, fieldName string, got apitypes.Damage) {
	test.compare(fieldName+" attack type", string(e.attackType), got.AttackType)
	test.compare(fieldName+" damage type", string(e.damageType), got.DamageType)
	test.compTestStructPtrs(fieldName+" hp", e.hp, got.HP)
	test.compTestStructPtrs(fieldName+" mp", e.mp, got.MP)
	test.compTestStructPtrs(fieldName+" ctb", e.ctb, got.CTB)
	test.compare(fieldName+" critical", e.critical, got.Critical)
	test.compare(fieldName+" criticalPlusVal", e.criticalPlusVal, got.CriticalPlusVal)
	test.compare(fieldName+" is piercing", e.isPiercing, got.IsPiercing)
	test.compare(fieldName+" break dmg lmt", e.breakDmgLmt, got.BreakDmgLimit)
	test.compIDAPIResourcePtrs(fieldName+" element", apilookups.E.Elements.Endpoint, e.element, got.Element)
}

type ClassDamage struct {
	damageFormula  database.DamageFormula
	damageConstant int32
}

func (e ClassDamage) Compare(test test, fieldName string, got apitypes.ClassDamage) {
	test.compare(fieldName+" - damage formula", string(e.damageFormula), got.DamageFormula)
	test.compare(fieldName+" - damage constant", e.damageConstant, got.DamageConstant)
}

type InflictedDelay struct {
	delayStrength string
}

func (e InflictedDelay) Compare(test test, fieldName string, got apitypes.InflictedDelay) {
	test.compare(fieldName+" delay strength", e.delayStrength, got.DelayStrength)
}

type StatChange struct {
	stat            int32
	calculationType string
	value           float32
}

func (e StatChange) Compare(test test, fieldName string, got apitypes.StatChange) {
	test.compIDAPIResource(fieldName+" - sc stat", apilookups.E.Stats.Endpoint, e.stat, got.Stat)
	test.compare(fieldName+" - calculation type", e.calculationType, got.CalculationType)
	test.compare(fieldName+" - value", e.value, got.Value)
}

type ModChange struct {
	modifier        int32
	calculationType string
	value           float32
}

func (e ModChange) Compare(test test, fieldName string, got apitypes.ModifierChange) {
	test.compIDAPIResource(fieldName+" - mc modifier", apilookups.E.Modifiers.Endpoint, e.modifier, got.Modifier)
	test.compare(fieldName+" - calculation type", e.calculationType, got.CalculationType)
	test.compare(fieldName+" - value", e.value, got.Value)
}
