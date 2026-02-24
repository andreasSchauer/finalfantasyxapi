package api

import (

)

type PlayerAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version"`
	Specification         	*string					`json:"specification"`
	Description           	*string					`json:"description"`
	Effect                	string					`json:"effect"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	CanUseOutsideBattle   	bool					`json:"can_use_outside_battle"`
	MpCost                	*int32					`json:"mp_cost"`
	Category              	NamedAPIResource		`json:"category"`
	AeonLearnItem         	*ItemAmount				`json:"aeon_learn_item"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Submenu               	*NamedAPIResource		`json:"submenu"`
	OpenSubmenu           	*NamedAPIResource		`json:"open_submenu"`
	Cursor                	*string					`json:"cursor"`
	LearnedBy             	[]NamedAPIResource		`json:"learned_by"`
	StandardGridCharacter 	*NamedAPIResource		`json:"standard_grid_character"`
	ExpertGridCharacter   	*NamedAPIResource		`json:"expert_grid_character"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}

