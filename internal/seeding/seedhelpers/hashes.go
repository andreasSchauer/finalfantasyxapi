// Package seedhelpers jklösfidjfsö
package seedhelpers

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

var Hashes = make(map[string]int32)

func GenerateDataHash(hs seedtypes.Hashable) string {
	return h.CombineFields(hs.ToHashFields())
}


func GetHashID(hs seedtypes.Hashable) (int32, error) {
	id, ok := Hashes[GenerateDataHash(hs)]
	if !ok {
		return 0, fmt.Errorf("no data hash available for %s", h.CombineFields(hs.ToHashFields()))
	}

	return id, nil
}
