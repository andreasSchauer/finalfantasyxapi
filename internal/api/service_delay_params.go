package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type DelayParams struct {
	IgnImmunities 	bool            `json:"ign_immunities"`
	Delay			CustomDelay		`json:"delay"`
	Target			DelayTarget		`json:"target"`
}

func (p DelayParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getDelayParamsDoc()
}

type CustomDelay struct {
	DelayType		string		`json:"delay_type"`
	AttackType		string		`json:"attack_type"`
	Strength		*string		`json:"strength"`
	DelayConstant	*int32		`json:"delay_constant"`
}

type DelayTarget struct {
	MonsterID       *int32  	`json:"monster_id"`
	AltState    	*int32  	`json:"alt_state"`
	Agility     	*int32  	`json:"agility"`
	Status      	*string 	`json:"status"`
	RemainingTicks	*int32		`json:"remaining_ticks"`
}

func (cfg *Config) getDelayParamsDoc() ParamsDoc {
	return ParamsDoc{
		GeneralRules: h.GetStrPtr(""),
		Fields: []FieldDoc{
			{
				Field:       pfnIgnImmunities,
				Type:        "bool",
				Description: "If this field is true, a looked up target monster's immunity to delay will be ignored.",
			},
			{
				Field: 			pfnDelay,
				Type: 			"object (CustomDelay)",
				Required: 		true,
				Description: 	"Defines the delay that is dealt to the target.",
				ChildProps: 	cfg.getFieldDocCustomDelay(),
			},
			{
				Field: pfnTarget,
				Type: "object (DelayTarget)",
				Required: true,
				Description: "Specifies the target of the delay.",
				ChildProps: cfg.getFieldDocDelayTarget(),
			},
		},
	}
}


func (cfg *Config) getFieldDocCustomDelay() []FieldDoc {
	return []FieldDoc{
		{
			Field: 			pfnDelayType,
			Type: 			"string (enum: delayType)",	
			EnumValues: 	cfg.t.DelayType.Strings(),
			Required: 		true,
			Description: 	"Define the type of the delay. Tickspeed-based delay is always dealt through attacks that deal delay. CTB-based delay is dealt, if the target is inflicted with 'haste' or 'slow'. In the case of 'haste', the 'delay' is actually a CTB-heal.",
		},
		{
			Field: 			pfnAttackType,
			Type: 			"string (enum: ctbAttackType)",	
			EnumValues: 	cfg.t.CtbAttackType.Strings(),
			DefaultVal: 	string(database.CtbAttackTypeAttack),
			Description: 	"Define whether the delay deals positive ('attack') or negative ('heal') CTB damage. In the game, CTB heals only occur, when getting inflicted with the 'haste' status.",
		},
		{
			Field: 			pfnStrength,
			Type: 			"string (enum: delayStrength)",	
			EnumValues:		cfg.t.DelayStrength.Strings(),
			RequiredOr: 	[]FieldName{pfnDelayConstant},
			ConflictsWith: 	[]FieldName{pfnDelayConstant},
			Description: 	"Define the strength of the delay. In the game, all delays actually only use four different delay constants, a weak and a strong version for both delays. Those are 8 and 16 for CTB-based delays and 24 and 48 for tickspeed-based delays. Based on the delay type and this value, the equivalent delay constant will be applied automatically, where 'weak' represents the lower constant and 'strong' the higher one.",
		},
		{
			Field: 			pfnDelayConstant,
			Type: 			"int",	
			MinVal: 		h.GetInt32Ptr(0),
			MaxVal: 		h.GetInt32Ptr(255),
			RequiredOr: 	[]FieldName{pfnStrength},
			ConflictsWith: 	[]FieldName{pfnStrength},
			Description: 	"Applies a custom delay constant.",
		},
	}
}


func (cfg *Config) getFieldDocDelayTarget() []FieldDoc {
	return []FieldDoc{
		{
			Field:         pfnMonsterID,
			Type:          "int (id: monster)",
			RequiredOr:    []FieldName{pfnAgility},
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(cfg.l.Monsters))),
			Description:   "To define a target, you can specify a monster ID and look it up. Its agility stat will be applied automatically. It can be overridden however, through the use of the agility field.",
		},
		{
			Field:       pfnAltState,
			Type:        "int (idx +1: monster altered states)",
			RequiresAll: []FieldName{pfnMonsterID},
			MinVal:      h.GetInt32Ptr(1),
			Description: "If the looked up monster has altered states, they can be applied, but will only have an effect, if they change its agility stat, or apply 'haste'. Returns an error, if the monster doesn't have as many altered states as given. The maxVal of this field is always the total amount of the monster's altered states.",
		},
		{
			Field:       pfnAgility,
			Type:        "int",
			RequiredOr:  []FieldName{pfnMonsterID},
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The targets's agility stat. If a monster was looked up, and this field is active, this value will override the monster's original agility stat.",
		},
		{
			Field:       pfnStatus,
			Type:        "string (enum: hasteStatus)",
			EnumValues:  []string{string(database.HasteStatusHaste), string(database.HasteStatusSlow)},
			Description: "Specify, whether the target carries the 'haste' or 'slow' status. Since these two statuses affect the target's tick speed, they will also affect the severity of the delay. 'Haste' halves the delay, while 'slow' doubles it.",
		},
		{
			Field:       pfnRemainingTicks,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(10000),
			Description: "Specify the amount of ticks until the target's next turn. If the delay is ctb-based, then this value will be used to calculate it.",
		},
	}
}