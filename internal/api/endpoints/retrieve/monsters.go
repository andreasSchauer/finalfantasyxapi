package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func RetrieveMonsters(r *http.Request, i handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryWrapper(r, ids, ts.QpnElementalResists, GetMonstersByElemResists),
		i.QueryIDListWrapper(r, ids, ts.QpnStatusResists, lookup.L.StatusConditions, GetMonstersByStatusResists),

		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, GetMonstersByItem),
		i.QueryID(r, ids, ts.QpnRonsoRage, lookup.L.RonsoRages, config.Cfg.DB.GetMonsterIDsByRonsoRage),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetMonsterIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetMonsterIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetMonsterIDsByArea),
		i.QueryIDWrapper(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, GetMonstersByAutoAbility),

		i.QueryIntList(r, ids, ts.QpnEmptySlots, config.Cfg.DB.GetMonsterIDsByEmptySlots),
		i.QueryIntList(r, ids, ts.QpnDistance, config.Cfg.DB.GetMonsterIDsByDistance),

		i.QueryEnumList(r, enums.T.MonsterCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMonsterIDsByCategory),
		i.QueryEnum(r, enums.T.MonsterSpecies, ids, ts.QpnSpecies, config.Cfg.DB.GetMonsterIDsBySpecies),
		i.QueryEnum(r, enums.T.CreationArea, ids, ts.QpnCreationArea, handlerinput.ToEnumQuery(enums.T.CreationArea, config.Cfg.DB.GetMonsterIDsByMaCreationArea)),

		i.QueryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetMonsterIDsByIsRepeatable),
		i.QueryBool(r, ids, ts.QpnCapture, config.Cfg.DB.GetMonsterIDsByCanBeCaptured),
		i.QueryBool(r, ids, ts.QpnHasOverdrive, config.Cfg.DB.GetMonsterIDsByHasOverdrive),
		i.QueryBool(r, ids, ts.QpnUnderwater, config.Cfg.DB.GetMonsterIDsByIsUnderwater),
		i.QueryBool(r, ids, ts.QpnZombie, config.Cfg.DB.GetMonsterIDsByIsZombie),
	})
}
