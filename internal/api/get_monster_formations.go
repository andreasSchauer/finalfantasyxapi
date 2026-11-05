package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getMonsterFormation(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource], id int32) (MonsterFormation, error) {
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
		FormationLoot:   getFormationLoot(formation),
		BossMusic:       convertObjPtr(srv, formation.FormationData.BossMusic, convertFormationBossSong),
		Monsters:        E.Monsters.toResAmtsFromNames(srv, formation.Monsters),
		Areas:           convertObjSlice(srv, formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: convertObjSlice(srv, formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func (srv *Server) retrieveMonsterFormations(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetMonsterFormationIDsByMonster),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetMonsterFormationIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetMonsterFormationIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetMonsterFormationIDsByArea),
		i.queryBool(r, ids, ts.QpnAmbush, config.Cfg.DB.GetMonsterFormationIDsByForcedAmbush),
		i.queryBool(r, ids, ts.QpnEscape, config.Cfg.DB.GetMonsterFormationIDsByCanEscape),
		i.queryEnumList(srv, r, enums.T.MonsterFormationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMonsterFormationIDsByCategory),
	})
}
