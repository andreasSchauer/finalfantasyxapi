package api


type TurnOrderResponse struct {
	URL            	string 			`json:"url"`
	Next			string			`json:"next"`
	IgnFirstTurn	bool			`json:"ign_first_turn"`
	BattleStart		string			`json:"battle_start"`
	PlayerParty		[]Participant	`json:"player_party"`
	OpponentParty	[]Participant	`json:"opponent_party"`
	TurnOrder		[]BattleTurn	`json:"turn_order"`
}

func (r TurnOrderResponse) GetURL() string {
	return r.URL
}

type Participant struct {
	Name			string		`json:"name"`
	Agility			int32		`json:"agility"`
	TickSpeed		int32		`json:"tick_speed"`
	FirstStrike		bool		`json:"first_strike"`
	Status			*string		`json:"status"`
	AltState		*int32		`json:"alt_state,omitempty"`
	TurnsReceived	int32		`json:"turns_received"`
	TurnsPercentage	float64		`json:"turns_percentage"`
}

type BattleParty string

const (
	battlePartyPlayer		BattleParty = "player"
	battlePartyOpponent		BattleParty = "opponent"
)

type BattleTurn struct {
	Name				string		`json:"name"`
	Party				BattleParty	`json:"party"`
	TicksNextTurn		int			`json:"ticks_next_turn"`
	TotalTicksPassed	int			`json:"total_ticks_passed"`
}

func calcTurnOrder(cfg *Config, params TurnOrderParams, url string) (TurnOrderResponse, error) {
	return TurnOrderResponse{}, nil
}