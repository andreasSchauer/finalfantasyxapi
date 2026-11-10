package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func RetrieveCharacters(r *http.Request, i handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryBool(r, ids, ts.QpnStoryBased, config.Cfg.DB.GetCharacterIDsStoryBased),
		i.QueryBool(r, ids, ts.QpnUnderwater, config.Cfg.DB.GetCharacterIDsCanFightUnderwater),
	})
}
