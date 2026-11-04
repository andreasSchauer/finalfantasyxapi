package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop3SeedPrimers(qtx *database.Queries, ctx context.Context) error {
	primers, err := extractPrimers()
	if err != nil {
		return err
	}

	params := database.CreatePrimerBulkParams{
		DataHash:      make([]string, len(primers)),
		KeyItemID:     make([]int32, len(primers)),
		AlBhedLetter:  make([]string, len(primers)),
		EnglishLetter: make([]string, len(primers)),
	}

	for i, p := range primers {
		params.DataHash[i] = generateDataHash(p)
		params.KeyItemID[i] = p.KeyItemID
		params.AlBhedLetter[i] = p.AlBhedLetter
		params.EnglishLetter[i] = p.EnglishLetter
	}

	dbRows, err := qtx.CreatePrimerBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create primers: %v", err)
	}

	for i, row := range dbRows {
		primers[i].ID = row.ID
		rawJSON.primers[i].ID = row.ID
		lookup.L.Primers[seedtypes.Key(primers[i])] = primers[i]
		lookup.L.PrimersID[row.ID] = primers[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractPrimers() ([]seedtypes.Primer, error) {
	primers := []seedtypes.Primer{}
	var err error

	for i := range rawJSON.primers {
		primer := &rawJSON.primers[i]

		primer.KeyItemID, err = assignFK(primer.Name, lookup.L.KeyItems)
		if err != nil {
			return nil, err
		}

		primers = append(primers, *primer)
	}

	return dedupeRows(primers, hashes), nil
}
