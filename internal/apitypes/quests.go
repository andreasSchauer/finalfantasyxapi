package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Quest struct {
	ID           int32              `json:"id"`
	Name         string             `json:"name"`
	Type         database.QuestType `json:"type"`
	TypedQuest   apiresource.Quest  `json:"typed_quest"`
	Availability string             `json:"availability"`
	IsRepeatable bool               `json:"is_repeatable"`
	Completion   *QuestCompletion   `json:"completion"`
}

type Sidequest struct {
	ID           int32               `json:"id"`
	Name         string              `json:"name"`
	UntypedQuest apiresource.Quest   `json:"untyped_quest"`
	Availability string              `json:"availability"`
	IsRepeatable bool                `json:"is_repeatable"`
	Completion   *QuestCompletion    `json:"completion"`
	Subquests    []apiresource.Quest `json:"subquests"`
}

type Subquest struct {
	ID              int32              `json:"id"`
	Name            string             `json:"name"`
	UntypedQuest    apiresource.Quest  `json:"untyped_quest"`
	Availability    string             `json:"availability"`
	IsRepeatable    bool               `json:"is_repeatable"`
	ParentSidequest apiresource.Quest  `json:"parent_sidequest"`
	Completion      QuestCompletion    `json:"completion"`
	ArenaCreation   *apiresource.Named `json:"arena_creation,omitempty"`
}

type QuestCompletion struct {
	Condition *string                                       `json:"condition"`
	Areas     []CompletionArea                              `json:"areas"`
	Reward    apiresource.ResourceAmount[apiresource.Typed] `json:"reward"`
}

type CompletionArea struct {
	Area  apiresource.Area `json:"area"`
	Notes *string          `json:"notes,omitempty"`
}

func (ca CompletionArea) GetAPIResource() apiresource.APIResource {
	return ca.Area
}
