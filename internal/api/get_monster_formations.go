package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonsterFormation(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource], id int32) (MonsterFormation, error) {
	formation, err := i.queryParamsVerifyGet(r, id)
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
		Monsters:        cfg.E.Monsters.toResAmtsFromNames(cfg, formation.Monsters),
		Areas:           convertObjSlice(cfg, formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: convertObjSlice(cfg, formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func (cfg *Config) retrieveMonsterFormations(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetMonsterFormationIDsByMonster),
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetMonsterFormationIDsByLocation),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetMonsterFormationIDsBySublocation),
		i.queryID(r, ids, QpnArea, cfg.L.Areas, cfg.db.GetMonsterFormationIDsByArea),
		i.queryBool(r, ids, QpnAmbush, cfg.db.GetMonsterFormationIDsByForcedAmbush),
		i.queryBool(r, ids, QpnEscape, cfg.db.GetMonsterFormationIDsByCanEscape),
		i.queryEnumList(cfg, r, cfg.T.MonsterFormationCategory, ids, QpnCategory, cfg.db.GetMonsterFormationIDsByCategory),
	})
}
