package servicetypes

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type TurnOrderParams struct {
	TurnsAmt      int32            `json:"turns_amt,omitempty"`
	IgnFirstTurn  bool             `json:"ign_first_turn,omitempty"`
	IgnImmunities bool             `json:"ign_immunities,omitempty"`
	RNG           string           `json:"rng,omitempty"`
	BattleStart   string           `json:"battle_start,omitempty"`
	FormationID   *int32           `json:"formation_id,omitempty"`
	Party         []TurnOrderParty `json:"party,omitempty"`
	Mons          []TurnOrderMon   `json:"mons,omitempty"`
}

type TurnOrderParty struct {
	UnitID      int32   `json:"unit_id,omitempty"`
	Agility     int32   `json:"agility,omitempty"`
	FirstStrike bool    `json:"first_strike,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (p TurnOrderParty) GetDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		p.UnitID,
	})
}

type TurnOrderMon struct {
	Lookup *TurnOrderMonLookup `json:"lookup,omitempty"`
	Custom *TurnOrderMonCustom `json:"custom,omitempty"`
}

func (p TurnOrderMon) GetDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		safeDuplicateKey(p.Lookup),
		safeDuplicateKey(p.Custom),
	})
}

type TurnOrderMonLookup struct {
	MonsterID int32                  `json:"monster_id"`
	AltState  *int32                 `json:"alt_state"`
	Overrides *TurnOrderMonOverrides `json:"overrides,omitempty"`
}

func (p TurnOrderMonLookup) GetDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		p.MonsterID,
		h.DerefOrNil(p.AltState),
		safeDuplicateKey(p.Overrides),
	})
}

type TurnOrderMonOverrides struct {
	Agility     *int32  `json:"agility,omitempty"`
	FirstStrike *bool   `json:"first_strike,omitempty"`
	Status      *string `json:"status,omitempty"`
}

func (p TurnOrderMonOverrides) GetDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		h.DerefOrNil(p.Agility),
		h.DerefOrNil(p.FirstStrike),
		h.DerefOrNil(p.Status),
	})
}

type TurnOrderMonCustom struct {
	Name        string  `json:"name"`
	Agility     int32   `json:"agility"`
	FirstStrike bool    `json:"first_strike"`
	Status      *string `json:"status,omitempty"`
}

func (p TurnOrderMonCustom) GetDuplicateKey() string {
	return h.CombineFields([]any{
		fmt.Sprintf("%T", p),
		p.Name,
		p.Agility,
		p.FirstStrike,
		h.DerefOrNil(p.Status),
	})
}
