package api

import h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"


type CritChanceParams struct {
	Ability		AbilityCritChance	`json:"ability"`
	User		UserCritChance		`json:"user"`
	Target		TargetCritChance	`json:"target"`
}

// ======== ability ========

type AbilityCritChance struct {
	Lookup			*AbilityLookupCritChance	`json:"lookup,omitempty"`
	Custom			*AbilityCustomCritChance	`json:"custom,omitempty"`
}

type AbilityLookupCritChance struct {
	AbilityID		int32						`json:"ability_id"`
	Interaction		*int32						`json:"interaction"`
	Overrides		*AbilityOverridesCritChance `json:"overrides,omitempty"`
}

type AbilityOverridesCritChance struct {
	Critical		*string		`json:"critical,omitempty"`
	CriticalPlusVal	*int32		`json:"critical_plus_val,omitempty"`
}

type AbilityCustomCritChance struct {
	Critical		string		`json:"critical"`
	CriticalPlusVal	int32		`json:"critical_plus_val"`
}

// ======== user ========

type UserCritChance struct {
	Lookup			*UserLookupCritChance	`json:"lookup,omitempty"`
	Custom			*UserCustomCritChance	`json:"custom,omitempty"`
	LuckStacks		int32					`json:"luck_stacks"`
}

type UserLookupCritChance struct {
	MonsterID		int32						`json:"monster_id"`
	Overrides		*UserOverridesCritChance	`json:"overrides,omitempty"`
}

type UserOverridesCritChance struct {
	Luck			*int32		`json:"luck,omitempty"`
	CriticalPlusVal	*int32		`json:"critical_plus_val,omitempty"`
}

type UserCustomCritChance struct {
	Luck			int32		`json:"luck"`
	CriticalPlusVal	int32		`json:"critical_plus_val"`
}

// ======== target ========

type TargetCritChance struct {
	Lookup			*TargetLookupCritChance		`json:"lookup,omitempty"`
	Custom			*TargetCustomCritChance		`json:"custom,omitempty"`
	JinxStacks		int32						`json:"jinx_stacks"`
}

type TargetLookupCritChance struct {
	MonsterID		int32						`json:"monster_id"`
	Overrides		*TargetOverridesCritChance	`json:"overrides,omitempty"`
}

type TargetOverridesCritChance struct {
	Luck			*int32		`json:"luck,omitempty"`
}

type TargetCustomCritChance struct {
	Luck			int32		`json:"luck"`
}

func getCritChanceParamsDoc(cfg *Config) ParamsDoc {
	return ParamsDoc{
		GeneralRules: h.GetStrPtr(""),
		Fields: []FieldDoc{
			{
				Field: 		 PfnAbility,
				Type: 		 "obj (abilityCritChance)",
				Required: 	 true,
				Description: "Define the ability that is being used for the attack.",
				ChildProps:  cfg.getFieldDocAbilityCritChance(),
			},
			{
				Field: 		 PfnUser,
				Type: 		 "obj (userCritChance)",
				Required: 	 true,
				Description: "Define the unit that uses the ability.",
				ChildProps:  cfg.getFieldDocUserCritChance(),
			},
			{
				Field: 		 PfnTarget,
				Type: 		 "obj (targetCritChance)",
				Required: 	 true,
				Description: "Define the unit that gets hit by the ability.",
				ChildProps:  cfg.getFieldDocTargetCritChance(),
			},
		},
	}
}

func (cfg *Config) getFieldDocAbilityCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field: 		   PfnLookup,
			Type: 		   "obj (abilityLookupCritChance)",
			RequiredOr:    []FieldName{PfnCustom},
			Description:   "Define the ability via lookup. You can override calculation-relevant values as well.",
			ChildProps:    cfg.getFieldDocAbilityLookupCritChance(),
		},
		{
			Field: 		   PfnCustom,
			Type: 		   "obj (abilityCustomCritChance)",
			RequiredOr:    []FieldName{PfnLookup},
			Description:   "Define a custom ability.",
			ChildProps:    cfg.getFieldDocAbilityCustomCritChance(),
		},
	}
}

