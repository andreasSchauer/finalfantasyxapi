package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type Quest struct {
	ID           int32                        `json:"id"`
	Name         string                       `json:"name"`
	Type         database.QuestType           `json:"type"`
	TypedQuest   apiresource.QuestAPIResource `json:"typed_quest"`
	Availability string                       `json:"availability"`
	IsRepeatable bool                         `json:"is_repeatable"`
	Completion   *QuestCompletion             `json:"completion"`
}

type Sidequest struct {
	ID           int32                          `json:"id"`
	Name         string                         `json:"name"`
	UntypedQuest apiresource.QuestAPIResource   `json:"untyped_quest"`
	Availability string                         `json:"availability"`
	IsRepeatable bool                           `json:"is_repeatable"`
	Completion   *QuestCompletion               `json:"completion"`
	Subquests    []apiresource.QuestAPIResource `json:"subquests"`
}

type Subquest struct {
	ID              int32                         `json:"id"`
	Name            string                        `json:"name"`
	UntypedQuest    apiresource.QuestAPIResource  `json:"untyped_quest"`
	Availability    string                        `json:"availability"`
	IsRepeatable    bool                          `json:"is_repeatable"`
	ParentSidequest apiresource.QuestAPIResource  `json:"parent_sidequest"`
	Completion      QuestCompletion               `json:"completion"`
	ArenaCreation   *apiresource.NamedAPIResource `json:"arena_creation,omitempty"`
}

type QuestCompletion struct {
	Condition *string                                                  `json:"condition"`
	Areas     []CompletionArea                                         `json:"areas"`
	Reward    apiresource.ResourceAmount[apiresource.TypedAPIResource] `json:"reward"`
}

func convertQuestCompletion(qc seedtypes.QuestCompletion) QuestCompletion {
	return QuestCompletion{
		Condition: qc.Condition,
		Areas:     ah.ConvertObjSlice(qc.Areas, convertCompletionArea),
		Reward:    E.AllItems.toResAmtFromName(qc.Reward),
	}
}

type CompletionArea struct {
	Area  apiresource.AreaAPIResource `json:"area"`
	Notes *string                     `json:"notes,omitempty"`
}

func (ca CompletionArea) GetAPIResource() apiresource.APIResource {
	return ca.Area
}

func convertCompletionArea(cl seedtypes.CompletionArea) CompletionArea {
	return CompletionArea{
		Area:  locAreaToAreaAPIResource(E.Areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}
