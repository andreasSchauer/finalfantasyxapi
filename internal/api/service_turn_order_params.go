package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TurnOrderParams struct {
	TurnsAmt     int32                `json:"turns_amt,omitempty"`
	IgnFirstTurn bool                 `json:"ign_first_turn,omitempty"`
	RNG          string               `json:"rng,omitempty"`
	BattleStart  string               `json:"battle_start,omitempty"`
	AglK         int32                `json:"agl_k,omitempty"`
	AglY         int32                `json:"agl_y,omitempty"`
	Battles      int32                `json:"battles,omitempty"`
	Formation    *int32               `json:"formation,omitempty"`
	Party        []turnOrderParty     `json:"party,omitempty"`
	Mons         []turnOrderMon       `json:"mons,omitempty"`
	MonsCustom   []turnOrderMonCustom `json:"mons_custom,omitempty"`
}

func (p TurnOrderParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getTurnOrderParamsDoc()
}

type turnOrderParty struct {
	Name   	string  `json:"name,omitempty"`
	Agl    	int32   `json:"agl"`
	FS     	bool    `json:"fs,omitempty"`
	Status 	*string `json:"status,omitempty"`
	Offset	int32	`json:"offset,omitempty"`
}

type turnOrderMon struct {
	ID       	int32   `json:"id"`
	AltState 	*int32  `json:"alt_state,omitempty"`
	Status   	*string `json:"status,omitempty"`
	Offset		int32	`json:"offset,omitempty"`
}

type turnOrderMonCustom struct {
	Name   	string  `json:"name,omitempty"`
	Agl    	int32   `json:"agl"`
	FS     	bool    `json:"fs,omitempty"`
	Status 	*string `json:"status,omitempty"`
	Offset	int32	`json:"offset,omitempty"`
}

func (cfg *Config) getTurnOrderParamsDoc() ParamsDoc {
	return ParamsDoc{
		GeneralRules: h.GetStrPtr(""),
		Fields: []FieldDoc{
			{
				Field:       pfnTurnsAmt,
				Type:        "int",
				DefaultVal:  30,
				MinVal:      h.GetInt32Ptr(1),
				MaxVal:      h.GetInt32Ptr(1000),
				Description: "The amount of turns to be calculated.",
			},
			{
				Field:       pfnIgnFirstTurn,
				Type:        "bool",
				Description: "If this field is true, the ICV calculation for the start of the battle is skipped and each participant in the battle starts with an ICV of 0. This allows for more direct speed comparisons between participants.",
			},
			{
				Field:       pfnRNG,
				Type:        "string (enum: turnOrderRNG)",
				DefaultVal:  string(database.TurnOrderRngMedian),
				EnumValues:  createEnumStringSlice(cfg.t.TurnOrderRNG.lookup),
				Description: "Specify the rng at the start of the battle, which is used for ICV calculation. This field will have no effect, if 'ign_first_turn' is set to 'true'.",
			},
			{
				Field:       pfnBattleStart,
				Type:        "string (enum: battleStart)",
				DefaultVal:  string(database.BattleStartNormal),
				EnumValues:  createEnumStringSlice(cfg.t.BattleStart.lookup),
				Description: "Specify, if the battle starts normally, or if it starts as an ambush or preemptive strike. This field will have no effect, if 'ign_first_turn' is set to 'true'.",
			},
			{
				Field:       pfnAglK,
				Type:        "int",
				DefaultVal:  6,
				MinVal:      h.GetInt32Ptr(0),
				MaxVal:      h.GetInt32Ptr(255),
				Description: "Sets Kimahri's agility. When you look up Biran and Yenke Ronso (ids 167/168), their agility stat will get calculated using this value.",
			},
			{
				Field:       pfnAglY,
				Type:        "int",
				DefaultVal:  10,
				MinVal:      h.GetInt32Ptr(0),
				MaxVal:      h.GetInt32Ptr(255),
				Description: "Sets Yuna's agility. When you look up the possessed aeons (ids 216-225), their agility stat will automatically get calculated using this value and the value given by the 'battles' field.",
			},
			{
				Field:       pfnBattles,
				Type:        "int",
				DefaultVal:  0,
				MinVal:      h.GetInt32Ptr(0),
				Description: "Sets the amount of battles the player has taken part in. When you look up the possessed aeons (ids 216-225), their agility stat will automatically get calculated using this value and the value given by the 'agl_y' field. The highest value that has an impact is 600, an amount that should have been comfortably reached by this point in the story on a casual playthrough. If you're unsure, just manually create a custom monster with your own aeon's agility stat.",
			},
			{
				Field:         pfnFormation,
				Type:          "int (id: monster formation)",
				RequiredOr:    []FieldName{pfnParty, pfnMons, pfnMonsCustom},
				ConflictsWith: []FieldName{pfnMons, pfnMonsCustom},
				MinVal:        h.GetInt32Ptr(1),
				MaxVal:        h.GetInt32Ptr(int32(len(cfg.l.MonsterFormations))),
				Description:   "Instead of manually selecting monsters, you can look up a monster-formation via id. The monster(s) it contains will automatically get added as participants.",
			},
			{
				Field:       pfnParty,
				Type:        "array[turnOrderParty]",
				RequiredOr:  []FieldName{pfnFormation, pfnMons, pfnMonsCustom},
				MaxArrayLen: h.GetIntPtr(3),
				Description: "Defines the participants of the player party.",
				ChildProps:  cfg.getFieldDocTurnOrderParty(),
			},
			{
				Field:         pfnMons,
				Type:          "array[turnOrderMons]",
				RequiredOr:    []FieldName{pfnFormation, pfnParty, pfnMonsCustom},
				ConflictsWith: []FieldName{pfnFormation},
				MaxArrayLen:   h.GetIntPtr(10),
				Description:   "Defines the participants of the opponent party by looking up monsters via id. The monster's own agility stat is used for the calculation. 'First strike' is considered, if 'ign_first_turn' is set to false.",
				ChildProps:    cfg.getFieldDocTurnOrderMon(),
			},
			{
				Field:         pfnMonsCustom,
				Type:          "array[turnOrderMonsCustom]",
				RequiredOr:    []FieldName{pfnFormation, pfnParty, pfnMons},
				ConflictsWith: []FieldName{pfnFormation},
				MaxArrayLen:   h.GetIntPtr(10),
				Description:   "Defines the participants of the opponent party by defining custom monsters with custom agility stats.",
				ChildProps:    cfg.getFieldDocTurnOrderMonCustom(),
			},
		},
	}
}