func (cfg *Config) getFieldDocAbilityLookupCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:         PfnAbilityID,
			Type:          "int (id: ability)",
			Required: 	   true,
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(cfg.L.Abilities))),
			Description:   "To define an ability, you can specify an ability ID and look up the corresponding ability. Its calculation-relevant values will be applied automatically. These values can be overridden however, through the use of the 'overrides' field.",
		},
		{
			Field:         PfnInteraction,
			Type:          "int (idx +1: ability battle interactions)",
			RequiresAll:   []FieldName{PfnAbilityID},
			Description:   "The calculator can only take one battle interaction into account. If an ability has more than one battle interaction, this field lets you specify the battle interaction to perform the calculation on.",
		},
		{
			Field: 		   PfnOverrides,
			Type: 		   "obj (abilityOverridesCritChance)",
			Description:   "Override the looked up ability's calculation-relevant values.",
			ChildProps:    cfg.getFieldDocAbilityOverridesCritChance(),
		},
	}
}

func (cfg *Config) getFieldDocAbilityOverridesCritChance() []FieldDoc {
	return []FieldDoc{
		
		{
			Field:       PfnCritical,
			Type:        "string (enum: criticalType)",
			EnumValues:  cfg.T.CriticalType.Strings(),
			Description: "Override the critical-type of the looked up ability.",
		},
		{
			Field:       PfnCriticalPlusVal,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(100),
			Description: "Override the looked up ability's own critical+ value. This value will be used when an abilitiy's crit-chance relies on it ('crit-ability'). The critical+ value will be added to the crit-chance, so a critical+ value of 6 will increase the chance by 6%.",
		},
	}
}

func (cfg *Config) getFieldDocAbilityCustomCritChance() []FieldDoc {
	return []FieldDoc{
		
		{
			Field:       PfnCritical,
			Type:        "string (enum: criticalType)",
			Required: 	 true,
			EnumValues:  cfg.T.CriticalType.Strings(),
			Description: "Specify the critical-type of the ability.",
		},
		{
			Field:       PfnCriticalPlusVal,
			Type:        "int",
			DefaultVal:  int32(0),
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(100),
			Description: "Specify the ability's own critical+ value. This value will be used when an abilitiy's crit-chance relies on it ('crit-ability'). The critical+ value will be added to the crit-chance, so a critical+ value of 6 will increase the chance by 6%.",
		},
	}
}

func (cfg *Config) getFieldDocUserCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field: 		   PfnLookup,
			Type: 		   "obj (userLookupCritChance)",
			RequiredOr:    []FieldName{PfnCustom},
			Description:   "Define the ability's user via monster-lookup. You can override calculation-relevant values as well.",
			ChildProps:    cfg.getFieldDocUserLookupCritChance(),
		},
		{
			Field: 		   PfnCustom,
			Type: 		   "obj (userCustomCritChance)",
			RequiredOr:    []FieldName{PfnLookup},
			Description:   "Define a custom ability user.",
			ChildProps:    cfg.getFieldDocUserCustomCritChance(),
		},
		{
			Field:       PfnLuckStacks,
			Type:        "int",
			DefaultVal:  int32(0),
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(5),
			Description: "Specify how many stacks of 'luck' were applied to the user. Each stack raises the chance of a crit landing by 1%.",
		},
	}
}

func (cfg *Config) getFieldDocUserLookupCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:         PfnMonsterID,
			Type:          "int (id: monster)",
			Required:      true,
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(cfg.L.Monsters))),
			Description:   "To define a user, you can specify a monster ID and look up the corresponding monster. Its luck stat, as well as its innate critical+ value will be applied automatically. These values can be overridden however, through the use of the 'overrides field. Note: I haven't found out the innate critical+ values of monsters so far, so for now, I'm assuming that it's the same as a monster's equipment drops would have (usually +3). But in general, there's only a handful of abilities that make use of a monster's critical+ value. Most of the time, enemy abilities rely on the ability's own critical+ value.",
		},
		{
			Field: 		   PfnOverrides,
			Type: 		   "obj (userOverridesCritChance)",
			Description:   "Override the monster's calculation-relevant values.",
			ChildProps:    cfg.getFieldDocUserOverridesCritChance(),
		},
	}
}

