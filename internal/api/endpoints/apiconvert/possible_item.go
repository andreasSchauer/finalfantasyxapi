package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func PossibleItem(item seedtypes.PossibleItem) apitypes.PossibleItem {
	return th.NewPossibleItem(item.ItemAmount, item.Chance)
}
