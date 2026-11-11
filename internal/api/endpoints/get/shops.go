package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Shop(r *http.Request, i handlerinput.Main[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed], id int32) (apitypes.Shop, error) {
	shop, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Shop{}, err
	}

	response := apitypes.Shop{
		ID:          shop.ID,
		Area:        handlerinput.IDToAreaAPIResource(apilookups.E.Areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  ah.ConvertObjPtr(shop.PreAirship, apiconvert.SubShop),
		PostAirship: ah.ConvertObjPtr(shop.PostAirship, apiconvert.SubShop),
	}

	return response, nil
}
