package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMonsterFormation(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, apiresource.UnnamedAPIResource], id int32) (MonsterFormation, error) {
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
		BossMusic:       ah.ConvertObjPtr(formation.FormationData.BossMusic, convertFormationBossSong),
		Monsters:        E.Monsters.toResAmtsFromNames(formation.Monsters),
		Areas:           ah.ConvertObjSlice(formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: ah.ConvertObjSlice(formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func retrieveMonsterFormations(r *http.Request, i handlerInput[seedtypes.MonsterFormation, MonsterFormation, apiresource.UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetMonsterFormationIDsByMonster),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetMonsterFormationIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetMonsterFormationIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetMonsterFormationIDsByArea),
		i.queryBool(r, ids, ts.QpnAmbush, config.Cfg.DB.GetMonsterFormationIDsByForcedAmbush),
		i.queryBool(r, ids, ts.QpnEscape, config.Cfg.DB.GetMonsterFormationIDsByCanEscape),
		i.queryEnumList(r, enums.T.MonsterFormationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMonsterFormationIDsByCategory),
	})
}
