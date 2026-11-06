package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertBlitzballItem(bi seedtypes.BlitzballItem) apitypes.PossibleItem {
	return convertPossibleItem(bi.PossibleItem)
}
