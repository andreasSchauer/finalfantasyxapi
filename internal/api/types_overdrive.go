package api

type Overdrive struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Version               	*int32					`json:"version,omitempty"`
	Specification         	*string					`json:"specification,omitempty"`
	Description           	string					`json:"description"`
	Effect                	string					`json:"effect"`
	Rank                  	*int32					`json:"rank"`
	AppearsInHelpBar      	bool					`json:"appears_in_help_bar"`
	CanCopycat            	bool					`json:"can_copycat"`
	UnlockCondition			*string					`json:"unlock_condition"`
	CountdownInSec			*int32					`json:"countdown_in_sec"`
	User             		NamedAPIResource		`json:"user"`
	OverdriveCommand		*NamedAPIResource		`json:"overdrive_command"`
	OverdriveAbilities		[]NamedAPIResource		`json:"overdrive_abilities"`
}