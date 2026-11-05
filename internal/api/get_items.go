package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getItem(r *http.Request, i handlerInput[seedtypes.Item, Item, NamedAPIResource], id int32) (Item, error) {
	item, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Item{}, err
	}

	var itemAbilityNamePtr *string
	if len(item.BattleInteractions) > 0 {
		itemAbilityNamePtr = &item.Name
	}

	rel, err := getItemRelationships(srv, r, item)
	if err != nil {
		return Item{}, err
	}

	response := Item{
		ID:                 item.ID,
		Name:               item.Name,
		UntypedItem:        idToTypedAPIResource(srv, E.AllItems, item.MasterItem.ID),
		Category:           item.Category,
		Description:        item.Description,
		SgDescription:      item.SphereGridDescription,
		Effect:             item.Effect,
		Usability:          item.Usability,
		BasePrice:          item.BasePrice,
		SellValue:          item.SellValue,
		ItemAbility:        namePtrToNamedAPIResPtr(srv, E.ItemAbilities, itemAbilityNamePtr, nil),
		Sphere:             rel.Sphere,
		AvailableMenus:     namesToNamedAPIResources(srv, E.Submenus, item.AvailableMenus),
		RelatedStats:       namesToNamedAPIResources(srv, E.Stats, item.RelatedStats),
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

func (srv *Server) retrieveItems(r *http.Request, i handlerInput[seedtypes.Item, Item, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.ItemCategory, ids, ts.QpnCategory, config.Cfg.DB.GetItemIDsCategory),
		i.queryBool2(r, ids, ts.QpnHasAbility, config.Cfg.DB.GetItemIDsWithAbility),
		i.queryNameID(r, ids, ts.QpnRelatedStat, E.Stats.resTypeSingle, lookup.L.Stats, config.Cfg.DB.GetItemIDsByRelatedStat),
		i.queryValueList(srv, r, ids, ts.QpnMethods, config.Cfg.DB.GetItemIDsByMethods),
		i.queryIDWrapper(srv, r, ids, ts.QpnLocation, E.Locations.objLookup, getItemsByLocation),
		i.queryIDWrapper(srv, r, ids, ts.QpnSublocation, E.Sublocations.objLookup, getItemsBySublocation),
		i.queryIDWrapper(srv, r, ids, ts.QpnArea, E.Areas.objLookup, getItemsByArea),
	})
}
