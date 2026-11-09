package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getItem(r *http.Request, i handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named], id int32) (apitypes.Item, error) {
	item, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Item{}, err
	}

	rel, err := assemble.GetItemRelationships(r, item)
	if err != nil {
		return apitypes.Item{}, err
	}

	response := apitypes.Item{
		ID:                 item.ID,
		Name:               item.Name,
		UntypedItem:        handlerinput.IDToTypedAPIResource(apilookups.E.AllItems, item.MasterItem.ID),
		Category:           item.Category,
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		ItemAbility:        rel.ItemAbility,
		Sphere:             rel.Sphere,
		AvailableMenus:     handlerinput.NamesToNamedAPIResources(apilookups.E.Submenus, item.AvailableMenus),
		RelatedStats:       handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, item.RelatedStats),
		Monsters:           rel.Monsters,
		Treasures:          rel.Treasures,
		Shops:              rel.Shops,
		Quests:             rel.Quests,
		BlitzballPrizes:    rel.BlitzballPrizes,
		AeonLearnAbilities: rel.AeonLearnAbilities,
		AutoAbilities:      rel.AutoAbilities,
		Mixes:              rel.Mixes,
	}

	return response, nil
}

func retrieveItems(r *http.Request, i handlerinput.Main[seedtypes.Item, apitypes.Item, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemIDsCategory),
		i.QueryBool2(r, ids, ts.QpnHasAbility, config.Cfg.DB.GetItemIDsWithAbility),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemIDsByRelatedStat),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetItemIDsByMethods),
		i.QueryIDWrapper(r, ids, ts.QpnLocation, apilookups.E.Locations.ObjLookup, getItemsByLocation),
		i.QueryIDWrapper(r, ids, ts.QpnSublocation, apilookups.E.Sublocations.ObjLookup, getItemsBySublocation),
		i.QueryIDWrapper(r, ids, ts.QpnArea, apilookups.E.Areas.ObjLookup, getItemsByArea),
	})
}
