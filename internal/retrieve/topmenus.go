package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func RetrieveTopmenus(r *http.Request, i handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
