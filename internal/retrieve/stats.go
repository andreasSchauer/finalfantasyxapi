package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func RetrieveStats(r *http.Request, i handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
