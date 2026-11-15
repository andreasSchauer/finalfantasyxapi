package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedPrimers(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = sh.GenerateDataHash(p)
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
		sh.RawJSON.Primers[i].ID = row.ID
		lookup.L.Primers[seedtypes.Key(primers[i])] = primers[i]
		lookup.L.PrimersID[row.ID] = primers[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractPrimers() ([]seedtypes.Primer, error) {
	primers := []seedtypes.Primer{}
	var err error

	for i := range sh.RawJSON.Primers {
		primer := &sh.RawJSON.Primers[i]

		primer.KeyItemID, err = sh.AssignFK(primer.Name, lookup.L.KeyItems)
		if err != nil {
			return nil, err
		}

		primers = append(primers, *primer)
	}

	return sh.DedupeRows(primers, sh.Hashes), nil
}
