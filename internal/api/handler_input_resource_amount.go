package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (i handlerInput[T, R, A]) toResAmt(id, amount int32) apiresource.ResourceAmount[A] {
	return apiresource.ResourceAmount[A]{
		Resource: i.idToResFunc(i, id),
		Amount:   amount,
	}
}

func (i handlerInput[T, R, A]) toResAmtFromNamePtr[NA apiresource.NameAmount](itemPtr *NA) *apiresource.ResourceAmount[A] {
	if itemPtr == nil {
		return nil
	}

	resAmt := i.toResAmtFromName(*itemPtr)

	return &resAmt
}

func (i handlerInput[T, R, A]) toResAmtFromName[NA apiresource.NameAmount](item NA) apiresource.ResourceAmount[A] {
	var parseResp ah.ParseResponse
	switch item.GetVersion() {
	case nil:
		parseResp, _ = ah.CheckUniqueName(item.GetName(), i.objLookup)
	default:
		parseResp, _ = ah.CheckNameVersion(item.GetName(), item.GetVersion(), i.objLookup)
	}

	return apiresource.ResourceAmount[A]{
		Resource: i.idToResFunc(i, parseResp.ID),
		Amount:   item.GetVal(),
	}
}

func (i handlerInput[T, R, A]) toResAmtsFromNames[NA apiresource.NameAmount](items []NA) []apiresource.ResourceAmount[A] {
	results := []apiresource.ResourceAmount[A]{}

	for _, item := range items {
		ra := i.toResAmtFromName(item)
		results = append(results, ra)
	}

	return results
}

func (i handlerInput[T, R, A]) toResAmtType[NA apiresource.NameAmount, RA apiresource.ResourceAmountType[A]](old []NA, fn func(A, int32) RA) []RA {
	resAmts := []RA{}

	for _, item := range old {
		ra := i.toResAmtFromName(item)
		resAmtType := fn(ra.Resource, ra.Amount)
		resAmts = append(resAmts, resAmtType)
	}

	return resAmts
}

// takes a slice of parent resources and for each of them, attaches the amount of the resource's seedtypes.ItemAmount field to the resource.
func itemAmtsToChildResAmts[T seedtypes.HasItemAmount, R any, A apiresource.APIResource](i handlerInput[T, R, A], resources []A) []apiresource.ResourceAmount[A] {
	resAmts := []apiresource.ResourceAmount[A]{}

	for _, res := range resources {
		lookupRes, _ := lookup.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmount := lookupRes.GetItemAmount()
		resAmt := apiresource.NewResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

// takes a slice of parent resources and for each of them, searches the resource's seedtypes.ItemAmount slice for the target item and attaches its amount to the resource.
func itemAmtsToChildResAmts2[T seedtypes.HasItemAmounts, R any, A apiresource.APIResource](i handlerInput[T, R, A], resources []A, itemID int32) []apiresource.ResourceAmount[A] {
	resAmts := []apiresource.ResourceAmount[A]{}

	for _, res := range resources {
		item, _ := lookup.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmounts := item.GetItemAmounts()
		var itemAmount seedtypes.ItemAmount

		for _, ia := range itemAmounts {
			item, _ := lookup.GetResource(ia.ItemName, lookup.L.Items)
			if item.ID == itemID {
				itemAmount = ia
			}
		}

		resAmt := apiresource.NewResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}
