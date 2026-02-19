package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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

func convertSubShopSimple(cfg *Config, ss seeding.SubShop) SubShopSimple {
	return SubShopSimple{
		Items:     convertObjSlice(cfg, ss.Items, shopItemNameString),
		Equipment: convertObjSlice(cfg, ss.Equipment, convertShopEquipmentSimple),
	}
}

type ShopEquipmentSimple struct {
	EquipmentName string  `json:"equipment_name"`
	Abilities     *string `json:"abilities"`
}

func convertShopEquipmentSimple(cfg *Config, se seeding.ShopEquipment) ShopEquipmentSimple {
	return ShopEquipmentSimple{
		EquipmentName: shopEquipmentNameString(cfg, se),
		Abilities:     foundEquipmentAbilitiesStringPtr(se.FoundEquipment),
	}
}

func shopItemNameString(_ *Config, si seeding.ShopItem) string {
	return fmt.Sprintf("%s - %d Gil", si.Name, si.Price)
}

func shopEquipmentNameString(_ *Config, se seeding.ShopEquipment) string {
	return fmt.Sprintf("%s - %d Gil", se.Name, se.Price)
}

func foundEquipmentAbilitiesStringPtr(fe seeding.FoundEquipment) *string {
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

func createShopSub(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.shops
	shop, _ := seeding.GetResourceByID(id, i.objLookupID)

	shopSimple := ShopSimple{
		ID:          shop.ID,
		URL:         createResourceURL(cfg, i.endpoint, id),
		Area:        idToLocAreaString(cfg, shop.AreaID),
		Category:    shop.Category,
		Notes:       shop.Notes,
		PreAirship:  convertObjPtr(cfg, shop.PreAirship, convertSubShopSimple),
		PostAirship: convertObjPtr(cfg, shop.PostAirship, convertSubShopSimple),
	}

	return shopSimple, nil
}
