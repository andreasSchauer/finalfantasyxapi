package api


type OtherAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Submenu               	*NamedAPIResource		`json:"submenu"`
	OpenSubmenu           	*NamedAPIResource		`json:"open_submenu,omitempty"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Cursor                	*string					`json:"cursor"`
	LearnedBy             	[]NamedAPIResource		`json:"learned_by"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}

