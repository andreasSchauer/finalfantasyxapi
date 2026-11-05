package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type Quest struct {
	ID           int32              `json:"id"`
	Name         string             `json:"name"`
	Type         database.QuestType `json:"type"`
	TypedQuest   QuestAPIResource   `json:"typed_quest"`
	Availability string             `json:"availability"`
	IsRepeatable bool               `json:"is_repeatable"`
	Completion   *QuestCompletion   `json:"completion"`
}

type Sidequest struct {
	ID           int32              `json:"id"`
	Name         string             `json:"name"`
	UntypedQuest QuestAPIResource   `json:"untyped_quest"`
	Availability string             `json:"availability"`
	IsRepeatable bool               `json:"is_repeatable"`
	Completion   *QuestCompletion   `json:"completion"`
	Subquests    []QuestAPIResource `json:"subquests"`
}

type Subquest struct {
	ID              int32             `json:"id"`
	Name            string            `json:"name"`
	UntypedQuest    QuestAPIResource  `json:"untyped_quest"`
	Availability    string            `json:"availability"`
	IsRepeatable    bool              `json:"is_repeatable"`
	ParentSidequest QuestAPIResource  `json:"parent_sidequest"`
	Completion      QuestCompletion   `json:"completion"`
	ArenaCreation   *NamedAPIResource `json:"arena_creation,omitempty"`
}

type QuestCompletion struct {
	Condition *string                          `json:"condition"`
	Areas     []CompletionArea                 `json:"areas"`
	Reward    ResourceAmount[TypedAPIResource] `json:"reward"`
}

func convertQuestCompletion(qc seedtypes.QuestCompletion) QuestCompletion {
	return QuestCompletion{
		Condition: qc.Condition,
		Areas:     convertObjSlice(qc.Areas, convertCompletionArea),
		Reward:    E.AllItems.toResAmtFromName(qc.Reward),
	}
}

type CompletionArea struct {
	Area  AreaAPIResource `json:"area"`
	Notes *string         `json:"notes,omitempty"`
}

func (ca CompletionArea) GetAPIResource() APIResource {
	return ca.Area
}

func convertCompletionArea(cl seedtypes.CompletionArea) CompletionArea {
	return CompletionArea{
		Area:  locAreaToAreaAPIResource(E.Areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}
