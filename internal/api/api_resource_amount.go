package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type ResourceAmount[A APIResource] struct {
	Resource A     `json:"resource"`
	Amount   int32 `json:"amount"`
}

func (ra ResourceAmount[A]) GetAPIResource() APIResource {
	return ra.Resource
}

func newResourceAmount[A APIResource](resource A, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: resource,
		Amount:   amount,
	}
}

func (i handlerInput[T, R, A]) toResAmt(server *Server, id, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: i.idToResFunc(server, i, id),
		Amount:   amount,
	}
}

func (i handlerInput[T, R, A]) toResAmtFromNamePtr[NA NameAmount](server *Server, itemPtr *NA) *ResourceAmount[A] {
	if itemPtr == nil {
		return nil
	}

	resAmt := i.toResAmtFromName(server, *itemPtr)

	return &resAmt
}

func (i handlerInput[T, R, A]) toResAmtFromName[NA NameAmount](server *Server, item NA) ResourceAmount[A] {
	var parseResp parseResponse
	switch item.GetVersion() {
	case nil:
		parseResp, _ = checkUniqueName(item.GetName(), i.objLookup)
	default:
		parseResp, _ = checkNameVersion(item.GetName(), item.GetVersion(), i.objLookup)
	}

	return ResourceAmount[A]{
		Resource: i.idToResFunc(server, i, parseResp.ID),
		Amount:   item.GetVal(),
	}
}

func (i handlerInput[T, R, A]) toResAmtsFromNames[NA NameAmount](server *Server, items []NA) []ResourceAmount[A] {
	results := []ResourceAmount[A]{}

	for _, item := range items {
		ra := i.toResAmtFromName(server, item)
		results = append(results, ra)
	}

	return results
}

func (i handlerInput[T, R, A]) toResAmtType[NA NameAmount, RA ResourceAmountType[A]](server *Server, old []NA, fn func(A, int32) RA) []RA {
	resAmts := []RA{}

	for _, item := range old {
		ra := i.toResAmtFromName(server, item)
		resAmtType := fn(ra.Resource, ra.Amount)
		resAmts = append(resAmts, resAmtType)
	}

	return resAmts
}

func GetResAmtTypeMap[T ResourceAmountType[A], A APIResource](items []T) map[string]int32 {
	resAmts := ResAmtTypesToStructs(items)
	return GetResAmtMap(resAmts)
}

func GetResAmtMap[A APIResource](items []ResourceAmount[A]) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.Resource.GetTestKey()
		amountMap[key] = item.Amount
	}

	return amountMap
}

func GetResAmtIDMap[A APIResource](items []ResourceAmount[A]) map[int32]int32 {
	amountMap := make(map[int32]int32)

	for _, item := range items {
		url := item.Resource.GetURL()
		id := getIDFromURL(url)
		amountMap[id] = item.Amount
	}

	return amountMap
}

func ResAmtTypesToStructs[T ResourceAmountType[A], A APIResource](items []T) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, item := range items {
		resAmts = append(resAmts, item.ToResAmount())
	}

	return resAmts
}

func resAmtMapToFloat(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64)

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}

	return floatMap
}

// takes a slice of parent resources and for each of them, attaches the amount of the resource's seedtypes.ItemAmount field to the resource.
func itemAmtsToChildResAmts[T seedtypes.HasItemAmount, R any, A APIResource](i handlerInput[T, R, A], resources []A) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, res := range resources {
		lookup, _ := lookup.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmount := lookup.GetItemAmount()
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
