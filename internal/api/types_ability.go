package api

// overdrive abilities need different query to get their ability attributes (GetOverdriveAbilityAttributes)
// it could be the case that the unspecified ability types don't even need a db query for that

type Ability struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	Type               NamedAPIResource    `json:"type"`
	TypedAbility       NamedAPIResource    `json:"typed_ability"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	Monsters           []NamedAPIResource  `json:"monsters"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}


type EnemyAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	UntypedAbility			NamedAPIResource		`json:"untyped_ability"`
	Effect                	*string					`json:"effect"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}


type ItemAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Item					NamedAPIResource		`json:"item"`
	UntypedAbility			NamedAPIResource		`json:"untyped_ability"`
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


type OverdriveAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	UntypedAbility			NamedAPIResource		`json:"untyped_ability"`
	User             		NamedAPIResource		`json:"user"`
	OverdriveCommand		*NamedAPIResource		`json:"overdrive_command"`
	Overdrives				[]NamedAPIResource		`json:"overdrives"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Cursor                	*string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}


type PlayerAbility struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	UntypedAbility			NamedAPIResource		`json:"untyped_ability"`
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


type TriggerCommand struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	UntypedAbility			NamedAPIResource		`json:"untyped_ability"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	Topmenu               	*NamedAPIResource		`json:"topmenu"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	Cursor                	string					`json:"cursor"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	UsedBy             		[]NamedAPIResource		`json:"used_by"`
	MonsterFormations		[]UnnamedAPIResource	`json:"monster_formations"`
	BattleInteractions		[]BattleInteraction		`json:"battle_interactions"`
}


type UnspecifiedAbility struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	UntypedAbility	   NamedAPIResource	   `json:"untyped_ability"`
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	Topmenu            *NamedAPIResource   `json:"topmenu"`
	Submenu            *NamedAPIResource   `json:"submenu"`
	OpenSubmenu        *NamedAPIResource   `json:"open_submenu,omitempty"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	Cursor             *string             `json:"cursor"`
	LearnedBy          []NamedAPIResource  `json:"learned_by"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}
