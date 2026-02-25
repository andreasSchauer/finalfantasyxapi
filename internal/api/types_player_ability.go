package api


type PlayerAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Description           	*string					`json:"description"`
	Effect                	string					`json:"effect"`
	Category              	NamedAPIResource		`json:"category"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Submenu               	*NamedAPIResource		`json:"submenu"`
	OpenSubmenu           	*NamedAPIResource		`json:"open_submenu,omitempty"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	CanUseOutsideBattle   	bool					`json:"can_use_outside_battle"`
	MpCost                	int32					`json:"mp_cost"`
	Cursor                	*string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	LearnedBy             	[]NamedAPIResource		`json:"learned_by"`
	StandardGridCharacter 	*NamedAPIResource		`json:"standard_grid_character,omitempty"`
	ExpertGridCharacter   	*NamedAPIResource		`json:"expert_grid_character,omitempty"`
	AeonLearnItem         	*ItemAmount				`json:"aeon_learn_item,omitempty"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}

