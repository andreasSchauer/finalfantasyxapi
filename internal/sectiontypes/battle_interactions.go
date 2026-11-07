package sectiontypes

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

type DamageSimple struct {
	AttackType    string  `json:"attack_type"`
	DamageType    string  `json:"damage_type"`
	HP            *string `json:"hp"`
	MP            *string `json:"mp"`
	CTB           *string `json:"ctb"`
	CanCritical   bool    `json:"can_critical"`
	BreakDmgLimit *string `json:"break_dmg_limit,omitempty"`
	Element       *string `json:"element,omitempty"`
}
