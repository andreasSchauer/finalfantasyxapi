package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getKeyItem(r *http.Request, i HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource], id int32) (apitypes.KeyItem, error) {
	keyItem, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.KeyItem{}, err
	}

	rel, err := getKeyItemRelationships(r, keyItem)
	if err != nil {
		return apitypes.KeyItem{}, err
	}

	response := apitypes.KeyItem{
		ID:              keyItem.ID,
		Name:            keyItem.Name,
		UntypedItem:     IDToTypedAPIResource(E.AllItems, keyItem.MasterItem.ID),
		Category:        keyItem.Category,
		Description:     keyItem.Description,
		Effect:          keyItem.Effect,
		Primer:          rel.Primer,
		CelestialWeapon: rel.CelestialWeapon,
		Treasures:       rel.Treasures,
		Quests:          rel.Quests,
		Areas:           rel.Areas,
	}

	return response, nil
}

func retrieveKeyItems(r *http.Request, i HandlerInput[seedtypes.KeyItem, apitypes.KeyItem, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnumList(r, enums.T.KeyItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetKeyItemIDsCategory),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetKeyItemIDsByMethods),
		i.QueryID(r, ids, ts.QpnLocation, E.Locations.objLookup, config.Cfg.DB.GetKeyItemIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, E.Sublocations.objLookup, config.Cfg.DB.GetKeyItemIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, E.Areas.objLookup, config.Cfg.DB.GetKeyItemIDsByArea),
	})
}
