package apiconv

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ConvertSubShop(ss seedtypes.SubShop) apitypes.SubShop {
	return apitypes.SubShop{
		Items:     ah.ConvertObjSlice(ss.Items, ConvertShopItem),
		Equipment: ah.ConvertObjSlice(ss.Equipment, ConvertShopEquipment),
	}
}

func ConvertShopItem(si seedtypes.ShopItem) apitypes.ShopItem {
	return apitypes.ShopItem{
		Item:  handlerinput.NameToNamedAPIResource(apilookups.E.Items, si.Name, nil),
		Price: si.Price,
	}
}

func ConvertShopEquipment(se seedtypes.ShopEquipment) apitypes.ShopEquipment {
	return apitypes.ShopEquipment{
		Equipment: ConvertFoundEquipment(se.TreasureEquipment),
		Price:     se.Price,
	}
}
