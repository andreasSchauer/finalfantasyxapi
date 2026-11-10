package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetShop(r *http.Request, i handlerinput.Main[seedtypes.Shop, apitypes.Shop, apiresource.Unnamed], id int32) (apitypes.Shop, error) {
	shop, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Shop{}, err
	}

	response := apitypes.Shop{
		ID:          shop.ID,
		Area:        handlerinput.IDToAreaAPIResource(apilookups.E.Areas, shop.AreaID),
		Notes:       shop.Notes,
		Category:    shop.Category,
		PreAirship:  ah.ConvertObjPtr(shop.PreAirship, apiconv.ConvertSubShop),
		PostAirship: ah.ConvertObjPtr(shop.PostAirship, apiconv.ConvertSubShop),
	}

	return response, nil
}
