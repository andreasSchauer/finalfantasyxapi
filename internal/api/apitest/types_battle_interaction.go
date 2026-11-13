package apitest

import (
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

type Accuracy struct {
	accSource   string
	hitChance   *int32
	accModifier *float32
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

type ClassDamage struct {
	damageFormula  database.DamageFormula
	damageConstant int32
}

type InflictedDelay struct {
	delayStrength string
}

type StatChange struct {
	stat            int32
	calculationType string
	value           float32
}

type ModChange struct {
	modifier        int32
	calculationType string
	value           float32
}
