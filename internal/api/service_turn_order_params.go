package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type TurnOrderParams struct {
	TurnsAmt      int32            `json:"turns_amt,omitempty"`
	IgnFirstTurn  bool             `json:"ign_first_turn,omitempty"`
	IgnImmunities bool             `json:"ign_immunities,omitempty"`
	RNG           string           `json:"rng,omitempty"`
	BattleStart   string           `json:"battle_start,omitempty"`
	FormationID   *int32           `json:"formation_id,omitempty"`
	Party         []turnOrderParty `json:"party,omitempty"`
	Mons          []turnOrderMon   `json:"mons,omitempty"`
}

type turnOrderParty struct {
	UnitID      int32   `json:"unit_id,omitempty"`
	Agility     int32   `json:"agility,omitempty"`
	FirstStrike bool    `json:"first_strike,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (p turnOrderParty) getDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		p.UnitID,
	})
}

/*

turnOrderMon
- *turnOrderMonLookup
	- monsterID
	- *altState
	- *turnOderMonOverrides
		- *agility
		- *firstStrike
- *turnOrderMonCustom
	- name
	- agility
	- firstStrike
- *status

- look up the monster and assign name, agility, firstStrike via id and altState
- override those values based on overrides
- do the other path where you can directly assign the three values
- add the status
-

*/

// use the same separation that is used for crit chance
type turnOrderMon struct {
	MonsterID   *int32  `json:"monster_id,omitempty"`
	AltState    *int32  `json:"alt_state,omitempty"`
	Name        *string `json:"name,omitempty"`
	Agility     *int32  `json:"agility,omitempty"`
	FirstStrike *bool   `json:"first_strike,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (p turnOrderMon) getDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		h.DerefOrNil(p.MonsterID),
		h.DerefOrNil(p.AltState),
		h.DerefOrNil(p.Name),
		h.DerefOrNil(p.Agility),
		h.DerefOrNil(p.FirstStrike),
		h.DerefOrNil(p.Status),
	})
}

func getTurnOrderParamsDoc() ParamsDoc {
	return ParamsDoc{
		GeneralRules: h.GetStrPtr(""),
		Fields: []FieldDoc{
			{
				Field:       ts.PfnTurnsAmt,
				Type:        "int",
				DefaultVal:  int32(30),
				MinVal:      h.GetInt32Ptr(1),
				MaxVal:      h.GetInt32Ptr(1000),
				Description: "The amount of turns to be calculated.",
			},
			{
				Field:       ts.PfnIgnFirstTurn,
				Type:        "bool",
				Description: "If this field is true, the ICV calculation for the start of the battle is skipped and each participant in the battle starts with an ICV of 0. This allows for more direct speed comparisons between participants. This means that the values within 'rng', 'battle_start', and a participant's 'first-strike' are also ignored.",
			},
			{
				Field:       ts.PfnIgnImmunities,
				Type:        "bool",
				Description: "If this field is true, a looked up monster's immunities to 'slow' and 'haste' will be ignored when applying a status.",
			},
			{
				Field:       ts.PfnRNG,
				Type:        "string (enum: turnOrderRNG)",
				DefaultVal:  string(database.TurnOrderRngMedian),
				EnumValues:  enums.T.TurnOrderRNG.Strings(),
				Description: "Specify the rng at the start of the battle, which is used for ICV calculation. This field will have no effect, if 'ign_first_turn' is set to 'true'.",
			},
			{
				Field:       ts.PfnBattleStart,
				Type:        "string (enum: battleStart)",
				DefaultVal:  string(database.BattleStartNormal),
				EnumValues:  enums.T.BattleStart.Strings(),
				Description: "Specify, if the battle starts normally, or if it starts as an ambush or preemptive strike. This field will have no effect, if 'ign_first_turn' is set to 'true'.",
			},
			{
				Field:         ts.PfnFormationID,
				Type:          "int (id: monster formation)",
				RequiredOr:    []ts.FieldName{ts.PfnParty, ts.PfnMons},
				ConflictsWith: []ts.FieldName{ts.PfnMons},
				MinVal:        h.GetInt32Ptr(1),
				MaxVal:        h.GetInt32Ptr(int32(len(lookup.L.MonsterFormations))),
				Description:   "Instead of manually selecting monsters, you can look up a monster-formation via id. The monster(s) it contains will automatically get added as participants, along with their agility and first-strike values.",
			},
			{
				Field:       ts.PfnParty,
				Type:        "array[turnOrderParty]",
				RequiredOr:  []ts.FieldName{ts.PfnFormationID, ts.PfnMons},
				MaxArrayLen: h.GetIntPtr(3),
				Description: "Defines the participants of the player party.",
				ChildProps:  getFieldDocTurnOrderParty(),
			},
			{
				Field:         ts.PfnMons,
				Type:          "array[turnOrderMon]",
				RequiredOr:    []ts.FieldName{ts.PfnFormationID, ts.PfnParty},
				ConflictsWith: []ts.FieldName{ts.PfnFormationID},
				MaxArrayLen:   h.GetIntPtr(10),
				Description:   "Defines the participants of the opponent party by looking up monsters via id or defining custom monsters.",
				ChildProps:    getFieldDocTurnOrderMon(),
			},
		},
	}
}

