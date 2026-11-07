package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMasterItem(r *http.Request, i HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource], id int32) (apitypes.MasterItem, error) {
	masterItem, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	obtainable, err := getMasterItemObtainable(r, masterItem)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	response := apitypes.MasterItem{
		ID:             masterItem.ID,
		Name:           masterItem.Name,
		Type:           masterItem.Type,
		ObtainableFrom: obtainable,
	}

	switch masterItem.Type {
	case database.ItemTypeItem:
		item, _ := lookup.GetResource(masterItem, lookup.L.Items)

		response.Description = item.Description
		response.Effect = item.Effect
		response.TypedItem = nameToNamedAPIResource(E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := lookup.GetResource(masterItem, lookup.L.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = nameToNamedAPIResource(E.KeyItems, keyItem.Name, nil)
	}

	return response, nil
}

func retrieveMasterItems(r *http.Request, i HandlerInput[seedtypes.MasterItem, apitypes.MasterItem, apiresource.TypedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnumList(r, enums.T.ItemType, ids, ts.QpnType, config.Cfg.DB.GetMasterItemIDsByType),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetMasterItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, E.Locations.objLookup, getMasterItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, E.Sublocations.objLookup, getMasterItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, E.Areas.objLookup, getMasterItemsByArea),
	})
}
