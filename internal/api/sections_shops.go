package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertSubShopSimple(ss seedtypes.SubShop) sectiontypes.SubShopSimple {
	return sectiontypes.SubShopSimple{
		Items:     ah.ConvertObjSlice(ss.Items, shopItemNameString),
		Equipment: ah.ConvertObjSlice(ss.Equipment, convertShopEquipmentSimple),
	}
}

func convertShopEquipmentSimple(se seedtypes.ShopEquipment) sectiontypes.ShopEquipmentSimple {
	return sectiontypes.ShopEquipmentSimple{
		EquipmentName: shopEquipmentNameString(se),
		Abilities:     foundEquipmentAbilitiesStringPtr(se.TreasureEquipment),
	}
}

func shopItemNameString(si seedtypes.ShopItem) string {
	return fmt.Sprintf("%s - %d Gil", si.Name, si.Price)
}

func shopEquipmentNameString(se seedtypes.ShopEquipment) string {
	return fmt.Sprintf("%s - %d Gil", se.Name, se.Price)
}

func foundEquipmentAbilitiesStringPtr(fe seedtypes.TreasureEquipment) *string {
	zeroAbilities := len(fe.Abilities) == 0
	zeroSlots := fe.EmptySlotsAmount == 0
	slotsStr := fmt.Sprintf("(%d)", fe.EmptySlotsAmount)
	abilitiesStr := h.StringSliceToListString(fe.Abilities)

	switch {
	case zeroAbilities && zeroSlots:
		return nil

	case zeroAbilities:
		return &slotsStr

	case zeroSlots:
		return &abilitiesStr

	default:
		s := fmt.Sprintf("%s, %s", abilitiesStr, slotsStr)
		return &s
	}
}

func createShopSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Shops
	shop, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	shopSimple := sectiontypes.ShopSimple{
		ID:          shop.ID,
		URL:         ah.CreateResourceURL(i.Endpoint, id),
		Area:        idToLocAreaString(shop.AreaID),
		Category:    shop.Category,
		Notes:       shop.Notes,
		PreAirship:  ah.ConvertObjPtr(shop.PreAirship, convertSubShopSimple),
		PostAirship: ah.ConvertObjPtr(shop.PostAirship, convertSubShopSimple),
	}

	return shopSimple, nil
}
