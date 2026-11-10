package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveMonsterFormations(r *http.Request, i handlerinput.Main[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetMonsterFormationIDsByMonster),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetMonsterFormationIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetMonsterFormationIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetMonsterFormationIDsByArea),
		i.QueryBool(r, ids, ts.QpnAmbush, config.Cfg.DB.GetMonsterFormationIDsByForcedAmbush),
		i.QueryBool(r, ids, ts.QpnEscape, config.Cfg.DB.GetMonsterFormationIDsByCanEscape),
		i.QueryEnumList(r, enums.T.MonsterFormationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMonsterFormationIDsByCategory),
	})
}
