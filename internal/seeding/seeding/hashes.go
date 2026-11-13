package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

var hashes = make(map[string]int32)

func generateDataHash(hs seedtypes.Hashable) string {
	return h.CombineFields(hs.ToHashFields())
}

func generateJunctionHash(j Junction, desc string) string {
	return h.CombineFields(j.ToHashFieldsJ(desc))
}

func getHashID(hs seedtypes.Hashable) (int32, error) {
	id, ok := hashes[generateDataHash(hs)]
	if !ok {
		return 0, fmt.Errorf("no data hash available for %s", h.CombineFields(hs.ToHashFields()))
	}

	return id, nil
}
