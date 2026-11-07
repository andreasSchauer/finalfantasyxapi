package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertSubShop(ss seedtypes.SubShop) apitypes.SubShop {
	return apitypes.SubShop{
		Items:     ah.ConvertObjSlice(ss.Items, convertShopItem),
		Equipment: ah.ConvertObjSlice(ss.Equipment, convertShopEquipment),
	}
}

func convertShopItem(si seedtypes.ShopItem) apitypes.ShopItem {
	return apitypes.ShopItem{
		Item:  handlerinput.NameToNamedAPIResource(E.Items, si.Name, nil),
		Price: si.Price,
	}
}

func convertShopEquipment(se seedtypes.ShopEquipment) apitypes.ShopEquipment {
	return apitypes.ShopEquipment{
		Equipment: convertFoundEquipment(se.TreasureEquipment),
		Price:     se.Price,
	}
}
