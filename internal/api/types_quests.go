package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Quest struct {
	ID           int32            `json:"id"`
	Name         string           `json:"name"`
	Type         NamedAPIResource `json:"type"`
	TypedQuest   QuestAPIResource `json:"typed_quest"`
	Availability NamedAPIResource `json:"availability"`
	IsRepeatable bool             `json:"is_repeatable"`
}

type Sidequest struct {
	ID           int32              `json:"id"`
	Name         string             `json:"name"`
	UntypedQuest QuestAPIResource   `json:"untyped_quest"`
	Availability NamedAPIResource   `json:"availability"`
	IsRepeatable bool               `json:"is_repeatable"`
	Completion   *QuestCompletion   `json:"completion"`
	Subquests    []QuestAPIResource `json:"subquests"`
}

type Subquest struct {
	ID              int32             `json:"id"`
	Name            string            `json:"name"`
	UntypedQuest    QuestAPIResource  `json:"untyped_quest"`
	Availability    NamedAPIResource  `json:"availability"`
	IsRepeatable    bool              `json:"is_repeatable"`
	ParentSidequest QuestAPIResource  `json:"parent_sidequest"`
	Completion      QuestCompletion   `json:"completion"`
	ArenaCreation   *NamedAPIResource `json:"arena_creation,omitempty"`
}

type QuestCompletion struct {
	Condition    *string                          `json:"condition"`
	IsRepeatable bool                             `json:"is_repeatable"`
	Areas        []CompletionArea                 `json:"areas"`
	Reward       ResourceAmount[TypedAPIResource] `json:"reward"`
}

func convertQuestCompletion(cfg *Config, qc seeding.QuestCompletion) QuestCompletion {
	return QuestCompletion{
		Condition: qc.Condition,
		Areas:     convertObjSlice(cfg, qc.Areas, convertCompletionArea),
		Reward:    nameAmountToResourceAmount(cfg, cfg.e.allItems, qc.Reward),
	}
}

type CompletionArea struct {
	Area  AreaAPIResource `json:"area"`
	Notes *string         `json:"notes,omitempty"`
}

func (ca CompletionArea) GetAPIResource() APIResource {
	return ca.Area
}

func convertCompletionArea(cfg *Config, cl seeding.CompletionArea) CompletionArea {
	return CompletionArea{
		Area:  locAreaToAreaAPIResource(cfg, cfg.e.areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}
