package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getMasterItem(r *http.Request, i handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource], id int32) (MasterItem, error) {
	masterItem, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return MasterItem{}, err
	}

	obtainable, err := getMasterItemObtainable(s, r, masterItem)
	if err != nil {
		return MasterItem{}, err
	}

	response := MasterItem{
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
		response.TypedItem = nameToNamedAPIResource(s, E.Items, item.Name, nil)

	case database.ItemTypeKeyItem:
		keyItem, _ := lookup.GetResource(masterItem, lookup.L.KeyItems)

		response.Description = keyItem.Description
		response.Effect = keyItem.Effect
		response.TypedItem = nameToNamedAPIResource(s, E.KeyItems, keyItem.Name, nil)
	}

	return response, nil
}

func (s *Server) retrieveMasterItems(r *http.Request, i handlerInput[seedtypes.MasterItem, MasterItem, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.ItemType, ids, ts.QpnType, s.Cfg.DB.GetMasterItemIDsByType),
		i.queryValueList(s, r, ids, ts.QpnMethods, s.Cfg.DB.GetMasterItemIDsByMethods),
		i.queryIDWrapper(s, r, ids, ts.QpnLocation, E.Locations.objLookup, getMasterItemsByLocation),
		i.queryIDWrapper(s, r, ids, ts.QpnSublocation, E.Sublocations.objLookup, getMasterItemsBySublocation),
		i.queryIDWrapper(s, r, ids, ts.QpnArea, E.Areas.objLookup, getMasterItemsByArea),
	})
}
