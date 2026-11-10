package servicetypes

import (
	"fmt"
)

type TurnOrderResponse struct {
	TurnsAmt      int32         `json:"turns_amt"`
	IgnFirstTurn  bool          `json:"ign_first_turn"`
	IgnImmunities bool          `json:"ign_immunities"`
	BattleStart   string        `json:"battle_start"`
	RNG           string        `json:"rng"`
	PlayerParty   []Participant `json:"player_party"`
	OpponentParty []Participant `json:"opponent_party"`
	TurnOrder     []BattleTurn  `json:"turn_order"`
}

type Participant struct {
	Name    string      `json:"name"`
	Party   BattleParty `json:"-"`
	Agility int32       `json:"agility"`
	AgilityVals
	FirstStrike   bool    `json:"first_strike"`
	Status        *string `json:"status"`
	AltState      *int32  `json:"alt_state,omitempty"`
	TurnsReceived int32   `json:"turns_received"`
	TurnsPercent  float64 `json:"turns_percent"`
}

func (p Participant) GetKey() string {
	if p.Party == BattlePartyPlayer {
		return p.Name
	}

	return fmt.Sprintf("mon|%s", p.Name)
}

type BattleParty string

const (
	BattlePartyPlayer   BattleParty = "player"
	BattlePartyOpponent BattleParty = "opponent"
)

type AgilityVals struct {
	TickSpeed int32  `json:"tick_speed"`
	MinICV    *int32 `json:"min_icv"`
	MaxICV    *int32 `json:"max_icv"`
}

type BattleTurn struct {
	Turn        int32       `json:"turn"`
	CurrentTick int32       `json:"current_tick"`
	Name        string      `json:"name"`
	Party       BattleParty `json:"party"`
	PriorityKey string      `json:"-"`
}