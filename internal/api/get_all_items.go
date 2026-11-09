package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMasterItem(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed], id int32) (apitypes.MasterItem, error) {
	masterItem, err := i.VerifyQueriesGet(r, id)
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
		response.TypedItem = handlerinput.NameToNamedAPIResource(apilookups.E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := lookup.GetResource(masterItem, lookup.L.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = handlerinput.NameToNamedAPIResource(apilookups.E.KeyItems, keyItem.Name, nil)
	}

	return response, nil
}

func retrieveMasterItems(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemType, ids, ts.QpnType, config.Cfg.DB.GetMasterItemIDsByType),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetMasterItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, getMasterItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, getMasterItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, getMasterItemsByArea),
	})
}
