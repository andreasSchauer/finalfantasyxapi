package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getKeyItem(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource], id int32) (KeyItem, error) {
	keyItem, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return KeyItem{}, err
	}

	rel, err := getKeyItemRelationships(s, r, keyItem)
	if err != nil {
		return KeyItem{}, err
	}

	response := KeyItem{
		ID:              keyItem.ID,
		Name:            keyItem.Name,
		UntypedItem:     idToTypedAPIResource(s, E.AllItems, keyItem.MasterItem.ID),
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

func (s *Server) retrieveKeyItems(r *http.Request, i handlerInput[seedtypes.KeyItem, KeyItem, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.KeyItemCategory, ids, ts.QpnCategory, s.Cfg.DB.GetKeyItemIDsCategory),
		i.queryValueList(s, r, ids, ts.QpnMethods, s.Cfg.DB.GetKeyItemIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, E.Locations.objLookup, s.Cfg.DB.GetKeyItemIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, E.Sublocations.objLookup, s.Cfg.DB.GetKeyItemIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, E.Areas.objLookup, s.Cfg.DB.GetKeyItemIDsByArea),
	})
}
