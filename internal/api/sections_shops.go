package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type ShopSimple struct {
	ID          int32          `json:"id"`
	URL         string         `json:"url"`
	Area        string         `json:"area"`
	Category    string         `json:"category"`
	Notes       *string        `json:"notes,omitempty"`
	PreAirship  *SubShopSimple `json:"pre_airship"`
	PostAirship *SubShopSimple `json:"post_airship"`
}

func (s ShopSimple) GetURL() string {
	return s.URL
}

type SubShopSimple struct {
	Items     []string              `json:"items"`
	Equipment []ShopEquipmentSimple `json:"equipment"`
}

func convertSubShopSimple(ss seedtypes.SubShop) SubShopSimple {
	return SubShopSimple{
		Items:     ConvertObjSlice(ss.Items, shopItemNameString),
		Equipment: ConvertObjSlice(ss.Equipment, convertShopEquipmentSimple),
	}
}

type ShopEquipmentSimple struct {
	EquipmentName string  `json:"equipment_name"`
	Abilities     *string `json:"abilities"`
}

func convertShopEquipmentSimple(se seedtypes.ShopEquipment) ShopEquipmentSimple {
	return ShopEquipmentSimple{
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

func createShopSimple(_ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.Shops
	shop, _ := lookup.GetResourceByID(id, i.objLookupID)

	shopSimple := ShopSimple{
		ID:          shop.ID,
		URL:         CreateResourceURL(i.Endpoint, id),
		Area:        idToLocAreaString(shop.AreaID),
		Category:    shop.Category,
		Notes:       shop.Notes,
		PreAirship:  ConvertObjPtr(shop.PreAirship, convertSubShopSimple),
		PostAirship: ConvertObjPtr(shop.PostAirship, convertSubShopSimple),
	}

	return shopSimple, nil
}
