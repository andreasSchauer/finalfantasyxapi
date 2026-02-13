package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type Sidequest struct {
	ID         int32              `json:"id"`
	Name       string             `json:"name"`
	Completion *QuestCompletion   `json:"completion"`
	Subquests  []NamedAPIResource `json:"subquests"`
}


type Subquest struct {
	ID              int32             `json:"id"`
	Name            string            `json:"name"`
	ParentSidequest NamedAPIResource  `json:"parent_sidequest"`
	Completions     []QuestCompletion `json:"completions"`
}



type QuestCompletion struct {
	Areas  []CompletionArea `json:"areas"`
	Reward ItemAmount       `json:"reward"`
}

func convertQuestCompletion(cfg *Config, qc seeding.QuestCompletion) QuestCompletion {
	return QuestCompletion{
		Areas:  convertObjSlice(cfg, qc.Areas, convertCompletionArea),
		Reward: convertItemAmount(cfg, qc.Reward),
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