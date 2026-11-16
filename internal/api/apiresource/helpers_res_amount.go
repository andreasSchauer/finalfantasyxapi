package apiresource

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
)

func GetResAmtTypeMap[T ResourceAmountType[A], A APIResource](items []T) map[string]int32 {
	resAmts := ResAmtTypesToStructs(items)
	return GetResAmtMap(resAmts)
}

func GetResAmtMap[A APIResource](items []ResourceAmount[A]) map[string]int32 {
	amountMap := make(map[string]int32, len(items))

	for _, item := range items {
		key := item.Resource.GetTestKey()
		amountMap[key] = item.Amount
	}

	return amountMap
}

func GetResAmtIDMap[A APIResource](items []ResourceAmount[A]) map[int32]int32 {
	amountMap := make(map[int32]int32, len(items))

	for _, item := range items {
		url := item.Resource.GetURL()
		id := ah.GetIDFromURL(url)
		amountMap[id] = item.Amount
	}

	return amountMap
}

func ResAmtTypesToStructs[T ResourceAmountType[A], A APIResource](items []T) []ResourceAmount[A] {
	resAmts := make([]ResourceAmount[A], len(items))

	for i, item := range items {
		resAmts[i] = item.ToResAmount()
	}

	return resAmts
}

func ResAmtMapToFloat(intMap map[string]int32) map[string]float64 {
	floatMap := make(map[string]float64, len(intMap))

	for key := range intMap {
		float := float64(intMap[key])
		floatMap[key] = float
	}

	return floatMap
}
