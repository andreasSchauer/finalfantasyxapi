package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func OverdriveCommands(r *http.Request, i handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
