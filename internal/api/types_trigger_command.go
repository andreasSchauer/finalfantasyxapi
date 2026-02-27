package api


type TriggerCommand struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	Category              	NamedAPIResource		`json:"category"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Cursor                	*string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	UsedBy             		[]NamedAPIResource		`json:"used_by"`
	MonsterFormations		[]UnnamedAPIResource	`json:"monster_formations"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}

