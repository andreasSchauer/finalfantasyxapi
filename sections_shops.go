package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ShopSub struct {
	ID          int32       `json:"id"`
	URL         string      `json:"url"`
	Area        string      `json:"area"`
	Category	string		`json:"category"`
	Notes       *string     `json:"notes,omitempty"`
	PreAirship	*SubShopSub	`json:"pre_airship"`
	PostAirship	*SubShopSub	`json:"post_airship"`
}

func (s ShopSub) GetURL() string {
	return s.URL
}

type SubShopSub struct {
	Items     []string  			`json:"items"`
	Equipment []ShopEquipmentSub 	`json:"equipment"`
}

func convertSubShopSub(cfg *Config, ss seeding.SubShop) SubShopSub {
	return SubShopSub{
		Items: 		convertObjSlice(cfg, ss.Items, shopItemNameString),
		Equipment: 	convertObjSlice(cfg, ss.Equipment, convertShopEquipmentSub),
	}
}

type ShopEquipmentSub struct {
	EquipmentName	string		`json:"equipment_name"`
	Abilities		*string		`json:"abilities"` // needs to be pointer
}

func convertShopEquipmentSub(cfg *Config, se seeding.ShopEquipment) ShopEquipmentSub {
	return ShopEquipmentSub{
		EquipmentName: 	shopEquipmentNameString(cfg, se),
		Abilities: 		foundEquipmentAbilitiesStringPtr(se.FoundEquipment),
	}
}

func shopItemNameString(_ *Config, si seeding.ShopItem) string {
	return fmt.Sprintf("%s (%d Gil)", si.Name, si.Price) 
}

func shopEquipmentNameString(_ *Config, se seeding.ShopEquipment) string {
	return fmt.Sprintf("%s (%d Gil)", se.Name, se.Price)
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

func handleShopsSection(cfg *Config, _ *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.shops
	shops := []ShopSub{}

	for _, shopID := range dbIDs {
		shop, _ := seeding.GetResourceByID(shopID, i.objLookupID)

		shopSub := ShopSub{
			ID:            	shop.ID,
			URL:           	createResourceURL(cfg, i.endpoint, shopID),
			Area:          	idToLocAreaString(cfg, shop.AreaID),
			Category: 		shop.Category,
			Notes:         	shop.Notes,
			PreAirship: 	convertObjPtr(cfg, shop.PreAirship, convertSubShopSub),
			PostAirship: 	convertObjPtr(cfg, shop.PostAirship, convertSubShopSub),
		}

		shops = append(shops, shopSub)
	}

	return toSubResourceSlice(shops), nil
}