func (cfg *Config) getFieldDocTurnOrderParty() []FieldDoc {
	return []FieldDoc{
		cfg.getFieldDocName("char"),
		cfg.getFieldDocAgl(),
		cfg.getFieldDocFS(),
		cfg.getFieldDocStatus(),
		cfg.getFieldDocOffset(),
	}
}

func (cfg *Config) getFieldDocTurnOrderMon() []FieldDoc {
	return []FieldDoc{
		{
			Field:       pfnID,
			Type:        "int (id: monster)",
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(int32(len(cfg.l.Monsters))),
			Description: "Specifies the id of the monster to be looked up.",
		},
		{
			Field:       pfnAltState,
			Type:        "int (id: monster altered states)",
			MinVal:      h.GetInt32Ptr(1),
			Description: "If a monster has altered states, they can be applied, but will only have an effect, if they change its agility stat, or apply 'haste'. If a monster has a different agility stat, only on its first turn, this logic is applied automatically, except when 'ign_first_turn' is true. Specifying that specific altered state will have no effect, but will not result in an error. The same is true for penance's arms who get 'haste' status only during their own turns.",
		},
		cfg.getFieldDocStatus(),
		cfg.getFieldDocOffset(),
	}
}

func (cfg *Config) getFieldDocTurnOrderMonCustom() []FieldDoc {
	return []FieldDoc{
		cfg.getFieldDocName("custom mon"),
		cfg.getFieldDocAgl(),
		cfg.getFieldDocFS(),
		cfg.getFieldDocStatus(),
		cfg.getFieldDocOffset(),
	}
}

func (cfg *Config) getFieldDocName(defName string) FieldDoc {
	return FieldDoc{
		Field:       pfnName,
		Type:        "string",
		DefaultVal:  fmt.Sprintf("%s {idx + 1}", defName),
		Description: fmt.Sprintf("The name of the participant that will show up in the list. If no name was given, the name '%s', along with its index in the list + 1 will be used.", defName),
	}
}

func (cfg *Config) getFieldDocAgl() FieldDoc {
	return FieldDoc{
		Field:       pfnAgl,
		Type:        "int",
		Required:    true,
		MinVal:      h.GetInt32Ptr(0),
		MaxVal:      h.GetInt32Ptr(255),
		Description: "The agility stat of the participant.",
	}
}

func (cfg *Config) getFieldDocFS() FieldDoc {
	return FieldDoc{
		Field:       pfnFS,
		Type:        "bool",
		Description: "If this field is true, the participant will have the 'first strike' auto-ability. This sets a party member's ICV to 0 and a monster's ICV to -1. 'First strike' won't have an effect, if 'ign_first_turn' is set to true.",
	}
}

func (cfg *Config) getFieldDocStatus() FieldDoc {
	return FieldDoc{
		Field:       pfnStatus,
		Type:        "string (enum: hasteStatus)",
		EnumValues:  createEnumStringSlice(cfg.t.HasteStatus.lookup),
		Description: "Specify, whether a participant carries the 'haste' or 'slow' status. If 'haste' or 'slow' are selected, it is assumed, that the participant gets this status on the first turn, so ICV calculation is unaffected. The ctb alterations that come from applying 'haste' or 'slow' are ignored. If 'auto-haste' is selected, the participant starts this battle with 'haste' (due to 'auto-haste', or 'sos-haste') which will affect ICV calculation.",
	}
}

func (cfg *Config) getFieldDocOffset() FieldDoc {
	return FieldDoc{
		Field:       pfnOffset,
		Type:        "int",
		DefaultVal:  0,
		Description: "If 'ign_first_turn' is set to true, this field specifies the starting tick of the participant.",
	}
}