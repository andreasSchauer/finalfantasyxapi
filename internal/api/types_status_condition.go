package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type StatusCondition struct {
	ID                  	int32             	`json:"id"`
	Name                	string              `json:"name"`
	Category				string			 	`json:"category"`
	IsPermanent				bool			 	`json:"is_permanent"`
	Effect                  string           	`json:"effect"`
	Visualization           *string          	`json:"visualization"`
	RelatedStats            []NamedAPIResource	`json:"related_stats"`
	RemovedStatusConditions []NamedAPIResource  `json:"removed_status_conditions"`
	AddedElemResist         *ElementalResist 	`json:"added_elem_resist"`
	CtbOnInfliction			*InflictedDelay	 	`json:"ctb_on_infliction"`
	NullifyArmored          *string          	`json:"h.Nullify_armored"`
	StatChanges             []StatChange     	`json:"stat_changes"`
	ModifierChanges         []ModifierChange 	`json:"modifier_changes"`
	AutoAbilities			[]NamedAPIResource	`json:"auto_abilities"`
	InflictedBy				StatusInteractions	`json:"inflicted_by"`
	RemovedBy				StatusInteractions	`json:"removed_by"`
	MonstersResist			[]NamedAPIResource	`json:"monsters_resist"`
	MonstersImmune			[]NamedAPIResource	`json:"monsters_immune"`
}

type StatusInteractions struct {
	PlayerAbilities			[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities		[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities			[]NamedAPIResource	`json:"item_abilities"`
	UnspecifiedAbilities	[]NamedAPIResource	`json:"unspecified_abilities"`
	EnemyAbilities			[]NamedAPIResource	`json:"enemy_abilities"`
	StatusConditions		[]NamedAPIResource	`json:"status_conditions,omitempty"`
}

type StatusInteractionQueries struct {
	PlayerAbilities			DbQueryIntMany
	OverdriveAbilities		DbQueryIntMany
	ItemAbilities			DbQueryIntMany
	UnspecifiedAbilities	DbQueryIntMany
	EnemyAbilities			DbQueryIntMany
	StatusConditions 		DbQueryIntMany
}

func getStatusInteractions(cfg *Config, r *http.Request, condition seeding.StatusCondition, queries StatusInteractionQueries) (StatusInteractions, error) {
	var err error
	statusInteractions := StatusInteractions{}

	statusInteractions.PlayerAbilities, err = getResourcesDbItem(cfg, r, cfg.e.playerAbilities, condition, queries.PlayerAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.OverdriveAbilities, err = getResourcesDbItem(cfg, r, cfg.e.overdriveAbilities, condition, queries.OverdriveAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.ItemAbilities, err = getResourcesDbItem(cfg, r, cfg.e.itemAbilities, condition, queries.ItemAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.UnspecifiedAbilities, err = getResourcesDbItem(cfg, r, cfg.e.unspecifiedAbilities, condition, queries.UnspecifiedAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.EnemyAbilities, err = getResourcesDbItem(cfg, r, cfg.e.enemyAbilities, condition, queries.EnemyAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	if queries.StatusConditions != nil {
		statusInteractions.StatusConditions, err = getResourcesDbItem(cfg, r, cfg.e.statusConditions, condition, queries.StatusConditions)
		if err != nil {
			return StatusInteractions{}, err
		}
	}

	return statusInteractions, nil
}