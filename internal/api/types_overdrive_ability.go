package api


type OverdriveAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	User             		NamedAPIResource		`json:"user"`
	OverdriveCommand		*NamedAPIResource		`json:"overdrive_command"`
	Overdrive				NamedAPIResource		`json:"overdrive"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Cursor                	*string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}
