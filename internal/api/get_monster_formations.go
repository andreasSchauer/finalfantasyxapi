package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getMonsterFormation(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource], id int32) (MonsterFormation, error) {
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
		BossMusic:       convertObjPtr(s, formation.FormationData.BossMusic, convertFormationBossSong),
		Monsters:        E.Monsters.toResAmtsFromNames(s, formation.Monsters),
		Areas:           convertObjSlice(s, formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: convertObjSlice(s, formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func (s *Server) retrieveMonsterFormations(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetMonsterFormationIDsByMonster),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetMonsterFormationIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetMonsterFormationIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, s.Cfg.DB.GetMonsterFormationIDsByArea),
		i.queryBool(r, ids, ts.QpnAmbush, s.Cfg.DB.GetMonsterFormationIDsByForcedAmbush),
		i.queryBool(r, ids, ts.QpnEscape, s.Cfg.DB.GetMonsterFormationIDsByCanEscape),
		i.queryEnumList(s, r, enums.T.MonsterFormationCategory, ids, ts.QpnCategory, s.Cfg.DB.GetMonsterFormationIDsByCategory),
	})
}
