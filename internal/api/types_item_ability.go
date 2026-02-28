package api


type ItemAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Item					NamedAPIResource		`json:"item"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	Category              	NamedAPIResource		`json:"category"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	CanUseOutsideBattle   	bool					`json:"can_use_outside_battle"`
	Cursor                	string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}