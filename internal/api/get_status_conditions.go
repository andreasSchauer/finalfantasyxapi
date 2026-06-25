package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getStatusCondition(r *http.Request, i handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList], id int32) (StatusCondition, error) {
	status, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return StatusCondition{}, err
	}

	rel, err := getStatusConditionRelationships(cfg, r, status)
	if err != nil {
		return StatusCondition{}, err
	}

	response := StatusCondition{
		ID:                      status.ID,
		Name:                    status.Name,
		Category:                enumToNamedAPIResource(cfg, cfg.e.statusConditionCategory.endpoint, status.Category, cfg.t.StatusConditionCategory),
		IsPermanent:             status.IsPermanent,
		Visualization:           status.Visualization,
		Effect:                  status.Effect,
		NullifyArmored:          status.NullifyArmored,
		RelatedStats:            namesToNamedAPIResources(cfg, cfg.e.stats, status.RelatedStats),
		RemovedStatusConditions: namesToNamedAPIResources(cfg, cfg.e.statusConditions, status.RemovedStatusConditions),
		AddedElemResist:         convertObjPtr(cfg, status.AddedElemResist, convertElemResist),
		CtbOnInfliction:         convertObjPtr(cfg, status.CtbOnInfliction, convertInflictedDelay),
		StatChanges:             convertObjSlice(cfg, status.StatChanges, convertStatChange),
		ModifierChanges:         convertObjSlice(cfg, status.ModifierChanges, convertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func (cfg *Config) retrieveStatusConditions(r *http.Request, i handlerInput[seeding.StatusCondition, StatusCondition, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.StatusConditionCategory, ids, qpnCategory, cfg.db.GetStatusConditionIDsByCategory)),
	})
}
