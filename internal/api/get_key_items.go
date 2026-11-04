package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getKeyItem(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource], id int32) (KeyItem, error) {
	keyItem, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return KeyItem{}, err
	}

	rel, err := getKeyItemRelationships(srv, r, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	response := KeyItem{
		ID:              keyItem.ID,
		Name:            keyItem.Name,
		UntypedItem:     idToTypedAPIResource(srv, E.AllItems, keyItem.MasterItem.ID),
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

func (srv *Server) retrieveKeyItems(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.KeyItemCategory, ids, ts.QpnCategory, srv.Cfg.DB.GetKeyItemIDsCategory),
		i.queryValueList(srv, r, ids, ts.QpnMethods, srv.Cfg.DB.GetKeyItemIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, E.Locations.objLookup, srv.Cfg.DB.GetKeyItemIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, E.Sublocations.objLookup, srv.Cfg.DB.GetKeyItemIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, E.Areas.objLookup, srv.Cfg.DB.GetKeyItemIDsByArea),
	})
}
