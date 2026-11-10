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

func RetrieveAutoAbilities(r *http.Request, i handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetAutoAbilityIDsByCategory),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetAutoAbilityIDsByEquipType),
		i.QueryIDWrapper(r, ids, ts.QpnMonster, lookup.L.Monsters, GetAutoAbilitiesByMonster),
		i.QueryID(r, ids, "monster_items", lookup.L.Monsters, config.Cfg.DB.GetAutoAbilityIDsByMonsterItems),
		i.QueryIDWrapper(r, ids, ts.QpnShop, lookup.L.Shops, GetAutoAbilitiesByShop),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetAutoAbilityIDsByMethods),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetAutoAbilityIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetAutoAbilityIDsBySublocation),
		i.QueryID(r, ids, "areas", lookup.L.Areas, config.Cfg.DB.GetAutoAbilityIDsByArea),
	})
}
