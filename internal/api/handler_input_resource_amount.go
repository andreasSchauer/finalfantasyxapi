package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (i handlerInput[T, R, A]) toResAmt(id, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: i.idToResFunc(i, id),
		Amount:   amount,
	}
}

func (i handlerInput[T, R, A]) toResAmtFromNamePtr[NA NameAmount](itemPtr *NA) *ResourceAmount[A] {
	if itemPtr == nil {
		return nil
	}

	resAmt := i.toResAmtFromName(*itemPtr)

	return &resAmt
}

func (i handlerInput[T, R, A]) toResAmtFromName[NA NameAmount](item NA) ResourceAmount[A] {
	var parseResp parseResponse
	switch item.GetVersion() {
	case nil:
		parseResp, _ = checkUniqueName(item.GetName(), i.objLookup)
	default:
		parseResp, _ = checkNameVersion(item.GetName(), item.GetVersion(), i.objLookup)
	}

	return ResourceAmount[A]{
		Resource: i.idToResFunc(i, parseResp.ID),
		Amount:   item.GetVal(),
	}
}

func (i handlerInput[T, R, A]) toResAmtsFromNames[NA NameAmount](items []NA) []ResourceAmount[A] {
	results := []ResourceAmount[A]{}

	for _, item := range items {
		ra := i.toResAmtFromName(item)
		results = append(results, ra)
	}

	return results
}

func (i handlerInput[T, R, A]) toResAmtType[NA NameAmount, RA ResourceAmountType[A]](old []NA, fn func(A, int32) RA) []RA {
	resAmts := []RA{}

	for _, item := range old {
		ra := i.toResAmtFromName(item)
		resAmtType := fn(ra.Resource, ra.Amount)
		resAmts = append(resAmts, resAmtType)
	}

	return resAmts
}


// takes a slice of parent resources and for each of them, attaches the amount of the resource's seedtypes.ItemAmount field to the resource.
func itemAmtsToChildResAmts[T seedtypes.HasItemAmount, R any, A APIResource](i handlerInput[T, R, A], resources []A) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, res := range resources {
		lookupRes, _ := lookup.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmount := lookupRes.GetItemAmount()
		resAmt := newResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

// takes a slice of parent resources and for each of them, searches the resource's seedtypes.ItemAmount slice for the target item and attaches its amount to the resource.
func itemAmtsToChildResAmts2[T seedtypes.HasItemAmounts, R any, A APIResource](i handlerInput[T, R, A], resources []A, itemID int32) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

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

		resAmt := newResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}