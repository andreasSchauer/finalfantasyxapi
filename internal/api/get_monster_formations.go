package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonsterFormation(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedAPIResourceList], id int32) (MonsterFormation, error) {
	formation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return MonsterFormation{}, err
	}

	response := MonsterFormation{
		ID:              formation.ID,
		Category:        formation.FormationData.Category,
		Availability:    formation.FormationData.Availability,
		IsForcedAmbush:  formation.FormationData.IsForcedAmbush,
		CanEscape:       formation.FormationData.CanEscape,
		Notes:           formation.FormationData.Notes,
		FormationLoot:   getFormationLoot(cfg, formation),
		BossMusic:       convertObjPtr(cfg, formation.FormationData.BossMusic, convertFormationBossSong),
		Monsters:        nameAmtsToResAmts(cfg, cfg.E.Monsters, formation.Monsters),
		Areas:           convertObjSlice(cfg, formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: convertObjSlice(cfg, formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func (cfg *Config) retrieveMonsterFormations(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetMonsterFormationIDsByMonster),
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetMonsterFormationIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetMonsterFormationIDsBySublocation),
		idQuery(r, i, ids, QpnArea, cfg.L.Areas, cfg.db.GetMonsterFormationIDsByArea),
		boolQuery(r, i, ids, QpnAmbush, cfg.db.GetMonsterFormationIDsByForcedAmbush),
		boolQuery(r, i, ids, QpnEscape, cfg.db.GetMonsterFormationIDsByCanEscape),
		enumListQuery(cfg, r, i, cfg.T.MonsterFormationCategory, ids, QpnCategory, cfg.db.GetMonsterFormationIDsByCategory),
	})
}