func getFieldDocTurnOrderParty() []FieldDoc {
	return []FieldDoc{
		{
			Field:       ts.PfnUnitID,
			Type:        "int (id: playerUnit)",
			Required:    true,
			MinVal:      h.GetInt32Ptr(1),
			MaxVal:      h.GetInt32Ptr(int32(len(lookup.L.PlayerUnits))),
			Description: "Specifies the id of the party member to be looked up.",
		},
		{
			Field:       ts.PfnAgility,
			Type:        "int",
			Required:    true,
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The agility stat of the party member.",
		},
		{
			Field:       ts.PfnFirstStrike,
			Type:        "bool",
			Description: "If this field is true, the party member will have the 'first strike' auto-ability, which sets their ICV to 0. 'First strike' won't have an effect, if 'ign_first_turn' is set to true.",
		},
		{
			Field:       ts.PfnStatus,
			Type:        "string (enum: hasteStatus)",
			EnumValues:  enums.T.HasteStatus.Strings(),
			Description: "Specify, whether a party member carries the 'haste' or 'slow' status. If 'haste' or 'slow' are selected, it is assumed, that the party member gets this status on the first turn, so ICV calculation is unaffected. The ctb alterations that come from applying 'haste' or 'slow' are ignored. If 'auto-haste' is selected, the party member starts this battle with 'haste' (due to 'auto-haste', or 'sos-haste') which will affect ICV calculation. 'Haste' halves the bearer's tick speed, while 'slow' doubles it.",
		},
	}
}

func getFieldDocTurnOrderMon() []FieldDoc {
	return []FieldDoc{
		{
			Field:         ts.PfnMonsterID,
			Type:          "int (id: monster)",
			RequiredOr:    []ts.FieldName{ts.PfnName},
			ConflictsWith: []ts.FieldName{ts.PfnName},
			MinVal:        h.GetInt32Ptr(1),
			MaxVal:        h.GetInt32Ptr(int32(len(lookup.L.Monsters))),
			Description:   "To define a monster, you can specify an ID and look up the corresponding monster. Its agility stat, first strike, as well as its innate 'haste' status will be applied automatically. These values can be overridden however, through the use of the other fields.",
		},
		{
			Field:         ts.PfnAltState,
			Type:          "int (idx +1: monster altered states)",
			RequiresAll:   []ts.FieldName{ts.PfnMonsterID},
			ParentIDField: ts.PfnPtr(ts.PfnMonsterID),
			AllowedIDValPairs: []AllowedIDValPairs{
				{
					ID:     32,
					Values: []int32{2, 3},
					MaxVal: getLenMonAltStates(32),
				},
				{
					ID:     139,
					Values: []int32{1, 2, 3},
					MaxVal: getLenMonAltStates(139),
				},
				{
					ID:     160,
					Values: []int32{1},
					MaxVal: getLenMonAltStates(160),
				},
				{
					ID:     198,
					Values: []int32{1},
					MaxVal: getLenMonAltStates(198),
				},
				{
					ID:     285,
					Values: []int32{1},
					MaxVal: getLenMonAltStates(285),
				},
				{
					ID:     287,
					Values: []int32{2},
					MaxVal: getLenMonAltStates(287),
				},
			},
			Description: "If the looked up monster has altered states, they can be applied, but will only have an effect, if they change its agility stat, or apply 'haste'. If a monster has a different agility stat on only its first turn, this altered state is applied automatically, except when 'ign_first_turn' is true. Specifying that specific altered state won't have any effect, nor will it result in an error. The same is true for Penance's arms who get the 'haste' status only during their own turns. Returns an error, if the monster doesn't have as many altered states as given. The maxVal of this field is always the total amount of the monster's altered states.",
		},
		{
			Field:         ts.PfnName,
			Type:          "string",
			RequiredOr:    []ts.FieldName{ts.PfnMonsterID},
			ConflictsWith: []ts.FieldName{ts.PfnMonsterID},
			RequiresAll:   []ts.FieldName{ts.PfnAgility},
			Description:   "Use this field, if you want to define a custom monster. The given name will show up in the list to represent it. Multiple monsters with the same name will be numbered automatically.",
		},
		{
			Field:       ts.PfnAgility,
			Type:        "int",
			MinVal:      h.GetInt32Ptr(0),
			MaxVal:      h.GetInt32Ptr(255),
			Description: "The monsters's agility stat. If a monster was looked up, and this field is active, this value will override the monster's original agility stat.",
		},
		{
			Field:       ts.PfnFirstStrike,
			Type:        "bool",
			Description: "If this field is true, the monster will have the 'first strike' auto-ability, which sets its ICV to -1. If a monster was looked up, and this field is active, this value will override the monster's original (lack of) access to this ability. 'First strike' won't have an effect, if 'ign_first_turn' is set to true.",
		},
		{
			Field:       ts.PfnStatus,
			Type:        "string (enum: hasteStatus)",
			EnumValues:  enums.T.HasteStatus.Strings(),
			Description: "Specify, whether the monster carries the 'haste' or 'slow' status. If 'haste' or 'slow' are selected, it is assumed, that the monster gets this status on its first turn, so ICV calculation is unaffected. The ctb alterations that come from applying 'haste' or 'slow' are ignored. If 'auto-haste' is selected, the monster starts this battle with 'haste' (due to 'auto-haste', or 'sos-haste') which will affect ICV calculation. If a monster was looked up, and this field is active, this value will override the monster's original haste status, though in practice, this will only affect Penance's arms (IDs 306/307). 'Haste' halves the bearer's tick speed, while 'slow' doubles it.",
		},
	}
}

func getLenMonAltStates(id int32) int32 {
	mon, _ := lookup.GetResourceByID(id, lookup.L.MonstersID)
	return h.Len32(mon.AlteredStates)
}
