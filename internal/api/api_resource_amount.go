package api

import (
	"errors"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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


func idAmountToResourceAmount[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id, amount int32) ResourceAmount[A] {
	return ResourceAmount[A]{
		Resource: i.idToResFunc(cfg, i, id),
		Amount:   amount,
	}
}

func nameAmountPtrToResAmtPtr[NA NameAmount, T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], itemPtr *NA) *ResourceAmount[A] {
	if itemPtr == nil {
		return nil
	}

	resAmt := nameAmountToResourceAmount(cfg, i, *itemPtr)

	return &resAmt
}

func nameAmountToResourceAmount[NA NameAmount, T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], item NA) ResourceAmount[A] {
	var parseResp parseResponse
	switch item.GetVersion() {
	case nil:
		parseResp, _ = checkUniqueName(item.GetName(), i.objLookup)
	default:
		parseResp, _ = checkNameVersion(item.GetName(), item.GetVersion(), i.objLookup)
	}

	return ResourceAmount[A]{
		Resource: i.idToResFunc(cfg, i, parseResp.ID),
		Amount:   item.GetVal(),
	}
}

func nameAmtsToResAmts[NA NameAmount, T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], items []NA) []ResourceAmount[A] {
	results := []ResourceAmount[A]{}

	for _, item := range items {
		ra := nameAmountToResourceAmount(cfg, i, item)
		results = append(results, ra)
	}

	return results
}

func toResAmtType[NA NameAmount, RA ResourceAmountType[A], T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], old []NA, fn func(A, int32) RA) []RA {
	resAmts := []RA{}

	for _, item := range old {
		ra := nameAmountToResourceAmount(cfg, i, item)
		resAmtType := fn(ra.Resource, ra.Amount)
		resAmts = append(resAmts, resAmtType)
	}

	return resAmts
}

func getResAmtTypeMap[T ResourceAmountType[A], A APIResource](items []T) map[string]int32 {
	resAmts := resAmtTypesToStructs(items)
	return getResAmtMap(resAmts)
}

func getResAmtMap[A APIResource](items []ResourceAmount[A]) map[string]int32 {
	amountMap := make(map[string]int32)

	for _, item := range items {
		key := item.Resource.GetKey()
		amountMap[key] = item.Amount
	}

	return amountMap
}

func resAmtTypesToStructs[T ResourceAmountType[A], A APIResource](items []T) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, item := range items {
		resAmts = append(resAmts, item.ToResAmount())
	}

	return resAmts
}

func getForeignResAmts[T seeding.HasItemAmount, R any, A APIResource, L APIResourceList](i handlerInput[T, R, A, L], resources []A) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, res := range resources {
		lookup, _ := seeding.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmount := lookup.GetItemAmount()
		resAmt := newResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}

func getForeignResAmts2[T seeding.HasItemAmounts, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], resources []A, id int32) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	for _, res := range resources {
		lookup, _ := seeding.GetResourceByID(res.GetID(), i.objLookupID)
		itemAmounts := lookup.GetItemAmounts()
		var itemAmount seeding.ItemAmount

		for _, ia := range itemAmounts {
			item, _ := seeding.GetResource(ia.ItemName, cfg.l.Items)
			if item.ID == id {
				itemAmount = ia
			}
		}

		resAmt := newResourceAmount(res, itemAmount.Amount)
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}


func getForeignSliceResAmts[K seeding.LookupableID ,T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], targetItem K, skipCondition bool, fn func(*Config, K, int32) (ResourceAmount[A], error)) []ResourceAmount[A] {
	resAmts := []ResourceAmount[A]{}

	if skipCondition {
		return resAmts
	}

	for idx := range len(i.objLookup) {
		id := int32(idx + 1)
		resAmt, err := fn(cfg, targetItem, id)
		if errors.Is(err, errContinue) {
			continue
		}
		resAmts = append(resAmts, resAmt)
	}

	return resAmts
}