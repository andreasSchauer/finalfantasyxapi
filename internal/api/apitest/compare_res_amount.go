package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareResAmounts[A apiresource.APIResource](test Test, fieldName string, exp TestResAmount[A], got apiresource.ResourceAmount[A]) {
	test.t.Helper()
	test.compIDAPIResource(fieldName+" - resource", apilookups.E.AllItems.Endpoint, exp.item, got)
	test.compare(fieldName+" - amount", exp.amount, got.Amount)
}

func comparePossibleItems(test Test, fieldName string, exp TestPossibleItem, got apitypes.PossibleItem) {
	test.t.Helper()
	compareResAmounts(test, fieldName+" - itemAmount", exp.TestResAmount, got.ResourceAmount)
	test.compare(fieldName+" - chance", exp.chance, got.Chance)
}
