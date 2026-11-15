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

func (e BattleInteraction) Compare(test Test, fieldName string, got apitypes.BattleInteraction) {
	test.T.Helper()
	test.compare(fieldName+" target", e.Target, got.Target)
	test.compare(fieldName+" range", e.AbilityRange, got.Range)
	test.compare(fieldName+" hit amount", e.HitAmount, got.HitAmount)
	test.compare(fieldName+" shatter rate", e.ShatterRate, got.ShatterRate)
	test.compare(fieldName+" special action", e.SpecialAction, got.SpecialAction)
	test.compare(fieldName+" based on phys atk", e.BasedOnUserAtk, got.BasedOnUserAttack)
	test.compare(fieldName+" darkable", e.Darkable, got.Darkable)
	test.compare(fieldName+" silenceable", e.Silenceable, got.Silenceable)
	test.compare(fieldName+" reflectable", e.Reflectable, got.Reflectable)
	e.Accuracy.Compare(test, fieldName+" accuracy", got.Accuracy)
	test.compTestStructPtrs(fieldName+" damage", e.Damage, got.Damage)
	test.compTestStructPtrs(fieldName+" inflicted delay", e.InflictedDelay, got.InflictedDelay)
	test.CheckResIDsInSlice(fieldName+" inflicted status conditions", apilookups.E.StatusConditions.Endpoint, e.InflictedStatusConditions, got.InflictedStatusConditions)
	test.CheckResIDsInSlice(fieldName+" removed status conditions", apilookups.E.StatusConditions.Endpoint, e.RemovedStatusConditions, got.RemovedStatusConditions)
	test.CheckResIDsInSlice(fieldName+" copied status conditions", apilookups.E.StatusConditions.Endpoint, e.CopiedStatusConditions, got.CopiedStatusConditions)
	test.compTestStructSlices(fieldName+" stat changes", e.StatChanges, got.StatChanges)
	test.compTestStructSlices(fieldName+" mod changes", e.ModifierChanges, got.ModifierChanges)
}

type Accuracy struct {
	AccSource   string
	HitChance   *int32
	AccModifier *float32
}

func (e Accuracy) Compare(test Test, fieldName string, got apitypes.Accuracy) {
	test.compare(fieldName+" acc source", e.AccSource, got.AccSource)
	test.compare(fieldName+" hit chance", e.HitChance, got.HitChance)
	test.compare(fieldName+" acc modifier", e.AccModifier, got.AccModifier)
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

func (e Damage) Compare(test Test, fieldName string, got apitypes.Damage) {
	test.compare(fieldName+" attack type", string(e.AttackType), got.AttackType)
	test.compare(fieldName+" damage type", string(e.DamageType), got.DamageType)
	test.compTestStructPtrs(fieldName+" hp", e.Hp, got.HP)
	test.compTestStructPtrs(fieldName+" mp", e.Mp, got.MP)
	test.compTestStructPtrs(fieldName+" ctb", e.Ctb, got.CTB)
	test.compare(fieldName+" critical", e.Critical, got.Critical)
	test.compare(fieldName+" criticalPlusVal", e.CriticalPlusVal, got.CriticalPlusVal)
	test.compare(fieldName+" is piercing", e.IsPiercing, got.IsPiercing)
	test.compare(fieldName+" break dmg lmt", e.BreakDmgLmt, got.BreakDmgLimit)
	test.compIDAPIResourcePtrs(fieldName+" element", apilookups.E.Elements.Endpoint, e.Element, got.Element)
}

type ClassDamage struct {
	DamageFormula  database.DamageFormula
	DamageConstant int32
}

func (e ClassDamage) Compare(test Test, fieldName string, got apitypes.ClassDamage) {
	test.compare(fieldName+" - damage formula", string(e.DamageFormula), got.DamageFormula)
	test.compare(fieldName+" - damage constant", e.DamageConstant, got.DamageConstant)
}

type InflictedDelay struct {
	DelayStrength string
}

func (e InflictedDelay) Compare(test Test, fieldName string, got apitypes.InflictedDelay) {
	test.compare(fieldName+" delay strength", e.DelayStrength, got.DelayStrength)
}

type StatChange struct {
	Stat            int32
	CalculationType string
	Value           float32
}

func (e StatChange) Compare(test Test, fieldName string, got apitypes.StatChange) {
	test.compIDAPIResource(fieldName+" - sc stat", apilookups.E.Stats.Endpoint, e.Stat, got.Stat)
	test.compare(fieldName+" - calculation type", e.CalculationType, got.CalculationType)
	test.compare(fieldName+" - value", e.Value, got.Value)
}

type ModChange struct {
	Modifier        int32
	CalculationType string
	Value           float32
}

func (e ModChange) Compare(test Test, fieldName string, got apitypes.ModifierChange) {
	test.compIDAPIResource(fieldName+" - mc modifier", apilookups.E.Modifiers.Endpoint, e.Modifier, got.Modifier)
	test.compare(fieldName+" - calculation type", e.CalculationType, got.CalculationType)
	test.compare(fieldName+" - value", e.Value, got.Value)
}
