package api

// overdrive abilities need different query to get their ability attributes (GetOverdriveAbilityAttributes)
// it could be the case that the other ability types don't even need a db query for that

type Ability struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Type					NamedAPIResource		`json:"type"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Users					[]NamedAPIResource		`json:"users"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}