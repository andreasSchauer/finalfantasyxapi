package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type BattleInteraction struct {
	Target                    string
	AbilityRange              *int32
	HitAmount                 int32
	ShatterRate               int32
	SpecialAction             *string
	BasedOnUserAtk            bool
	Darkable                  bool
	Silenceable               bool
	Reflectable               bool
	Accuracy                  Accuracy
	Damage                    *Damage
	InflictedDelay            *InflictedDelay
	InflictedStatusConditions []int32
	RemovedStatusConditions   []int32
	CopiedStatusConditions    []int32
	StatChanges               []StatChange
	ModifierChanges           []ModChange
}

func (e BattleInteraction) Compare(test Test, fn string, got apitypes.BattleInteraction) {
	test.T.Helper()
	test.compare(childFn(fn, "target"), e.Target, got.Target)
	test.compare(childFn(fn, "range"), e.AbilityRange, got.Range)
	test.compare(childFn(fn, "hit amount"), e.HitAmount, got.HitAmount)
	test.compare(childFn(fn, "shatter rate"), e.ShatterRate, got.ShatterRate)
	test.compare(childFn(fn, "special action"), e.SpecialAction, got.SpecialAction)
	test.compare(childFn(fn, "based on phys atk"), e.BasedOnUserAtk, got.BasedOnUserAttack)
	test.compare(childFn(fn, "darkable"), e.Darkable, got.Darkable)
	test.compare(childFn(fn, "silenceable"), e.Silenceable, got.Silenceable)
	test.compare(childFn(fn, "reflectable"), e.Reflectable, got.Reflectable)
	e.Accuracy.Compare(test, childFn(fn, "accuracy"), got.Accuracy)
	test.compTestStructPtrs(childFn(fn, "damage"), e.Damage, got.Damage)
	test.compTestStructPtrs(childFn(fn, "inflicted delay"), e.InflictedDelay, got.InflictedDelay)
	test.CheckResIDsInSlice(childFn(fn, "inflicted status conditions"), apilookups.E.StatusConditions.Endpoint, e.InflictedStatusConditions, got.InflictedStatusConditions)
	test.CheckResIDsInSlice(childFn(fn, "removed status conditions"), apilookups.E.StatusConditions.Endpoint, e.RemovedStatusConditions, got.RemovedStatusConditions)
	test.CheckResIDsInSlice(childFn(fn, "copied status conditions"), apilookups.E.StatusConditions.Endpoint, e.CopiedStatusConditions, got.CopiedStatusConditions)
	test.compTestStructSlices(childFn(fn, "stat changes"), e.StatChanges, got.StatChanges)
	test.compTestStructSlices(childFn(fn, "mod changes"), e.ModifierChanges, got.ModifierChanges)
}

type Accuracy struct {
	AccSource   string
	HitChance   *int32
	AccModifier *float32
}

func (e Accuracy) Compare(test Test, fn string, got apitypes.Accuracy) {
	test.compare(childFn(fn, "acc source"), e.AccSource, got.AccSource)
	test.compare(childFn(fn, "hit chance"), e.HitChance, got.HitChance)
	test.compare(childFn(fn, "acc modifier"), e.AccModifier, got.AccModifier)
}

type Damage struct {
	AttackType      database.AttackType
	DamageType      database.DamageType
	Hp              *ClassDamage
	Mp              *ClassDamage
	Ctb             *ClassDamage
	Critical        *string
	CriticalPlusVal *int32
	IsPiercing      bool
	BreakDmgLmt     *string
	Element         *int32
}

func (e Damage) Compare(test Test, fn string, got apitypes.Damage) {
	test.compare(childFn(fn, "attack type"), string(e.AttackType), got.AttackType)
	test.compare(childFn(fn, "damage type"), string(e.DamageType), got.DamageType)
	test.compTestStructPtrs(childFn(fn, "hp"), e.Hp, got.HP)
	test.compTestStructPtrs(childFn(fn, "mp"), e.Mp, got.MP)
	test.compTestStructPtrs(childFn(fn, "ctb"), e.Ctb, got.CTB)
	test.compare(childFn(fn, "critical"), e.Critical, got.Critical)
	test.compare(childFn(fn, "criticalPlusVal"), e.CriticalPlusVal, got.CriticalPlusVal)
	test.compare(childFn(fn, "is piercing"), e.IsPiercing, got.IsPiercing)
	test.compare(childFn(fn, "break dmg lmt"), e.BreakDmgLmt, got.BreakDmgLimit)
	test.compIDAPIResourcePtrs(childFn(fn, "element"), apilookups.E.Elements.Endpoint, e.Element, got.Element)
}

type ClassDamage struct {
	DamageFormula  database.DamageFormula
	DamageConstant int32
}

func (e ClassDamage) Compare(test Test, fn string, got apitypes.ClassDamage) {
	test.compare(childFn(fn, "damage formula"), string(e.DamageFormula), got.DamageFormula)
	test.compare(childFn(fn, "damage constant"), e.DamageConstant, got.DamageConstant)
}

type InflictedDelay struct {
	DelayStrength string
}

func (e InflictedDelay) Compare(test Test, fn string, got apitypes.InflictedDelay) {
	test.compare(childFn(fn, "delay strength"), e.DelayStrength, got.DelayStrength)
}

type StatChange struct {
	Stat            int32
	CalculationType string
	Value           float32
}

func (e StatChange) Compare(test Test, fn string, got apitypes.StatChange) {
	test.compIDAPIResource(childFn(fn, "sc stat"), apilookups.E.Stats.Endpoint, e.Stat, got.Stat)
	test.compare(childFn(fn, "calculation type"), e.CalculationType, got.CalculationType)
	test.compare(childFn(fn, "value"), e.Value, got.Value)
}

type ModChange struct {
	Modifier        int32
	CalculationType string
	Value           float32
}

func (e ModChange) Compare(test Test, fn string, got apitypes.ModifierChange) {
	test.compIDAPIResource(childFn(fn, "mc modifier"), apilookups.E.Modifiers.Endpoint, e.Modifier, got.Modifier)
	test.compare(childFn(fn, "calculation type"), e.CalculationType, got.CalculationType)
	test.compare(childFn(fn, "value"), e.Value, got.Value)
}
