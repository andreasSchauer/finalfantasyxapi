package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/filters"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetMasterItem(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed], id int32) (apitypes.MasterItem, error) {
	masterItem, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	rel, err := assemble.GetMasterItemRelationships(r, masterItem)
	if err != nil {
		return apitypes.MasterItem{}, err
	}

	response := apitypes.MasterItem{
		ID:             masterItem.ID,
		Name:           masterItem.Name,
		Type:           masterItem.Type,
		ObtainableFrom: rel.ObtainableFrom,
	}

	return assemble.CompleteMasterItemResponse(r, response), nil
}

func RetrieveMasterItems(r *http.Request, i handlerinput.Main[seedtypes.MasterItem, apitypes.MasterItem, apiresource.Typed]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemType, ids, ts.QpnType, config.Cfg.DB.GetMasterItemIDsByType),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetMasterItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, filters.GetMasterItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, filters.GetMasterItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, filters.GetMasterItemsByArea),
	})
}
