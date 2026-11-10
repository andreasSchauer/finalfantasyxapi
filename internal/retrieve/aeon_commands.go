package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func RetrieveAeonCommands(r *http.Request, i handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
