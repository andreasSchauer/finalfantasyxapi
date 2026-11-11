package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func BlitzballItem(bi seedtypes.BlitzballItem) apitypes.PossibleItem {
	return PossibleItem(bi.PossibleItem)
}
