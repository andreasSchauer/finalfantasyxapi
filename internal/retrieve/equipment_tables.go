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

func RetrieveEquipmentTables(r *http.Request, i handlerinput.Main[seedtypes.EquipmentTable, apitypes.EquipmentTable, apiresource.Unnamed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryIDList(r, ids, ts.QpnAutoAbilities, lookup.L.AutoAbilities, config.Cfg.DB.GetEquipmentTableIDsByAutoAbility),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetEquipmentTableIDsEquipType),
		i.QueryBool2(r, ids, ts.QpnCelestialWeapon, config.Cfg.DB.GetEquipmentTableIDsCelestialWeapon),
	})
}
