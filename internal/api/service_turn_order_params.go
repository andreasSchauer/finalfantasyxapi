package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type TurnOrderParams struct {
	TurnsAmt      int32            `json:"turns_amt,omitempty"`
	IgnFirstTurn  bool             `json:"ign_first_turn,omitempty"`
	IgnImmunities bool             `json:"ign_immunities,omitempty"`
	RNG           string           `json:"rng,omitempty"`
	BattleStart   string           `json:"battle_start,omitempty"`
	Formation     *int32           `json:"formation,omitempty"`
	Party         []turnOrderParty `json:"party,omitempty"`
	Mons          []turnOrderMon   `json:"mons,omitempty"`
}

func (p TurnOrderParams) GetDoc(cfg *Config) ParamsDoc {
	return cfg.getTurnOrderParamsDoc()
}

type turnOrderParty struct {
	ID      int32   `json:"id"`
	Agility int32   `json:"agility"`
	FS      bool    `json:"first_strike,omitempty"`
	Status  *string `json:"status,omitempty"`
}

func (p turnOrderParty) getDuplicateKey() string {
	return seeding.CombineFields([]any{
		fmt.Sprintf("%T", p),
		p.ID,
	})
}

type turnOrderMon struct {
	ID          *int32  `json:"id"`
	AltState    *int32  `json:"alt_state,omitempty"`
	Name        *string `json:"name,omitempty"`
	Agility     *int32  `json:"agility"`
	FirstStrike *bool   `json:"first_strike,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (p turnOrderMon) getDuplicateKey() string {
	return seeding.CombineFields([]any{
		fmt.Sprintf("%T", p),
		h.DerefOrNil(p.ID),
		h.DerefOrNil(p.AltState),
		h.DerefOrNil(p.Name),
		h.DerefOrNil(p.Agility),
		h.DerefOrNil(p.FirstStrike),
		h.DerefOrNil(p.Status),
	})
}

func (cfg *Config) getTurnOrderParamsDoc() ParamsDoc {
	return ParamsDoc{
		GeneralRules: h.GetStrPtr(""),
		Fields: []FieldDoc{
			{
				Field:       pfnTurnsAmt,
				Type:        "int",
				DefaultVal:  int32(30),
				MinVal:      h.GetInt32Ptr(1),
				MaxVal:      h.GetInt32Ptr(1000),
				Description: "The amount of turns to be calculated.",
			},
			{
				Field:       pfnIgnFirstTurn,
				Type:        "bool",
				Description: "If this field is true, the ICV calculation for the start of the battle is skipped and each participant in the battle starts with an ICV of 0. This allows for more direct speed comparisons between participants. This means that the values within 'rng', 'battle_start', and a participant's 'first-strike' are also ignored.",
			},
			{
				Field:       pfnIgnImmunities,
				Type:        "bool",
				Description: "If this field is true, a looked up monster's immunities to 'slow' and 'haste' will be ignored when applying a status.",
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
				Field:         pfnFormation,
				Type:          "int (id: monster formation)",
				RequiredOr:    []FieldName{pfnParty, pfnMons},
				ConflictsWith: []FieldName{pfnMons},
				MinVal:        h.GetInt32Ptr(1),
				MaxVal:        h.GetInt32Ptr(int32(len(cfg.l.MonsterFormations))),
				Description:   "Instead of manually selecting monsters, you can look up a monster-formation via id. The monster(s) it contains will automatically get added as participants, along with their agility and first-strike values.",
			},
			{
				Field:       pfnParty,
				Type:        "array[turnOrderParty]",
				RequiredOr:  []FieldName{pfnFormation, pfnMons},
				MaxArrayLen: h.GetIntPtr(3),
				Description: "Defines the participants of the player party.",
				ChildProps:  cfg.getFieldDocTurnOrderParty(),
			},
			{
				Field:         pfnMons,
				Type:          "array[turnOrderMon]",
				RequiredOr:    []FieldName{pfnFormation, pfnParty},
				ConflictsWith: []FieldName{pfnFormation},
				MaxArrayLen:   h.GetIntPtr(10),
				Description:   "Defines the participants of the opponent party by looking up monsters via id or defining custom monsters.",
				ChildProps:    cfg.getFieldDocTurnOrderMon(),
			},
		},
	}
}

func (cfg *Config) getFieldDocTurnOrderParty() []FieldDoc {
	return []FieldDoc{
		{
			Field:       pfnID,
			Type:        "int (id: playerUnit)",
			Required:    true,
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(int32(len(cfg.l.PlayerUnits))),
			Description: "Specifies the id of the party member to be looked up.",
		},
		{
			Field:       pfnAgility,
			Type:        "int",
			Required:    true,
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The agility stat of the party member.",
		},
		{
			Field:       pfnFirstStrike,
			Type:        "bool",
			Description: "If this field is true, the party member will have the 'first strike' auto-ability, which sets their ICV to 0. 'First strike' won't have an effect, if 'ign_first_turn' is set to true.",
		},
		{
			Field:       pfnStatus,
			Type:        "string (enum: hasteStatus)",
			EnumValues:  createEnumStringSlice(cfg.t.HasteStatus.lookup),
			Description: "Specify, whether a party member carries the 'haste' or 'slow' status. If 'haste' or 'slow' are selected, it is assumed, that the party member gets this status on the first turn, so ICV calculation is unaffected. The ctb alterations that come from applying 'haste' or 'slow' are ignored. If 'auto-haste' is selected, the party member starts this battle with 'haste' (due to 'auto-haste', or 'sos-haste') which will affect ICV calculation.",
		},
	}
}

func (cfg *Config) getFieldDocTurnOrderMon() []FieldDoc {
	return []FieldDoc{
		{
			Field:         pfnID,
			Type:          "int (id: monster)",
			RequiredOr:    []FieldName{pfnName},
			ConflictsWith: []FieldName{pfnName},
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(cfg.l.Monsters))),
			Description:   "To define a monster, you can specify an ID and look up the corresponding monster. Its agility stat, first strike, as well as its innate 'haste' status will be applied automatically. These values can be overridden however, through the use of the other fields.",
		},
		{
			Field:       pfnAltState,
			Type:        "int (id: monster altered states)",
			RequiresAll: []FieldName{pfnID},
			MinVal:      h.GetInt32Ptr(1),
			Description: "If the looked up monster has altered states, they can be applied, but will only have an effect, if they change its agility stat, or apply 'haste'. If a monster has a different agility stat on only its first turn, this altered state is applied automatically, except when 'ign_first_turn' is true. Specifying that specific altered state won't have any effect, nor will it result in an error. The same is true for Penance's arms who get the 'haste' status only during their own turns.",
		},
		{
			Field:         pfnName,
			Type:          "string",
			RequiredOr:    []FieldName{pfnID},
			ConflictsWith: []FieldName{pfnID},
			RequiresAll:   []FieldName{pfnAgility},
			Description:   "Use this field, if you want to define a custom monster. The given name will show up in the list to represent it. Multiple monsters with the same name will be numbered automatically.",
		},
		{
			Field:       pfnAgility,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The monsters's agility stat. If a monster was looked up, and this field is active, this value will override the monster's original agility stat.",
		},
		{
			Field:       pfnFirstStrike,
			Type:        "bool",
			Description: "If this field is true, the monster will have the 'first strike' auto-ability, which sets its ICV to -1. If a monster was looked up, and this field is active, this value will override the monster's original (lack of) access to this ability. 'First strike' won't have an effect, if 'ign_first_turn' is set to true.",
		},
		{
			Field:       pfnStatus,
			Type:        "string (enum: hasteStatus)",
			EnumValues:  createEnumStringSlice(cfg.t.HasteStatus.lookup),
			Description: "Specify, whether the monster carries the 'haste' or 'slow' status. If 'haste' or 'slow' are selected, it is assumed, that the monster gets this status on the first turn, so ICV calculation is unaffected. The ctb alterations that come from applying 'haste' or 'slow' are ignored. If 'auto-haste' is selected, the monster starts this battle with 'haste' (due to 'auto-haste', or 'sos-haste') which will affect ICV calculation. If a monster was looked up, and this field is active, this value will override the monster's original haste status, though in practice, this will only affect Penance's arms (IDs 306/307).",
		},
	}
}