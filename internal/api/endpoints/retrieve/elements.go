package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Elements(r *http.Request, i handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
