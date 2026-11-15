package seedhelpers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func DedupeRows[T seedtypes.Hashable](rows []T, hashes map[string]int32) []T {
	seen := make(map[string]bool)
	new := []T{}

	for _, row := range rows {
		hash := GenerateDataHash(row)

		_, ok := hashes[hash]
		if ok || seen[hash] {
			continue
		}

		seen[hash] = true
		new = append(new, row)
	}
	return new
}