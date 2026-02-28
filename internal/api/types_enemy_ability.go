package api


type EnemyAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Effect                	*string					`json:"effect"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}