package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func applyOsgStats(r *http.Request, char apitypes.Character) (apitypes.Character, error) {
	if char.IsStoryBased {
		return char, nil
	}

	var statMap map[string]int32
	queryParam := apilookups.Q.Characters[ts.QpnOsgStats]

	sgType, err := queryParam.ParseEnum(r, enums.T.SphereGridType)
	if errs.QueryIsEmpty(err) {
		return char, nil
	}
	if err != nil {
		return apitypes.Character{}, err
	}

	switch sgType.Name {
	case string(database.SphereGridTypeStandard):
		statMap = char.StdSphereGrid.ToMap()

	case string(database.SphereGridTypeExpert):
		statMap = char.ExpSphereGrid.ToMap()
	}

	char.BaseStats = th.AddToBaseStats(char.BaseStats, statMap)

	return char, nil
}
