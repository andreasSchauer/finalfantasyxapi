package servicetypes

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

func (p Participant) GetTurnParams(rng string) TurnParams {
	params := TurnParams{
		Name:        p.Name,
		PriorityKey: p.GetKey(),
		Party:       p.Party,
		Agility:     p.Agility,
		TickSpeed:   p.TickSpeed,
	}

	switch rng {
	case string(database.TurnOrderRngBest):
		switch p.Party {
		case BattlePartyPlayer:
			params.TicksNextTurn = *p.MinICV

		case BattlePartyOpponent:
			params.TicksNextTurn = *p.MaxICV
		}

	case string(database.TurnOrderRngWorst):
		switch p.Party {
		case BattlePartyPlayer:
			params.TicksNextTurn = *p.MaxICV

		case BattlePartyOpponent:
			params.TicksNextTurn = *p.MinICV
		}

	case string(database.TurnOrderRngMedian):
		params.TicksNextTurn = (*p.MinICV + *p.MaxICV) / 2
	}

	return params
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

type TurnParams struct {
	Name          string
	Party         BattleParty
	PriorityKey   string
	Agility       int32
	TickSpeed     int32
	TicksNextTurn int32
}