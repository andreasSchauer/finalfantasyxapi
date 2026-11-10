package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Overdrive struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Version            *int32              `json:"version,omitempty"`
	Specification      *string             `json:"specification,omitempty"`
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	Rank               *int32              `json:"rank"`
	AppearsInHelpBar   bool                `json:"appears_in_help_bar"`
	CanCopycat         bool                `json:"can_copycat"`
	UnlockCondition    *string             `json:"unlock_condition"`
	CountdownInSec     *int32              `json:"countdown_in_sec"`
	Cursor             *string             `json:"cursor"`
	User               apiresource.Named   `json:"user"`
	OverdriveCommand   *apiresource.Named  `json:"overdrive_command"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
}
