package handlerinput

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (i HandlerInput[T, R, A]) ToResAmt(id, amount int32) apiresource.ResourceAmount[A] {
	return apiresource.ResourceAmount[A]{
		Resource: i.IDToResFunc(i, id),
		Amount:   amount,
	}
}

func (i HandlerInput[T, R, A]) ToResAmtFromNamePtr[NA apiresource.NameAmount](itemPtr *NA) *apiresource.ResourceAmount[A] {
	if itemPtr == nil {
		return nil
	}

	resAmt := i.ToResAmtFromName(*itemPtr)

	return &resAmt
}

func (i HandlerInput[T, R, A]) ToResAmtFromName[NA apiresource.NameAmount](item NA) apiresource.ResourceAmount[A] {
	var parseResp ah.ParseResponse
	switch item.GetVersion() {
	case nil:
		parseResp, _ = ah.CheckUniqueName(item.GetName(), i.ObjLookup)
	default:
		parseResp, _ = ah.CheckNameVersion(item.GetName(), item.GetVersion(), i.ObjLookup)
	}

	return apiresource.ResourceAmount[A]{
		Resource: i.IDToResFunc(i, parseResp.ID),
		Amount:   item.GetVal(),
	}
}

func (i HandlerInput[T, R, A]) ToResAmtsFromNames[NA apiresource.NameAmount](items []NA) []apiresource.ResourceAmount[A] {
	results := []apiresource.ResourceAmount[A]{}

	for _, item := range items {
		ra := i.ToResAmtFromName(item)
		results = append(results, ra)
	}

	return results
}

func (i HandlerInput[T, R, A]) ToResAmtType[NA apiresource.NameAmount, RA apiresource.ResourceAmountType[A]](old []NA, fn func(A, int32) RA) []RA {
	resAmts := []RA{}

	for _, item := range old {
		ra := i.ToResAmtFromName(item)
		resAmtType := fn(ra.Resource, ra.Amount)
		resAmts = append(resAmts, resAmtType)
	}

	return resAmts
}

// ItemAmtsToChildResAmts takes a slice of parent resources and for each of them, attaches the amount of the resource's seedtypes.ItemAmount field to the resource.
func ItemAmtsToChildResAmts[T seedtypes.HasItemAmount, R any, A apiresource.APIResource](i HandlerInput[T, R, A], resources []A) []apiresource.ResourceAmount[A] {
	resAmts := []apiresource.ResourceAmount[A]{}

	for _, res := range resources {
		lookupRes, _ := lookup.GetResourceByID(res.GetID(), i.ObjLookupID)
		itemAmount := lookupRes.GetItemAmount()
		resAmt := apiresource.NewResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

// ItemAmtsToChildResAmts2 takes a slice of parent resources and for each of them, searches the resource's seedtypes.ItemAmount slice for the target item and attaches its amount to the resource.
func ItemAmtsToChildResAmts2[T seedtypes.HasItemAmounts, R any, A apiresource.APIResource](i HandlerInput[T, R, A], resources []A, itemID int32) []apiresource.ResourceAmount[A] {
	resAmts := []apiresource.ResourceAmount[A]{}

	for _, res := range resources {
		item, _ := lookup.GetResourceByID(res.GetID(), i.ObjLookupID)
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
