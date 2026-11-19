package servicetypes


type CritChanceParams struct {
	Ability AbilityCritChance `json:"ability"`
	User    UserCritChance    `json:"user"`
	Target  TargetCritChance  `json:"target"`
}

// ======== ability ========

type AbilityCritChance struct {
	Lookup *AbilityLookupCritChance `json:"lookup,omitempty"`
	Custom *AbilityCustomCritChance `json:"custom,omitempty"`
}

type AbilityLookupCritChance struct {
	AbilityID   int32                       `json:"ability_id"`
	Interaction *int32                      `json:"interaction"`
	Overrides   *AbilityOverridesCritChance `json:"overrides,omitempty"`
}

type AbilityOverridesCritChance struct {
	Critical        *string `json:"critical,omitempty"`
	CriticalPlusVal *int32  `json:"critical_plus_val,omitempty"`
}

type AbilityCustomCritChance struct {
	Critical        string `json:"critical"`
	CriticalPlusVal int32  `json:"critical_plus_val"`
}

// ======== user ========

type UserCritChance struct {
	Lookup     *UserLookupCritChance `json:"lookup,omitempty"`
	Custom     *UserCustomCritChance `json:"custom,omitempty"`
	LuckStacks int32                 `json:"luck_stacks"`
}

type UserLookupCritChance struct {
	MonsterID int32                    `json:"monster_id"`
	Overrides *UserOverridesCritChance `json:"overrides,omitempty"`
}

type UserOverridesCritChance struct {
	Luck            *int32 `json:"luck,omitempty"`
	CriticalPlusVal *int32 `json:"critical_plus_val,omitempty"`
}

type UserCustomCritChance struct {
	Luck            int32 `json:"luck"`
	CriticalPlusVal int32 `json:"critical_plus_val"`
}

// ======== target ========

type TargetCritChance struct {
	Lookup     *TargetLookupCritChance `json:"lookup,omitempty"`
	Custom     *TargetCustomCritChance `json:"custom,omitempty"`
	JinxStacks int32                   `json:"jinx_stacks"`
}

type TargetLookupCritChance struct {
	MonsterID int32                      `json:"monster_id"`
	Overrides *TargetOverridesCritChance `json:"overrides,omitempty"`
}

type TargetOverridesCritChance struct {
	Luck *int32 `json:"luck,omitempty"`
}

type TargetCustomCritChance struct {
	Luck int32 `json:"luck"`
}