func (cfg *Config) getFieldDocUserOverridesCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:       PfnLuck,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "Override the luck stat of the looked up monster (user).",
		},
		{
			Field:       PfnCriticalPlusVal,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(100),
			Description: "Override the looked up monster's own critical+ value. This value will be used when an abilitiy's crit-chance relies on the user rather than the ability itself ('crit-equip'). The critical+ value will be added to the crit-chance, so a critical+ value of 6 will increase the chance by 6%. Note: I haven't found out the innate critical+ values of monsters so far, so for now, I'm assuming that it's the same as a monster's equipment drops would have (usually +3). But in general, there's only a handful of abilities that make use of a monster's critical+ value. Most of the time, enemy abilities rely on the ability's critical+ value.",
		},
	}
}

func (cfg *Config) getFieldDocUserCustomCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:       PfnLuck,
			Type:        "int",
			Required: 	 true,
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The user's luck stat.",
		},
		{
			Field:       PfnCriticalPlusVal,
			Type:        "int",
			DefaultVal:  int32(6),
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(100),
			Description: "Specify the user's own critical+ value. This value will be used when an abilitiy's crit-chance relies on the user rather than the ability itself ('crit-equip'). The critical+ value will be added to the crit-chance, so a critical+ value of 6 will increase the chance by 6%. Each piece of equipment (both weapon and armor) has its own critical+ value. Usually this value is 3, so for most of the game, a character will have a critical+ value of 6. Some arena creations and dark aeons can drop equipment with critical+ = 6, celestial weapons also have critical+ = 6, and penance drops equipment with critical+ = 10. So a character's maximum is +16.",
		},
	}
}

func (cfg *Config) getFieldDocTargetCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field: 		   PfnLookup,
			Type: 		   "obj (targetLookupCritChance)",
			RequiredOr:    []FieldName{PfnCustom},
			Description:   "Define the ability's target via monster-lookup. You can override calculation-relevant values as well.",
			ChildProps:    cfg.getFieldDocTargetLookupCritChance(),
		},
		{
			Field: 		   PfnCustom,
			Type: 		   "obj (targetCustomCritChance)",
			RequiredOr:    []FieldName{PfnLookup},
			Description:   "Define a custom target.",
			ChildProps:    cfg.getFieldDocTargetCustomCritChance(),
		},
		{
			Field:       PfnJinxStacks,
			Type:        "int",
			DefaultVal:  int32(0),
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(5),
			Description: "Specify how many stacks of 'jinx' were applied to the target. Each stack raises the chance of a crit landing by 1%.",
		},
	}
}

func (cfg *Config) getFieldDocTargetLookupCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:         PfnMonsterID,
			Type:          "int (id: monster)",
			RequiredOr:    []FieldName{PfnLuck},
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(cfg.L.Monsters))),
			Description:   "To define a target, you can specify a monster ID and look up the corresponding monster. Its luck stat will be applied automatically. This value can be overridden however, through the use of the 'luck' field.",
		},
		{
			Field: 		   PfnOverrides,
			Type: 		   "obj (targetOverridesCritChance)",
			Description:   "Override the monster's calculation-relevant values.",
			ChildProps:    cfg.getFieldDocTargetOverridesCritChance(),
		},
	}
}

func (cfg *Config) getFieldDocTargetOverridesCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:       PfnLuck,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "Override the luck stat of the looked up monster (target).",
		},
	}
}

func (cfg *Config) getFieldDocTargetCustomCritChance() []FieldDoc {
	return []FieldDoc{
		{
			Field:       PfnLuck,
			Type:        "int",
			Required:    true,
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The target's luck stat.",
		},
	}
}