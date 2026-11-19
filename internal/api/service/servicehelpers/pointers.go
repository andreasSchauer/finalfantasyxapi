package servicehelpers

import h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"


func OverrideNonePtr(ptr *string) *string {
	if ptr == nil || *ptr == "none" {
		return nil
	}

	return h.CopyPtr(ptr)
}