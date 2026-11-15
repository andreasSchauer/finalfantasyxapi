package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedElements(qtx *database.Queries, ctx context.Context) error {
	elements := dedupeRows(seedhelpers.RawJSON.Elements, seedhelpers.Hashes)

	params := database.CreateElementBulkParams{
		DataHash: make([]string, len(elements)),
		Name:     make([]string, len(elements)),
	}

	for i, e := range elements {
		params.DataHash[i] = seedhelpers.GenerateDataHash(e)
		params.Name[i] = e.Name
	}

	dbRows, err := qtx.CreateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elements: %v", err)
	}

	for i, row := range dbRows {
		elements[i].ID = row.ID
		seedhelpers.RawJSON.Elements[i].ID = row.ID
		lookup.L.Elements[seedtypes.Key(elements[i])] = elements[i]
		lookup.L.ElementsID[row.ID] = elements[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func loop2UpdateElements(qtx *database.Queries, ctx context.Context) error {
	elements, err := extractUpdatedElements()
	if err != nil {
		return err
	}

	params := database.UpdateElementBulkParams{
		ID:                make([]int32, len(elements)),
		DataHash:          make([]string, len(elements)),
		OppositeElementID: make([]sql.NullInt32, len(elements)),
	}

	for i, e := range elements {
		params.ID[i] = e.ID
		params.DataHash[i] = seedhelpers.GenerateDataHash(e)
		params.OppositeElementID[i] = h.GetNullInt32(e.OppositeElementID)
	}

	dbRows, err := qtx.UpdateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't update elements: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractUpdatedElements() ([]seedtypes.Element, error) {
	elements := []seedtypes.Element{}
	var err error

	for i := range seedhelpers.RawJSON.Elements {
		element := &seedhelpers.RawJSON.Elements[i]

		if element.OppositeElement == nil {
			continue
		}

		delete(seedhelpers.Hashes, seedhelpers.GenerateDataHash(element))

		element.OppositeElementID, err = seedhelpers.AssignFKPtr(element.OppositeElement, lookup.L.Elements)
		if err != nil {
			return nil, err
		}

		elements = append(elements, *element)
	}

	return dedupeRows(elements, seedhelpers.Hashes), nil
}
